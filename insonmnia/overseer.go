package insonmnia

import (
	"context"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	log "github.com/noxiouz/zapctx/ctxlog"
)

const overseerTag = "sonm.overseer"
const dieEvent = "die"

// Description for a target application
// name, version, hash etc.
type Description struct {
	Registry string
	Image    string
}

// Overseer watches all miners' applications
type Overseer interface {
	Spool(ctx context.Context, d Description) error
	Spawn(ctx context.Context, d Description) error
	Close() error
}

type overseer struct {
	ctx    context.Context
	cancel context.CancelFunc

	client *client.Client

	registryAuth map[string]string

	// protects containers map
	mu         sync.Mutex
	containers map[string]*dcontainer
}

func NewOverseer(ctx context.Context) (Overseer, error) {
	ctx, cancel := context.WithCancel(ctx)
	ovr := &overseer{
		ctx:    ctx,
		cancel: cancel,
	}

	return ovr, nil
}

func (o *overseer) Close() error {
	o.cancel()
	return nil
}

func (o *overseer) handleStreamingEvents(ctx context.Context, sinceUnix int64, filterArgs filters.Args) (last int64, err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// No messages handled
	// so in the worst case restart from Since
	last = sinceUnix
	options := types.EventsOptions{
		Since:   strconv.FormatInt(sinceUnix, 10),
		Filters: filterArgs,
	}
	log.G(ctx).Info("subscribe to Docker events", zap.String("since", options.Since))
	messages, errors := o.client.Events(ctx, options)
	for {
		select {
		case err = <-errors:
			return last, err

		case message := <-messages:
			last = message.TimeNano

			switch message.Status {
			case dieEvent:
				id := message.Actor.ID
				log.G(ctx).Info("container has died", zap.String("id", id))

				var c *dcontainer
				o.mu.Lock()
				c, ok := o.containers[id]
				delete(o.containers, id)
				o.mu.Unlock()
				if ok {
					c.remove()
				} else {
					// NOTE: it could be orphaned container from our previous launch
					log.G(ctx).Warn("unknown container with sonm tag will be removed", zap.String("id", id))
					containerRemove(o.ctx, o.client, id)
				}
			default:
				log.G(ctx).Warn("received unknown event", zap.String("status", message.Status))
			}
		}
	}
}

func (o *overseer) watchEvents() {
	backoff := NewBackoffTimer(time.Second, time.Second*32)
	defer backoff.Stop()

	sinceUnix := time.Now().Unix()

	filterArgs := filters.NewArgs()
	filterArgs.Add("event", dieEvent)
	filterArgs.Add("label", overseerTag)

	var err error
	for {
		sinceUnix, err = o.handleStreamingEvents(o.ctx, sinceUnix, filterArgs)
		switch err {
		// NOTE: seems no nil-error case needed
		// case nil:
		// 	// pass
		case context.Canceled, context.DeadlineExceeded:
			log.G(o.ctx).Info("event listenening has been cancelled")
		default:
			log.G(o.ctx).Warn("failed to attach to a Docker events stream. Retry later")
			select {
			case <-backoff.C():
				//pass
			case <-o.ctx.Done():
				log.G(o.ctx).Info("event listenening has been cancelled during sleep")
				return
			}
		}
	}
}

func (o *overseer) Spool(ctx context.Context, d Description) error {
	log.G(ctx).Info("pull the application image")
	options := types.ImagePullOptions{
		All: false,
	}

	if registryAuth, ok := o.registryAuth[d.Registry]; ok {
		log.G(ctx).Info("use credentials for the registry", zap.String("registry", d.Registry))
		options.RegistryAuth = registryAuth
	}

	refStr := filepath.Join(d.Registry, d.Image)

	body, err := o.client.ImagePull(ctx, refStr, options)
	if err != nil {
		log.G(ctx).Error("ImagePull failed", zap.Error(err))
		return err
	}

	if err = decodeImagePull(body); err != nil {
		log.G(ctx).Error("failed to pull an image", zap.Error(err))
		return err
	}

	return nil
}

func (o *overseer) Spawn(ctx context.Context, d Description) error {
	pr, err := newContainer(ctx, o.client, d)
	if err != nil {
		return err
	}

	o.mu.Lock()
	o.containers[pr.ID] = pr
	o.mu.Unlock()

	if err = pr.startContainer(); err != nil {
		return err
	}
	return nil
}

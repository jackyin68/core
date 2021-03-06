// Code generated by protoc-gen-go. DO NOT EDIT.
// source: container.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Registry struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *Registry) Reset()                    { *m = Registry{} }
func (m *Registry) String() string            { return proto.CompactTextString(m) }
func (*Registry) ProtoMessage()               {}
func (*Registry) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Registry) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Registry) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// ContainerRestartPolicy represents the restart policies of the container.
type ContainerRestartPolicy struct {
	// Name can be either "always" to always restart or "on-failure" to restart
	// only when the container exit code is non-zero. If on-failure is used,
	// MaximumRetryCount controls the number of times to retry before giving up.
	// The default is not to restart.
	Name              string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	MaximumRetryCount uint32 `protobuf:"varint,2,opt,name=maximumRetryCount" json:"maximumRetryCount,omitempty"`
}

func (m *ContainerRestartPolicy) Reset()                    { *m = ContainerRestartPolicy{} }
func (m *ContainerRestartPolicy) String() string            { return proto.CompactTextString(m) }
func (*ContainerRestartPolicy) ProtoMessage()               {}
func (*ContainerRestartPolicy) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *ContainerRestartPolicy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerRestartPolicy) GetMaximumRetryCount() uint32 {
	if m != nil {
		return m.MaximumRetryCount
	}
	return 0
}

type NetworkSpec struct {
	Type    string            `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Options map[string]string `protobuf:"bytes,2,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Subnet  string            `protobuf:"bytes,3,opt,name=subnet" json:"subnet,omitempty"`
	Addr    string            `protobuf:"bytes,4,opt,name=addr" json:"addr,omitempty"`
}

func (m *NetworkSpec) Reset()                    { *m = NetworkSpec{} }
func (m *NetworkSpec) String() string            { return proto.CompactTextString(m) }
func (*NetworkSpec) ProtoMessage()               {}
func (*NetworkSpec) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *NetworkSpec) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NetworkSpec) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *NetworkSpec) GetSubnet() string {
	if m != nil {
		return m.Subnet
	}
	return ""
}

func (m *NetworkSpec) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type Container struct {
	// Image describes a Docker image name. Required.
	Image string `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	// SSH public key used to attach to the container.
	SshKey string `protobuf:"bytes,2,opt,name=sshKey" json:"sshKey,omitempty"`
	// CommitOnStop points whether a container should commit when stopped.
	// Committed containers can be fetched later while there is an active
	// deal.
	CommitOnStop bool `protobuf:"varint,3,opt,name=commitOnStop" json:"commitOnStop,omitempty"`
	// Env describes environment variables forwarded into the container.
	Env map[string]string `protobuf:"bytes,4,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Volumes describes network volumes that are used to be mounted inside
	// the container.
	// Mapping from the volume type (cifs, nfs, etc.) to its settings.
	Volumes map[string]*Volume `protobuf:"bytes,5,rep,name=volumes" json:"volumes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Mounts describes mount points from the volume name to the container.
	Mounts   []string       `protobuf:"bytes,6,rep,name=mounts" json:"mounts,omitempty"`
	Networks []*NetworkSpec `protobuf:"bytes,7,rep,name=networks" json:"networks,omitempty"`
	// ContainerRestartPolicy describes the restart policies of the container.
	RestartPolicy *ContainerRestartPolicy `protobuf:"bytes,8,opt,name=restartPolicy" json:"restartPolicy,omitempty"`
	// Expose controls how container ports are exposed.
	// Format is "public_ip:public_port:private_port/protocol"
	// Protocol can be "tcp", "udp", "sctp".
	// If the "protocol" parameter is ommited "tcp" is implied.
	// If the "public_ip" parameter is ommited then the port is being exposed on all available ips.
	Expose []string `protobuf:"bytes,10,rep,name=expose" json:"expose,omitempty"`
	// Push the committed image to remote repository (works only if CommitOnStop is set to `true`).
	PushOnStop bool `protobuf:"varint,11,opt,name=pushOnStop" json:"pushOnStop,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *Container) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Container) GetSshKey() string {
	if m != nil {
		return m.SshKey
	}
	return ""
}

func (m *Container) GetCommitOnStop() bool {
	if m != nil {
		return m.CommitOnStop
	}
	return false
}

func (m *Container) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Container) GetVolumes() map[string]*Volume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *Container) GetMounts() []string {
	if m != nil {
		return m.Mounts
	}
	return nil
}

func (m *Container) GetNetworks() []*NetworkSpec {
	if m != nil {
		return m.Networks
	}
	return nil
}

func (m *Container) GetRestartPolicy() *ContainerRestartPolicy {
	if m != nil {
		return m.RestartPolicy
	}
	return nil
}

func (m *Container) GetExpose() []string {
	if m != nil {
		return m.Expose
	}
	return nil
}

func (m *Container) GetPushOnStop() bool {
	if m != nil {
		return m.PushOnStop
	}
	return false
}

func init() {
	proto.RegisterType((*Registry)(nil), "sonm.Registry")
	proto.RegisterType((*ContainerRestartPolicy)(nil), "sonm.ContainerRestartPolicy")
	proto.RegisterType((*NetworkSpec)(nil), "sonm.NetworkSpec")
	proto.RegisterType((*Container)(nil), "sonm.Container")
}

func init() { proto.RegisterFile("container.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xd1, 0x6a, 0xdb, 0x30,
	0x14, 0xc5, 0xb1, 0x9b, 0xb8, 0x37, 0x2e, 0x5b, 0xc5, 0x28, 0xc2, 0x8c, 0x12, 0xcc, 0x1e, 0xca,
	0x58, 0x1d, 0xe8, 0xa0, 0x94, 0xee, 0x2d, 0xa5, 0x30, 0x18, 0xac, 0x43, 0x85, 0x3d, 0xf4, 0xcd,
	0x71, 0x44, 0x22, 0x1a, 0x49, 0x46, 0x92, 0xd3, 0xfa, 0x07, 0xf6, 0x63, 0xfb, 0xb1, 0x21, 0xc9,
	0x36, 0xce, 0xda, 0x97, 0xbe, 0xdd, 0x23, 0x9d, 0x7b, 0xee, 0xb9, 0xc7, 0x32, 0xbc, 0x2b, 0xa5,
	0x30, 0x05, 0x13, 0x54, 0xe5, 0x95, 0x92, 0x46, 0xa2, 0x48, 0x4b, 0xc1, 0xd3, 0x64, 0x27, 0xb7,
	0x35, 0xa7, 0xfe, 0x2c, 0x5b, 0x40, 0x4c, 0xe8, 0x9a, 0x69, 0xa3, 0x1a, 0x94, 0x42, 0x5c, 0x6b,
	0xaa, 0x44, 0xc1, 0x29, 0x0e, 0x66, 0xc1, 0xd9, 0x21, 0xe9, 0xb1, 0xbd, 0xab, 0x0a, 0xad, 0x9f,
	0xa4, 0x5a, 0xe1, 0x91, 0xbf, 0xeb, 0x70, 0xf6, 0x00, 0x27, 0x37, 0xdd, 0x28, 0x42, 0xb5, 0x29,
	0x94, 0xf9, 0x25, 0xb7, 0xac, 0x6c, 0x10, 0x82, 0x68, 0xa0, 0xe6, 0x6a, 0xf4, 0x05, 0x8e, 0x79,
	0xf1, 0xcc, 0x78, 0xcd, 0x09, 0x35, 0xaa, 0xb9, 0x91, 0xb5, 0x30, 0x4e, 0xf2, 0x88, 0xbc, 0xbc,
	0xc8, 0xfe, 0x06, 0x30, 0xfd, 0x49, 0xcd, 0x93, 0x54, 0x8f, 0xf7, 0x15, 0x2d, 0xad, 0xa2, 0x69,
	0xaa, 0x5e, 0xd1, 0xd6, 0xe8, 0x0a, 0x26, 0xb2, 0x32, 0x4c, 0x0a, 0x8d, 0x47, 0xb3, 0xf0, 0x6c,
	0x7a, 0x71, 0x9a, 0xdb, 0x4d, 0xf3, 0x41, 0x5f, 0x7e, 0xe7, 0x09, 0xb7, 0xc2, 0xa8, 0x86, 0x74,
	0x74, 0x74, 0x02, 0x63, 0x5d, 0x2f, 0x05, 0x35, 0x38, 0x74, 0x7a, 0x2d, 0xb2, 0x53, 0x8a, 0xd5,
	0x4a, 0xe1, 0xc8, 0x4f, 0xb1, 0x75, 0x7a, 0x0d, 0xc9, 0x50, 0x04, 0xbd, 0x87, 0xf0, 0x91, 0x36,
	0xad, 0x11, 0x5b, 0xa2, 0x0f, 0x70, 0xb0, 0x2b, 0xb6, 0x35, 0x6d, 0x03, 0xf2, 0xe0, 0x7a, 0x74,
	0x15, 0x64, 0x7f, 0x22, 0x38, 0xec, 0x23, 0xb2, 0x3c, 0xc6, 0x8b, 0x75, 0xb7, 0x84, 0x07, 0xce,
	0x8b, 0xde, 0xfc, 0xa0, 0x4d, 0xdb, 0xde, 0x22, 0x94, 0x41, 0x52, 0x4a, 0xce, 0x99, 0xb9, 0x13,
	0xf7, 0x46, 0x56, 0xce, 0x69, 0x4c, 0xf6, 0xce, 0xd0, 0x67, 0x08, 0xa9, 0xd8, 0xe1, 0xc8, 0x6d,
	0x8f, 0xfd, 0xf6, 0xfd, 0xbc, 0xfc, 0x56, 0xec, 0xfc, 0xde, 0x96, 0x84, 0x2e, 0x61, 0xe2, 0x5f,
	0x80, 0xc6, 0x07, 0x8e, 0xff, 0xf1, 0x7f, 0xfe, 0x6f, 0x7f, 0xdd, 0x66, 0xd5, 0x92, 0xad, 0x3f,
	0x6e, 0x3f, 0x89, 0xc6, 0xe3, 0x59, 0x68, 0xfd, 0x79, 0x84, 0xce, 0x21, 0x16, 0x3e, 0x68, 0x8d,
	0x27, 0x4e, 0xf0, 0xf8, 0x45, 0xfc, 0xa4, 0xa7, 0xa0, 0x05, 0x1c, 0xa9, 0xe1, 0x1b, 0xc1, 0xf1,
	0x2c, 0x78, 0xc5, 0xc4, 0xde, 0x3b, 0x22, 0xfb, 0x2d, 0xd6, 0x0a, 0x7d, 0xae, 0xa4, 0xa6, 0x18,
	0xbc, 0x15, 0x8f, 0xd0, 0x29, 0x40, 0x55, 0xeb, 0x4d, 0x1b, 0xd4, 0xd4, 0x05, 0x35, 0x38, 0x49,
	0x2f, 0x21, 0xee, 0xb2, 0x78, 0xcb, 0xe7, 0x4b, 0xbf, 0x43, 0x32, 0xcc, 0xe4, 0x95, 0xde, 0x6c,
	0xd8, 0x3b, 0xbd, 0x48, 0xfc, 0x36, 0xbe, 0x69, 0xa0, 0xb4, 0xf8, 0xf4, 0x90, 0xad, 0x99, 0xd9,
	0xd4, 0xcb, 0xbc, 0x94, 0x7c, 0x6e, 0x49, 0xe7, 0x4c, 0xce, 0x4b, 0xa9, 0xe8, 0xdc, 0xfd, 0x8f,
	0xdf, 0xec, 0xd1, 0x72, 0xec, 0xea, 0xaf, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x15, 0xa4, 0xe7,
	0xcf, 0xc2, 0x03, 0x00, 0x00,
}

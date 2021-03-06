module github.com/sonm-io/core

require (
	bazil.org/fuse v0.0.0-20160811212531-371fbbdaa898
	github.com/Azure/go-autorest v8.1.1+incompatible
	github.com/BurntSushi/toml v0.0.0-20170626110600-a368813c5e64
	github.com/ContainX/docker-volume-netshare v0.0.0-20180128064639-3b432b74a4be
	github.com/Masterminds/squirrel v0.0.0-20180612165926-7d201d09bfa8
	github.com/MaxHalford/gago v0.0.0-20180803141938-502b393f59df
	github.com/NVIDIA/nvidia-docker v0.0.0-20170722195403-d8e8727788aa
	github.com/Sirupsen/logrus v1.1.0
	github.com/StackExchange/wmi v0.0.0-20170410192909-ea383cf3ba6e
	github.com/anmitsu/go-shlex v0.0.0-20161002113705-648efa622239
	github.com/aristanetworks/goarista v0.0.0-20180907105523-ff33da284e76
	github.com/armon/go-metrics v0.0.0-20180221182744-783273d70314
	github.com/asaskevich/govalidator v0.0.0-20180319081651-7d2e70ef918f
	github.com/beorn7/perks v0.0.0-20160804104726-4c0e84591b9a
	github.com/bifurcation/mint v0.0.0-20181105071958-a14404e9a861 // indirect
	github.com/blang/semver v3.5.1+incompatible
	github.com/boltdb/bolt v1.3.1
	github.com/btcsuite/btcd v0.0.0-20171023093315-c7588cbf7690
	github.com/c2h5oh/datasize v0.0.0-20171227191756-4eba002a5eae
	github.com/ccding/go-stun v0.0.0-20170323223013-04a4eed61c57
	github.com/cdipaolo/goml v0.0.0-20161030204843-e2d8def04c9a
	github.com/certifi/gocertifi v0.0.0-20180905225744-ee1a9a0726d2
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/cnf/structhash v0.0.0-20170702194520-7710f1f78fb9
	github.com/containerd/cgroups v0.0.0-20170714210333-6d5c608c203d
	github.com/coreos/go-systemd v0.0.0-20170609144627-24036eb3df68
	github.com/davecgh/go-spew v1.1.1
	github.com/deckarep/golang-set v0.0.0-20180927150649-699df6a3acf6
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dickeyxxx/netrc v0.0.0-20150924214217-3acf1b3de25d
	github.com/docker/distribution v2.7.0-rc.0+incompatible
	github.com/docker/docker v0.0.0-20170608045038-cd35e4beee13
	github.com/docker/go-connections v0.4.0
	github.com/docker/go-plugins-helpers v0.0.0-20180116160015-61cb8e233420
	github.com/docker/go-units v0.3.3
	github.com/docker/libkv v0.0.0-20170701232644-93ab0e6c056d
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7
	github.com/edsrzf/mmap-go v0.0.0-20170320065105-0bce6a688712
	github.com/elastic/gosigar v0.0.0-20180330100440-37f05ff46ffa
	github.com/ethereum/go-ethereum v0.0.0-20180929205331-b69942befeb9
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052
	github.com/fatih/structs v0.0.0-20180123065059-ebf56d35bba7
	github.com/fjl/memsize v0.0.0-20180929194037-2a09253e352a
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gliderlabs/ssh v0.0.0-20170731185311-9ccc7bbb6481
	github.com/go-ole/go-ole v0.0.0-20170712174622-085abb85892d
	github.com/go-playground/locales v0.12.1
	github.com/go-playground/universal-translator v0.0.0-20170327191703-71201497bace
	github.com/go-stack/stack v1.8.0
	github.com/godbus/dbus v0.0.0-20170707174628-bd29ed602e2c
	github.com/gogo/protobuf v0.0.0-20171213104750-35b81a066e52
	github.com/golang/mock v1.0.0
	github.com/golang/protobuf v0.0.0-20180328163153-e09c5db29600
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db
	github.com/gosuri/uilive v0.0.0-20170323041506-ac356e6e42cd
	github.com/gosuri/uiprogress v0.0.0-20170224063937-d0567a9d84a1
	github.com/grpc-ecosystem/go-grpc-middleware v0.0.0-20171215095238-967bee733a73
	github.com/grpc-ecosystem/go-grpc-prometheus v0.0.0-20170826090648-0dafe0d496ea
	github.com/gxed/GoEndian v0.0.0-20160916112711-0f5c6873267e
	github.com/gxed/eventfd v0.0.0-20160916113412-80a92cca79a8
	github.com/hashicorp/errwrap v0.0.0-20141028054710-7554cd9344ce
	github.com/hashicorp/go-immutable-radix v0.0.0-20180129170900-7f3cd4390caa
	github.com/hashicorp/go-msgpack v0.0.0-20150518234257-fa3f63826f7c
	github.com/hashicorp/go-multierror v0.0.0-20171204182908-b7773ae21874
	github.com/hashicorp/go-sockaddr v0.0.0-20180320115054-6d291a969b86
	github.com/hashicorp/golang-lru v0.5.0
	github.com/hashicorp/hcl v0.0.0-20180404174102-ef8a98b0bbce
	github.com/hashicorp/memberlist v0.0.0-20180209033901-2288bf30e9c8
	github.com/howeyc/gopass v0.0.0-20170109162249-bf9dde6d0d2c
	github.com/huin/goupnp v0.0.0-20180415215157-1395d1447324
	github.com/inconshreveable/mousetrap v1.0.0
	github.com/influxdata/influxdb v0.0.0-20180412224233-7ebfc9c544e0
	github.com/ipfs/go-log v0.0.0-20180131202911-0e2a17b81af4
	github.com/jackpal/go-nat-pmp v0.0.0-20170405195558-28a68d0c24ad
	github.com/jinzhu/configor v0.0.0-20171024081003-6ecfe629230f
	github.com/json-iterator/go v0.0.0-20180424004623-2ddf6d758266
	github.com/karalabe/hid v0.0.0-20180420081245-2b4488a37358
	github.com/kr/pty v1.1.3
	github.com/lann/builder v0.0.0-20180216234317-1b87b36280d0
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0
	github.com/lib/pq v0.0.0-20180327071824-d34b9ff171c2
	github.com/libp2p/go-reuseport v0.1.10
	github.com/libp2p/go-sockaddr v0.0.0-20180128062301-09ae606455f8
	github.com/lucas-clemente/aes12 v0.0.0-20171027163421-cd47fb39b79f // indirect
	github.com/lucas-clemente/quic-go v0.10.0
	github.com/lucas-clemente/quic-go-certificates v0.0.0-20160823095156-d2f86524cced // indirect
	github.com/magiconair/properties v0.0.0-20180217134545-2c9e95027885
	github.com/mattn/go-colorable v0.0.0-20180310133214-efa589957cd0
	github.com/mattn/go-isatty v0.0.4
	github.com/mattn/go-runewidth v0.0.0-20170510074858-97311d9f7767
	github.com/matttproud/golang_protobuf_extensions v1.0.1
	github.com/miekg/dns v1.0.4
	github.com/mitchellh/go-homedir v0.0.0-20180523094522-3864e76763d9
	github.com/mitchellh/mapstructure v0.0.0-20170523030023-d0303fe80992
	github.com/mmcloughlin/geohash v0.0.0-20180909114810-59020f29e94a
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 v0.0.0-20180320133207-05fbef0ca5da
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe
	github.com/noxiouz/zapctx v0.0.0-20170511100814-b19b4472cfa5
	github.com/ogier/pflag v0.0.0-20160129220114-45c278ab3607
	github.com/olekukonko/tablewriter v0.0.0-20180506121414-d4647c9c7a84
	github.com/opencontainers/go-digest v1.0.0-rc1
	github.com/opencontainers/image-spec v1.0.1
	github.com/opencontainers/runtime-spec v0.0.0-20170728172216-6d696f524b25
	github.com/opentracing/basictracer-go v0.0.0-20171205173151-7be4bbce182d
	github.com/opentracing/opentracing-go v0.0.0-20171003133519-1361b9cd60be
	github.com/oschwald/geoip2-golang v1.2.1
	github.com/oschwald/maxminddb-golang v1.3.0
	github.com/pborman/uuid v0.0.0-20160216163710-c55201b03606
	github.com/pelletier/go-toml v0.0.0-20180323185243-66540cf1fcd2
	github.com/pkg/errors v0.8.0
	github.com/pmezard/go-difflib v1.0.0
	github.com/pquerna/ffjson v0.0.0-20171002144729-d49c2bc1aa13
	github.com/prometheus/client_golang v0.0.0-20180120141031-06bc6e01f4ba
	github.com/prometheus/client_model v0.0.0-20171117100541-99fa1f4be8e5
	github.com/prometheus/common v0.0.0-20180110214958-89604d197083
	github.com/prometheus/procfs v0.0.0-20171226183907-b15cd069a834
	github.com/rcrowley/go-metrics v0.0.0-20180503174638-e2704e165165
	github.com/rjeczalik/notify v0.9.2
	github.com/robertkrimen/otto v0.0.0-20180617131154-15f95af6e78d
	github.com/rs/cors v0.0.0-20180826180256-dc7332ab32be
	github.com/satori/uuid v0.0.0-20170321230731-5bf94b69c6b6
	github.com/sean-/seed v0.0.0-20170313163322-e2103e2c3529
	github.com/serialx/hashring v0.0.0-20170811022404-6a9381c5a83e
	github.com/sevlyar/retag v0.0.0-20171212084444-de3f0096c81e
	github.com/shirou/gopsutil v0.0.0-20170804030934-3aa2ffab12a1
	github.com/spf13/afero v1.1.0
	github.com/spf13/cast v1.2.0
	github.com/spf13/cobra v0.0.0-20170731170427-b26b538f6930
	github.com/spf13/jwalterweatherman v0.0.0-20180109140146-7c0cea34c8ec
	github.com/spf13/pflag v1.0.0
	github.com/spf13/viper v0.0.0-20180404183325-8dc2790b029d
	github.com/sshaman1101/nvidia-docker v0.0.3
	github.com/stretchr/testify v1.2.2
	github.com/syndtr/goleveldb v0.0.0-20180815032940-ae2bd5eed72d
	github.com/tcnksm/go-input v0.0.0-20180404061846-548a7d7a8ee8
	github.com/tehnerd/gnl2go v0.0.0-20161218223753-101b5c6e2d44
	github.com/uber-go/atomic v1.3.1
	github.com/veandco/go-sdl2 v0.0.0-20180925095440-75ff82abc4e3
	github.com/vishvananda/netlink v0.0.0-20180515155430-7c0b5944a303
	github.com/vishvananda/netns v0.0.0-20171111001504-be1fbeda1936
	github.com/whyrusleeping/go-logging v0.0.0-20170515211332-0457bb6b88fc
	github.com/yandex/pandora v0.1.3
	go.uber.org/atomic v0.0.0-20170719224650-70bd1261d36b
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.9.1
	golang.org/x/crypto v0.0.0-20180927165925-5295e8364332
	golang.org/x/net v0.0.0-20180926154720-4dfa2610cdf3
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f
	golang.org/x/sys v0.0.0-20180928133829-e4b3c5e90611
	golang.org/x/text v0.0.0-20180911161511-905a57155faa
	golang.org/x/tools v0.0.0-20180928181343-b3c0be4c978b
	google.golang.org/genproto v0.0.0-20180206005123-2b5a72b8730b
	google.golang.org/grpc v0.0.0-20180801224056-b20cbb449d97
	gopkg.in/bluesuncorp/validator.v9 v9.15.0
	gopkg.in/go-playground/validator.v8 v8.18.2
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce
	gopkg.in/olebedev/go-duktape.v3 v3.0.0-20180302121509-abf0ba0be5d5
	gopkg.in/oleiade/lane.v1 v1.0.0
	gopkg.in/sourcemap.v1 v1.0.5
	gopkg.in/urfave/cli.v1 v1.20.0
	gopkg.in/yaml.v1 v1.0.0-20140924161607-9f9df34309c0
	gopkg.in/yaml.v2 v2.2.1
)

replace github.com/Sirupsen/logrus v1.1.0 => github.com/sirupsen/logrus v1.1.0

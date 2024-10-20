module github.com/amarps/joebot

go 1.23.2

replace github.com/filebrowser/filebrowser/v2 v2.31.2 => ./external/filebrowser

require (
	github.com/amarps/gotty v0.0.0-20241015150555-406cf3cdb0ed
	github.com/asdine/storm/v3 v3.2.1
	github.com/filebrowser/filebrowser/v2 v2.31.2
	github.com/ginuerzh/gost v2.6.1+incompatible
	github.com/hashicorp/yamux v0.1.2
	github.com/labstack/echo v3.3.10+incompatible
	github.com/pkg/errors v0.9.1
	github.com/pkg/sftp v1.13.6
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/afero v1.11.0
	github.com/spf13/cobra v1.8.1
	github.com/twinj/uuid v1.0.0
	golang.org/x/crypto v0.28.0
	golang.org/x/net v0.30.0
)

require (
	git.torproject.org/pluggable-transports/goptlib.git v1.3.0 // indirect
	git.torproject.org/pluggable-transports/obfs4.git v0.0.0-20181103133120-08f4d470188e // indirect
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/aead/chacha20 v0.0.0-20180709150244-8b13a72661da // indirect
	github.com/agl/ed25519 v0.0.0-20170116200512-5312a6153412 // indirect
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/asticode/go-astikit v0.42.0 // indirect
	github.com/asticode/go-astisub v0.26.2 // indirect
	github.com/asticode/go-astits v1.13.0 // indirect
	github.com/bifurcation/mint v0.0.0-20181105071958-a14404e9a861 // indirect
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.5 // indirect
	github.com/creack/pty v1.1.9 // indirect
	github.com/dchest/siphash v1.2.2 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/dsnet/compress v0.0.2-0.20210315054119-f66993602bf5 // indirect
	github.com/dsoprea/go-exif/v3 v3.0.1 // indirect
	github.com/dsoprea/go-logging v0.0.0-20200710184922-b02d349568dd // indirect
	github.com/dsoprea/go-utility/v2 v2.0.0-20221003172846-a3e1774ef349 // indirect
	github.com/elazarl/go-bindata-assetfs v1.0.1 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/flynn/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/ginuerzh/gosocks4 v0.0.1 // indirect
	github.com/ginuerzh/gosocks5 v0.2.0 // indirect
	github.com/ginuerzh/tls-dissector v0.0.2-0.20200224064855-24ab2b3a3796 // indirect
	github.com/go-errors/errors v1.5.1 // indirect
	github.com/go-log/log v0.2.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/geo v0.0.0-20230421003525-6adc56603217 // indirect
	github.com/golang/mock v1.1.1 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/klauspost/cpuid/v2 v2.2.6 // indirect
	github.com/klauspost/crc32 v1.2.0 // indirect
	github.com/klauspost/pgzip v1.2.6 // indirect
	github.com/klauspost/reedsolomon v1.12.0 // indirect
	github.com/kr/fs v0.1.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/pty v1.1.8 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/lucas-clemente/aes12 v0.0.0-20171027163421-cd47fb39b79f // indirect
	github.com/lucas-clemente/quic-go v0.10.0 // indirect
	github.com/lucas-clemente/quic-go-certificates v0.0.0-20160823095156-d2f86524cced // indirect
	github.com/maruel/natural v1.1.1 // indirect
	github.com/marusama/semaphore/v2 v2.5.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mholt/archiver/v3 v3.5.1 // indirect
	github.com/miekg/dns v1.1.58 // indirect
	github.com/myesui/uuid v1.0.0 // indirect
	github.com/nwaples/rardecode v1.1.3 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.34.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/shadowsocks/shadowsocks-go v0.0.0-20200409064450-3e585ff90601 // indirect
	github.com/shirou/gopsutil/v3 v3.24.5 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tomasen/realip v0.0.0-20180522021738-f0c99a92ddce // indirect
	github.com/ulikunitz/xz v0.5.11 // indirect
	github.com/urfave/cli v1.22.16 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	github.com/yudai/hcl v0.0.0-20151013225006-5fa2393b3552 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.etcd.io/bbolt v1.3.11 // indirect
	golang.org/x/image v0.19.0 // indirect
	golang.org/x/mod v0.20.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	golang.org/x/tools v0.24.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	gopkg.in/gorilla/websocket.v1 v1.4.0 // indirect
	gopkg.in/stretchr/testify.v1 v1.2.2 // indirect
	gopkg.in/xtaci/kcp-go.v2 v2.0.3 // indirect
	gopkg.in/xtaci/smux.v1 v1.4.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

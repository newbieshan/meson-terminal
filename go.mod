module github.com/daqnext/meson-terminal

go 1.15

require (
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/daqnext/meson-common v1.0.26
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-contrib/gzip v0.0.3
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/imroc/req v0.3.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/shirou/gopsutil/v3 v3.21.3
	github.com/sirupsen/logrus v1.7.0
	github.com/syndtr/goleveldb v1.0.1-0.20200815110645-5c35d600f0ca
	github.com/takama/daemon v1.0.0
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f
)

//replace github.com/daqnext/meson-common => /Users/zhangzhenbo/workspace/go/project/meson-common

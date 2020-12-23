package statemgr

import (
	"encoding/json"
	"fmt"
	"github.com/daqnext/meson-common/common/accountmgr"
	"github.com/daqnext/meson-common/common/commonmsg"
	"github.com/daqnext/meson-common/common/httputils"
	"github.com/daqnext/meson-common/common/logger"
	"github.com/daqnext/meson-common/common/resp"
	"github.com/daqnext/meson-common/common/utils"
	"github.com/daqnext/meson-terminal/terminal/manager/config"
	"github.com/daqnext/meson-terminal/terminal/manager/filemgr"
	"github.com/daqnext/meson-terminal/terminal/manager/global"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

var State = &commonmsg.TerminalStatesMsg{}

func GetMachineState() (*commonmsg.TerminalStatesMsg, error) {
	if State.OS == "" {
		if h, err := host.Info(); err == nil {
			State.OS = fmt.Sprintf("%v:%v(%v):%v", h.OS, h.Platform, h.PlatformFamily, h.PlatformVersion)
		}
	}

	if State.CPU == "" {
		if c, err := cpu.Info(); err == nil {
			State.CPU = c[0].ModelName
		}
	}

	if v, err := mem.VirtualMemory(); err == nil {
		State.MemTotal = v.Total
		State.MemAvailable = v.Available
	}

	if d, err := disk.Usage("./"); err == nil {
		State.DiskTotal = d.Total
		State.DiskAvailable = d.Free
	}

	State.CdnDiskTotal = uint64(filemgr.CdnSpaceLimit)
	State.CdnDiskAvailable = State.CdnDiskTotal - uint64(filemgr.CdnSpaceUsed)

	if State.MacAddr == "" {
		if macAddr, err := utils.GetMainMacAddress(); err != nil {
			logger.Error("failed to get mac address", "err", err)
		} else {
			State.MacAddr = macAddr
		}
	}

	if State.Port == "" {
		State.Port = config.UsingPort
	}

	State.Version = global.Version

	return State, nil
}

func SendStateToServer() {
	machineState, err := GetMachineState()
	if err != nil {
		return
	}
	header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + accountmgr.Token,
	}

	//提交请求
	content, err := httputils.Request("POST", global.SendHeartBeatUrl, machineState, header)
	if err != nil {
		logger.Error("send terminalState to server error", "err", err)
		return
	}
	//logger.Debug("response form server", "response string", string(content))
	var respBody resp.RespBody
	if err := json.Unmarshal(content, &respBody); err != nil {
		logger.Error("response from terminal unmarshal error", "err", err)
		return
	}

	switch respBody.Status {
	case 0:
		//logger.Debug("send State success")
	case 101: //auth error
		logger.Fatal("auth error,please restart terminal with correct username and password")
	case 106: //low version
		logger.Fatal("Your version need upgrade. Please download new version from meson.network ")
	default:
		logger.Error("server error")
	}
}

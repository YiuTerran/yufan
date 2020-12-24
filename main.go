package main

import (
	"fmt"
	"net"
	"strings"
	"time"

	"yufan/lib"

	"github.com/imroc/req"
)

type envInfo struct {
	Name string
	url  string
}

const (
	heartbeatSuffix = "/iot-router/face/yufan/heartbeat"
	identifySuffix  = "/iot-router/face/yufan/identify"
)

var (
	envs = []*envInfo{
		{"建造汇", "https://www.whjzh.cn"},
		{"仙桃建管通", "http://119.36.247.108"},
		{"车都集团", "http://58.49.51.38"},
	}
)

func main() {
	req.SetTimeout(5 * time.Second)
	var idx int
	for {
		fmt.Println("请选择需要对接的环境：")
		for idx, info := range envs {
			fmt.Printf("[%d]%s: %s\n", idx, info.Name, info.url)
		}
		_, err := fmt.Scanln(&idx)
		if err != nil || idx < 0 || idx >= len(envs) {
			fmt.Println("输入错误，请输入选项前括号内的数字")
		} else {
			break
		}
	}
	prefix := envs[idx].url
	var targets []string
	fmt.Println("请输入考勤机ip地址，每行输入一个，直接回车结束输入：")
	for {
		var ipstr string
		_, _ = fmt.Scanln(&ipstr)
		ipstr = strings.TrimSpace(ipstr)
		if ipstr == "" {
			break
		}
		addr := net.ParseIP(ipstr)
		if addr == nil {
			fmt.Println("ip地址格式错误，请重新输入")
			continue
		}
		targets = append(targets, ipstr)
	}
	if len(targets) == 0 {
		fmt.Println("无有效输入，程序结束")
		return
	}
	var pass string
	fmt.Println("请输入设备密码(直接回车，则默认为755626):")
	_, _ = fmt.Scanln(&pass)
	if pass == "" {
		pass = lib.DefaultPass
	}
	fmt.Println("开始设置回调地址...")
	for _, ip := range targets {
		c := &lib.Device{
			Pass: pass,
			IP:   ip,
		}
		err := c.SetHeartBeatURL(prefix + heartbeatSuffix)
		if err != nil {
			fmt.Printf("%s设置回调失败，请检查输入以及网络情况\n", ip)
			continue
		}
		err = c.SetIdentifyCallBack(prefix + identifySuffix)
		if err != nil {
			fmt.Printf("%s设置回调失败，请检查输入以及网络情况\n", ip)
		}
	}
	fmt.Println("设置完成，按回车键退出")
	var exit string
	_, _ = fmt.Scanln(&exit)
}

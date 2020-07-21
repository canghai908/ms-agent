package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/canghai908/ms-agent/logging"
	"github.com/canghai908/ms-agent/setting"
)

//AppVersion version
var (
	vers *bool
	help *bool

	version   = "No Version Provided"
	gitHash   = "No GitHash Provided"
	buildTime = "No BuildTime Provided"
)

func init() {
	vers = flag.Bool("v", false, "display the version.")
	help = flag.Bool("h", false, "print this help.")
	flag.Parse()

	if *vers {
		fmt.Println("Version:", version)
		fmt.Println("Git Commit Hash:", gitHash)
		fmt.Println("UTC Build Time:", buildTime)
		os.Exit(0)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	setting.Setup()
	logging.Setup()
}

func main() {
	if len(os.Args) < 3 {
		logging.Error("参数不正确")
		return
	}
	if setting.AppSetting.Debug == "0" {
		logging.Debug("os.Args[1]:", os.Args[1])
		logging.Debug("os.Args[2]:", os.Args[2])
		logging.Debug("os.Args[3]:", os.Args[3])
	}
	post := []byte(os.Args[3])
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	addr := setting.AppSetting.Host
	req, err := http.NewRequest("POST", addr, bytes.NewReader(post))
	if err != nil {
		logging.Error(err)
		return
	}
	if setting.AppSetting.Debug == "0" {
		logging.Debug("Send --->>>", string(post))
	}
	token := setting.AppSetting.Token
	req.Header.Add("Token", token)
	req.Header.Add("User-Agent", "Zabbix-Client")
	resp, err := client.Do(req)
	if err != nil {
		logging.Error(err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Error(err)
		return
	}
	if setting.AppSetting.Debug == "0" {
		logging.Debug("Receive <<<---", string(b))
	}
	logging.Info("send ok!")
}

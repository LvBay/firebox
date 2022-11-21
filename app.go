package main

import (
	"context"
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os/exec"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetMouseClick() int {
	return 1
}

func (a *App) GreetAsyncViaEvent() {
	count := 0
	go func() {
		ch := robotgo.EventStart()
		for event := range ch {
			if event.Kind == hook.MouseUp {
				count++
				runtime.EventsEmit(a.ctx, "rcv:greet", count)
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

func (a *App) OpenFileDialog() error {
	//p, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{Title: "选取lol路径"})
	//cmd := exec.Command("cmd", "/c", "start", p)
	////  D:\lol\LOL(26)\TCLS\Client.exe
	//err = cmd.Run()
	//fmt.Println(p, err)
	cmd := exec.Command("cmd", "/c", "wmic PROCESS WHERE name='LeagueClientUx.exe'")
	out, err := cmd.Output()
	fmt.Println(string(out), err)
	return err
}

//  CreationClassName  CreationDate               CSCreationClassName   CSName           Description         ExecutablePath                                  ExecutionState  Handle  HandleCount  InstallDate  KernelModeTime  MaximumWorkingSetSize  MinimumWorkingSetSize  Name                OSCreationClassName    OSName                                                               OtherOperationCount  OtherTransferCount  PageFaults  PageFileUsage  ParentProcessId  PeakPageFileUsage  PeakVirtualSize  PeakWorkingSetSize  Priority  PrivatePageCount  ProcessId  QuotaNonPagedPoolUsage  QuotaPagedPoolUsage  QuotaPeakNonPagedPoolUsage  QuotaPeakPagedPoolUsage  ReadOperationCount  ReadTransferCount  SessionId  Status  TerminationDate  ThreadCount  UserModeTime  VirtualSize  WindowsVersion  WorkingSetSize  WriteOperationCount  WriteTransferCount
//LeagueClientUx.exe  d:/lol/lol(26)/LeagueClient/LeagueClientUx.exe "--riotclient-auth-token=6L79Q-wRsQCLf6wsXvbARw" "--riotclient-app-port=63927" "--riotclient-tencent" "--no-rads" "--disable-self-update" "--region=TENCENT" "--locale=zh_CN" "--t.lcdshost=hn1-sz-feapp.lol.qq.com" "--t.chathost=hn1-sz-ejabberd.lol.qq.com" "--t.lq=https://hn1-sz-login.lol.qq.com:8443" "--t.storeurl=https://hn1-sr.lol.qq.com:8443" "--t.rmsurl=wss://sz-rms-bcs.lol.qq.com:443" "--rso-auth.url=https://prod-rso.lol.qq.com:3000" "--rso_platform_id=HN1" "--rso-auth.client=lol" "--t.location=loltencent.sz.HN1" "--tglog-endpoint=https://tglogsz.datamore.qq.com/lolcli/report/" "--t.league_edge_url=https://ledge-hn1.lol.qq.com:22019" "--ccs=https://cc-hn1.lol.qq.com:8093" "--entitlements-url=https://hn1-entitlements.lol.qq.com:28088/api/token/v1" "--dradis-endpoint=http://some.url" "--remoting-auth-token=CWRt9oBVjG_eEtdu6QN8NA" "--app-port=63967" "--install-directory=d:\lol\lol(26)\LeagueClient" "--app-name=LeagueClient" "--ux-name=LeagueClientUx" "--ux-helper-name=LeagueClientUxHelper" "--log-dir=LeagueClient Logs" "--crash-reporting=" "--crash-environment=HN1" "--app-log-file-path=d:/lol/lol(26)/LeagueClient/../Game/Logs/LeagueClient Logs/2022-11-15T23-08-01_14576_LeagueClient.log" "--app-pid=14576" "--output-base-dir=d:/lol/lol(26)/LeagueClient/../Game" "--no-proxy-server" "--ignore-certificate-errors"  Win32_Process      20221115230801.763717+480  Win32_ComputerSystem  CHINAMI-2TSIBCP  LeagueClientUx.exe  d:\lol\lol(26)\LeagueClient\LeagueClientUx.exe                  14092   827                       6718750         1380                   200                    LeagueClientUx.exe  Win32_OperatingSystem  Microsoft Windows 11 רҵ▒▒|C:\Windows|\Device\Harddisk0\Partition3  28002                530228              60558       66476          14576            98272              520646656        114028              8         68071424          14092      49                      766                  3558                        807                      40400               68098464           1                                   25           21875000      505933824    10.0.22000      87216128        36125                7004038
//
// <nil>

// remoting-auth-token CWRt9oBVjG_eEtdu6QN8NA
// app-pid 14576
// app-port 63927
// curl https://127.0.0.1:63927/lol-perks/v1/currentpage
// https://127.0.0.1:63927/lol-banners/v1/current-summoner/flags

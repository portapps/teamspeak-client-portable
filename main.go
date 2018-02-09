//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	"os"
	"runtime"

	. "github.com/portapps/portapps"
)

func init() {
	Papp.ID = "teamspeak-client-portable"
	Papp.Name = "TeamSpeak Client"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = AppPathJoin("data")

	if runtime.GOARCH == "386" {
		Papp.Process = PathJoin(Papp.AppPath, "ts3client_win32.exe")
	} else {
		Papp.Process = PathJoin(Papp.AppPath, "ts3client_win64.exe")
	}

	Papp.Args = []string{"-nosingleinstance"}
	Papp.WorkingDir = Papp.AppPath

	OverrideEnv("TS3_CONFIG_DIR", Papp.DataPath)
	Launch(os.Args[1:])
}

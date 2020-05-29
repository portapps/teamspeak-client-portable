//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	"os"
	"runtime"

	"github.com/portapps/portapps/v2"
	"github.com/portapps/portapps/v2/pkg/log"
	"github.com/portapps/portapps/v2/pkg/utl"
)

var (
	app *portapps.App
)

func init() {
	var err error

	// Init app
	if app, err = portapps.New("teamspeak-client-portable", "TeamSpeak Client"); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	utl.CreateFolder(app.DataPath)
	if runtime.GOARCH == "386" {
		app.Process = utl.PathJoin(app.AppPath, "ts3client_win32.exe")
	} else {
		app.Process = utl.PathJoin(app.AppPath, "ts3client_win64.exe")
	}

	app.Args = []string{
		"-nosingleinstance",
	}

	utl.OverrideEnv("TS3_CONFIG_DIR", app.DataPath)

	defer app.Close()
	app.Launch(os.Args[1:])
}

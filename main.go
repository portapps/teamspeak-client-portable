//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	"os"
	"runtime"

	. "github.com/portapps/portapps"
	"github.com/portapps/portapps/pkg/utl"
)

var (
	app *App
)

func init() {
	var err error

	// Init app
	if app, err = New("teamspeak-client-portable", "TeamSpeak Client"); err != nil {
		Log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
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

	app.Launch(os.Args[1:])
}

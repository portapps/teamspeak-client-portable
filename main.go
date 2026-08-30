//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
package main

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/portapps/portapps/v3"
	"github.com/portapps/portapps/v3/pkg/log"
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
	if err := os.MkdirAll(app.DataPath, os.ModePerm); err != nil {
		log.Fatal().Err(err).Msg("Cannot create data directory.")
	}
	if runtime.GOARCH == "386" {
		app.Process = filepath.Join(app.AppPath, "ts3client_win32.exe")
	} else {
		app.Process = filepath.Join(app.AppPath, "ts3client_win64.exe")
	}

	app.Args = []string{
		"-nosingleinstance",
	}

	os.Setenv("TS3_CONFIG_DIR", app.DataPath)

	defer app.Close()
	app.Launch(os.Args[1:])
}

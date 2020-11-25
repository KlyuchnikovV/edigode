package main

import (
	"log"
	"os"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

var logger *log.Logger

func init() {
	file, err := os.OpenFile("./logs/log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0766)
	if err != nil {
		file = os.Stdout
		log.Printf("ERROR can't open file '%s'", err.Error())
	}
	logger = log.New(file, log.Prefix(), log.Flags())
	log.Printf("INFO logger initted")
	// TODO: load and parse config

	// VS Code debug workaround
	// TODO: if debug
	os.Unsetenv("ELECTRON_RUN_AS_NODE")
}

func main() {
	// Initialize astilectron
	var a, err = astilectron.New(logger, astilectron.Options{
		// TODO: from config
		AppName:            "Electron",
		BaseDirectoryPath:  "./build",
		VersionAstilectron: "0.33.0",
		VersionElectron:    "6.1.2",
	})
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer a.Close()

	// Start astilectron
	if err := a.Start(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	// TODO: provide path via config?
	// Create a new window
	w, err := a.NewWindow("./ui/build/index.html", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(600),
		Width:  astikit.IntPtr(800),
	})
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	if err := w.Create(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	if err := w.OpenDevTools(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	// Blocking pattern
	a.Wait()
}

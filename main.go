package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/blrsn/zentile/state"
	log "github.com/sirupsen/logrus"
)

func main() {
	setLogLevel()
	state.Populate()

	t := initTracker(CreateWorkspaces())
	bindKeys(t)

	// Run X event loop
	xevent.Main(state.X)
}

func generateStatus(t *tracker) {
	fname := Config.StatusFname
	if fname == "" {
		return
	}

	blob, _ := json.MarshalIndent(t.workspaces, "", "    ")
	os.Create(fname)
	err := os.WriteFile(fname, blob, 0600)
	if err != nil {
		log.Fatal(err)
	}
}

func setLogLevel() {
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "verbose mode")
	flag.Parse()

	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}

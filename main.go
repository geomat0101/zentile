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

	startups := Config.TileStartup
	for i := 0; i < len(startups); i += 1 {
		ws := startups[i]
		if ws >= len(t.workspaces) {
			log.Warn("Invalid workspace number in tile_workspaces: ", ws)
			continue
		}

		t.workspaces[uint(ws)].IsTiling = true
		t.workspaces[uint(ws)].Tile()
	}

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

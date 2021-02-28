// Package restart - quick and simple restart app
package restart

import (
	"flag"
	"os"
	"os/exec"

	"github.com/marstr/guid"
)

var (
	// OFF - disable flag of restart
	OFF bool
	// Logger - Logger
	Logger logger
	// DebugMode = do not restart app, only logging
	DebugMode bool
	cmd       *exec.Cmd // cmd - command being prepared for run
	start     chan int  // start - semaphore for start cmdStart() with exit code
	id        string    // id - unique ID (or name) for app
)

func init() {
	Logger = &defaultLogger{}
	
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	id = getID()
	cmd = exec.Command(exe, os.Args[1:]...)
	start = make(chan int, 1)
	go cmdStart()
}

func getID() string {
	id := flag.String("id", "", "unique ID (or name) for app")
	flag.Parse()
	if id == nil || *id == "" {
		*id = generateID()
		os.Args = append(os.Args, "-id", *id)
	}
	if DebugMode {
		Logger.Printf("restart::getID: app id: %v", *id)
	}
	return *id
}

func generateID() string {
	return guid.NewGUID().String()
}

// cmdStart - for once quick start of cmd.Start() without call func
func cmdStart() {
	defer os.Exit(<-start)
	if !DebugMode {
		err := cmd.Start()
		if err != nil {
			panic(err)
		}
	}
}

// Now - restarting current app and return a exitCode (via os.Exit(exitCode))
func Now(exitCode int, err error) {
	Logger.Printf("restart::Now: exitCode: %v, error: %v", exitCode, err)
	if !OFF {
		start <- exitCode
	}
}

// ID - return a unique id of crawler's manager, which used for setup job for monitorus server
func ID() string {
	return id
}

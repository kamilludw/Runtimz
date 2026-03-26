package main

import (
	"os"

	
	"github.com/kamilludw/runtimz/cmd"
	"github.com/kamilludw/runtimz/internal/logger"
	"github.com/kamilludw/runtimz/internal/output"
	"github.com/kamilludw/runtimz/internal/runtime"
	"github.com/kamilludw/runtimz/internal/state"
	"github.com/kamilludw/runtimz/internal/update"
)

func main() {
	if update.RunUpdaterIfRequested() {
		return
	}

	logger.Init()
	logger.Debug("main started", "args", os.Args)

	state := state.NewState()
	if err := state.Load(); err != nil {
		logger.Error("failed to load state", "err", err)
		output.Error("Failed to load state: " + err.Error())
		return
	}

	logger.Debug("state loaded")

	goRuntime := runtime.Init(state)
	runtime.Register(goRuntime)

	nodeRuntime := runtime.InitNode(state)
	runtime.Register(nodeRuntime)

	pythonRuntime := runtime.InitPython(state)
	runtime.Register(pythonRuntime)

	logger.Debug("runtimes registered")

	cmd.Run(os.Args, state)
}

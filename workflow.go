package main

import (
	"os"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	if os.Getenv("alfred_workflow_bundleid") == "" {
		return
	}
	wf = aw.New()
}

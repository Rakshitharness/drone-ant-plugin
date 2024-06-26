// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`

	// Goals defines the Ant targets/goals to execute.
	Goals string `envconfig:"PLUGIN_GOALS"`
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {

	// Split the goals into individual targets
	goals := strings.Fields(args.Goals)

	// Run `ant` command with specified goals
	antCmd := exec.Command("ant", goals...)
	antOutput, antErr := antCmd.CombinedOutput()
	if antErr != nil {
		fmt.Println("Error running 'ant "+args.Goals+"':", antErr)
		fmt.Println(string(antOutput))
		return fmt.Errorf("error running 'ant %s': %w", args.Goals, antErr)
	}
	fmt.Println("Output of 'ant "+args.Goals+"':", string(antOutput))

	return nil
}

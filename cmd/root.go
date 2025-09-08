// Copyright 2025 Victor Palma <victor.palma@rackspace.com>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "openCenter",
    Short: "openCenter CLI manages cluster configurations and GitOps scaffolding",
    PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
        // Handle config-dir flag
        if cfgDir, _ := cmd.Flags().GetString("config-dir"); cfgDir != "" {
            // Set environment variable so that config.ResolveConfigDir picks it up
            if err := os.Setenv("OPENCENTER_CONFIG_DIR", cfgDir); err != nil {
                return err
            }
        }
        return nil
    },
}

// Execute runs the root command and returns any error. This is the main
// entrypoint for the openCenter CLI. It is called by main.main().
//
// Inputs:
//   - version: The version string for the application.
//
// Outputs:
//   - error: An error if one occurred during execution.
func Execute(version string) error {
    rootCmd.Version = version
    // Register subcommands
    rootCmd.AddCommand(newClusterCmd())
    // Global persistent flag for config-dir
    rootCmd.PersistentFlags().String("config-dir", "", "configuration directory (defaults to ~/.config/openCenter on Linux/macOS)")
    return rootCmd.Execute()
}

// helpers for printing errors. In Cobra commands, returning an error
// will cause it to be printed and the process to exit with a non-zero
// code. Use fmt.Errorf to wrap underlying errors.
func failf(format string, a ...interface{}) error {
    return fmt.Errorf(format, a...)
}

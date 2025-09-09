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
	"os/exec"
	"path/filepath"

	"github.com/rackerlabs/openCenter/internal/ansible"
	"github.com/rackerlabs/openCenter/internal/config"
	"github.com/rackerlabs/openCenter/internal/gitops"
	"github.com/rackerlabs/openCenter/internal/terraform"
	"github.com/spf13/cobra"
)

func newClusterSetupCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "setup [name]",
        Short: "Setup GitOps directory (copy or render templates and initialise git)",
        Args:  cobra.MaximumNArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            var name string
            if len(args) > 0 {
                name = args[0]
            } else {
                var err error
                name, err = config.GetActive()
                if err != nil {
                    return err
                }
                if name == "" {
                    return fmt.Errorf("no active cluster; specify name")
                }
            }
            cfg, err := config.Load(name)
            if err != nil {
                return err
            }
            // Validate config
            if errs := config.Validate(cfg); len(errs) > 0 {
                for _, e := range errs {
                    fmt.Fprintln(cmd.ErrOrStderr(), e)
                }
                return fmt.Errorf("validation failed")
            }
            // Flag: --render
            render, _ := cmd.Flags().GetBool("render")
            if err := gitops.CopyBase(cfg, render); err != nil {
                return fmt.Errorf("failed to prepare gitops directory: %w", err)
            }
            // Provision terraform
            if err := terraform.Provision(cfg); err != nil {
                return fmt.Errorf("failed to provision terraform: %w", err)
            }
            // Provision ansible
            if err := ansible.Provision(cfg); err != nil {
                return fmt.Errorf("failed to provision ansible: %w", err)
            }
            // Write .opencenter marker
            markerPath := filepath.Join(cfg.GitOps.GitDir, ".opencenter")
            if err := os.WriteFile(markerPath, []byte(cfg.ClusterName), 0o644); err != nil {
                return err
            }
            // Initialise git repo if not present
            if _, statErr := os.Stat(filepath.Join(cfg.GitOps.GitDir, ".git")); os.IsNotExist(statErr) {
                cmdGit := exec.Command("git", "init", "-b", "main")
                cmdGit.Dir = cfg.GitOps.GitDir
                cmdGit.Stdout = cmd.OutOrStdout()
                cmdGit.Stderr = cmd.ErrOrStderr()
                if err := cmdGit.Run(); err != nil {
                    return fmt.Errorf("git init failed: %w", err)
                }
            }
            // Add and commit
            // git add .
            if err := runGit(cfg.GitOps.GitDir, []string{"add", "."}, cmd); err != nil {
                return err
            }
            // git commit -m "Initial commit" (allow empty)
            if err := runGit(cfg.GitOps.GitDir, []string{"commit", "-m", "Initial commit", "--allow-empty"}, cmd); err != nil {
                return err
            }
            fmt.Fprintln(cmd.OutOrStdout(), "Setup complete.")
            return nil
        },
    }
    cmd.Flags().Bool("render", false, "render templates (rather than copy)")
    return cmd
}

// runGit executes a git command in the given directory, streaming
// output to the provided cobra.Command's stdout/stderr. It returns an
// error if the command fails.
func runGit(dir string, args []string, cmd *cobra.Command) error {
    g := exec.Command("git", args...)
    g.Dir = dir
    g.Stdout = cmd.OutOrStdout()
    g.Stderr = cmd.ErrOrStderr()
    return g.Run()
}

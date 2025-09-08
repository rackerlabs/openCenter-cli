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
	"bytes"
	"fmt"
	"os/exec"

	"github.com/rackerlabs/openCenter/internal/config"
	"github.com/spf13/cobra"
)

func newClusterBootstrapCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bootstrap",
		Short: "Commit and push cluster configuration to the GitOps repository",
		RunE: func(cmd *cobra.Command, args []string) error {
			var name string
			var err error
			if len(args) > 0 {
				name = args[0]
			} else {
				name, err = config.GetActive()
				if err != nil {
					return err
				}
			}
			if name == "" {
				return fmt.Errorf("no active cluster; specify name or use 'select' to set it")
			}
			cfg, err := config.Load(name)
			if err != nil {
				return err
			}
			if cfg.GitOps.GitURL == "" {
				return fmt.Errorf("gitops.git_url not set; cannot bootstrap")
			}
			dir := cfg.GitOps.GitDir
			// Add and commit any changes
			if err := runGit(dir, []string{"add", "."}, cmd); err != nil {
				return err
			}
			if err := runGit(dir, []string{"commit", "-m", "Bootstrap commit", "--allow-empty"}, cmd); err != nil {
				// commit may fail if no git repo, but we proceed
				return err
			}
			// Set remote origin if not present
			if ok, err := hasOrigin(dir); err != nil {
				return err
			} else if !ok {
				add := exec.Command("git", "remote", "add", "origin", cfg.GitOps.GitURL)
				add.Dir = dir
				if err := add.Run(); err != nil {
					return fmt.Errorf("failed to add remote: %w", err)
				}
			}
			// Push main
			pushArgs := []string{"push", "-u", "origin", "main"}
			if force, _ := cmd.Flags().GetBool("force"); force {
				pushArgs = append(pushArgs, "--force")
			}
			if err := runGit(dir, pushArgs, cmd); err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), "Bootstrap push complete.")
			return nil
		},
	}
	cmd.Flags().Bool("force", false, "force push to the git remote")
	return cmd
}

// hasOrigin returns true if the git repository at `dir` has a remote
// named `origin`.
func hasOrigin(dir string) (bool, error) {
	remotes := exec.Command("git", "remote")
	remotes.Dir = dir
	out, err := remotes.Output()
	if err != nil {
		return false, err
	}
	//
	for _, line := range bytes.Split(out, []byte("\n")) {
		if string(line) == "origin" {
			return true, nil
		}
	}
	return false, nil
}

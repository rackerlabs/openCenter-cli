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

package terraform

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rackerlabs/openCenter/internal/config"
	"github.com/rackerlabs/openCenter/internal/provision"
)

// Provision generates Terraform files from templates. It creates the `main.tf`
// and `variables.tf` files in the directory specified by `cfg.Terraform.Path`
// within the GitOps repository.
//
// Inputs:
//   - cfg: The cluster configuration.
//
// Outputs:
//   - error: An error if one occurred during file generation.
func Provision(cfg config.Config) error {
	if !cfg.Terraform.Enabled {
		return nil
	}

	terraformDir := filepath.Join(cfg.GitOps.GitDir, cfg.Terraform.Path)
	if err := os.MkdirAll(terraformDir, 0755); err != nil {
		return fmt.Errorf("failed to create terraform directory: %w", err)
	}

	// Create main.tf
	mainTfPath := filepath.Join(terraformDir, "main.tf")
	mainTfFile, err := os.Create(mainTfPath)
	if err != nil {
		return fmt.Errorf("failed to create main.tf: %w", err)
	}
	defer mainTfFile.Close()

	if err := provision.Templates.ExecuteTemplate(mainTfFile, "main.tf.tmpl", cfg); err != nil {
		return fmt.Errorf("failed to execute main.tf template: %w", err)
	}

	// Create variables.tf
	variablesTfPath := filepath.Join(terraformDir, "variables.tf")
	variablesTfFile, err := os.Create(variablesTfPath)
	if err != nil {
		return fmt.Errorf("failed to create variables.tf: %w", err)
	}
	defer variablesTfFile.Close()

	if err := provision.Templates.ExecuteTemplate(variablesTfFile, "variables.tf.tmpl", cfg); err != nil {
		return fmt.Errorf("failed to execute variables.tf template: %w", err)
	}

	return nil
}

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

package ansible

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/rackerlabs/openCenter/internal/config"
	"github.com/rackerlabs/openCenter/internal/provision"
)

func TestMain(m *testing.M) {
	// The templates are in a different package, so we need to initialize them explicitly.
	// This is a bit of a hack, but it's the easiest way to test this.
	if err := provision.Init(); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestProvision(t *testing.T) {
	dir := t.TempDir()
	cfg := config.NewDefault("test")
	cfg.GitOps.GitDir = dir
    cfg.IAC.Counts["master"] = 1
    cfg.IAC.Counts["worker"] = 2

	if err := Provision(cfg); err != nil {
		t.Fatal(err)
	}

	inventoryPath := filepath.Join(dir, "ansible", "inventory")
	if _, err := os.Stat(inventoryPath); os.IsNotExist(err) {
		t.Error("inventory file was not created")
	}

	data, err := os.ReadFile(inventoryPath)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(data, []byte("[master]")) {
		t.Error("inventory does not contain master group")
	}
	if !bytes.Contains(data, []byte("[worker]")) {
		t.Error("inventory does not contain worker group")
	}
	t.Logf("inventory: %s", data)
}

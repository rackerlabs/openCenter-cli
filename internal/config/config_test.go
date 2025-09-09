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

package config

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	dir := t.TempDir()
	os.Setenv("OPENCENTER_CONFIG_DIR", dir)
	defer os.Unsetenv("OPENCENTER_CONFIG_DIR")

	t.Run("Save and Load", func(t *testing.T) {
		cfg := NewDefault("test")
		if err := Save(cfg); err != nil {
			t.Fatal(err)
		}

		loaded, err := Load("test")
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(cfg, loaded) {
			t.Errorf("loaded config does not match saved config")
		}
	})

	t.Run("List", func(t *testing.T) {
		names, err := List()
		if err != nil {
			t.Fatal(err)
		}
		if len(names) != 1 || names[0] != "test" {
			t.Errorf("unexpected list result: %v", names)
		}
	})

	t.Run("SetActive and GetActive", func(t *testing.T) {
		if err := SetActive("test"); err != nil {
			t.Fatal(err)
		}
		active, err := GetActive()
		if err != nil {
			t.Fatal(err)
		}
		if active != "test" {
			t.Errorf("unexpected active cluster: %s", active)
		}
	})

	t.Run("Validate", func(t *testing.T) {
		cfg := NewDefault("test")
		cfg.GitOps.GitDir = "/tmp/gitops"
        cfg.IAC.Networking.DNSZoneName = "example.com"
		errs := Validate(cfg)
		if len(errs) != 0 {
			t.Errorf("unexpected validation errors: %v", errs)
		}

		// Test validation failure
        cfg.IAC.Networking.UseOctavia = true
        cfg.IAC.Networking.VRRPEnabled = true
		errs = Validate(cfg)
		if len(errs) == 0 {
			t.Error("expected validation error, but got none")
		}
	})
}

func TestResolveConfigDir(t *testing.T) {
	// Unset env var to test default behavior
	os.Unsetenv("OPENCENTER_CONFIG_DIR")

	dir, err := ResolveConfigDir()
	if err != nil {
		t.Fatal(err)
	}
	home, _ := os.UserHomeDir()
	expected := filepath.Join(home, ".config", "openCenter")
	if dir != expected {
		t.Errorf("expected config dir %s, but got %s", expected, dir)
	}

	// Set env var to test override
	testDir := "/tmp/openCenter-test"
	os.Setenv("OPENCENTER_CONFIG_DIR", testDir)
	defer os.Unsetenv("OPENCENTER_CONFIG_DIR")

	dir, err = ResolveConfigDir()
	if err != nil {
		t.Fatal(err)
	}
	if dir != testDir {
		t.Errorf("expected config dir %s, but got %s", testDir, dir)
	}
}

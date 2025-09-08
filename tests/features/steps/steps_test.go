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

package steps

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

// TestFeatures runs the BDD scenarios. It uses Godog’s suite to
// register steps defined in helpers.go. Running `go test` in this
// package will execute the feature files automatically.
func TestFeatures(t *testing.T) {
	opts := godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "pretty",
		Paths:  []string{".."},
		Tags:   "~@wip",
	}

	w, err := newWorld()
	if err != nil {
		t.Fatalf("failed to create world: %v", err)
	}

	suite := godog.TestSuite{
		Name: "openCenter",
		ScenarioInitializer: func(s *godog.ScenarioContext) {
			s.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
				tmp, err := ioutil.TempDir("", "opencenter-test-")
				if err != nil {
					t.Fatalf("failed to create temp dir: %v", err)
				}
				w.tmpDir = tmp
				if err := w.isolateConfigDir(); err != nil {
					t.Fatalf("failed to isolate config dir: %v", err)
				}
				return ctx, nil
			})

			RegisterSteps(s, t, w)

			s.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
				os.RemoveAll(w.tmpDir)
				os.RemoveAll(w.configDir)
				return ctx, err
			})
		},
		Options: &opts,
	}
	if suite.Run() != 0 {
		t.Fail()
	}
}

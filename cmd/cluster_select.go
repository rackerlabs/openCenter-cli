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
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rackerlabs/openCenter/internal/config"
	"github.com/spf13/cobra"
)

var (
	docStyle    = lipgloss.NewStyle().Margin(1, 2)
	titleStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFDF5")).Background(lipgloss.Color("#25A065")).Padding(0, 1)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
)

type item struct {
	title string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return i.title }

type model struct {
	list     list.Model
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = i.title
			}
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != "" {
		return ""
	}
	if m.quitting {
		return ""
	}
	return docStyle.Render(m.list.View())
}

func newClusterSelectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "select [name]",
		Short: "Select the active cluster",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var name string
			if len(args) > 0 {
				name = args[0]
			}
			// If name not provided, prompt
			if name == "" {
				names, err := config.List()
				if err != nil {
					return err
				}
				if len(names) == 0 {
					return errors.New("no clusters defined")
				}

				items := []list.Item{}
				for _, name := range names {
					items = append(items, item{title: name})
				}

				delegate := list.NewDefaultDelegate()
				delegate.Styles.SelectedTitle = selectedItemStyle
				delegate.Styles.NormalTitle = itemStyle

				l := list.New(items, delegate, 0, 0)
				l.Title = "Select a cluster"
				l.Styles.Title = titleStyle

				m := model{list: l}
				p := tea.NewProgram(m, tea.WithAltScreen())

				finalModel, err := p.Run()
				if err != nil {
					return err
				}

				m, ok := finalModel.(model)
				if !ok {
					return errors.New("could not cast model")
				}
				name = m.choice
			}

			if name == "" {
				return nil
			}

			// Check that cluster config file exists
			path, err := config.ConfigPath(name)
			if err != nil {
				return err
			}
			if _, err := os.Stat(path); err != nil {
				return fmt.Errorf("cluster '%s' not found. Use 'openCenter cluster list' to see available clusters", name)
			}
			// Set active
			if err := config.SetActive(name); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Active cluster set to %s\n", name)
			return nil
		},
	}
}

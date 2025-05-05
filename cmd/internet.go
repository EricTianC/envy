/*
Copyright © 2025 Eric Tian <erictianc@outlook.com>
*/
package cmd

import (
	"log/slog"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// internetCmd represents the internet command
var internetCmd = &cobra.Command{
	Use:   "internet",
	Short: "Check internet connections",
	Long: `Run internet connection diagnostic programms.
By running this, envy will know whether to use a mirror
or connect directly in the following solving stages.

Can also work as a independent internet doctor panel.
(It is designed to have a pretty tui)`,
	Run: internetHandler,
}

func internetHandler(cmd *cobra.Command, args []string) {
	slog.Debug("running internet command:", "arguments", args)
	items := []list.Item{
		item{title: "a", desc: "aaaaa"},
		item{title: "b", desc: "bbbbb"},
	}
	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "My title"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		slog.Error("Error running command:", "error", err)
		os.Exit(1)
	}
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
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
	return docStyle.Render(m.list.View())
}

func init() {
	checkCmd.AddCommand(internetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// internetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// internetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

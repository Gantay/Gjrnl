package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
	"os"
)

type model struct {
	textInput textinput.Model
	err       error
}
type (
	errMsg error
)

func bb() {

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

func initialModel() model {
	// lipgloss
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80 // default to 80 if there's an error
	}
	var Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Underline(true).
		Align(lipgloss.Center).
		PaddingTop(0).
		PaddingBottom(0).
		PaddingRight(0).
		PaddingLeft(0).
		MarginTop(1).
		MarginRight(0).
		MarginLeft(0).
		MarginBottom(0).
		Width(width - 1)
	fmt.Println(Title.Render("Gjrnl"))

	// lipgloss_END

	ti := textinput.New()
	ti.Placeholder = "Title"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 156
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4"))
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#5d6d7e"))
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA"))

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"Jurnal Intrey...\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"

}

func CSS() {

	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80  // default to 80 if there's an error
		height = 40 // default to 40 if there's an error
	}
	// var Title = lipgloss.NewStyle().
	// 	Bold(true).
	// 	Foreground(lipgloss.Color("#FAFAFA")).
	// 	Background(lipgloss.Color("#7D56F4")).
	// 	Underline(true).
	// 	Align(lipgloss.Center).
	// 	//

	// 	PaddingTop(0).
	// 	PaddingBottom(0).
	// 	PaddingRight(0).
	// 	PaddingLeft(0).
	// 	//
	// 	MarginTop(2).
	// 	MarginRight(0).
	// 	MarginLeft(0).
	// 	MarginBottom(0).
	// 	Width(width - 1)

	// fmt.Println(Title.Render("Gjrnl"))

	var Borders = lipgloss.NewStyle().
		Width(width-3).
		Height(height-9).
		Border(lipgloss.BlockBorder(), false, true, true, true).
		BorderForeground(lipgloss.Color("63"))
	fmt.Println(Borders.Render())

}

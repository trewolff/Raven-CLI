package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type screen int

var tickets = [...]string{"Ticket 1", "Ticket 2", "Ticket 3", "Ticket 4"}
var alerts = [...]string{"Alert 1", "Alert 2", "Alert 3", "Alert 4"}
var automations = [...]string{"Playbooks", "Actions", "Enrichments"}

const (
	splashScreen screen = iota
	ticketsScreen
	workTicketScreen
	alertsScreen
	automationsScreen
)

type Model struct {
	currentScreen screen
	cursor        int
	choices       []string
}

func InitialModel() Model {
	choices := []string{"Tickets", "Alerts", "Automations", "Exit"}
	return Model{
		currentScreen: splashScreen,
		choices:       choices,
	}
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			m.currentScreen = splashScreen
		case "esc":
			if m.currentScreen == ticketsScreen {
				m.currentScreen = splashScreen
			} else if m.currentScreen == workTicketScreen {
				m.currentScreen = ticketsScreen

			} else if m.currentScreen == alertsScreen {
				m.currentScreen = splashScreen

			} else if m.currentScreen == automationsScreen {
				m.currentScreen = splashScreen

			}
		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.currentScreen == ticketsScreen && m.cursor > 0 {
				m.cursor--
			}
			if m.currentScreen == alertsScreen && m.cursor > 0 {
				m.cursor--
			}
			if m.currentScreen == splashScreen && m.cursor > 0 {
				m.cursor--
			}
			if m.currentScreen == automationsScreen && m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.currentScreen == ticketsScreen && m.cursor < len(tickets)-1 {
				m.cursor++
			}
			if m.currentScreen == alertsScreen && m.cursor < len(tickets)-1 {
				m.cursor++
			}
			if m.currentScreen == splashScreen && m.cursor < len(m.choices)-1 {
				m.cursor++
			}
			if m.currentScreen == automationsScreen && m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			if m.currentScreen == ticketsScreen {
				m.currentScreen = workTicketScreen
			}
			if m.currentScreen == splashScreen {
				switch m.cursor {
				case 0:
					m.currentScreen = ticketsScreen
				case 1:
					m.currentScreen = alertsScreen
				case 2:
					m.currentScreen = automationsScreen
				case 3:
					return m, tea.Quit
				}
			}
		}
	}

	// Return the updated Model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

var (
	boarderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true).
			Align(lipgloss.Center)

	menuStyle = lipgloss.NewStyle().
			Margin(1, 2)
)

func (m Model) splashView() string {
	title := titleStyle.Render("Welcome to the ticketing system")
	instructions := "Select one..."
	menu := ""
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		menu += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	return menuStyle.Render(boarderStyle.Render(fmt.Sprintf("%s\n\n%s%s", title, menu, instructions)))
}

func (m Model) ticketsView() string {
	title := titleStyle.Render("Tickets screen:")
	instructions := "Use k/j to navigate and enter to select"
	menu := ""
	for i, ticket := range tickets {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		menu += fmt.Sprintf("%s %s\n", cursor, ticket)
	}
	return menuStyle.Render(boarderStyle.Render(fmt.Sprintf("%s\n\n%s%s", title, menu, instructions)))
}

func (m Model) alertsView() string {
	title := titleStyle.Render("Alerts screen:")
	instructions := "Use k/j to navigate and enter to select"
	menu := ""
	for i, alert := range alerts {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		menu += fmt.Sprintf("%s %s\n", cursor, alert)
	}
	return menuStyle.Render(boarderStyle.Render(fmt.Sprintf("%s\n\n%s%s", title, menu, instructions)))
}

func (m Model) automationsView() string {
	title := titleStyle.Render("Automations screen:")
	instructions := "Use k/j to navigate and enter to select"
	menu := ""
	for i, automation := range automations {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		menu += fmt.Sprintf("%s %s\n", cursor, automation)
	}
	return menuStyle.Render(boarderStyle.Render(fmt.Sprintf("%s\n\n%s%s", title, menu, instructions)))
}

func (m Model) workTicketView() string {
	title := titleStyle.Render("Working on a ticket...")
	instructions := "Press q to quit to the main menu"
	// TODO:
	return menuStyle.Render(boarderStyle.Render(fmt.Sprintf("%s\n\n%s", title, instructions)))
}

func (m Model) View() string {
	switch m.currentScreen {
	case splashScreen:
		return m.splashView()
	case ticketsScreen:
		return m.ticketsView()
	case alertsScreen:
		return m.alertsView()
	case automationsScreen:
		return m.automationsView()
	case workTicketScreen:
		return m.workTicketView()
	default:
		return ""
	}
}

const (
	padding  = 2
	maxWidth = 50
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

type tickMsg time.Time

type Model2 struct {
	progress progress.Model
}

func (m Model2) Init() tea.Cmd {
	return tickCmd()
}

func (m Model2) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		if m.progress.Percent() == 1.0 {
			return m, tea.Quit
		}

		// Note that you can also use progress.Model.SetPercent to set the
		// percentage value explicitly, too.
		cmd := m.progress.IncrPercent(0.20)
		return m, tea.Batch(tickCmd(), cmd)

	// FrameMsg is sent when the progress bar wants to animate itself
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}
}

func (m Model2) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + m.progress.View() + "\n\n" +
		pad + helpStyle("Press any key to quit")
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*400, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func SecondModel() Model2 {
	return Model2{
		progress: progress.New(progress.WithDefaultGradient()),
	}
}

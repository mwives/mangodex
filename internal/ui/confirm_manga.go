package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/mwives/mangodex/pkg/mangadex"
)

const maxWidth = 80

var (
	indigo = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	green  = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
)

type Styles struct {
	Base,
	MangaInfoBox,
	MangaInfoHeader lipgloss.Style
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.MangaInfoBox = lg.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(indigo).
		PaddingLeft(1).
		MarginTop(1)
	s.MangaInfoHeader = lg.NewStyle().Foreground(green).Bold(true)
	return &s
}

type Model struct {
	lg     *lipgloss.Renderer
	styles *Styles
	form   *huh.Form
	manga  mangadex.MangaResult
	width  int
}

func NewModel(manga mangadex.MangaResult) Model {
	m := Model{
		width: maxWidth,
		manga: manga,
	}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Key("confirm").
				Title("Confirm the manga selection?").
				Validate(func(v bool) error {
					if !v {
						return fmt.Errorf("please confirm the manga selection")
					}
					return nil
				}).
				Affirmative("Yes").
				Negative("No"),
		),
	).
		WithWidth(45)

	return m
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd

	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		// Quit when the form is done.
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := m.styles

	switch m.form.State {
	case huh.StateCompleted:
		var b strings.Builder
		fmt.Fprintf(&b, "Manga selection confirmed!")
		return s.MangaInfoBox.Margin(0, 1).Padding(1, 2).Width(48).Render(b.String()) + "\n\n"
	default:
		// Form
		v := strings.TrimSuffix(m.form.View(), "\n\n")
		form := m.lg.NewStyle().Margin(1, 0).Render(v)

		// Manga details
		var mangaInfo string
		{
			mangaInfo = s.MangaInfoBox.
				Width(m.width).
				Render(
					s.MangaInfoHeader.Render("Title: ") + m.manga.Title + "\n" +
						s.MangaInfoHeader.Render("Author: ") + m.manga.Author + "\n" +
						s.MangaInfoHeader.Render("Genres: ") + m.manga.Genres + "\n" +
						s.MangaInfoHeader.Render("Alt Titles: ") + m.manga.AltTitles + "\n" +
						s.MangaInfoHeader.Render("Year: ") + fmt.Sprintf("%d", m.manga.Year) + "\n" +
						s.MangaInfoHeader.Render("Status: ") + m.manga.Status + "\n",
				)
		}

		errors := m.form.Errors()
		if len(errors) > 0 {
			fmt.Println(errors[0].Error())
		}
		body := lipgloss.JoinVertical(lipgloss.Top, mangaInfo, form)

		return s.Base.Render(body)
	}
}

func ConfirmMangaSelection(manga mangadex.MangaResult) {
	_, err := tea.NewProgram(NewModel(manga)).Run()
	if err != nil {
		fmt.Println(err)
	}
}

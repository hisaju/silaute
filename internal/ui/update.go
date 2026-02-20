package ui

import (
	"math/rand/v2"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/hisaju/silaute/internal/responses"
)

type tickMsg struct{}

func tickCmd() tea.Cmd {
	return tea.Tick(40*time.Millisecond, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

// thinkTickInterval はスピナーの更新間隔（100ms）に対するtickMsg回数。
// tickは40msごとなので、100ms / 40ms ≒ 2〜3 tick でスピナー1フレーム。
const spinnerTickInterval = 3

// Update はBubble Teaのイベントを処理する。
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKey(msg)

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tickMsg:
		return m.handleTick()
	}

	if m.state == stateInput {
		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) handleTick() (tea.Model, tea.Cmd) {
	switch m.state {
	case stateThinking:
		m.thinkTicks++
		if m.thinkTicks >= m.thinkTarget {
			// 思考完了 → ストリーミング開始
			m.state = stateTyping
			m.answer = []rune(responses.Random())
			m.displayLen = 0
		}
		return m, tickCmd()

	case stateTyping:
		if m.displayLen < len(m.answer) {
			m.displayLen++
			return m, tickCmd()
		}
		// 全文表示完了
		return m.finishTyping(), nil
	}

	return m, nil
}

func (m Model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyCtrlC:
		return m, tea.Quit

	case tea.KeyEnter:
		if m.state == stateThinking {
			// 思考中にEnterで即ストリーミングへ
			m.state = stateTyping
			m.answer = []rune(responses.Random())
			m.displayLen = 0
			return m, tickCmd()
		}
		if m.state == stateTyping {
			return m.finishTyping(), nil
		}
		q := m.textInput.Value()
		if q == "" {
			return m, nil
		}
		m.textInput.SetValue("")
		m.history = append(m.history, exchange{question: q})
		// 思考フェーズに入る（1.5〜4秒 = 37〜100 tick @40ms）
		m.state = stateThinking
		m.thinkTicks = 0
		m.thinkTarget = 37 + rand.IntN(63)
		return m, tickCmd()

	case tea.KeyEsc:
		if m.state == stateThinking {
			m.state = stateTyping
			m.answer = []rune(responses.Random())
			m.displayLen = 0
			return m, tickCmd()
		}
		if m.state == stateTyping {
			return m.finishTyping(), nil
		}
		return m, nil
	}

	if m.state == stateInput {
		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) finishTyping() Model {
	m.displayLen = len(m.answer)
	if len(m.history) > 0 {
		m.history[len(m.history)-1].answer = string(m.answer)
	}
	m.answer = nil
	m.displayLen = 0
	m.state = stateInput
	if len(m.history) > maxHistory {
		m.history = m.history[len(m.history)-maxHistory:]
	}
	m.textInput.Focus()
	return m
}

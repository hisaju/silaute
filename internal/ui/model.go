package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type state int

const (
	stateInput    state = iota
	stateThinking       // 思考中（スピナー表示）
	stateTyping         // 回答ストリーミング中
)

const maxHistory = 5

// スピナーのフレーム（Claude Code風）
var spinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

type exchange struct {
	question string
	answer   string
}

// Model はBubble TeaのModelインターフェースを実装する。
type Model struct {
	state        state
	textInput    textinput.Model
	answer       []rune
	displayLen   int
	history      []exchange
	width        int
	height       int
	thinkTicks   int // 思考中の経過ティック数
	thinkTarget  int // 何ティックで思考完了するか
}

// New は初期化済みのModelを返す。
func New() Model {
	ti := textinput.New()
	ti.Placeholder = "何か聞いてみてください..."
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 60
	ti.PromptStyle = inputPromptStyle
	ti.Prompt = "❯ "

	return Model{
		state:     stateInput,
		textInput: ti,
		width:     80,
		height:    24,
	}
}

// Init はBubble Teaの初期コマンドを返す。
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

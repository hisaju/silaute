package ui

import "github.com/charmbracelet/lipgloss"

var (
	// ヘッダーボックス
	headerBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("205")).
			Padding(0, 1).
			MarginBottom(1)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("244"))

	// ユーザーメッセージ
	userPromptStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true)

	userTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("15")).
			Bold(true)

	// 回答ブロック（左ボーダー）
	botBorderStyle = lipgloss.NewStyle().
			BorderLeft(true).
			BorderStyle(lipgloss.ThickBorder()).
			BorderForeground(lipgloss.Color("205")).
			PaddingLeft(1)

	botTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252"))

	// 履歴（薄め表示）
	historyPromptStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("240")).
				Bold(true)

	historyUserStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("245"))

	historyBotBorderStyle = lipgloss.NewStyle().
				BorderLeft(true).
				BorderStyle(lipgloss.ThickBorder()).
				BorderForeground(lipgloss.Color("238")).
				PaddingLeft(1)

	historyBotStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))

	// 入力プロンプト
	inputPromptStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("205")).
				Bold(true)

	// 思考中
	thinkingSpinnerStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("205")).
				Bold(true)

	thinkingLabelStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("205")).
				Italic(true)

	thinkingSecStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("241"))

	// フッター
	footerKeyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205"))

	footerDescStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))

	footerSepStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("236"))
)

package ui

import (
	"fmt"
	"strings"
)

// View はBubble Teaの画面描画を行う。
func (m Model) View() string {
	var b strings.Builder

	// ヘッダーボックス
	header := fmt.Sprintf(
		"%s\n%s",
		titleStyle.Render("✻ Silaute Code v0.1.0"),
		subtitleStyle.Render("素人AIアシスタント — /help でヘルプ（嘘）"),
	)
	b.WriteString(headerBoxStyle.Render(header))
	b.WriteString("\n")

	// 会話履歴
	for i, ex := range m.history {
		isActive := (i == len(m.history)-1) && m.state != stateInput
		if isActive {
			m.renderActiveExchange(&b, ex)
		} else {
			m.renderHistoryExchange(&b, ex)
		}
	}

	// 入力プロンプト（thinking/typing中は非表示）
	if m.state == stateInput {
		b.WriteString("  " + m.textInput.View())
		b.WriteString("\n")
	}

	// フッター
	b.WriteString("\n")
	b.WriteString(m.renderFooter())
	b.WriteString("\n")

	return b.String()
}

func (m Model) renderActiveExchange(b *strings.Builder, ex exchange) {
	// ユーザーメッセージ
	b.WriteString(fmt.Sprintf("  %s %s\n",
		userPromptStyle.Render("❯"),
		userTextStyle.Render(ex.question),
	))
	b.WriteString("\n")

	if m.state == stateThinking {
		// 思考中: スピナー + 経過秒数
		frame := spinnerFrames[(m.thinkTicks/spinnerTickInterval)%len(spinnerFrames)]
		elapsed := float64(m.thinkTicks*40) / 1000.0
		spinner := thinkingSpinnerStyle.Render(frame)
		label := thinkingLabelStyle.Render(" 考え中...")
		sec := thinkingSecStyle.Render(fmt.Sprintf(" %.1fs", elapsed))
		b.WriteString(fmt.Sprintf("  %s%s%s\n\n", spinner, label, sec))
	} else {
		// ストリーミング中の回答
		displayed := string(m.answer[:m.displayLen])
		cursor := ""
		if m.displayLen < len(m.answer) {
			cursor = "▌"
		}
		content := botTextStyle.Render(displayed + cursor)
		b.WriteString("  " + botBorderStyle.Render(content))
		b.WriteString("\n\n")
	}
}

func (m Model) renderHistoryExchange(b *strings.Builder, ex exchange) {
	b.WriteString(fmt.Sprintf("  %s %s\n",
		historyPromptStyle.Render("❯"),
		historyUserStyle.Render(ex.question),
	))
	b.WriteString("\n")

	content := historyBotStyle.Render(ex.answer)
	b.WriteString("  " + historyBotBorderStyle.Render(content))
	b.WriteString("\n\n")
}

func (m Model) renderFooter() string {
	sep := footerSepStyle.Render(" | ")

	items := []string{
		footerKeyStyle.Render("ctrl+c") + footerDescStyle.Render(" 終了"),
	}
	if m.state == stateThinking || m.state == stateTyping {
		items = append(items,
			footerKeyStyle.Render("enter") + footerDescStyle.Render(" スキップ"),
		)
	}

	return "  " + strings.Join(items, sep)
}

# Silaute Code

Claude Code パロディ版の対話型CLIツール。何を聞いても「素人なので分かりません」的な回答を返す素人AIアシスタント。

```
╭──────────────────────────────────────────────╮
│ ✻ Silaute Code v0.1.0                       │
│ 素人AIアシスタント — /help でヘルプ（嘘）        │
╰──────────────────────────────────────────────╯

  ❯ Goでソートアルゴリズムを実装して

  ⠹ 考え中... 2.3s

  ❯ Pythonの基本を教えて

  ┃ 素人なのでさっぱり分かりません!

  ❯ _

  ctrl+c 終了 | enter スキップ
```

## インストール

### Homebrew

```bash
brew tap hisaju/tap
brew install silaute
```

### Docker

```bash
docker compose run --rm silaute
```

### ソースから

```bash
go install github.com/hisaju/silaute/cmd/silaute@latest
```

## 機能

- Claude Code風の対話UI（Bubble Tea + Lip Gloss）
- 思考中スピナー付きの回答演出
- 一文字ずつストリーミング表示
- 8カテゴリ31種類のランダム回答
- 直近5件の会話履歴表示

## 開発

```bash
# ローカルビルド
go build -o silaute ./cmd/silaute/

# 実行
./silaute

# Dockerでビルド＆実行
docker compose run --rm silaute
```

## リリース

```bash
git tag v0.x.x
git push origin v0.x.x
# GitHub Actions が自動でリリース＆Homebrew formula更新
```

## License

MIT

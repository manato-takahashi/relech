# relech

Release Checker - 複数リポジトリの develop → main(master) 間の差分を一括チェックし、Draft PR の作成まで行うCLIツール。

毎週のリリース作業で「どのリポジトリにリリースが必要か」を一発で確認できる。

## 前提

- **Go 1.22+**

  ```bash
  brew install go
  ```

  Go 1.21以降はtoolchain自動管理があるので、バージョンマネージャ（goenv等）は不要。

- **[GitHub CLI (`gh`)](https://cli.github.com/)**

  ```bash
  brew install gh
  gh auth login
  ```

  relechは `gh` コマンド経由でGitHub APIを叩くため、認証済みである必要がある。

## Install

```bash
go install github.com/manato-takahashi/relech@latest
```

`go install` で入ったバイナリは `$HOME/go/bin`（= `$(go env GOPATH)/bin`）に置かれる。このディレクトリがPATHに通っていない場合、`relech` コマンドが見つからない。

### PATHを通す（初回のみ）

使っているシェルに応じて追記：

```bash
# zsh (macOSデフォルト)
echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc

# bash
echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

### 動作確認

```bash
which relech
# → /Users/xxx/go/bin/relech が返ればOK

relech --help
```

## Setup

設定ファイルを `~/.config/relech/config.yaml` に作成する:

```bash
mkdir -p ~/.config/relech
```

```yaml
# ~/.config/relech/config.yaml
repositories:
  - name: frontend-app
    owner: your-org
    base: main
    head: develop

  - name: backend-api
    owner: your-org
    base: master       # リポジトリごとにbase branchを変更可能
    head: develop

pr_template:
  title: "本番リリース({{.Date}})"
```

デフォルトで `~/.config/relech/config.yaml` を読む。別のファイルを指定する場合:

```bash
relech check --config /path/to/config.yaml
```

## Usage

### 差分チェック

```bash
# リッチ表示（PRタイトル込み）
$ relech check

foodslabo  56 commits / 14 PRs ⚠
  - #3386 reserve_agent_scoutのエラーハンドリング改善
  - #3405 hotfixの差分埋め
  - #3311 DEV-2191 PF Admin エージェント企業一覧で企業のフリーワード検索

foodslabo-pf-ag  1 commits / 1 PRs ✓
  - #400 feat: Sentryエラー調査スキルを追加

admin-front no changes
toB-front no changes

Summary: 2/4 repositories need release
```

```bash
# 省略表示
$ relech check --short

foodslabo  56 commits / 14 PRs ⚠
foodslabo-pf-ag  1 commits / 1 PRs ✓
admin-front no changes
toB-front no changes

Summary: 2/4 repositories need release
```

commits と PRs の数が一致していれば ✓（スカッシュマージ済み）、不一致なら ⚠ を表示。

### Draft PR 作成

```bash
$ relech prepare

Creating draft PRs...

foodslabo  https://github.com/org/foodslabo/pull/3420
foodslabo-pf-ag  https://github.com/org/foodslabo-pf-ag/pull/401
admin-front no changes
toB-front no changes

Created 2 draft PRs
```

- 常に Draft PR として作成（通常PRの作成機能はなし）
- PRタイトルは `本番リリース(YYYY-MM-DD)` で自動生成
- PR本文にリリース対象のPR番号リストを自動挿入

## うまくいかないとき

| 症状 | 確認 |
|------|------|
| `command not found: relech` | `echo $PATH` で `$HOME/go/bin` が含まれているか。シェルを再起動したか |
| `gh: not found` | `gh` CLIがインストール・認証済みか（`gh auth status` で確認） |
| `failed to read config` | `~/.config/relech/config.yaml` が存在するか、YAML形式が正しいか |
| `relech check` が固まったように見える | 初回はGitHub API呼び出しで数秒かかる。リポ数×1-2秒が目安 |

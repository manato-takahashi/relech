package github

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// CompareResult は2ブランチ間の差分情報を保持する
type CompareResult struct {
	AheadBy int      // head が base より何コミット先行しているか
	PRs     []PRInfo // コミットメッセージから抽出したPR一覧
}

// PRInfo はマージ済みPRの情報を保持する
type PRInfo struct {
	Number int
	Title  string
}

// Compare は gh CLI を使って2ブランチ間の差分を取得する
//
// TODO: 自分で実装するパート
//
// やること:
//
//	gh api repos/{owner}/{repo}/compare/{base}...{head}
//	を os/exec で呼び出し、JSONレスポンスから ahead_by を取得する
//
// 参考: gh api の出力は JSON なので、encoding/json でパースする
//
// 戻り値:
//   - CompareResult: 差分情報
//   - error: エラー（gh コマンド失敗、JSONパース失敗など）
func Compare(owner, repo, base, head string) (*CompareResult, error) {
	path := fmt.Sprintf("repos/%s/%s/compare/%s...%s", owner, repo, base, head)
	jq := `{ahead_by: .ahead_by, messages: [.commits[].commit.message]}`
	out, err := exec.Command("gh", "api", path, "--jq", jq).Output()
	if err != nil {
		return nil, err
	}

	var result struct {
		AheadBy  int      `json:"ahead_by"`
		Messages []string `json:"messages"`
	}
	err = json.Unmarshal(out, &result)
	if err != nil {
		return nil, err
	}

	mergeRe := regexp.MustCompile(`Merge pull request #(\d+)`)
	squashRe := regexp.MustCompile(`\(#(\d+)\)`)
	var prs []PRInfo
	for _, msg := range result.Messages {
		if matches := mergeRe.FindStringSubmatch(msg); matches != nil {
			// 通常マージの処理
			num, _ := strconv.Atoi(matches[1])
			parts := strings.SplitN(msg, "\n\n", 2)
			title := ""
			if len(parts) >= 2 {
				title = parts[1]
			}
			prs = append(prs, PRInfo{Number: num, Title: title})
		} else if matches := squashRe.FindStringSubmatch(msg); matches != nil {
			// スカッシュマージの処理
			num, _ := strconv.Atoi(matches[1])
			loc := squashRe.FindStringIndex(msg)
			title := strings.TrimSpace(msg[:loc[0]])
			prs = append(prs, PRInfo{Number: num, Title: title})
		}
	}

	return &CompareResult{AheadBy: result.AheadBy, PRs: prs}, nil
}

// GetMergedPRs は head → base 方向にマージ済みのPR一覧を取得する
//
// TODO: 自分で実装するパート
//
// やること:
//
//	gh pr list --repo {owner}/{repo} --base {head} --state merged
//	を os/exec で呼び出し、PR番号とタイトルを取得する
//
// 注意:
//
//	ここで取得するのは「developにマージされたPR」= feature → develop のPR
//	base に head(develop) を指定するのがポイント
//
// 戻り値:
//   - []PRInfo: PR一覧
//   - error: エラー
func GetMergedPRs(owner, repo, head string) ([]PRInfo, error) {
	// TODO: 実装する
	return nil, nil
}

// CreateDraftPR は Draft PR を作成し、PRのURLを返す
//
// TODO: 自分で実装するパート
//
// やること:
//
//	gh pr create --repo {owner}/{repo} --base {base} --head {head}
//	  --title {title} --body {body} --draft
//	を os/exec で呼び出す
//
// 戻り値:
//   - string: 作成されたPRのURL
//   - error: エラー
func CreateDraftPR(owner, repo, base, head, title, body string) (string, error) {
	// TODO: 実装する
	return "", nil
}

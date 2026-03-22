package cmd

import (
	"fmt"
	"os"

	"github.com/manato-takahashi/relech/internal/config"
	"github.com/manato-takahashi/relech/internal/github"

	// "github.com/manato-takahashi/relech/internal/github" // TODO: 実装時にコメント解除
	"github.com/spf13/cobra"
)

var shortOutput bool

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "リポジトリの差分をチェックする",
	Run:   runCheck,
}

func init() {
	checkCmd.Flags().BoolVar(&shortOutput, "short", false, "省略表示（サマリーのみ）")
	rootCmd.AddCommand(checkCmd)
}

func runCheck(cmd *cobra.Command, args []string) {
	cfg, err := config.Load(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "設定ファイルの読み込みに失敗: %v\n", err)
		os.Exit(1)
	}

	needRelease := 0
	for _, repo := range cfg.Repositories {
		result, err := github.Compare(repo.Owner, repo.Name, repo.Base, repo.Head)
		if err != nil {
			fmt.Fprintf(os.Stderr, " %s: エラー: %v\n", repo.Name, err)
			continue
		}
		if result.AheadBy == 0 {
			fmt.Printf("%s no changes\n", repo.Name)
		} else {
			needRelease++
			mark := ""
			if result.AheadBy == len(result.PRs) {
				mark = "\033[32m✓\033[0m"
			} else {
				mark = "\033[33m⚠\033[0m"
			}
			fmt.Printf("%s  %d commits / %d PRs %s\n", repo.Name, result.AheadBy, len(result.PRs), mark)
		}
		if !shortOutput {
			for _, pr := range result.PRs {
				fmt.Printf("  - #%d %s\n", pr.Number, pr.Title)
			}
		}
		if !shortOutput && result.AheadBy > 0 {
			fmt.Println()
		}
	}
	fmt.Println()
	fmt.Printf("Summary: %d/%d repositories need release\n", needRelease, len(cfg.Repositories))
}

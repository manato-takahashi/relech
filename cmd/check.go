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

	for _, repo := range cfg.Repositories {
		result, err := github.Compare(repo.Owner, repo.Name, repo.Base, repo.Head)
		if err != nil {
			fmt.Fprintf(os.Stderr, " %s: エラー: %v\n", repo.Name, err)
			continue
		}
		fmt.Printf("%s  %d commits / %d PRs\n", repo.Name, result.AheadBy, len(result.PRs))
	}
}

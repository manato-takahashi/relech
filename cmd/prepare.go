package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/manato-takahashi/relech/internal/config"
	"github.com/manato-takahashi/relech/internal/github"

	// "github.com/manato-takahashi/relech/internal/github" // TODO: 実装時にコメント解除
	"github.com/spf13/cobra"
)

var prepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "差分のあるリポジトリにDraft PRを作成する",
	Run:   runPrepare,
}

func init() {
	rootCmd.AddCommand(prepareCmd)
}

func runPrepare(cmd *cobra.Command, args []string) {
	cfg, err := config.Load(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "設定ファイルの読み込みに失敗: %v\n", err)
		os.Exit(1)
	}

	needRelease := 0
	fmt.Println("Creating draft PRs...")
	fmt.Println()
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
			title := fmt.Sprintf("本番リリース(%s)", time.Now().Format("2006-01-02"))
			body := ""
			for _, pr := range result.PRs {
				body += fmt.Sprintf("- #%d\n", pr.Number)
			}
			url, err := github.CreateDraftPR(repo.Owner, repo.Name, repo.Base, repo.Head, title, body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "  %s: PR作成時エラー: %v\n", repo.Name, err)
				continue
			}
			fmt.Printf("%s  %s\n", repo.Name, url)
		}
	}
	fmt.Println()
	fmt.Printf("Created %d draft PRs\n", needRelease)
}

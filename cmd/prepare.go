package cmd

import (
	"fmt"
	"os"

	"github.com/manato-takahashi/relech/internal/config"
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

	// ============================================================
	// TODO: ここからがコアロジック（自分で実装するパート）
	// ============================================================
	//
	// 1. cfg.Repositories をループして差分チェック（check と同様）
	//    - 差分がないリポジトリはスキップ
	//
	// 2. 差分があるリポジトリに Draft PR を作成
	//    - github.CreateDraftPR() を呼ぶ
	//    - PRタイトル: cfg.PRTemplate.Title のテンプレートに日付を埋める
	//    - PR本文: 差分PRタイトル一覧を含める
	//
	// 3. 作成したPRのリンク一覧を表示
	//    - "Created N draft PRs" のサマリーも出力
	//
	// ============================================================

	_ = cfg // 使ったら消してね
}

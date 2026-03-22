package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "relech",
	Short: "Release Checker - リリース前の差分チェックツール",
	Long:  `relech は指定リポジトリの develop → main(master) 間の差分を一括チェックし、Draft PR の作成まで行うCLIツールです。`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "config.yaml", "設定ファイルのパス")
}

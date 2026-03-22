package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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

func defaultConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "config.yaml"
	}
	return filepath.Join(home, ".config", "relech", "config.yaml")
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", defaultConfigPath(), "設定ファイルのパス")
}

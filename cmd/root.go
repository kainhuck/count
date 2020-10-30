package cmd

import (
	"count/setting"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"count/core"
)

var ignoreHide bool
var suffix string
var ignoreSpaceLine bool
var showInfo bool
var pwd string

var rootCmd = &cobra.Command{
	Use:   "count",
	Short: "统计行数",
	Long:  "对代码项目进行行数的统计",
	Run: func(cmd *cobra.Command, args []string) {
		var specifiedSuffix []string
		if suffix != ""{
			specifiedSuffix = strings.Split(suffix, ",")
		}

		var set = setting.New(ignoreHide, specifiedSuffix, ignoreSpaceLine, showInfo)
		count := core.Count(pwd, set)
		fmt.Printf("[%s -> total lines]: %d\n", pwd, count)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&ignoreHide, "hide", "i", true, "是否忽略隐藏文件")
	rootCmd.Flags().StringVarP(&suffix, "suffix", "s", "", "请指定后缀，省略查找全部")
	rootCmd.Flags().BoolVarP(&ignoreSpaceLine, "empty", "e", false, "是否忽略空行")
	rootCmd.Flags().BoolVarP(&showInfo, "view", "v", false, "是否显示过程")
	rootCmd.Flags().StringVarP(&pwd, "path", "p", ".", "指定项目路径")
}

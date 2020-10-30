package cmd

import (
	"count/core"
	"count/file"
	"count/setting"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var hide bool
var suffix string
var empty bool
var view bool
var pwd string
var but string

var rootCmd = &cobra.Command{
	Use:   "count",
	Short: "统计行数",
	Long:  "对代码项目进行行数的统计",
	Run: func(cmd *cobra.Command, args []string) {
		var specifiedSuffix []string
		if suffix != "" {
			specifiedSuffix = strings.Split(suffix, ",")
		}
		var butSuffix []string
		if but != "" {
			butSuffix = strings.Split(but, ",")
		}

		var set = setting.New(hide, specifiedSuffix, empty, view, butSuffix)
		count := core.Count(&file.Folder{
			FullPath: pwd,
			IsHide: false,
		}, set)
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
	rootCmd.Flags().BoolVarP(&hide, "hide", "i", true, "是否忽略隐藏文件")
	rootCmd.Flags().StringVarP(&suffix, "suffix", "s", "", "请指定后缀，省略查找全部")
	rootCmd.Flags().BoolVarP(&empty, "empty", "e", false, "是否忽略空行")
	rootCmd.Flags().BoolVarP(&view, "view", "v", false, "是否显示过程")
	rootCmd.Flags().StringVarP(&pwd, "path", "p", ".", "指定项目路径")
	rootCmd.Flags().StringVarP(&but, "but", "b", "exe", "指定不统计的文件后缀")
}

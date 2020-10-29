package main

import (
	"count/file"
	"count/setting"
	"flag"
	"fmt"
	"strings"
)

func main() {
	var ignoreHide bool
	var suffix string
	var ignoreSpaceLine bool
	var showInfo bool
	var pwd string

	flag.BoolVar(&ignoreHide, "hide", true, "是否忽略隐藏文件 true | false, 默认（true）忽略隐藏文件")
	flag.StringVar(&suffix, "suffix", "", "指定要统计的文件后缀，用 英文逗号 分隔, 默认空（全部文件）")
	flag.BoolVar(&ignoreSpaceLine, "spaceline", false, "是否要忽略文件中的空行 true | false, 默认（false）不忽略")
	flag.BoolVar(&showInfo, "info", false, "是否要显示过程 true | false, 默认（false）不显示")
	flag.StringVar(&pwd, "path", ".", "要统计的项目目录，默认当前目录")

	flag.Parse()

	var specifiedSuffix []string
	if suffix != ""{
		specifiedSuffix = strings.Split(suffix, ",")
	}

	var set = setting.New(ignoreHide, specifiedSuffix, ignoreSpaceLine, showInfo)
	count := Count(pwd, set)

	fmt.Printf("[total lines]: %d\n", count)
}

func Count(pwd string, set *setting.Setting) int {
	fileList, folderList := file.List(pwd)

	count := CountFileList(fileList, set)

	for _, folder := range folderList {
		count += Count(folder.FullPath, set)
	}

	return count
}

func CountFileList(fileList []*file.File, set *setting.Setting) (count int) {
	for _, f := range fileList {
		count += file.Count(f, set)
	}

	return count
}

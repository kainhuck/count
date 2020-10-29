package main

import (
	"count/file"
	"count/setting"
	"fmt"
)

func main() {
	var set = setting.NewDefault()
	count := Count("C:\\Users\\kainhuck\\Documents\\github\\gin_slim", set)

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

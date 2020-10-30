package core

import (
	"count/file"
	"count/setting"
)

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

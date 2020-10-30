package core

import (
	"count/file"
	"count/setting"
)

func Count(f *file.Folder, set *setting.Setting) int {
	if set.IgnoreHide && f.IsHide{
		return 0
	}
	fileList, folderList := file.List(f.FullPath)

	count := CountFileList(fileList, set)

	for _, folder := range folderList {
		count += Count(folder, set)
	}

	return count
}

func CountFileList(fileList []*file.File, set *setting.Setting) (count int) {
	for _, f := range fileList {
		count += file.Count(f, set)
	}

	return count
}

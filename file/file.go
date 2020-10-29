package file

import (
	"bufio"
	"count/setting"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type File struct {
	Name     string
	FullPath string
	IsHide   bool
	Suffix   string
}

type Folder struct {
	Name     string
	FullPath string
	IsHide   bool
}

// 列出指定路径下的文件和文件夹
func List(pwd string) (fileList []*File, folderList []*Folder) {
	resources, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}

	for _, each := range resources {
		isHide := string([]rune(each.Name())[0]) == "."
		if each.IsDir() {
			folderList = append(folderList, &Folder{
				Name:     each.Name(),
				FullPath: path.Join(pwd, each.Name()),
				IsHide:   isHide,
			})
		} else {
			fileList = append(fileList, &File{
				Name:     each.Name(),
				FullPath: path.Join(pwd, each.Name()),
				IsHide:   isHide,
				Suffix:   ext(each.Name()),
			})
		}
	}

	return
}

// 取后缀
func ext(pwd string) string {
	tmp := strings.Split(pwd, ".")
	if len(tmp) >= 2 && tmp[0] != "" {
		return strings.ToLower(tmp[len(tmp)-1])
	}
	return ""
}

// 统计行数
func Count(file *File, set *setting.Setting) int {
	if (set.IgnoreHide && file.IsHide) || (set.SpecifiedSuffix != nil && !in(set.SpecifiedSuffix, file.Suffix)) {
		return 0
	}
	count, _ := readLine(file.FullPath, set.IgnoreSpaceLine)
	if set.ShowInfo{
		fmt.Printf("[%s]: %d\n", file.FullPath, count)
	}
	return count
}

// 逐行读取文件
func readLine(fileName string, ignoreSpaceLine bool) (int, error) {
	count := 0
	f, err := os.Open(fileName)
	if err != nil {
		return count, err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if ignoreSpaceLine {
			if len(line) != 0 {
				count++
			}
		} else {
			count++
		}
		if err != nil {
			if err == io.EOF {
				return count, nil
			}
			return count, err
		}
	}
}

func in(strings []string, str string) bool {
	for _, each := range strings {
		if each == str {
			return true
		}
	}

	return false
}

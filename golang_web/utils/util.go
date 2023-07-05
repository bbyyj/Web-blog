package utils

import (
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// 提取路径中的路径名和后缀
func FileSuffixSplit(filepath string) (string, string) {
	for i := 0; i < len(filepath); i++ {
		if filepath[i] == '.' {
			return filepath[:i], filepath[i:]
		}
	}

	return filepath, ""
}

func GetFileSize(filename string) (int, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return int(fi.Size()), nil
}

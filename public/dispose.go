package public

import (
	"os"
	"strings"
)

//提取的字符串切片
func SS(s [][]string) [][]string {
	f := func(c rune) bool {
		if c == '<' || c == '"' || c == '>' {
			return true
		} else {
			return false
		}
	}
	var erweis [][]string
	for _, ss := range s {
		str := strings.FieldsFunc(ss[1], f) //按自定义函数进行拆分得到的字符切片
		erweis = append(erweis, str)
	}
	return erweis
}

// 判断文件（小说）是否存在
func FilePath(path string) (*os.File, error) {
	_, err := os.Stat(path)
	var file *os.File
	if err == nil {
		file, err = os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
		return file, nil
	}
	if os.IsNotExist(err) {
		file, err = os.Create(path)
		return file, nil
	}
	return nil, err
}

//去除空格
func DeleteSpare(ss string) string {
	ss = strings.Replace(ss, " ", "", -1)  //去空格
	ss = strings.Replace(ss, "\n", "", -1) //去换行符
	ss = strings.Replace(ss, "\r", "", -1)
	return ss
}

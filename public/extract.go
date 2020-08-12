package public

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"pachong/configure"
	"pachong/modal"
)

////输入书籍链接获取目录地址
func BookUrl(url string) (s string) {
	result, err := HttpGet(url)
	if err != nil {
		//fmt.Println("HttpGet", err)
		configure.SugaredLogger.Errorf("BookUrl HttpGet: Error = %s", err)
		return
	}
	result = DeleteSpare(result)
	ret := regexp.MustCompile(`<aclass="all-catalog"target="_blank"href="(.*?)"`) //提取目录链接
	shu := ret.FindAllStringSubmatch(result, 1)
	s = shu[0][1]
	return
}

//输入目录链接获取所有章节链接
func BookDir(url string) []*modal.Zhang {
	result, err := HttpGet(url)
	if err != nil {
		//fmt.Println("HttpGet", err)
		configure.SugaredLogger.Errorf("BookUrl HttpGet: Error = %s", err)
		return nil
	}
	result = DeleteSpare(result)
	ret := regexp.MustCompile(`<liclass="col-4">(?s:(.*?))</a>`) //提取章节
	shu := ret.FindAllStringSubmatch(result, -1)
	ss := SS(shu)
	var zhangs []*modal.Zhang
	for _, v := range ss {
		zhang := &modal.Zhang{}
		if v[0] == "aclass=" {
			zhang.Link = v[3]
			zhang.ZhangName = v[8]
		} else {
			zhang.Link = v[1]
			zhang.ZhangName = v[6]
		}
		zhangs = append(zhangs, zhang)
	}
	return zhangs
}

//输入章节链接保存章节内容
func ZhangJ(url string) string {
	result, err := HttpGet(url)
	if err != nil {
		//fmt.Println("HttpGet", err)
		configure.SugaredLogger.Errorf("ZhangJ HttpGet: Error = %s", err)
		return ""
	}
	ret := regexp.MustCompile(`<div class="content" itemprop="acticleBody">(?s:(.*?))</div>`)
	shu := ret.FindAllStringSubmatch(result, 1)
	str := DeleteSpare(shu[0][1])
	str = strings.Replace(shu[0][1], "</p>", "\n\r", -1)
	str = str + "\r\n爬取下一章\r\n"
	return str
}

//book结构体要有目录链接和书名 ，爬取小说章节
func Xiaoshuo(book *modal.Paqu) error {
	zhang := BookDir(book.DirAddress) //获取章节链接
	fmt.Printf("共有%d章\n", len(zhang))
	fmt.Println("请输入爬取起始章节(一次爬取10章)")
	var ii int
	fmt.Scan(&ii)
	jj := ii + 9
	for {
		if len(zhang) >= jj && ii > 0 {
			break
		}
		fmt.Println("输入有误，请重新输入")
		fmt.Scan(&ii)
	}
	path := fmt.Sprintf("../小说/%s.txt", book.Title)
	fmt.Println(path)
	file, err := FilePath(path)
	defer file.Close()
	if err != nil {
		//fmt.Print("Xiaoshuo FilePath err:", err)
		configure.SugaredLogger.Errorf("Xiaoshuo FilePath: Error = %s", err)
		return err
	}
	for i := (ii - 1); i < jj; i++ {
		fmt.Printf("%d...", i+1)
		shu := zhang[i]
		str := ZhangJ(shu.Link)
		_, err := file.Write([]byte(shu.ZhangName))
		_, err = file.Write([]byte(str))
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 5)
	}
	book.FileLocation = path
	book.CreateTime = time.Now()
	book.UpdateTime = book.CreateTime
	book.AddLatest = strconv.Itoa(jj)
	err = book.Add()
	if err != nil {
		//fmt.Println("book add err", err)
		configure.SugaredLogger.Errorf("book add er: Error = %s", err)
	}
	return nil
}

package page

import (
	"fmt"
	"pachong/configure"
	"pachong/modal"
	"pachong/public"
	"regexp"
	"strconv"
)

//书库爬取
func ToWork(start int) {
	fmt.Printf("正在爬取第 %d 。\n", start)
	shuku := SpiderPage(start)
	if len(shuku) == 0 {
		return
	}
	fmt.Println("请输入要爬取的书籍编号（第几本）")
	var ii int
	fmt.Scan(&ii)
	for {
		if ii > 0 && ii <= len(shuku) {
			break
		}
		fmt.Println("输入有误，请重新输入")
		fmt.Scan(&ii)
	}
	book := shuku[(ii - 1)]
	book.DirAddress = public.BookUrl(book.BookAddress) //输入书籍链接获取目录地址
	err := public.Xiaoshuo(book)
	//book.AddLatest = string(n)
	if err == nil {
		fmt.Printf("\n%s 爬取成功\n", book.Title)
	} else {
		fmt.Printf("\n%s 爬取失败\n", book.Title)
	}
}

//抓取书库页内容
func SpiderPage(i int) []*modal.Paqu {
	url := "http://book.zongheng.com/store/c0/c0/b0/u0/p" + strconv.Itoa(i) + "/v9/s9/t0/u0/i1/ALL.html"
	result, err := public.HttpGet(url) //获取书库url的全部内容
	if err != nil {
		//fmt.Println("HttpGet", err)
		configure.SugaredLogger.Errorf("SpiderPage HttpGet: Error = %s", err)
		return nil
	}
	result = public.DeleteSpare(result)                               //去除空格
	ret1 := regexp.MustCompile(`<divclass="bookname">(?s:(.*?))</a>`) //书籍链接和书名
	shu1 := ret1.FindAllStringSubmatch(result, -1)
	ret2 := regexp.MustCompile(`<divclass="bookilnk">(?s:(.*?))<span>`) //作者/书籍分类
	shu2 := ret2.FindAllStringSubmatch(result, -1)
	ret3 := regexp.MustCompile(`最新章节：(?s:(.*?))</a>`) //最新章节
	shu3 := ret3.FindAllStringSubmatch(result, -1)
	s1 := public.SS(shu1)
	s2 := public.SS(shu2)
	var shuku []*modal.Paqu
	for i, v := range s1 {
		paqu := &modal.Paqu{
			Title:        v[4],
			Author:       s2[i][4],
			BookClassify: s2[i][11],
			Latest:       shu3[i][1],
			BookAddress:  v[1],
		}
		shuku = append(shuku, paqu)
		fmt.Printf("第%d本 链接：%s 书名：%s 作者：%s 类型：%s 最新章节：%s\n", i+1, v[1], v[4], s2[i][4], s2[i][11], shu3[i][1])
	}
	return shuku
}

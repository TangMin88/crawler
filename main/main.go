package main

import (
	"fmt"
	"pachong/configure"
	"pachong/page"
)

func main() {
	configure.InitLogger()
	defer configure.SugaredLogger.Sync()
	var start int
	for {
		fmt.Print("\n--------1 书库爬取 --------\n")
		fmt.Print("--------2   退出   --------\n")
		fmt.Scan(&start)
		switch start {
		case 1:
			var ii int
			fmt.Print("请输入爬取页(>=1)\n")
			fmt.Scan(&ii)
			page.ToWork(start)
		case 2:
			fmt.Println("退出程序")
			return
		default:
			fmt.Println("输入有误，请重新输入")
			fmt.Scan(&start)
		}
	}

}

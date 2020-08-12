package modal

import (
	"pachong/db"
	"time"
)

type Paqu struct {
	ID           int64
	Title        string    //书名 //
	Author       string    //作者 //
	Latest       string    `gorm:"-" ` //最新章节 //
	AddLatest    string    //保存的最新章节 //
	BookClassify string    //书籍分类 //
	DirAddress   string    //目录网页地址 //
	BookAddress  string    //书籍链接地址 //
	CreateTime   time.Time //创建时间 //
	UpdateTime   time.Time //更新时间 //
	FileLocation string    //文件位置 //
	Zhang        []*Zhang  `gorm:"-"` //章节切片

}

func (p *Paqu) Add() error {
	return db.Db.Create(p).Error
}

// func (p *Paqu)Query(title string) []*Paqu{

// }

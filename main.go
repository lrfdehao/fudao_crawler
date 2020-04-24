package main

import (
	"crawler/worker"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 初始化db
	db, err := gorm.Open("mysql", "root:root@/fudao?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)

	// 获取年级课程信息
	gsWorker := worker.NewGradeSubjWorker()
	gsWorker.Proccess()

	// 处理课程信息
	scWorker := worker.NewSubjectCourseWorker(db)
	scWorker.Proccess(gsWorker.Data)

	// 处理课程包数据
	//pcWorker := worker.NewPackageCourseWorker(db)
	//pcWorker.Proccess(scWorker.PkgIdList)

	// 处理课程详情
	cdWorker := worker.NewCourseDetailWorker(db)
	cdWorker.Proccess(scWorker.Cids)

}

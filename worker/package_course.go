package worker

import (
	"crawler/model"
	"crawler/util"
	"encoding/json"
	"log"

	"github.com/jinzhu/gorm"
)

type PackageCourseWorker struct {
	db *gorm.DB
}

func NewPackageCourseWorker(db *gorm.DB) *PackageCourseWorker {
	return &PackageCourseWorker{db: db}
}

func (w *PackageCourseWorker) Proccess(pkgIdList []string) {
	for _, pkgId := range pkgIdList {
		resp := util.GetPackageCourseInfo(pkgId)
		packageCourse := &model.PackageCourseInfo{}
		err := json.Unmarshal(resp, packageCourse)
		if err != nil {
			log.Printf("GetPackageCourseInfo error[%v], package_id[%v]", err, pkgId)
			return
		}

		util.SaveCourseToDB(packageCourse.Result.Courses, w.db)
	}

}

package worker

import (
	"crawler/model"
	"crawler/util"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

type SubjectCourseWorker struct {
	db   *gorm.DB
	Cids []int64
}

func NewSubjectCourseWorker(db *gorm.DB) *SubjectCourseWorker {
	return &SubjectCourseWorker{db: db}
}

func (w *SubjectCourseWorker) Proccess(gsData *model.GradeSubjResp) {
	var cidList []int64
	for _, gs := range gsData.Result.GradeSubjects {
		g := gs.Grade
		for _, s := range gs.Subject {
			// 处理一组grade,subject
			log.Printf("ready to request course, grade[%v], subject[%v]", g, s)
			resp := util.GetDiscoverSubjResp(g, s, 1, 0)
			subj := &model.DiscoverSubjResp{}
			err := json.Unmarshal(resp, subj)
			if err != nil {
				log.Printf("GetDiscoverSubjResp error[%v], grade[%v], subject[%v]", err, g, s)
				return
			}

			// 添加所有cid
			for _, pkg := range subj.Result.SysCoursePkgList {
				cidsStr := strings.Split(pkg.CidList, ",")
				for _, cidStr := range cidsStr {
					cid, _ := strconv.ParseInt(cidStr, 10, 64)
					cidList = append(cidList, cid)
				}
			}

			for _, cData := range subj.Result.SpeCourseList.Data {
				cidList = append(cidList, cData.Cid)
			}

			//log.Printf("ready to save db, grade[%v], subject[%v]", g, s)
			//util.SaveCourseToDB(subj.Result.SpeCourseList.Data, w.db)
		}
	}

	log.Printf("total cid len[%v], cidlist[%v]", len(cidList), cidList)
	w.Cids = cidList

}

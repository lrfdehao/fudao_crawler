package worker

import (
	"crawler/model"
	"crawler/util"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type CourseDetailWorker struct {
	db *gorm.DB
}

func NewCourseDetailWorker(db *gorm.DB) *CourseDetailWorker {
	return &CourseDetailWorker{db: db}
}

func (w *CourseDetailWorker) Proccess(cidList []int64) {

	for _, cid := range cidList {
		resp := util.GetCourseDetail(cid)
		cDetail := &model.CourseDetailResp{}
		err := json.Unmarshal(resp, cDetail)
		if err != nil {
			log.Printf("GetCourseDetail error[%v], course_id[%v]", err, cid)
			return
		}

		// 保存课程信息到db
		w.saveCourseInfo(cDetail)
		// 保存教师信息到db
		w.saveTeacherInfo(cDetail)
		// 保存目录信息到db
		w.saveDirector(cDetail)
	}
}

func (w *CourseDetailWorker) saveDirector(cDetail *model.CourseDetailResp) {
	var tid int64
	if len(cDetail.Result.Teacher) > 0 {
		tid = cDetail.Result.Teacher[0].Tid
	}
	for _, d := range cDetail.Result.Directory {
		director := &model.CourseDirectory{}
		w.db.Where("csid = ?", d.Csid).First(director)
		director.Csid = d.Csid
		director.CsType = d.CsType
		director.ExamDuration = d.ExamDuration
		director.Title = d.Title
		director.Tid = tid
		director.BeginTime = time.Unix(d.Bgtime, 0)
		director.EndTime = time.Unix(d.Endtime, 0)
		director.Cid = cDetail.Result.Cid
		grade, err := strconv.ParseInt(cDetail.Result.Grade, 10, 64)
		if err == nil {
			director.Grade = grade
		}
		director.Subject = cDetail.Result.Subject
		err = w.db.Save(director).Error
		if err != nil {
			log.Print("SaveCourseToDB error ", err)
		} else {
			log.Printf("save director success, name[%v], csid[%v]", d.Title, d.Csid)
		}
	}
}

func (w *CourseDetailWorker) saveTeacherInfo(cDetail *model.CourseDetailResp) {
	for _, t := range cDetail.Result.Teacher {
		teacher := &model.Teacher{}
		w.db.Where("tid = ?", t.Tid).First(teacher)
		teacher.Desc = t.Desc
		teacher.Name = t.Name
		teacher.Pic = t.Pic
		teacher.Tid = t.Tid
		err := w.db.Save(teacher).Error
		if err != nil {
			log.Print("SaveCourseToDB error ", err)
		} else {
			log.Printf("save teacher success, name[%v], tid[%v]", t.Name, t.Tid)
		}
	}
}

func (w *CourseDetailWorker) saveCourseInfo(cDetail *model.CourseDetailResp) {
	course := &model.CourseInfo{}
	w.db.Where("cid = ?", cDetail.Result.Cid).First(course)
	result := cDetail.Result
	course.Cid = result.Cid
	course.CourseName = result.Name
	course.PreAmount = result.Price
	course.Subject = result.Subject
	grade, err := strconv.ParseInt(result.Grade, 10, 64)
	if err == nil {
		course.Grade = grade
	}

	course.BeginTime = time.Unix(result.BgTime, 0)
	course.EndTime = time.Unix(result.EndTime, 0)

	err = w.db.Save(course).Error
	if err != nil {
		log.Print("SaveCourseToDB error ", err)
	} else {
		log.Printf("save course success, name[%v], cid[%v]", course.CourseName, course.Cid)
	}
}

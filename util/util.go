package util

import (
	"crawler/constant"
	"crawler/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

func SaveCourseToDB(speCourseList []model.SpeCourse, db *gorm.DB) {
	for _, speCourse := range speCourseList {
		course := &model.CourseInfo{}
		err := db.Where("cid = ?", speCourse.Cid).First(course).Error

		course.Cid = speCourse.Cid
		course.CourseName = speCourse.Name
		course.ApplyNum = speCourse.ApplyNum
		course.PreAmount = speCourse.PreAmount
		course.AfAmount = speCourse.AfAmount
		course.Grade = speCourse.Grade
		course.Subject = speCourse.Subject
		course.BeginTime = time.Unix(speCourse.FirstSubBgtime, 0)
		course.EndTime = time.Unix(speCourse.LastSubEndtime, 0)

		err = db.Save(course).Error
		if err != nil {
			log.Print("saveCourseToDB error ", err)
		} else {
			log.Printf("save course success, name[%v], cid[%v]", course.CourseName, course.Cid)
		}
	}
}

// GetGradeSubjResp 获取grade_subject接口数据
func GetGradeSubjResp() []byte {
	request, err := http.NewRequest("GET", constant.API_GRADE_SUBJECT, nil)
	if err != nil {
		return nil
	}
	request.Header.Add("referer", "https://fudao.qq.com")

	return sendRequest(request)
}

// GetPackageCourseInfo 获取get_course_package_info接口数据
func GetPackageCourseInfo(pkg_id string) []byte {
	request, err := http.NewRequest("GET", constant.API_GET_COURSE_PACKAGE, nil)
	if err != nil {
		return nil
	}

	request.Header.Add("referer", "https://fudao.qq.com")
	q := request.URL.Query()
	q.Add("subject_package_id", pkg_id)
	request.URL.RawQuery = q.Encode()

	return sendRequest(request)
}

// GetCourseDetail 获取courseStaticDetail接口数据
func GetCourseDetail(cid int64) []byte {
	request, err := http.NewRequest("GET", constant.API_COURSE_DETAIL, nil)
	if err != nil {
		return nil
	}

	request.Header.Add("referer", "https://fudao.qq.com")
	q := request.URL.Query()
	q.Add("cid", strconv.FormatInt(cid, 10))
	request.URL.RawQuery = q.Encode()

	return sendRequest(request)
}

// GetDiscoverSubjResp 获取discover_subject接口数据
func GetDiscoverSubjResp(grade, subject int64, page, size int) []byte {
	request, err := http.NewRequest("GET", constant.API_DISCOVER_SUBJECT, nil)
	if err != nil {
		return nil
	}

	request.Header.Add("referer", "https://fudao.qq.com")
	q := request.URL.Query()
	q.Add("grade", strconv.FormatInt(grade, 10))
	q.Add("subject", strconv.FormatInt(subject, 10))
	q.Add("page", strconv.Itoa(page))
	q.Add("size", strconv.Itoa(size))
	q.Add("showid", "0")
	request.URL.RawQuery = q.Encode()

	return sendRequest(request)
}

// sendRequest common send request func
func sendRequest(request *http.Request) []byte {
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Print("http request error ", err)
		return nil
	}

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("http parse body error ", err)
		return nil
	}

	return bodyData
}

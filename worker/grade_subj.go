package worker

import (
	"crawler/model"
	"crawler/util"
	"encoding/json"
	"log"
)

type GradeSubjWorker struct {
	Data *model.GradeSubjResp
}

func NewGradeSubjWorker() *GradeSubjWorker {
	return &GradeSubjWorker{}
}

func (w *GradeSubjWorker) Proccess() {
	resp := util.GetGradeSubjResp()
	gradeSubj := &model.GradeSubjResp{}
	err := json.Unmarshal(resp, gradeSubj)
	if err != nil {
		log.Fatal(err)
		return
	}

	if gradeSubj.Retcode != 0 {
		log.Fatalf("GradeSubjWorker error retcode: %v", gradeSubj.Retcode)
	}

	w.Data = gradeSubj
}

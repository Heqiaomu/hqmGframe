package entity

import "time"

// Work 作业列表
type Work struct {
	Model
	WorkName       string `json:"work_name" gorm:"work_name"`             // 课程作业名称
	WorkDetail     string `json:"work_detail" gorm:"work_detail"`         // 课程作业详情
	ClassID        int    `json:"class_id" gorm:"class_id"`               //  关联班级
	CourseID       int    `json:"course_id" gorm:"course_id"`             // 关联课程
	PublishTeacher int    `json:"publish_teacher" gorm:"publish_teacher"` // 发布老师
}

func (w Work) TableName() string {
	return "tb_work"
}

// SubmitWork 作业提交列表
type SubmitWork struct {
	Model
	WorkID         int       `json:"work_id" gorm:"work_id"`                                                 // 关联作业ID
	WorkName       string    `json:"work_name" gorm:"work_name"`                                             // 作业名
	ClassID        int       `json:"class_id" gorm:"class_id"`                                               // 关联班级
	Submitter      int       `json:"submitter" gorm:"submitter"`                                             // 提交人ID
	SubmitVersion  string    `json:"submit_version" gorm:"submit_version"`                                   // 提交版本
	SubmitLastTime time.Time `json:"submit_last_time" gorm:"submit_last_time,default:'2022-11-11 00:00:00'"` // 最后一次提交时间
}

func (sw SubmitWork) TableName() string {
	return "tb_submit_work"
}

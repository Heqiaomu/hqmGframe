package request

type Class struct {
	ClassName string `json:"c_lass_name" form:"class_name"`
	Description string `json:"description" form:"description"`
	ClassLeader int `json:"class_leader" form:"class_leader"`
}

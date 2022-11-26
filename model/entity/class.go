package entity

// 班级信息

type Class struct {
	Model
	ClassName   string `json:"class_name" gorm:"class_name"` // 班级名
	Description string `json:"description" gorm:"description"`
	ClassLeader int    `json:"class_leader" gorm:"class_leader"`
}

func (c Class) TableName() string {
	return "tb_class"
}

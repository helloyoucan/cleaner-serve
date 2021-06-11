package models

type BranchQuery struct {
	Name string `form:"name"`
	Province string `form:"province"`
	City     string `form:"city"`
	Area     string `form:"area"`
	ContactPerson  string `form:"contact_person"` //联系人
	Status uint8 `form:"status"` // 网点状态 0关闭，1营业中，2休息中
	CreatedStartTime uint `form:"created_start_time"`
	CreatedEndTime uint `form:"created_end_time"`
}

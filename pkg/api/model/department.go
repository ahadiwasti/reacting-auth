package model

type Department struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	OrderNum    int    `json:"order_num"`
	ParentId    int    `json:"parent_id"`
	ExtendField string `json:"extend_field"`
}

func (Department) TableName() string {
	return "department"
}

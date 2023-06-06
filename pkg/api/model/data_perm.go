package model

type DataPerm struct {
	Id        int    `json:"id"`
	ParentId  int    `json:"parent_id"`
	Name      string `json:"name"`
	Perms     string `json:"perms"`
	PermsRule string `json:"perms_rule"`
	PermsType int    `json:"perms_type"`
	OrderNum  int    `json:"order_num"`
	DomainId  int    `json:"domain_id"`
	Remarks   string `json:"remarks"`
}

type DataPermQuery struct {
	DomainId   int
	Name       string
	Pagination *Pagination
}

func (dp *DataPerm) TableName() string {
	return "data_perm"
}

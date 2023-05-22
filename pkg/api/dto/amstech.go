package dto


// Amsrequest - dto for payload creation
type AmsPayload struct {
	Service        string `form:"service" json:"service" binding:"required"`
	Method        string `form:"method" json:"method" binding:"required"`
	Request        string `form:"request" json:"request" binding:"required"`
}
type CRYPTODTO struct {
	Content string `form:"content" json:"content" binding:"content"`
	Iv  string `form:"iv" json:"iv" binding:"iv"`
}

type ErpDtoPayload struct {
	Item        string `form:"item" json:"item" binding:"required"`
}


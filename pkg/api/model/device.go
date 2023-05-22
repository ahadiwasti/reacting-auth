package model

// Device model
type Device struct {
	iddevice        int        `json:"iddevice" example:"1"`
	devicename      string     `json:"devicename" example:"wutongci"`
	deviceid        string     `json:"deviceid" example:"186000000"`
	deviceserial    string     `json:"deviceserial" example:"1"`
	assettag      	string     `json:"assettag" example:"assettag"`
}

func (Device) TableName() string {
	return "device"
}

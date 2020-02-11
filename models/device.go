package models

// Device device struct
type Device struct {
	DevID      string `json:"id" example:"TD2BDhvjKEaZo37c4DLAq6" xorm:"'dev_id'"`
	DevNm      string `json:"name" example:"device-37" xorm:"'dev_nm'"`
	DevTag     string `json:"tag" example:"tag" xorm:"'dev_tag'"`
	PrtcType   string `json:"prtctype" example:"prtctype" xorm:"'prtc_type'"`
	DevRec     string `json:"rec" example:"비고" xorm:"'dev_rec'"`
	Firmware   string `json:"firmware" example:"firmware version" xorm:"'firmware'"`
	Status     string `json:"status" example:"1" xorm:"'status'"`
	AuthID     string `json:"authid" example:"?" xorm:"'auth_id'"`
	AuthKey    string `json:"authkey" example:"??" xorm:"'auth_key'"`
	SerialNo   string `json:"serialno" example:"102312-12312" xorm:"'serial_no'"`
	ModelID    int    `json:"modelid" example:"4" xorm:"'model_id'"`
	EqupID     int    `json:"equpid" example:"1" xorm:"'equp_id'"`
	SiteID     int    `json:"siteid" example:"1" xorm:"'site_id'"`
	CustID     int    `json:"custid" example:"4" xorm:"'cust_id'"`
	DevType    string `json:"type" example:"??" xorm:"'dev_type'"`
	UseYN      string `json:"useyn" example:"Y" xorm:"'use_yn'"`
	CreatedAt  string `json:"createdat" example:"2020-01-29 13:10:39" xorm:"created 'created_at'"`
	UpdateedAt string `json:"updatedat" example:"2020-01-29 13:10:39" xorm:"updated 'updated_at'"`
}

// TableName 테이블명 지정
func (Device) TableName() string {
	return "device"
}

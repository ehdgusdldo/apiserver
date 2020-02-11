package models

// Model model struct
type Model struct {
	ModelID    int    `json:"id" example:"4" xorm:"'model_id'"`
	ModelNm    string `json:"name" example:"viux-38" xorm:"'model_nm'"`
	ModelTag   string `json:"tag" example:"tag" xorm:"'model_tag'"`
	PrtcType   string `json:"prtctype" example:"prtctype" xorm:"'prtc_type'"`
	PrtcIo     string `json:"prtcio" example:"1" xorm:"'prtc_io'"`
	CustID     int    `json:"custid" example:"4" xorm:"'cust_id'"`
	Firmware   string `json:"firmware" example:"firmware version" xorm:"'firmware'"`
	ModelRec   string `json:"rec" example:"비고" xorm:"'model_rec'"`
	ModelType  string `json:"type" example:"??" xorm:"'model_type'"`
	UseYN      string `json:"useyn" example:"Y" xorm:"'use_yn'"`
	CreatedAt  string `json:"createdat" example:"2020-01-29 13:10:39" xorm:"created 'created_at'"`
	UpdateedAt string `json:"updatedat" example:"2020-01-29 13:10:39" xorm:"updated 'updated_at'"`
}

// TableName 테이블명 지정
func (Model) TableName() string {
	return "model"
}

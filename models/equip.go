package models

// Equip equip struct
type Equip struct {
	EqupID     int    `json:"id" example:"1" xorm:"'equp_id'"`
	EqupNm     string `json:"name" example:"2공장 설비1" xorm:"'equp_nm'"`
	CustID     int    `json:"custid" example:"4" xorm:"'cust_id'"`
	SiteID     int    `json:"siteid" example:"13" xorm:"'site_id'"`
	EqupRec    string `json:"rec" example:"비고" xorm:"'equp_rec'"`
	EqupType   string `json:"type" example:"??" xorm:"'equp_type'"`
	UseYN      string `json:"useyn" example:"Y" xorm:"'use_yn'"`
	CreatedAt  string `json:"createdat" example:"2020-01-29 13:10:39" xorm:"created 'created_at'"`
	UpdateedAt string `json:"updatedat" example:"2020-01-29 13:10:39" xorm:"updated 'updated_at'"`
}

// TableName 테이블명 지정
func (Equip) TableName() string {
	return "equip"
}

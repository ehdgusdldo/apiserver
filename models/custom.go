package models

import (
	"fmt"

	"github.com/ehdgusdldo/APIServer/util"
)

type Custom struct {
	CustID     int    `json:"id" example:"1" xorm:"'cust_id'"`
	CustEID    string `json:"eid" example:"SAMSUNG" xorm:"'cust_eid'"`
	CustNm     string `json:"name" example:"삼성" xorm:"'cust_nm'"`
	PostNo     string `json:"postno" example:"10414" xorm:"'post_no'"`
	BasAdr     string `json:"basadr" example:"서울특별시 강서구 화곡로 12" xorm:"'bas_adr'"`
	DtlAdr     string `json:"dtladr" example:"심당빌딩 3층 302호" xorm:"'dtl_adr'"`
	CustType   int    `json:"type" example:"1" xorm:"'cust_type'"`
	CompNo     string `json:"compno" example:"17282718127" xorm:"'comp_no'"`
	ReptNm     string `json:"reptnm" example:"고영호" xorm:"'rept_nm'"`
	ApprYN     string `json:"appryn" example:"Y" xorm:"'appr_yn'"`
	CustRec    string `json:"custrec" example:"--" xorm:"'cust_rec'"`
	UseYN      string `json:"useyn" example:"Y" xorm:"'use_yn'"`
	CreatedAt  string `json:"createdat" example:"2020-01-29 13:10:39" xorm:"created 'created_at'"`
	UpdateedAt string `json:"updatedat" example:"2020-01-29 13:10:39" xorm:"updated 'updated_at'"`
	OrgID      int    `json:"orgid" example:"2" xorm:"'org_id'"`
}

func (Custom) TableName() string {
	return "custom"
}

// 단일get
func (c *Custom) Get() (bool, error) {

	has, err := util.Engine.Get(c)

	if !has || err != nil {
		fmt.Println(err)
		return false, err
	}

	return true, err
}

// 전체조회
func (c Custom) GetAll() (customs []Custom, err error) {

	err = util.Engine.Find(&customs)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return customs, err
}

// 수정
func (c Custom) update() (rows int, err error) {

	return
}

// 고객추가(spring에서 selectkey 사용하고 ai기능을 안쓰기때문에 현재문제있음)
// affected는 insert된 행갯수 xorm엔 insert된 pk를 반환하는 func이 없는거같음
func (c Custom) Add() (affected int64, err error) {

	affected, err = util.Engine.Insert(&c)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

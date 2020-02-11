package models

import (
	"encoding/json"
	"fmt"

	"github.com/ehdgusdldo/APIServer/util"
	"golang.org/x/crypto/bcrypt"
)

// Users uesrs Struct
type Users struct {
	UserID    int    `json:"id" example:"4" xorm:"'user_id'"`
	Pwd       string `json:"pwd" example:"abcd1234" xorm:"'pwd'"`
	Email     string `json:"email" example:"example@example.com" xorm:"'email'"`
	Name      string `json:"name" example:"강준구" xorm:"'user_nm'"`
	Phone     string `json:"phone" example:"01012341234" xorm:"'phone'"`
	Dept      string `json:"dept" example:"경영지원팀" xorm:"'dept'"`
	Position  string `json:"position" example:"사원" xorm:"'position'"`
	Admin     int    `json:"admin" example:"1" xorm:"'admin'"`
	CustID    int    `json:"custid" example:"23" xorm:"'cust_id'"`
	UserRec   string `json:"rec" example:"비고" xorm:"'user_rec'"`
	UseYN     string `json:"useyn" example:"Y" xorm:"'use_yn'"`
	CreatedAt string `json:"createdat" example:"TIMESTAMP" xorm:"created 'created_at'"`
	UpdatedAt string `json:"updatedat" example:"SAMSUNG" xorm:"updated 'updated_at'"`
	GfID      int    `json:"gfid" example:"12" xorm:"'gf_id'"`
}

// TableName 테이블명 지정
func (Users) TableName() string {
	return "users"
}

// GetAll 전체 사용자목록 select후 반환
func (u Users) GetAll() (users []Users, err error) {

	err = util.Engine.Find(&users)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users, err
}

// Get 단일사용자 조회
func (u *Users) Get() (bool, error) {

	has, err := util.Engine.Get(u)

	if !has || err != nil {
		fmt.Println(err)
		return false, err
	}

	return true, err
}

// Add 사용자추가
func (u Users) Add() (id int64, err error) {
	out, err := json.Marshal(u)
	fmt.Println(string(out))
	pwHash, _ := bcrypt.GenerateFromPassword([]byte(u.Pwd), bcrypt.DefaultCost)

	affected, err := util.Engine.Exec("INSERT INTO users ( user_nm, pwd, phone, email, dept, position, admin, cust_id, user_rec, gf_id) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		u.Name,
		string(pwHash),
		u.Phone,
		u.Email,
		util.NewNullString(u.Dept),
		util.NewNullString(u.Position),
		u.Admin,
		u.CustID,
		util.NewNullString(u.UserRec),
		u.GfID,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	id, err = affected.LastInsertId()

	return
}

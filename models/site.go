package models

// Site Site struct
type Site struct {
	SiteID     int    `json:"id" example:"1" xorm:"'site_id'"`
	SiteNm     string `json:"name" example:"제주1공장" xorm:"'site_nm'"`
	PostNo     string `json:"postno" example:"10232" xorm:"'post_no'"`
	BasAdr     string `json:"basadr" example:"서울특별시 강서구 화곡로 132" xorm:"'bas_adr'"`
	DtlAdr     string `json:"dtladr" example:"심당빌딩 3층 301호" xorm:"'dtl_adr'"`
	Lat        string `json:"lat" example:"32°" xorm:"'lat'"`
	Lon        string `json:"lon" example:"24°" xorm:"'lon'"`
	CustID     int    `json:"custid" example:"1" xorm:"'cust_id'"`
	SiteRec    string `json:"rec" example:"비고" xorm:"'site_rec'"`
	SiteType   string `json:"type" example:"??" xorm:"'site_type'"`
	UseYN      string `json:"useyn" example:"Y" xorm:"'use_yn'"`
	CreatedAt  string `json:"createdat" example:"2020-01-29 13:10:39" xorm:"created 'created_at'"`
	UpdateedAt string `json:"updatedat" example:"2020-01-29 13:10:39" xorm:"updated 'updated_at'"`
}

// TableName 테이블명 지정
func (Site) TableName() string {
	return "site"
}

/* // Get 사이트 단일조회
func (s *Site) Get() (bool, error) {

	has, err := util.Engine.Get(s)

	if !has || err != nil {
		fmt.Println(err)
		return false, err
	}

	return true, err
} */

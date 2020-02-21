package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ehdgusdldo/APIServer/util"
	"github.com/gin-gonic/gin"
)

// TimeValue  Equip:  iotdata5 측정값
type TimeValue struct {
	Time     string `json:"t" example:"2019-02-02"`
	Measures int64  `json:"m" example:"302"`
}

// EquipResponse Equip :  iotdata5 응답용 struct I : 장치아이디 K : 측정키 Value: 측정값리스트
type EquipResponse struct {
	DevID string `json:"i" example:"KyvrjRACQJGd3Q7Q9udzo4"`
	Key   string `json:"k" example:"co2"`
	Value []TimeValue
}

// EquipSeries .. I : 장치아이디 , K : 측정키 , M : 측정값
type EquipSeries struct {
	DevID    string `json:"i" example:"KyvrjRACQJGd3Q7Q9udzo4"`
	Key      string `json:"k" example:"co2"`
	Time     string `json:"t" example:"2019-02-02"`
	Measures int64  `json:"m" example:"302"`
}

// SiteLastDataResponse .. SID : 사이트아이디
type SiteLastDataResponse struct {
	SID      string `json:"SID" example:"site004"`
	LastData []EquipSeries
}

// EventData ..eventdata struct
type EventData struct {
	Time     string `json:"t" example:"2019-02-02"`
	ActKey   string `json:"actkey" example:"acky"`
	ActVal   string `json:"actval" example:"20"`
	CurVal   int64  `json:"curval" example:"812"`
	Etype    string `json:"etype" example:"err"`
	EventKey string `json:"eventkey" example:"co2"`
	Limit    string `json:"limit" example:"50"`
	Model    string `json:"model" example:"14"`
	NotiKey  string `json:"notikey" example:"1"`
	Site     string `json:"site" example:"site004"`
}

// ResponseEventData ..eventdata 응답용 struct
type ResponseEventData struct {
	DevID     string `json:"i" example:"KyvrjRACQJGd3Q7Q9udzo4"`
	EventList []EventData
}

// Equip ..
// @Description iotdata5 조회 limit offset 을 넣지않으면 최근 10건
// @Router /equip [get]
// @Tags equip
// @Accept  json
// @Produce  json
// @Param id query string true "장치아이디"
// @Param key query string true "측정키"
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {object} EquipResponse
func Equip(c *gin.Context) {
	// query string parameter bind
	id := c.Query("id")
	key := c.Query("key")
	limit := c.Query("limit")
	// limit이 파라미터로 안넘어올시 default값설정
	if limit == "" {
		limit = "10"
	}
	offset := c.Query("offset")
	// offset이 파라미터로 안넘어올시 default값설정
	if offset == "" {
		offset = "0"
	}
	log.Println(offset)
	// res, err := util.QueryDB("select * from iotdata5 where i = 'TD2BDhvjKEaZo37c4DLAq6'  and k = 'lx2' order by time desc limit 10;")
	res, err := util.QueryDB("select * from iotdata5 where i = '" + id + "' and k = '" + key + "' order by time desc limit " + limit + " offset " + offset + " ;")
	// queryDB err처리
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return
	}
	// 조회된 데이터가 없을시 500 메세지 반환
	if len(res[0].Series) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "no data check parameter",
		})
		return
	}
	// fmt.Println(res[0].Series[0].Columns)
	// fmt.Println(res[0].Series[0].Name)
	// fmt.Println(res[0].Series[0].Partial)
	// fmt.Println(res[0].Series[0].Tags)

	// res[0].Series[0].Values  Values는 [][]interface{} 로 모든데이터가 들어가있음
	values := res[0].Series[0].Values
	log.Println(values)
	var value []TimeValue
	// Values를 순회하며 응답할 데이터형태로 조립
	// iotdata5 의 values[i]의 구조예시 -> [[2020-02-10T07:27:09.875302157Z] [TD2BDhvjKEaZo37c4DLAq6] [lx2] [521] [site003]]
	for i := 0; i < len(values); i++ {
		// [i][0] 시간
		time := fmt.Sprint(values[i][0])
		// [i][3] 측정값 JSON.Number -> int64
		v, err := values[i][3].(json.Number).Int64()
		// convert err 처리
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed",
			})
			return
		}
		influx := TimeValue{time, v}

		value = append(value, influx)
	}
	// 장치id와 tag값을 0번째데이터에서 가져와 응답 k, i에 셋팅lets incrytp
	k := fmt.Sprint(values[0][2])
	i := fmt.Sprint(values[0][1])
	var response EquipResponse
	response = EquipResponse{i, k, value}

	c.JSON(http.StatusOK, response)

}

// Event ..
// @Description 장치의 최근이벤트 히스토리 조회
// @Router /event [get]
// @Tags equip
// @Accept  json
// @Produce  json
// @Param id query string true "장치아이디""
// @Param limit query int false "limit"
// @Success 200 {object} ResponseEventData
func Event(c *gin.Context) {
	// 장치아이디 query string parameter bind
	id := c.Query("id")
	limit := c.Query("limit")
	// limit이 파라미터로 안넘어올시 default값설정
	if limit == "" {
		limit = "10"
	}
	res, err := util.QueryDB(" select * from eventdata where id='" + id + "' order by time desc limit " + limit + " ;")
	// queryDB err처리
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return
	}
	/*
		index : 											        0              1        2      3       4        5               6              7      8       9        10
		key : 													  Time          actkey   actval  curval  etype  eventkey            id           limit  model   notikey    site
		eventdata measurement의경우 Values[i]의 구조 [2020-02-21T01:58:44.537Z    acky      20     629      ov      hm1   KyvrjRACQJGd3Q7Q9udzo4   50      30       1      site006]
	*/
	log.Println(res[0].Series[0].Values[4])
	// values 를순회하며 응답할 데이터형태로 조립
	var eventList []EventData
	values := res[0].Series[0].Values
	for i := 0; i < len(values); i++ {
		time := fmt.Sprint(values[i][0])
		actkey := fmt.Sprint(values[i][1])
		actval := fmt.Sprint(values[i][2])
		curval, err := values[i][3].(json.Number).Int64()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed",
			})
			return
		}
		etype := fmt.Sprint(values[i][4])
		eventkey := fmt.Sprint(values[i][5])
		limit := fmt.Sprint(values[i][7])
		model := fmt.Sprint(values[i][8])
		notikey := fmt.Sprint(values[i][9])
		site := fmt.Sprint(values[i][10])

		event := EventData{time, actkey, actval, curval, etype, eventkey, limit, model, notikey, site}
		eventList = append(eventList, event)

	}
	response := ResponseEventData{id, eventList}

	c.JSON(http.StatusOK, response)
}

// Site ..
// @Description 해당사이트에 속한 장치들의 최근측정데이터
// @Router /site [get]
// @Tags equip
// @Accept  json
// @Produce  json
// @Param id query string true "사이트아이디"
// @Success 200 {object} SiteLastDataResponse
func Site(c *gin.Context) {
	// 사이트아이디 query string parameter bind
	id := c.Query("id")
	// where 조건 사이트아이디  측정키, 장치아이디로 그룹해 가장최근데이타 (장치,측정키)별 가장최근 데이터 조회
	res, err := util.QueryDB("select * from iotdata5 where s = '" + id + "' group by k,i order by time desc limit 1 ;")
	// queryDB err처리
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return
	}
	// 데이터리스트
	var list []EquipSeries

	/*
		                     message              Tags                     Column           Values [][]interface{}          Partial
			series[0] ex)  {iotdata5 map[i:zo7ic7jJZmaxGzBjNjASRk k:tm2] [time m s] [[2020-02-21T04:13:50.78Z 286 site004]] false}

	*/
	// 시리즈를 순회하며 응답할형태의 데이터로 조립
	series := res[0].Series
	log.Println(res[0].Series[0])
	for i := 0; i < len(series); i++ {
		devID := series[i].Tags["i"]
		key := series[i].Tags["k"]
		time := fmt.Sprint(series[i].Values[0][0])
		measure, err := series[i].Values[0][1].(json.Number).Int64()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed",
			})
		}
		equipseries := EquipSeries{devID, key, time, measure}
		list = append(list, equipseries)
	}
	response := SiteLastDataResponse{id, list}

	c.JSON(http.StatusOK, response)
}

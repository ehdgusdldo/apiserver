package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ehdgusdldo/APIServer/util"
	"github.com/gin-gonic/gin"
)

// Influx value struct
type Influx struct {
	Time  string `json:"t"`
	Value string `json:"v"`
}

// Response response struct
type Response struct {
	I     string `json:"i"`
	K     string `json:"k"`
	Value []Influx
}

// Equip ..
// @Description influx 조회
// @Router /equip [get]
// @Tags equip
// @Accept  json
// @Produce  json
// @Param id query string true "장치아이디"
// @Param tag query string true "태그명"
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {object} Response
func Equip(c *gin.Context) {
	// query string parameter bind
	id := c.Query("id")
	tag := c.Query("tag")
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
	res, err := util.QueryDB("select * from iotdata5 where i = '" + id + "' and k = '" + tag + "' order by time desc limit " + limit + " offset " + offset + " ;")

	// queryDB err처리
	if err != nil {
		fmt.Println(err)
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

	var value []Influx
	// Values를 반복돌며 time value를 뽑아내  []value생성
	// values[i]의 구조예시 -> [[2020-02-10T07:27:09.875302157Z] [TD2BDhvjKEaZo37c4DLAq6] [lx2] [521] [site003]]
	for i := 0; i < len(values); i++ {
		// [i][0] 시간
		time := fmt.Sprint(values[i][0])
		// [i][3] 측정값
		v := fmt.Sprint(values[i][3])

		influx := Influx{time, v}

		value = append(value, influx)
	}
	// 장치id와 tag값을 0번째데이터에서 가져와 응답 k, i에 셋팅
	k := fmt.Sprint(values[0][2])
	i := fmt.Sprint(values[0][1])
	var response Response
	response = Response{i, k, value}

	c.JSON(http.StatusOK, response)

}

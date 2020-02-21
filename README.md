# ziumks API SERVER

## setting value (환경변수)

1. GoDotEnv(현재 채택)
<https://github.com/joho/godotenv>

2. env 주석으로사용
<https://github.com/caarlos0/env>


## go-gin   gin-gonic webframework
<https://github.com/gin-gonic/gin>

## MySQL Driver
<https://github.com/go-sql-driver/mysql>

## DBORM XORM
<https://github.com/go-xorm/xorm>
### xorm insert시 db에 디폴트값이있어도 not null 체크되있을경우 반드시 입력해야함


### gin으로 web서버만들어보기 블로그 

1.   <https://dejavuqa.tistory.com/330?category=320633> 

2.  <https://bourbonkk.tistory.com/60?category=814937>


## jwt middleware
<https://github.com/appleboy/gin-jwt>

## influxdb

1. git <https://github.com/influxdata/influxdb1-client>

2. 예제 <https://medium.com/spankie/getting-started-with-influxdb-with-golang-example-10990c5efee7>

influxDB result구조

```goalng
type Result struct {
	StatementId int `json:"statement_id"`
	Series      []models.Row
	Messages    []*Message
	Err         string `json:"error,omitempty"`
}
```
result.Series[] 의구조
```golang
type Row struct {
	Name    string            `json:"name,omitempty"`
	Tags    map[string]string `json:"tags,omitempty"`
	Columns []string          `json:"columns,omitempty"`
	Values  [][]interface{}   `json:"values,omitempty"`
	Partial bool              `json:"partial,omitempty"`
}
```
influxdb의 measurement 의 종류에따라 select했을때 result의 형태가 제각각이라 원하는대로 조합해서 사용하기위해서는 먼저 찍어봐서 확인해야함
Group by를 사용하면 Result의 []Series 로 나뉘고 일반 select 는 Series[0].Values 로 대부분의 데이터가들어감

Result.Series.Values 는 interface{} 형태이기때문에 원하는 타입으로 convert 하기위해 string 은 fmt.Sprint를 사용했고 int값이들어있을땐 Json.Number 이기에 .(jsonNumber).int64() 로사용했음 int같은경우 err처리도해야함
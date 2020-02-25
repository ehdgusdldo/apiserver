package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// Command ..
type Command struct {
	DevID     string `form:"dev-id" json:"dev-id" binding:"required" example:"Wxldi39DJecl2dUdlJWL34"`
	ActionKey string `form:"actionkey" json:"actionkey" binding:"required" example:"acky"`
	ActionVal int    `form:"actionval" json:"actionval" binding:"required" example:"200"`
}

// CommandResponse swagger 응답표시용 struct
type CommandResponse struct {
	Message        string `example:"success"`
	SubscribeCount int64  `example:"5"`
}

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis-addr"),
		Password: "", // no password set
		DB:       0,  // use default DB
		// DialTimeout:  10 * time.Second,
		// ReadTimeout:  30 * time.Second,
		// WriteTimeout: 30 * time.Second,
		// PoolSize:     10,
		// PoolTimeout:  30 * time.Second,
	})
}

// HelloRedis 레디스 연결테스트
func HelloRedis(c *gin.Context) {
	pong, err := rdb.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

	// 스트림테스트
	// values := map[string]interface{}{
	// 	"doc": "deux",
	// 	"hi":  "hello",
	// }
	// log.Println(values)
	// err = rdb.XAdd(&redis.XAddArgs{
	// 	Stream: "stream2",
	// 	ID:     "2-1",
	// 	Values: values,
	// }).Err()

	// if err != nil {
	// 	log.Println(err)
	// }
}

// PubCommand request로 들어온 장치아이디, actionkey, actionVal 를 redis에 publish
// @Description redis publish
// @Router /command [post]
// @Tags redis
// @Accept  json
// @Produce  json
// @Param command body Command true "command"
// @success 200 {object} CommandResponse
func PubCommand(c *gin.Context) {
	var command Command
	// request body 바인딩
	c.ShouldBindJSON(&command)

	// redis publish
	count, err := rdb.Publish("usercommand", command.DevID+","+command.ActionKey+","+strconv.Itoa(command.ActionVal)).Result()

	log.Println("subscribe count :" + strconv.FormatInt(count, 10))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
			"error":   err,
		})
	}

	// 성공메시지와 subscribe 된 count 응답
	c.JSON(http.StatusOK, gin.H{
		"message":        "success",
		"SubscribeCount": count,
	})

}

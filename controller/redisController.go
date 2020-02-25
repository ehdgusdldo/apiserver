package controller

import (
	"fmt"
	"log"
	"net/http"
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

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.0.102:16379",
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
	rdb.XAdd(&redis.XAddArgs{
		Stream: "stream1",
		ID:     "2-0",
		Values: map[string]interface{}{"doc": "deux"},
	}).Result()
}

// PubCommand command  ..
// @Description redis publish
// @Router /command [post]
// @Tags redis
// @Accept  json
// @Produce  json
// @Param command body Command true "command"
func PubCommand(c *gin.Context) {
	var command Command
	c.ShouldBindJSON(&command)

	log.Println(command)
	count, err := rdb.Publish("usercommand", command.DevID+","+command.ActionKey+","+strconv.Itoa(command.ActionVal)).Result()
	log.Println("subscribe count :" + strconv.FormatInt(count, 10))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
			"error":   err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "success",
		"SubscribeCount": count,
	})

}

package util

import (
	"log"
	"os"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

// 추후 dbanme_custID 처럼들어오면 QueryDB에서 파라미터를 하나더받는식으로 수정?
const (
	// MyDB specifies name of database
	MyDB = "mydb"
)

var influxurl string

func init() {
	influxurl = os.Getenv("influx-url")
}

// Insert saves points to database
func Insert(productMeasurement map[string]interface{}) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: influxurl,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	tags := map[string]string{"productView": productMeasurement["ProductName"].(string)}
	fields := productMeasurement

	pt, err := client.NewPoint("products", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}

	// Close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}

// QueryDB convenience function to query the database
// http방식으로 쿼리를 날린후 result반환
// cmd : query문
func QueryDB(cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:   cmd,
		Database:  MyDB,
		Precision: "ns",
	}
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: influxurl,
	})

	if response, err := c.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

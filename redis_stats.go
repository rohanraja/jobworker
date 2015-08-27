package jobworker

import (
	"encoding/json"

	"github.com/rohan1020/retry"
)

func UpdateRedisStats() {

	obj := GetInfoObj()

	js, _ := json.Marshal(obj)

	outStr := string(js)

	SetInfoHash(obj.Host, outStr)

}

func SetInfoHash(host, resultStr string) {

	key := "job:workers"

	err := retry.Do(func() (err error) {
		_, err = Redis.Client.HSet(key, host, resultStr).Result()
		return

	}, func() {
		Redis_dispatch.InitClient()
	})

	if err != nil {
		panic(err)
	}

	return

}

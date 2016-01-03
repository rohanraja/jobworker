package jobworker

import (
	"encoding/json"
	"fmt"

	"github.com/rohan1020/retry"
)

func UpdateRedisStats() {

	obj := GetInfoObj()

	js, _ := json.Marshal(obj)

	outStr := string(js)

	key := fmt.Sprintf("%s_%d", obj.Host, obj.Pid)

	SetInfoHash(key, outStr)

}

func SetInfoHash(host, resultStr string) {

	key := "job:workers"

	err := retry.Do(func() (err error) {
		_, err = Redis.Client.HSet(key, host, resultStr).Result()
		return

	}, func() {
		Redis.InitClient()
	})

	if err != nil {
		panic(err)
	}

	return

}

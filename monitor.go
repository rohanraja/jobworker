package jobworker

import (
	"encoding/json"
	"time"

	"github.com/rohan1020/retry"
	"gopkg.in/redis.v3"
)

func GetActiveJobs() (jinfos []Info) {

	strs := GetInfoStrings()

	for _, s := range strs {

		var jinfo Info
		_ = json.Unmarshal([]byte(s), &jinfo)
		if time.Since(jinfo.TimeUpdate).Seconds() < 10 {
			jinfos = append(jinfos, jinfo)
		}
	}

	return
}

func GetInfoStrings() (val []string) {

	key := "job:workers"

	err := retry.Do(func() (err error) {
		val, err = Redis_fetch.Client.HVals(key).Result()
		if err == redis.Nil {
			err = nil
		}
		return

	}, func() {
		Redis_fetch.InitClient()
	})

	if err != nil {
		panic(err)
	}
	return
}

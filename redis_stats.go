package jobworker

import "github.com/rohan1020/retry"

func SetInfoHash(key, host, resultStr string) {

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

package jobworker

import (
	"github.com/rohan1020/retry"
	"gopkg.in/redis.v3"
)

func FetchJob(binkey string) (out string, err error) {

	defer func() {
		if r := recover(); r != nil {
			err, _ = r.(error)
		}
	}()

	jid := GetPendingJids(binkey)
	if jid == "" {
		out = ""
		return
	}
	out = GetJobInfo(binkey, jid)
	MoveJidToProcessingSet(binkey, jid)

	return
}

func GetPendingJids(binkey string) (val string) {

	key := "job:" + binkey + ":" + "pending"

	err := retry.Do(func() (err error) {
		val, err = Redis_fetch.Client.SRandMember(key).Result()
		if err == redis.Nil {
			err = nil
			val = ""
			return
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
func MoveJidToProcessingSet(binkey, jid string) (val string) {

	key_pend := "job:" + binkey + ":" + "pending"
	key_process := "job:" + binkey + ":" + "processing"

	err := retry.Do(func() (err error) {
		_, err = Redis_fetch.Client.SMove(key_pend, key_process, jid).Result()
		return

	}, func() {
		Redis_fetch.InitClient()
	})

	if err != nil {
		panic(err)
	}

	return

}
func GetJobInfo(binkey, jid string) (val string) {

	key := "job:" + binkey + ":" + "args"

	err := retry.Do(func() (err error) {
		val, err = Redis_fetch.Client.HGet(key, jid).Result()
		return

	}, func() {
		Redis_fetch.InitClient()
	})

	if err != nil {
		panic(err)
	}

	return

}

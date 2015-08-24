package jobworker

import "crypto/md5"
import "encoding/json"
import "fmt"
import "github.com/rohan1020/retry"

func EnqueueJob(args ...string) {

	binNext := ""
	if len(args) > 2 {
		binNext = args[2]
	}
	jinfo, jid := GenerateJobInfoString(args[0], args[1], binNext)

	QueueJobInRedis(args[0], jid, jinfo)

}

func QueueJobInRedis(binkey, jid, jinfo string) (err error) {

	defer func() {
		if r := recover(); r != nil {
			err, _ = r.(error)
		}
	}()

	AddToPendingSet(binkey, jid)
	AddToArgHash(binkey, jid, jinfo)

	return nil
}

func AddToPendingSet(binkey, jid string) {

	key := "job:" + binkey + ":" + "pending"

	err := retry.Do(func() (err error) {
		_, err = Redis.Client.SAdd(key, jid).Result()
		return

	}, func() {
		Redis.InitClient()
	})

	if err != nil {
		panic(err)
	}

}
func AddToArgHash(binkey, jid, jinfo string) {

	key := "job:" + binkey + ":" + "args"

	err := retry.Do(func() (err error) {
		_, err = Redis.Client.HSet(key, jid, jinfo).Result()
		return

	}, func() {
		Redis.InitClient()
	})

	if err != nil {
		panic(err)
	}
}

func GenerateJobInfoString(binkey, args, binNext string) (string, string) {

	jid := generateJobId(args)
	jInfo := JobInfo{args, jid, binkey, binNext}

	js, err := json.Marshal(jInfo)

	if err != nil {
		panic(err)
	}

	return string(js), jid
}

func generateJobId(keyStr string) string {

	data := []byte(keyStr)
	return fmt.Sprintf("%x", md5.Sum(data))
}

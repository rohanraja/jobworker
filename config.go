package jobworker

import (
	"time"

	"github.com/rohan1020/redisutils"
)

type ConfigInfo struct {
	BinaryPath       string
	OS_Prefix        string
	NumWorkers       int
	FetchPollDelay   time.Duration
	REDIS_Fetch_HOST string

	Fetch_Binkey string
	NumFetches   int
}

var Redis_fetch *redisutils.RedisConn
var Config ConfigInfo

func init() {

	Config.BinaryPath = GetBinaryPath()
	Config.OS_Prefix = GetOSPrefix()
	Config.NumWorkers = 40
	Config.FetchPollDelay = 5 * time.Second
	Config.REDIS_Fetch_HOST = "localhost:6379"
	Config.Fetch_Binkey = "parsebin"
	Config.NumFetches = 20

	Redis_fetch = redisutils.New(Config.REDIS_Fetch_HOST)
}

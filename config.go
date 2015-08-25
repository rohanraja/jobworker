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

	Fetch_Binkey       string
	NumFetches         int
	DispatchBufferSize int
}

var Redis_fetch *redisutils.RedisConn
var Redis_dispatch *redisutils.RedisConn
var Redis *redisutils.RedisConn
var Config ConfigInfo

func init() {

	Config.BinaryPath = GetBinaryPath()
	Config.OS_Prefix = GetOSPrefix()
	Config.NumWorkers = 5
	Config.FetchPollDelay = 0 * time.Second
	Config.REDIS_Fetch_HOST = "localhost:6379"
	Config.Fetch_Binkey = "bookinfocrawl"
	Config.NumFetches = 100

	Config.DispatchBufferSize = 100

	Redis_fetch = redisutils.New(Config.REDIS_Fetch_HOST)
	Redis_dispatch = redisutils.New(Config.REDIS_Fetch_HOST)
	Redis = redisutils.New(Config.REDIS_Fetch_HOST)
}

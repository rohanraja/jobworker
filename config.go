package jobworker

import (
	"os"
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
	ListenPort         int
}

var Redis_fetch *redisutils.RedisConn
var Redis_dispatch *redisutils.RedisConn
var Redis *redisutils.RedisConn
var Config ConfigInfo

func ChangeRedisHost(ipaddr string) {

	Config.REDIS_Fetch_HOST = ipaddr + ":6379"
	Redis_dispatch = redisutils.New(Config.REDIS_Fetch_HOST)
	Redis_fetch = redisutils.New(Config.REDIS_Fetch_HOST)
	Redis = redisutils.New(Config.REDIS_Fetch_HOST)
}

func init() {

	Config.BinaryPath = GetBinaryPath()
	os.Setenv("http_proxy", "http://10.3.100.207:8080")
	Config.OS_Prefix = GetOSPrefix()
	Config.NumWorkers = 40
	Config.FetchPollDelay = 1 * time.Second
	Config.Fetch_Binkey = "bookcrawl"
	Config.NumFetches = 1000
	Config.ListenPort = 8081

	Config.DispatchBufferSize = 500
	ChangeRedisHost("127.0.0.1")
	// Todo: Save dipatch buffer to file on quit

}

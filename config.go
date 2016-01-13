package jobworker

import (
	"flag"
	"os"
	"time"

	"github.com/rohan1020/redisutils"
	"github.com/vharitonsky/iniflags"
)

type ConfigInfo struct {
	BinaryPath         string
	OS_Prefix          string
	NumWorkers         int
	FetchPollDelay     time.Duration
	REDIS_HOST         string
	SERVER_IP          string
	Fetch_Binkey       string
	NumFetches         int
	DispatchBufferSize int
	ListenPort         int
	RequestQueueSize   int
}

var Redis_fetch *redisutils.RedisConn
var Redis_dispatch *redisutils.RedisConn
var Redis *redisutils.RedisConn
var Config ConfigInfo

func ChangeRedisHost(ipaddr string) {

	Config.REDIS_HOST = ipaddr + ":6379"
	Redis_dispatch = redisutils.New(Config.REDIS_HOST)
	Redis_fetch = redisutils.New(Config.REDIS_HOST)
	Redis = redisutils.New(Config.REDIS_HOST)
}

var fetchpoll string
var proxyset bool

func flagsSetup() {

	flag.IntVar(&Config.ListenPort, "port", 3015, "http listen port for the web interface")
	flag.IntVar(&Config.ListenPort, "p", 3015, "http listen port for the web interface")
	flag.IntVar(&Config.NumWorkers, "num", 1, "number of workers")
	flag.IntVar(&Config.NumWorkers, "n", 1, "number of workers")
	flag.IntVar(&Config.NumFetches, "numfetches", 1, "number of redis fetches in one go")
	flag.StringVar(&fetchpoll, "time", "5s", "time interval in seconds of polling redis")
	flag.StringVar(&Config.SERVER_IP, "ip", "127.0.0.1", "ip address of the host")
	flag.StringVar(&Config.Fetch_Binkey, "bin", "a.out", "name of binary/queue to crawl")
	flag.IntVar(&Config.RequestQueueSize, "reqsize", 4, "buffer size of request channel")
	flag.IntVar(&Config.DispatchBufferSize, "dispatch", 4, "buffer size of dispatch array")
	flag.BoolVar(&proxyset, "proxy", false, "set proxy server for IIT Kharagpur network")

	iniflags.Parse() // Instead of flag.Parse()

}

func init() {

	flagsSetup()

	Config.BinaryPath = GetBinaryPath()
	Config.OS_Prefix = GetOSPrefix()

	Config.FetchPollDelay, _ = time.ParseDuration(fetchpoll) // * time.Second
	ChangeRedisHost(Config.SERVER_IP)
	if proxyset == true {
		os.Setenv("http_proxy", "http://10.3.100.207:8080")

	}

}

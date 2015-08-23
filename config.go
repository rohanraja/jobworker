package jobworker

import "time"

type ConfigInfo struct {
	BinaryPath     string
	OS_Prefix      string
	NumWorkers     int
	FetchPollDelay time.Duration
}

var Config ConfigInfo

func init() {

	Config.BinaryPath = GetBinaryPath()
	Config.OS_Prefix = GetOSPrefix()
	Config.NumWorkers = 40
	Config.FetchPollDelay = 10 * time.Second

}

package jobworker

type ConfigInfo struct {
	BinaryPath string
	OS_Prefix  string
	NumWorkers int
}

var Config ConfigInfo

func init() {

	Config.BinaryPath = GetBinaryPath()
	Config.OS_Prefix = GetOSPrefix()
	Config.NumWorkers = 5

}

package jobworker

type ConfigInfo struct {
	BinaryPath string
	OS_Prefix  string
}

var Config ConfigInfo

func init() {

	Config.BinaryPath = GetBinaryPath()
	Config.OS_Prefix = GetOSPrefix()

}

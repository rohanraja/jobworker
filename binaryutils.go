package jobworker

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteBinary(filepath, arguments string) (string, error) {

	cmd := exec.Command(filepath, arguments)
	outBytes, err := cmd.Output()

	return string(outBytes), err
}

func CheckIfFileExists(filepath string) bool {

	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}

func GetOSPrefix() string {

	return "osx"
}
func GetBinaryPath() string {

	return "./"
}

func GetBinaryFilePath(binKey string) (fpath string) {

	fpath = fmt.Sprintf("%s%s", GetBinaryPath(), binKey)
	return
}

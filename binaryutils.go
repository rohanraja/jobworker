package jobworker

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
)

func ExecuteBinary(filepath, arguments string) (outStr string, err error) {

	cmd := exec.Command(filepath, arguments)
	// color.Magenta("Running Command %s", filepath)

	stOut, _ := cmd.StdoutPipe()
	stErr, _ := cmd.StderrPipe()
	cmd.Start()
	time.Sleep(2 * time.Second)
	// color.Yellow("Sending SIGINT")

	cmd.Process.Signal(os.Interrupt)
	time.Sleep(1 * time.Second)
	cmd.Process.Signal(os.Interrupt)
	m := io.MultiReader(stOut, stErr)
	r := bufio.NewReader(m)
	for i := 0; ; i++ {
		line, err := r.ReadString(byte('\n'))
		color.Yellow("%s", string(line))
		if err != nil {
			break
		}
	}
	outBytes, _ := ioutil.ReadAll(stOut)
	cmd.Wait()
	outStr = string(outBytes)

	return
}

func ExecuteBinary_WAIT(filepath, arguments string) (string, error) {

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

	if Config.OS_Prefix == "win" {
		return ""
	}

	return "./"
}

func GetBinaryFilePath(binKey string) (fpath string) {

	fpath = fmt.Sprintf("%s%s", GetBinaryPath(), binKey)
	return
}

package jobworker

import "os/exec"

func ExecuteBinary(filepath, arguments string) (string, error) {

	cmd := exec.Command(filepath, arguments)
	outBytes, err := cmd.Output()

	return string(outBytes), err
}

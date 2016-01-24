package jobworker

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	"github.com/fatih/color"
)

type CgtJob struct {
	TrainingJobId string
	FolderPath    string
	Output        string
	Cmd           *exec.Cmd
	Stderr        io.ReadCloser
	Stdout        io.ReadCloser
}

func ProcessCGTJob(job *JobRequest) (res string, err error) {

	res = "CGT Job Processed"

	var cgtJob CgtJob
	cgtJob.TrainingJobId = job.Jobinfo.Args
	cgtPath, _ := filepath.Abs("./cgtjobs")

	LogMsg("Processing " + cgtJob.TrainingJobId)

	cgtJob.FolderPath = cgtPath + "/" + cgtJob.TrainingJobId
	// cgtJob.FolderPath = "/Users/rraja/cgtjobs/" + cgtJob.TrainingJobId
	// cgtJob.FolderPath = cgtJob.TrainingJobId
	job.cgtJob = &cgtJob

	cgtJob.DownloadJobFiles()
	cgtJob.Execute()

	res = cgtJob.Output
	return
}

func (c *CgtJob) DownloadJobFiles() {

	isExist, _ := FileExists(c.FolderPath)
	if isExist == true {
		return
	}

	isExist, _ = FileExists("./cgtjobs")
	if isExist == false {
		os.MkdirAll("./cgtjobs", 0755)
	}

	url := fmt.Sprintf("http://%s:8000/cgtjobs/%s.zip", Config.SERVER_IP, c.TrainingJobId)
	fpath := fmt.Sprintf("%s.zip", c.FolderPath)
	// color.Yellow("Downloading From %s", url)
	// color.Yellow("Saving to %s", fpath)

	err := downloadFile(fpath, url)
	if err != nil {
		color.Red("%v", err)
		panic(err)
	}

	err = Unzip(fpath, c.FolderPath)
	if err != nil {
		color.Red("%v", err)
		panic(err)
	}

}

func (c *CgtJob) Execute() {

	filepath := "./a.out"
	arguments := c.FolderPath
	cmd := exec.Command(filepath, arguments)

	c.Cmd = cmd
	// color.Magenta("Starting Execution of CGT Job #%s", c.TrainingJobId)

	// stOut, _ := cmd.StdoutPipe()
	stErr, _ := cmd.StderrPipe()
	outfile, _ := os.Create("./out.txt")
	cmd.Stdout = outfile
	cmd.Stderr = outfile
	defer outfile.Close()

	// writer := bufio.NewWriter(outfile)
	// defer writer.Flush()

	c.Stderr = stErr

	cmd.Start()

	// go io.Copy(writer, stOut)

	// outBytes, _ := ioutil.ReadAll(stOut)
	//
	// c.Output = string(outBytes)

	cmd.Wait()

	LogMsg("Finished " + c.TrainingJobId)
	// color.Red("Finished Execution of CGT Job #%s", c.TrainingJobId)

}

type CgtStatus struct {
	Accuracy string
	BatchNo  int
	EpochNo  int
}

func (c *CgtJob) GetStatus() (res string, err error) {

	path := c.FolderPath + "/log"
	lines, err := readLines(path)

	if err != nil {
		return
		res = fmt.Sprintf("%v", err)
	}

	res = fmt.Sprintf("%s - Acc: %s, Epc: %s", c.TrainingJobId, lines[len(lines)-1], lines[len(lines)-2])

	return
}

func (c *CgtJob) UpdateRedisStatus() (err error) {
	status, _ := c.GetStatus()
	params_path := c.FolderPath + "/params_out"
	content, _ := ioutil.ReadFile(params_path)
	SetInfoHash("job:cgt:status", c.TrainingJobId, status)
	SetInfoHash("job:cgt:params", c.TrainingJobId, string(content))
	// LogMsg(fmt.Sprintf("%s PUpdated", c.TrainingJobId))
	return
}

func (c *CgtJob) GracefullyExit() (err error) {

	color.Red("Sending signal Quitting")
	cmd := c.Cmd
	// cmd.Process.Signal(syscall.Signal(31))
	// time.Sleep(2 * time.Second)
	// cmd.Process.Signal(syscall.Signal(31))
	// cmd.Process.Signal(os.Interrupt)
	sig := 31
	for i := 0; i < 10; i++ {
		err = syscall.Kill(cmd.Process.Pid, syscall.Signal(sig))
		if err != nil {
			return
		}
		if i == 3 {
			sig = 9
		}
		color.Red("Pid: %d, Error: %v", cmd.Process.Pid, err)
		time.Sleep(2 * time.Second)
	}
	return
}

func (c *CgtJob) GetStatus_old() (res string, err error) {

	cmd := c.Cmd
	// cmd.Process.Signal(os.Interrupt)
	cmd.Process.Signal(syscall.Signal(30))
	stErr := c.Stderr
	// color.Yellow("%v\n%v", cmd.Process, stErr)
	r := bufio.NewReader(stErr)

	out, _, err := r.ReadLine()
	res = string(out)

	return
}

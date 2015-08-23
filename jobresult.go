package jobworker

import "fmt"

type JobResult struct {
	Jobid     string
	ResultStr string
	Status    int
	ErrorMsg  string
	Jobinfo   *JobInfo
}

func NewJobResult(jinfo *JobInfo, rstr string, err error) (jr JobResult) {

	errno := 0
	if err != nil {
		errno = 1
	}

	jout := JobResult{
		jinfo.Jobid,
		rstr,
		errno,
		fmt.Sprintf("%v", err),
		jinfo,
	}

	return jout
}

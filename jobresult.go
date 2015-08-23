package jobworker

type JobResult struct {
	Jobid     string
	ResultStr string
	Status    int
	ErrorMsg  string
	Jobinfo   *JobInfo
}

func NewJobResult(jinfo *JobInfo, rstr string, err error) (jr *JobResult) {

	return nil
}

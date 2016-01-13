package jobworker

type JobRequest struct {
	Jobinfo        JobInfo
	Result         JobResult
	ResultsChannel chan JobResult
	cgtJob         *CgtJob
}

func (r *JobRequest) ProcessRequest() {

	resultStr, err := r.Executor()

	jobresult := NewJobResult(&r.Jobinfo, resultStr, err)

	r.ResultsChannel <- jobresult

}

func (r *JobRequest) Executor() (res string, err error) {

	if r.Jobinfo.BinaryKey == "a.out" {
		res, err = ProcessCGTJob(r)
		return
	}

	binPath := GetBinaryFilePath(r.Jobinfo.BinaryKey)

	res, err = ExecuteBinary(binPath, r.Jobinfo.Args)

	return
}

func (r *JobRequest) GetStatus() (res string, err error) {

	// color.Red("%s", r.cgtJob.FolderPath)

	if r.Jobinfo.BinaryKey == "a.out" {
		res, err = r.cgtJob.GetStatus()
		return
	}
	return
}
func (r *JobRequest) ExitGracefully() (err error) {

	// color.Red("%s", r.cgtJob.FolderPath)

	if r.Jobinfo.BinaryKey == "a.out" {
		err = r.cgtJob.GracefullyExit()
		return
	}
	return
}

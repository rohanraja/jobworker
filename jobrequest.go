package jobworker

type JobRequest struct {
	Jobinfo        JobInfo
	Result         JobResult
	ResultsChannel chan JobResult
}

func (r *JobRequest) ProcessRequest() {

	resultStr, err := r.Executor()

	jobresult := NewJobResult(&r.Jobinfo, resultStr, err)

	r.ResultsChannel <- jobresult

}

func (r *JobRequest) Executor() (res string, err error) {

	binPath := GetBinaryFilePath(r.Jobinfo.BinaryKey)

	res, err = ExecuteBinary(binPath, r.Jobinfo.Args)

	return
}

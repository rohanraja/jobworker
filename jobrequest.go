package jobworker

type JobRequest struct {
	Jobinfo        JobInfo
	Result         JobResult
	ResultsChannel JobResultsQueue
}

func (r *JobRequest) ProcessRequest() {

	resultStr, err := r.Executor()

	jobresult := NewJobResult(&r.Jobinfo, resultStr, err)

	r.ResultsChannel <- jobresult

}

func (r *JobRequest) Executor() (string, error) {

	return "", nil
}

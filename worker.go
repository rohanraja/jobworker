package jobworker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

type WorkForce struct {
	ExitSignalChannel chan int
	NumWorkers        int
	JobRequestQueue   chan JobRequest
	JobResultQueue    chan JobResult
	ActiveJobs        []*JobRequest
	// ActiveJobs        map[*JobRequest]bool
	sync.RWMutex
}

var workForce *WorkForce

func NewWorkForce(jrq chan JobRequest, jresq chan JobResult) *WorkForce {

	reqSignal := make(chan int)

	var jobsmap []*JobRequest
	// var jobsmap map[*JobRequest]bool
	var sm sync.RWMutex
	wf := WorkForce{reqSignal, Config.NumWorkers, jrq, jresq, jobsmap, sm}

	return &wf

}

func (w *WorkForce) ChangeNumWorkers(num int) {

	if num > w.NumWorkers {
		w.NumWorkers = num - w.NumWorkers
		w.StartWorking()

	} else {

		w.NumWorkers = w.NumWorkers - num
		w.StopJobWorkers()
	}
	w.NumWorkers = num

}

func (w *WorkForce) StartWorking() {

	// w.ActiveJobs = make(map[*JobRequest]bool)

	for i := 0; i < w.NumWorkers; i++ {
		go w.ProcessQueue()
	}
}

func (w *WorkForce) StopJobWorkers() {

	for i := 0; i < w.NumWorkers; i++ {
		w.ExitSignalChannel <- 1
	}
}

func (w *WorkForce) ProcessQueue() {

	for {
		select {

		case job := <-w.JobRequestQueue:

			w.Lock()
			// LogMsg(fmt.Sprintf("Adding Job %p", &job))
			w.ActiveJobs = append(w.ActiveJobs, &job)
			// w.ActiveJobs[&job] = true
			w.SaveActiveJobs()
			w.Unlock()
			job.ProcessRequest()
			w.Lock()
			LogMsg(fmt.Sprintf("Deleting Job %v", &job))
			// delete(w.ActiveJobs, &job)
			var tmpJobs []*JobRequest
			for _, j := range w.ActiveJobs {
				if j != &job {
					tmpJobs = append(tmpJobs, j)
				}
			}
			w.ActiveJobs = tmpJobs
			w.Unlock()

			select {
			case <-w.ExitSignalChannel:
				return

			default:
			}

		case <-w.ExitSignalChannel:
			return

		}

	}

}
func (w *WorkForce) UpdateStatusAll() {

	w.RLock()
	// LogMsg(fmt.Sprintf("#%v", w.ActiveJobs))
	for _, job := range w.ActiveJobs {
		_ = job.cgtJob.UpdateRedisStatus()
	}
	w.RUnlock()
}

func (w *WorkForce) GetStatusAll() (res string) {

	res = ""

	w.RLock()
	for _, job := range w.ActiveJobs {
		outb, _ := job.GetStatus()
		LogMsg(outb)
		res = res + string(outb) + "\n"
	}
	w.RUnlock()
	return
}
func (w *WorkForce) ExitGracefullyAll() (res string) {

	res = ""

	for _, job := range w.ActiveJobs {
		go job.ExitGracefully()
	}
	return
}

func (w *WorkForce) SaveActiveJobs() (res string) {

	res = ""

	var toSave []JobRequest
	for _, job := range w.ActiveJobs {
		toSave = append(toSave, *job)
	}
	// color.Yellow("Saving Jobs \n%v", toSave)
	byts, _ := json.Marshal(toSave)
	// color.Red("Error: %v", err)
	// color.Yellow("Bytes for Jobs %d", len(toSave))
	ioutil.WriteFile("activejobs.json", byts, 0644)

	return
}

func (w *WorkForce) LoadActiveJobs() (res string) {

	res = ""

	var toSave []JobRequest
	content, _ := ioutil.ReadFile("activejobs.json")
	_ = json.Unmarshal(content, &toSave)

	// color.Yellow("Loaded Jobs \n %v", toSave)

	for _, job := range toSave {
		job.ResultsChannel = w.JobResultQueue
		w.JobRequestQueue <- job
	}

	return
}

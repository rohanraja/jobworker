package jobworker

import (
	"fmt"
	"time"

	ui "github.com/gizak/termui"
)

var Messages []string

func RunTerminalUI() {

	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	// var ps []float64
	//
	// jobrateStream := (func() []float64 {
	//
	// 	jInfo := GetInfoObj()
	// 	ps = append(ps, float64(jInfo.JobRate))
	// 	return ps
	//
	// })
	// ui.UseTheme("helloworld")
	// lc := ui.NewLineChart()
	// lc.Border.Label = "JobRate"
	// lc.Data = jobrateStream()
	// lc.Width = 100
	// lc.Height = 20
	// lc.X = 0
	// lc.Y = 0
	// lc.AxesColor = ui.ColorWhite
	// lc.LineColor = ui.ColorRed | ui.AttrBold
	// lc.Mode = "dot"
	BufferPercent := func() int {

		return int(100 * len(resultsToDispatch) / Config.DispatchBufferSize)

	}

	gspeed := ui.NewGauge()
	gspeed.Percent = 50
	gspeed.Width = 50
	gspeed.Height = 3
	gspeed.Y = 8
	gspeed.X = 0
	gspeed.Border.Label = "Job Rate"
	gspeed.BarColor = ui.ColorGreen
	gspeed.Border.FgColor = ui.ColorWhite
	gspeed.Border.LabelFgColor = ui.ColorYellow

	greq := ui.NewGauge()
	greq.Percent = 50
	greq.Width = 50
	greq.Height = 3
	greq.Y = 0
	greq.X = 0
	greq.Border.Label = "Requests Buffer"
	greq.BarColor = ui.ColorYellow
	greq.Border.FgColor = ui.ColorWhite
	greq.Border.LabelFgColor = ui.ColorCyan

	g := ui.NewGauge()
	g.Percent = 50
	g.Width = 50
	g.Height = 3
	g.Y = 4
	g.Border.Label = "Results Buffer"
	g.BarColor = ui.ColorRed
	g.Border.FgColor = ui.ColorWhite
	g.Border.LabelFgColor = ui.ColorCyan

	list_msg := ui.NewList()
	list_msg.ItemFgColor = ui.ColorYellow
	list_msg.Border.Label = "Log"
	list_msg.Height = 20
	list_msg.Y = 14
	list_msg.X = 30
	list_msg.Width = 35

	list := ui.NewList()
	list.ItemFgColor = ui.ColorYellow
	list.Border.Label = "Info"
	list.Height = 10
	list.Y = 14
	list.Width = 25

	listItems := func() (out []string) {
		inf := GetInfoObj()
		out = append(out, fmt.Sprintf(" %d Jobs/Second", inf.JobRate))
		out = append(out, fmt.Sprintf(" %d  Workers", workForce.NumWorkers))
		out = append(out, fmt.Sprintf(" %d  Jobs Processed", TotalDone))
		out = append(out, fmt.Sprintf(" %s  QueueBinary", Config.Fetch_Binkey))
		out = append(out, fmt.Sprintf(" %s", inf.Host))
		out = append(out, fmt.Sprintf(" %v", inf.IpAddresses))
		return
	}

	p := ui.NewPar("0 Jobs per second")
	p.Height = 3
	p.Width = 50
	p.TextFgColor = ui.ColorWhite
	p.Border.Label = "Text Box"
	p.Border.FgColor = ui.ColorCyan

	draw := func(t int) {
		// offset := int((len(ps) / (lc.Width / 2)) * (lc.Width / 2))
		// lc.Data = jobrateStream()[offset:]

		list.Items = listItems()
		list_msg.Items = Messages
		gspeed.Percent = GetInfoObj().JobRate
		g.Percent = BufferPercent()
		greq.Percent = int(100 * len(workForce.JobRequestQueue) / Config.RequestQueueSize)
		// p.Text = fmt.Sprintf("%d Jobs/Second | %d Workers | %d Jobs Done", GetInfoObj().JobRate, workForce.NumWorkers, TotalDone)
		ui.Render(list_msg, list, g, greq, gspeed)
	}

	evt := ui.EventCh()

	i := 0
	drawEnabled := true
	for {
		select {
		case e := <-evt:

			if e.Type == ui.EventKey {
				switch e.Ch {

				case 'd':
					drawEnabled = !drawEnabled

				}

			}

			if e.Type == ui.EventKey && e.Ch == 'J' {
				workForce.ChangeNumWorkers(workForce.NumWorkers - 10)
			}
			if e.Type == ui.EventKey && e.Ch == 'K' {
				workForce.ChangeNumWorkers(workForce.NumWorkers + 10)
			}
			if e.Type == ui.EventKey && e.Ch == 'j' {
				workForce.ChangeNumWorkers(workForce.NumWorkers - 1)
			}
			if e.Type == ui.EventKey && e.Ch == 'k' {
				workForce.ChangeNumWorkers(workForce.NumWorkers + 1)
			}
			if e.Type == ui.EventKey && e.Ch == 's' {
				funcc := func() {
					statusStr := workForce.GetStatusAll()
					Messages = append(Messages, statusStr)
					draw(i)
				}
				go funcc()
			}
			if e.Type == ui.EventKey && e.Ch == 'q' {
				return
			}
		default:
			if drawEnabled == true {

				draw(i)
			}
			i++
			time.Sleep(time.Second / 2)
		}
	}

}

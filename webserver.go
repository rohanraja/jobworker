package jobworker

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Info struct {
	JobRate   float64
	TotalDone int
}

func NumWorkersHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	num := r.PostFormValue("numworkers")
	binkey := r.PostFormValue("binkey")
	numInt, _ := strconv.Atoi(num)

	Config.NumFetches = numInt
	Config.Fetch_Binkey = binkey

	StopJobWorkers()

	fmt.Fprintf(w, "Num Changed to %d", numInt)

}
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/info.html")

	j := Info{Rate, TotalDone}

	t.Execute(w, &j)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func StartWebServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/info", InfoHandler)
	http.HandleFunc("/changenum", NumWorkersHandler)
	Log.Fatal(http.ListenAndServe(":8080", nil))
}

package jobworker

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"golang.org/x/net/websocket"
)

type Info struct {
	JobRate    int
	NumWorkers int
	TotalDone  int
	BinKey     string
}

func NumWorkersHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	num := r.PostFormValue("numworkers")
	binkey := r.PostFormValue("binkey")
	numInt, _ := strconv.Atoi(num)

	// Config.NumFetches = numInt
	Config.Fetch_Binkey = binkey

	workForce.ChangeNumWorkers(numInt)

	fmt.Fprintf(w, "<script>window.location = '/'</script>")

	// fmt.Fprintf(w, "Num Changed to %d", numInt)

}
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/info.html")

	j := Info{int(Rate), workForce.NumWorkers, TotalDone, Config.Fetch_Binkey}

	t.Execute(w, &j)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func WebsocketHandler(ws *websocket.Conn) {
	// io.Copy(ws, ws)
	ws.Write([]byte("Hey there.. Got some data"))
}
func StartWebServer() {
	http.HandleFunc("/", InfoHandler)
	http.HandleFunc("/changenum", NumWorkersHandler)
	// http.Handle("/websocket", websocket.Handler(WebsocketHandler))
	Log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Config.ListenPort), nil))
}

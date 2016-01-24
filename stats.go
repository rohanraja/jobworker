package jobworker

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

type Info struct {
	JobRate     int
	NumWorkers  int
	NumWorking  int
	TotalDone   int
	Progress    string
	BinKey      string
	Host        string
	IpAddresses []string
	Pid         int
	TimeUpdate  time.Time
}

func PeriodicInfoUpdater() {

	for {
		UpdateRedisStats()
		time.Sleep(5 * time.Second)
	}

}

func UpdateRedisStats() {

	obj := GetInfoObj()

	if obj.BinKey == "a.out" {
		workForce.UpdateStatusAll()
	}

	js, _ := json.Marshal(obj)

	outStr := string(js)

	key := fmt.Sprintf("%s_%d", obj.Host, obj.Pid)

	SetInfoHash("job:workers", key, outStr)

}
func GetInfoObj() Info {

	host, ips := GetNetworkStats()
	timenow := time.Now() //.MarshalJSON()
	j := Info{
		int(Rate),
		workForce.NumWorkers,
		len(workForce.ActiveJobs),
		TotalDone,
		workForce.GetStatusAll(),
		Config.Fetch_Binkey,
		host,
		ips,
		os.Getpid(),
		timenow}

	return j
}

func GetStats() {

}

func GetNetworkStats() (host string, ips []string) {

	host, _ = os.Hostname()
	// color.Green(host)
	addrs, _ := net.LookupIP(host)

	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			// fmt.Println("IPv4: ", ipv4)
			ips = append(ips, fmt.Sprintf("%v", ipv4))
		}
	}
	return

}

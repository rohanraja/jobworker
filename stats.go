package jobworker

import (
	"fmt"
	"net"
	"os"
	"time"
)

type Info struct {
	JobRate     int
	NumWorkers  int
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

func GetInfoObj() Info {

	host, ips := GetNetworkStats()
	timenow := time.Now() //.MarshalJSON()
	j := Info{
		int(Rate),
		workForce.NumWorkers,
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

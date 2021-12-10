package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"github.com/alexpetrean80/ipspy/lib"
)

func readIPFile(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var ips []string

	s := bufio.NewScanner(f)
	for s.Scan() {
		ips = append(ips, s.Text())
	}

	if err = s.Err(); err != nil {
		return nil, err
	}

	return ips, nil
}

func getIPStrings(fp string) (ipStrings []string, err error) {
	if fp != "" {
		return readIPFile(fp)
	}

	if len(flag.Args()) == 0 {
		return nil, fmt.Errorf("No IP address provided. Exiting...")
	}

	return flag.Args(), nil
}

func main() {
	var ips []net.IP
	var ipStrings []string

	ipFile := flag.String("f", "", "specify ip addresses from a file")
	allFlag := flag.Bool("a", false, "retrieve all data about the IP.")
	geoFlag := flag.Bool("geo", false, "retrieve geolocation data about the IP.")
	dnsFlag := flag.Bool("dns", false, "retrieve the reverse dns of the IP. empty if not found.")
	mobileFlag := flag.Bool("mobile", false, "check for cellular connection.")
	proxyFlag := flag.Bool("proxy", false, "check for proxy, VPN, or Tor exit address.")
	hostingFlag := flag.Bool("hosted", false, "check for hosting, colocated or data center")
	initIPFlag := flag.Bool("ip", false, "return the initial IP(useful for scans of multiple IPs).")
	ispFlag := flag.Bool("isp", false, "return the ISP of the IP.")
	flag.Parse()

	fields := lib.Fields{
		All:       *allFlag,
		Geo:       *geoFlag,
		DNS:       *dnsFlag,
		Mobile:    *mobileFlag,
		InitialIP: *initIPFlag,
		ISP:       *ispFlag,
		Proxy:     *proxyFlag,
		Hosting:   *hostingFlag,
	}

	ipStrings, err := getIPStrings(*ipFile)

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, ipStr := range ipStrings {
		ip, err := lib.ParseIP(ipStr)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		ips = append(ips, ip)
	}

	resCh := make(chan bytes.Buffer, 10)
	errCh := make(chan error, 10)
	quitCh := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		lib.Lookup(ips, fields, resCh, errCh, quitCh)
	}()

	for i := 0; i < len(ips); i++ {
		select {
		case res := <-resCh:
			fmt.Println(res.String())
		case err := <-errCh:
			fmt.Println(err.Error())
		}
	}

	wg.Wait()

}

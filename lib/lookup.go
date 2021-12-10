package lib

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"sync"
)

func Lookup(ips []net.IP, fields Fields, resCh chan bytes.Buffer, errCh chan error, quitCh chan int) {
	q := NewQuery(fields)
	wg := sync.WaitGroup{}
	ipLen := len(ips)

	wg.Add(ipLen)

	for _, ip := range ips {
		req := q.Create(ip)
		go request(req, resCh, errCh, &wg)
	}

	wg.Wait()
	close(resCh)
	close(errCh)
}

func request(req string, ch chan bytes.Buffer, errCh chan error, wg *sync.WaitGroup) {
	defer wg.Done()

	reader, err := makeReq(req)

	if err != nil {
		errCh <- err
		return
	}

	var buff bytes.Buffer
	if _, err := buff.ReadFrom(reader); err != nil {
		errCh <- err
		return
	}
	ch <- buff
}

func makeReq(q string) (io.Reader, error) {
	res, err := http.Get(q)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

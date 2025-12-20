package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[90m"
	bold   = "\033[1m"
)

const (
	address   = "127.0.0.1:8080"
	nThreads  = 50
	nRequests = 10000
)

var auth_cmd = []byte("PASS abc\n")
var cmd = []byte("SET key value\n")
var cmd2 = []byte("GET key\n")

func writeAndReadResponse(conn net.Conn, cmd []byte, reader *bufio.Reader) {
	_, err := conn.Write(cmd)
	if err != nil {
		panic(err)
	}
	_, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
}

func worker(wg *sync.WaitGroup, ops *int64, mu *sync.Mutex) {
	defer wg.Done()

	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	writeAndReadResponse(conn, auth_cmd, reader)

	for i := 0; i < nRequests; i++ {
		writeAndReadResponse(conn, cmd, reader)
		writeAndReadResponse(conn, cmd2, reader)
		mu.Lock()
		*ops++
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var ops int64

	start := time.Now()

	wg.Add(nThreads)
	for i := 0; i < nThreads; i++ {
		go worker(&wg, &ops, &mu)
	}
	wg.Wait()

	elapsed := time.Since(start)

	fmt.Println(bold + green + "✨ Benchmark:" + reset)
	fmt.Println(gray + "────────────────────────────" + reset)
	fmt.Println(cyan+"📊 Total operations:"+reset, yellow+fmt.Sprint(ops)+reset)
	fmt.Println(blue+"⏱️  Elapsed time:"+reset, elapsed)
	fmt.Printf(purple+"⚡ Ops/sec:"+reset+" "+bold+"%.0f"+reset+"\n", float64(ops)/elapsed.Seconds())
}

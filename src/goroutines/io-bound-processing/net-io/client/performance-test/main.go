package main

import (
	"fmt"
	"os"
	"runtime"
	"text/tabwriter"
	"time"

	"github.com/EnricoPicci/go-class.git/src/goroutines/io-bound-processing/net-io/client/send"
)

type test struct {
	requests      int           // number of requests to send to the server
	concurrent    int           // number of concurrent goroutines
	maxprocs      int           // if maxprocs is not specified, the max number of cores is used
	executionTime time.Duration // execution time of the test
}

func main() {
	tests := []test{
		// {requests: 100, concurrent: 1, maxprocs: 1},
		// {requests: 100, concurrent: 1, maxprocs: 4},
		// {requests: 100, concurrent: 1},
		{requests: 100, concurrent: 10, maxprocs: 1},
		{requests: 100, concurrent: 10, maxprocs: 4},
		{requests: 100, concurrent: 10},
		{requests: 100, concurrent: 100, maxprocs: 1},
		{requests: 100, concurrent: 100, maxprocs: 4},
		{requests: 100, concurrent: 100},
		{requests: 1000, concurrent: 100, maxprocs: 1},
		{requests: 1000, concurrent: 100},
		// {requests: 10000, concurrent: 1000, maxprocs: 1},
		// {requests: 10000, concurrent: 1000},
	}

	fmt.Println("*******************************************************************************************************************************")
	fmt.Println("****************************************         RESULTS      *****************************************************************")
	fmt.Println("*******************************************************************************************************************************")
	fmt.Println("")

	// w := new(tabwriter.Writer)
	// w.Init(os.Stdout, 0, 24, 0, '\t', 0)

	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.AlignRight)

	fmt.Fprintln(w, "Requests\tConcurrency\tMaxProcs\tExecutionTime\t")

	for i := range tests {
		t := tests[i]
		if t.maxprocs == 0 {
			t.maxprocs = runtime.NumCPU()
		}
		runtime.GOMAXPROCS(t.maxprocs)

		start := time.Now()
		send.CallServer(t.requests, t.concurrent)
		t.executionTime = time.Since(start)
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t\n", t.requests, t.concurrent, t.maxprocs, t.executionTime)

		fmt.Printf("Send %v requests to the server with %v concurrent goroutines and %v active cores - Execution time: %v \n",
			t.requests, t.concurrent, t.maxprocs, t.executionTime)
	}

	fmt.Print("\n\n")
	fmt.Println("*******************************************************************************************************************************")
	fmt.Println("****************************************         RESULTS      *****************************************************************")
	fmt.Println("*******************************************************************************************************************************")
	fmt.Println("")
	w.Flush()
}

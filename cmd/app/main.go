package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shv-ng/clip"
)

func main() {
	serverPtr := flag.Bool("server", false, "start the server")
	benchmarkPtr := flag.Bool("benchmark", false, "run benchmark on a file")
	filePtr := flag.String("file", "", "path to file contain url, one url in one line")

	flag.Parse()

	if *serverPtr {
		// TODO: make a http server for our shortner
		fmt.Println("server: server isn't setup yet")
		return
	}

	if *benchmarkPtr {
		if *filePtr == "" {
			fmt.Println("-file is required when use -benchmark")
			os.Exit(1)
		}
		err := clip.RunBenchmark(*filePtr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	flag.Usage()
}

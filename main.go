package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"
)

var usage = `usage: fileserver [-h] [--bind ADDRESS] [--directory DIRECTORY] [port]`

func main() {
	var (
		addr string
		path string
		help bool
	)

	flag.StringVar(&addr, "bind", "0.0.0.0", "server address")
	flag.StringVar(&path, "directory", "", "folder to serve")
	flag.BoolVarP(&help, "help", "h", false, "print help")
	flag.Parse()

	if help {
		fmt.Println("usage: fileserver [-h] [--bind ADDRESS] [--directory DIRECTORY] [port]")
		flag.PrintDefaults()
		os.Exit(0)
	}

	var (
		port = 8000
		err  error
	)
	switch len(flag.Args()) {
	case 0:
	case 1:
		port, err = strconv.Atoi(flag.Arg(0))
		if err != nil {
			fmt.Fprintln(os.Stderr, "positional argument #1 (port) must be a number")
			os.Exit(1)
		}
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	var message string
	if len(path) == 0 {
		path, err = os.Getwd()
		message = "get current directory failed: %v"
	} else {
		path, err = filepath.Abs(path)
		message = "path to absolute failed: %v"
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, message, err)
		os.Exit(1)
	}

	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir(path)))

	printAddr := addr
	if strings.HasPrefix(printAddr, "0.0.0.0") {
		printAddr = strings.Replace(printAddr, "0.0.0.0", "127.0.0.1", -1)
	}

	fmt.Printf("Serving HTTP on %s port %d -- http://%s:%d/ ...\n", addr, port, printAddr, port)

	httpAddr := fmt.Sprintf("%s:%d", addr, port)
	err = http.ListenAndServe(httpAddr, router)
	if err != nil {
		fmt.Fprintf(os.Stderr, "HTTP server failed: %v", err)
		os.Exit(1)
	}
}

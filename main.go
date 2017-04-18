package main

import (
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"time"
	"fmt"
	"path"
	"os/user"
	"runtime"
	"os"
)

func main() {
	for i := 1; i <= 5; i++ {
		go makeCalls(fmt.Sprintf("client%v", i))
	}

	time.Sleep(1000000000000)
}

func makeCalls(name string) {
	var url string
	if len(os.Args) > 1 {
		url = os.Args[1]
	} else {
		url = getDefaultSocketPath()
	}

	client := connect(url)
	defer client.Close()
	for {
		var result string
		err := client.Call(&result, "eth_blockNumber")
		if err != nil {
			log.Printf("%v failed %v", name, err)
		} else {
			log.Printf("%v answered %v", name, result)
		}
		time.Sleep(1000000000)
	}
}

func connect(url string) *rpc.Client {
	client, err := rpc.Dial(url)
	if err != nil {
		panic(err)
	}
	return client
}

func getDefaultSocketPath() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	var relativePath string
	switch runtime.GOOS {
	case "darwin":
		relativePath = "Library/Application Support/io.parity.ethereum/jsonrpc.ipc"
	default:
		relativePath = ".local/share/io.parity.ethereum/jsonrpc.ipc"
	}

	return path.Join(usr.HomeDir, relativePath)
}
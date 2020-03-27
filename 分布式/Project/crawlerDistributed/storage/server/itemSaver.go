package main

import (
	"../../../crawlerDistributed/rpcSupport"
	"../../../crawlerDistributed/storage"
	"flag"
	"fmt"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	err := rpcSupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		&storage.ItemSaverService{
			// item config
		})
	if err != nil {
		panic(err)
	}
}

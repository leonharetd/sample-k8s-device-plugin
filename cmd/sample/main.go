package main

import (
	"flag"
	"fmt"
	"log"
	"sample-k8s-device-plugin/pkg/sample"
)


var sugerNum *int

func init() {
    sugerNum = flag.Int("sugerNum", 2, "suger total num")
}

func main(){
    flag.Parse()
	stop := make(chan struct{})
	if err := sample.Register(); err != nil {
		log.Fatal(fmt.Errorf("failed to register with kubelet: %s", err))
	}
	go func() {
		if err := sample.WatchKubeletRestart(stop); err != nil {
			log.Fatal(fmt.Errorf("error watching kubelet restart: %s", err))
		}
	}()
	if err := sample.Serve(*sugerNum); err != nil {
		log.Fatal(fmt.Errorf("error running server: %s", err))
	}
	log.Printf("Suger Device plugin Start has %d suger", *sugerNum)
	stop <- struct{}{}
}
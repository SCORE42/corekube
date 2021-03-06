package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
	"travis/rax"
)

func getStackDetails() {

}

func deployStack() {
	readfile, _ := ioutil.ReadFile("../corekube-heat.yaml")
	template := string(readfile)
	templateRaw := fmt.Sprintf("template: %q\n", template)

	token := rax.IdentitySetup()
}

func waitForStackResult(heatTimeout int) []string {
	c1 := make(chan []string, 1)
	go func() {
		c1 <- []string{"1.2.3.4"}
	}()

	select {
	case res := <-c1:
		return res
	case <-time.After(time.Duration(heatTimeout) * time.Minute):
		log.Fatal("Timed out: Waiting for Heat Stack Result")
	}
	return nil
}

func testMinionsRegistered(machines []string, k8sTimeout int) {
	log.Printf("%s", machines)
}

func main() {
	//heatTimeout := 10 // minutes
	//k8sTimeout := 1   // minutes

	deployStack()
	//machines := waitForStackResult(heatTimeout)
	//testMinionsRegistered(machines, k8sTimeout)
}

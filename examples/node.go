package main

import (
  "time"
  "strconv"
  "fmt"
	rcfNode "rcf/rcfNode"
)

func main() {
  // creating node instance object which contains node struct in which all intern comm channels and topic/ action data maps are contained
  nodeInstance := rcfNode.Create(47)

  NserviceExeced := 0

  // initiating node by opening tcp server on node id
  // strarting action and topic handlers
  rcfNode.Init(nodeInstance)

  // adding action
  rcfNode.ActionCreate(nodeInstance, "testAction", func(params []byte, n rcfNode.Node){
    fmt.Println("---- ACTION TEST EXECUTED.")
    println(string(params))
  })


  // adding action
  rcfNode.ServiceCreate(nodeInstance, "testServiceDelay", func(params []byte, n rcfNode.Node) []byte {
    NserviceExeced += 1
    fmt.Println("---- Service delay TEST EXECUTED("+strconv.Itoa(NserviceExeced)+" times). Param: "+string(params))
    time.Sleep(1*time.Second)
    return params
  })

  // adding action
  rcfNode.ServiceCreate(nodeInstance, "testService", func(params []byte, n rcfNode.Node) []byte {
    NserviceExeced += 1
    fmt.Println("---- Service TEST EXECUTED TEST EXECUTED("+strconv.Itoa(NserviceExeced)+" times). Param: "+string(params))
    return params
  })

  // halting node so it doesn't quit
  rcfNode.NodeHalt()
}

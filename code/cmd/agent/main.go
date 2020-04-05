package main

import (
	"flag"
	"gitlab.bj.sensetime.com/zhangxiaohu/bootcamp/code/pkg/agent"
	"gitlab.bj.sensetime.com/zhangxiaohu/bootcamp/code/pkg/metrics"
	"strconv"
)
var (
	nameserverFile string
	informationMgr  * metrics.NodeInformationMgr
	// from command line
	serverAddress string
	serverPort int
)



func main() {
	flag.StringVar(&nameserverFile,"data-path","/tmp","set the nameserver file which operator will read /write")
	flag.StringVar(&serverAddress,"serverAddress","localhost","set the serverAddress")
	flag.IntVar(&serverPort,"serverPort",9000,"set the serverPort")
	flag.Parse()

	informationMgr = metrics.New(nameserverFile)

	if err := informationMgr.Init(); err != nil {
		metrics.Log.Errorf("can not init agent manager \n: %v", err)
	}



	err := agent.Run(informationMgr,serverAddress+":"+strconv.Itoa(serverPort))
	if err != nil {
		metrics.Log.Errorf("Run agent failed : %v",err)
	}

	// Change All fmt.Println to Log method
}

package main

import (
	"context"
	"flag"
	"gitlab.bj.sensetime.com/zhangxiaohu/bootcamp/code/pkg/metrics"
	"gitlab.bj.sensetime.com/zhangxiaohu/bootcamp/code/pkg/server"
	"sync"
)

/**
 *
 * @Author: zhangxiaohu
 * @File: main.go.go
 * @Version: 1.0.0
 * @Time: 2020/1/16
 */

var (
	dataPath string
	address string
	port int

	//informationMgr *metrics.NodeInformationMgr = make(map[string]*metrics.NodeInformationMgr)
	wg sync.WaitGroup
)

func main () {

	//  1. flags to parse command line
	flag.StringVar(&dataPath,"data-path","/tmp/server","set the nameserver file which operator will read /write")
	flag.StringVar(&address,"binding-address","localhost","set http server binding address")
	flag.IntVar(&port,"port",9000,"set http server port")
	flag.Parse()


	// 2. initalize server
	server := server.New(address,port,dataPath)

	if err := server.Init(); err != nil {
		metrics.Log.Error("can not init agent manager \n: %v", err)
	}


	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server.Run(ctx)

	<-ctx.Done()


}
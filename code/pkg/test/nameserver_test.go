package test

import (
	"fmt"
	"gitlab.bj.sensetime.com/zhangxiaohu/bootcamp/code/pkg/metrics"
	"testing"
)

/**
 *
 * @Author: zhangxiaohu
 * @File: nameserver_test.go
 * @Version: 1.0.0
 * @Time: 2020/1/15
 */
var(
	nameserverFile string
	mgr            *metrics.NodeInformationMgr
)
func TestNew(t *testing.T) {
	//flag.StringVar(&nameserverFile,"nameserver-file","/tmp/resolv.conf","set the nameserver file which operator will read /write")
	//flag.Parse()
	mgr = metrics.New(nameserverFile)
	mgr.Init()
	//mgr = Manager{path: nameserverFile}
}

func TestNameserverManager_List(t *testing.T) {
	TestNew(t)
	//fmt.Println(mgr)
	entry := mgr.Information.Nameserver.Entries
	fmt.Println(len(entry))
	for name, address := range entry {
		for _, value := range address.Address {
			println(name+" "+value)
		}
	}
}

func TestNameserverManager_Get(t *testing.T) {
	TestNew(t)
	address,_ := mgr.Information.Nameserver.Get("search")
	fmt.Println(address)
}

func TestNameserverManager_Set(t *testing.T) {
	TestNew(t)
	mgr.Information.Nameserver.AddEntries(nil)
	entry := mgr.Information.Nameserver.Entries
	fmt.Println(len(entry))
}

func TestNameserverManager_Delete(t *testing.T) {
	TestNew(t)
	mgr.Information.Nameserver.Delete("sensetime2")
	entry := mgr.Information.Nameserver.Entries
	fmt.Println(len(entry))
}

func TestNameserverManager_Update(t *testing.T) {
	TestNew(t)
	mgr.Information.Nameserver.UpdateEntries(nil)
	entry := mgr.Information.Nameserver.Entries
	fmt.Println(len(entry))
}


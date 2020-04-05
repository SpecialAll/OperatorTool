package metrics

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gitlab.bj.sensetime.com/zhangxiaohu/bootcamp/code/pkg/metrics/nameserver"
	"io"
	"os"
	"path/filepath"
	"strings"
)

/**
 *
 * @Author: zhangxiaohu
 * @File: metrics.go.go
 * @Version: 1.0.0
 * @Time: 2020/1/15
 */

const (
	nameserverPath = "resolv.conf"
	dnsPath = "hosts"
)

type NodeInfo struct {
	Nameserver *nameserver.Manager `json:"nameserver"`
}

func (info *NodeInfo)String()string{
	data,err := json.Marshal(info)
	if err != nil {
		return fmt.Sprintf("Marshal errors : %v", err)
	}
	return string(data)
}

type NodeInformationMgr struct{
	NameserverFilePath string
	Information *NodeInfo
}


func New(datapth string) *NodeInformationMgr{
	return & NodeInformationMgr{
		NameserverFilePath: datapth + "/" + nameserverPath,

		Information : new(NodeInfo),
	}
}

func (mgr * NodeInformationMgr) Init() error {
	return mgr.Load()
}


func ( mgr *NodeInformationMgr) Load() error {

	// open nameserver file
	nameserverEntries , err := mgr.loadNameserverFromFile(mgr.NameserverFilePath)
	if err != nil {
		return err
	}
	// load to mgr
	mgr.Information.Nameserver = nameserverEntries

	return  nil
}


func lineToEntry(line string )(string,*nameserver.NSEntry , bool){
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "#"){
		return "",nil,false
	}
	line = strings.TrimSpace(line)
	str := strings.Split(line, " ")
	entry := new(nameserver.NSEntry)
	if len(str) == 2 {
		entry.Address = append(entry.Address, str[1])
		return str[0],entry,true
	}

	return "",nil, false
}


func  ( mgr *NodeInformationMgr) loadNameserverFromFile(filePath string)( *nameserver.Manager,error ){
	//read all entries to mgr
	// O_RDWR
	f,err := os.OpenFile(filePath,os.O_RDONLY, 0666,)
	if err != nil{
		return  nil,err
	}
	defer f.Close()

	//entries := new(nameserver.Manager)
	entries := nameserver.New()
	buf := bufio.NewReader(f)

	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF  || err == nil {
			name,entry,ok := lineToEntry(line)
			if ok {
				if entries.Entries[name] != nil {
					for _, value := range entry.Address {
						entries.Entries[name].Address = append(entries.Entries[name].Address, value)
					}
				} else {
					entries.Entries[name] = entry
				}
			}
		}
		if err != nil{
			break
		}
	}
	return entries,nil
}

func (mgr *NodeInformationMgr) saveNameserverInfomation(f * os.File) error{
	entries := mgr.Information.Nameserver.Entries
	Log.Info(entries)
	writer := bufio.NewWriter(f)
	defer writer.Flush()
	for name,addressEntry := range entries {
		for _, address := range addressEntry.Address {
			line := fmt.Sprintf("%v %v\n",name,address)
			_,err := writer.WriteString(line)
			if err != nil {
				return err
			}
		}
	}
	return nil
}


func (mgr *NodeInformationMgr) Save() error{
	// O_RDWR
	// if the file is not exist, Openfile will create a file
	dir,_ := filepath.Split(mgr.NameserverFilePath)

	err := os.MkdirAll(dir,0777)
	if err != nil {
		return fmt.Errorf("Create node data directory failed : %v",err)
	}
	f,err := os.OpenFile(mgr.NameserverFilePath,os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil{
		return fmt.Errorf("openfile : %v",err)
		return  err
	}
	defer f.Close()

	mgr.saveNameserverInfomation(f)

	return nil
}


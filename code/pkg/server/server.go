package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gitlab.bj.sensetime.com/zhangxiaohu/bootcamp/code/pkg/metrics"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"sync"
)

/**
 *
 * @Author: zhangxiaohu
 * @File: server.go
 * @Version: 1.0.0
 * @Time: 2020/1/15
 */
var (
	lock            sync.Mutex
	informationsMgr = make(map[string]*metrics.NodeInformationMgr)
)

type Server struct {
	address string
	port int
	dataPath string
}

func New(address string, port int ,datePath string) *Server {
	return &Server{
		address: address,
		port:    port,
		dataPath: datePath,
	}
}

func (server * Server) Init() error {

	// 2. Load all Node information from datapath
	files,err := ioutil.ReadDir(server.dataPath)
	if err != nil{
		return err
	}
	// 2.1 check data directory, see some directory like node0 , node1
	for _,file := range files {
		if file.IsDir(){
			_ , fileName := filepath.Split(file.Name())
				informationsMgr[fileName] = metrics.New(server.dataPath+"/"+file.Name())
			err := informationsMgr[fileName].Load()
			if err != nil{
				return err
			}
			metrics.Log.Infof("Find node : %v : %v",fileName,informationsMgr[file.Name()].Information )
		}
	}
	return nil
}



func (server * Server) Run(ctx context.Context){
	go server.RunServer(ctx)
}

func (server * Server) RunServer(ctx context.Context){


	r := mux.NewRouter()
	r.HandleFunc("/nameserver", server.GetAllNodeNameserverInformation).Methods(http.MethodGet)
	r.HandleFunc("/nameserver", server.AddAllNameserverInformation).Methods(http.MethodPost)
	r.HandleFunc("/nameserver", server.UpdateAllNameserverInformation).Methods(http.MethodPut)
	r.HandleFunc("/nameserver", server.DeleteAllNameserverInformation).Methods(http.MethodDelete)
	r.HandleFunc("/nodes/{nodeName}/nameserver", server.GetNodeNameserverInformation).Methods(http.MethodGet)
	r.HandleFunc("/nodes/{nodeName}/nameserver", server.AddNodeNameserverInformation).Methods(http.MethodPost)
	r.HandleFunc("/nodes/{nodeName}/nameserver", server.UpdateNodeNameserverInformation).Methods(http.MethodPut)
	r.HandleFunc("/nodes/{nodeName}/nameserver", server.DeleteNodeNameserverInformation).Methods(http.MethodDelete)
	r.HandleFunc("/nodes/{nodeName}", server.RegisterAgent).Methods(http.MethodPost)

	http.ListenAndServe(server.address + ":" + strconv.Itoa(server.port),r)

}



func (server * Server)GetAllNodeNameserverInformation(w http.ResponseWriter, r *http.Request) {
	// 1. lock
	// 2. defer unlock
	// 3. send all information to agent

	lock.Lock()
	defer lock.Unlock()

	r.Close = true
	data,_ := json.Marshal(informationsMgr)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (server * Server)GetNodeNameserverInformation(w http.ResponseWriter, r *http.Request) {
	// 1. lock
	// 2. defer unlock
	// 3. send one node information to agent

	lock.Lock()
	defer lock.Unlock()

	vars := mux.Vars(r)
	data,_ := json.Marshal(informationsMgr[vars["nodeName"]])
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "node: %v\n", vars["node"]+string(data))
}



func (server *Server) AddAllNameserverInformation(writer http.ResponseWriter, request *http.Request) {
	// get add nameserver from request
	// lock
	// defer unlock
	// add nameserver to all nodes
	lock.Lock()
	defer lock.Unlock()

	content, err := ioutil.ReadAll(request.Body)

	if err != nil {
		metrics.Log.Error("Read request error : %v", err)
		writer.WriteHeader(402)
		writer.Write([]byte("Read Error "))
		return	}
	defer request.Body.Close()

	var tmp metrics.NodeInfo

	err = json.Unmarshal(content,&tmp)
	if err != nil {
		fmt.Printf("unmashal error : %v",err)
	}

	//update agent information
	for k,_ := range informationsMgr {
		informationsMgr[k].Information.Nameserver.AddEntries(tmp.Nameserver.Entries)
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("ok!\n"))


	for _,mgr := range informationsMgr {
		mgr.Save()
	}


}

func (server *Server) UpdateAllNameserverInformation(writer http.ResponseWriter, request *http.Request) {
	// get update nameserver from request
	// lock
	// defer unlock
	// update nameserver to all nodes
	lock.Lock()
	defer lock.Unlock()
	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		metrics.Log.Error("Read request error : %v", err)
		// follow api design return response
		writer.WriteHeader(402)
		writer.Write([]byte("Read Error "))
		return
	}
	defer request.Body.Close()

	var tmp metrics.NodeInfo

	err = json.Unmarshal(content,&tmp)
	if err != nil {
		metrics.Log.Error("Unmarshal request error : %v",err)
		// follow api design return response
		writer.WriteHeader(402)
		writer.Write([]byte("Unsupport message format "))
	}

	//update agent information
	for k,_ := range informationsMgr {
		informationsMgr[k].Information.Nameserver.UpdateEntries(tmp.Nameserver.Entries)
	}
	writer.WriteHeader(http.StatusOK)
	// follow api design return response
	writer.Write([]byte("ok!\n"))
}

func (server *Server) DeleteAllNameserverInformation(writer http.ResponseWriter, request *http.Request) {
	// get delete nameserver from request
	// lock
	// defer unlock
	// delete nameserver from all nodes
	lock.Lock()
	defer lock.Unlock()

	content, err := ioutil.ReadAll(request.Body)

	if err != nil {
		metrics.Log.Error("Read request error : %v", err)
		writer.WriteHeader(402)
		writer.Write([]byte("Read Error "))
		return
	}
	defer request.Body.Close()

	var tmp metrics.NodeInfo

	err = json.Unmarshal(content,&tmp)
	if err != nil {
		fmt.Printf("unmashal error : %v",err)
	}

	//delete agent nameserver information
	for k,_ := range informationsMgr {
		informationsMgr[k].Information.Nameserver.DeleteEntries(tmp.Nameserver.Entries)
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("ok!\n"))

}

func (server *Server) AddNodeNameserverInformation(writer http.ResponseWriter, request *http.Request) {
	// get add nameserver from request
	// lock
	// defer unlock
	// add nameserver to one node
	lock.Lock()
	defer lock.Unlock()

	content, err := ioutil.ReadAll(request.Body)

	if err != nil {
		metrics.Log.Error("Read request error : %v", err)
		writer.WriteHeader(402)
		writer.Write([]byte("Read Error "))
		return
	}
	defer request.Body.Close()

	var tmp metrics.NodeInfo

	vars := mux.Vars(request)
	nodeName := vars["nodeName"]
	err = json.Unmarshal(content,&tmp)
	if err != nil {
		fmt.Printf("unmashal error : %v",err)
	}

	//update agent information
	informationsMgr[nodeName].Information.Nameserver.AddEntries(tmp.Nameserver.Entries)
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("ok!\n"))

}

func (server *Server) UpdateNodeNameserverInformation(writer http.ResponseWriter, request *http.Request) {
	// get update nameserver from request
	// lock
	// defer unlock
	// update nameserver to one nodes
	lock.Lock()
	defer lock.Unlock()

	content, err := ioutil.ReadAll(request.Body)

	if err != nil {
		metrics.Log.Error("Read request error : %v", err)
		writer.WriteHeader(402)
		writer.Write([]byte("Read Error "))
		return
	}
	defer request.Body.Close()

	vars := mux.Vars(request)
	nodeName := vars["nodeName"]
	var tmp metrics.NodeInfo

	err = json.Unmarshal(content,&tmp)
	if err != nil {
		fmt.Printf("unmashal error : %v",err)
	}

	//update server information
	informationsMgr[nodeName].Information.Nameserver.UpdateEntries(tmp.Nameserver.Entries)

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("ok!\n"))

}

func (server *Server) DeleteNodeNameserverInformation(writer http.ResponseWriter, request *http.Request) {
	// get delete nameserver from request
	// lock
	// defer unlock
	// delete nameserver from one node
	lock.Lock()
	defer lock.Unlock()
	content, err := ioutil.ReadAll(request.Body)

	if err != nil {
		metrics.Log.Error("Read request error : %v", err)
		writer.WriteHeader(402)
		writer.Write([]byte("Read Error "))
		return
	}
	defer request.Body.Close()

	vars := mux.Vars(request)
	nodeName := vars["nodeName"]
	var tmp metrics.NodeInfo

	err = json.Unmarshal(content,&tmp)
	if err != nil {
		fmt.Printf("unmashal error : %v",err)
		// writer header unspported message format

	}

	//delete server information
	metrics.Log.Info(string(content))
	informationsMgr[nodeName].Information.Nameserver.DeleteEntries(tmp.Nameserver.Entries)


	writer.WriteHeader(http.StatusOK)
	// Define a status struct and a succsss const struct , return it
	writer.Write([]byte("state: success,message: success operation\n"))

}

func (server * Server) RegisterAgent(w http.ResponseWriter, r *http.Request){
	//1. get agent information form request
	//2. lock
	//3. defer unlock
	//4. update agent information
	lock.Lock()
	defer lock.Unlock()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		metrics.Log.Error("Read request error : %v", err)
		w.WriteHeader(402)
		w.Write([]byte("Read Error "))
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	nodeName := vars["nodeName"]
	var tmp *metrics.NodeInfo

	err = json.Unmarshal(content,&tmp)
	if err != nil {
		fmt.Printf("unmashal error : %v",err)
	}

	mgr,ok := informationsMgr[nodeName]
	if !ok{
		// first time register
		// change to real datapath
		mgr = metrics.New(server.dataPath+"/"+nodeName)
		mgr.Information = tmp
		informationsMgr[nodeName] = mgr
		metrics.Log.Infof("A new server : %v" , nodeName )
	}

	err = mgr.Save()
	if err !=  nil{
		metrics.Log.Errorf("Save config error : %v", err)
	}

	w.WriteHeader(http.StatusOK)
	data,_ := json.Marshal(informationsMgr[nodeName])
	w.Write(data)

}

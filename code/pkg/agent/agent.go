package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.bj.sensetime.com/zhangxiaohu/bootcamp/code/pkg/metrics"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)


// inject node_name to pod
//https://kubernetes.io/docs/tasks/inject-data-application/environment-variable-expose-pod-information/#use-pod-fields-as-values-for-environment-variables
const (
	nodeNameEnv = "NODE_NAME"
)

/**
 *
 * @Author: zhangxiaohu
 * @File: agent.go
 * @Version: 1.0.0
 * @Time: 2020/1/16
 */

func pingServer( mgr *metrics.NodeInformationMgr, serverURL string) error{

	data, err := json.Marshal(mgr.Information)

	if err != nil {
		return fmt.Errorf("Marshal information failed : %v",err)
	}

	req, err := http.NewRequest(http.MethodPost, serverURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	req.Close = true
	response, err := http.DefaultClient.Do(req)
	if err != nil{
		return  err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if body == nil  {
		return err
	}

	var tmp metrics.NodeInformationMgr

	err = json.Unmarshal(body,&tmp)
	if err != nil {
		return fmt.Errorf("Unmarshal failed : %v , %v ",err,string(body))
	}

	metrics.Log.Info(string(body))
	b := mgr.Information.Nameserver.UpdateAgentEntries(tmp.Information.Nameserver.Entries)

	if b {
		mgr.Save()
	}

	return nil

}

func Run(mgr *metrics.NodeInformationMgr, serverAddress string) error {

	nodeName := os.Getenv(nodeNameEnv)
	if nodeName == "" {
		return fmt.Errorf("Can not find node name from env : %v", nodeNameEnv )
	}

	for {
		//go connectWithServer(mgr, serverAddress)
		err := pingServer(mgr, "http://"+serverAddress+"/nodes/"+nodeName)
		if err != nil{
			metrics.Log.Errorf("Ping server error : %v\n",err)
		}
		time.Sleep(time.Second * 10)
	}

	return nil
}


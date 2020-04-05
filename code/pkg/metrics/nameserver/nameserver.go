package nameserver

import (
	"errors"
)

/**
 *
 * @Author: zhangxiaohu
 * @File: nameserver.go
 * @Version: 1.0.0
 * @Time: 2020/1/14
 */

// Manager mgr get/set/add/delete nameservers in node
type Manager struct {
	Entries map[string]*NSEntry `json:"entries"`
}

// nameserver entity
type NSEntry struct {
	Address []string `json:"address"`
}


func New() *Manager{
	return &Manager{
		Entries:make(map[string]*NSEntry),
	}
}



//get a nameserver from file
func (mgr *Manager) Get(key string)(*NSEntry ,error){
	ns := mgr.Entries

	for name, address := range ns {
		if key == name {
			return address, nil
		}
	}
	return nil, errors.New("the nameserver not exist")
}

// update one nameserver
// update all nameserver
func (mgr *Manager) UpdateEntries(entries map[string]*NSEntry) error {
	for name, entry := range entries {
		delete(mgr.Entries, name)
		mgr.Entries[name] = entry
	}
	return nil
}

// delete one nameserver
func (mgr *Manager) Delete(key string) error {
	delete(mgr.Entries,key)
	return nil
}

func (mgr *Manager) Contains(value string,name string) bool {
	for _,v := range mgr.Entries[name].Address {
		if value == v {
			return true
		}
	}
	return false
}

func (mgr * Manager) UpdateAgentEntries( entries map[string]*NSEntry) bool{

	b := false
	setmap := make(map[string] bool)
	//update
	for name, address := range entries {
		setmap[name] = true
		_, err := mgr.Get(name)
		if err == nil {
			//update
			delete(mgr.Entries, name)
			mgr.Entries[name] = address
			b = true
		} else {
			//add
			mgr.Entries[name] = address
			b = true
		}

		// the length is not equal
		if len(entries) != len(mgr.Entries) {
			for name, _ := range mgr.Entries {
				if !setmap[name] {
					mgr.Delete(name)
					b = true
				}
			}
		}
	}
	return b
}
func (mgr * Manager) AddEntries( entries map[string]*NSEntry) {
	//
	for name, entry := range entries {
		_,err := mgr.Get(name)
		if err == nil {
			for _, value := range entry.Address {
				mgr.Entries[name].Address = append(mgr.Entries[name].Address, value)
			}
		} else {
			mgr.Entries[name] = entry
		}
	}
}


func (mgr * Manager) DeleteEntries( entries map[string]*NSEntry) {
	for name,_ := range entries {
		_,err := mgr.Get(name)
		if err == nil {
			mgr.Delete(name)
		}
	}
}

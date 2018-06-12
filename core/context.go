package core

import (
	"encoding/json"
	"errors"
	"github.com/etherSync/common/config"
	"sync"
	"fmt"
)

var SystemContext *ProcessContext

type ProcessContext struct {
	sync.RWMutex
	EthHeight int64
}

func (this *ProcessContext) LoadContext() error {
	if SystemContext==nil{
		SystemContext=new(ProcessContext)
	}
	if str, err := DefaultStore.LoadContext(); err != nil {
		fmt.Printf("config.Parameters.SyncStart=%d",config.Parameters.SyncStart)
		SystemContext.EthHeight = config.Parameters.SyncStart
	} else {
		if err := json.Unmarshal(str, SystemContext); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (this *ProcessContext) SaveContext() error {
	str, err := json.Marshal(this)
	if err != nil {
		return err
	}
	if err := DefaultStore.SetContext(str); err != nil {
		return err
	}
	return nil
}

func (this *ProcessContext) SetEthHeight(height int64) error {
	this.Lock()
	defer this.Unlock()
	if height < this.EthHeight {
		return errors.New("[SetEthHeight] invalide height.")
	}
	this.EthHeight = height
	if err := this.SaveContext(); err != nil {
		return err
	}
	return nil
}

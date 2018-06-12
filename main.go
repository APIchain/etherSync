package main

import (
	"github.com/etherSync/common/config"
	"github.com/etherSync/common/log"
	"github.com/etherSync/core"
	"github.com/etherSync/sync"
	"github.com/etherSync/store/ChainStore"
	"os"
	"runtime"
	"time"
)

func main() {
	var err error
	runtime.GOMAXPROCS(8)
	log.Init(log.Path, log.Stdout)
	core.DefaultStore, err = ChainStore.NewLedgerStore()
	defer core.DefaultStore.Close()
	if err != nil {
		log.Error("open LedgerStore err:", err)
		os.Exit(1)
	}
	core.SystemContext.LoadContext()
	time.Sleep(2 * time.Second)

	go sync.SyncStart()
	go checklog()
	for {
		time.Sleep(5 * time.Second)
	}
}

func checklog() {
	for {
		time.Sleep(config.DEFAULTGENLOGTIME * time.Second)
		isNeedNewFile := log.CheckIfNeedNewFile()
		if isNeedNewFile == true {
			log.ClosePrintLog()
			log.Init(log.Path, os.Stdout)
		}
	}
}

package sync

import (
	"context"
	"fmt"
	"github.com/etherSync/core"
	"github.com/etherSync/metadata"
	"github.com/alanchchen/web3go/provider"
	"github.com/alanchchen/web3go/rpc"
	"github.com/alanchchen/web3go/web3"
	"github.com/etherSync/common/config"
	"github.com/etherSync/common/log"
	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"sync"
	"time"
)

var hostname = config.Parameters.SyncServer
var port = config.Parameters.SyncServerPort
var web3Api *web3.Web3
var blockTime int64
var blockPack int64 = 100000

var wgblock sync.WaitGroup
var rpcclient *ethrpc.Client
var syncBatchFlg = false

func SyncStart() {
	//first start
	log.Infof("Connect to %s:%s", hostname, port)
	RPCClient := provider.NewHTTPProvider(hostname+":"+port, rpc.GetDefaultMethod())
	web3Api = web3.NewWeb3(RPCClient)
	client, err := connectToRpc()
	if err != nil {
		panic(err.Error())
	}
	//syncBlock(client, 	4680308,	4680308)
	for true {
		onlineHeight, err := web3Api.Eth.BlockNumber()
		if err != nil {
			log.Error(err)
			time.Sleep(5 * time.Second)
			continue
		}
		localHeight := core.SystemContext.EthHeight
		for localHeight < onlineHeight.Int64() {
			localHeight++
			log.Tracef("[SyncStart] height is %d", localHeight)
			syncBlock(client, localHeight)
			core.SystemContext.EthHeight = core.SystemContext.EthHeight + 1
			core.SystemContext.SaveContext()
			log.Tracef("[SyncEnd  ] height is %d", localHeight)
		}
		time.Sleep(5 * time.Second)
	}
}

func syncBlock(client *ethclient.Client, start int64) error {
	log.Info("####################################################################################################")
	log.Infof("%s [syncBlock]Block=%d", time.Now().Format("2006-01-02 15:04:05"), start)
	if err := core.DefaultStore.NewBlock(); err != nil {
		log.Errorf("[syncBlock NewBlock] error=%s", err)
		time.Sleep(5 * time.Second)
		return err
	}
	syncEtherTransfer(client, start)
	syncTokenEvent(client, start)
	return nil
}

func logHandleTokenTransfer(hash, addr, from, to, value string, blockheight int64, removed bool) error {
	order := &metadata.TokenTransfer{
		Timestamp: blockTime,
		Height:    blockheight,
		TxHash:    hash,
		Token:     addr,
		From:      from,
		To:        to,
		Value:     value,
		Removed:   removed,
		//TransactionInfo: stats,
	}
	if err := core.SaveTokenTransferInfo(order); err != nil {
		log.Errorf("[SaveTokenTransferInfo] failed with hash=%s err=%s", hash, err)
		return err
	}
	return nil
}

func syncEtherTransfer(client *ethclient.Client, blockNum int64) error {
	log.Infof("%s [syncEtherTransfer] BlockNum=%d", time.Now().Format("2006-01-02 15:04:05"), blockNum)
	ctx := context.Background()
	block, err := getBlockInfo(ctx, blockNum)
	if err != nil {
		log.Errorf("[syncEtherTransfer] getBlockInfo failed with error=", err)
		return err
	}
	var count int64
	for _, v := range block.Transactions {
		realData := &metadata.TxTransfer{
			From:      v.From,
			To:        v.To,
			Value:     hexoToString(v.Value),
			TxHash:    v.Hash,
			Height:    hexoToInt(v.BlockNumber),
			Timestamp: hexoToInt(block.Timestamp),
		}
		blockTime = hexoToInt(block.Timestamp)
		if err := core.SaveEtherTransferInfo(realData); err != nil {
			log.Errorf("[syncEtherTransfer] SaveEtherTransferInfo failed with %s", err)
			return err
		}
		count++
	}
	log.Infof("%s [syncEtherTransfer] Block=%d (%d) transactions sync completed", time.Now().Format("2006-01-02 15:04:05"), blockNum, count)
	return nil
}

func syncTokenEvent(client *ethclient.Client, blockNum int64) error {
	log.Infof("%s [syncTokenEvent          ] Block=%d", time.Now().Format("2006-01-02 15:04:05"), blockNum)
	filter := ethereum.FilterQuery{}
	filter.Addresses = make([]common.Address, 0)
	filter.FromBlock = big.NewInt(blockNum)
	filter.ToBlock = big.NewInt(blockNum)
	hash := []common.Hash{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}
	filter.Topics = [][]common.Hash{hash}
	ctx := context.Background()
	logs, err := client.FilterLogs(ctx, filter)
	if err != nil {
		log.Errorf("[syncTokenEvent] FilterLogs failed with err=%s", err)
		return err
	}
	log.Infof("%s [syncTokenEvent          ] (%d) internal transaction found.", time.Now().Format("2006-01-02 15:04:05"), len(logs))
	for _, v := range logs {
		//str, _ := prettyjson.Marshal(v)
		//log.Debug(string(str))
		var from, to, value string
		if len(v.Topics) < 3 {
			if len(v.Topics) == 1 && len(v.Data) == 32*3 {
				detail := fmt.Sprintf("%x", v.Data)
				from, err = marchalAddr(detail[:64])
				if err != nil {
					log.Errorf("[marchalAddr] from failed with hash=%s,err=%s", v.TxHash.String(), err)
					continue
				}
				//log.Tracef("from =%s,to=%s,value=%d",from,to,value)
				to, err = marchalAddr(detail[64:128])
				if err != nil {
					log.Errorf("[marchalAddr] from failed with hash=%s,err=%s", v.TxHash.String(), err)
					continue
				}
				//log.Tracef("from =%s,to=%s,value=%s",from,to,value)
				value, err = marchalInt(detail[128:])
				if err != nil {
					log.Errorf("[marchalInt2] value failed with hash=%s,err=%s", v.TxHash.String(), err)
					continue
				}
				//log.Tracef("from =%s,to=%s,value=%d",from,to,value)

			} else {
				log.Errorf("error with hash=%s topic=%d,data=%d", v.TxHash.String(), len(v.Topics), len(v.Data))
				continue
			}
		} else {
			from, err = marchalAddr(v.Topics[1].String()[2:])
			if err != nil {
				log.Errorf("[marchalAddr] from failed with hash=%s,err=%s", v.TxHash.String(), err)
				continue
			}
			to, err = marchalAddr(v.Topics[2].String()[2:])
			if err != nil {
				log.Errorf("[marchalAddr] from failed with hash=%s,err=%s", v.TxHash.String(), err)
				continue
			}
			if len(v.Topics) < 4 {
				value, err = marchalInt(fmt.Sprintf("%x", v.Data))
				if err != nil {
					log.Errorf("[marchalInt2] value failed with hash=%s,err=%s", v.TxHash.String(), err)
					continue
				}
				//log.Infof("hash=%s,from=%s,to=%s,value=%s\n", v.TxHash.String(), from, to, value)
			} else {
				value, err = marchalInt(v.Topics[3].String()[2:])
				if err != nil {
					log.Errorf("[marchalInt3] value failed with hash=%s,err=%s", v.TxHash.String(), err)
					continue
				}
				log.Infof("hash=%s,from=%s,to=%s,value=%s", v.TxHash.String(), from, to, value)
			}
		}
		err = logHandleTokenTransfer(v.TxHash.String(), v.Address.String(), from, to, value, int64(v.BlockNumber), v.Removed)
		if err != nil {
			log.Errorf("[logHandleTokenTransfer] failed logHandleTokenTransfer with err=%s", err)
			return err
		}
	}
	log.Infof("%s [syncTokenEvent          ] Block=%d sync completed", time.Now().Format("2006-01-02 15:04:05"), blockNum)
	return nil
}

func connectToRpc() (*ethclient.Client, error) {
	client, err := ethrpc.Dial(fmt.Sprintf("http://%s:%s", hostname, port))
	if err != nil {
		return nil, err
	}
	rpcclient = client
	conn := ethclient.NewClient(client)
	log.Info("rpc connect completed.")
	return conn, nil
}
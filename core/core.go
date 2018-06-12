package core

import (
	"github.com/etherSync/common/log"
	"github.com/etherSync/metadata"
)

var DefaultStore IcoreStore

//token transfer
func SaveTokenTransferInfo(info *metadata.TokenTransfer) error {
	log.Debugf("[SaveTokenTransferInfo]transfer token saved =%v", info)
	//if err := DefaultStore.SaveTokenTransfer(info.Token, info.From, info.TxHash, info); err != nil {
	//	log.Error(err)
	//	return nil
	//}
	////
	//if info.From != info.To {
	//	if err := DefaultStore.SaveTokenTransfer(info.Token, info.To, info.TxHash, info); err != nil {
	//		log.Error(err)
	//		return nil
	//	}
	//}
	return nil
}

//eth transfer
func SaveEtherTransferInfo(info *metadata.TxTransfer) error {
	log.Debugf("[SaveEtherTransferInfo] save =%v", info)
	//if err := DefaultStore.SaveEtherTransfer(info.From, info.TxHash, info); err != nil {
	//	log.Error("[SaveEtherTransfer]failed.", err)
	//	return nil
	//}
	////
	//if info.From != info.To {
	//	if err := DefaultStore.SaveEtherTransfer(info.To, info.TxHash, info); err != nil {
	//		log.Error("[SaveEtherTransfer]failed.", err)
	//		return nil
	//	}
	//}
	//log.Debug("[SaveEtherTransferInfo] completed")
	return nil
	//}
}


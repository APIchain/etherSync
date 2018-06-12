package store

import "github.com/etherSync/metadata"

type IStore interface {
	SaveTokenTransfer(token, UserAddr, hash string, info *metadata.TokenTransfer) error

	//tx transfer
	SaveEtherTransfer(UserAddr, hash string, info *metadata.TxTransfer) error


	LoadContext() ([]byte, error)
	SetContext([]byte) error

	//close
	Close()

	NewBlock() error
	BlockCommit() error

}

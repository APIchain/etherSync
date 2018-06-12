package core

import "github.com/etherSync/metadata"

type IcoreStore interface {
	//token transfer
	//SaveTokenTransfer(token, UserAddr, hash string, data []byte) error
	//GetUserTokenTransferInfo(userAddr, token string) ([]*metadata.TokenTransfer, error)

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

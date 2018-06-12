package ChainStore

import (
	//"encoding/json"
	"github.com/etherSync/metadata"
	"github.com/etherSync/store"
	"github.com/etherSync/store/storage"
	. "github.com/etherSync/store/storage/LevelDBStore"
	"sync"
	//"math/big"
)

//**********************************************************************************************************************
// system function
//**********************************************************************************************************************
type ChainStore struct {
	st storage.IStore
	mu sync.RWMutex // guard the following var

}

func NewLedgerStore() (store.IStore, error) {
	// TODO: read config file decide which db to use.
	cs, err := NewChainStore("/opt/data/TokenServer")
	if err != nil {
		return nil, err
	}
	return cs, nil
}
func  (bd *ChainStore)SaveTokenTransfer(token, UserAddr, hash string, info *metadata.TokenTransfer) error{
	return nil
}

func  (bd *ChainStore)SaveEtherTransfer(UserAddr, hash string, info *metadata.TxTransfer) error{
	return nil
}

func NewChainStore(file string) (*ChainStore, error) {
	st, err := NewStore(file)
	if err != nil {
		return nil, err
	}
	chain := &ChainStore{
		st: st,
	}
	return chain, nil
}

func NewStore(file string) (storage.IStore, error) {
	ldbs, err := NewLevelDBStore(file)

	return ldbs, err
}

func (self *ChainStore) Close() {
	self.st.Close()
}
func (self *ChainStore) NewBlock() error{
	return self.st.NewBatch()
}
func (self *ChainStore) BlockCommit() error{
	return self.st.BatchCommit()
}

func (bd *ChainStore) LoadContext() ([]byte, error) {
	prefix := []byte{byte(metadata.DATA_CONTEXT)}
	data, err_get := bd.st.Get(append(prefix))
	if err_get != nil {
		return nil, err_get
	}
	return data, nil
}

func (bd *ChainStore) SetContext(data []byte) error {
	bd.mu.Lock()
	defer bd.mu.Unlock()
	prefix := []byte{byte(metadata.DATA_CONTEXT)}
	err := bd.st.Put(prefix, data)
	if err != nil {
		return err
	}
	return nil
}

func (bd *ChainStore) LoadKeyValue(setkey string) ([]byte, error) {
	//prefix := []byte{byte(metadata.DATA_TOKEN_LIST)}
	var key metadata.DBkey
	key = key.AppendPrefix(metadata.DATA_TOKEN_LIST).AppendString(setkey)
	data, err_get := bd.st.Get(append(key.GetData()))
	if err_get != nil {
		return nil, err_get
	}
	return data, nil
}

func (bd *ChainStore) SetKeyValue(setkey string, data []byte) error {
	bd.mu.Lock()
	defer bd.mu.Unlock()
	var key metadata.DBkey
	key = key.AppendPrefix(metadata.DATA_TOKEN_LIST).AppendString(setkey)
	//prefix := []byte{byte(metadata.DATA_TOKEN_LIST)}
	err := bd.st.Put(key.GetData(), data)
	if err != nil {
		return err
	}
	return nil
}


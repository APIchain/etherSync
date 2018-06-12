package metadata

import (
	"encoding/json"
)

type TokenTransfer struct {
	Timestamp       int64
	Height          int64
	TxHash          string
	Token           string
	From            string
	To              string
	Value           string
	Removed         bool
	//TransactionInfo *Receipt `json:"TransactionInfo" `
}

func (this *TokenTransfer) Marshal() ([]byte, error) {
	str, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return str, nil
}

func UnMarshalTokenTransfer(str []byte) (*TokenTransfer, error) {
	info := new(TokenTransfer)
	if err := json.Unmarshal(str, info); err != nil {
		return nil, err
	}
	return info, nil
}

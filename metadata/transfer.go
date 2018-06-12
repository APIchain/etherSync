package metadata

import (
	"encoding/json"
)

type TxTransfer struct {
	Timestamp       int64    `json:"Timestamp" `
	Height          int64    `json:"Height" `
	TxHash          string   `json:"TxHash" `
	From            string   `json:"From"     gencodec:"required"`
	To              string   `json:"To"       gencodec:"required"`
	Value           string   `json:"Value"    gencodec:"required"`
}

func (this *TxTransfer) Marshal() ([]byte, error) {
	str, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return str, nil
}

func UnMarshalTxTransfer(str []byte) (*TxTransfer, error) {
	info := new(TxTransfer)
	if err := json.Unmarshal(str, info); err != nil {
		return nil, err
	}
	return info, nil
}

package metadata

import (
	"strconv"
	"strings"
)

type DBkey struct {
	Key []byte
}

func (key DBkey) AppendString(str string) DBkey {
	var r []byte
	r = key.Key
	r = append(r, strings.ToUpper(str)...)
	key.Key = r
	return key
}

func (key DBkey) AppendInt64(num int64) DBkey {
	var r []byte
	r = key.Key
	r = append(r, strconv.FormatInt(num, 10)...)
	key.Key = r
	return key
}

func (key DBkey) AppendInt(num int) DBkey {
	var r []byte
	r = key.Key
	r = append(r, strconv.Itoa(num)...)
	key.Key = r
	return key
}

func (key DBkey) AppendPrefix(prefix DataEntryPrefix) DBkey {
	var r []byte
	r = key.Key
	r = []byte{byte(prefix)}
	key.Key = r
	return key
}

func (key DBkey) GetData() []byte {
	return key.Key
}

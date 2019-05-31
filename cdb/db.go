package cdb

import (
	"fmt"
	"github.com/lvchengchang/cbtree"
	"os"
	"sync"
)

const btreeDegrees = 64
const fileSuffix = ".dat"

type Cdb struct {
	mu       sync.Mutex
	keys     *cbtree.BTree
	file     *os.File
	closed   bool
	lastSeek int64
}

type Item struct {
	Key, Val string
}

func Open(dbName string) (*Cdb, error) {
	var err error
	cdb := &Cdb{}
	cdb.keys = cbtree.New(btreeDegrees, "keys")

	cdb.file, err = os.OpenFile(dbName+fileSuffix, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	// reload db file 实现重启然后读取硬盘数据到内存中
	cdb.reload()
	cdb.lastSeek, err = cdb.file.Seek(0, 2)
	go cdb.syncDisk()
	return cdb, nil
}

func (cdb *Cdb) Put(key string, val string) {
	fmt.Printf("key 值是%s, val值是%s\n", key, val)
	data := Item{key, val}
	cdb.keys.ReplaceOrInsert(&data)
}

// return cbtree.Item need Type Assertion
func (cdb *Cdb) Get(key string) *Item {
	data := cdb.keys.Get(&Item{Key: key})
	return data.(*Item)
}

func (cdb *Cdb) reload() {

}

func (cdb *Cdb) Close() {
	cdb.file.Close()
}

func (cdb *Cdb) syncDisk() {

}

func (i1 *Item) Less(item cbtree.Item, ctx interface{}) bool {
	i2 := item.(*Item)
	switch tag := ctx.(type) {
	case string:
		if tag == "vals" {
			if i1.Val < i2.Val {
				return true
			} else if i1.Val > i2.Val {
				return false
			}
			// Both vals are equal so we should fall though
			// and let the key comparison take over.
		}
	}
	return i1.Key < i2.Key
}

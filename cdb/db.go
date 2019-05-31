package cdb

import (
	"github.com/lvchengchang/cbtree"
	"os"
	"sync"
)

const btreeDegrees = 64
const fileSuffix = ".dat"

type Cdb struct {
	mu     sync.Mutex
	keys   *cbtree.BTree
	file   *os.File
	closed bool
}

type Item struct {
	Key, Val string
}

func Open(dbName string) (*Cdb, error) {
	var err error
	cdb := &Cdb{}
	cdb.keys = cbtree.New(btreeDegrees, nil)

	cdb.file, err = os.OpenFile(dbName+fileSuffix, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	// reload db file 实现重启然后读取硬盘数据到内存中
	cdb.load()
	go cdb.syncDisk()
	return cdb, nil
}

func (cdb *Cdb) Put(key string, val string) {
	data := Item{key, val}
	cdb.keys.ReplaceOrInsert(data)
}

func (cdb *Cdb) load() {

}

func (cdb *Cdb) Close() {
	cdb.file.Close()
}

func (cdb *Cdb) syncDisk() {

}

func (i1 Item) Less(item cbtree.Item, ctx interface{}) bool {
	return true
}

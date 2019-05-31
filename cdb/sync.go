package cdb

import (
	"bufio"
	"fmt"
	"github.com/lvchengchang/cbtree"
	"io"
	"log"
	"os"
	"time"
)

var (
	syncFail *os.File
	err      error
)

const TMPTREE = 64

// reload disk file
func (cdb *Cdb) reload() error {
	file, err := os.Open("cdb.dat")
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		fmt.Println(line)

		if err != nil || io.EOF == err {
			break
		}
	}

	return nil
}

func (cdb *Cdb) syncDisk() {
	t := time.NewTicker(time.Second)
	tmpTree := cbtree.New(TMPTREE, nil)

	// start ticker
	for range t.C {
		// range valid data in disk
		cdb.keys.Ascend(func(item cbtree.Item) bool {
			kvi := item.(*Item)
			tmpTree.ReplaceOrInsert(kvi)

			return true
		})

		// data write to new file
		syncFail, err = os.OpenFile("cdb.dat.bak", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Println(err)
			return
		}

		// writeNewFile
		var tmpNum int
		tmpTree.Ascend(func(i cbtree.Item) bool {
			kvi := i.(*Item)

			_, err = syncFail.WriteString("gocdb:k-" + kvi.Key + "-v" + kvi.Val)
			tmpNum++
			if tmpNum == 3 {
				syncFail.Write([]byte("\n"))
				tmpNum = 0
			}
			if err != nil {
				log.Println(err)
			}

			return true
		})

		// complete rename file
		err = os.Remove("cdb.dat")
		if err != nil {
			log.Println(err)
		}

		os.Rename("cdb.dat.bak", "cdb.dat")
		fmt.Println("complete one sync")
	}
}

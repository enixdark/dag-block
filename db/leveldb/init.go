package dag_leveldb

import (
	"os"
	"fmt"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type LeveldbDatabase struct {
	Session     *leveldb.DB
	Options *LevelOption
}

type LevelOption struct {
	Path   string
	Options *opt.Options
}

func (d *LeveldbDatabase) Connection() {

	database, err := leveldb.OpenFile(d.Options.Path, d.Options.Options)

	if err != nil {
		os.Exit(0)
	}

	d.Session = database
}

func (d *LeveldbDatabase) Close() {
	d.Session.Close()
}

func (d *LeveldbDatabase) Get(key string) ([]byte, error) {

	data, err := d.Session.Get([]byte(key), nil)
	if err != nil {
		fmt.Println("Can't get data")
		return nil, err
	} else {
		fmt.Println(data)
		return data, nil
	}
}

func (d *LeveldbDatabase) Put(key string, value string, unique bool) bool {
	key_store := key
	if unique == false {
		key_store += fmt.Sprintf(":%d", time.Now().Unix())
	}
	err := d.Session.Put([]byte(key_store), []byte(value), nil)

	if err != nil {
		fmt.Println("Can't not insert data")
		return false
	}

	return true
}

func (d *LeveldbDatabase) Seek(regex_key string) {

	iter := d.Session.NewIterator(util.BytesPrefix([]byte(regex_key)), nil)

	for iter.Next() {
		fmt.Println(string(iter.Key()))
	}

	iter.Release()
	err := iter.Error()

	if err != nil {
		fmt.Println(err)
	}

}

func (d *LeveldbDatabase) Traversal() {
	iter := d.Session.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		fmt.Println(string(iter.Key()) + "  " + string(iter.Value()))
	}

	iter.Release()
	err := iter.Error()

	if err != nil {
		fmt.Println(err)
	}
}

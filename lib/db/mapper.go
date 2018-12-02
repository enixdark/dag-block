package db

import (
	"github.com/enixdark/dag-block/lib/db/leveldb"
)

type DB struct {
	Db IDB
	Adapter string
}

type IDB interface {
	Connection()
	Close()
	//GetDatabase() *interface{}
	Get(key string) ([]byte, error)
	Put(key string, value string, unique bool) bool
	Seek(regex_key string)
	Traversal()
}

func NewDatabase(adapter string, opt interface{}) DB {
	database := DB{}
	if adapter == "leveldb" {
		options := opt.(dag_leveldb.LevelOption)
		database.Adapter = adapter
		database.Db = &dag_leveldb.LeveldbDatabase{
			Options: &options,
		}
	}

	return database
}







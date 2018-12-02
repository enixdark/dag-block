package tests

import (
	"testing"
	"os"
	"github.com/enixdark/dag-block/lib/db/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/enixdark/dag-block/lib"
	"github.com/enixdark/dag-block/lib/db"
	"github.com/enixdark/dag-block/lib/dag"
)

func TestLeveldbDatabase(t *testing.T) {

	dbPath := "/tmp/dag/db"
	os.RemoveAll(dbPath)

	option := dag_leveldb.LevelOption{
		Options: &opt.Options{
			Filter: filter.NewBloomFilter(10),
		},
		Path: dbPath,
	}

	db := db.NewDatabase("leveldb", option)

	dg := lib.DAG{
		Db: &db,
		Worker: make(chan string),
	}

	dg.Memory = dag.NewDAG()
	dg.Db.Db.Connection()

	defer dg.Db.Db.Close()

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		t.Errorf("The path %s does not exists", dbPath)
	}
}
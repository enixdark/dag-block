package main

import (

	"fmt"
	"os"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/enixdark/dag-block/lib"
	"github.com/enixdark/dag-block/lib/dag"
	"github.com/enixdark/dag-block/lib/db"
	"github.com/enixdark/dag-block/lib/db/leveldb"
)

func main() {

	//var err error
	os.RemoveAll("./data/db")

	option := dag_leveldb.LevelOption{
		Options: &opt.Options{
			Filter: filter.NewBloomFilter(10),
		},
		Path: "./data/db",
	}

	db := db.NewDatabase("leveldb", option)

	dg := lib.DAG{
		Db: &db,
		Worker: make(chan string),
	}

	dg.Memory = dag.NewDAG()
	dg.Db.Db.Connection()

	defer dg.Db.Db.Close()

	for i := 1; i < 100; i++ {
		vertex := dag.NewVertex(strconv.Itoa(i), nil)
		dg.Memory.AddVertex(vertex)
		dg.Insert(vertex)
	}

	fmt.Println(dg.List("4"))




}

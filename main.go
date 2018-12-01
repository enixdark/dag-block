package main

import (

	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/enixdark/dag-block/dag"
	"github.com/enixdark/dag-block/utils"
	"github.com/enixdark/dag-block/db"
	"github.com/enixdark/dag-block/db/leveldb"
)

type ID string

type DAG struct {
	db     *db.DB
	memory *dag.DAG
	worker chan string
}

//// Count number of ancestors/progeny (children of children) of any given vertex.
func (dg *DAG) Reach(id interface{}) int {
	return dg.ConditionalReach(id, true) + dg.ConditionalReach(id, false)
}
//
//// Count number of ancestors/progeny that have their "flag" set to true, or false for any given vertex.
func (dg *DAG) ConditionalReach(id interface{}, flagCondition bool) int {
	return dg.ConditionalList(id, flagCondition).Size()
}

//
//// List the ancestors/progeny with the requirements denoted in algorithms 1) and 2).
func (dg *DAG) List(id interface{}) *utils.OrderedSet {
	set := dg.ConditionalList(id, false)
	set_true := dg.ConditionalList(id, true)
	for _, e := range set_true.Values() {
		set.Add(e)
	}

	return set
}

func (dg *DAG) ConditionalList(id interface{}, flagCondition bool) *utils.OrderedSet {
	var set * utils.OrderedSet = utils.NewOrderedSet()
	vid, _ := dg.db.Db.Get("C:"+id.(string))

	if vid != nil {
		ids := strings.Split(string(vid), ",")

		if flagCondition {
			for _, id := range ids {
				vertex := dag.NewVertex(id, nil)
				set.Add(vertex)
			}
			return set
		} else {
			vertex, _ := dg.memory.GetVertex(id)
			for _, e := range vertex.Childrens().Values() {
				v := e.(*dag.Vertex)
				set.Add(v.ID)
			}

			not_intersection_size := set.StringNotContains(ids)
			return not_intersection_size
		}
		//f := flow.New()
		//sources.GenerateVertexSource(f,
		//	vertex.Childrens(), 1).Map(func(v dag.Vertex){
		//		dg.worker <- v.ID
		//}).Run()
	}

	return nil
}
//
//// Insert vertex to DAG and automatically construct necessary parent/children edges. Note that this DAG may be "incomplete" at any moment in time of its construction.
func (dg *DAG) Insert(vertex *dag.Vertex) {
	dg.memory.AddVertex(vertex)
	size := dg.memory.Order()
	// check if the elem vertex that added which's not a genesis vertex
	if size > 1 {
		// simulate pick a random
		parent_vertex := dg.memory.GetRamdomVertex()
		if parent_vertex != vertex {
			dg.memory.AddEdge(parent_vertex, vertex)
			dg.SyncDb(parent_vertex, vertex)
		}
	}
}

func (dg *DAG) SyncDb(parent *dag.Vertex, child *dag.Vertex) {

	parent_data, _ := dg.db.Db.Get("C:"+parent.ID)

	if parent_data != nil {
		dg.db.Db.Put("C:"+parent.ID, string(parent_data) + "," + child.ID, true)
	} else {
		dg.db.Db.Put("C:"+parent.ID, child.ID, true)
	}

	child_data, _ := dg.db.Db.Get("P:"+child.ID)

	if child_data != nil {
		dg.db.Db.Put("P:"+child.ID, string(child_data) + "," + parent.ID, true)
	} else {
		dg.db.Db.Put("P:"+child.ID, parent.ID, true)
	}
}

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

	dg := DAG{
		db: &db,
		worker: make(chan string),
	}

	dg.memory = dag.NewDAG()
	dg.db.Db.Connection()

	defer dg.db.Db.Close()

	for i := 1; i < 100; i++ {
		vertex := dag.NewVertex(strconv.Itoa(i), nil)
		dg.memory.AddVertex(vertex)
		dg.Insert(vertex)
	}

	fmt.Println(dg.List("4"))




}

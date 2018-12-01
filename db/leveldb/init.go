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

////// Count number of ancestors/progeny (children of children) of any given vertex.
//func (dg *DAG) Reach(id interface{}) int {
//	return dg.ConditionalReach(id, true) + dg.ConditionalReach(id, false)
//}
////
////// Count number of ancestors/progeny that have their "flag" set to true, or false for any given vertex.
//func (dg *DAG) ConditionalReach(id interface{}, flagCondition bool) int {
//	return dg.ConditionalList(id, flagCondition).Size()
//}
//
////
////// List the ancestors/progeny with the requirements denoted in algorithms 1) and 2).
//func (dg *DAG) List(id interface{}) *utils.OrderedSet {
//	set := dg.ConditionalList(id, false)
//	set_true := dg.ConditionalList(id, true)
//	for _, e := range set_true.Values() {
//		set.Add(e)
//	}
//
//	return set
//}
//
//func (dg *DAG) ConditionalList(id interface{}, flagCondition bool) *utils.OrderedSet {
//	var set * utils.OrderedSet = utils.NewOrderedSet()
//	vid, _ := dg.Get("C:"+id.(string))
//
//	if vid != nil {
//		ids := strings.Split(string(vid), ",")
//
//		if flagCondition {
//			for _, id := range ids {
//				vertex := dag.NewVertex(id, nil)
//				set.Add(vertex)
//			}
//			return set
//		} else {
//			vertex, _ := dg.memory.GetVertex(id)
//			for _, e := range vertex.Childrens().Values() {
//				v := e.(*dag.Vertex)
//				set.Add(v.ID)
//			}
//
//			not_intersection_size := set.StringNotContains(ids)
//			return not_intersection_size
//		}
//		//f := flow.New()
//		//sources.GenerateVertexSource(f,
//		//	vertex.Childrens(), 1).Map(func(v dag.Vertex){
//		//		dg.worker <- v.ID
//		//}).Run()
//	}
//
//	return nil
//}
////
////// Insert vertex to DAG and automatically construct necessary parent/children edges. Note that this DAG may be "incomplete" at any moment in time of its construction.
//func (dg *DAG) Insert(vertex *dag.Vertex) {
//	dg.memory.AddVertex(vertex)
//	size := dg.memory.Order()
//	// check if the elem vertex that added which's not a genesis vertex
//	if size > 1 {
//		// simulate pick a random
//		parent_vertex := dg.memory.GetRamdomVertex()
//		if parent_vertex != vertex {
//			dg.memory.AddEdge(parent_vertex, vertex)
//			dg.SyncDb(parent_vertex, vertex)
//		}
//	}
//}
//
//func (dg *DAG) SyncDb(parent *dag.Vertex, child *dag.Vertex) {
//
//	parent_data, _ := dg.Get("C:"+parent.ID)
//
//	if parent_data != nil {
//		dg.Put("C:"+parent.ID, string(parent_data) + "," + child.ID, true)
//	} else {
//		dg.Put("C:"+parent.ID, child.ID, true)
//	}
//
//	child_data, _ := dg.Get("P:"+child.ID)
//
//	if child_data != nil {
//		dg.Put("P:"+child.ID, string(child_data) + "," + parent.ID, true)
//	} else {
//		dg.Put("P:"+child.ID, parent.ID, true)
//	}
//}
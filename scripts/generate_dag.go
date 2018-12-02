package scripts

import (
	"math/rand"
	"os"
	"strconv"
	"encoding/json"
)

var (
	v           uint64 = 1000
	e           uint64 = 1250
	randomGraph [][]uint64
)


type key struct {
	child  uint64
	parent uint64
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}


func contains(arr []uint64, value uint64) bool {
	for _, a := range arr {
		if a == value {
			return true
		}
	}
	return false
}


func main() {

	edges := make(map[key]struct{})

	list_edges_with_parents := make(map[uint64][]uint64)

	var nodes []uint64

	for i := uint64(0); i < v; i++ {
		nodes = append(nodes, i)
		var nodeParents []uint64
		if i != 0 {
			parent := rand.Uint64() % i
			nodeParents = append(nodeParents, parent)
			e--

			edges[key{ child:i, parent: parent}] = struct{}{}
			if len(list_edges_with_parents[parent]) > 0 {
				if !contains(list_edges_with_parents[parent], i) {
					list_edges_with_parents[parent] = append(list_edges_with_parents[parent], i)
				}
			} else {
				list_edges_with_parents[parent] = []uint64{i}
			}
		}

		randomGraph = append(randomGraph, nodeParents)
	}

	for e > 0 {
		child := rand.Uint64()%(v-1) + 1
		parent := rand.Uint64() % child

		if _, exists := edges[key{parent, child}]; !exists {
			edges[key{ child:child, parent: parent}] = struct{}{}

			randomGraph[child] = append(randomGraph[child], parent)
			e--
		}
	}

	f, err := os.Create("./tmp/file.api")
	check(err)
	defer f.Close()


	sep := "\n"
	f.WriteString("child -> parent" + sep)
	for key, value := range list_edges_with_parents {

		v, _ := json.Marshal(value)

		if _, err = f.WriteString(strconv.FormatInt(int64(key), 10) + " -> " + string(v) + sep); err != nil {
			panic(err)
		}
	}

}
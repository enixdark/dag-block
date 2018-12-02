package sources

import (

	"github.com/chrislusf/glow/flow"
	"github.com/enixdark/dag-block/lib/utils"
	"github.com/enixdark/dag-block/lib/dag"
)

func GenerateVertexSource(f *flow.FlowContext, data *utils.OrderedSet, shard int) *flow.Dataset {
	var arr []dag.Vertex
	for _, e := range data.Values() {
		v, _ := e.(*dag.Vertex)
		arr = append(arr, *v)
	}

	return f.Slice(arr).Partition(shard)
}


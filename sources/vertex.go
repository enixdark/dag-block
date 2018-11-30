package sources

import (

	"github.com/chrislusf/glow/flow"
	//"fmt"
	//"github.com/chrislusf/gleam/pb"
	//"github.com/chrislusf/gleam/util"
	//"github.com/gocql/gocql"
	//"github.com/enixdark/dag-block/dag"
	"github.com/enixdark/dag-block/utils"
	//"reflect"
	"github.com/enixdark/dag-block/dag"
)



func GenerateVertexSource(f *flow.FlowContext, data *utils.OrderedSet, shard int) *flow.Dataset {
	var arr []dag.Vertex
	//orders := data.Values()
	for _, e := range data.Values() {
		v, _ := e.(*dag.Vertex)
		arr = append(arr, *v)
	}

	return f.Slice(arr).Partition(shard)
}

//func genShardInfos(f *flow.Flow, data *orderedset.OrderedSet, partitionIds []int32) *flow.Dataset {
//	return f.Source(s.prefix+".list", func(writer io.Writer, stats *pb.InstructionStat) error {
//
//		stats.InputCounter++
//
//		for _, pid := range partitionIds {
//			stats.OutputCounter++
//			util.NewRow(util.Now(), encodeShardInfo(&KafkaPartitionInfo{
//				Brokers:        s.Brokers,
//				Topic:          s.Topic,
//				Group:          s.Group,
//				TimeoutSeconds: s.TimeoutSeconds,
//				PartitionId:    pid,
//			})).WriteTo(writer)
//		}
//
//		return nil
//	})
//}
package doc

import (
	"fmt"

	"github.com/pinezapple/LibraryProject20201/portal/core"
	"github.com/pinezapple/LibraryProject20201/portal/microservice"
)

func getDocMangerServiceByString(str string) (ser *microservice.DocmanagerShardServiceWrapper, err error) {
	var (
		id = core.GetHash(str)
	)
	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		return nil, fmt.Errorf("nil shardService")
	}
	shardID := core.GetShardID(uint32(id))
	ser, ok := shardService[uint64(shardID)]
	if !ok {
		fmt.Println("nil shardID")
		return nil, fmt.Errorf("no shard id")
	}
	return ser, nil
}

func getDocMangerServiceByUint64(num uint64) (ser *microservice.DocmanagerShardServiceWrapper, err error) {
	shardService := microservice.GetDocmanagerShardServices()
	if shardService == nil {
		return nil, fmt.Errorf("nil shardService")
	}
	shardID := core.GetShardID(uint32(num))
	ser, ok := shardService[uint64(shardID)]
	if !ok {
		fmt.Println("nil shardID")
		return nil, fmt.Errorf("no shard id")
	}
	return ser, nil
}

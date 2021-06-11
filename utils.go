package gomysql

import (
	
	"hash/fnv"
	"strconv"

	"time"
)

func GenerateHash() uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(strconv.FormatInt(time.Now().UnixNano(), 10)))
	return hash.Sum32()
}

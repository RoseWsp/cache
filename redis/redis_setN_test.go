package redis

import (
	"strconv"
	"testing"
	"time"
)

// info memory => used_memory_human:60.00M
func TestSetCacheN(t *testing.T) {
	for i := 0; i < 1000*90; i++ {
		rdb.Set(ctx, strconv.Itoa(i), strconv.Itoa(i), 30*time.Second).Err()
	}
}

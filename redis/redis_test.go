package redis

/*
redis ，随着 value 的size 变大，其 set 和 get 的耗时都会增加。但这个耗时增长和 value size 增长是线性关系的。直到 value size 从5k 涨到1M,其操作耗时就是 数量级的差距 。
*/
import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

var tenByte = [10]byte{9: 1}
var twentyByte = [20]byte{19: 1}
var fiftyByte = [50]byte{49: 1}
var hundred = [100]byte{99: 1}
var twoHunderd = [200]byte{199: 1}

var oneK = [1024]byte{1023: 1}
var fiveK = [5 * 1024]byte{1023: 1}

var oneM = [1 * 1024 * 1024]byte{1023: 1}

var ctx = context.Background()
var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func BenchmarkSet10Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		setCache("10b")
	}
}

func BenchmarkSet20Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		setCache("20b")
	}
}

func BenchmarkSet50Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		setCache("50b")
	}
}

func BenchmarkSet100Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		setCache("100b")
	}
}

func BenchmarkSet200Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		setCache("200b")
	}
}

func BenchmarkSet1K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		setCache("1k")
	}
}

func BenchmarkSet5K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		setCache("5k")
	}
}

func BenchmarkSet1M(b *testing.B) {
	for i := 0; i < b.N; i++ {
		setCache("1m")
	}
}

func BenchmarkGet10Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getCache("10b")
	}
}

func BenchmarkGet20Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getCache("20b")
	}
}

func BenchmarkGet50Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getCache("50b")
	}
}

func BenchmarkGet100Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getCache("100b")
	}
}

func BenchmarkGet200Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getCache("200b")
	}
}

func BenchmarkGet1K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getCache("1k")
	}
}

func BenchmarkGet5K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getCache("5k")
	}
}

func BenchmarkGet1M(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getCache("1m")
	}
}

// 10 20 50 100 200 1k 5k
func setCache(cls string) {
	var err error

	switch cls {
	case "10b":
		err = rdb.Set(ctx, cls, (string)(tenByte[:]), 30*time.Second).Err()
	case "20b":
		err = rdb.Set(ctx, cls, (string)(twentyByte[:]), 30*time.Second).Err()
	case "50b":
		err = rdb.Set(ctx, cls, (string)(fiftyByte[:]), 30*time.Second).Err()
	case "100b":
		err = rdb.Set(ctx, cls, (string)(hundred[:]), 30*time.Second).Err()
	case "200b":
		err = rdb.Set(ctx, cls, (string)(twoHunderd[:]), 30*time.Second).Err()
	case "1k":
		err = rdb.Set(ctx, cls, (string)(oneK[:]), 30*time.Second).Err()
	case "5k":
		err = rdb.Set(ctx, cls, (string)(fiveK[:]), 30*time.Second).Err()
	case "1m":
		err = rdb.Set(ctx, cls, (string)(oneM[:]), 30*time.Second).Err()
	default:
		panic("unkonw cls") // 这里不应该抛panic ,这里是偷懒了 ，不想在 函数签名加 error
	}

	if err != nil {
		panic(err)
	}
}

// 10 20 50 100 200 1k 5k
func getCache(cls string) {
	_, err := rdb.Get(ctx, cls).Result()
	if err != nil {
		panic(err)
	}
}

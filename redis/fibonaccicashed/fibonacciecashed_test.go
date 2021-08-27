package fibonaccicashed

import (
	"github.com/Vladimir77715/fibonacciservice/fibonacci"
	_ "github.com/Vladimir77715/fibonacciservice/fibonacci"
	"github.com/Vladimir77715/fibonacciservice/redis/client"
	_ "github.com/Vladimir77715/fibonacciservice/redis/client"
	"strconv"
	"testing"
)

func BenchmarkWithCache(b *testing.B) {

	var startIndex = 1000
	var endIndex = 5000
	var size = endIndex - startIndex
	var shift = 0
	mockData, _ := fibonacci.FibonacciSlice(startIndex, endIndex, nil)
	var mockRedisCache = make(map[string]*client.RedisEntity, size)
	for i, j := 0, 0; i < size; i++ {
		mockRedisCache[strconv.Itoa(startIndex+shift+i)] = &client.RedisEntity{Key: strconv.Itoa(startIndex + shift + i), Value: strconv.FormatUint(mockData[i], 10)}
		j++
	}
	cw := CashedService{}

	b.StartTimer()
	cw.fibonacciSlice(startIndex, endIndex, size, false, mockRedisCache)
	b.StopTimer()
}

func BenchmarkWithoutCache(b *testing.B) {

	var startIndex = 1000
	var endIndex = 5000
	var size = endIndex - startIndex
	cw := CashedService{}

	b.StartTimer()
	cw.fibonacciSlice(startIndex, endIndex, size, false, nil)
	b.StopTimer()
}

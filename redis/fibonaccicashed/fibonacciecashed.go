package fibonaccicashed

import (
	"fmt"
	"github.com/Vladimir77715/fibonacciservice/fibonacci"
	"github.com/Vladimir77715/fibonacciservice/redis/client"
	"strconv"
	"time"
)

type CashedService struct {
	rcw *client.RedisClientWrapper
}

func InitCashedService(rcw *client.RedisClientWrapper) *CashedService {
	return &CashedService{rcw: rcw}
}

func (cs *CashedService) getRedisEntityMap(x int, size int) (map[string]*client.RedisEntity, bool) {
	entitySlice := make(map[string]*client.RedisEntity, size)
	var isAllNil = true
	for i := 0; i < size; i++ {
		entity, e := cs.rcw.Get(strconv.Itoa(x + i))
		if e == nil {
			entitySlice[strconv.Itoa(i)] = entity
			isAllNil = false
		}
	}
	return entitySlice, isAllNil
}

func (cs *CashedService) fibonacciSlice(x int, y int, size int, isAllNil bool, entityMap map[string]*client.RedisEntity) ([]uint64, error) {
	var err error
	if !isAllNil {
		var resPointer = make([]*uint64, size)
		var res = make([]uint64, size)
		for i, _ := range resPointer {
			if v, ok := entityMap[strconv.Itoa(x+i)]; ok {
				if vi, e := strconv.ParseUint(v.Value, 10, 64); e == nil {
					resPointer[i] = &vi
					res[i] = vi
				}
			} else {
				var fis *fibonacci.FibonacciInitSection
				if i > 2 && resPointer[i-1] != nil && resPointer[i-2] != nil {
					pos1, pos2 := x+i-1, x+i-2
					first, second := resPointer[i-1], resPointer[i-2]
					Fn := fibonacci.FibonacciNumber{Value: *first, Position: pos2}
					Sn := fibonacci.FibonacciNumber{Value: *second, Position: pos1}
					fis = &fibonacci.FibonacciInitSection{Fn, Sn}
				}
				fSlice, e := fibonacci.FibonacciSlice(x+i, x+i, fis)
				if e == nil {
					res[i] = fSlice[0]
					if cs.rcw != nil {
						err = cs.rcw.Set(strconv.Itoa(x+i), fSlice[0], time.Minute)
					}
				} else {
					err = e
				}
			}
		}
		return res, err
	} else {

		fSlice, err := fibonacci.FibonacciSlice(x, y, nil)
		if err != nil {
			return nil, err
		}
		if cs.rcw != nil {
			for i := 0; i < size; i++ {
				err = cs.rcw.Set(strconv.Itoa(x+i), fSlice[i], time.Minute)
				if err != nil {
					break
				}
			}
		}
		return fSlice, err
	}
}

func (cs *CashedService) FibonacciSlice(x int, y int) ([]uint64, error) {
	var size = y - x + 1
	var entitySlice, isAllNil = cs.getRedisEntityMap(x, size)
	println(entitySlice)
	return cs.fibonacciSlice(x, y, size, isAllNil, nil)
}

func interfaceToUint64(v interface{}) (uint64, error) {
	vv, ee := strconv.ParseUint(fmt.Sprintf("%v", v), 10, 64)
	if ee == nil {
		return vv, nil
	}
	return 0, ee
}

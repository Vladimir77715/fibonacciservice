package main

import (
	"context"
)

var ctx = context.Background()
var (
	addr     = "0.0.0.0:6380"
	password = ""
	DB       = 0
)

func main() {
	//rdb := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6380",
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//})
	//
	//
	//rdb.FlushAll(ctx)
	//// Place data
	//ee := rdb.MSet(ctx,"1","1","2","1")
	//
	//s := rdb.MGet(ctx,"0","1","2","55")

	//cw,err := client.InitRedisClient(&ctx,true)
	//vals , err := cw.MGet("1","2")
	//s := rdb.Get(ctx,"55")
	//println(vals, err)
	//	"entities", &redis.Z{
	//	Score:  float64(0),
	//	Member: "data\nmoredata",
	//}).Err()
	//if err != nil {
	//	panic(err)
	//}

	// Get data
	//keys, v := rdb.ZScan(ctx,"entities", 0, "*", 0).Val()
	//if len(keys) != 2 {
	//	fmt.Printf("expected length of keys to be == 2, got: %v", len(keys))
	//} else {
	//	fmt.Println("Worked! len(keys) == ", len(keys), v)
	//}

	//err = rdb.Set(ctx, "key", "value", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//val, err := rdb.Get(ctx, "key").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key", val)
	//
	//val2, err := rdb.Get(ctx, "key2").Result()
	//if err == redis.Nil {
	//	fmt.Println("key2 does not exist")
	//} else if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("key2", val2)
	//}
}

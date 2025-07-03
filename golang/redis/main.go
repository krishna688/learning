package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	if err := login("kp02"); err != nil {
		log.Println("err", err)
	}

}

func login(key string) error {
	rds := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Default Redis port
	})

	key = fmt.Sprintf("adminuser:%s", key)

	ctx := context.TODO()
	id, err := rds.Incr(ctx, key).Uint64()
	log.Println(id)

	log.Println("count", id)

	if id == 1 {
		rds.Expire(ctx, key, time.Hour)
	}

	// if count == 0 {
	// 	if err := rds.Set(ctx, key, []byte{1}, time.Minute).Err(); err != nil && err != redis.Nil {
	// 		return err
	// 	}

	// 	return nil
	// }
	// if count > 3 {
	// 	return fmt.Errorf("limit reached")
	// }

	return err
}

package main

import (
    "fmt"
    "github.com/go-redis/redis"
)

func main() {
    client := redis.NewClient(&redis.Options{
        Addr:     "r-uf66a1eyqrrvqbcc6r.redis.rds.aliyuncs.com:6379",
        Password: "3Dclobest", // no password set
        DB:       15,  // use default DB
    })

    pong, err := client.Ping().Result()
    fmt.Println(pong, err)

    err = client.Set("key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := client.Get("key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)

    val2, err := client.Get("key2").Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exists")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("key2", val2)
    }
}

package redis

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

type Config struct {
    Host     string `json:"host"`
    Port     uint64 `json:"port"`
    Password string `json:"password"`
    DB       int    `json:"db"`
}

func New(c *Config) *redis.Client {
    r := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%v:%v", c.Host, c.Port),
        Password: c.Password,
        DB:       c.DB,
    })

    if err := r.Ping(context.Background()).Err(); err != nil {
        panic(err)
    } else {
        return r
    }
}

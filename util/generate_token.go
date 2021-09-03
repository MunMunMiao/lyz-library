package util

import (
    "math/rand"
    "strconv"
    "time"
)

func GenerateToken() string {
    encrypted := time.Now().UnixNano() + rand.Int63n(10000)
    return SumSha3([]byte(strconv.FormatInt(encrypted, 10)))
}

package util

import (
    "crypto/md5"
    "crypto/sha256"
    "encoding/hex"
)

func SumSha3(i []byte) string {
	sum := sha256.Sum256(i)
	return hex.EncodeToString(sum[:])
}

func SumMd5(i []byte) string {
	sum := md5.Sum(i)
	return hex.EncodeToString(sum[:])
}

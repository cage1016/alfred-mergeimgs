package lib

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int, charset string) string {
	var result []byte
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		index := rand.Intn(len(charset))
		result = append(result, charset[index])
	}
	return string(result)
}

func RandomImageFilename(length int) string {
	filename := randomString(length, charset)
	return filename + ".png"
}

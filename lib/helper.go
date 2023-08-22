package lib

import (
	"math/rand"
	"strings"
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

var ErrorMessage = []string{
	"error",
	"fd error",
}

func IsFdError(line string) bool {
	for _, msg := range ErrorMessage {
		if strings.Contains(line, msg) {
			return true
		}
	}
	return false
}

package utils

import (
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetEnv(key string, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}

func keyInArray(key string, arr []string) bool {
	for _, value := range arr {
		if value == key {
			return true
		}
	}
	return false
}

func StrToBool(key string) bool {
	return !keyInArray(strings.ToLower(key), []string{"0", "false", "no", "n"})
}

func StrToInt(key string) int64 {
	numVal, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		panic(err)
	}
	return numVal
}

func GenerateUniqID() string {
	return fmt.Sprintf(
		"%x%s",
		sha256.Sum256([]byte(strconv.FormatInt(time.Now().UnixMilli(), 10))),
		strings.ReplaceAll(uuid.New().String(), "-", ""),
	)
}

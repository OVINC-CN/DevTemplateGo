package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"runtime"
)

func loadEnv() {
	_, currentPath, _, ok := runtime.Caller(0)
	if !ok {
		panic("[SystemLoadFailed] LoadEnv")
	}
	envFilePath := fmt.Sprintf("%s/.env", filepath.Dir(filepath.Dir(filepath.Dir(currentPath))))
	if _, err := os.Stat(envFilePath); err == nil {
		_ = godotenv.Load(envFilePath)
	}
}

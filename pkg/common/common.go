package common

import (
	"fmt"
	"os"
	"time"
)

func TimeFormat(t time.Time, full bool) (ret string) {
	if full == true {
		ret = t.Format("2006-01-02 15:04:05")
	} else {
		ret = t.Format("2006-01-02")
	}

	return
}

func CheckAndMkdirAll(pathString string) error {
	_, err := os.Stat(pathString)
	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		err := os.MkdirAll(pathString, 0755)
		if err != nil {
			return fmt.Errorf("mkdirall %s failed", pathString)
		}
	}

	return nil
}

func CheckAndCreateFile(filename string) error {
	_, err := os.Stat(filename)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		f, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("create file %s failed", filename)
		}
		f.Close()
	}

	return nil
}

func GetAppEnv() string {
	// 运行模式: release debug test
	mode := os.Getenv("APP_ENV")
	if mode == "" {
		mode = "debug"
	}

	return mode
}
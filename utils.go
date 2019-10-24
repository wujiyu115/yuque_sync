package main

import (
	"fmt"
	"os"
)

//GenNameSpace gen namespace
func GenNameSpace(config SyncConfig) string {
	return fmt.Sprintf("%s/%s", config.Login, config.Repo)
}

//CreateFile create file
func CreateFile(path string) error {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

package support

import (
	"os"
	"time"
)

func GetFilesList(crt Crt, baseFolder string) []os.DirEntry {
	files, err := os.ReadDir(baseFolder)
	if err != nil {
		//fmt.Printf("%s Error reading folder %s: %v\n", chNormal, baseFolder, err)
		crt.Error("Error reading folder '"+baseFolder+"'", err)
		return nil
	}
	return files
}

func GetTimeStamp() string {
	return time.Now().Format("20060102")
}

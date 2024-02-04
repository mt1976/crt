package clean

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	uuid "github.com/lithammer/shortuuid/v3"
	support "github.com/mt1976/admin_me/support"
)

var itemCount int = 0
var debugMode bool = false
var crt support.Crt

func Run(crtIn support.Crt, debugModeIn bool, cleanPathIn, messageIn string) {

	debugMode = debugModeIn
	crt = crtIn

	crt.Print("Starting file name cleanse [" + crt.Bold(cleanPathIn) + "]")
	crt.Blank()

	baseFolder := "."

	fileList := support.GetFilesList(crt, baseFolder)
	if len(fileList) == 0 {
		crt.Shout(fmt.Sprintf("No files found in folder %s\n", baseFolder))
		return
	}

	crt.Print(fmt.Sprintf("Processing "+crt.Bold("%d")+" files "+messageIn, len(fileList)))
	crt.Blank()

	for _, file := range fileList {

		err := cleanFileName(file, baseFolder)

		if err != nil {
			crt.Error("ERROR processing files", err)
			return
		}
	}
	//fmt.Println(crt.BR(&T))
	crt.Break()

	if itemCount > 0 {
		crt.Print(fmt.Sprintf("Cleaned %d filenames in %s", itemCount, cleanPathIn))
	} else {
		crt.Print(fmt.Sprintf("No files cleaned in %s", cleanPathIn))
	}
}

func cleanFileName(info fs.DirEntry, path string) error {
	//fmt.Printf("%s Processing file %s\n", support.PFX, info.Name())
	cleanName, err := getCleanName(info.Name())
	if err != nil {
		crt.Error("Error cleaning file name '"+info.Name()+"'", err)
		return err
	}
	//fmt.Printf("%s Cleaned file name %s\n", support.PFX, cleanName)
	if cleanName == "OnlyFans.mp4" {
		// Rename the file to OnlyFans_Date_Time.mp4
		id := uuid.New()

		cleanName = "OnlyFans_" + time.Now().Format("060102150405") + id + ".mp4"
	}
	if err == nil && cleanName != info.Name() {
		renameFile(path, cleanName, info.Name())
		itemCount++
	}
	return nil
}

func getCleanName(fileName string) (string, error) {
	//fmt.Printf("%s Cleaning file name '%s'\n", support.PFX, name)
	newFileName := fileName

	// Remove all characters that are not in the ValidChars list
	for _, c := range fileName {
		if !strings.Contains(strings.Join(support.ValidFileNameCharacters, ""), string(c)) {
			newFileName = strings.ReplaceAll(newFileName, string(c), "")
		}
	}
	newFileName = strings.ReplaceAll(newFileName, "_", " ")
	newFileName = strings.ReplaceAll(newFileName, "-", " ")

	// Remove all double spaces
	newFileName = support.TrimRepeatingCharacters(newFileName, " ")
	newFileName = support.TrimRepeatingCharacters(newFileName, ".")
	newFileName = support.TrimRepeatingCharacters(newFileName, "-")
	newFileName = support.TrimRepeatingCharacters(newFileName, "*")
	newFileName = strings.TrimLeft(newFileName, " ")
	newFileName = strings.TrimLeft(newFileName, "-")
	newFileName = strings.TrimLeft(newFileName, " ")
	newFileName = strings.TrimLeft(newFileName, "-")
	//fmt.Printf("%s New file name '%s'\n", support.PFX, newName)
	return newFileName, nil
}

func renameFile(path string, newFileName string, oldFileName string) {
	newPath := filepath.Join(filepath.Dir(path), newFileName)
	oldPath := filepath.Join(filepath.Dir(path), oldFileName)
	err := error(nil)

	if !debugMode {
		err = os.Rename(oldPath, newPath)
	}

	if err != nil {
		crt.Error("Failed to rename file '"+path+"'", err)
	} else {
		crt.Print(fmt.Sprintf("Renamed file [%s -> %s]", oldFileName, newPath))
	}
}

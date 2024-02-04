package tidy

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/mt1976/admin_me/support"
	"github.com/ricochet2200/go-disk-usage/du"
)

var (
	fileExtensions = []string{"nfo", "jpeg", "jpg", "bif", "vob", "txt", "png", "me", "exe"}
)

var debugMode bool = true
var crt support.Crt

func Run(crtIn support.Crt, debugModeIn bool, pathIn string) {
	debugMode = debugModeIn
	crt = crtIn

	crt.Print("Media Management - Tidy Folders")

	if len(pathIn) < 1 {
		crt.Error("No path argument supplied", nil)
		return
	}

	//path := os.Args[1]

	if _, err := os.Stat(pathIn); os.IsNotExist(err) {
		crt.Error("The path provided is not valid", err)
		return
	}

	if pathIn == "/" || pathIn == "~" {
		crt.Error("The path provided is the root or home directory", nil)
		return
	}

	if !debugMode {
		crt.Special(crt.Bold(crt.Underline("This is a live run. Files will be deleted.")))
	} else {
		crt.Print(crt.Underline("This is a trial run. Files & Folders will not be deleted."))
	}

	crt.Print("Resolved path: " + realpath(pathIn))
	//fmt.Printf("%s File types to be removed: [%s]\n", support.CHnormal, strings.Join(types, " "))
	crt.Print("Starting file removal process for " + crt.Bold(strings.Join(fileExtensions, " ")))
	var userResponse string
	//fmt.Printf("%s Are you sure you want to proceed? %s(y/n) : %s", PFY, bold, normal)
	crt.Input("Are you sure you want to proceed", "y/n")
	fmt.Scanln(&userResponse)

	if strings.ToLower(userResponse) != "y" && strings.ToLower(userResponse) != "yes" {
		//fmt.Printf("%s Exiting\n", PFY)
		crt.Print(crt.Bold("Exiting"))
		return
	}
	crt.Blank()
	diskSizeTotalBefore, diskSizeFreeBefore, diskPercentUsedBefore := getDiskInfo(pathIn)

	//fmt.Printf("%s Changing directory to %s\n", PFY, path)
	crt.Print("Changing directory to " + crt.Bold(pathIn))
	crt.Blank()
	err := os.Chdir(pathIn)
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to change directory: %v", PFY, err))
		crt.Error("Unable to change directory", err)
		return
	}

	var wg sync.WaitGroup

	for _, fileExtension := range fileExtensions {
		wg.Add(1)
		//fmt.Printf("%s Operation on .%s files completed in %s seconds\n", support.CHspecial, fileExt, runtime)

		defer wg.Done()
		go processFileTypes(fileExtension)

	}
	wg.Wait()

	//fmt.Printf("%s Deleting empty directories\n", PFY)
	crt.Special("Deleting empty directories")
	startLoopIteration := time.Now()
	if !debugMode {
		removeEmptyDirectories()
	} else {
		findEmptyDirectories()
	}
	endLoopIteration := time.Now()
	runtime := endLoopIteration.Sub(startLoopIteration)
	crt.Special("Deletion of empty folders completed in " + crt.Bold(runtime.String()) + " seconds")
	diskSizeTotalAfter, diskSizeFreeAfter, diskPercentUsedAfter := getDiskInfo(pathIn)

	printStorageReport(pathIn, diskSizeTotalBefore, diskSizeFreeBefore, diskPercentUsedBefore, diskSizeTotalAfter, diskSizeFreeAfter, diskPercentUsedAfter)
}

func processFileTypes(fileExtension string) {
	startTime := time.Now()

	if !debugMode {
		crt.Special("Removing all files with extension ." + crt.Bold(fileExtension))
		removeFiles(fileExtension)
	} else {
		crt.Special("Finding all files with extension ." + crt.Bold(fileExtension))
		findFiles(fileExtension)
	}
	endTime := time.Now()
	runtime := endTime.Sub(startTime)

	crt.Special("Operation on ." + crt.Bold(fileExtension) + " files completed in " + crt.Bold(runtime.String()) + "")
	//crt.Blank()
}

func realpath(path string) string {

	realPathCmd := exec.Command("realpath", path)
	output, err := realPathCmd.Output()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to resolve path: %v", PFY, err))
		crt.Error("Unable to resolve path", err)
		return ""
	}
	return strings.TrimSpace(string(output))
}

// The function "getDiskInfo" returns the total disk size, free disk space, and percentage of disk
// space used for a given path.
func getDiskInfo(path string) (total, free, percentUsed string) {
	info := du.NewDiskUsage(path)
	total = support.DiskSizeHuman(info.Size())
	free = support.DiskSizeHuman(info.Available())
	percentUsed = support.Human(info.Usage())

	return total, free, percentUsed
}

func removeFiles(fileExtension string) {
	if debugMode {
		//fmt.Printf("%s DEBUG: Would have removed files\n", PFY)
		crt.Print("DEBUG: Would have removed files")
		return
	}
	findCmd := exec.Command("find", ".", "-type", "f", "-name", "*."+fileExtension, "-exec", "rm", "-f", "{}", ";")
	err := findCmd.Run()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to remove files: %v", PFY, err))
		crt.Error("Unable to remove files", err)
		return
	}
}

func findFiles(fileExt string) {
	findCmd := exec.Command("find", ".", "-type", "f", "-name", "*."+fileExt)
	output, err := findCmd.Output()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to find files: %v", PFY, err))
		crt.Error("Unable to find files", err)
		return
	}

	crt.Spool(output)
}

func removeEmptyDirectories() {
	if debugMode {
		//fmt.Printf("%s DEBUG: Would have removed empty directories\n", PFY)
		crt.Print("DEBUG: Would have removed empty directories")
		return
	}
	findCmd := exec.Command("find", ".", "-type", "d", "-exec", "rmdir", "{}", "+")
	err := findCmd.Run()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to remove empty directories: %v", PFY, err))
		crt.Error("Unable to remove empty directories", err)
		return
	}
}

func findEmptyDirectories() {
	findCmd := exec.Command("find", ".", "-type", "d", "-empty", "-print")
	output, err := findCmd.Output()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to find empty directories: %v", PFY, err))
		crt.Error("Unable to find empty directories", err)
		return
	}

	crt.Spool(output)
}

func printStorageReport(path, beforeDiskSizeTotal, beforeDiskSizeFree, beforeDiskPercentUsed, afterDiskSizeTotal, afterDiskSizeFree, afterDiskPercentUsed string) {
	crt.Break()
	crt.Print(crt.Bold(crt.Underline("STORAGE REPORT")))
	crt.Break()
	crt.Print("BEFORE  : " + crt.Bold(beforeDiskSizeFree) + " available out of " + crt.Bold(beforeDiskSizeTotal) + " total (" + crt.Bold(beforeDiskPercentUsed) + "% used)")
	crt.Print("AFTER   : " + crt.Bold(afterDiskSizeFree) + " available out of " + crt.Bold(afterDiskSizeTotal) + " total (" + crt.Bold(afterDiskPercentUsed) + "% used)")
	crt.Print("MACHINE : " + support.GetSystemInfo(crt))
	crt.Print("HOST    : " + support.GetHostname(crt))
	crt.Print("USER    : " + support.GetUsername(crt))
	mode := "DEBUG"
	if !debugMode {
		mode = "LIVE"
	}
	crt.Print("MODE    : " + mode)
	crt.Print("TYPES   : " + strings.Join(fileExtensions, " "))
	crt.Print("END     : " + time.Now().Format("02/01/2006 15:04:05"))
}

package filechooser

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/mt1976/crt"
	errs "github.com/mt1976/crt/errors"
	lang "github.com/mt1976/crt/language"
)

type File struct {
	Name     string
	Path     string
	Created  string
	Modified string
	Size     int64
	SizeTxt  string
	IsDir    bool
	Icon     string
	Mode     string
	Seq      int
}

type flagger struct {
	allowDirs    bool
	allowFiles   bool
	showDotFiles bool
	showFiles    bool
	showDirs     bool
}

var All = flagger{allowDirs: true, allowFiles: true, showDotFiles: true, showFiles: true, showDirs: true}
var DirectoriesOnly = flagger{allowDirs: true, allowFiles: false, showDotFiles: false, showFiles: false, showDirs: true}
var FilesOnly = flagger{allowDirs: false, allowFiles: true, showDotFiles: false, showFiles: true, showDirs: true}
var DirectoriesAll = flagger{allowDirs: true, allowFiles: false, showDotFiles: true, showFiles: false, showDirs: true}
var FilesAll = flagger{allowDirs: false, allowFiles: true, showDotFiles: true, showFiles: true, showDirs: true}

// var actionUp = "U"
// var actionUpDoubleDot = ".."
// var actionUpArrow = "^"
// var actionGo = "G"
var pathSeparator = string(os.PathSeparator)

// var actionSelect = "S"

// FileChooser is a function to choose a file or directory using the file chooser.
//
// Parameters:
//
//	searchPath (string): The directory to start browsing from.
//	flags (flagger): A set of flags that control which types of items are included in the list and other behavior.
//
// Returns:
//
//	(string, bool, error): The chosen file or directory, a boolean indicating whether it is a directory, and an error if one occurred.
func FileChooser(searchPath string, flags flagger) (string, bool, error) {

	if searchPath == "" {
		return "", false, errs.ErrInvalidPathSpecialDirectory
	}
	if searchPath == "." {
		// get the real wd directory
		usr, err := user.Current()
		if err != nil {
			return "", false, err
		}
		searchPath = usr.HomeDir
	}

	// Function to choose a file or directory using the file chooser
	term := crt.New()
	page := term.NewPage(lang.TxtFileChooserTitle)

	// Get a list of files in the specified directory
	files, err := GetFolderContent(searchPath, flags)
	if err != nil {
		return "", false, err
	}

	// Add information about the current user and directory to the page
	uh, _ := UserHome()
	un, _ := UserName()
	page.AddFieldValuePair("User Name", un)
	page.AddFieldValuePair("User Home", uh)
	page.AddFieldValuePair("Directory", searchPath)

	// Add a blank row to separate the header from the file list
	//page.AddBreakRow()
	brk(page, "-")

	// Add columns for the file list
	//page.AddColumnsTitle("Icon", "Name", "Mode", "Size", "Modified")

	//sufx := "%1s| %-30s | %-10s | %-12s | %-15s"
	format := "%-4s) %-1s  %-30s | %-10s | %-12s | %-15s"
	head := "%-4s| %-1s| %-30s | %-10s | %-12s | %-15s"
	// Add a title row for the file list
	title := fmt.Sprintf(head, " ", "T", "Name", "Mode", "Modified", "Size")
	page.Add(title, "", "")

	// Add a row for a separator between the header and the file list
	brk(page, "+")
	//page.AddBreakRow()

	// Add an option for the parent directory
	up := fmt.Sprintf(format, lang.Up.Action(), "", "..", "", "", "")
	page.Add(up, "", "")

	// Add actions for the parent directory, up arrow, and select
	page.AddAction(lang.Up)
	page.AddAction(lang.UpArrow)
	page.AddAction(lang.UpDoubleDot)
	page.AddAction(lang.Select)

	// Add options for each file or directory in the list
	for _, file := range files {
		// Create a row for the file or directory
		row := fmt.Sprintf("%-1s %-30s | %-10s | %-12s | %-15s", file.Icon, file.Name, file.Mode, file.Modified, file.SizeTxt)
		page.AddMenuOption(file.Seq+1, row, "", "")

		// Add an action for selecting the directory if it is a directory
		if file.IsDir {
			act := lang.NewAction(lang.Go.Action() + fmt.Sprintf("%v", file.Seq+1))
			page.AddAction(act)
		}
	}

	// Display the file chooser with actions
	nextAction := page.Display_Actions()
	if nextAction.Is(lang.Quit) {
		return "", false, nil
	}
	if nextAction.Is(lang.Select) {
		// The current folder has been selected
		return searchPath, true, nil
	}
	page.Dump(nextAction.Action(), lang.Up.Action(), lang.UpArrow.Action(), lang.UpDoubleDot.Action())
	// Handle actions for the parent directory, up arrow, and select
	if nextAction.Is(lang.Up) || nextAction.Is(lang.UpArrow) || nextAction.Is(lang.UpDoubleDot) {
		page.Dump("Up One Directory", searchPath, pathSeparator)
		upPath := strings.Split(searchPath, pathSeparator)
		page.Dump(fmt.Sprintf("b4 upPath: %v\n", upPath))

		if len(upPath) > 1 {
			upPath = upPath[:len(upPath)-1]
		}
		page.Dump(fmt.Sprintf("af upPath: %v\n", upPath))
		toPath := strings.Join(upPath, pathSeparator)
		page.Dump("Relaunch FileChooser", toPath, lang.Up.Action(), lang.UpArrow.Action(), lang.UpDoubleDot.Action())
		return FileChooser(toPath, flags)
	}

	// Split the action into its first character and the remaining characters
	first := upcase(nextAction.Action()[:1])
	remainder := nextAction.Action()[1:]

	// Handle actions for selecting a directory or file
	if lang.Go.Equals(first) && isInt(remainder) {
		r := files[term.Helpers.ToInt(remainder)-1]
		if !r.IsDir {
			page.Error(errs.ErrNotADirectory, r.Path)
			return FileChooser(searchPath, flags)
		}
		page.Dump("Drilldown", r.Path, first, lang.Go.Action())
		return FileChooser(r.Path, flags)
	}

	// Handle selection of a specific file or directory
	if nextAction.IsInt() {
		r := files[term.Helpers.ToInt(nextAction.Action())-1]
		if !r.IsDir && flags.allowDirs {
			page.Error(errs.ErrNotAFile, r.Path)
			return FileChooser(searchPath, flags)
		}
		if r.IsDir && flags.allowFiles {
			page.Error(errs.ErrNotADirectory, r.Path)
			return FileChooser(searchPath, flags)
		}
		return r.Path, r.IsDir, nil
	}

	return FileChooser(searchPath, flags)
}

func brk(page *crt.Page, breakChar string) {
	brk := "%-4s+%-2s+%-31s-+-%-10s-+-%-12s-+-%-5s"
	//replace + char with the breakChar
	brk = strings.Replace(brk, "+", breakChar, -1)
	breaker := fmt.Sprintf(brk, strings.Repeat("-", 4), strings.Repeat("-", 2), strings.Repeat("-", 31), strings.Repeat("-", 10), strings.Repeat("-", 12), strings.Repeat("-", 5))
	page.Add(breaker, "", "")
}

// GetFolderContent gets a list of files in the specified directory.
// It filters the list of files to only include directories that match the include flags.
// It returns a slice of File structs that contain information about each file or directory.
func GetFolderContent(dir string, include flagger) ([]File, error) {
	// Get a list of files in the specified directory
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// Filter the list of files to only include directories
	var directories []File

	itemNo := 0

	for _, file := range files {

		// Check if the file is a directory and should be included
		if file.IsDir() && !include.showDirs {
			continue
		}

		// Check if the file is a hidden file and should be included
		if file.Name()[0] == '.' && !include.showDotFiles {
			continue
		}

		// Check if the file is a regular file and should be included
		if !file.IsDir() && !include.allowFiles {
			continue
		}

		// Check if the file should be shown
		if !file.IsDir() && !include.showFiles {
			continue
		}

		// Create a new File struct and fill it with information about the file
		var this File
		this.Name = strings.Trim(file.Name(), " ")
		this.Path = dir + pathSeparator + file.Name()
		inf, _ := file.Info()
		this.Created = "N/A"
		this.Modified = crt.New().Formatters.HumanFromUnixDate(inf.ModTime().Local().Unix())
		this.Size = inf.Size()
		yy := fmt.Sprintf("%v", this.Size)
		this.SizeTxt = yy
		this.Mode = inf.Mode().String()
		this.IsDir = file.IsDir()
		if this.IsDir {
			this.Icon = lang.TxtFolderIcon
		} else {
			this.Icon = lang.TxtFileIcon
		}
		// Check if the file is a symbolic link
		if isSymLink(this.Mode) {
			this.Icon = lang.TxtSymLinkIcon
		}
		this.Icon = this.Icon + " "
		this.Seq = itemNo
		// Add the file to the list of directories
		directories = append(directories, this)
		itemNo++
	}
	return directories, nil
}

// isSymLink returns true if the input string can be converted to an integer.
func isSymLink(mode string) bool {
	return mode[0] == 'L' || mode[0] == 'l'
}

// ChooseDirectory is a function to choose a directory using the file chooser.
//
// Parameters:
//
//	root (string): The root directory to start browsing from.
//
// Returns:
//
//	(string, error): The chosen directory, or an error if one occurred.
func ChooseDirectory(root string) (string, error) {
	// Function to choose a directory using the file chooser
	item, _, err := FileChooser(root, DirectoriesOnly)
	if err != nil {
		return "", err
	}
	return item, err
}

// isInt returns true if the input string can be converted to an integer.
func isInt(s string) bool {
	return crt.New().Helpers.IsInt(s)
}

// upcase returns the input string with all characters converted to uppercase.
func upcase(s string) string {
	return crt.New().Formatters.Upcase(s)
}

// UserHome returns the home directory of the current user, or an error if it cannot be determined.
func UserHome() (string, error) {
	// Function gets the home directory of the current user, or returns an error if it cant.
	//
	// Returns:
	// The home directory of the current user, or an error if it cant.
	return os.UserHomeDir()
}

// UserName returns the name of the current user, or an error if it cannot be determined.
func UserName() (string, error) {
	// Get the current user
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	// Return the username
	return currentUser.Name, nil
}

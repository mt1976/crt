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
	directory bool
	file      bool
	dotfile   bool
	showFiles bool
}

var All = flagger{directory: true, file: true, dotfile: true, showFiles: true}
var DirectoriesOnly = flagger{directory: true, file: false, dotfile: false, showFiles: false}
var FilesOnly = flagger{directory: false, file: true, dotfile: false, showFiles: true}
var DirectoriesAll = flagger{directory: true, file: false, dotfile: true, showFiles: false}
var FilesAll = flagger{directory: false, file: true, dotfile: true, showFiles: true}

var actionUp = "U"
var actionUpDoubleDot = ".."
var actionUpArrow = "^"
var actionGo = "G"
var pathSeparator = string(os.PathSeparator)
var actionSelect = "S"

func FileChooser(searchPath string, flags flagger) (string, bool, error) {

	term := crt.New()
	page := term.NewPage(lang.TxtFileChooserTitle)

	files, err := GetFolderList(searchPath, flags)
	if err != nil {
		return "", false, err
	}
	uh, _ := UserHome()
	un, _ := UserName()
	page.AddFieldValuePair("User Name", un)
	page.AddFieldValuePair("User Home", uh)
	page.AddFieldValuePair("Directory", searchPath)
	page.AddBlankRow()
	//page.AddColumnsTitle("Name", "Mode", "Size", "Modified")
	formatter := "%1v %-30v %10v %12v %15v"
	tformatter := "%5v %1v-|%-30v|-%-10v-|-%-12v-|-%-15v"
	title := fmt.Sprintf(tformatter, "", "T", " Name", "Mode", "Modified", "Size")
	title = strings.ReplaceAll(title, "|", " ")
	title = strings.ReplaceAll(title, "-", " ")

	page.Add(title, "", "")
	breaker := fmt.Sprintf(tformatter, strings.Repeat(" ", 5), strings.Repeat("-", 1), strings.Repeat("-", 30), strings.Repeat("-", 10), strings.Repeat("-", 12), strings.Repeat("-", 15))
	//breaker = strings.Repeat("-", len(breaker))
	page.Add(breaker, "", "")
	up := fmt.Sprintf(formatter, "^", " ..", "", "", "")
	page.Add(actionUp+" "+up, "", "")
	page.AddAction(actionUp)
	page.AddAction(actionUpArrow)
	page.AddAction(actionUpDoubleDot)
	page.AddAction(actionSelect)

	for _, file := range files {
		row := fmt.Sprintf(formatter, file.Icon, file.Name, file.Mode, file.Modified, file.SizeTxt)
		page.AddMenuOption(file.Seq+1, row, "", "")
		if file.IsDir {
			page.AddAction(actionGo + fmt.Sprintf("%v", file.Seq+1))
		}
	}
	// spew.Dump(page)
	// os.Exit(0)
	na, _ := page.DisplayWithActions()
	if na == lang.SymActionQuit {
		return "", false, nil
	}

	if na == actionUp || na == actionUpArrow || na == actionUpDoubleDot {
		upPath := strings.Split(searchPath, pathSeparator)
		if len(upPath) > 1 {
			upPath = upPath[:len(upPath)-1]
		}
		toPath := strings.Join(upPath, pathSeparator)

		return FileChooser(toPath, flags)
	}
	// split na into first char and remainder
	first := upcase(na[:1])
	remainder := na[1:]
	if first == actionGo || isInt(remainder) {
		r := files[term.Helpers.ToInt(remainder)-1]
		if !r.IsDir {
			page.Error(errs.ErrNotADirectory, r.Path)
			return FileChooser(searchPath, flags)
		}
		return FileChooser(r.Path, flags)
	}

	if term.Helpers.IsInt(na) {
		// if a specific item has been selected, return the path of that item
		r := files[term.Helpers.ToInt(na)-1]
		if !r.IsDir && flags.directory {
			page.Error(errs.ErrNotAFile, r.Path)
			return FileChooser(searchPath, flags)
		}
		if r.IsDir && flags.file {
			page.Error(errs.ErrNotADirectory, r.Path)
			return FileChooser(searchPath, flags)
		}
		return r.Path, r.IsDir, nil
	}
	if upcase(na) == upcase(actionSelect) {
		// The current folder has been selected
		return searchPath, true, nil
	}

	return FileChooser(searchPath, flags)
}

func GetFolderList(dir string, include flagger) ([]File, error) {
	// Get a list of files in the specified directory
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// Filter the list of files to only include directories
	var directories []File
	//include := false
	itemNo := 0
	// upPath := strings.Split(dir, "/")
	// if len(upPath) > 1 {
	// 	upPath = upPath[:len(upPath)-1]
	// }
	//upPath = strings.Join(upPath, "/")
	//up := File{Name: "..", Path: strings.Join(upPath, "/"), Icon: lang.TxtFolderIcon, IsDir: true, Seq: -1}
	//directories = append(directories, up)
	for _, file := range files {
		// fmt.Printf("\"PROCESSING\": %v\n", file.Name())
		// fmt.Printf("file.IsDir(): %v\n", file.IsDir())
		// fmt.Printf("include.directory: %v\n", include.directory)
		if file.IsDir() && !include.directory {
			//include = true
			continue
		}
		// fmt.Printf("file.Name()[0]: %v\n", file.Name()[0])
		// fmt.Printf("include.dotfile: %v\n", include.dotfile)
		if file.Name()[0] == '.' && !include.dotfile {
			//include = false
			continue
		}
		// fmt.Printf("file.IsDir(): %v\n", file.IsDir())
		// fmt.Printf("include.file: %v\n", include.file)
		if !file.IsDir() && !include.file {
			//include = false
			continue
		}
		// fmt.Printf("file.IsDir(): %v\n", file.IsDir())
		// fmt.Printf("include.showFiles: %v\n", include.showFiles)
		if !file.IsDir() && !include.showFiles {
			//include = false
			continue
		}

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
		if isSymLink(this.Mode) {
			this.Icon = lang.TxtSymLinkIcon
		}
		this.Icon = this.Icon + " "
		this.Seq = itemNo
		//fmt.Printf("this: %v\n", this)
		directories = append(directories, this)
		itemNo++
	}
	//spew.Dump(directories, include, All)
	//os.Exit(0)
	return directories, nil
}

func isSymLink(mode string) bool {
	return mode[0] == 'L' || mode[0] == 'l'
}

func ChooseDirectory(root string) (string, error) {
	// Function to choose a directory using the file chooser
	item, _, err := FileChooser(root, DirectoriesOnly)
	if err != nil {
		return "", err
	}
	return item, err
}

func isInt(s string) bool {
	//_, err := crt.New().Helpers.IsInt(s)
	return crt.New().Helpers.IsInt(s)
}

func upcase(s string) string {
	return crt.New().Formatters.Upcase(s)
}

func UserHome() (string, error) {
	// Function gets the home directory of the current user, or returns an error if it cant.
	//
	// Returns:
	// The home directory of the current user, or an error if it cant.
	return os.UserHomeDir()
}

func UserName() (string, error) {
	// Function gets the name of the current user, or returns an error if it cant.
	//
	// Returns:
	// The name of the current user, or an error if it cant.
	currentUser, err := user.Current()
	return currentUser.Name, err
}

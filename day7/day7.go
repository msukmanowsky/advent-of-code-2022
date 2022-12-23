package day7

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/msukmanowsky/advent-of-code-2022/utils"
)

type File struct {
	Dir  *Dir
	Name string
	Size uint
}

func (f File) Path() string {
	dirPath := f.Dir.Path()
	return dirPath + f.Name
}

type Dir struct {
	Parent *Dir
	Name   string
	Files  map[string]File
	Dirs   map[string]Dir
}

func NewDir() *Dir {
	return &Dir{
		Files: make(map[string]File),
		Dirs:  make(map[string]Dir),
	}
}

func (d Dir) Size() uint {
	var size uint = 0
	for _, file := range d.Files {
		size += file.Size
	}
	for _, dir := range d.Dirs {
		size += dir.Size()
	}
	return size
}

func (d Dir) Path() string {
	if d.Parent == nil {
		return "/"
	}

	components := []string{d.Name}
	dir := d.Parent
	for dir != nil && dir.Parent != nil {
		components = append([]string{dir.Name}, components...)
		dir = dir.Parent
	}

	return "/" + strings.Join(components, "/")
}

type DirWalkFunc func(d Dir)

func WalkDirs(dir Dir, walkFn DirWalkFunc) {
	walkFn(dir)
	for _, d := range dir.Dirs {
		WalkDirs(d, walkFn)
	}
}

type Command struct {
	Name   string
	Args   []string
	Output string
}

var commandSep = []byte("\n$")

// A custom split function for parsing commands
func SplitCommand(data []byte, atEOF bool) (int, []byte, error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, commandSep); i >= 0 {
		// New command found, advance past the newline character, but not the
		// $ marker so it can be included in the next command
		return i + 1, data[:i], nil
	}

	// We're at EOF, return command as is
	if atEOF {
		return len(data), data, nil
	}

	// Request more data
	return 0, nil, nil
}

func parseCommand(rawCmd string) (*Command, error) {
	// first line will be the command itself
	rawCmd = strings.TrimPrefix(rawCmd, "$ ")
	lines := strings.Split(rawCmd, "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("empty command: \"%v\"", rawCmd)
	}
	parts := strings.Split(lines[0], " ")
	name := parts[0]
	args := parts[1:]
	output := strings.Join(lines[1:], "\n")
	return &Command{
		Name:   name,
		Args:   args,
		Output: output,
	}, nil
}

func buildTree(r io.Reader) (rootDir *Dir, err error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(SplitCommand)
	rootDir = NewDir()
	pwd := rootDir
	for scanner.Scan() {
		rawCmd := scanner.Text()
		cmd, err := parseCommand(rawCmd)
		if err != nil {
			return nil, fmt.Errorf("error parsing command: \"%s\"", rawCmd)
		}

		switch cmd.Name {
		case "cd":
			path := cmd.Args[0]
			var dir *Dir
			if path == "/" {
				dir = rootDir
			} else if path == ".." {
				dir = pwd.Parent
			} else {
				d, ok := pwd.Dirs[path]
				if !ok {
					return nil, fmt.Errorf("no dir named %s found in %s", path, pwd.Path())
				}
				dir = &d
			}
			pwd = dir
		case "ls":
			lines := strings.Split(cmd.Output, "\n")
			for _, line := range lines {
				if strings.HasPrefix(line, "dir ") {
					// Directory
					name := strings.TrimPrefix(line, "dir ")
					dir := NewDir()
					dir.Parent = pwd
					dir.Name = name
					pwd.Dirs[name] = *dir
				} else {
					// File
					parts := strings.SplitN(line, " ", 2)
					size, err := strconv.ParseUint(parts[0], 10, 32)
					if err != nil {
						return nil, fmt.Errorf("invalid filesize: %s", line)
					}
					name := parts[1]
					file := File{
						Name: name,
						Size: uint(size),
						Dir:  pwd,
					}
					pwd.Files[name] = file
				}
			}
		default:
			return nil, fmt.Errorf("invalid command: %v", cmd.Name)
		}
	}
	return rootDir, nil
}

func printTree(d Dir) {
	WalkDirs(d, func(d Dir) {
		log.Printf("%s", d.Path())
	})
}

func Day7_1() {
	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()

	totalSize := uint(0)
	dirMaxSize := uint(100000)
	rootDir, err := buildTree(file)
	if err != nil {
		log.Fatalln(err)
	}

	WalkDirs(*rootDir, func(d Dir) {
		size := d.Size()
		if size <= dirMaxSize {
			totalSize += size
		}
	})
	log.Printf("Total size of directories <= %d: %d\n", dirMaxSize, totalSize)
}

type DirsBySize []Dir

func (s DirsBySize) Len() int           { return len(s) }
func (s DirsBySize) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s DirsBySize) Less(i, j int) bool { return s[i].Size() < s[j].Size() }

func Day7_2() {
	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()

	rootDir, err := buildTree(file)
	if err != nil {
		log.Fatalln(err)
	}

	dirs := []Dir{}
	WalkDirs(*rootDir, func(d Dir) {
		dirs = append(dirs, d)
	})
	sort.Sort(DirsBySize(dirs))

	// Find the smallest directory that, if deleted, would free up enough space
	// on the filesystem to run the update. What is the total size of that
	// directory?
	totalDiskSpace := uint(70000000)
	updateSpaceNeeded := uint(30000000)
	freeSpace := totalDiskSpace - rootDir.Size()
	freeSpaceNeeded := updateSpaceNeeded - freeSpace

	var toDelete *Dir
	for _, dir := range dirs {
		size := dir.Size()
		log.Printf("%s (%d)", dir.Path(), size)
		if size >= freeSpaceNeeded {
			toDelete = &dir
			break
		}
	}

	if toDelete == nil {
		log.Fatalf("no directory found that would free up %d", freeSpaceNeeded)
	}

	log.Printf("To free up %d, directory %s should be deleted (%d)\n", freeSpaceNeeded, toDelete.Path(), toDelete.Size())
}

package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func IsFile(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	} else if stat.IsDir() {
		return false, nil
	} else {
		return true, nil
	}
}

func IsSymlink(path string) (bool, error) {
	stat, err := os.Readlink(path)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func IsDir(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	} else if stat.IsDir() {
		return true, nil
	} else {
		return false, nil
	}
}

func FileListing(endPattern string) []map[string]string {
	var paths []map[string]string

	filepath.Walk(
		".",
		func(path string, info os.FileInfo, err error) error {
			if strings.HasSuffix(path, endPattern) {
				dirname := filepath.Dir(path)
				linked, _ := os.Readlink(path)
				linked = strings.TrimSuffix(linked, ".placeholder")

				oldBase := filepath.Join(dirname, linked)
				oldPath, _ := filepath.Abs(oldBase)

				newBase := strings.Replace(path, "locations/", "", 1)
				newBase = strings.TrimSuffix(newBase, endPattern)

				currentUser, _ := user.Current()
				newPath, _ := filepath.Abs(
					filepath.Join(
						currentUser.HomeDir,
						"."+newBase,
					),
				)

				currentPath := make(map[string]string)

				currentPath["old_path"] = oldPath
				currentPath["new_path"] = newPath
				paths = append(paths, currentPath)
			}
			return nil
		},
	)

	return paths
}

func DirectoryListing() []map[string]string {
	return FileListing(".directory.symlink")
}

func LocalsListing() []map[string]string {
	var files []map[string]string
	allFiles := FileListing(".symlink")
	for i := range allFiles {
		if strings.HasSuffix(allFiles[i]["new_path"], ".local") {
			files = append(files, allFiles[i])
		}
	}

	return files
}

func DotfileListing() []map[string]string {
	var files []map[string]string
	allFiles := FileListing(".symlink")
	for i := range allFiles {
		if !strings.HasSuffix(allFiles[i]["new_path"], ".directory") &&
			!strings.HasSuffix(allFiles[i]["new_path"], ".local") {
			files = append(files, allFiles[i])
		}
	}

	return files
}

func CurrentFile(f map[string]string) map[string]string {
	_, newPath := f["old_path"], f["new_path"]
	if d, _ := IsDir(newPath); d {
		return map[string]string{
			"type": "directory",
			"path": newPath,
			"old_path": newPath,
			"new_path": newPath,
		}
	} else if d, _ := IsSymlink(newPath); d {
		linked, _ := os.Readlink(newPath)
		return map[string]string{
			"type": "symlink",
			"path": newPath,
			"old_path": linked,
			"new_path": newPath,
		}
	} else {
		return map[string]string{
			"type": "file",
			"path": newPath,
			"old_path": newPath,
			"new_path": newPath,
		}
	}
}

func MaybeOverwriteSymlink(oldLink, newLink, noInput) {
	if noInput {
		fmt.Println(
			"Warn!  | Overwrite | path: " +
			newLink["new_path"] + ", " +
			"old: " + oldLink["old_path"] + ", " +
			"new: " + newLink["old_path"],
		)
	}
}

func main() {
	IsSymlink("kdljaflksd.go")
	IsSymlink("locations/bash_profile.symlink")
	files := FileListing(".symlink")
	for i := range files {
		fmt.Printf(files[i]["old_path"])
		fmt.Printf(" | ")
		fmt.Println(files[i]["new_path"])
	}
}

package main

import (
	"path/filepath"
	"fmt"
	"os"
	"os/user"
	"strings"
)


func FileListing(endPattern string) []map[string]string {
	var paths []map[string]string

	filepath.Walk(
		".",
		func(path string, info os.FileInfo, err error) error {
			if strings.HasSuffix(path, endPattern) {
				dirname        := filepath.Dir(path)
				linked, _      := os.Readlink(path)
				linked          = strings.TrimSuffix(linked, ".placeholder")

				oldBase        := filepath.Join(dirname, linked)
				oldPath, _     := filepath.Abs(oldBase)

				newBase        := strings.Replace(path, "locations/", "", 1)
				newBase         = strings.TrimSuffix(newBase, endPattern)

				currentUser, _ := user.Current()
				newPath, _ := filepath.Abs(
					filepath.Join(
						currentUser.HomeDir,
						"." + newBase,
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


func main() {
	files := FileListing(".symlink")
	for i := range files {
		fmt.Printf(files[i]["old_path"])
		fmt.Printf(" | ")
		fmt.Println(files[i]["new_path"])
	}
}

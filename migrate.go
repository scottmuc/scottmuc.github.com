package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func MigratePost(path string, fileInfo os.FileInfo, err error) error {
	if fileInfo.IsDir() {
		return nil
	}

	octopressFilename := filepath.Base(path)
	hugoFilename := strings.Replace(octopressFilename[11:], ".markdown", ".md", -1)
	hugoFilePath := "hugo/content/blog/" + hugoFilename
	fmt.Printf("%s -> %s\n", path, hugoFilePath)


	octopressFileContents, e := os.ReadFile(path)
	if e != nil {
		log.Fatal(e)
	}

	frontMatter := strings.Split(string(octopressFileContents), "---")[1]
	fmt.Println(frontMatter)

	return nil
}

func main() {
	octopressPostsDir := "octopress/source/_posts"
	err := filepath.Walk(octopressPostsDir, MigratePost)
	if err != nil {
		log.Fatal(err)
	}
}

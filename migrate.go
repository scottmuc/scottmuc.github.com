package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type OctopressFrontMatter struct {
	Title string
	Date OctopressDateTime
	Categories []string
}

type OctopressDateTime struct {
	time.Time
}

// Dates in Octopress seem to be a mixture of the following date formats
const octopressTimeFormat = "2006-01-02 15:04"
const octopressTimeFormat2 = "2006-01-02 15:04:05 -0700"


// Behold the ugly!
func (odt *OctopressDateTime) UnmarshalYAML(value *yaml.Node) error {
	strValue := value.Value
	parsedTime, err := time.Parse(octopressTimeFormat, strValue)
	if err != nil {
		parsedTime, err = time.Parse(octopressTimeFormat2, strValue)
		if err != nil {
			return err
		} else {
			odt.Time = parsedTime
		}

	} else {
		odt.Time = parsedTime
	}

	return nil
}

func MigratePost(path string, fileInfo os.FileInfo, err error) error {
	if fileInfo.IsDir() {
		return nil
	}

	octopressFilename := filepath.Base(path)
	hugoFilename := strings.Replace(octopressFilename[11:], ".markdown", ".md", -1)
	hugoFilePath := "hugo/content/blog/" + hugoFilename

	octopressFileContents, e := os.ReadFile(path)
	if e != nil {
		log.Fatal(e)
	}

	frontMatter := strings.Split(string(octopressFileContents), "---")[1]

	t := OctopressFrontMatter{}

	e2 := yaml.Unmarshal([]byte(frontMatter), &t)
	if e2 != nil {
		log.Fatalf("error: %v", e2)
	}

	fmt.Printf("%s -> %s\n", path, hugoFilePath)
	fmt.Printf("title: %s\n", t.Title)
	fmt.Printf("date: %v\n", t.Date)
	fmt.Printf("categories: %v\n", t.Categories)

	return nil
}

func main() {
	octopressPostsDir := "octopress/source/_posts"
	err := filepath.Walk(octopressPostsDir, MigratePost)
	if err != nil {
		log.Fatal(err)
	}
}

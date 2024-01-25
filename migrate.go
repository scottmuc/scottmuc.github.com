package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type OctopressFrontMatter struct {
	Title string
	Date OctopressDateTime
	Categories []string
}

type HugoFrontMatter struct {
	Title string
	Date time.Time
}

type OctopressDateTime struct {
	time.Time
}

type ImageLink struct {
	Alignment string
	Width int
	ImgSrc string
	Href string
}


const octopressPostsDir = "octopress/source/_posts"
const hugoPostsDir = "hugo/content/blog"

func parseImageLink(octopressImg string) (*ImageLink, error) {
	pattern := `\{\%\s*img\s+[left|right|center]*\s*([\S]+)\s*\d*\s*(?P<skip>"[^"]*")?\s*\%\}`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(octopressImg)

	if len(matches) < 1 {
		return nil, fmt.Errorf("Could not parse link: %s", octopressImg)
	}

	imageLink := &ImageLink{
		ImgSrc: matches[1],
	}
	return imageLink, nil
}

func (odt *OctopressDateTime) UnmarshalYAML(value *yaml.Node) error {
	// Dates in Octopress seem to be a mixture of the following date formats
	parsedTime, err := time.Parse("2006-01-02 15:04", value.Value)
	if err != nil {
		parsedTime, err = time.Parse("2006-01-02 15:04:05 -0700", value.Value)
		if err != nil {
			return fmt.Errorf("Could not unmarshal date: %s", err)
		}
	}

	odt.Time = parsedTime
	return nil
}

func createHugoFile(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	
	return nil
}


func MigratePost(path string, fileInfo os.FileInfo, err error) error {
	if fileInfo.IsDir() {
		return nil
	}

	octopressFilename := filepath.Base(path)
	hugoFilename := strings.Replace(octopressFilename[11:], ".markdown", ".md", -1)
	hugoFilePath := hugoPostsDir + "/" + hugoFilename

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

	postContent := strings.Split(string(octopressFileContents), "---")[2]

	imgRegex := `\{\%\s*img.*\%\}`
	re := regexp.MustCompile(imgRegex)
	images := re.FindAllString(postContent, -1)

	fmt.Printf("%s -> %s\n", path, hugoFilePath)
	fmt.Printf("title: %s\n", t.Title)
	fmt.Printf("date: %v\n", t.Date)
	fmt.Printf("categories: %v\n", t.Categories)
	fmt.Printf("imgs: \n")
	for _, image := range images {
		imageLink, e := parseImageLink(image)

		newImg := "![test image](https://scottmuc.com"+imageLink.ImgSrc+")"
		postContent = strings.Replace(postContent, image, newImg, -1)

		if e != nil {
			log.Fatalln(e)
		}
		fmt.Printf(" - %v\n", imageLink)
	}

	hugoFrontMatter := HugoFrontMatter{
		Title: t.Title,
		Date:  t.Date.Time,
	}

	hugoYamlFrontMatter, _ := yaml.Marshal(&hugoFrontMatter)
	hugoPost := fmt.Sprintf("---\n%s\n---\n%s\n", hugoYamlFrontMatter, postContent)

	e3 := createHugoFile(hugoFilePath, hugoPost)
	if e3 != nil {
		log.Fatalln("Error writing hugo file:", e3)
	}
	return nil
}

func main() {
	os.RemoveAll(hugoPostsDir)
	os.Mkdir(hugoPostsDir, 0755)

	err := filepath.Walk(octopressPostsDir, MigratePost)
	if err != nil {
		log.Fatal(err)
	}
}

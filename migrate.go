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
	Title      string
	Date       OctopressDateTime
	Categories []string
}

type HugoFrontMatter struct {
	Title string
	Date  time.Time
	Tags  []string
}

type OctopressDateTime struct {
	time.Time
}

const octopressPostsDir = "octopress/source/_posts"
const hugoPostsDir = "hugo/content/blog"

func parseImageSrc(octopressImg string) (string, error) {
	pattern := `\{\%\s*img\s+[left|right|center]*\s*([\S]+)\s*\d*\s*(?P<skip>"[^"]*")?\s*\%\}`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(octopressImg)

	if len(matches) < 1 {
		pattern := `src="([^"]*)"`
		re := regexp.MustCompile(pattern)
		matches = re.FindStringSubmatch(octopressImg)
		if len(matches) < 1 {
			return "", fmt.Errorf("Could not parse link: %s", octopressImg)
		}
	}

	return matches[1], nil
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

func writeFile(filePath, content string) error {
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

func parseImages(content string) []string {
	re := regexp.MustCompile(`\{\%\s*img.*\%\}`)
	imagesOcto := re.FindAllString(content, -1)

	re = regexp.MustCompile(`<img[^>]*>`)
	imagesImg := re.FindAllString(content, -1)

	images := append(imagesOcto, imagesImg...)
	return images
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

	ofm := OctopressFrontMatter{}

	e2 := yaml.Unmarshal([]byte(frontMatter), &ofm)
	if e2 != nil {
		log.Fatalf("error: %v", e2)
	}

	postContent := strings.Split(string(octopressFileContents), "---")[2]

	images := parseImages(postContent)
	for _, image := range images {
		imgRelSrc, e := parseImageSrc(image)
		if e != nil {
			log.Fatalln(e)
		}

		newImg := "![test image](https://scottmuc.com" + imgRelSrc + ")"
		postContent = strings.Replace(postContent, image, newImg, -1)
	}

	hugoFrontMatter := HugoFrontMatter{
		Title: ofm.Title,
		Date:  ofm.Date.Time,
		Tags:  ofm.Categories,
	}

	hugoYamlFrontMatter, _ := yaml.Marshal(&hugoFrontMatter)
	hugoPost := fmt.Sprintf("---\n%s\n---\n%s\n", hugoYamlFrontMatter, postContent)

	e3 := writeFile(hugoFilePath, hugoPost)
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

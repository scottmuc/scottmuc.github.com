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

const octopressPostsDir = "octopress/source/_posts"
const hugoPostsDir = "hugo/content/blog"

type OctopressFrontMatter struct {
	Title      string
	Date       OctopressDateTime
	Categories []string
}

type HugoFrontMatter struct {
	Title   string
	Date    time.Time
	Tags    []string
	Aliases []string
}

type OctopressDateTime struct {
	time.Time
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

func createPostBundle(bundleDir, content string) error {
	os.Mkdir(bundleDir, 0755)
	postPath := bundleDir + "/index.md"

	file, err := os.Create(postPath)
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

func processContent(content string) string {
	newContent := content
	images := parseImages(content)
	for _, image := range images {
		imgRelSrc, e := parseImageSrc(image)
		if e != nil {
			log.Fatalln(e)
		}

		newImg := "![test image](https://scottmuc.com" + imgRelSrc + ")"
		newContent = strings.Replace(newContent, image, newImg, -1)
	}

	youtubePattern := `\{% youtube ([^%]+) %\}`
	re := regexp.MustCompile(youtubePattern)
	newContent = re.ReplaceAllString(newContent, "{{< youtube $1 >}}")
	return newContent
}

func MigratePost(path string, fileInfo os.FileInfo, err error) error {
	if fileInfo.IsDir() {
		return nil
	}

	octopressFilename := filepath.Base(path)
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

	hugoPageBundleName := strings.Replace(octopressFilename[11:], ".markdown", "", -1)

	hugoFrontMatter := HugoFrontMatter{
		Title:   ofm.Title,
		Date:    ofm.Date.Time,
		Tags:    ofm.Categories,
		Aliases: []string{"/" + hugoPageBundleName},
	}

	postContent := strings.Split(string(octopressFileContents), "---")[2]
	processedContent := processContent(postContent)

	hugoYamlFrontMatter, _ := yaml.Marshal(&hugoFrontMatter)
	hugoPost := fmt.Sprintf("---\n%s\n---\n%s\n", hugoYamlFrontMatter, processedContent)

	hugoFilePath := hugoPostsDir + "/" + hugoPageBundleName
	e3 := createPostBundle(hugoFilePath, hugoPost)
	if e3 != nil {
		log.Fatalln("Error writing hugo file:", e3)
	}
	return nil
}

func MigrateImage(path string, fileInfo os.FileInfo, err error) error {
	if fileInfo.IsDir() {
		return nil
	}

	fileName := filepath.Base(path)
	if fileName != "index.md" {
		return nil
	}

	contents, e := os.ReadFile(path)
	if e != nil {
		log.Fatal(e)
	}

	// pattern for image links
	pattern := `\[\!\[test image\]\(https:\/\/scottmuc\.com\/images.+\)\]\(/images(.+)\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(string(contents), -1)

	const octopressImagesDir = "octopress/source/images"
	for _, match := range matches {
		fmt.Println(match[0])
		imagePath := octopressImagesDir + match[1]
		hugoImagePath := filepath.Dir(path) + "/" + filepath.Base(imagePath)
		copyFile(imagePath, hugoImagePath)
	}

	return nil
}

func copyFile(srcPath string, dstPath string) {
	in, err := os.Open(srcPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer in.Close()

	out, err := os.Create(dstPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer out.Close()

	_, err = out.ReadFrom(in)
	if err != nil {
		log.Fatalln(err)
	}
}

func doPostMigration() {
	os.RemoveAll(hugoPostsDir)
	os.Mkdir(hugoPostsDir, 0755)

	err := filepath.Walk(octopressPostsDir, MigratePost)
	if err != nil {
		log.Fatal(err)
	}
}

func doImageMigration() {
	err := filepath.Walk(hugoPostsDir, MigrateImage)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// no longer running post migrations since I've now modified the contents
	// of the migrated posts. I didn't realize I could do the image migration
	// until after the fact... so hence the commenting out of the function.
	//doPostMigration()
	doImageMigration()
}

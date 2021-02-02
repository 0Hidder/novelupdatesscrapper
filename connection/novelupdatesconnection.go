package connection

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	baseChapterURL       = "www.novelupdates.com/extnu/"
	baseChapterURLIndex  = 27
	standardChapterClass = "chp-release"
	rand518ChapterClass  = "chp-release rand518"
)

// Novel is a struct representing the novel we get from the URL from NovelUpdates
type Novel struct {
	Name          string `json:"name"`
	LatestChapter string `json:"latestChapter"`
	ChapterURL    string `json:"chapterURL"`
}

// Extracts the body of the HTML page from an URL
func extractBodyFromURL(URL string) string {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	response, err := client.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Convert the response to string format
	dataInBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)
	return pageContent
}

// GetChapterFromString Function that gets the latest chapter
func GetChapterFromString(URL string) string {
	URLAsString := extractBodyFromURL(URL)
	// Find the index of the opening tag
	chapterStartIndex := strings.Index(URLAsString, "<a title=\"")
	if chapterStartIndex == -1 {
		// generate log here
		os.Exit(0)
	}

	// increase the index. I`m only interested in getting the content of the tag
	// no the tags themselves
	chapterStartIndex += 10

	// Find the index of the closing tag
	chapterEndIndex := strings.Index(URLAsString, "\" class=\"chp-release")
	if chapterEndIndex == -1 {
		//Generate log here
		os.Exit(0)
	}

	chapterTitle := []byte(URLAsString[chapterStartIndex:chapterEndIndex])

	return string(chapterTitle)
}

//GetTitleFromString gets the title of the novel from an URL
func GetTitleFromString(URL string) string {
	URLAsString := extractBodyFromURL(URL)
	titleStartIndex := strings.Index(URLAsString, "<title>")
	if titleStartIndex == -1 {
		// generate log here
		os.Exit(0)
	}

	// The start index of the title is the index of the first
	// character, the < symbol. We don't want to include
	// <title> as part of the final value, so let's offset
	// the index by the number of characers in <title>
	titleStartIndex += 7

	// Find the index of the closing tag
	titleEndIndex := strings.Index(URLAsString, "</title>")
	if titleEndIndex == -1 {
		// generate log here
		os.Exit(0)
	}

	// Copy the substring in to a separate variable so the
	// variables with the full document data can be garbage collected
	pageTitle := []byte(URLAsString[titleStartIndex:titleEndIndex])

	return string(pageTitle)
}

// GetChapterURLFromString Function that gets the URL for the chapter
func GetChapterURLFromString(URL string) string {
	// Find the index of the opening tag
	URLAsString := extractBodyFromURL(URL)
	chapterStartIndex := strings.Index(URLAsString, baseChapterURL)
	if chapterStartIndex == -1 {
		// generate log here
		os.Exit(0)
	}

	// increase the index. I`m only interested in getting the content of the tag
	// no the tags themselves
	chapterStartIndex += baseChapterURLIndex

	// Find the index of the closing tag
	chapterEndIndex := strings.Index(URLAsString, "\" rel")
	if chapterEndIndex == -1 {
		//Generate log here
		os.Exit(0)
	}

	chapterURL := []byte(URLAsString[chapterStartIndex:chapterEndIndex])

	finalURL := "https://" + baseChapterURL + string(chapterURL)

	return finalURL

}

// GetBodyAsJSON Extracts the elements of the HTML Page, puts it on a struct and return the JSON of said struct
func (n Novel) GetBodyAsJSON(URL string) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	response, err := client.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Convert the response to string format
	dataInBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)
	exportAsJSON(pageContent)
}

func exportAsJSON(URLAsString string) []byte {

	var novelStruct *Novel
	novelStruct = new(Novel)

	novelStruct.LatestChapter = GetChapterFromString(URLAsString)
	novelStruct.Name = GetTitleFromString(URLAsString)
	novelStruct.ChapterURL = GetChapterURLFromString(URLAsString)

	j, err := json.Marshal(novelStruct)
	if err != nil {
		// generate log here
	}
	if j != nil {
		// error while parsing
	}

	return j

}

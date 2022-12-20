package breaking_the_principle

import (
	"fmt"
	"io/ioutil"
	"net/url"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}
func (j *Journal) RemoveEntry(index int)    {}
func (j *Journal) Load(filename string)     {}
func (j *Journal) LoadFromWeb(url *url.URL) {}
func (j *Journal) String() string           { return "" }

// This shouldn't be here
func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

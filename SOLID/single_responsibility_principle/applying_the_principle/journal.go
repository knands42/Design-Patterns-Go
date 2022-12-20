package applying_the_principle

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
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
func (j *Journal) RemoveEntry(index int) {}

func (j *Journal) Load(filename string)     {}
func (j *Journal) LoadFromWeb(url *url.URL) {}
func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// Separations of Concerns
type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

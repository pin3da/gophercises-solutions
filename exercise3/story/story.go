package story

import (
	"encoding/json"
	"os"
)

// Chapter of CYOA
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

// Adventure represents all the chapters of CYOA
type Adventure map[string]Chapter

// LoadArcsFromFile loads all the chapters from a json file
func LoadArcsFromFile(filename string) (Adventure, error) {

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	var result Adventure
	d := json.NewDecoder(file)
	if err := d.Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

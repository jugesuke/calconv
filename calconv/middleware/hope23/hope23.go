package hope23

import (
	"bytes"
	_ "embed"
	"encoding/csv"
)

//go:embed hope23.csv
var hope23 []byte

type Links struct {
	LinkList map[string]string
}

func New() *Links {
	r := bytes.NewReader(hope23)
	reader := csv.NewReader(r)
	linkList := make(map[string]string)
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}
		linkList[row[0]] = row[1]
	}
	return &Links{LinkList: linkList}
}

func (l *Links) GetLMSLink(subjectName string) (string, error) {
	return l.LinkList[subjectName], nil
}

func (l *Links) GetLMSName() string {
	return "HOPE"
}

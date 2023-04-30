package stu23

// FUN Students 学務システム(2023) 用

import (
	"bufio"
	"errors"
	"io"
	"strings"

	"github.com/minojirapid/calconv/calconv"
)

type Parser struct {
	l calconv.LMSLinkGetter
}

func New(l calconv.LMSLinkGetter) *Parser {
	return &Parser{l: l}
}

func (p *Parser) ParseTimeTable(r io.Reader) ([]calconv.ClassMetaData, error) {
	var classMetaData []calconv.ClassMetaData

	table := parseTimeTable(r)
	for hour_index, row := range table {
		if hour_index < 1 {
			continue
		}
		for date, v := range row {
			if date < 1 || v == "" {
				continue
			}
			md, err := parseMetaData(p.l, v)
			if err != nil {
				return nil, err
			}
			md.Hour = hour_index
			md.Date = calconv.Date(date)
			classMetaData = append(classMetaData, md)
		}
	}
	return classMetaData, nil

}

func parseTimeTable(r io.Reader) [][]string {
	scanner := bufio.NewScanner(r)

	var table [][]string
	for i := 0; i < 7; i += 1 {
		if !scanner.Scan() {
			break
		}
		line := strings.Split(scanner.Text(), "\t")
		line[0] = line[0][0:1]
		table = append(table, line)
	}
	return table
}
func parseMetaData(l calconv.LMSLinkGetter, input string) (calconv.ClassMetaData, error) {
	rPerlenPlace := strings.Index(input, "(")
	lPerlenPlace := strings.Index(input, ")")
	// title 抽出
	if rPerlenPlace <= 1 {
		err := errors.New("parseMetaData: failed to get subject information: " + input)
		return calconv.ClassMetaData{}, err
	}
	subject := input[0:rPerlenPlace]
	if subject == "" {
		err := errors.New("parseMetaData: failed to get subject information: " + input)
		return calconv.ClassMetaData{}, err
	}

	// 教室 抽出
	room := input[rPerlenPlace+1 : lPerlenPlace]
	if room == "" {
		err := errors.New("parseMetaData: failed to get room information: " + input)
		return calconv.ClassMetaData{}, err
	}

	// LMS Link
	link, err := l.GetLMSLink(subject)
	if err != nil {
		return calconv.ClassMetaData{}, err
	}

	return calconv.ClassMetaData{
		Subject: subject,
		Room:    room,
		LMSLink: link,
	}, nil
}

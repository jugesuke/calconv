package calconv

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Flags struct {
	DaysFilePath   *string
	OutputFilePath *string
}

type Container struct {
	l LMSLinkGetter
	t TimeTableParser
	s SchoolScheduleGetter
	e Exporter
}

func New(l LMSLinkGetter, t TimeTableParser, s SchoolScheduleGetter, e Exporter) *Container {
	return &Container{l: l, t: t, s: s, e: e}
}

func (c *Container) Calconv(f Flags) (int, error) {
	var err error

	// load timetable from stdin
	fmt.Println("Please paste your timetable")

	var classMetaData []ClassMetaData
	{
		// set metadata from "students" system
		classMetaData, err = c.t.ParseTimeTable(os.Stdin)
		if err != nil {
			fmt.Println(err)
			return 1, err
		}
	}
	fmt.Printf("parsed timetable successfully\n\n")

	// convert
	fmt.Println("convert to the calendar events...")
	var events []Event
	{
		// convert to calendar data

		zero := time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)

		for _, class := range classMetaData {
			classDays, err := c.s.GetClassDay(class.Date)
			if err != nil {
				return 1, err
			}

			start, end, err := c.s.GetClassHour(class.Hour)
			if err != nil {
				return 1, err
			}

			fmt.Printf("Subject: %s\nRoom: %s\n%s: %s\n", class.Subject, class.Room, c.l.GetLMSName(), class.LMSLink)
			for t, day := range classDays {
				event := Event{
					Subject:     fmt.Sprintf("%s #%d", class.Subject, t+1),
					StartTime:   day.Add(start.Sub(zero)),
					EndTime:     day.Add(end.Sub(zero)),
					AllDayEvent: false,
					Description: fmt.Sprintf("教室: %s\n%s: %s", class.Room, c.l.GetLMSName(), class.LMSLink),
					Location:    class.Room,
				}
				events = append(events, event)

				fmt.Printf("#%2s %s ~ %s\n", strconv.Itoa(t+1), event.StartTime.Format("2006-01-02 15:04"), event.EndTime.Format("15:04"))
			}
			fmt.Println()
		}
	}
	fmt.Printf("Success!\n\n")

	// export
	fmt.Print("Export...")
	{
		f, err := os.Create(*f.OutputFilePath)
		if err != nil {
			return 1, err
		}
		defer f.Close()
		c.e.Export(f, events)
	}

	fmt.Println("Success!")

	return 0, nil
}

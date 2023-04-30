package google_calendar

import (
	"encoding/csv"
	"io"
	"strconv"

	"github.com/minojirapid/calconv/calconv"
)

type CsvExporter struct{}

func NewCsvExporter() *CsvExporter {
	return &CsvExporter{}
}

func (exp *CsvExporter) Export(w io.Writer, e []calconv.Event) error {
	cw := csv.NewWriter(w)
	defer cw.Flush()

	err := cw.Write([]string{"Subject", "StartDate", "StartTime", "EndDate", "EndTime", "AllDayEvent", "Description", "Location", "Private"})
	if err != nil {
		return err
	}
	for _, v := range e {
		subject := v.Subject
		startDate := v.StartTime.Format("01/02/2006")
		startTime := v.StartTime.Format("03:04 PM")
		endDate := v.EndTime.Format("01/02/2006")
		endTime := v.EndTime.Format("03:04 PM")
		allDayEvent := strconv.FormatBool(v.AllDayEvent)
		description := v.Description
		location := v.Location
		private := strconv.FormatBool(true)
		err := cw.Write([]string{
			subject,
			startDate, startTime,
			endDate, endTime,
			allDayEvent, description, location, private,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

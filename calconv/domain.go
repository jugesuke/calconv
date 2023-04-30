package calconv

import (
	"io"
	"time"
)

type LMSLinkGetter interface {
	GetLMSLink(subjectName string) (string, error)
	GetLMSName() string
}

type TimeTableParser interface {
	ParseTimeTable(r io.Reader) ([]ClassMetaData, error)
}

type SchoolScheduleGetter interface {
	GetClassDay(date Date) ([]time.Time, error)
	GetClassHour(nth int) (time.Time, time.Time, error)
}

type EventGenerator interface {
	Generate(m []ClassMetaData) ([]Event, error)
}

type Exporter interface {
	Export(w io.Writer, e []Event) error
}

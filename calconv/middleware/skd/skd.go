package skd

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/minojirapid/calconv/calconv"
)

type Schedule struct {
	days [7][]time.Time
}

func New(skdFilePath string) (*Schedule, error) {
	f, err := os.Open(skdFilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var days [7][]time.Time
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		line := strings.Split(scanner.Text(), ",")

		day, err := time.Parse("2006/1/2", line[0])
		if err != nil {
			return nil, err
		}

		date, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, err
		}

		days[date] = append(days[date], day)
	}
	return &Schedule{days: days}, nil
}

func (s *Schedule) GetClassDay(date calconv.Date) ([]time.Time, error) {
	return s.days[date], nil
}

func (s *Schedule) GetClassHour(nth int) (time.Time, time.Time, error) {
	start := time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
	end := time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
	switch nth {
	case 1:
		start = time.Date(0, 0, 0, 9, 0, 0, 0, time.Local)
		end = time.Date(0, 0, 0, 10, 30, 0, 0, time.Local)
	case 2:
		start = time.Date(0, 0, 0, 10, 40, 0, 0, time.Local)
		end = time.Date(0, 0, 0, 12, 10, 0, 0, time.Local)
	case 3:
		start = time.Date(0, 0, 0, 13, 10, 0, 0, time.Local)
		end = time.Date(0, 0, 0, 14, 40, 0, 0, time.Local)
	case 4:
		start = time.Date(0, 0, 0, 14, 50, 0, 0, time.Local)
		end = time.Date(0, 0, 0, 16, 20, 0, 0, time.Local)
	case 5:
		start = time.Date(0, 0, 0, 16, 30, 0, 0, time.Local)
		end = time.Date(0, 0, 0, 18, 0, 0, 0, time.Local)
	case 6:
		start = time.Date(0, 0, 0, 18, 10, 0, 0, time.Local)
		end = time.Date(0, 0, 0, 19, 40, 0, 0, time.Local)
	default:
		err := errors.New("writeSchedule: parsing hour " + strconv.Itoa(nth) + ": invalid number.")
		return start, end, err
	}
	return start, end, nil
}

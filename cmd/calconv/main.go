package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/minojirapid/calconv/calconv"
	"github.com/minojirapid/calconv/calconv/middleware/google_calendar"
	"github.com/minojirapid/calconv/calconv/middleware/hope23"
	"github.com/minojirapid/calconv/calconv/middleware/skd"
	"github.com/minojirapid/calconv/calconv/middleware/stu23"
)

var ver string = "23.04 (Beta)"

func main() {
	fmt.Println("This is calconv for FUN " + ver + ".")
	// flags
	var flags calconv.Flags
	flags.DaysFilePath = flag.String("d", "days.skd", "school days list (txt) path")
	flags.OutputFilePath = flag.String("o", "your_schedule.csv", "out put file (csv) path")
	flag.Parse()

	// DI
	l := hope23.New()
	t := stu23.New(l)
	s, err := skd.New(*flags.DaysFilePath)
	if err != nil {
		panic(err)
	}
	e := google_calendar.NewCsvExporter()
	c := calconv.New(l, t, s, e)

	// Execute
	code, err := c.Calconv(flags)
	if err != nil {
		fmt.Println(err)
		os.Exit(code)
	}
}

package main

import (
	"flag"
	"time"
)

func main() {
	var intflagvar int
	var boolflagvar bool
	var stringflagvar string
	var durationflagvar time.Duration

	flag.IntVar(&intflagvar, "intflag", 0, "int flag value")
	flag.BoolVar(&boolflagvar, "boolflag", false, "bool flag value")
	flag.StringVar(&stringflagvar, "stringflag", "default", "string flag value")
	flag.DurationVar(&durationflagvar, "durationflag", 0, "duration flag value")

	_ = flag.Int("intflag", 0, "int flag value")
	_ = flag.Bool("boolflag", false, "bool flag value")
	_ = flag.String("stringflag", "default", "string flag value")
	_ = flag.Duration("durationflag", 0, "duration flag value")

	flag.Parse()
}

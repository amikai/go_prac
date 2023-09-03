package main

import (
	"flag"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFsetBindVar(t *testing.T) {
	args := []string{"-intflag", "12", "-stringflag", "test", "-boolflag", "-durationflag", "10s"}

	var intflag int
	var boolflag bool
	var stringflag string
	var durationflag time.Duration

	fs := flag.NewFlagSet("MyFlagSet", flag.PanicOnError)
	fs.IntVar(&intflag, "intflag", 0, "int flag value")
	fs.BoolVar(&boolflag, "boolflag", false, "bool flag value")
	fs.StringVar(&stringflag, "stringflag", "default", "string flag value")
	fs.DurationVar(&durationflag, "durationflag", 0, "duration flag value")
	fs.Parse(args)

	assert.Equal(t, 12, intflag)
	assert.True(t, boolflag)
	assert.Equal(t, "test", stringflag)
	assert.Equal(t, 10*time.Second, durationflag)
}

func TestFsetAssign(t *testing.T) {
	args := []string{"-intflag", "12", "-stringflag", "test", "-boolflag", "-durationflag", "10s"}

	fs := flag.NewFlagSet("MyFlagSet", flag.PanicOnError)
	intflag := fs.Int("intflag", 0, "int flag value")
	boolflag := fs.Bool("boolflag", false, "bool flag value")
	stringflag := fs.String("stringflag", "default", "string flag value")
	durationflag := fs.Duration("durationflag", 0, "duration flag value")
	fs.Parse(args)

	assert.Equal(t, 12, *intflag)
	assert.True(t, *boolflag)
	assert.Equal(t, "test", *stringflag)
	assert.Equal(t, 10*time.Second, *durationflag)
}

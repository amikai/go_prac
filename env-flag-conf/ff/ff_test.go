package ffex

import (
	"flag"
	"testing"
	"time"

	"github.com/peterbourgon/ff/v4"
	"github.com/peterbourgon/ff/v4/ffyaml"
	"github.com/stretchr/testify/assert"
)

func TestFFlag(t *testing.T) {
	args := []string{"-intflag", "12", "-stringflag", "test", "-boolflag", "-durationflag", "10s"}

	var intflag int
	var boolflag bool
	var stringflag string
	var durationflag time.Duration

	fs := flag.NewFlagSet("my_program", flag.ContinueOnError)
	fs.IntVar(&intflag, "intflag", 0, "int flag value")
	fs.BoolVar(&boolflag, "boolflag", false, "bool flag value")
	fs.StringVar(&stringflag, "stringflag", "default", "string flag value")
	fs.DurationVar(&durationflag, "durationflag", 0, "duration flag value")
	err := ff.Parse(fs, args)

	assert.NoError(t, err)
	assert.Equal(t, 12, intflag)
	assert.True(t, boolflag)
	assert.Equal(t, "test", stringflag)
	assert.Equal(t, 10*time.Second, durationflag)
}

func TestFFlagWithEnvPrefix(t *testing.T) {
	t.Setenv("MY_PROGRAM_INTFLAG", "12")
	t.Setenv("MY_PROGRAM_BOOLFLAG", "true")
	t.Setenv("MY_PROGRAM_STRINGFLAG", "test")
	t.Setenv("MY_PROGRAM_DURATIONFLAG", "10s")

	fs := flag.NewFlagSet("my_program", flag.ContinueOnError)
	intflag := fs.Int("intflag", 0, "int flag value")
	boolflag := fs.Bool("boolflag", false, "bool flag value")
	stringflag := fs.String("stringflag", "default", "string flag value")
	durationflag := fs.Duration("durationflag", 0, "duration flag value")
	err := ff.Parse(fs, []string{}, ff.WithEnvVarPrefix("MY_PROGRAM"))

	assert.NoError(t, err)
	assert.Equal(t, 12, *intflag)
	assert.True(t, *boolflag)
	assert.Equal(t, "test", *stringflag)
	assert.Equal(t, 10*time.Second, *durationflag)
}

func TestFFlagWithEnvConfig(t *testing.T) {
	fs := flag.NewFlagSet("my_program", flag.ContinueOnError)
	intflag := fs.Int("intflag", 0, "int flag value")
	boolflag := fs.Bool("boolflag", false, "bool flag value")
	stringflag := fs.String("stringflag", "default", "string flag value")
	durationflag := fs.Duration("durationflag", 0, "duration flag value")
	_ = fs.String("config", "", "config file")
	err := ff.Parse(fs, []string{"-config", "testdata/config.yaml"}, ff.WithConfigFileFlag("config"), ff.WithConfigFileParser(ffyaml.Parse))

	assert.NoError(t, err)
	assert.Equal(t, 12, *intflag)
	assert.True(t, *boolflag)
	assert.Equal(t, "test", *stringflag)
	assert.Equal(t, 10*time.Second, *durationflag)
}

func TestFFLognOption(t *testing.T) {
	// TODO: implement the long option
}

func TestFFMixOptionEnvConf(t *testing.T) {
	// TODO: implement mixing options, env variable and config file
}

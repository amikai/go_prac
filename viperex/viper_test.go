package viperex

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoadEnvFile(t *testing.T) {
	// See https://github.com/spf13/viper#why-viper
	// Precedence of explicit call to Set higher than config
	config := []byte(`
var:
  key: "VALUE_IN_CONF"
`)

	os.Setenv("MYAPP_VAR_KEY", "VALUE")

	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.SetEnvPrefix("MYAPP")
	viper.AutomaticEnv()
	err := viper.ReadConfig(bytes.NewBuffer(config))
	assert.NoError(t, err)

	expect := "VALUE"
	got := viper.GetString("var.key")
	assert.Equal(t, expect, got)
}

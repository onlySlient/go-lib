package conf

import (
	"os"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGetEnv(t *testing.T) {
	if err := os.Setenv("PG_HOST", "localhost"); err != nil {
		t.Fatal(err)
	}

	val := GetEnv("PG_HOST")
	t.Log(val)
	assert.Equal(t, val, "localhost")
}

func TestGetString(t *testing.T) {
	if err := os.Setenv("PG_HOST", ""); err != nil {
		t.Fatal(err)
	}

	val := GetString("PG_HOST")
	t.Log(val)
	assert.Equal(t, val, "127.0.0.1")
}

func TestEnvAndConfig(t *testing.T) {
	// assert env priority over config
	if err := os.Setenv("PG_HOST", "localhost"); err != nil {
		t.Fatal(err)
	}

	val := PG_HOST.Value()
	t.Log(val)
	assert.Equal(t, val, "localhost")
}

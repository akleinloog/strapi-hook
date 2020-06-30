package config

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDefaultPortIs8080(t *testing.T) {
	config := New()
	assert.Equal(t, 8080, config.Port())
}

func TestPortCanBeSetWithViper(t *testing.T) {
	viper.Set("port", 500)
	config := New()
	assert.Equal(t, 500, config.Port())
	viper.Set("port", nil)
}

func TestPortCanBeSetWithEnvironmentVariable(t *testing.T) {
	os.Setenv("STRAPI_HOOK_PORT", "10080")
	Initialize()
	config := New()
	assert.Equal(t, 10080, config.Port())
	os.Setenv("STRAPI_HOOK_PORT", "")
}

func TestDefaultTargetIsLocalHost10080(t *testing.T) {
	config := New()
	assert.Equal(t, "http://localhost:10080/api", config.Target())
}

func TestTargetCanBeSetWithViper(t *testing.T) {
	viper.Set("target", "http://sometest:8080/api")
	config := New()
	assert.Equal(t, "http://sometest:8080/api", config.Target())
	viper.Set("target", nil)
}

func TestTargetCanBeSetWithEnvironmentVariable(t *testing.T) {
	os.Setenv("STRAPI_HOOK_TARGET", "http://env-server:8080/api")
	Initialize()
	config := New()
	assert.Equal(t, "http://env-server:8080/api", config.Target())
	os.Setenv("STRAPI_HOOK_TARGET", "")
}

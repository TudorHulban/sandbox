package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const _path = "configs/config.json"
const _emailSender = "golang@libre.eu"

func TestConfigurationWrite(t *testing.T) {
	configEmail := ConfigurationEmail{
		EmailSender: _emailSender,
	}

	config := Configuration{
		ConfigurationEmail: configEmail,
	}

	f, errCr := os.Create(_path)
	require.NoError(t, errCr)

	defer f.Close()

	config.WriteTo(f)
}

func TestNewConfigurationFromFS(t *testing.T) {
	config, errCr := NewConfigurationFromFS("configs/config.json")
	require.NoError(t, errCr)
	require.NotZero(t, config)

	require.Equal(t, _emailSender, config.EmailSender)
}

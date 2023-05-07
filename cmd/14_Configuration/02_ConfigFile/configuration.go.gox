package main

import (
	"encoding/json"
	"io"
)

type ConfigurationEmail struct {
	EmailSender string `valid:"required,email"`
}

type Configuration struct {
	// monorepo configuration

	// microservice (services) common configuration

	// microservice configuration
	ConfigurationEmail
}

func NewConfigurationFromFS(fsPath string) (*Configuration, error) {
	content, errRead := _fs.ReadFile(fsPath)
	if errRead != nil {
		return nil, errRead
	}

	var res Configuration

	_, errWr := res.Write(content)
	if errWr != nil {
		return nil, errWr
	}

	return &res, nil
}

func (cfg *Configuration) WriteTo(w io.Writer) (int, error) {
	data, errMa := json.MarshalIndent(cfg, "", " ")
	if errMa != nil {
		return 0, errMa
	}

	return w.Write(data)
}

func (cfg *Configuration) Write(p []byte) (int, error) {
	var config Configuration

	if errUn := json.Unmarshal(p, &config); errUn != nil {
		return 0, errUn
	}

	*cfg = config

	return len(p), nil
}

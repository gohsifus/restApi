package configs

import (
	"encoding/json"
	"os"
	"restApi/errs"
)

// ServerConfig конфигурация сервера
type ServerConfig struct {
	Host      string `json:"host"`
	Port      string `json:"port"`
	PathToLog string `json:"path_to_log"`
}

// NewConfig ...
func NewConfig() *ServerConfig {
	return &ServerConfig{}
}

// LoadConfigs загрузит логи из json файла
func (s *ServerConfig) LoadConfigs(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return errs.Wrap(err)
	}

	err = json.Unmarshal(data, s)
	if err != nil {
		errs.Wrap(err)
	}

	return nil
}

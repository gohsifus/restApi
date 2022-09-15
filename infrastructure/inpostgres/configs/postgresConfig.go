package configs

import (
	"encoding/json"
	"fmt"
	"os"
	"restApi/errs"
)

// PostgresConfig конфигурации для подключения к postgres
type PostgresConfig struct {
	Host     string `json:"Host"`
	Port     string `json:"Port"`
	User     string `json:"User"`
	Password string `json:"Password"`
	DBName   string `json:"db_name"`
}

// NewConfig ...
func NewConfig() *PostgresConfig {
	return &PostgresConfig{}
}

// LoadConfigs загрузит логи из json файла
func (p *PostgresConfig) LoadConfigs(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return errs.Wrap(err)
	}

	err = json.Unmarshal(data, p)
	if err != nil {
		errs.Wrap(err)
	}

	return nil
}

// GetConnectionString вернет строку подключения к бд, на основании данных в конфиге
func (p *PostgresConfig) GetConnectionString() string {
	return fmt.Sprintf(
		"Host=%s Port=%s User=%s Password=%s dbname=%s sslmode=disable",
		p.Host,
		p.Port,
		p.User,
		p.Password,
		p.DBName,
	)
}

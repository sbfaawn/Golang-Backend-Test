package config

import "fmt"

const (
	DBUser     = "postgres"
	DBPassword = "dika1dika"
	DBName     = "dummy"
	DBHost     = "localhost"
	DBPort     = "5432"
	DBType     = "postgres"
)

func GetDBType() string {
	return DBType
}

func GetDBConfigurationString() string {
	database := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		DBType, DBUser, DBPassword, DBHost, DBPort, DBName)

	return database
}

package databases

import (
	"database/sql"

	"com.github/mrthoabby/m-authentication/globalConfig"
	"com.github/mrthoabby/m-authentication/types"
)

type iDatabaseConnectionFactory interface {
	CreateConnection() (IDatabaseConnectionRepository, error)
}

type IDatabaseConnectionRepository interface {
	Query(query string) (*sql.Rows, error)
	ValidAuthentication(credentials types.Credentials) bool
}

type ConcreteDatabaseFactorystruct struct {
}

func NewConcreteDatabaseFactory(databaseType, connectionString string) iDatabaseConnectionFactory {
	switch databaseType {
	case globalConfig.CONNECTION_DATABASE_TYPE_SQL:
		return NewMycrosoftSQLFactory(connectionString)
	default:
		return nil
	}
}

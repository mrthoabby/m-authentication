package databases

import (
	"database/sql"

	"com.github/mrthoabby/m-authentication/globalConfig"
	"com.github/mrthoabby/m-authentication/types"
	"com.github/mrthoabby/m-authentication/types/basic"
)

type iDatabaseConnectionFactory interface {
	CreateConnection(tableMapper *types.TableMapper) (IDatabaseConnectionRepository, error)
}

type IDatabaseConnectionRepository interface {
	Query(query string, args ...any) (*sql.Rows, error)
	GetPasswordHash(credentials types.Credentials) (string, error)
}

type ConcreteDatabaseFactorystruct struct {
}

func NewConcreteDatabaseFactory(connectionData basic.Connection) iDatabaseConnectionFactory {
	switch connectionData.Type {
	case globalConfig.CONNECTION_DATABASE_TYPE_SQL:
		return NewMycrosoftSQLFactory(connectionData)
	default:
		return nil
	}
}

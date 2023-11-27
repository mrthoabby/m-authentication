package databases

import (
	"database/sql"

	"com.github/mrthoabby/m-authentication/types"
	"com.github/mrthoabby/m-authentication/util"
)

type microsoftSQL struct {
	connection *sql.DB
}

func newMicrosoftSQL(connectionString *sql.DB) *microsoftSQL {
	return &microsoftSQL{connection: connectionString}
}

func (s *microsoftSQL) Disconnect() {
	if s.connection != nil {
		s.connection.Close()
	}
}

func (s *microsoftSQL) Query(query string) (*sql.Rows, error) {
	rows, err := s.connection.Query(query)
	defer rows.Close()
	if err != nil {
		util.LoggerHandler().Error("Error querying database", "error", err.Error())
		return nil, err
	}
	return rows, types.NewCustomError("Error querying database")
}

// Is not functional is just an example, because i nned to use encrypted passwords
func (s *microsoftSQL) ValidAuthentication(credentials types.Credentials) bool {
	rows, errorQuerying := s.Query("SELECT * FROM users WHERE user = '" + credentials.User + "' AND password = '" + credentials.Password + "'")
	if errorQuerying != nil {
		return false
	}
	return rows.Next()
}

type mycrosoftSQLFactory struct {
	connectionString string
}

func NewMycrosoftSQLFactory(connectionString string) *mycrosoftSQLFactory {
	return &mycrosoftSQLFactory{connectionString: connectionString}
}

func (m *mycrosoftSQLFactory) CreateConnection() (IDatabaseConnectionRepository, error) {
	database, errorConnecting := sql.Open("sql", m.connectionString)
	if errorConnecting != nil {
		util.LoggerHandler().Error("Error connecting to database", "error", errorConnecting.Error())
		return nil, errorConnecting
	}
	if pingError := database.Ping(); pingError != nil {
		util.LoggerHandler().Error("Error pinging database", "error", pingError.Error())
		return nil, pingError
	}
	return newMicrosoftSQL(database), nil
}

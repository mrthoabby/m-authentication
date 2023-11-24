package databases

import (
	"database/sql"
	logHandler "log/slog"
)

type MSSql struct {
	database *sql.DB
}

func (s *MSSql) Connect(connectionString string) (bool, error) {
	database, err := sql.Open("mysql", connectionString)
	defer database.Close()
	if err != nil {
		logHandler.Error("Error connecting to database", "error", err.Error())
		return false, err
	}
	if pingError := database.Ping(); pingError != nil {
		logHandler.Error("Error pinging database", "error", pingError.Error())
		return false, pingError
	}
	s.database = database
	return true, nil
}

func (s *MSSql) Disconnect() {
	if s.database != nil {
		s.database.Close()
	}
}

func (s *MSSql) Query(query string) *sql.Rows {
	rows, err := s.database.Query(query)
	if err != nil {
		logHandler.Error("Error querying database", "error", err.Error())
	}
	return rows
}

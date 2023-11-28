package databases

import (
	"database/sql"
	"net/url"
	"strconv"

	"com.github/mrthoabby/m-authentication/globalConfig"
	"com.github/mrthoabby/m-authentication/services"
	"com.github/mrthoabby/m-authentication/types"
	"com.github/mrthoabby/m-authentication/types/basic"
	"com.github/mrthoabby/m-authentication/util"
	_ "github.com/microsoft/go-mssqldb"
)

type microsoftSQL struct {
	connection   *sql.DB
	tablesMapper *types.TableMapper
}

func newMicrosoftSQL(connection *sql.DB, tableMapper *types.TableMapper) (*microsoftSQL, error) {

	tableValidatorService, erroGettingService := services.NewTableValidatorService(connection, tableMapper)
	if erroGettingService != nil {
		util.LoggerHandler().Error("Error creating table validator service", "error", erroGettingService.Error())
		return nil, erroGettingService
	}

	if !(tableValidatorService.IsAnValidColumn(tableMapper.AuthTable, tableMapper.UserColumn) &&
		tableValidatorService.IsAnValidColumn(tableMapper.AuthTable, tableMapper.PasswordColumn)) {
		util.LoggerHandler().Error("Error validating table", "error", "Invalid table: "+tableMapper.AuthTable)
		return nil, types.NewCustomError("Invalid table: " + tableMapper.AuthTable)
	}

	return &microsoftSQL{
		connection:   connection,
		tablesMapper: tableMapper,
	}, nil
}

func (s *microsoftSQL) Disconnect() {
	if s.connection != nil {
		s.connection.Close()
	}
}

func (s *microsoftSQL) Query(query string, args ...any) (*sql.Rows, error) {
	rows, err := s.connection.Query(query)
	if err != nil {
		util.LoggerHandler().Error("Error querying database", "error", err.Error())
		return nil, err
	}
	defer rows.Close()
	return rows, types.NewCustomError("Error querying database")
}

func (s *microsoftSQL) GetPasswordHash(credentials types.Credentials) (string, error) {

	query := `SELECT ` + s.tablesMapper.PasswordColumn + ` FROM ` + s.tablesMapper.AuthTable + ` WHERE LOWER(` + s.tablesMapper.UserColumn + `) = LOWER(@username)`
	var passwordHash string
	errorQuerying := s.connection.QueryRow(query, sql.Named("username", credentials.User)).Scan(&passwordHash)
	if errorQuerying != nil {
		util.LoggerHandler().Error("Error querying database [GetPasswordHash]", "error", errorQuerying.Error())
		return "", errorQuerying
	}
	if passwordHash == "" {
		return "", types.NewCustomError("There is no user with that username")
	}
	return passwordHash, nil
}

type mycrosoftSQLFactory struct {
	connectionData basic.Connection
}

func NewMycrosoftSQLFactory(connectionData basic.Connection) *mycrosoftSQLFactory {
	return &mycrosoftSQLFactory{
		connectionData: connectionData,
	}
}

func (m *mycrosoftSQLFactory) CreateConnection(tableMapper *types.TableMapper) (IDatabaseConnectionRepository, error) {
	if m.connectionData.Type != globalConfig.CONNECTION_DATABASE_TYPE_SQL {
		util.LoggerHandler().Error("Error creating database connection", "error", "Invalid database type")
		return nil, types.NewCustomError("Invalid database type")
	}

	var hostToUse = m.connectionData.Host
	if m.connectionData.Port != 0 {
		hostToUse += ":" + strconv.Itoa(m.connectionData.Port)
	}

	query := url.Values{}
	query.Add("database", m.connectionData.Database)
	connectionUrl := &url.URL{
		Scheme:   m.connectionData.Type,
		User:     url.UserPassword(m.connectionData.User, m.connectionData.Password),
		Host:     hostToUse,
		RawQuery: query.Encode(),
	}

	database, errorConnecting := sql.Open(m.connectionData.Type, connectionUrl.String())
	if errorConnecting != nil {
		util.LoggerHandler().Error("Error connecting to database", "error", errorConnecting.Error())
		return nil, errorConnecting
	}
	if pingError := database.Ping(); pingError != nil {
		util.LoggerHandler().Error("Error pinging database", "error", pingError.Error())
		return nil, pingError
	}
	sqlDatabase, errorGettingDb := newMicrosoftSQL(database, tableMapper)
	if errorGettingDb != nil {
		util.LoggerHandler().Error("Error getting database", "error", errorGettingDb.Error())
		return nil, errorGettingDb
	}

	return sqlDatabase, nil
}

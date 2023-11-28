package services

import (
	"database/sql"

	"com.github/mrthoabby/m-authentication/types"
	"com.github/mrthoabby/m-authentication/util"
)

type TableValidatorService struct {
	databaseConnection *sql.DB
	validTables        []string
}

func NewTableValidatorService(databaseConnection *sql.DB, tableMapper *types.TableMapper) (*TableValidatorService, error) {
	tables := append(tableMapper.DataSourceTables, tableMapper.AuthTable)
	for _, table := range tables {
		if !isAnValidTable(table, databaseConnection) {
			util.LoggerHandler().Error("Error validating table", "error", "Invalid table: "+table)
			return nil, types.NewCustomError("Invalid table: " + table)
		}
	}
	return &TableValidatorService{
		databaseConnection: databaseConnection,
		validTables:        tables,
	}, nil
}

func isAnValidTable(tableName string, databaseConnection *sql.DB) bool {
	query := "SELECT CASE WHEN EXISTS (SELECT 1 FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = @tableName) THEN 1 ELSE 0 END"
	var exist bool
	err := databaseConnection.QueryRow(query, sql.Named("tableName", tableName)).Scan(&exist)
	if err != nil {
		util.LoggerHandler().Error("Error querying database", "error", err.Error())
		return false
	}
	return exist
}

func (t *TableValidatorService) IsAnValidColumn(tableName string, columnName string) bool {
	if !util.ContainsString(t.validTables, tableName) {
		util.LoggerHandler().Error("Error validating table", "error", "Invalid table: "+tableName)
		return false
	}

	query := `SELECT CASE WHEN COUNT(COLUMN_NAME) > 0 THEN 1 ELSE 0 END AS Result FROM information_schema.columns WHERE TABLE_NAME = @tableName AND COLUMN_NAME = @columnName`
	var exists bool
	err := t.databaseConnection.QueryRow(query, sql.Named("tableName", tableName), sql.Named("columnName", columnName)).Scan(&exists)
	if err != nil {
		util.LoggerHandler().Error("Error querying database", "error", err.Error())
		return false
	}
	if !exists {
		util.LoggerHandler().Error("Error validating column", "error", "Invalid column: "+columnName)
	}

	return exists
}

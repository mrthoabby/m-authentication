package basic

import (
	"reflect"

	"com.github/mrthoabby/m-authentication/globalConfig"
	"com.github/mrthoabby/m-authentication/util"
	"github.com/go-playground/validator/v10"
)

type Config struct {
	Connection []connection `xml:"connection" validate:"required,allConnectionsCanBeAnUniqueId"`
	Auth       auth         `xml:"auth" validate:"required"`
}

func (c *Config) GetType() reflect.Type {
	return reflect.TypeOf(*c)
}

func allConnectionsCanBeAnUniqueId(fl validator.FieldLevel) bool {
	config := fl.Parent().Interface().(Config)
	ids := make(map[string]bool)
	for _, connection := range config.Connection {
		if _, exists := ids[connection.Id]; exists {
			return false
		}
		ids[connection.Id] = true
	}
	return true
}

type connection struct {
	Id       string `xml:"id,attr" validate:"required"`
	Type     string `xml:"type,attr" validate:"required,validTypeConnection"`
	Host     string `xml:"host,attr" validate:"required"`
	Port     int    `xml:"port,attr"`
	Database string `xml:"database,attr" validate:"required"`
	User     string `xml:"user,attr" validate:"required"`
	Password string `xml:"password,attr" validate:"required"`
}

func isAnValidTypeConnection(fl validator.FieldLevel) bool {
	validTypes := globalConfig.CurrentConnectionDatabaseType
	for _, validType := range validTypes {
		if fl.Field().String() == validType {
			return true
		}
	}
	return false
}

type auth struct {
	UseRoles   bool   `xml:"useRoles,attr"`
	RouterName string `xml:"routerName,attr" validate:"required"`
	Table      table  `xml:"table" validate:"required"`
	Roles      roles  `xml:"roles" validate:"rolesRequiredAuth"`
}

func isRolesRequiredAuth(fl validator.FieldLevel) bool {
	auth := fl.Parent().Interface().(auth)
	return auth.UseRoles
}

type table struct {
	Name     string   `xml:"name,attr" validate:"required"`
	User     user     `xml:"user" validate:"required,userColumnIsDiferentPasswordTable"`
	Password password `xml:"password" validate:"required"`
}

func userColumnIsDiferentPasswordTable(fl validator.FieldLevel) bool {
	table := fl.Parent().Interface().(table)
	return table.User.Column != table.Password.Column
}

type user struct {
	Column string `xml:"column,attr" validate:"required"`
}

type password struct {
	Column  string  `xml:"column,attr" validate:"required"`
	Encrypt encrypt `xml:"encrypt" validate:"required"`
}

type encrypt struct {
	Algorithm string `xml:"algorithm,attr" validate:"required,validAlgorithm"`
	Source    string `xml:"source,attr" validate:"required,validSource"`
	Key       string `xml:"key,attr" validate:"required,KeyRequired"`
}

func isAnValidAlgorithm(fl validator.FieldLevel) bool {
	validAlgorithms := globalConfig.CurrentAlgorithms
	for _, validAlgorithm := range validAlgorithms {
		if fl.Field().String() == validAlgorithm {
			return true
		}
	}
	return false
}

func isAnValidSource(fl validator.FieldLevel) bool {
	validSources := globalConfig.CurrentSources
	for _, validSource := range validSources {
		if fl.Field().String() == validSource {
			return true
		}
	}
	return false
}

func isAnRequiredKey(fl validator.FieldLevel) bool {
	encrypt := fl.Parent().Interface().(encrypt)
	if encrypt.Source == globalConfig.SOURCE_ENCRYPTION_LOCAL {
		return encrypt.Key != ""
	}
	return true
}

type roles struct {
	Global global `xml:"global"`
	Role   []role `xml:"role" validate:"required"`
}

type role struct {
	Name   string `xml:"name,attr" validate:"required"`
	Claims claims `xml:"claims" validate:"required"`
}

type global struct {
	Claims claims `xml:"claims" validate:"required"`
}

type claims struct {
	DataSource []dataSource `xml:"DataSource" validate:"required"`
}

type dataSource struct {
	Type         string     `xml:"type,attr" validate:"required,validTypeDataSource"`
	ConnectionId string     `xml:"connectionId,attr" validate:"IdConnectionRequired"`
	AsideTable   asideTable `xml:"aside_table" validate:"AsideTableRequired"`
	Claim        []claim    `xml:"claim" validate:"ClaimRequired"`
}

func isAnValidTypeDataSource(fl validator.FieldLevel) bool {
	validTypes := globalConfig.CurrentDataSourcesType
	for _, validType := range validTypes {
		if fl.Field().String() == validType {
			return true
		}
	}
	return false
}

func isIdConnectionRequiredIfTypeIsConnectionDataSource(fl validator.FieldLevel) bool {
	dataSource := fl.Parent().Interface().(dataSource)
	return dataSource.Type == globalConfig.DATASOURCE_TYPE_CONNECTION
}

func isAsideTableRequiredIfTypeIsConnectionDataSource(fl validator.FieldLevel) bool {
	dataSource := fl.Parent().Interface().(dataSource)
	return dataSource.Type == globalConfig.DATASOURCE_TYPE_CONNECTION
}

func isClaimRequiredIfTypeIsNotConnectionDataSource(fl validator.FieldLevel) bool {
	dataSource := fl.Parent().Interface().(dataSource)
	return dataSource.Type != globalConfig.DATASOURCE_TYPE_CONNECTION
}

type asideTable struct {
	Name          string  `xml:"name,attr" validate:"required"`
	Column        string  `xml:"column,attr" validate:"required,isColumnDifferentSessionColumn"`
	SessionColumn string  `xml:"sessionColumn,attr" validate:"required"`
	Claim         []claim `xml:"claim" validate:"required"`
}

func isColumnDifferentSessionColumnAsideTable(fl validator.FieldLevel) bool {
	asideTable := fl.Parent().Interface().(asideTable)
	return asideTable.Column != asideTable.SessionColumn
}

type claim struct {
	Name   string `xml:"name,attr" validate:"required"`
	Format string `xml:"format,attr" validate:"required,validFormatClaim"`
}

func isAnValidFormatClaim(fl validator.FieldLevel) bool {
	validFormats := globalConfig.CurrentClaimValidFormats
	for _, validFormat := range validFormats {
		if fl.Field().String() == validFormat {
			return true
		}
	}
	return false
}

func init() {
	util.GetValidator().RegisterValidation("IdConnectionRequired", isIdConnectionRequiredIfTypeIsConnectionDataSource)
	util.GetValidator().RegisterValidation("isColumnDifferentSessionColumn", isColumnDifferentSessionColumnAsideTable)
	util.GetValidator().RegisterValidation("AsideTableRequired", isAsideTableRequiredIfTypeIsConnectionDataSource)
	util.GetValidator().RegisterValidation("userColumnIsDiferentPasswordTable", userColumnIsDiferentPasswordTable)
	util.GetValidator().RegisterValidation("ClaimRequired", isClaimRequiredIfTypeIsNotConnectionDataSource)
	util.GetValidator().RegisterValidation("allConnectionsCanBeAnUniqueId", allConnectionsCanBeAnUniqueId)
	util.GetValidator().RegisterValidation("validTypeDataSource", isAnValidTypeDataSource)
	util.GetValidator().RegisterValidation("validTypeConnection", isAnValidTypeConnection)
	util.GetValidator().RegisterValidation("validFormatClaim", isAnValidFormatClaim)
	util.GetValidator().RegisterValidation("rolesRequiredAuth", isRolesRequiredAuth)
	util.GetValidator().RegisterValidation("validAlgorithm", isAnValidAlgorithm)
	util.GetValidator().RegisterValidation("validSource", isAnValidSource)
	util.GetValidator().RegisterValidation("KeyRequired", isAnRequiredKey)
}

package basicAuth

import (
	"com.github/mrthoabby/m-authentication/global"
	"com.github/mrthoabby/m-authentication/utils"
	"github.com/go-playground/validator/v10"
)

type Config struct {
	Connection []Connection `xml:"connection" validate:"required,allConnectionsCanBeAnUniqueId"`
	Auth       Auth         `xml:"auth" validate:"required"`
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

type Connection struct {
	Id       string `xml:"id,attr" validate:"required"`
	Type     string `xml:"type,attr" validate:"required,validTypeConnection"`
	Host     string `xml:"host,attr" validate:"required"`
	Port     int    `xml:"port,attr"`
	Database string `xml:"database,attr" validate:"required"`
	User     string `xml:"user,attr" validate:"required"`
	Password string `xml:"password,attr" validate:"required"`
}

func isAnValidTypeConnection(fl validator.FieldLevel) bool {
	validTypes := []string{"sql"}
	for _, validType := range validTypes {
		if fl.Field().String() == validType {
			return true
		}
	}
	return false
}

type Auth struct {
	UseRoles bool  `xml:"useRoles,attr"`
	Table    Table `xml:"table" validate:"required"`
	Roles    Roles `xml:"roles" validate:"rolesRequiredAuth"`
}

func isRolesRequiredAuth(fl validator.FieldLevel) bool {
	auth := fl.Parent().Interface().(Auth)
	return auth.UseRoles
}

type Table struct {
	Name     string   `xml:"name,attr" validate:"required"`
	User     User     `xml:"user" validate:"required,userColumnIsDiferentPasswordTable"`
	Password Password `xml:"password" validate:"required"`
}

func userColumnIsDiferentPasswordTable(fl validator.FieldLevel) bool {
	table := fl.Parent().Interface().(Table)
	return table.User.Colum != table.Password.Colum
}

type User struct {
	Colum string `xml:"colum,attr" validate:"required"`
}

type Password struct {
	Colum   string  `xml:"colum,attr" validate:"required"`
	Encrypt Encrypt `xml:"encrypt" validate:"required"`
}

type Encrypt struct {
	Algorithm string `xml:"algorithm,attr" validate:"required,validAlgorithm"`
	Source    string `xml:"source,attr" validate:"required,validSource"`
	Key       string `xml:"key,attr" validate:"required,KeyRequired"`
}

func isAnValidAlgorithm(fl validator.FieldLevel) bool {
	validAlgorithms := global.CurrentAlgorithms
	for _, validAlgorithm := range validAlgorithms {
		if fl.Field().String() == validAlgorithm {
			return true
		}
	}
	return false
}

func isAnValidSource(fl validator.FieldLevel) bool {
	validSources := global.CurrentSources
	for _, validSource := range validSources {
		if fl.Field().String() == validSource {
			return true
		}
	}
	return false
}

func isAnRequiredKey(fl validator.FieldLevel) bool {
	encrypt := fl.Parent().Interface().(Encrypt)
	return encrypt.Source == global.SOURCE_ENCRYPTION_LOCAL
}

type Roles struct {
	Global Global `xml:"global"`
	Role   []Role `xml:"role" validate:"required"`
}

type Role struct {
	Name   string `xml:"name,attr" validate:"required"`
	Claims Claims `xml:"claims" validate:"required"`
}

type Global struct {
	Claims Claims `xml:"claims" validate:"required"`
}

type Claims struct {
	DataSource []DataSource `xml:"DataSource" validate:"required"`
}

type DataSource struct {
	Type         string     `xml:"type,attr" validate:"required,validTypeDataSource"`
	ConnectionId string     `xml:"connectionId,attr" validate:"IdConnectionRequired"`
	AsideTable   AsideTable `xml:"aside_table" validate:"AsideTableRequired"`
	Claim        []Claim    `xml:"claim" validate:"ClaimRequired"`
}

func isAnValidTypeDataSource(fl validator.FieldLevel) bool {
	validTypes := global.CurrentDataSourcesType
	for _, validType := range validTypes {
		if fl.Field().String() == validType {
			return true
		}
	}
	return false
}

func isIdConnectionRequiredIfTypeIsConnectionDataSource(fl validator.FieldLevel) bool {
	dataSource := fl.Parent().Interface().(DataSource)
	return dataSource.Type == global.DATASOURCE_TYPE_CONNECTION
}

func isAsideTableRequiredIfTypeIsConnectionDataSource(fl validator.FieldLevel) bool {
	dataSource := fl.Parent().Interface().(DataSource)
	return dataSource.Type == global.DATASOURCE_TYPE_CONNECTION
}

func isClaimRequiredIfTypeIsNotConnectionDataSource(fl validator.FieldLevel) bool {
	dataSource := fl.Parent().Interface().(DataSource)
	return dataSource.Type != global.DATASOURCE_TYPE_CONNECTION
}

type AsideTable struct {
	Name          string  `xml:"name,attr" validate:"required"`
	Column        string  `xml:"column,attr" validate:"required,isColumnDifferentSessionColumn"`
	SessionColumn string  `xml:"sessionColumn,attr" validate:"required"`
	Claim         []Claim `xml:"claim" validate:"required"`
}

func isColumnDifferentSessionColumnAsideTable(fl validator.FieldLevel) bool {
	asideTable := fl.Parent().Interface().(AsideTable)
	return asideTable.Column != asideTable.SessionColumn
}

type Claim struct {
	Name   string `xml:"name,attr" validate:"required"`
	Format string `xml:"format,attr" validate:"required,validFormatClaim"`
}

func isAnValidFormatClaim(fl validator.FieldLevel) bool {
	validFormats := global.CurrentClaimValidFormats
	for _, validFormat := range validFormats {
		if fl.Field().String() == validFormat {
			return true
		}
	}
	return false
}

func init() {
	validator := utils.GetValidator()
	global.Once.Do(func() {
		validator.RegisterValidation("IdConnectionRequired", isIdConnectionRequiredIfTypeIsConnectionDataSource)
		validator.RegisterValidation("isColumnDifferentSessionColumn", isColumnDifferentSessionColumnAsideTable)
		validator.RegisterValidation("AsideTableRequired", isAsideTableRequiredIfTypeIsConnectionDataSource)
		validator.RegisterValidation("userColumnIsDiferentPasswordTable", userColumnIsDiferentPasswordTable)
		validator.RegisterValidation("ClaimRequired", isClaimRequiredIfTypeIsNotConnectionDataSource)
		validator.RegisterValidation("allConnectionsCanBeAnUniqueId", allConnectionsCanBeAnUniqueId)
		validator.RegisterValidation("validTypeDataSource", isAnValidTypeDataSource)
		validator.RegisterValidation("validTypeConnection", isAnValidTypeConnection)
		validator.RegisterValidation("validFormatClaim", isAnValidFormatClaim)
		validator.RegisterValidation("rolesRequiredAuth", isRolesRequiredAuth)
		validator.RegisterValidation("validAlgorithm", isAnValidAlgorithm)
		validator.RegisterValidation("validSource", isAnValidSource)
		validator.RegisterValidation("KeyRequired", isAnRequiredKey)
	})
}

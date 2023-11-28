package types

type TableMapper struct {
	AuthTable        string
	UserColumn       string
	PasswordColumn   string
	DataSourceTables []string
	DataSourcColumns map[string][]string
}

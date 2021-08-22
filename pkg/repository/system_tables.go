package repository

const (
	db_type_postgres = "postgres"
)

type SystemTablesConf struct {
	sequences string
}

func initSystemTablesConf(db_type string) *SystemTablesConf {
	switch db_type {
	case db_type_postgres:
		{
			return &SystemTablesConf{
				sequences: "pg_sequences",
			}
		}
	default:
		{
			return &SystemTablesConf{}
		}
	}

}

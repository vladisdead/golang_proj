package cfg

func New() (*CFG, error) {
	cfg := CFG{
		Storage: Storage{
			Connstring:     "user=api-user password=password host=localhost port=5432 database=kursovaya sslmode=disable",
			MigrationsPath: "storage/migrations/",
		},
	}
	return &cfg, nil
}

db_login:
	psql ${DATABASE_URL}

db_create_migration:
	migrate create -ext sql -dir migrations -seq $(name)

db_migrate:
	migrate -database ${DATABASE_URL} -path migrations up

db_force:
	migrate -database ${DATABASE_URL} -path migrations force 1

db_verbose_down:
	migrate -database ${DATABASE_URL} -path migrations -verbose down 1

db_verbose_up:
	migrate -database ${DATABASE_URL} -path migrations -verbose up

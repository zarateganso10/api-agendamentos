migrate-up:
	migrate -path adapter/postgres/migration/ -database "postgres://postgres:postgres@192.168.68.107:5433/agendamentos?sslmode=disable" -verbose up

migrate-down:
	migrate -path adapter/postgres/migration/ -database "postgres://zarate:b9AF89YHHQkqv6Nj0SGrcrXVQkvMcGTs@dpg-cmikdtn109ks739mnrfg-a.oregon-postgres.render.com/agendamentos" -verbose down

run:
	go run cmd/api/main.go
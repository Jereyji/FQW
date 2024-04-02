migrate create -ext sql -dir ./schema -seq init 

docker run --name=db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
## What is this?
A simple Todo API, testing chi router

## How to run
### 1. Up database
~~~~
*Necessary docker
Inside the folder build/postgres/ run:

docker build -t db_todos .
docker-compose up -d
docker exec -it db_todos psql -p 5431 -U postgres -a -f docker-entrypoint-initdb.d/init.sql

Or, on linux, you can run:
sudo chmod +x db-up.sh && ./db-up.sh
~~~~
### 2. Run the go code 
~~~~
Inside the root folder, run:
go run main.go
~~~~

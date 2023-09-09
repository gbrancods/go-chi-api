docker build -t db_todos .
docker-compose up -d
docker exec -it db_todos psql -p 5431 -U postgres -a -f docker-entrypoint-initdb.d/init.sql

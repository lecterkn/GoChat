# 作業記録

## sql-migrateのインストール

```shell
go get -v github.com/rubenv/sql-migrate/...
go install github.com/rubenv/sql-migrate/...@latest
```

## dockerでのDBへの接続
```shell
docker exec -it my_postgres psql -U postgres
```

## DB作成
```sql
CREATE DATABASE mydb;
```
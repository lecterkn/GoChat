# 作業記録

## sql-migrateのインストール

```shell
go get -v github.com/rubenv/sql-migrate/...
go install github.com/rubenv/sql-migrate/...@latest
```

## sql-migrate新規作成

```shell
sql-migrate new <action>
```

例：ユーザーテーブルを作成
```shell
sql-migrate new create_user
```

## PostgresSQL
```shell
go get github.com/lib/pq
```

## dockerでのDBへの接続
```shell
docker exec -it my_postgres psql -U postgres
```

## DB作成
```sql
CREATE DATABASE mydb;
```

## UUID
```shell
go get github.com/google/uuid
```
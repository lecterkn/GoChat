# UniversalTalk Backend

## DB起動
```shell
docker compose -f build/package/docker-compose.yml up
```

## サーバー起動
```shell
go run cmd/gochat/main.go
```

## マイグレーション
```shell
cd buikd/package
sql-migrate up
```
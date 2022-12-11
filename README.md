# sqlx-xo-example

Go + sqlx + xo のお試しリポジトリ

```
go version
1.19 darwin/amd64

mysql version
8.0

docker version
20.10.17

docker compose version
v2.7.0
```

## Usage

```
docker compose build --no-cache
// develop environment
docker compose up

// production environment
docker build --target production -t go-app:latest . -f docker/backend/Dockerfile

curl http://localhost:8000/
```

## description

- 開発環境ではホットリロード対応
- 本番環境用イメージは軽量イメージを使用

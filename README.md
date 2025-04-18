# Goの学習用リポジトリ

## 環境構築
- ローカルにGo、golang-migrateのインストール
- リポジトリをクローン
- リポジトリのディレクトリに移動
- `cp .env.example .env` で環境変数ファイルを作成
- `docker compose up -d --build` でコンテナを立ち上げ
- ブラウザで`http://localhost:8080/ping`にアクセス

## DBマイグレーションのupコマンド
- `migrate -database "mysql://root:password@tcp(localhost:3306)/emo_tracking" -path db/migrations up`
- `migrate -database "mysql://root:password@tcp(localhost:3306)/emo_tracking" -path db/migrations down`

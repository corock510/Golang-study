FROM golang:1.20-alpine AS build

WORKDIR /app/api

# 依存関係ファイルを先にコピー
COPY api/go.mod api/go.sum ./

# 依存関係をインストール
RUN go mod download

# アプリケーションのソースコードをコピー
COPY api/ ./

# アプリケーションをビルド
RUN go build -o emo_tracking .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=build /app/api/emo_tracking /usr/local/bin/emo_tracking

EXPOSE 8080

CMD ["emo_tracking"]
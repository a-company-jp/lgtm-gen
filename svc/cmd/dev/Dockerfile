FROM ghcr.io/a-company-jp/imagemagick:latest

WORKDIR /app

COPY ./ ./

COPY .env .env

RUN CGO_ENABLED=1 GOOS=linux go build -o main ./svc/cmd/dev/main.go

# ポートの公開
EXPOSE 8080

# アプリケーションの実行
CMD ["./main"]

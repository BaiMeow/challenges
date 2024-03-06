FROM golang:1.21.3 as builder

COPY src /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o app .

FROM chromedp/headless-shell:latest

COPY --from=builder /app/app /app/app

WORKDIR /app

ENTRYPOINT ["./app"]

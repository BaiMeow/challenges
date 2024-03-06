FROM golang:1.16 as builder

COPY src /src

WORKDIR /src

RUN go build -o app

FROM ubuntu:devel

RUN apt-get update \
    && apt-get install -y unzip \
    && apt-get clean

COPY --from=builder /src/app /app/app

WORKDIR /app

CMD ./app

EXPOSE 8080
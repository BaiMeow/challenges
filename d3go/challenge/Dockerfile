FROM golang:1.19.0 as builder

COPY backend /src/

RUN cd /src && go build -o /app/d3go

FROM ubuntu:devel

ENV Flag=d3ctf{flag_test}

RUN groupadd -r d3ctf && useradd -r -g d3ctf d3ctf && mkdir /app && chown d3ctf:d3ctf /app

USER d3ctf

COPY --from=builder --chown=d3ctf:d3ctf /app/d3go /app/d3go

COPY --chown=d3ctf:d3ctf  files/config.yaml /app/config.yaml

EXPOSE 8080

WORKDIR "/app"

CMD ["./d3go"]
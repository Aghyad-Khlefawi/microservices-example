FROM alpine:latest AS base
EXPOSE 8080
EXPOSE 5001
RUN apt-get update && apt-get install -y curl

FROM golang:latest AS build

ENV GOOS=linux
ENV CGO_ENABLED=0
WORKDIR /src
COPY src/identity/go.mod .
RUN go mod download
COPY src/identity .
RUN go build cmd/main.go

FROM base AS final

COPY --from=build /src/main /app/main


HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD curl -f http://localhost:8080/hc || exit 1 

ENTRYPOINT [ "/app/main" ]

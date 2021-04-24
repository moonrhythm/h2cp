FROM golang:1.16.3

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
RUN mkdir -p /workspace
WORKDIR /workspace
ADD go.mod go.sum ./
RUN go mod download
ADD . .
RUN go build -o .build/h2cp -ldflags "-w -s" .

FROM gcr.io/distroless/static

WORKDIR /app

COPY --from=0 /workspace/.build/* ./
ENTRYPOINT ["/app/h2cp"]

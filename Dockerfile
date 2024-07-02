FROM golang:latest

WORKDIR /app
COPY go.mod .
# COPY go.sum .
RUN go mod downloadjjj
COPY . .

CMD ["go", "run", "main.go"]
FROM golang:1.18
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY view ./view
COPY util ./util
COPY queries ./queries
COPY model ./model
RUN go build .
CMD ["./db-proj"]
FROM golang:1.14.15
WORKDIR /app
RUN mkdir -p /app/logs
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bitspawn-api .
EXPOSE 8080
VOLUME /app/logs
CMD ["./bitspawn-api"]

FROM golang:1.21.1-alpine as dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]

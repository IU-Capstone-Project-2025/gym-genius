FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum assets/exercises.json assets/users.json assets/workouts.json  ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o admin-backend ./app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/admin-backend /app/admin-backend

EXPOSE 3000

CMD ["./admin-backend"]

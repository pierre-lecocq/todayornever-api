# Build

FROM golang:1.22 AS build-stage

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/todayornever-api .

# Run

FROM scratch AS run-stage

WORKDIR /app

COPY --from=build-stage /app/todayornever-api /app/todayornever-api

EXPOSE 8080

ENTRYPOINT ["/app/todayornever-api"]

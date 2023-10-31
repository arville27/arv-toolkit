FROM golang:1.21-alpine AS builder

WORKDIR /build

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o arv-toolkit


FROM alpine:3.18 AS runner

RUN addgroup toolkit && \
    adduser -S -G toolkit toolkit

RUN mkdir -p /app/data && \
    mkdir -p /app/config && \
    chown -R toolkit:toolkit /app

WORKDIR /app

COPY --from=builder /build/arv-toolkit arv-toolkit

EXPOSE 8080

USER toolkit:toolkit

CMD ["/app/arv-toolkit"]
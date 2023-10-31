FROM golang:1.21-alpine AS base



FROM base AS builder

WORKDIR /build

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o arv-toolkit



FROM base AS runner

RUN addgroup toolkit && \
    adduser -S -G toolkit toolkit

RUN mkdir -p /app/data && \
    mkdir -p /app/config && \
    chown -R toolkit:toolkit /app

WORKDIR /app
COPY --from=builder --chown=toolkit:toolkit /build/arv-toolkit arv-toolkit

EXPOSE 8080

USER toolkit:toolkit

CMD ["/app/arv-toolkit"]
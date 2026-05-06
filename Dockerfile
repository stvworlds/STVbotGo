# syntax=docker/dockerfile:1

ARG GO_VERSION=1.26.0

FROM golang:${GO_VERSION} AS build

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/server ./cmd/STVbotGo

FROM alpine:latest AS final

RUN apk --no-cache add \
    ca-certificates \
    tzdata \
    ffmpeg \
    && update-ca-certificates

ARG UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser

USER appuser

COPY --from=build /bin/server /bin/server
COPY --from=build /src/configs /configs

EXPOSE 1234

ENTRYPOINT ["/bin/server"]

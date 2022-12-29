# syntax=docker/dockerfile:1
# Step 1 - Build
FROM golang:1.19-alpine AS build

ARG TRASH_SIGN
LABEL sign=$TRASH_SIGN

WORKDIR /github.com/achillescres/saina-api/

COPY . .
RUN go mod download
RUN go mod tidy

RUN GOOS=linux go build -o saina-api cmd/main.go
RUN ls

# Step 2 - prepare
FROM alpine

WORKDIR /app

COPY /.aws ./.aws

COPY /.aws /.aws

COPY /.aws ~/.aws

COPY /external ./external
COPY .env .
COPY --from=build /github.com/achillescres/saina-api/saina-api .

EXPOSE 7771

ENV PROJECT_ABS_PATH "/app"
ENV AWS_CONFIG_FILE "/.aws/config"
ENV AWS_SHARED_CREDENTIALS_FILE="/.aws/credentials"

RUN cat ./.aws/credentials
RUN cat /.aws/credentials

CMD ["./saina-api"]

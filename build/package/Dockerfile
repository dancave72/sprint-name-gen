FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
COPY go.mod /Users/dancave/projects/sprint-name-generator/cmd/name-generator/go.mod
COPY go.sum /Users/dancave/projects/sprint-name-generator/cmd/name-generator/go.sum
RUN apk add git
WORKDIR /build
RUN go mod download
RUN go build -o main github.com/dancave/sprint-name-gen/cmd/name-generator
FROM alpine
EXPOSE 80
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]

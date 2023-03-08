FROM golang:1.19-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /blog-build

EXPOSE 8081

CMD [ "/blog-build" ]
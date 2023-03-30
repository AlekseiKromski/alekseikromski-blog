FROM golang:1.19.4-alpine
WORKDIR /app
ENV GO111MODULE=on BLOG_ADDRESS=":3000" DB_USERNAME="" DB_PASSWORD="" DB_HOSTNAME="" DB_DATABASE="" DB_PORT=""
COPY go.mod ./
RUN go mod download
COPY . ./
RUN go build -o /blog-build

EXPOSE 3000

CMD [ "/blog-build" ]
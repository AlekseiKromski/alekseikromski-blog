FROM node:16.15-alpine
WORKDIR /app
COPY ./front-end ./
RUN npm install
RUN npm install @tiptap/extension-highlight
RUN npm install @tiptap/extension-image
RUN npm run build

FROM golang:1.19.4-alpine
WORKDIR /app
ENV GO111MODULE=on BLOG_ADDRESS=":3001" DB_USERNAME="postgres" DB_PASSWORD="postgres" DB_HOSTNAME="192.168.0.101" DB_DATABASE="blog" DB_PORT="5432"
COPY go.mod ./
RUN go mod download
ADD . .
COPY --from=0 /app/build ./front-end/build
RUN go build -o /blog-build

EXPOSE 3001

CMD [ "/blog-build" ]
FROM golang:1.18

WORKDIR usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

ARG PORT=3000
ENV PORT=$PORT


RUN go build -o /usr/local/bin/app 

EXPOSE $PORT

CMD ["app"]
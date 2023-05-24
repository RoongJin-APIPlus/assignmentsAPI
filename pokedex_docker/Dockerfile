FROM golang

ENV POSTGRES_PASSWORD=test

WORKDIR /app

COPY . ./

RUN go mod download

CMD ["go", "run", "./server.go"]
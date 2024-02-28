FROM golang:latest

WORKDIR /api

COPY . .

RUN make tidy

RUN make build

EXPOSE 5555

CMD make run


FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
RUN go mod download

COPY . .

#build
RUN go build -o sense-docker .

WORKDIR /dist

RUN mkdir tmp

VOLUME [ "/dist/tmp" ]

RUN cp /build/sense-docker .

CMD [ "/dist/sense-docker", "/dist/tmp" ]



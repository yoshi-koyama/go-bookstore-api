FROM golang:1.25
ENV TZ="Asia/Tokyo"
WORKDIR /go/src/app
COPY . .
RUN go install github.com/air-verse/air@v1.61.1

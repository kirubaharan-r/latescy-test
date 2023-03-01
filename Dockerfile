FROM golang:latest
RUN apt update|apt install -y curl 
WORKDIR /test
COPY . /test
#RUN go mod tidy
RUN go get -d -v
RUN CGO_ENABLED=0 go build -a -installsuffix cgo main.go
RUN rm main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
RUN mkdir /Published
COPY --from=0 /test /Published
WORKDIR /Published
EXPOSE 8084
ENV GIN_MODE release
CMD ["./main"]
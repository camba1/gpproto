FROM golang AS Dev
WORKDIR /go/src/gpProto
RUN go get github.com/githubnemo/CompileDaemon
COPY ./price ./price
RUN go get -d  -v ./...
RUN go build -o priceServer ./price/server/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o priceServerAlp ./price/server/
EXPOSE 50051
CMD ["./priceServer"]


FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /goProto
COPY --from=Dev /go/src/gpProto/priceServerAlp priceServerAlp
RUN ls
CMD ["./priceServerAlp"]
<code style="font-size: 10.2375px; font-family: Consolas, Menlo, Monaco, "Courier New", monospace; color: rgba(51, 51, 51, 0.85); background-color: rgb(246, 247, 248); border-radius: 2px; display: block; line-height: 1.5; position: relative; top: 0px; outline: 0px !important;"> 
FROM golang:alpine AS build-env
RUN mkdir /go/src/app && apk update && apk add git
ADD main.go /go/src/app/
WORKDIR /go/src/app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo 
-ldflags '-extldflags "-static"' -o app .
FROM scratch
WORKDIR /app
COPY --from=build-env /go/src/app/app .
ENTRYPOINT [ "./app" ]
</code>
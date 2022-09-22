FROM registry.cn-hangzhou.aliyuncs.com/liy36/golang:1.18.3 as builder
WORKDIR /app
ADD . /app
RUN go mod tidy
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags='-s -w' -o app .


FROM registry.cn-hangzhou.aliyuncs.com/liy36/alpine:3.15
WORKDIR /apps
COPY --from=builder /app/app /apps/app
EXPOSE 8080
CMD /apps/app
# 一阶构建 缓存mod编译二进制文件
FROM harbor.vchangyi.com/library/golang:1.1 AS compile

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /go/compile

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd ./cmd/
COPY internal ./internal/

RUN go build -ldflags="-s -w" -o http  cmd/http/main.go
RUN go build -ldflags="-s -w" -o job  cmd/job/main.go
RUN go build -ldflags="-s -w" -o binlog  cmd/binlog/main.go

# 二阶构建 压缩可执行文件、瘦身
FROM gruebel/upx AS upx
WORKDIR /upx
COPY --from=compile /go/compile/http .
COPY --from=compile /go/compile/job .
COPY --from=compile /go/compile/binlog .
RUN upx http job binlog

# 三阶构建 复制上层可执行文件、配置文件、静态文件
FROM harbor.vchangyi.com/library/golang:1.0

### 环境
ENV env='test'

ENV server='http'

COPY deploy/build/supervisor/application.conf /etc/supervisor/conf.d/

WORKDIR /app

COPY cmd ./cmd/
COPY internal ./internal/
COPY storage ./storage/
COPY db ./db/
COPY source ./source/
COPY storage ./storage/
COPY app.*.yaml ./

COPY --from=upx /upx/http .
COPY --from=upx /upx/job .
COPY --from=upx /upx/binlog .

COPY deploy/build/docker-entrypoint.sh /opt/

RUN chmod +x /opt/docker-entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["/opt/docker-entrypoint.sh"]

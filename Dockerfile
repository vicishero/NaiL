# syntax=docker/dockerfile:1

# build go mod cache
FROM bitbus/paopao-ce-backend-builder:latest AS gomodcache
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /paopao-ce
COPY go.mod .
COPY go.sum .
RUN go mod download

# build backend
FROM gomodcache AS backend
WORKDIR /paopao-ce/server
COPY server/ .
WORKDIR /paopao-ce
ENV GOPROXY=https://goproxy.cn,direct
RUN make -C server build TAGS='slim go_json'

FROM bitbus/paopao-ce-backend-runner:latest
ENV TZ=Asia/Shanghai
WORKDIR /app/paopao-ce
COPY --from=backend /paopao-ce/release/paopao .
COPY --from=backend /paopao-ce/server/config.yaml.sample config.yaml
VOLUME ["/app/paopao-ce/custom"]
EXPOSE 8008
HEALTHCHECK --interval=5s --timeout=3s  --retries=3  CMD ps -ef | grep paopao || exit 1
ENTRYPOINT ["/app/paopao-ce/paopao"]
CMD ["serve"]

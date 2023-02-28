FROM ttbb/base:go AS build
COPY . /opt/compile
WORKDIR /opt/compile/cmd/config
RUN go build -o config_gen .


FROM ttbb/bookkeeper:nake

COPY docker-build /opt/bookkeeper/mate

COPY --from=build /opt/compile/cmd/config/config_gen /opt/bookkeeper/mate/config_gen

COPY config/bk_server_original.conf /opt/bookkeeper/conf/bk_server_original.conf
COPY config/standalone_original.conf /opt/bookkeeper/conf/standalone_original.conf
COPY lib/* /opt/bookkeeper/lib

WORKDIR /opt/bookkeeper

CMD ["/usr/bin/dumb-init", "bash", "-vx", "/opt/bookkeeper/mate/scripts/start.sh"]

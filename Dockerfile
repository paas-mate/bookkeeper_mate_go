FROM ttbb/base:go AS build
COPY . /opt/sh/compile
WORKDIR /opt/sh/compile/pkg
RUN go build -o bookkeeper_mate .
WORKDIR /opt/sh/compile/cmd/config
RUN go build -o config_gen .


FROM ttbb/bookkeeper:nake

COPY docker-build /opt/sh/bookkeeper/mate

COPY --from=build /opt/sh/compile/pkg/bookkeeper_mate /opt/sh/bookkeeper/mate/bookkeeper_mate
COPY --from=build /opt/sh/compile/cmd/config/config_gen /opt/sh/bookkeeper/mate/config_gen

COPY config/bk_server_original.conf /opt/sh/bookkeeper/conf/bk_server_original.conf
COPY config/standalone_original.conf /opt/sh/bookkeeper/conf/standalone_original.conf

# due to Unrecognized VM option 'AggressiveOpts'
COPY config/common_mod.sh /opt/sh/bookkeeper/bin/common.sh

WORKDIR /opt/sh/bookkeeper

CMD ["/usr/bin/dumb-init", "bash", "-vx", "/opt/sh/bookkeeper/mate/scripts/start.sh"]

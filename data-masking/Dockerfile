FROM acr.aishu.cn/public/ubuntu:22.04.20230316

#RUN cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN apt-get update && \
    apt-get install -y curl

COPY cc /usr/local/bin/af/

COPY cmd/server/config /usr/local/bin/af/cmd/server/config/
COPY cmd/server/docs /usr/local/bin/af/cmd/server/docs/
COPY infrastructure/repository/db/gen/migration /usr/local/bin/af/infrastructure/repository/db/gen/migration/
WORKDIR /usr/local/bin/af/
RUN chmod 757 -R  /usr/local/bin/af/  \
   && mkdir -p /opt/log/af  \
   && chmod 757 -R /opt/log/af

ENTRYPOINT ["/usr/local/bin/af/cc","serve","-c/usr/local/bin/af/cmd/server/config/config.yaml"]

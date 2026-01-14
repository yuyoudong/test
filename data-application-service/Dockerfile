FROM acr.aishu.cn/public/ubuntu:22.04.20230329

RUN #cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY data-application-service /usr/local/bin/af/

COPY cmd/server/config /usr/local/bin/af/cmd/server/config/
COPY cmd/server/docs /usr/local/bin/af/cmd/server/docs/
COPY cmd/server/fonts /usr/local/bin/af/cmd/server/fonts/
WORKDIR /usr/local/bin/af/
RUN chmod 757 -R  /usr/local/bin/af/  \
    && mkdir -p /opt/log/af  \
    && chmod 757 -R /opt/log/af

ENTRYPOINT ["/usr/local/bin/af/data-application-service","--confPath","/usr/local/bin/af/cmd/server/config/"]

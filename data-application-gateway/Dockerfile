FROM acr.aishu.cn/public/ubuntu:22.04.20230329

RUN #cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY data-application-gateway /usr/local/bin/af/

COPY cmd/server/config /usr/local/bin/af/cmd/server/config/
RUN chmod 757 -R  /usr/local/bin/af/  \
    && mkdir -p /opt/log/af  \
    && chmod 757 -R /opt/log/af

ENTRYPOINT ["/usr/local/bin/af/data-application-gateway","--confPath","/usr/local/bin/af/cmd/server/config/"]

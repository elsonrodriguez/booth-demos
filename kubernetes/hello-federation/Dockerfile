FROM scratch
MAINTAINER Elson Rodriguez <elson.rodriguez@gmail.com>

ADD hello-federation /

ENV LISTEN_ADDRESS 0.0.0.0
ENV LISTEN_PORT 80

ENV POD_NODENAME pod_nodename
ENV POD_NAME pod_name

ENTRYPOINT ["/hello-federation"]

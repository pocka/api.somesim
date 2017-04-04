FROM scratch
MAINTAINER pocka <pockawoooh@gmail.com>


ADD dist /work

ADD docs /work/docs

WORKDIR /work

ENTRYPOINT ["/work/somesim", "--tls-port", "443", "--tls-host", "0.0.0.0"]

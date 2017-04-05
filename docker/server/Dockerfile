FROM scratch
MAINTAINER pocka <pockawoooh@gmail.com>


ADD dist /work

ADD docs /work/docs

WORKDIR /work

ENTRYPOINT ["/work/somesim", "--port", "80", "--host", "0.0.0.0"]

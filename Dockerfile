FROM alpine:3.4
MAINTAINER Wilson Zhang <topagentwilson@gmail.com>
Add multican /usr/bin/multican
EXPOSE 8100
WORKDIR /usr/bin
ENTRYPOINT ["multican"]


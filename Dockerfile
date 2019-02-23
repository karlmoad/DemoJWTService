FROM alpine:3.7
MAINTAINER Karl Moad <https:/github.com/karlmoad>
RUN apk add --no-cache bash
ADD dist /service
EXPOSE 30200
WORKDIR /service
ENTRYPOINT /service/DemoJWTService
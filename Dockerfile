FROM --platform=linux/arm64 debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

ADD my-blog /usr/bin/my-blog

CMD ["my-blog"]
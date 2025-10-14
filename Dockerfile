FROM --platform=linux/amd64 debian:stable-slim



RUN apt-get update && apt-get install -y ca-certificates

RUN ls -l

ADD my-blog /usr/bin/my-blog

CMD ["my-blog"]

FROM golang:1.23.1
WORKDIR /root/
RUN apt-get update -y && apt-get upgrade -y
COPY .  /root/
CMD ["go","run","."]

FROM telegraf:1.32.0
RUN apt-get update -y && apt-get upgrade -y
RUN apt-get install vim -y

COPY ./telegraf.conf /etc/telegraf/

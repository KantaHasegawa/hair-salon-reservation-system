FROM --platform=linux/x86_64 mysql
RUN mkdir -p /var/lib/mysql-files
RUN apt update
RUN apt install -y tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
ADD ./my.cnf /etc/mysql/my.cnf

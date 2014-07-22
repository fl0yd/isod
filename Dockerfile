FROM google/debian:wheezy
RUN apt-get install -y genisoimage
ADD isod /usr/local/bin/isod
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/isod"]

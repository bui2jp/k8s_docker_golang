FROM ubuntu

ADD ./cktk_post2slack.sh /opt/cktk_post2slack.sh
RUN chmod +x /opt/cktk_post2slack.sh
RUN apt-get update
RUN apt-get install -y curl 

WORKDIR /opt/

# EXPOSE 8080

CMD ["./cktk_post2slack.sh"]

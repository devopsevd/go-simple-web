FROM golang:1.7.3-alpine
ADD ca-certificates.crt /etc/ssl/certs/
ADD microservice /
RUN mkdir /assets
ADD assets /assets
EXPOSE 3000
CMD ["/microservice"]

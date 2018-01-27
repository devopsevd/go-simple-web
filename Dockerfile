FROM golang:1.7.3-alpine
ADD ca-certificates.crt /etc/ssl/certs/
ADD microservice /
EXPOSE 3000
CMD ["/microservice"]

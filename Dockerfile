FROM golang:1.7.3-alpine
ADD ca-certificates.crt /etc/ssl/certs/
RUN mkdir /js
COPY assets/js/app.jsx /js/app.jsx
ADD microservice /
EXPOSE 3000
CMD ["/microservice"]

# go-simple-web

Read Me:

command for build. #CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o microservice .

Add ca-certificates.crt to current folder: #cp /etc/ssl/certs/ca-certificates.crt .

Docker build #docker build -t go_microservice .

Run container #docker run -it --rm -p 8080:8080 go_microservice

Using scratch :

Build using scratch #docker build -t go_microsvc_scratch -f Dockerfile.scratch .
Run container #docker run -it --rm -p 8080:8080 go_microsvc_scratch

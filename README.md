# go-simple-web

Read Me:
1) command for build.
#CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o microservice .

2) Add ca-certificates.crt to current folder: 
#cp /etc/ssl/certs/ca-certificates.crt .

3) Docker build
#docker build -t go_microservice .

4) Run container
#docker run -it --rm -p 3000:3000 --network=postgres_default go_microsvc_scratch

Using scratch :

3) Build using scratch
#docker build -t go_microsvc_scratch -f Dockerfile.scratch .

4) Run container
#docker run -it --rm -p 3000:3000 --network=postgres_default go_microsvc_scratch

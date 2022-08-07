# TO-DO:
# Create docker image with the following commands:
FROM golang:1.18-alpine
WORKDIR $(pwd)/src/bin
ENV NFT_DB_HOST=db
EXPOSE 8080
COPY . .  
RUN go build -o /usr/bin/nft main.go server.go

CMD ["nft"]



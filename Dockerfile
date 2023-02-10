# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.17-alpine

WORKDIR /app

# Download necessary Go modules
#COPY go.mod ./
#COPY go.sum ./

COPY . ./
RUN go mod download
#COPY handlers ./
#COPY data ./

RUN go build -o /product-api

CMD [ "/product-api" ]
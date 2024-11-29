FROM golang:latest

# Install Node.js and npm
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - \
    && apt-get install -y nodejs

RUN npm install -g npm@10.9.1

WORKDIR /app
COPY ../ ./

#React Frontend

WORKDIR /app/website

RUN npm -v

RUN npm install
RUN npm run build


#Go Backend
WORKDIR /app

RUN go mod download

RUN go build -o server/main ./server

EXPOSE 8080

CMD ["./server/main"]

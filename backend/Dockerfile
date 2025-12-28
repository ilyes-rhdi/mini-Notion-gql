# Specifies a parent image
FROM golang:1.21
 
# Creates an app directory to hold your app’s source code
WORKDIR /app
 
# Copies everything from your root directory into /app
COPY . .
 


# Installs Go dependencies
RUN go mod tidy
 


# Builds your app
RUN go build -o server ./cmd/server

EXPOSE 8080

CMD ["./server"]

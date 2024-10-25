# Use the Golang Alpine image as the base
FROM golang:1.23-alpine



# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code and Air configuration file into the container
COPY go.mod ./
RUN go mod download

# Install git, bash, and Delve for debugging
RUN apk add --no-cache git bash && \
    go install github.com/go-delve/delve/cmd/dlv@latest
# Install CA certificates
RUN apk --no-cache add ca-certificates

RUN go install github.com/air-verse/air@latest

# Expose app and Delve debugger ports
EXPOSE 8080 2345

# CMD [ "ls","/app/bin" ]


# Start Air with Delve in headless mode
CMD ["air","-c",".air.toml"]
# CMD [ "dlv" , "debug" ,"--headless" ,"--listen=:2345" ,"--api-version=2" ,"--accept-multiclient", "./"]

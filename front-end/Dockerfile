# We specify the base image we need for our
# go application
FROM public.ecr.aws/docker/library/golang:1.21 as builder
# We create an /app directory within our
# image that will hold our application source
# files
RUN mkdir /app
# We copy everything in the root directory
# into our /app directory
ADD . /app
# We specify that we now wish to execute
# any further commands inside our /app
# directory
WORKDIR /app
# Avoid the GoProxy
ENV GOPROXY=direct
# we run go build to compile the binary
# executable of our Go program
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o service ./

# To make our images smaller, use alpine and copy in the service binary.
FROM alpine:latest
# Install certs
RUN apk --no-cache add ca-certificates
# Copy the binary from the builder image
COPY --from=builder /app ./
# Make the binary executable
RUN chmod +x ./service

# Start the service
ENTRYPOINT ["./service"]
# The service runs on port 80, 8080, and 8081
EXPOSE 80 8082
EXPOSE 8081/udp

# Copy index.html for verifying HTML traffic
FROM public.ecr.aws/nginx/nginx:1.19
COPY index.html /usr/share/nginx/html

# Start from the latest golang base image
FROM golang:1.15
ARG APP=bidder

# Setting ENv
ENV APP ${APP}

# Add Maintainer Info
LABEL maintainer="Rishabh Mishra <rishabhmishra131@gmail.com>"
 
# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .

WORKDIR /app/${APP}

# Build the Go app
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags static_all -a -installsuffix cgo -ldflags '-extldflags "-static"' -o './bin/'${APP} cmd/main/main.go
 
# Expose port 8080 to the outside world
#EXPOSE 8080


# Command to run the executable 
CMD ["bash", "-c", "./bin/${APP}"]
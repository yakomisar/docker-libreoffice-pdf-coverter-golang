# Start from the latest Golang base image based on Debian
FROM golang:1.16-buster

# Add Maintainer Info
LABEL maintainer="example@example.com"

# Disable prompts on apt-get install
ENV DEBIAN_FRONTEND noninteractive

# Set the Current Working Directory inside the container
WORKDIR /app

# Install latest stable LibreOffice
RUN apt-get update -qq \
    && apt-get install -y -q libreoffice \
    && apt-get remove -q -y libreoffice-gnome

# Cleanup after apt-get commands
RUN apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* /var/cache/apt/archives/*.deb /var/cache/apt/*cache.bin

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 6000 to the outside world
EXPOSE 6000

# Command to run the executable
CMD ["./main"]

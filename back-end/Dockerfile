FROM golang:1.14-alpine

# Set the Current Working Directory inside the container
WORKDIR /app/back-end

COPY . .

# This container exposes port 8081
EXPOSE 8081

# RUN main.go
CMD ["go","run","main.go"]

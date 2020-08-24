FROM golang:1.14-alpine

# Set the Current Working Directory inside the container
WORKDIR /app/front-end

COPY . .
RUN apk --no-cache add curl

# This container exposes port 8080
EXPOSE 8080

# RUN main.go
CMD ["go","run","main.go"]
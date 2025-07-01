FROM golang:1.17.2-bullseye as builder
LABEL maintainer="System-Y-Team"
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ARG Active_Profile
#ENV Active_Profile "$Active_Profile"
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .
#CMD ["/app/main"]

# Create the final image using debian base image
FROM debian:bullseye-slim

# Set the timezone environment variable and install tzdata package
ENV TZ=Asia/Karachi
RUN apt-get update && apt-get install -y tzdata && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone
    
# Import the user and group files from the builder.
#COPY --from=builder /etc/passwd /etc/passwd
#COPY --from=builder /etc/group /etc/group
# Copy our static executable.
COPY --from=builder /app/main /app/main
#COPY /test/main /app/main
COPY app_prod.env /app_prod.env
# Use an unprivileged user.
#USER appuser:appuser
# Run the hello binary.
CMD ["/app/main"]

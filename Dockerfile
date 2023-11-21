FROM golang:1.21-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
# Copy data file
COPY users_with_password_dummy.csv ./
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /worker-request-login
# Run
CMD ["/docker-gs-ping"]
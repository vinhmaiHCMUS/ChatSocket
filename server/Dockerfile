# Sử dụng image có sẵn của Golang để build ứng dụng
FROM golang:1.17

# Sao chép mã nguồn vào container
COPY server.go /app/

# Thiết lập thư mục làm việc
WORKDIR /app

# Build ứng dụng
RUN go build server.go

# Chạy server khi container được khởi chạy
CMD ["./server"]

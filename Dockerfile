# Sử dụng ổ đĩa cơ sở làm cơ sở hình ảnh
FROM golang:1.19
# Thiết lập thư mục làm việc

RUN apt-get update && apt-get install -y curl
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
RUN chmod +x kubectl
RUN mv kubectl /usr/local/bin/

# Thiết lập biến môi trường
ENV APP_DIR /app
WORKDIR $APP_DIR

# Tạo thư mục làm việc
RUN mkdir -p $APP_DIR

# Sao chép mã nguồn ứng dụng vào thư mục làm việc
COPY src/k8s $APP_DIR

# Biên dịch ứng dụng
COPY go.mod go.sum $APP_DIR
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/k8s/main .


# Mở cổng 8080
EXPOSE 8080

# Khởi chạy ứng dụng
CMD ["/app/k8s/main"]

package main

import (
    "fmt"
)

// 1.Consumer là một struct có phụ thuộc vào Service.
type Consumer struct {
	service Service
}

// 2.Service là một interface định nghĩa hành vi của dịch vụ cần sử dụng.
type Service interface {
    Serve()
}

// 3.Implementation là một cài đặt cụ thể của Service.
type Implementation struct{}

// 4.NewConsumer tạo một instance của Consumer với Service đã cung cấp.
func NewConsumer(service Service) *Consumer {
	return &Consumer{service: service}
}

// 5.Serve triển khai hành vi Serve của Service interface.
func (impl *Implementation) Serve() {
    fmt.Println("Service is serving.")
}

// 6.Consume sử dụng dịch vụ đã cung cấp.
func (c *Consumer) Consume() {
    c.service.Serve()
}

func main() {
    // Tạo một instance của Service Implementation.
    service := &Implementation{}

    // Tạo một instance của Consumer với Service đã cung cấp.
    consumer := NewConsumer(service)

    // Sử dụng Consumer.
    consumer.Consume()
}

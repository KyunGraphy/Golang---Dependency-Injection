# Overview

Trong Golang, Dependency Injection (DI) thường được thực hiện bằng cách truyền các phụ thuộc vào hàm hoặc struct thông qua các tham số hoặc trường (field). Điều này cho phép bạn cung cấp các phụ thuộc từ bên ngoài, thay vì cứng cáp kết nối chúng bên trong hàm hoặc struct.

### Example 1
Trong ví dụ này, Circle Rectangle là một struct có phụ thuộc vào Calc. Thay vì khởi tạo Calc trực tiếp trong Circle và Rectangle, chúng ta đã chuyển Calc vào Circle và Rectangle thông qua NewCircle và NewRectangle. Điều này cho phép chúng ta thay đổi cài đặt Calc mà Circle và Rectangle sử dụng mà không cần sửa đổi Calc chính.

### Example 2
Trong ví dụ này, Consumer là một struct có phụ thuộc vào Service. Thay vì khởi tạo Service trực tiếp trong Consumer, chúng ta đã chuyển Service vào Consumer thông qua NewConsumer. Điều này cho phép chúng ta thay đổi cài đặt Service mà Consumer sử dụng mà không cần sửa đổi Consumer chính.

# Work flow
1. Định nghĩa cấu trúc đối tượng - **struct**
2. Định nghĩa phương thức đối tượng - **interface**
3. Tạo cài đặt cụ thể của đối tượng - **struct**
4. Viết phương thức khởi tạo đối tượng - **pointer**
5. Tiến hành xây dựng các phương thức cho đối tượng - **declared in the interface**
6. Sử dụng dịch vụ cung cấp

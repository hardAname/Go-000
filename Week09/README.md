学习笔记

作业：
1. 用 Go 实现一个 tcp server ，用两个 goroutine 读写 conn，两个 goroutine 通过 chan 可以传递 message，能够正确退出。

    作业说明：
    进入tcptest目录，运行:
    ```
    go test -v tcpServer_test.go
    ```
    另起一个终端,进入tcptest目录,运行:
    ```
    go test -v tcpClient_test.go
    ```
    可以定时收发消息, 'ctrl + c' 退出。
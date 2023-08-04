# ssh tunnel demo

## 简述
简单使用 go 语言来做 ssh tunnel 连接到数据库

## 使用

1. 下载依赖
    ```bash
    go get golang.org/x/text
    go get github.com/golang/snappy
    go get golang.org/x/crypto/ssh
    ```
2. 将 dm 驱动放到 GOPATH/src 目录下
3. 先执行 `ssh_tunnel.go` 文件
    ```bash
    go run ssh_tunnel.go
    ```
4. 再执行 `main.go` 文件
    ```shell
    go run main.go
    ```

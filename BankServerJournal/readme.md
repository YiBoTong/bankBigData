# 构建服务
1.打包(一定要改日志输出文件夹配置)
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bank_server_journal main.go
```
2.构建docker镜像（一定要改config.js中的数据库地址）
```
docker build -t bank_server_journal:1.0 .
```
3.导出镜像(在`_docker`文件夹中)
```
docker save -o bank_server_journal_1.0.tar bank_server_journal:1.0
```
4.导入镜像
```docker
docker load -i bank_server_journal_1.0.tar
```
5.运行
```docker
docker run --name bank_server_journal --restart always -v $PWD/run:/app/run -v $PWD/ftp:/app/ftp -d bank_server_journal:1.0
```
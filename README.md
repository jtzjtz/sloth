sloth代码自动生成
====
##支持java、golang模板生成代码
### 一、安装使用
####1.docker安装
######docker push jtzjtz/sloth:tagname

####2.源码安装（[下载地址](https://github.com/jtzjtz/sloth)
######生成java模板：编译条件golang、git
######生成go模板：编译条件golang、protoc、protoc-gen-go、git
###二、java模板自动生成代码，自动上传到使用的项目git
###三、golang模板代码自动生成两层代码，并上传到git，client和server通过go mod 安装包，使用对应代码
######1.自动生成的包文件，[demo地址](https://github.com/jtzjtz/ys_pack)
######2.client：消费者，引入包即可使用[demo地址](https://github.com/jtzjtz/ys_server)
######3.server：grpc服务提供者，引入包即可使用 [demo地址](https://github.com/jtzjtz/ys_api)




![](https://github.com/jtzjtz/sloth/blob/main/jt.png)
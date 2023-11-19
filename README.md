golang-docker
this is a test project,build develop env for golang in docker
## golang 环境设置
* 基于ubuntu
* 具体docker构建见 go-backend/Dockerfile

## docker命令解释
- FROM 命令，docker构建的基础镜像
- WORKDIR 命令，docker的工作目录，绝对路径
- ENV 命令，设置环境变量，这里也是绝对路径。其中GOPATH是下载go组件时存放的路径，GOCACHE是构建go组件时的缓存目录。
- CPYOPY命令，复制命令，目录是相对WORKDDIR的相对路径。
- RUN 命令，执行一些shell或者系统的指令，第一个参数是执行的命令或者可执行文件的路径，后面是命令的参数
- CMD 命令，docker启动时执行的默认命令。

## docker-compose 配置文件
因为go主要作为开发环境，涉及到文件的更新，所以在yml文件中建立了本地目录和docker路径之间的映射。

因为docker是监听了http端口80，所以建立docker和主机之间的端口映射
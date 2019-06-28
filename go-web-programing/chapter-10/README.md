#chapter-10

## 将应用部署到独立的服务器

美国国家技术标准技术研究所定义了当今广为使用的 3 种服务模型（National Institute of Standrds and Technology, US Department of COmmerce, NIST）

Iaas -- Infrastructure-as-a-Service  -- 基础设施即服务
Paas -- Platform-as-a-Service  -- 平台即服务
Saas -- Software-as-a-Service  -- 应用即服务

Upstart 和 systemd 守护进程。

## 将应用部署到 google app engine

## 将应用部署到 docker
docker 守护进程运行在宿主操作系统之上，该进程会对客户端发送的服务请求进行应答，并对容器进行管理。

docker 容器是对运行特定应用所需的全部程序（包括操作系统）的一种轻量级虚拟化。轻量级容器会让应用以及与之相关联的其他程序认为自己独占了整个操作系统以及所有硬件，但是实际上并非如此，多个应用共享同一个宿主操作系统。

docker 容器都基于 docker 镜像构建，后者是辅助容器进行启动的只读模板，所有容器都需要通过镜像启动。有好几种不同的方法可以创建 docker 镜像，其中一种就是在一个名为 dockerfile 的文件里包含一系列指令。

docker 镜像既可以以本地形式存储在运行着 docker 守护进程的机器上，也可以被托管至名为 docker 注册中心的 docker 镜像资源库里面。

`https://github.com/docker/machine` docker 部署到特定的一些云供应商。
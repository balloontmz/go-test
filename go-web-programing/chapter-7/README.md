# chapter-7

soap 风格和 restful 风格

## web 服务

web 服务是与其他软件程序进行交互的采用 http 进行交互的程序。

soap 通常是 rpc 的，功能驱动。restful 服务通常是数据驱动。

restful：速度快，构建简单

soap：安全且健壮。

## 基于 soap 的 web 服务

## 基于 rest 的 web 服务

Representational State Transfer -- 具象状态传输 -- 设计理念

PUT 方法是幂等的，POST 方法不是。多次请求，PUT 只会创建一次资源，POST 会有多次。

REST 完成理论上的任意动作：
1. 把过程抽象成对象，或者把动作转换为名词，然后将其用作为资源
2. 将动作用作资源的属性。

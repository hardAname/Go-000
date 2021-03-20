学习笔记

# 毕业总结：
        我当前对于微服务的理解：微服务主要是为了便于整体项目的迭代和管理，需要在理解整体业务范围的情况下做拆分，
    便于子模块的迭代，以及在项目发展得越来越庞杂的情况下，便于整体项目正常运行的管理。但是缺点是微服务的基础设施搭建会比较
    难度比较高，而且提高了迭代的灵活性的代价就是降低了整体服务的性能。
    
# 毕业项目：
        当前负责的模块是一个长连接推送服务，目前掺杂了一点业务逻辑，不是个存粹的保证可达的长连接推送服务。
    项目结构设计中考虑了代码分层，DI，使用了wire来做项目的启动入口；目前没有使用orm框架，只是在api层和service层，
    以及service和dao层间做了下结构体的转换。从代码可见性上，外部调用只能访问到api目录和cmd目录，其余的项目逻辑都
    放在了internal目录中，无需提供给外部服务调用。
    
    以下是项目的结构树：
    ```
    ├─api
    │  └─pushMsgRecvServer
    ├─cmd
    ├─configs
    ├─doc
    ├─internal
    │  ├─dao
    │  ├─di
    │  ├─model
    │  │  ├─comet
    │  │  ├─cometAclient
    │  │  └─cometAinner
    │  ├─server
    │  │  ├─grpc
    │  │  │  └─pushMsgRecvServer
    │  │  ├─http
    │  │  ├─kcp
    │  │  ├─tcp
    │  │  └─websock
    │  │      └─lib
    │  │          └─websocket
    │  ├─service
    │  └─util
    │      ├─alert
    │      ├─jwtUtil
    │      ├─MPrometheus
    │      ├─xlog
    │      └─zookeeper
    └─test
        ├─kcptest
        ├─mock
        ├─outOfOrder
        └─timestamp
    ```
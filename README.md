# Cloudgo
## 1.编程 web 服务程序 类似 cloudgo 应用。
###   要求有详细的注释
###   是否使用框架、选哪个框架自己决定请在 README.md 说明你决策的依据

```javascript
package main

import "github.com/go-martini/martini"  //使用martini框架

func main() {
  m := martini.Classic()//获取一个martini实例释
  m.Get("/hello/:name", func(params martini.Params) string {// 用户自定义路由规则
    return "Hello " + params["name"]
  })
  m.Run()// 运行服务器
}
```

#### Martini.Classic()里，创建了一个经典的martini对象。
```javascript
func Classic() *ClassicMartini { 
    r := NewRouter() 新建标准路由 
    m := New() 新建Martini对象 
    m.Use(Logger())//注册中间件，标准logger 
    m.Use(Recovery()) //注册中间件，标准recover 
    m.Use(Static("public")) //注册中间件Static 
    m.MapTo(r, (*Routes)(nil)) //Injector的Mapto方法，实现类型和对象的关联注入 
    m.Action(r.Handle) //设置处理函数为路由处理函数 
    return &ClassicMartini{m, r} //返回一个ClassMartini实例，继承了Martini的结构， 提升了martini以及Router的相关方法 
}
```

#### Run()
```javascript
func (m *Martini) Run() {
    port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000" //默认是3000端口
	}
	host := os.Getenv("HOST")
	m.RunOnAddr(host + ":" + port)
}
```

## 2.使用 curl 测试，将测试结果写入 README.md

简单的curl测试一：

    可以看到，输入命令行curl -v http:localhost:3000/hello/testusername后，收到了报文内容中有Hello testusername,在服务端可以看到 接收到来自127.0.0.1的请求和完成请求所需要的时间：275us
    
    ![image](https://github.com/YlingMA/Cloudgo/raw/master/image/图2-1.png)
    
curl测试一千次：

    通过图2-2的命令进行1000次测试（注意这里需要进入root模式）
     ![image](https://github.com/YlingMA/Cloudgo/raw/master/image/图2-2.png)
     
    结果见图2-3：
    可以看到，每一次都成功了并且花费时间不多
     ![image](https://github.com/YlingMA/Cloudgo/raw/master/image/图2-3.png)

## 3.使用 ab 测试，将测试结果写入 README.md。并解释重要参数。
ab测试样例一：

    每组10个请求，总共100个请求。结果如图3-1，3-2：
     ![image](https://github.com/YlingMA/Cloudgo/raw/master/image/图3-1.png)
     
     图3-1
     
     ![image](https://github.com/YlingMA/Cloudgo/raw/master/image/图3-2.png)
     
     图3-2
    
ab测试样例二：

    每组10个请求，总共100个请求。结果如图3-3，3-4：
     ![image](https://github.com/YlingMA/Cloudgo/raw/master/image/图3-3.png)
     
    图3-3
     ![image](https://github.com/YlingMA/Cloudgo/raw/master/image/图3-4.png)
     
    图3-4

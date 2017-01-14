# ngo-blog

A little blog web frame created by Golang and Angularjs. It is easy to setup one-page blog Web site by Golang and Angularjs. It is a MVC blog web frame.   

# Why should you want this blog web frame?

- It is simple. It only use Golang and Angularjs and Bootstrap. 
- It is easy to deploy. Just make and run. 
- It is extended. You can do more things than just blog.
- Full support Angularjs. Server and client are separate.

# Supported features
- Article list, view, writer
- Up and down statistics
- Pretty code

# Usage:

## Dependencies:
- Golang http://golang.org/
- jQuery http://jquery.com/
- Angularjs https://angularjs.org/
- Twitter Bootstrap http://twitter.github.com/bootstrap/

## Install and run:

    mkdir ~/go_codes/src
    cd ~/go_codes/src
    git clone https://github.com/DoBetter-pan/ngo-blog.git
    cd ngo-blog
    source setenv.sh
    You should have installed mysql. Then import the data in server/data/ngo-blog.sql
    vim server/datawrapper/datawrapper.go
    change the following line using your database user and password:
    var WrapperConfig *DbConfig = &DbConfig{"tcp", "127.0.0.1", "3306", "test", "123456", "ngo-blog", "charset=utf8"}
    go get github.com/go-sql-driver/mysql 
    go build
    ./ngo-blog -host=127.0.0.1 -port=9898 
Then you can visit http://127.0.0.1:9898/ to view the blogs.
Then you can visit http://127.0.0.1:9898/blog/admin to write a new blog.
The default user and password are yingx and 123456.

Note:
You must setup your golang environment first and will change your GOPATH when executing "source setenv.sh".
You must install mysql first.

## Develop:

You can create a new module by modulehelper.sh.
    
# What should we do next?
- up and down statistics
- comments
- markdown
- pretty code 
- ...

# Support and contact

If you have any question and advice, you can cantact me in QQ Group: 536069420.


# ngo-blog

ngo-blog是使用Golang与Angularjs编写的一个小的博客框架。利用ngo-blog可以很容易的构建一个单页的博客，同时ngo-blog也是一个MVC的博客框架。   

# 为什么我们要使用该博客框架呢？

- 他是简单的。 该框架在后端仅仅使用了Golang；在前端使用了Angularjs与Bootstrap等前端框架. 
- 他是非常简单部署的。 仅仅需要设置了Golang环境，然后 go build，然后就可以启动主服务了。
- 他是可扩展的。 你可以使用该框架做更多的事情，不仅仅事博客。
- 完全支持Angularjs。 服务与客户端事完全分离的。

# 支持的功能
- 博客的列表显示、单个显示、撰写
- 赞踩功能
- 美化显示代码

# 使用与安装:

## 依赖:
- Golang http://golang.org/
- jQuery http://jquery.com/
- Angularjs https://angularjs.org/
- Twitter Bootstrap http://twitter.github.com/bootstrap/

## 安装与运行:

    mkdir ~/go_codes/src
    cd ~/go_codes/src
    git clone https://github.com/DoBetter-pan/ngo-blog.git
    cd ngo-blog
    source setenv.sh
    假设你已经安装了mysql. 执行server/data/ngo-blog.sql脚本导入数据。
    vim server/datawrapper/datawrapper.go
    使用自己数据库的密码代替下行中的用户（test）与密码（123456）：
    var WrapperConfig *DbConfig = &DbConfig{"tcp", "127.0.0.1", "3306", "test", "123456", "ngo-blog", "charset=utf8"} 
    go get github.com/go-sql-driver/mysql 
    go build
    ./ngo-blog -host=127.0.0.1 -port=9898 
访问 http://127.0.0.1:9898/ 可以浏览博客.
访问 http://127.0.0.1:9898/blog/admin 可以撰写新的博客.
默认的登陆用户与密码是： yingx 与 123456.

注:
你必须先安装Golang的环境。
你必须先安装mysql。

## 开发:

你可以使用modulehelper.sh创建新的模块。
    
# 接下来将做什么?
- 点赞功能
- 评论功能
- markdown支持
- 源代码美化显示
- 管理页面的完善 
- ...

# 支持与联系

假如你有任何问题与建议，可以在QQ群：536069420联系我。


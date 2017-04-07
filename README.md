[TOC]


#LVBU
##首次运行
在app.conf里面配置的mysql主机上面执行:


	root@localhost:~#mysql -u username -p
	mysql>create lvbu;
	mysql>exit;

	root@localhost:~#mysql -u username -p lvbu < path/to/lvbu.sql
	root@localhost:~#go get github.com/astaxie/beego
	root@localhost:~#cd lvbu
	root@localhost:~#go build main.go
	root@localhost:~#./main

##初始帐号密码
>website  :   127.0.0.1
username:***root***
password:***1***  **(数字壹)**
##运行说明
***1.***首次访问可以直接访问http://localhost/, 点立即体验
***2.***自已发现

##主要目录讲解

***1.***  code:存放每个项目的从GIT仓库download的文件
***2.*** .code:存放部署时，会在目录下面创建本地仓库临时文件
***3.*** .prohisconf:每次版本升级，会把上个版本的配置项生成文件，保存到该目录下，文件命名规则

 pro_sign(目录)
: DE(环境标识)_版本号.conf

***4.*** .profile : 项目编译依赖如：java6,java7,tomcat
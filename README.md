#LVBU
##首次运行
在app.conf里面配置的mysql主机上面执行:
	root@localhost:~#mysql -u username -p
	mysql>create lvbu;
	mysql>exit;

	root@localhost:~#mysql -u username -p lvbu < path/to/lvbu.sql
	root@localhost:~#cd lvbu
	root@localhost:~#go run main.go
##初始帐号密码
username:***root***
password:***1***  **(数字壹)**
##运行说明
***1.***首次访问可以直接访问http://localhost/, 点立即体验
***2.***自已发现

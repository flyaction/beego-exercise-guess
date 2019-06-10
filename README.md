# beego-exercise-guess
# beego-url
https://www.kancloud.cn/hello123/beego/126086
# 常用命令
bee new 新建项目结构
bee run 自动编译部署
bee generate 自动生成代码
bee pack 打包
bee generate scaffold user -fields="id:int64,name:string,gender:int,age:int" -driver=mysql -conn="root:@tcp(127.0.0.1:3306)/beego"  //脚手架
# 监控检查
app.conf加入EnableAdmin = true
http://localhost:8088/healthcheck

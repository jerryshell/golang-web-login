# Golang login demo

Golang 用户登录 Demo

数据库驱动：[mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)

# 功能列表

- [x] 登录
- [x] 注销
- [ ] 注册
- [ ] 更新信息

# 建表语句

```sql
CREATE TABLE `user` ( `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, `username` TEXT NOT NULL, `password` TEXT NOT NULL, `email` TEXT )
```
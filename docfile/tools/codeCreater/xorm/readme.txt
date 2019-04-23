使用xorm工具，根据数据库自动生成 go 代码 - artong0416 - 博客园 https://www.cnblogs.com/artong0416/p/7456674.html
go get github.com/go-xorm/cmd/xorm

go get github.com/go-sql-driver/mysql  //MyMysql
go get github.com/ziutek/mymysql/godrv  //MyMysql
go get github.com/lib/pq  //Postgres
go get github.com/mattn/go-sqlite3  //SQLite

go get github.com/go-xorm/xorm


编译github.com/go-xorm/cmd/xorm时需要cloud.google.com/go/civil
但因墙无法下载，所以手动使用以下地址：
https://github.com/GoogleCloudPlatform/google-cloud-go

git clone https://github.com/GoogleCloudPlatform/google-cloud-go.git


环境配置好后使用
build.go
created.bat即可




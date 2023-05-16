module github.com/server/txs-analysis

go 1.16

require (
	github.com/astaxie/beego v1.12.3
	github.com/ethereum/go-ethereum v1.11.6
	github.com/go-sql-driver/mysql v1.7.1
	github.com/vesoft-inc/nebula-go/v3 v3.4.0
	github.com/zhihu/norm v0.1.11
)

replace github.com/zhihu/norm => ../norm

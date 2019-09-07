module github.com/Myriad-Dreamin/nyan-oj-backend

go 1.12

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43

require (
	github.com/Myriad-Dreamin/core-oj v1.0.1-0.20190902031246-7d62d2a419be
	github.com/Myriad-Dreamin/gin-middleware v0.0.0-20190902015905-44035a61dbe2
	github.com/casbin/casbin/v2 v2.0.2
	github.com/casbin/xorm-adapter v0.0.0-20190806085643-0629743c2857
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.6
)

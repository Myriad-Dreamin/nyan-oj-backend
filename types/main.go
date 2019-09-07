package main

import (
	"fmt"
	"os"

	"github.com/Myriad-Dreamin/core-oj/log"
	morm "github.com/Myriad-Dreamin/gin-middleware/sample/user/orm"
	rbac "github.com/Myriad-Dreamin/gin-middleware/sample/user/rbac"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Server struct {
	engine *xorm.Engine
	logger log.TendermintLogger
}

func NewServer() (srv *Server, err error) {
	srv = new(Server)

	srv.logger, err = log.NewZapColorfulDevelopmentSugarLogger()
	if err != nil {
		return nil, err
	}

	return
}

func (srv *Server) prepareDatabase(driver, connection string) error {
	var err error

	srv.engine, err = xorm.NewEngine(driver, connection)
	if err != nil {
		srv.logger.Error("prepare failed", "error", err)
		return err
	}

	err = rbac.Init(srv.engine)
	if err != nil {
		srv.logger.Error("prepare failed", "error", err)
		return err
	}

	morm.RegisterEngine(srv.engine)

	srv.engine.ShowSQL(true)
	return nil
}

func (srv *Server) Close() error {
	return nil
}

func (srv *Server) Serve(port string) error {

}

func main() {
	var srv, err = NewServer()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = srv.prepareDatabase("mysql", "coreoj-admin:123456@tcp(127.0.0.1:3306)/coreoj?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err := srv.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	if err = srv.Serve(":23336"); err != nil {
		fmt.Println(err)
		return
	}
}

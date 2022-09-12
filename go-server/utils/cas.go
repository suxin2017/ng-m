package utils

import (
	"log"
	"os"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"suxin2017.com/server/models"
)

var casLog = log.New(os.Stdout, "[cas] ", log.Lshortfile|log.Ldate|log.Ltime)
var casEnforcer *casbin.Enforcer

func InitCasForGorm() {
	a, dbErr := gormadapter.NewAdapterByDB(models.DB)
	casLog.Println("初始化连接数据库")
	PanicErr(dbErr)
	e, casErr := casbin.NewEnforcer("rbac.conf", a)
	casLog.Println("初始化casbin与数据库连接")
	PanicErr(casErr)
	casEnforcer = e
	casEnforcer.LoadPolicy()
	casLog.Println("初始化策略")
	if hasAdmin, _ := e.Enforce("user:admin", "*", "*"); !hasAdmin {
		e.AddPolicy("user:admin", "*")
		casEnforcer.SavePolicy()
	}
}

func AddPathPolice(role string, path string) {

}

var userPrefix = "user:"

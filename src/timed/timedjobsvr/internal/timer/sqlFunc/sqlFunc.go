package sqlFunc

import (
	"context"
	"fmt"
	"github.com/dop251/goja"
	"github.com/i-Things/things/shared/conf"
	"github.com/i-Things/things/shared/domain/task"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/stores"
	"github.com/i-Things/things/src/timed/internal/domain"
	"github.com/i-Things/things/src/timed/timedjobsvr/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"strings"
	"time"
)

type SqlFunc struct {
	SvcCtx     *svc.ServiceContext
	ctx        context.Context
	Task       *domain.TaskInfo
	vm         *goja.Runtime
	ExecuteLog []*domain.ScriptLog
	SelectNum  int64 //查询的数量
	ExecNum    int64 //执行的数量
	kvKeyPre   string
	logx.Logger
}

func NewSqlFunc(ctx context.Context, svcCtx *svc.ServiceContext, task *domain.TaskInfo, vm *goja.Runtime) *SqlFunc {
	kvKeyPre := fmt.Sprintf("timed:sql:%s:", task.GroupCode)
	if code := task.Env["code"]; code != "" {
		kvKeyPre = fmt.Sprintf("timed:sql:%s:", task.GroupCode)
	}
	return &SqlFunc{SvcCtx: svcCtx, ctx: ctx, Logger: logx.WithContext(ctx), Task: task, vm: vm, kvKeyPre: kvKeyPre}
}

func (s *SqlFunc) Register() error {
	var funcList = []struct {
		Name string
		f    func(in goja.FunctionCall) goja.Value
	}{
		{"Set", s.Set()},
		{"Get", s.Get()},
		{"Select", s.Select()},
		{"Exec", s.Exec()},
		{"LogError", s.LogError()},
		{"LogInfo", s.LogInfo()},
		{"GetEnv", s.GetEnv()},
		{"Hexists", s.Hexists()},
		{"Hdel", s.Hdel()},
		{"Hget", s.Hget()},
		{"Hset", s.Hset()},
		{"HgetAll", s.HGetAll()},
		{"CreateOne", s.CreateOne()},
		{"TaskSendSqlJs", s.TaskSendSqlJs()},
	}
	for _, f := range funcList {
		err := s.vm.Set(f.Name, f.f)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SqlFunc) getConn(in goja.FunctionCall, tp string) (*gorm.DB, func() error) {
	dsn := s.Task.Env[task.SqlEnvDsn]
	dbType := s.Task.Env[task.SqlEnvDBType]
	if len(in.Arguments) > 1 {
		dbName := in.Arguments[1].String()
		c, ok := s.Task.Sql.Config.Database[dbName]
		if ok {
			dsn = c.DSN
			dbType = c.DBType
		}
	}
	if dsn == "" { //判断系统配置
		c, ok := s.Task.Sql.Config.Database[tp]
		if ok {
			dsn = c.DSN
			dbType = c.DBType
		} else {
			return stores.GetCommonConn(s.ctx), func() error {
				return nil
			}
		}
	}
	conn, err := stores.GetConn(conf.Database{
		DBType: dbType,
		DSN:    dsn,
	})
	if err != nil {
		panic(errors.Database.AddMsgf("getConn.GetConn failure dsn:%v dbType:%v err:%v", dsn, dbType, err))
	}
	db, err := conn.DB()
	if err != nil {
		panic(errors.Database.AddMsgf("getConn.conn.DB failure dsn:%v dbType:%v err:%v", dsn, dbType, err))
	}
	return conn, db.Close
}
func (s *SqlFunc) GetStringKey(key string) string {
	return s.kvKeyPre + "string:" + key
}
func (s *SqlFunc) GetHashKey(key string) string {
	return s.kvKeyPre + "hash:" + key
}

func (s *SqlFunc) GetHashField(field string) string {
	date := time.Now().Format("2006-01-02")
	return fmt.Sprintf("%s:%s", date, field)
}
func (s *SqlFunc) GetHashFieldWithDay(field string, day int) string {
	date := time.Now().Add(time.Hour * 24 * time.Duration(day)).Format("2006-01-02")
	return fmt.Sprintf("%s:%s", date, field)
}
func (s *SqlFunc) ToRealHashField(field string) string {
	_, ret, find := strings.Cut(field, ":")
	if !find {
		return field
	}
	return ret
}

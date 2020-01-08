package repos

import (
	"github.com/jinzhu/gorm"
	"goim-pro/internal/app/repos/address"
	"goim-pro/internal/app/repos/user"
	"goim-pro/pkg/logs"
)

var logger = logs.GetLogger("ERROR")

type RepoServer struct {
	UserRepo    user.IUserRepo
	AddressRepo address.IAddress
}

func New() *RepoServer {
	//redisDB := redsrv.GetRedisClient()
	return &RepoServer{
		UserRepo:    user.NewUserRepo(),
		AddressRepo: address.New(),
	}
}

func initialMysqlTables(db *gorm.DB) (err error) {
	if !db.HasTable(user.User{}) {
		err = db.Set(
			"gorm:table_options",
			"ENGINE=InnoDB DEFAULT CHARSET=utf8",
		).CreateTable(
			user.User{},
		).Error
		if err != nil {
			logger.Errorf("initial mysql tables [users] error: %v\n", err)
			return
		}
	}
	return
}

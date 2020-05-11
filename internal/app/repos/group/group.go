package group

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	. "goim-pro/internal/app/models"
	"goim-pro/pkg/logs"
	"goim-pro/pkg/utils"
	"time"
)

type GroupImpl Group

type MemberImpl Member

type IGroupRepo interface {
	CreateGroup(groupProfile *Group) (newGroup *Group, err error)
	RemoveGroupByGroupId(groupId string, isForce bool) (err error)

	FindOneGroupMember(condition map[string]interface{}) (memberProfile *Member, err error)
	InsertMembers(members ...*Member) (err error)
	FindOneGroup(condition map[string]interface{}) (groupProfile *Group, err error)
}

var logger = logs.GetLogger("ERROR")
var mysqlDB *gorm.DB

func NewGroupRepo(db *gorm.DB) IGroupRepo {
	mysqlDB = db
	return &GroupImpl{}
}

// CreateGroup - create group with group profile and members
func (i *GroupImpl) CreateGroup(groupProfile *Group) (newGroup *Group, err error) {
	_db := mysqlDB.Create(groupProfile)
	if _db.Error != nil {
		err = _db.Error
		logger.Errorf("create group error: %s", err.Error())

		return nil, err
	}
	return _db.Value.(*Group), nil
}

func (i *GroupImpl) FindOneGroup(condition map[string]interface{}) (groupProfile *Group, err error) {
	err = mysqlDB.Preload("Member").Find(&groupProfile, condition).Error
	return
}

func (i *GroupImpl) InsertMembers(members ...*Member) (err error) {
	var buffer bytes.Buffer
	sql := "INSERT INTO `members` (`userId`, `alias`, `role`, `status`, `createdAt`, `updatedAt`) values"
	if _, err := buffer.WriteString(sql); err != nil {
		return err
	}
	for i, e := range members {
		nowDateTime := utils.TimeFormat(time.Now(), utils.MysqlDateTimeFormat)
		if i == len(members)-1 {
			buffer.WriteString(fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s');", e.UserId, e.Alias, "1", "NORMAL", nowDateTime, nowDateTime))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s'),", e.UserId, e.Alias, "1", "NORMAL", nowDateTime, nowDateTime))
		}
	}
	if err := mysqlDB.Exec(buffer.String()).Error; err != nil {
		logger.Errorf("exec insert members error: %v", err)
		return err
	}
	return nil
}

func (i *GroupImpl) FindOneGroupMember(condition map[string]interface{}) (memberProfile *Member, err error) {
	memberProfile = &Member{}
	db := mysqlDB.Where(condition).First(&memberProfile)
	if db.RecordNotFound() {
		err = utils.ErrUserNotExists
	} else if err = db.Error; err != nil {
		logger.Errorf("error happened to query group information: %s", err.Error())
	}
	return
}

func (i *GroupImpl) RemoveGroupByGroupId(groupId string, isForce bool) (err error) {
	_db := mysqlDB
	if isForce {
		_db = mysqlDB.Unscoped()
	}
	_db = _db.Delete(&Group{}, "groupId = ?", groupId)
	if _db.RecordNotFound() {
		logger.Warningln("remove group fail, groupId not found")
	} else if _db.Error != nil {
		logger.Errorf("error happened to remove group: %v", _db.Error)
		err = _db.Error
	}
	return
}

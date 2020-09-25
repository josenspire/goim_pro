package contactsrv

import (
	"fmt"
	"github.com/jinzhu/gorm"
	protos "goim-pro/api/protos/salty"
	. "goim-pro/internal/app/constants"
	"goim-pro/internal/app/models"
	. "goim-pro/internal/app/models/errors"
	. "goim-pro/internal/app/repos/contact"
	. "goim-pro/internal/app/repos/user"
	"goim-pro/internal/app/services/converters"
	mysqlsrv "goim-pro/pkg/db/mysql"
	redsrv "goim-pro/pkg/db/redis"
	"goim-pro/pkg/errors"
	"goim-pro/pkg/logs"
	"goim-pro/pkg/utils"
)

var (
	logger  = logs.GetLogger("INFO")
	myRedis redsrv.IMyRedis
	mysqlDB *gorm.DB

	userRepo    IUserRepo
	contactRepo IContactRepo
	//notificationLogRepo INotificationLogRepo
)

type IContactService interface {
	RequestContact(userId, contactId, reqReason string) (tErr *TError)
	RefusedContact(userId, contactId, refusedReason string) (tErr *TError)
	AcceptContact(userId, contactId string) (tErr *TError)
	DeleteContact(userId, contactId string) (tErr *TError)
	UpdateRemarkInfo(userId, contactId string, contactRemark *protos.ContactRemark) (tErr *TError)
	GetContacts(userId string) (contacts []models.Contact, tErr *TError)
	GetContactOperationMessageList(userId string, maxMessageTime int64) (contactOptsList []models.Contact, tErr *TError)
}

type ContactService struct {
}

func New() IContactService {
	myRedis = redsrv.NewRedis()
	mysqlDB = mysqlsrv.NewMysql()

	userRepo = NewUserRepo(mysqlDB)
	contactRepo = NewContactRepo(mysqlDB)
	//notificationLogRepo = NewNotificationLogRepo(mysqlDB)

	return &ContactService{}
}

// RequestContact: request add contact
func (cs *ContactService) RequestContact(userId, contactId, reqReason string) (tErr *TError) {
	contact, err := userRepo.FindByUserId(contactId)
	if err != nil {
		logger.Errorf("find contact by userId error: %s", err.Error())
		return NewTError(protos.StatusCode_STATUS_INTERNAL_SERVER_ERROR, err)
	}
	if contact == nil {
		return NewTError(protos.StatusCode_STATUS_BAD_REQUEST, errmsg.ErrInvalidContact)
	}

	isExists, err := contactRepo.IsContactExists(userId, contactId)
	if err != nil {
		logger.Errorf("checking contact error: %s", err.Error())
		return NewTError(protos.StatusCode_STATUS_INTERNAL_SERVER_ERROR, err)
	}
	if isExists {
		return NewTError(protos.StatusCode_STATUS_BAD_REQUEST, errmsg.ErrContactAlreadyExists)
	}

	// TODO: cache in redis, should replace to Push notification server
	ctKey := fmt.Sprintf("CT-REQ-%s-%s", userId, contactId)
	cacheContent := make(map[string]interface{})
	cacheContent["userId"] = userId
	cacheContent["contactId"] = contactId
	cacheContent["addReason"] = reqReason
	cacheContent["rejectReason"] = ""
	cacheContent["messageId"] = utils.NewULID()
	cacheContent["createdTime"] = utils.MakeTimestamp()
	cacheContent["isNeedRemind"] = true
	cacheContent["operationType"] = protos.ContactOperationMessage_ACCEPT_ACTIVE

	jsonStr, err := utils.TransformMapToJSONString(cacheContent)
	if err != nil {
		logger.Errorf("redis operations error, transform map error: %s", err.Error())
	}
	err = myRedis.RSet(ctKey, jsonStr, DefaultExpiresTime)

	if err != nil {
		logger.Errorf("redis cache error: %s", err.Error())
		return NewTError(protos.StatusCode_STATUS_INTERNAL_SERVER_ERROR, err)
	}
	return
}

// refused request contact
func (cs *ContactService) RefusedContact(userId, contactId, refusedReason string) (tErr *TError) {
	contact, err := userRepo.FindByUserId(contactId)
	if err != nil {
		logger.Errorf("find contact by userId error: %s", err.Error())
		return NewTError(protos.StatusCode_STATUS_INTERNAL_SERVER_ERROR, err)
	}
	if contact == nil {
		return NewTError(protos.StatusCode_STATUS_BAD_REQUEST, errmsg.ErrInvalidContact)
	}

	isExists, err := contactRepo.IsContactExists(userId, contactId)
	if err != nil {
		logger.Errorf("checking contact error: %s", err.Error())
		return NewTError(protos.StatusCode_STATUS_INTERNAL_SERVER_ERROR, err)
	}
	if isExists {
		return NewTError(protos.StatusCode_STATUS_BAD_REQUEST, errmsg.ErrContactAlreadyExists)
	}

	// TODO: cache in redis, should replace to Push notification server
	ctKey := fmt.Sprintf("CT-REF-%s-%s", userId, contactId)
	cacheContent := make(map[string]interface{})
	cacheContent["userId"] = userId
	cacheContent["contactId"] = contactId
	cacheContent["addReason"] = ""
	cacheContent["rejectReason"] = refusedReason
	cacheContent["messageId"] = utils.NewULID()
	cacheContent["createdTime"] = utils.MakeTimestamp()
	cacheContent["isNeedRemind"] = true
	cacheContent["operationType"] = protos.ContactOperationMessage_REJECT_PASSIVE

	jsonStr, err := utils.TransformMapToJSONString(cacheContent)
	if err != nil {
		logger.Errorf("redis operations error, transform map error: %s", err.Error())
	}
	if len(myRedis.RGet(ctKey)) > 0 { // already send notification
		return NewTError(protos.StatusCode_STATUS_BAD_REQUEST, errmsg.ErrContactRepeatOperation)
	}
	err = myRedis.RSet(ctKey, jsonStr, DefaultExpiresTime)
	if err != nil {
		logger.Errorf("redis cache error: %s", err.Error())
		return NewTError(protos.StatusCode_STATUS_INTERNAL_SERVER_ERROR, err)
	}
	return
}

// accept contact request
func (cs *ContactService) AcceptContact(userId, contactId string) (tErr *TError) {
	contact, err := userRepo.FindByUserId(contactId)
	if err != nil {
		logger.Errorf("find contact by userId error: %s", err.Error())
		return NewTError(protos.StatusCode_STATUS_INTERNAL_SERVER_ERROR, err)
	}
	if contact == nil {
		return NewTError(protos.StatusCode_STATUS_BAD_REQUEST, errmsg.ErrInvalidContact)
	}

	isExists, err := contactRepo.IsContactExists(userId, contactId)
	if err != nil {
		logger.Errorf("checking contact error: %s", err.Error())
		return NewTError(protos.StatusCode_STATUS_INTERNAL_SERVER_ERROR, err)
	}
	if isExists {
		return NewTError(protos.StatusCode_STATUS_BAD_REQUEST, errmsg.ErrContactAlreadyExists)
	}

	// TODO: cache in redis, should replace to Push notification server
	ctKey := fmt.Sprintf("CT-ACP-%s-%s", userId, contactId)
	if len(myRedis.RGet(ctKey)) > 0 { // already send notification
		return NewTError(protos.StatusCode_STATUS_BAD_REQUEST, errmsg.ErrContactRepeatOperation)
	}

	cacheContent := make(map[string]interface{})
	cacheContent["userId"] = userId
	cacheContent["contactId"] = contactId
	cacheContent["addReason"] = ""
	cacheContent["rejectReason"] = ""
	cacheContent["messageId"] = utils.NewULID()
	cacheContent["createdTime"] = utils.MakeTimestamp()
	cacheContent["isNeedRemind"] = true
	cacheContent["operationType"] = protos.ContactOperationMessage_ACCEPT_PASSIVE
	jsonStr, err := utils.TransformMapToJSONString(cacheContent)

	err = myRedis.RSet(ctKey, jsonStr, DefaultExpiresTime)
	if err != nil {
		logger.Errorf("redis cache error: %s", err.Error())
		return NewTError(protos.StatusCode_STATUS_INTERNAL_SERVER_ERROR, err)
	}

	if err = handleAcceptContact(userId, contactId); err != nil {
		logger.Errorf("insert contact error: %s", err.Error())
		return NewTError(protos.StatusCode_STATUS_INTERNAL_SERVER_ERROR, err)
	}
	return
}

// delete contact
func (cs *ContactService) DeleteContact(userId, contactId string) (tErr *TError) {
	isExists, err := contactRepo.IsContactExists(userId, contactId)
	if err != nil {
		logger.Errorf("checking contact error: %s", err.Error())
		return NewTError(protos.StatusCode_STATUS_INTERNAL_SERVER_ERROR, err)
	}
	if !isExists {
		return NewTError(protos.StatusCode_STATUS_BAD_REQUEST, errmsg.ErrContactNotExists)
	}

	// TODO: cache in redis, should replace to Push notification server
	ctKey := fmt.Sprintf("CT-DEL-%s-%s", userId, contactId)
	if len(myRedis.RGet(ctKey)) > 0 { // already send notification
		return NewTError(protos.StatusCode_STATUS_BAD_REQUEST, errmsg.ErrContactRepeatOperation)
	}

	cacheContent := make(map[string]interface{})
	cacheContent["userId"] = userId
	cacheContent["contactId"] = contactId
	cacheContent["addReason"] = ""
	cacheContent["rejectReason"] = ""
	cacheContent["messageId"] = utils.NewULID()
	cacheContent["createdTime"] = utils.MakeTimestamp()
	cacheContent["isNeedRemind"] = true
	cacheContent["operationType"] = protos.ContactOperationMessage_DELETE_ACTIVE
	jsonStr, err := utils.TransformMapToJSONString(cacheContent)

	err = myRedis.RSet(ctKey, jsonStr, DefaultExpiresTime)
	if err != nil {
		// TODO: should log down exception information
		logger.Errorf("redis cache error: %s", err.Error())
	}

	if err = handleDeleteContact(userId, contactId); err != nil {
		logger.Errorf("remove contact error: %s", err.Error())
		return NewTError(protos.StatusCode_STATUS_INTERNAL_SERVER_ERROR, err)
	}

	return
}

// update contact remark profile
func (cs *ContactService) UpdateRemarkInfo(userId, contactId string, contactRemark *protos.ContactRemark) (tErr *TError) {
	isExists, err := contactRepo.IsContactExists(userId, contactId)
	if err != nil {
		logger.Errorf("checking contact error: %s", err.Error())
		return NewTError(protos.StatusCode_STATUS_INTERNAL_SERVER_ERROR, err)
	}
	if !isExists {
		return NewTError(protos.StatusCode_STATUS_BAD_REQUEST, errmsg.ErrContactNotExists)
	}

	if err = handleUpdateContactRemark(userId, contactId, contactRemark); err != nil {
		logger.Errorf("update contact remark error: %s", err.Error())
		return NewTError(protos.StatusCode_STATUS_INTERNAL_SERVER_ERROR, err)
	}
	return
}

// query user's contact list
func (cs *ContactService) GetContacts(userId string) (contacts []models.Contact, tErr *TError) {
	// TODO: should consider the black list function
	criteria := map[string]interface{}{
		"userId": userId,
	}
	contacts, err := contactRepo.FindAll(criteria)
	if err != nil {
		logger.Errorf("query user contacts error: %s", err.Error())
		return nil, NewTError(protos.StatusCode_STATUS_INTERNAL_SERVER_ERROR, err)
	}

	return contacts, nil
}

func (cs *ContactService) GetContactOperationMessageList(userId string, maxMessageTime int64) (contacts []models.Contact, tErr *TError) {
	// TODO: should check the regx pattern
	ctKey := fmt.Sprintf("^CT-.*%s*", userId)
	_, err := myRedis.RHGetAll(ctKey)
	if err != nil {
		logger.Errorf("query user contacts operation message list error: %s", err.Error())
		return nil, NewTError(protos.StatusCode_STATUS_INTERNAL_SERVER_ERROR, err)
	}



	return nil, nil
}

func handleAcceptContact(userId string, contactId string) (err error) {
	newContact1 := &models.Contact{
		UserId:    userId,
		ContactId: contactId,
	}
	newContact2 := &models.Contact{
		UserId:    contactId,
		ContactId: userId,
	}

	err = contactRepo.InsertContacts(newContact1, newContact2)
	return err
}

func handleDeleteContact(userId string, contactId string) (err error) {
	tx := mysqlDB.Begin()
	if err = contactRepo.RemoveContactsByIds(userId, contactId); err != nil {
		logger.Errorf("remove contact err: %s", err.Error())
		tx.Rollback()
		return
	}
	if err = contactRepo.RemoveContactsByIds(contactId, userId); err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return err
}

func handleUpdateContactRemark(userId string, contactId string, pbProfile *protos.ContactRemark) (err error) {
	criteria := map[string]interface{}{
		"UserId":    userId,
		"ContactId": contactId,
	}

	remarkProfile := converters.ConvertProto2EntityForRemarkProfile(pbProfile)
	updateMap := utils.TransformStructToMap(remarkProfile)

	err = contactRepo.FindOneAndUpdateRemark(criteria, updateMap)

	return
}

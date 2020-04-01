package contact

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	mysqlsrv "goim-pro/pkg/db/mysql"
	"testing"
)

func TestContact_InsertContacts(t *testing.T) {
	mysqlDB := mysqlsrv.NewMysqlConnection()
	_ = mysqlDB.Connect()
	NewContactRepo(mysqlsrv.NewMysqlConnection().GetMysqlInstance())

	newContact1 := &Contact{
		UserId:    "TEST001",
		ContactId: "TEST002",
	}
	newContact2 := &Contact{
		UserId:    "TEST003",
		ContactId: "TEST004",
	}

	ct := &Contact{}
	err := ct.InsertContacts(newContact1, newContact2)

	Convey("Test_InsertContacts", t, func() {
		Convey("should_insert_multiple_contacts", func() {
			ShouldBeNil(err)
		})
	})

	_ = ct.RemoveContactsByIds("TEST001", "TEST002")
	_ = ct.RemoveContactsByIds("TEST003", "TEST004")
}

func TestContact_RemoveContactsByIds(t *testing.T) {
	mysqlDB := mysqlsrv.NewMysqlConnection()
	_ = mysqlDB.Connect()
	NewContactRepo(mysqlsrv.NewMysqlConnection().GetMysqlInstance())

	newContact1 := &Contact{
		UserId:    "TEST001",
		ContactId: "TEST002",
	}
	newContact2 := &Contact{
		UserId:    "TEST001",
		ContactId: "TEST003",
	}
	newContact3 := &Contact{
		UserId:    "TEST003",
		ContactId: "TEST004",
	}

	ct := &Contact{}
	err := ct.InsertContacts(newContact1, newContact2, newContact3)
	if err != nil {
		fmt.Println(err)
	}
	Convey("Test_RemoveContacts", t, func() {
		Convey("should_remove_contacts_by_contactIds", func() {
			err := ct.RemoveContactsByIds("TEST001", "TEST002", "TEST003")
			ShouldBeNil(err)

			contact, err := ct.FindOne(&Contact{UserId: "TEST001"})
			ShouldBeNil(err)
			ShouldBeNil(contact)

			contact2, err := ct.FindOne(&Contact{UserId: "TEST003"})
			ShouldBeNil(err)
			So(contact2.UserId, ShouldEqual, "TEST003")

			_ = ct.RemoveContactsByIds("TEST003", "TEST004")

			err = ct.RemoveContactsByIds("TEST003", "TEST004")

			ShouldBeNil(err)
		})
	})

}

func TestContact_FindOneAndUpdateRemark(t *testing.T) {
	mysqlDB := mysqlsrv.NewMysqlConnection()
	_ = mysqlDB.Connect()
	NewContactRepo(mysqlsrv.NewMysqlConnection().GetMysqlInstance())

	newContact1 := &Contact{
		UserId:    "TEST001",
		ContactId: "TEST002",
	}

	ct := &Contact{}
	_ = ct.InsertContacts(newContact1)

	Convey("Test_FindOneAndUpdateRemark", t, func() {
		Convey("should_update_contact_remark_profile_successfully", func() {
			criteria := &Contact{}
			criteria.UserId = "TEST001"
			criteria.ContactId = "TEST002"

			remarkProfile := map[string]interface{}{
				"RemarkName":  "JAMES001",
				"Telephone":   "13631210001;13631210001",
				"Description": "Crazy boy..",
				"Tags":        "Friend;Boy",
			}
			err := ct.FindOneAndUpdateRemark(criteria, remarkProfile)
			ShouldBeNil(err)

			result, _ := ct.FindOne(criteria)
			So(result.RemarkName, ShouldEqual, "JAMES001")
			So(result.Telephone, ShouldEqual, "13631210001;13631210001")
			So(result.Description, ShouldEqual, "Crazy boy..")
			So(result.Tags, ShouldEqual, "Friend;Boy")
		})
	})

	_ = ct.RemoveContactsByIds("TEST001", "TEST002")
}
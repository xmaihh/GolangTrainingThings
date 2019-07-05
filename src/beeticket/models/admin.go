package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	AdminList map[string]*Admin
)

type Admin struct {
	ID        string `required:"true" description:"admin id"`
	Username  string `required:"true" description:"admin username"`
	Password  string `required:"true" description:"admin password"`
	Totalcode int64  `required:"true" description:"total code"`
	Exp       string `required:"true" description:"expired"`
	Userlimit int64  `required:"true" description:"userlimit"`
}

func AddAdmin(admin Admin) string {
	admin.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	AdminList[admin.ID] = &admin
	return admin.ID
}

func GetAdmin(uid string) (admin *Admin, err error) {
	if admin, ok := AdminList[uid]; ok {
		return admin, nil
	}
	return nil, errors.New("Admin not exists")
}

func GetAllAdmins() map[string]*Admin {
	return AdminList
}

func UpdateAdmin(uid string, ad *Admin) (a *Admin, err error) {
	if admin, ok := AdminList[uid]; ok {
		if ad.Username != "" {
			admin.Username = ad.Username

		}
		if ad.Password != "" {
			admin.Password = ad.Password
		}
		if ad.Totalcode != 0 {
			admin.Totalcode = ad.Totalcode
		}
		if ad.Exp != "" {
			admin.Exp = ad.Exp
		}
		if ad.Userlimit != -1 {
			admin.Userlimit = ad.Userlimit
		}
		return admin, nil
	}
	return nil, errors.New("Admin Not Exist")
}

func Logind(username, passwd string) bool {
	for _, u := range AdminList {
		if u.Username == username && u.Password == passwd {
			return true
		}
	}
	return false
}

func DeleteAdmin(uid string) {
	delete(AdminList, uid)
}

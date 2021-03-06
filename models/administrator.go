package models

import "github.com/sunyatsuntobee/server/logger"

// Administrator Model
type Administrator struct {
	ID       int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Name     string `xorm:"name VARCHAR(20) NOTNULL"`
	Password string `xorm:"password VARCHAR(50) NOTNULL"`
	Level    int    `xorm:"level INT NOTNULL"`
}

// NewAdministrator creates a new administrator
func NewAdministrator(name string, password string, level int) *Administrator {
	return &Administrator{
		Name:     name,
		Password: password,
		Level:    level,
	}
}

// AdministratorDataAccessObject provides database access for Model
// Administrator
type AdministratorDataAccessObject struct{}

// AdministratorDAO instance of AdministratorDataAccessObject
var AdministratorDAO *AdministratorDataAccessObject

// TableName returns table name
func (*AdministratorDataAccessObject) TableName() string {
	return "administrators"
}

// FindByUsername finds an administrator by its name
func (*AdministratorDataAccessObject) FindByUsername(
	username string) (Administrator, bool) {

	var admin Administrator
	has, err := orm.Table(AdministratorDAO.TableName()).
		Where("name=?", username).Get(&admin)
	logger.LogIfError(err)
	return admin, has
}

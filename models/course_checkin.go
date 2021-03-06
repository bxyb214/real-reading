package models

import (
	"time"

	"github.com/reechou/holmes"
)

type UserCourseCheckin struct {
	ID                   int64 `xorm:"pk autoincr" json:"id"`
	UserId               int64 `xorm:"not null default 0 int index" json:"userId"`
	CourseId             int64 `xorm:"not null default 0 int index" json:"courseId"`
	MonthCourseCatalogId int64 `xorm:"not null default 0 int" json:"monthCourseCatalogId"`
	CreatedAt            int64 `xorm:"not null default 0 int" json:"createdAt"`
	UpdatedAt            int64 `xorm:"not null default 0 int" json:"-"`
}

func CreateUserCourseCheckin(info *UserCourseCheckin) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create user course checkin error: %v", err)
		return err
	}
	holmes.Info("create user course checkin[%v] success.", info)

	return nil
}

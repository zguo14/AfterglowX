package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/wwkeyboard/sunsetwx/logs"
)

type RuleEntity struct {
	ID         int     `json:"id"`
	RuleID     string  `json:"rule_id"`
	UserID     string  `json:"user_id"`
	Coordinate string  `json:"coordinate"`
	Distance   float32 `json:"distance"`
	Quality    float32 `json:"quality"`
	Type       string  `json:"type"`
	Time       int64   `json:"time"`
	Frequency  int     `json:"frequency"`
	State      int     `json:"state"`
	CreatedOn  int64   `json:"created_on"`
	ModifiedOn int64   `json:"modified_on"`
}

func (rule *RuleEntity) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("CreatedOn", time.Now().Unix())
}

func (rule *RuleEntity) BeforeUpdate(scope *gorm.Scope) {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
}

func GetRuleEntities(data map[string]interface{}) (res []RuleEntity, err error) {

	err = db.Where("user_id = ?", data["user_id"].(string)).Find(&res).Error
	if err != nil {
		logs.Log.Error("GetRuleEntities error, error message: %s", err)
		return nil, err
	}

	return
}

func CreateRuleEntity(data RuleEntity) error {
	data.RuleID = uuid.NewString()

	err := db.Create(&data).Error
	if err != nil {
		logs.Log.Error("CreateRuleEntity error, error message: %s", err)
		return err
	}

	return nil
}

func UpdateRuleEntity(data RuleEntity) error {

	err := db.Model(&RuleEntity{}).Where("rule_id = ?", data.RuleID).Update(&data).Error

	if err != nil {
		logs.Log.Error("UpdateRuleEntity error, error message: %s", err)
		return err
	}

	return nil
}

func DeleteRuleEntity(data map[string]interface{}) error {

	err := db.Where("rule_id = ?", data["rule_id"].(string)).Delete(&RuleEntity{}).Error

	if err != nil {
		logs.Log.Error("DeleteRuleEntity error, error message: %s", err)
		return err
	}

	return nil
}

package model

import "github.com/jinzhu/gorm"

// Commit 提交
type Commit struct {
	gorm.Model
	UserID uint   `gorm:"index"` // 用户id
	Status int    // 0:pending 1:correct -1:wrong
	Code   string `gorm:"type:text(50000)"`
	Output string `gorm:"type:text(50000)"`
}

// CreateCommit 创建commit
func CreateCommit(id uint, code string) (Commit, error) {
	commit := Commit{
		UserID: id,
		Code:   code,
		Status: 0,
	}
	err := DB.Create(&commit).Error
	return commit, err
}

// SetCommit 设置commit状态
func SetCommit(commitID int, output string, status int) error {
	var commit Commit
	err := DB.Where("id = ?", commitID).First(&commit).Error
	if err != nil {
		return err
	}
	commit.Output = output
	commit.Status = status
	err = DB.Save(&commit).Error
	return err
}

// GetCommitByID 获取commit
func GetCommitByID(commitID int) (Commit, error) {
	var commit Commit
	err := DB.Where("id = ?", commitID).First(&commit).Error
	return commit, err
}

// GetCommitList 获取用户的commit列表
func GetCommitList(id uint) ([]Commit, error) {
	var commits []Commit
	err := DB.Where("user_id = ?", id).Find(&commits).Order("id DESC", true).Error
	return commits, err
}

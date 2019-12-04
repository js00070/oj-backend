package model

import "github.com/jinzhu/gorm"

// Commit 提交
type Commit struct {
	gorm.Model
	Status int    // 0:pending 1:correct -1:wrong
	Code   string `gorm:"size:32767"`
	Output string `gorm:"size:32767"`
}

// CreateCommit 创建commit
func CreateCommit(code string) (Commit, error) {
	commit := Commit{
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

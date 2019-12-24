package model

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Problem 题目
type Problem struct {
	gorm.Model
	Title      string `gorm:"varchar(64)"`
	Desciption string `gorm:"type:text(50000)"`
	StdInput   string `gorm:"type:text(50000)"`
	StdOutput  string `gorm:"type:text(50000)"`
}

// Commit 提交
type Commit struct {
	gorm.Model
	UserID    uint `gorm:"index"` // 用户id
	ProblemID uint `gorm:"index"` // 题目id
	Language  string
	Status    int    // 0:pending 1:correct -1:wrong
	Code      string `gorm:"type:text(50000)"`
	Output    string `gorm:"type:text(50000)"`
}

// CreateCommit 创建commit
func CreateCommit(id uint, pid uint, code string, lang string) (Commit, error) {
	commit := Commit{
		UserID:    id,
		Code:      code,
		ProblemID: pid,
		Language:  lang,
		Status:    0,
	}
	err := DB.Create(&commit).Error
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	var prob Problem
	err = DB.Where("id = ?", pid).First(&prob).Error
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	go Judge(&prob, &commit)
	return commit, err
}

// SetCommit 设置commit状态
func SetCommit(commitID uint, output string, status int) error {
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

// Judge 判题线程
func Judge(p *Problem, c *Commit) {
	time.Sleep(time.Second * 1)
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
			SetCommit(c.ID, "", -1)
		}
	}()
	cmd := exec.Command("python", "-c", c.Code)
	cmd.Stdin = strings.NewReader(p.StdInput)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		SetCommit(c.ID, "", -1)
		log.Fatal(err)
	}
	outStr := strings.Trim(strings.TrimSpace(out.String()), "\n")
	log.Println("output is ", outStr)
	if outStr == p.StdOutput {
		SetCommit(c.ID, outStr, 1)
	} else {
		SetCommit(c.ID, outStr, -1)
	}
}

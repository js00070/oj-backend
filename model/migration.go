package model

//执行数据迁移

func migration() {
	// 清空所有历史数据
	DB.DropTableIfExists(&User{}, &Commit{}, &Problem{})
	// 自动迁移模式
	DB.AutoMigrate(&User{}, &Commit{}, &Problem{})
	p1 := Problem{
		Title:      "计算a+a",
		Desciption: "计算a+a",
		StdInput:   "12",
		StdOutput:  "24",
	}
	DB.Create(&p1)
	p2 := Problem{
		Title:      "计算a*a",
		Desciption: "计算a*a",
		StdInput:   "12",
		StdOutput:  "144",
	}
	DB.Create(&p2)
}

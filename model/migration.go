package model

//执行数据迁移

func migration() {
	// 清空所有历史数据
	DB.DropTableIfExists(&User{}, &Commit{}, &Problem{})
	// 自动迁移模式
	DB.AutoMigrate(&User{}, &Commit{}, &Problem{})
	p1 := Problem{
		Title:      "计算a+a",
		Desciption: "输入一个整数a，请你计算并输出a+a的结果",
		StdInput:   "12",
		StdOutput:  "24",
	}
	DB.Create(&p1)
	p2 := Problem{
		Title:      "计算正整数的平方",
		Desciption: "输入一个整数a，请你输出a的平方的结果",
		StdInput:   "12",
		StdOutput:  "144",
	}
	DB.Create(&p2)
	DB.Create(&Problem{
		Title:      "计算正整数的阶乘",
		Desciption: "输入一个整数n，请你输出n的阶乘的结果",
		StdInput:   "8",
		StdOutput:  "40320",
	})
}

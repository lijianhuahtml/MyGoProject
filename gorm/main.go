package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 是用于映射数据库中的user表的模型
type User struct {
	gorm.Model
	Name string
	Age  int
}

// Equals 方法用于比较两个 User 对象是否相等
//func (u User) Equals(other User) bool {
//      return u.ID == other.ID && u.Name == other.Name
//}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/demo1_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// 自动迁移模型，确保表结构与模型对应
	err = db.AutoMigrate(&User{})
	if err != nil {
		return
	}

	// 创建用户
	newUser := User{Name: "Alice", Age: 25}
	db.Create(&newUser)

	// 查询用户
	var queriedUser User
	db.First(&queriedUser, 1) // 查询ID为1的用户
	fmt.Println("Queried User:", queriedUser)

	// 更新用户
	db.Model(&queriedUser).Update("Age", 26)

	// 删除用户
	db.Delete(&queriedUser)

	// 查询所有用户
	var allUsers []User
	db.Find(&allUsers)
	fmt.Println("All Users:", allUsers)
}

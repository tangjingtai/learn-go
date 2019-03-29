package main

import "fmt"

type Role struct {
	RoleId   int
	RoleName string
}

type User struct {
	Id       int
	Name     string
	UserName string
	Age      int
	Role     *Role
}

func ValueOrReferenceTransmit() {
	user1 := User{
		Id:       1,
		Name:     "张三",
		UserName: "zhangsan",
		Age:      20,
		Role: &Role{
			RoleId:   1,
			RoleName: "admin",
		},
	}
	user2 := &User{
		Id:       2,
		Name:     "李四",
		UserName: "lisi",
		Age:      30,
		Role:     user1.Role,
	}
	ValueTransmit(user1)
	ReferenceTransmit(user2)

	fmt.Printf("user1:%v, user1 type:%T, user2:%v, user2 type:%T\n", user1, user1, user2, user2)

	fmt.Printf("user1.Role:%v,  user2.Role:%v\n", user1.Role, user2.Role)
}

func ValueTransmit(user User) {
	user.Name = user.Name + "_update"
	user.Age = user.Age + 20
	user.Role.RoleName = user.Role.RoleName + "_update"
}

func ReferenceTransmit(user *User) {
	user.Name = user.Name + "_update"
	user.Age = user.Age + 20
	user.Role.RoleName = user.Role.RoleName + "_update"
}

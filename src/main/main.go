package main

import (
	"fmt"
	"otherpkg"
)

func main() {
	ad := otherpkg.Admin{
		Level: 1,
	}
	ad.UserName = "tangjingtai"
	ad.Email = "a253210810@qq.com"
	fmt.Printf("%v\n", ad)
}

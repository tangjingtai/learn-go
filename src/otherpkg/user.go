package otherpkg

type user struct {
	UserName string
	Email string
}

type Admin struct {
	user
	Level int
}

type Student struct {
	user
	CourseType int8
	Grade int8
}

type Teacher struct {
	user
	CourseType int8
}

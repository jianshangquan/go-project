package passbyref


type User struct{
	Name string
}


func modifyUserWithRef(user *User){
	user.Name = "John Doe updated"
}



func modifyUser(user User){
	user.Name = "John Doe updated"
}





func modifySlice(slice []int){
	slice[0] = 100;
}


func PassByRefTest(){
	var slice = []int{1, 2, 3}
	modifySlice(slice)
	println(slice[0]);


	user := User{Name: "Jane Doe"}
	modifyUserWithRef(&user)
	println(user.Name)

	user2 := User{Name: "Jane Doe2"}
	modifyUser(user2)
	println(user2.Name)
}
package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	// user 类型的值可以用来调用使用值接收者声明的方法
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	// 指向 user 类型值的指针也可以用来调用使用值接收者声明的方法
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	// user 类型的值可以用来调用使用指针接收者声明的方法
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// 指向 user 类型值的指针可以用来调用使用之神接收者声明的方法
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}

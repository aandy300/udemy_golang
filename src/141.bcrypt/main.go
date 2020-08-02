// 141. bcrypt 加密 和 確認密碼
// https://godoc.org/golang.org/x/crypto/bcrypt
package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	// cost  =  加密等級?
	// func GenerateFromPassword（password [] byte，cost int）（[] byte，error）
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPW := `password123`
	// 確認密碼 登入 loginPW = bs > login
	// func CompareHashAndPassword(hashedPassword, password []byte) error
	// CompareHashAndPassword將bcrypt hashed密碼與其可能的等效純文本進行比較。如果成功，則返回nil；如果失敗，則返回錯誤。
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPW))
	if err != nil {
		fmt.Println("u can't login")
		return
	}
	fmt.Println("u are login")
}

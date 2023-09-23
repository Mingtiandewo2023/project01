package main

import (
	"fmt" 
	"time" 
	"math/rand"
	"regexp"
)

var (
	phoneNumber string
	num1 int 
	num2 string
	num3 = fmt.Sprintf(sendVerifiCode(6))
	num4 int
	w int
)	

func checkMobile(phoneNumber string) bool{
	regRule:= "^1[345789]{1}\\d{9}$"
	reg := regexp.MustCompile(regRule)
	return reg.MatchString(phoneNumber)
}

func sendVerifiCode(n int) string{
	verifiCode := ""
	rand.Seed(time.Now().UnixNano())
	for j := 0 ; j < n ; j++{
		verifiCode = fmt.Sprintf("%s%d",verifiCode,rand.Intn(10))
	}
	return verifiCode
}

func main(){
	for i :=1 ; i <=3.14E10; i++{
		fmt.Print("请输入您的手机号：")
		fmt.Scanln(&phoneNumber)
		if checkMobile(phoneNumber) {
			for y := 1 ; y <=3.14E10 ; y++{
				fmt.Println("1:输入验证码进行登录 2：获取验证码")
				fmt.Scanln(&num1)
				if num1 == 2 {
					fmt.Print("您的验证码为:")
					fmt.Println(num3)
					break
				}else {
					fmt.Println("请输入正确的指令!")
				}
			}
			for x := 1 ; x <=3.14E10 ; x++{
				fmt.Println("1:输入验证码进行登录 2：获取验证码")
				fmt.Scanln(&num4)
				if num4 == 2 {
					fmt.Println("一分钟内请勿重复获取验证码！")
				}
				if num4 == 1 {
					for w= 4;w >=0;w--{
						fmt.Print("请输入验证码:")
						fmt.Scanln(&num2)
						if num2 == num3{
							fmt.Println("恭喜您，登入成功！")
							break		
						}else {
							fmt.Println("对不起，验证码错误！")
							fmt.Printf("您当前还可以输入%d次验证码",w)
							fmt.Println()
						}
					}
					break
				}
				if num4 != 1 && num4 != 2{
					fmt.Println("请输入正确的指令！")
				}
			}	
			break
		}else {
			fmt.Println("请输入正确的手机号！")
		}
	}
}	
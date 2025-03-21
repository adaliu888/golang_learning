package main

import (
	"fmt"
	"time"
)

func OnceTime() {
	// 1.timer基本使用
	//timer1 := time.NewTimer(2 * time.Second)
	//t1 := time.Now()
	//fmt.Printf("t1:%v\n", t1)
	//t2 := <-timer1.C
	//fmt.Printf("t2:%v\n", t2)

	// 2.验证timer只能响应1次
	//timer2 := time.NewTimer(time.Second)
	//for {
	// <-timer2.C
	// fmt.Println("时间到")
	//}

	// 3.timer实现延时的功能
	//(1)
	//time.Sleep(time.Second)
	//(2)
	//timer3 := time.NewTimer(2 * time.Second)
	//<-timer3.C
	//fmt.Println("2秒到")
	//(3)
	//<-time.After(2*time.Second)
	//fmt.Println("2秒到")

	// 4.停止定时器
	//timer4 := time.NewTimer(2 * time.Second)
	//go func() {
	// <-timer4.C
	// fmt.Println("定时器执行了")
	//}()
	//b := timer4.Stop()
	//if b {
	// fmt.Println("timer4已经关闭")
	//}

	// 5.重置定时器
	// 5. 创建并重置定时器
	timer5 := time.NewTimer(3 * time.Second)
	defer timer5.Stop() // 确保定时器在函数结束时停止

	// 重置定时器为1秒
	timer5.Reset(1 * time.Second)

	// 打印当前时间
	fmt.Println("当前时间:", time.Now())

	// 等待定时器到期
	select {
	case <-timer5.C:
		fmt.Println("定时器到期:", time.Now())
	}

	// 防止程序立即退出
	for {
	}
}

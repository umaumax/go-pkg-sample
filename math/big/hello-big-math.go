package main

import (
	"fmt"
	"math/big"
)

//	参考URL
//	http://www.geocities.jp/m_hiroi/golang/yagp03.html#ans27

func main() {
	{
		fmt.Println("----big.Rat----")
		fmt.Println("RAT	rational number, 有理数")
		r := big.NewRat(100, 30)
		fmt.Println("100 / 30 = ", r)
		fmt.Println("分子(num)", r.Num())
		fmt.Println("分母(denom)", r.Denom())
		fmt.Println()
	}
	{
		fmt.Println("----big.Rat----")
		a := &big.Rat{}
		b := big.NewRat(100, 200)
		c := big.NewRat(100, 400)
		fmt.Printf("a:%v, b:%v, c:%v\n", a, b, c)
		fmt.Printf("	a.Add(b, c)\n")
		//	a = b + c
		a.Add(b, c)
		fmt.Printf("a:%v, b:%v, c:%v\n", a, b, c)
		fmt.Printf("	a = a.Add(b, c)\n")
		//	返り値はメソッドのインスタンスと同一の参照先
		d := a.Add(b, c)
		if a == d {
			fmt.Println("same!")
		}
		//	ためしにdにsetするとaの値も変化
		d.Set(big.NewRat(1, 8))
		fmt.Printf("a:%v, b:%v, c:%v\n", a, b, c)
		fmt.Println()
	}
	{
		fmt.Println("----big.Int----")
		a := big.NewInt(10)
		b := big.NewInt(20)
		c := big.NewInt(30)
		a = b.Add(b, a)
		d := c
		d.Set(big.NewInt(200))
		fmt.Printf("a:%v, b:%v, c:%v, d:%v\n", a, b, c, d)
		fmt.Println()
	}

	fmt.Println("----power----")
	fmt.Println("3 ^ 100 = ", power(3, 100))
	fmt.Println("7 ^ 100 = ", power(7, 100))
	fmt.Println("----fact----")
	fmt.Println("300! = ", fact(300))
	fmt.Println("700! = ", fact(700))
	fmt.Println("----fibo----")
	fmt.Println("fibo(300) = ", fibo(300))
	fmt.Println("fibo(700) = ", fibo(700))
	fmt.Println("----cacheFibo----")
	fmt.Println("fibo(300) = ", cacheFibo(200000))
	fmt.Println("fibo(300) = ", cacheFibo(210000))
	fmt.Println("fibo(300) = ", cacheFibo(220000))
}

var (
	fiboMap = make([]big.Int, 0x100000)
	fiboLen = int64(2)
)

func init() {
	fiboMap[0] = *big.NewInt(0)
	fiboMap[1] = *big.NewInt(1)
	fiboMap[2] = *big.NewInt(1)
}

func cacheFibo(n int64) *big.Int {
	//	chache
	if n <= fiboLen {
		tmp := fiboMap[n]
		return &tmp
	}
	n -= fiboLen

	c := big.NewInt(0)
	btmp := fiboMap[fiboLen]
	//	big.Intの構造体の中ではスライスを参照しているので単純に[=]を用いてコピーするのではなく以下の様な方法で新規作成する必要性がある。
	b := (&big.Int{}).Set(&btmp)
	atmp := fiboMap[fiboLen-1]
	a := (&big.Int{}).Set(&atmp)
	for i := int64(1); i <= n; i++ {
		c.Set(a)
		a.Add(a, b)
		fiboMap[fiboLen+i] = *(&big.Int{}).Set(a)
		b.Set(c)
	}
	fiboLen += n
	return a
}

// 繰り返し
func fibo(n int64) *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)
	c := big.NewInt(0)
	for ; n > 0; n-- {
		c.Set(a)
		a.Add(a, b)
		b.Set(c)
	}
	return a
}

// 階乗
func fact(n int64) *big.Int {
	a := big.NewInt(1)
	for ; n > 0; n-- {
		a.Mul(a, big.NewInt(n))
	}
	return a
}

// 累乗
func power(x, y int64) *big.Int {
	switch {
	case y == 0:
		return big.NewInt(1)
	case y == 1:
		return big.NewInt(x)
	default:
		z := power(x, y/2)
		z.Mul(z, z) // z *= z
		if y%2 == 0 {
			return z
		} else {
			return z.Mul(z, big.NewInt(x))
		}
	}
}

func bigunm() {
	verybig := big.NewInt(1)
	ten := big.NewInt(10)
	for i := 0; i < 100000; i++ {
		temp := new(big.Int)
		temp.Mul(verybig, ten)
		verybig = temp
	}
	fmt.Println(verybig)
}

// Code generated by "stringer -type=Fruit"; DO NOT EDIT

package main

import "fmt"

const _Fruit_name = "AppleOrangeBanana"

var _Fruit_index = [...]uint8{0, 5, 11, 17}

func (i Fruit) String() string {
	if i < 0 || i >= Fruit(len(_Fruit_index)-1) {
		return fmt.Sprintf("Fruit(%d)", i)
	}
	return _Fruit_name[_Fruit_index[i]:_Fruit_index[i+1]]
}

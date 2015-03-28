package main

import (
	"fmt"
	"github.com/reidiiius/zaiyu/zhu"
)

func main() {
	var f float64

	zhu.Zaiyu[0] = 1

	for i := 0; i < 425; i++ {
		f += 1
		zhu.Zaiyu = append(zhu.Zaiyu, zhu.ChingFang(f))
	}

	for i := 0; i < 425; i++ {
		fmt.Println(zhu.Zaiyu[i])
	}
}

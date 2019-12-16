package main

import (
	"fmt"
	"github.com/mingzhehao/goutils/systool"
	"github.com/mingzhehao/goutils/timetool"
	"time"
)

func main() {
	fmt.Println(systool.IntranetIP())
	fmt.Println(timetool.DateFormat(time.Now(), "YYYY-MM-DD"))
}

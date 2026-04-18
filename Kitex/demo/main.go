package main

import (
	calc "demo/kitex_gen/demo/calc/calcservice"
	"log"
)

func main() {
	svr := calc.NewServer(new(CalcServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

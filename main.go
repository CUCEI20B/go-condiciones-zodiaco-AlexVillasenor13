package main

import "fmt"

func main()  {
	var dia int
	var mes int
	fmt.Scan(&dia)
	fmt.Scan(&mes)
	fecha := (mes * 100) + dia
	switch {
	case dia <=0 || mes <=0 || dia>31:
		fmt.Println("")
	case fecha >= 321 && fecha < 421:
		fmt.Println("aries")
	case fecha >= 421 && fecha < 521:
		fmt.Println("tauro")
	case fecha >= 521 && fecha < 622:
		fmt.Println("géminis")
	case fecha >= 622 && fecha < 723:
		fmt.Println("cáncer")
	
	case fecha >= 723 && fecha < 822:
		fmt.Println("leo")
	
	case fecha >= 822 && fecha < 923:
		fmt.Println("virgo")
	case fecha >= 923 && fecha < 1023:
		fmt.Println("libra")
	case fecha >= 1023 && fecha < 1123:
		fmt.Println("escorpio")
	case fecha >= 1123 && fecha < 1222:
		fmt.Println("sagitario")

	case (fecha >= 1222 && fecha < 1232) || (fecha >= 101 && fecha < 121):
		fmt.Println("capricornio")

	case fecha >= 121 && fecha < 219:
		fmt.Println("acuario")
	case fecha >= 219 && fecha < 321:
		fmt.Println("piscis")
	default:
		fmt.Println("")
	}
}
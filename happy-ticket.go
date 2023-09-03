package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 5
	fullLen := n * 2
	cnt := 0

	//сгенерированное максимальное число для генератора
	intNGen := builderNForGen(fullLen)

	for i := 0; i < intNGen; i++ {
		//fmt.Println("i", i)
		s := builder(i, fullLen)

		if checkIfIsLucky(fullLen, s) {
			cnt += 1
			fmt.Println(cnt)
		}
	}

	fmt.Println("cnt happy: ", cnt)
}

func builderNForGen(fullLen int) int {
	var sb strings.Builder
	sb.Write([]byte("1"))
	for i := 0; i < fullLen; i++ {
		sb.Write([]byte("0"))
	}
	intNGen, _ := strconv.Atoi(sb.String())
	//if err != nil {
	//	log.Fatal("builderNForGen error: ", err)
	//}

	return intNGen
}

func builder(i, n int) string {
	var sb strings.Builder

	strI := strconv.Itoa(i)

	//fmt.Println("strI: ", strI)
	//fmt.Println("len strI: ", len(strI))

	//если строка меньше N то восполняем нулями спереди
	if len(strI) < n {
		for j := 0; j < n-len(strI); j++ {
			sb.Write([]byte("0"))
		}
	}

	//fmt.Println("sb.String(): ", sb.String())
	s := sb.String() + strI
	//fmt.Println("s:", s)

	return s
}

// проверка билета на счастливый номер
func checkIfIsLucky(fullLen int, ticket string) bool {

	leftSum := 0
	rightSum := 0

	for i := 0; i < fullLen; i++ {
		if i < fullLen/2 {
			leftInt, _ := strconv.Atoi(string(ticket[i]))
			//if err != nil {
			//	log.Fatal("checkIfIsLucky leftInt err: ", err)
			//}
			leftSum += leftInt
		} else {
			rightInt, _ := strconv.Atoi(string(ticket[i]))
			//if err != nil {
			//	log.Fatal("checkIfIsLucky rightInt err: ", err)
			//}
			rightSum += rightInt
		}
	}

	if leftSum == rightSum {
		return true
	} else {
		return false
	}
}

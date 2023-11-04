// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

//Тут проблема в том что v[:100] создает новую подстроку, которая ссылается на тот же массив. что и исходная, от чего исходная строка v продолжит хранить все данные
//Таким образом, она не будет освобождена, и большое количество памяти будет занято неиспользуемыми данными, что может привести к проблемам с использованием памяти.

package main

import (
	"bytes"
	"fmt"
	"math/rand"
)

func createHugeString(size int) string {
	// Создаем строку заданного размера и заполняем ее данными.
	data := make([]byte, size)
	for i := range data {
		data[i] = byte('A' + rand.Intn(26)) // Заполняем строку случайными буквами A-Z
	}
	return string(data)
}

func someFunc(s *string) {
	v := createHugeString(1 << 10)
	buffer := bytes.NewBufferString(v)
	*s = string(buffer.Next(100))
}

func main() {
	var justString string
	someFunc(&justString)
	fmt.Println(justString)
}

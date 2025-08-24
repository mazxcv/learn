package designnumbers

import "slices"

// Разработать контейнер для хранения числа.
// методы:
// 1. find(n int) - поиск числа в контейнере, -1, если не найден
// 2. change(i int, n int)  - установить по индексу i новое число

// На вход подается массив указаний и массив действий
// ["NumberContainers","find","change","change","change","change","find","change","find"]
// [[],[10],[2,10],[1,10],[3,10],[5,10],[10],[1,20],[10]]
// 1. Иницииализация контейнера жзначением - операция возвращает Null
// 2. Поиск числа - в контейнере пусто - операция возвращает -1
// 3. Установка числа по индексу - операция возвращает Null теперь по индексу 2 лежит 10
// 4. Установка числа по индексу - операция возвращает Null теперь по индексу 1 лежит 10
// 5. Установка числа по индексу - операция возвращает Null теперь по индексу 3 лежит 10
// 6. Установка числа по индексу - операция возвращает Null теперь по индексу 5 лежит 10
// 7. Поиск числа 10 - в контейнере под индексом 1, 2, 3, 5 лежит 10 - ищем минимальный индекс и возвращаем 1
// 8. Установка числа по индексу - операция возвращает Null теперь по индексу 1 лежит 20
// 9. Поиск числа 10- в контейнере под индексом  2, 3, 5 лежит 10 - ищем минимальный индекс и возвращаем 2
// результат [null,-1,null,null,null,null,1,null,2]

// Идея: map[int][]int - здесь ключ - значение, а значение - упорядоченные индексы
// Для реализации change(i int, n int)  - установить по индексу i новое число потребуется ввести еще одну структуру для хранения значения текущего индекса

type NumberContainers struct {
	// map[int][]int для хранения упорядоченных индексов у значения
	res map[int][]int
	// map[int]int для хранения значений индексов
	ind map[int]int
	// отсортированы ли индексы
	isSorted map[int]bool
}

func Constructor() NumberContainers {
	return NumberContainers{
		res:      make(map[int][]int),
		ind:      make(map[int]int),
		isSorted: make(map[int]bool),
	}
}

func (nc *NumberContainers) Change(index int, number int) {
	oldNumber, exists := nc.ind[index]
	if oldNumber == number {
		return
	}
	if exists {
		// remove old number
		for i, current := range nc.res[oldNumber] {
			if current == index {
				nc.res[oldNumber] = append(nc.res[oldNumber][:i], nc.res[oldNumber][i+1:]...)
				break
			}
		}
	}
	nc.ind[index] = number
	nc.res[number] = append(nc.res[number], index)
	nc.isSorted[number] = false
}

func (nc *NumberContainers) Find(number int) int {
	if !nc.isSorted[number] {
		slices.Sort(nc.res[number])
		nc.isSorted[number] = true
	}
	numberIndices, ok := nc.res[number]

	if !ok || len(numberIndices) == 0 {
		return -1
	}
	return numberIndices[0]
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
func designnumbers(command []string, values [][]int) []int {
	nc := Constructor()
	result := make([]int, len(command))
	for i := 1; i < len(command); i++ {
		switch command[i] {
		case "change":
			nc.Change(values[i][0], values[i][1])
		case "find":
			res := nc.Find(values[i][0])
			result[i] = res
		}
	}
	return result
}

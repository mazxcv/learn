package findsumpairs

// Дано  два целочисленных массива nums1и nums2. Реализовать структуру данных, которая поддерживает запросы двух типов:
//  Добавить положительное целое число к элементу заданного индекса в массиве nums2.
// void add(int index, int val) применяет nums2[index] += val.
// Подсчитайть количество пар (i, j), для которых nums1[i] + nums2[j]равны заданному значению ( 0 <= i < nums1.lengthи 0 <= j < nums2.length).
// int count(int tot)Возвращает количество пар, (i, j)таких что nums1[i] + nums2[j] == tot.
// FindSumPairs(int[] nums1, int[] nums2)Инициализирует FindSumPairs объект двумя целочисленными массивами nums1и nums2.

// Главня идея - пока не хранить в summary суммы, а просто сохранять количество цифр
// Когда захочется поискать сумму - надо будет пройтись по первому массиву и поискать в summary разность
// Просуммировать результаты из summary и вернуть

type FindSumPairs struct {
	num1 []int
	num2 []int

	summary map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	fsp := FindSumPairs{
		num1:    nums1,
		num2:    nums2,
		summary: make(map[int]int, len(nums2)),
	}

	for _, num := range nums2 {
		fsp.summary[num]++
	}

	return fsp
}

func (fsp *FindSumPairs) Add(index, value int) {
	originalValue := fsp.num2[index]
	newValue := originalValue + value

	fsp.num2[index] = newValue
	fsp.summary[originalValue]--
	fsp.summary[newValue]++
}

func (fsp *FindSumPairs) Count(target int) int {
	count := 0
	for _, num := range fsp.num1 {
		count += fsp.summary[target-num]
	}
	return count
}

func findSumPairs(init [][]int, commands []string, values [][]int) []int {
	fsp := Constructor(init[0], init[1])
	result := make([]int, len(commands)+1)

	for i, command := range commands {
		switch command {
		case "add":
			fsp.Add(values[i][0], values[i][1])
		case "count":
			result[i+1] = fsp.Count(values[i][0])
		}
	}

	return result
}

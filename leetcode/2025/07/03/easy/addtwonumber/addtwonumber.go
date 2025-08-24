package addtwonumber

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  Основная идея: пройтись по связному списку и дополнительно хранить выход метку выхода суммы за 10, чтобы увеличить сумму следущих разрядов. Применить подход итеративно ко всем узлам списков.
// Если один узел закончился, а второй еще нет, то считаем, что первый узел содержит нули
// если узлы кончились у обоих списков, но метка переключатель регистра содержит true, то добавить еще один узел в сумму и назначить его единицей
func addTwoNumbers(l *ListNode, r *ListNode) *ListNode {
	res := &ListNode{}
	head := res
	registerUp := 0

	for l != nil || r != nil || registerUp == 1 {
		res.Val = registerUp
		if l != nil {
			res.Val += l.Val
			l = l.Next
		}
		if r != nil {
			res.Val += r.Val
			r = r.Next
		}
		registerUp = res.Val / 10
		res.Val %= 10

		if l != nil || r != nil || registerUp == 1 {
			res.Next = &ListNode{}
			res = res.Next
		}
	}

	return head
}

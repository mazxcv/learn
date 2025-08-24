package _672__Richest_Customer_Wealth

func maximumWealth(accounts [][]int) int {
	m := 0
	for i := 0; i < len(accounts); i++ {
		sumJ := 0
		for j := 0; j < len(accounts[i]); j++ {
			sumJ += accounts[i][j]
		}
		if m < sumJ {
			m = sumJ
		}
	}

	return m
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	r1 := head
	r2 := head
	for r1 != nil && r2 != nil {
		r1 = r1.Next
		r2 = r2.Next
		if r2 != nil {
			r2 = r2.Next
		}
	}
	return r1
}

package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unit := make(map[string]int)
	unit["quarter_of_a_dozen"] = 3
	unit["half_of_a_dozen"] = 6
	unit["dozen"] = 12
	unit["small_gross"] = 120
	unit["gross"] = 144
	unit["great_gross"] = 1728
	return unit
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	dozen, exists := units[unit]
	if !exists {
		return exists
	} else {
		bill[item] += dozen
		return true
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, exists := bill[item]
	if !exists {
		return exists
	}
	dozen, exists := units[unit]
	if !exists {
		return exists
	}
	if bill[item] < dozen {
		return false
	}
	if bill[item] == dozen {
		delete(bill, item)
		return true
	}
	bill[item] -= dozen
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	billDozen, exists := bill[item]
	if !exists {
		return 0, exists
	} else {
		return billDozen, exists
	}
}

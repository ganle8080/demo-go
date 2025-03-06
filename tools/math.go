package tools

// IfTree 三目运算
func IfTree[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

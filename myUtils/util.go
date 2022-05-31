package myUtils

func ToInterfaceArray(vals ...interface{}) []interface{} {
	sliceI := make([]interface{}, len(vals))
	copy(sliceI, vals)
	return sliceI
}

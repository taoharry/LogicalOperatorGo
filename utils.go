/**
 * @author Co2
 * @file_name utils.go
 */

package LogicalOperatorGo

func RemoveKong(strs []string) []string {
	result := []string{}
	for _, k := range strs {
		if k != "" && k != LuKong {
			result = append(result, k)
		}
	}
	return result
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

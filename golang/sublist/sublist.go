package sublist

type Relation = string

func checkSubList(ll1 int, ll2 int, l1 []int, l2 []int) bool {
	for i := 0; i < ll2-ll1+1; i++ {
		isSubList := true
		for j := i; j < i+ll1; j++ {
			if l1[j-i] != l2[j] {
				isSubList = false
				continue
			}
		}
		if isSubList {
			return isSubList
		}
	}
	return false
}

func Sublist(l1, l2 []int) string {
	ll1, ll2 := len(l1), len(l2)
	if ll1 == ll2 && ll1 == 0 {
		return "equal"
	}
	if ll1 != ll2 {
		if ll1 < ll2 && checkSubList(ll1, ll2, l1, l2) {
			return "sublist"
		} else if ll1 > ll2 && checkSubList(ll2, ll1, l2, l1) {
			return "superlist"
		}
		return "unequal"
	} else {
		if checkSubList(ll1, ll2, l1, l2) {
			return "equal"
		} else {
			return "unequal"
		}
	}

}

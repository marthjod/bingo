// Code generated by "stringer -type=Case"; DO NOT EDIT.

package wordtype

import "fmt"

const _Case_name = "NefnifallÞolfallÞágufallEignarfall"

var _Case_index = [...]uint8{0, 9, 17, 27, 37}

func (i Case) String() string {
	if i < 0 || i >= Case(len(_Case_index)-1) {
		return fmt.Sprintf("Case(%d)", i)
	}
	return _Case_name[_Case_index[i]:_Case_index[i+1]]
}

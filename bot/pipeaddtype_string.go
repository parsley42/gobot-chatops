// Code generated by "stringer -type=pipeAddType constants.go"; DO NOT EDIT.

package bot

import "strconv"

const _pipeAddType_name = "typeTasktypePlugintypeJob"

var _pipeAddType_index = [...]uint8{0, 8, 18, 25}

func (i pipeAddType) String() string {
	if i < 0 || i >= pipeAddType(len(_pipeAddType_index)-1) {
		return "pipeAddType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _pipeAddType_name[_pipeAddType_index[i]:_pipeAddType_index[i+1]]
}
// Code generated by "stringer -type=Protocol"; DO NOT EDIT.

package bot

import "strconv"

const _Protocol_name = "SlackTerminalTest"

var _Protocol_index = [...]uint8{0, 5, 13, 17}

func (i Protocol) String() string {
	if i < 0 || i >= Protocol(len(_Protocol_index)-1) {
		return "Protocol(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Protocol_name[_Protocol_index[i]:_Protocol_index[i+1]]
}

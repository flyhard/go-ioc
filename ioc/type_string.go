// Code generated by "stringer -type=Type"; DO NOT EDIT.

package ioc

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Bitcoin-0]
	_ = x[MD5-1]
	_ = x[SHA1-2]
	_ = x[SHA256-3]
	_ = x[SHA512-4]
	_ = x[Domain-5]
	_ = x[Email-6]
	_ = x[IPv4-7]
	_ = x[IPv6-8]
	_ = x[URL-9]
	_ = x[File-10]
	_ = x[CVE-11]
}

const _Type_name = "BitcoinMD5SHA1SHA256SHA512DomainEmailIPv4IPv6URLFileCVE"

var _Type_index = [...]uint8{0, 7, 10, 14, 20, 26, 32, 37, 41, 45, 48, 52, 55}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
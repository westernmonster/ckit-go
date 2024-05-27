package ecode

var (
	Nil                = localCode{-1, "", nil}
	OK                 = localCode{200, "OK", nil}
	MethodNoPermission = localCode{4, "Method has no permission", nil}
	RequestErr         = localCode{400, "Invalid Request", nil}
	NothingFound       = localCode{404, "Nothing Found", nil}
	ServerErr          = localCode{500, "Internal Server Error", nil}
)

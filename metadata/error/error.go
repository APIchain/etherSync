package error

const (
	SUCCESS            int = 0
	SESSION_EXPIRED    int = 41001
	SERVICE_CEILING    int = 41002
	ILLEGAL_DATAFORMAT int = 41003
	OAUTH_TIMEOUT      int = 41004
	DUPLICATE_ORDER    int = 41005
	EXPIRE_ORDER       int = 41006
	ORDER_NOT_FOUND    int = 41007
	NRF                int = 41008

	INVALID_METHOD int = 42001
	INVALID_PARAMS int = 42002
	INVALID_TOKEN  int = 42003

	INVALID_VERSION int = 45001
	INTERNAL_ERROR  int = 45002
	TEST_HELLO      int = 99999
)

var ErrMap = map[int]string{
	SUCCESS:            "SUCCESS",
	NRF:                "NO RECORED FOUND",
	SESSION_EXPIRED:    "SESSION EXPIRED",
	SERVICE_CEILING:    "SERVICE CEILING",
	ILLEGAL_DATAFORMAT: "ILLEGAL DATAFORMAT",
	OAUTH_TIMEOUT:      "CONNECT TO OAUTH TIMEOUT",

	INVALID_METHOD: "INVALID METHOD",
	INVALID_PARAMS: "INVALID PARAMS",
	INVALID_TOKEN:  "VERIFY TOKEN ERROR",

	INVALID_VERSION: "INVALID VERSION",
	INTERNAL_ERROR:  "INTERNAL ERROR",
	DUPLICATE_ORDER: "DUPLICATE ERROR",
	EXPIRE_ORDER:    "EXPIRE_ORDER",
	ORDER_NOT_FOUND: "ORDER_NOT_FOUND",
	TEST_HELLO:      "TEST_HELLO",
}
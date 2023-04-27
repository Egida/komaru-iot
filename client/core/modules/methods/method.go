package methods

var (
	methods = make([]*Method, 0)
)

var (
	FlagSize      = 0
	FlagUserAgent = 1 // This disables random useragent
	FlagMethod    = 2
	FlagPayload   = 3 // This will overwrite the random payload
	FlagThreads   = 4 // This will overwrite the default which is 1
)

type Method struct {
	Id       int
	Executor func(targets []string, port int, duration int, flags map[int]interface{}) error
}

func Make(method *Method) {
	methods = append(methods, method)
}

func Clone() []*Method {
	return methods
}

func Get(id int) *Method {
	for _, method := range methods {
		if method.Id == id {
			return method
		}
	}
	return nil
}

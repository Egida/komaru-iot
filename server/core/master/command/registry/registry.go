package registry

import (
	"server/core/master/command/admin"
	"server/core/master/command/all"
)

func Init() {
	admin.Register()
	all.Register()
}

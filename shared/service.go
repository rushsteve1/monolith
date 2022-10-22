package shared

import (
	"fmt"

	"github.com/thejerf/suture/v4"
)

type Service interface {
	suture.Service
	fmt.Stringer
	Addr() string
	Name() string
	UseFcgi() bool
}

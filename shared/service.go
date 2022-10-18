package shared

type Service interface {
	Addr() string
	Name() string
	UseFcgi() bool
}

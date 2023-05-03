package interfaces

type Repository interface {
}

type Repositories interface {
	Repository
}

type Service interface {
}

type Services interface {
	Service
}

type Controller interface {
}

type Controllers interface {
	Controller
}

type Router interface {
	RegisterRoutes()
}

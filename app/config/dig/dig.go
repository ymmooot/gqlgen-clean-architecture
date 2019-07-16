package dig

import (
	"github.com/pkg/errors"
	"github.com/ymmooot/gqlgen-clean-architecture/app/infrastructure"
	"github.com/ymmooot/gqlgen-clean-architecture/app/interfaceadapter/gateway"
	"github.com/ymmooot/gqlgen-clean-architecture/app/interfaceadapter/resolver"
	"github.com/ymmooot/gqlgen-clean-architecture/app/usecase"
	"go.uber.org/dig"
)

func BuildDigDependencies() (*dig.Container, error) {
	c := dig.New()

	registerDependencies(
		c,
		infrastructure.NewSqlHandler,
		gateway.NewArticleGateway,
		usecase.NewArticleUseCaseImpl,
		resolver.NewResolver,
	)

	return c, nil
}

func registerDependencies(c *dig.Container, constructors ...interface{}) {
	for i := 0; i < len(constructors); i++ {
		err := c.Provide(constructors[i])
		if err != nil {
			panic(errors.Wrapf(err, "On #%v/%v elms:", i, len(constructors)))
		}
	}
}

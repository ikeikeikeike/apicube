package trans

import (
	"github.com/facebookgo/inject"

	"github.com/ikeikeikeike/apicube/base/util"
	"github.com/ikeikeikeike/apicube/base/util/logger"
)

// Inject injects dependencies
func Inject(env util.Environment, g *inject.Graph, rt interface{}) {
	// inject
	err := g.Provide(
		&inject.Object{Value: &products{}},
	)
	if err != nil {
		logger.Panicf("failed to process injection: %s", err)
	}

}

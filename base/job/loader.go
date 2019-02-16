package job

import (
	"github.com/facebookgo/inject"
	"github.com/ikeikeikeike/apicube/base/job/loades"
	"github.com/ikeikeikeike/apicube/base/util"
)

// Inject injects dependencies
func Inject(env util.Environment, g *inject.Graph, rt interface{}) {
	loades.Inject(env, g, rt)
}

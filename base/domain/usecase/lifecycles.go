package usecase

import (
	"database/sql"
	"fmt"

	"github.com/ikeikeikeike/apicube/base/util"
	pb "github.com/ikeikeikeike/apicube/rpc/protocol/apicube/lifecycle"
)

type (
	// Lifecycles manifests ...
	Lifecycles interface {
		Ping(*pb.Ping) (*pb.Ping, error)
	}

	lifecycles struct {
		Env util.Environment `inject:""`
		DB  *sql.DB          `inject:""`
	}
)

// Ping returns pb message
//
func (p *lifecycles) Ping(req *pb.Ping) (*pb.Ping, error) {
	// make sure db connection available
	if err := p.DB.Ping(); err != nil {
		return nil, fmt.Errorf("it was unable to connect the db: %s", err)
	}

	return &pb.Ping{Msg: fmt.Sprintf("pong - %v", req.Msg)}, nil

}

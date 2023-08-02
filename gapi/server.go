package gapi

import (
	"fmt"

	db "github.com/hingebingedurgen/simplebank/db/sqlc"
	"github.com/hingebingedurgen/simplebank/pb"
	"github.com/hingebingedurgen/simplebank/token"
	"github.com/hingebingedurgen/simplebank/util"
)

// Server serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server and setup routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create a token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

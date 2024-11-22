package main

import (
	"git.itzana.me/strafesnet/data-service/internal/controller"
	"git.itzana.me/strafesnet/data-service/internal/datastore"
	"git.itzana.me/strafesnet/go-grpc/bots"
	"git.itzana.me/strafesnet/go-grpc/events"
	"git.itzana.me/strafesnet/go-grpc/maps"
	"git.itzana.me/strafesnet/go-grpc/ranks"
	"git.itzana.me/strafesnet/go-grpc/servers"
	"git.itzana.me/strafesnet/go-grpc/times"
	"git.itzana.me/strafesnet/go-grpc/transactions"
	"git.itzana.me/strafesnet/go-grpc/users"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
)

func rpc(db datastore.Datastore, debug bool) {
	// Open gRPC Listener
	port := ":9000"
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.WithField("error", err).Fatalln("failed to net.Listen")
	}

	// Spawn grpc server instance
	grpcServer := grpc.NewServer()

	// Service Registrations
	bots.RegisterBotsServiceServer(grpcServer, &controller.Bots{Store: db})
	events.RegisterEventsServiceServer(grpcServer, &controller.Event{Store: db})
	maps.RegisterMapsServiceServer(grpcServer, &controller.Maps{Store: db})
	servers.RegisterServerServiceServer(grpcServer, &controller.Servers{Store: db})
	times.RegisterTimesServiceServer(grpcServer, &controller.Times{Store: db})
	users.RegisterUsersServiceServer(grpcServer, &controller.Users{Store: db})
	transactions.RegisterTransactionsServiceServer(grpcServer, &controller.Transactions{Store: db})
	ranks.RegisterRanksServiceServer(grpcServer, &controller.Ranks{Store: db})

	if err := grpcServer.Serve(lis); err != nil {
		log.WithField("error", err).Fatalf("failed to serve")
	}
}

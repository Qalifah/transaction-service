package handler

import (
	"github.com/Qalifah/grey-challenge/transaction/database/postgres"
	"github.com/Qalifah/grey-challenge/transaction/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"net"

	"context"
	"testing"

	log "github.com/sirupsen/logrus"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	ctx := context.Background()
	conn, err := postgres.CreateTestDB(ctx)
	if err != nil {
		log.Fatalf("Unable to create test database %v", err)
	}
	transferRepo := postgres.NewTransferRepository(conn)
	proto.RegisterTransactionServer(s, New(transferRepo))
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestHandler_CreditAccount(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer conn.Close()

	tsClient := proto.NewTransactionClient(conn)
	resp, err := tsClient.CreditAccount(ctx, &proto.Transfer{To: 1, From: 7, Amount: 100})
	require.NoError(t, err)
	require.Equal(t, true, resp.IsSuccessful)
}

func TestHandler_DebitAccount(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer conn.Close()

	tsClient := proto.NewTransactionClient(conn)
	resp, err := tsClient.DebitAccount(ctx, &proto.Transfer{To: 1, From: 7, Amount: 100})
	require.NoError(t, err)
	require.Equal(t, true, resp.IsSuccessful)
}

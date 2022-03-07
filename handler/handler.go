package handler

import (
	"context"
	core "github.com/Qalifah/grey-challenge/transaction"
	"github.com/Qalifah/grey-challenge/transaction/proto"
	walletpb "github.com/Qalifah/grey-challenge/wallet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type handler struct {
	transferRepo core.TransferRepository
	proto.UnimplementedTransactionServer
}

var walletAddr = "127.0.0.1:50051"

func New(transferRepo core.TransferRepository) *handler {
	return &handler{transferRepo: transferRepo}
}

func (h *handler) CreditAccount(ctx context.Context, transfer *proto.Transfer) (*proto.TransferResponse, error) {
	tfs := unMarshaTransfer(transfer)
	tfs.Type = "credit"
	err := h.transferRepo.Add(ctx, tfs)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.DialContext(ctx, walletAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	walletClient := walletpb.NewWalletClient(conn)
	resp, err := walletClient.GetBalance(ctx, &walletpb.GetBalanceRequest{UserId: transfer.From})
	if err != nil {
		return nil, err
	}
	newAmount := resp.Amount - transfer.Amount
	_, err = walletClient.UpdateBalance(ctx, &walletpb.UpdateBalanceRequest{NewBalance: newAmount, UserId: transfer.From})
	if err != nil {
		return nil, err
	}

	resp, err = walletClient.GetBalance(ctx, &walletpb.GetBalanceRequest{UserId: transfer.To})
	if err != nil {
		return nil, err
	}
	newAmount = resp.Amount + transfer.Amount
	_, err = walletClient.UpdateBalance(ctx, &walletpb.UpdateBalanceRequest{NewBalance: newAmount, UserId: transfer.To})
	if err != nil {
		return nil, err
	}

	return &proto.TransferResponse{IsSuccessful: true}, nil
}

func (h *handler) DebitAccount(ctx context.Context, transfer *proto.Transfer) (*proto.TransferResponse, error) {
	tfs := unMarshaTransfer(transfer)
	tfs.Type = "debit"
	err := h.transferRepo.Add(ctx, tfs)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.DialContext(ctx, walletAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	walletClient := walletpb.NewWalletClient(conn)
	resp, err := walletClient.GetBalance(ctx, &walletpb.GetBalanceRequest{UserId: transfer.From})
	if err != nil {
		return nil, err
	}
	newAmount := resp.Amount + transfer.Amount
	_, err = walletClient.UpdateBalance(ctx, &walletpb.UpdateBalanceRequest{NewBalance: newAmount, UserId: transfer.From})
	if err != nil {
		return nil, err
	}

	resp, err = walletClient.GetBalance(ctx, &walletpb.GetBalanceRequest{UserId: transfer.To})
	if err != nil {
		return nil, err
	}
	newAmount = resp.Amount - transfer.Amount
	_, err = walletClient.UpdateBalance(ctx, &walletpb.UpdateBalanceRequest{NewBalance: newAmount, UserId: transfer.To})
	if err != nil {
		return nil, err
	}

	return &proto.TransferResponse{IsSuccessful: true}, nil
}

func unMarshaTransfer(transfer *proto.Transfer) *core.Transfer {
	return &core.Transfer{
		To:          core.UserID(transfer.To),
		From:        core.UserID(transfer.From),
		Amount:      int(transfer.Amount),
		PerformedAt: transfer.PerformedAt.AsTime(),
	}
}

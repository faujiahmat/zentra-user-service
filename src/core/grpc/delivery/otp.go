package delivery

import (
	"context"

	pb "github.com/faujiahmat/zentra-proto/protogen/otp"
	"github.com/faujiahmat/zentra-user-service/src/common/log"
	"github.com/faujiahmat/zentra-user-service/src/infrastructure/config"
	"github.com/faujiahmat/zentra-user-service/src/interface/delivery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OtpGrpcImpl struct {
	client pb.OtpServiceClient
}

func NewOtpGrpc() (delivery.OtpGrpc, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(
		opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	conn, err := grpc.NewClient(config.Conf.ApiGateway.BaseUrl, opts...)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "delivery.NewOtpGrpc", "section": "grpc.NewClient"}).Fatal(err)
	}

	client := pb.NewOtpServiceClient(conn)

	return &OtpGrpcImpl{
		client: client,
	}, nil
}

func (u *OtpGrpcImpl) Send(ctx context.Context, email string) error {
	_, err := u.client.Send(ctx, &pb.SendReq{
		Email: email,
	})

	return err
}

func (u *OtpGrpcImpl) Verify(ctx context.Context, data *pb.VerifyReq) (*pb.VerifyRes, error) {
	res, err := u.client.Verify(ctx, &pb.VerifyReq{
		Email: data.Email,
		Otp:   data.Otp,
	})

	if err != nil {
		return nil, err
	}

	return res, err
}

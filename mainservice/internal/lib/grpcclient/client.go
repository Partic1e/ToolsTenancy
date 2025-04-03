package grpcclient

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	pbAd "mainservice/api/ad"
	pbPayment "mainservice/api/payment"
	pbRent "mainservice/api/rent"
	pbUser "mainservice/api/user"
	"sync"
)

type GrpcClient struct {
	conns *sync.Pool
}

func NewGrpcClient(host string, port string) *GrpcClient {
	pool := &sync.Pool{
		New: func() interface{} {
			conn, err := grpc.Dial(
				fmt.Sprintf("%s:%s", host, port),
				grpc.WithTransportCredentials(insecure.NewCredentials()),
			)
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}

			return conn
		},
	}

	return &GrpcClient{conns: pool}
}

func (c *GrpcClient) GetOrCreateUser(ctx context.Context, tgId int64) (*pbUser.GetOrCreateUserResponse, error) {
	clientConn := c.conns.Get().(*grpc.ClientConn)
	if clientConn == nil {
		return nil, fmt.Errorf("не удалось получить соединение из пула")
	}
	defer c.conns.Put(clientConn)

	userClient := pbUser.NewUserServiceClient(clientConn)

	req := &pbUser.GetOrCreateUserRequest{TgId: tgId}

	resp, err := userClient.GetOrCreateUser(ctx, req)
	if err != nil {
		log.Printf("Ошибка при вызове GetOrCreateUser: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *GrpcClient) Deposit(ctx context.Context, userId int64, amount string) (bool, error) {
	clientConn := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(clientConn)

	paymentClient := pbPayment.NewPaymentServiceClient(clientConn)

	req := &pbPayment.DepositRequest{
		UserId: userId,
		Amount: &wrapperspb.StringValue{Value: amount},
	}

	resp, err := paymentClient.Deposit(ctx, req)
	if err != nil {
		log.Printf("Ошибка при вызове Deposit: %v", err)
		return false, err
	}

	return resp.Success, nil
}

func (c *GrpcClient) WithDraw(ctx context.Context, userId int64, amount string) (bool, error) {
	clientConn := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(clientConn)

	paymentClient := pbPayment.NewPaymentServiceClient(clientConn)

	req := &pbPayment.WithDrawRequest{
		UserId: userId,
		Amount: &wrapperspb.StringValue{Value: amount},
	}

	resp, err := paymentClient.Withdraw(ctx, req)
	if err != nil {
		log.Printf("Ошибка при вызове WithDraw: %v", err)
		return false, err
	}

	return resp.Success, nil
}

func (c *GrpcClient) CreateAd(ctx context.Context, req *pbAd.CreateAdRequest) (*pbAd.Ad, error) {
	clientConn := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(clientConn)

	adClient := pbAd.NewAdServiceClient(clientConn)

	resp, err := adClient.CreateAd(ctx, req)
	if err != nil {
		log.Printf("Ошибка при вызове CreateAd: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *GrpcClient) DeleteAd(ctx context.Context, req *pbAd.DeleteAdRequest) (bool, error) {
	clientConn := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(clientConn)

	adClient := pbAd.NewAdServiceClient(clientConn)

	resp, err := adClient.DeleteAd(ctx, req)
	if err != nil {
		log.Printf("Ошибка при вызове DeleteAd: %v", err)
		return false, err
	}

	return resp.Success, nil
}

func (c *GrpcClient) UpdateAd(ctx context.Context, req *pbAd.Ad) (*pbAd.Ad, error) {
	clientConn := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(clientConn)

	adClient := pbAd.NewAdServiceClient(clientConn)

	resp, err := adClient.UpdateAd(ctx, req)
	if err != nil {
		log.Printf("Ошибка при вызове UpdateAd: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *GrpcClient) GetAllCategories(ctx context.Context) (*pbAd.CategoryList, error) {
	clientConn := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(clientConn)

	adClient := pbAd.NewAdServiceClient(clientConn)

	resp, err := adClient.GetAllCategories(ctx, &pbAd.Empty{})
	if err != nil {
		log.Printf("Ошибка при вызове GetAllCategories: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *GrpcClient) GetAdsByCategory(ctx context.Context, categoryId int64) (*pbAd.GetAdsByCategoryResponse, error) {
	clientConn := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(clientConn)

	adClient := pbAd.NewAdServiceClient(clientConn)

	req := &pbAd.GetAdsByCategoryRequest{CategoryId: categoryId}

	resp, err := adClient.GetAdsByCategory(ctx, req)
	if err != nil {
		log.Printf("Ошибка при вызове GetAdsByCategory: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *GrpcClient) GetAdsByLandlord(ctx context.Context, landlordId int64) (*pbAd.GetAdsByLandlordResponse, error) {
	clientConn := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(clientConn)

	adClient := pbAd.NewAdServiceClient(clientConn)

	req := &pbAd.GetAdsByLandlordRequest{LandlordId: landlordId}

	resp, err := adClient.GetAdsByLandlord(ctx, req)
	if err != nil {
		log.Printf("Ошибка при вызове GetAdsByLandlord: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *GrpcClient) GetRentsByLandlord(ctx context.Context, landlordId int64) (*pbRent.GetResponse, error) {
	clientConn := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(clientConn)

	adClient := pbRent.NewRentServiceClient(clientConn)

	req := &pbRent.GetRentByLandlordRequest{LandlordId: landlordId}

	resp, err := adClient.GetRentsByLandlord(ctx, req)
	if err != nil {
		log.Printf("Ошибка при вызове GetRentsByLandlord: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *GrpcClient) GetRentsByRenter(ctx context.Context, renterId int64) (*pbRent.GetResponse, error) {
	clientConn := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(clientConn)

	adClient := pbRent.NewRentServiceClient(clientConn)

	req := &pbRent.GetRentByRenterRequest{RenterId: renterId}

	resp, err := adClient.GetRentsByRenter(ctx, req)
	if err != nil {
		log.Printf("Ошибка при вызове GetRentsByRenter: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *GrpcClient) GetRentedDates(ctx context.Context, adId int64) (*pbRent.GetRentedDatesResponse, error) {
	clientConn := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(clientConn)

	adClient := pbRent.NewRentServiceClient(clientConn)

	req := &pbRent.GetRentedDatesRequest{AdId: adId}

	resp, err := adClient.GetRentedDates(ctx, req)
	if err != nil {
		log.Printf("Ошибка при вызове GetRentedDates: %v", err)
		return nil, err
	}

	return resp, nil
}

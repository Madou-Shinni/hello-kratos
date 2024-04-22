package biz

import (
	"context"
	v1 "helloword/api/helloworld/v1"
	stockV1 "helloword/api/stock/v1"

	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	// 导入 kratos 的 dtm 驱动
	_ "github.com/dtm-labs/driver-kratos"
)

var dtmServer = "localhost:36790"

// 业务地址，下面的 busi 换成实际在 server 初始化设置的名字, dtm将采用registry来访问busi
var stockServer = "localhost:9000"

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	Update(context.Context, *Greeter) (*Greeter, error)
	FindByID(context.Context, int64) (*Greeter, error)
	ListByHello(context.Context, string) ([]*Greeter, error)
	ListAll(context.Context) ([]*Greeter, error)
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo        GreeterRepo
	log         *log.Helper
	stockClient stockV1.StockClient
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo, logger log.Logger, stockClient stockV1.StockClient) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger), stockClient: stockClient}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}

func (uc *GreeterUsecase) DeductStock(ctx context.Context, req *stockV1.DeductStockRequest) error {
	uc.log.Infof("DeductStock: %v", req)
	gid := dtmgrpc.MustGenGid(dtmServer)
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(stockServer+stockV1.Stock_DeductStock_FullMethodName, stockServer+stockV1.Stock_IncreaseStock_FullMethodName, req)

	saga.WaitResult = true
	return saga.Submit()
}

func (uc *GreeterUsecase) AddStock(ctx context.Context, req *stockV1.IncreaseStockRequest) error {
	uc.log.Infof("AddStock: %v", req)
	_, err := uc.stockClient.IncreaseStock(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

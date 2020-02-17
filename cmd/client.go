package main 


//测试负载均衡
import (
	"context"

	"fmt"
	"os"
	
	
	"time"
"google.golang.org/grpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	//grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	
	
	
	//"google.golang.org/grpc/codes"
	
	golog "log"

	"huzhu_service/register/consul"
	service "huzhu_service/pkg/svc"
	 "huzhu_service/pkg/transports"
)
	const (
	target      = "consul://127.0.0.1:8500/grpc.health.v1.Addsvc"
	defaultName = "world"
 )
func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = level.NewFilter(logger, level.AllowInfo())
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	
	
	var (
		svc service.AddsvcService
		err error
	)

	consul.Init(logger);

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	conn, err := grpc.DialContext(ctx, target, grpc.WithBlock(), grpc.WithInsecure(), grpc.WithBalancerName("round_robin"))
	if err != nil {
		golog.Fatalf("did not connect: %v", err)
	}




	svc = transports.NewGRPCClient(conn, logger)
	var a int64 =14;
		var b int64 =10;
		ctx2, _ := context.WithTimeout(context.Background(), time.Second)
		v, err := svc.Sum(ctx2, a, b)
		fmt.Println(err);
		fmt.Fprintf(os.Stdout, "%d + %d = %d\n", a, b, v)
}


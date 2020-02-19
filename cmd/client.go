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
	service "huzhu_service/pkg/svc/message"
	"huzhu_service/pb"
	addS "huzhu_service/pkg/svc/add"
	transports "huzhu_service/pkg/transports/message"
	addtransports "huzhu_service/pkg/transports/add"
)
	const (
	target      = "consul://127.0.0.1:8500/microsoft_micro_service"
	defaultName = "world"
 )
func main() {
	//addsvc();
	msgsvc();
	
}

func msgsvc(){

	//file,_:=os.OpenFile("/tmp/log.log",os.O_CREATE|os.O_APPEND,777);
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		//logger = log.NewLogfmtLogger(os.Stderr)
		logger = level.NewFilter(logger, level.AllowAll())
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	logger.Log("foo", "bar") // as normal, no level
	level.Debug(logger).Log("request_id", 123, "trace_data", 1111)
	 level.Error(logger).Log("value", 123)
	
	var (
		svc service.MsgsvcService
		err error
	)

	consul.Init(logger);

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	conn, err := grpc.DialContext(ctx, target, grpc.WithBlock(), grpc.WithInsecure(), grpc.WithBalancerName("round_robin"))
	if err != nil {
		golog.Fatalf("did not connect: %v", err)
	}


	c:=pb.NewMsgsvcClient(conn);

	ctx2, _ := context.WithTimeout(context.Background(), time.Second)

	r, err := c.Echo(ctx2, &pb.EchoRequest{Word: "你妈"})
	fmt.Println("grpc r,",r,err);

	svc = transports.NewGRPCClient(conn, logger)
	fmt.Println(svc);
	
		//ctx2, _ := context.WithTimeout(context.Background(), time.Second)
		v, err := svc.Echo(ctx2, "dfdf")
		fmt.Println(v);
}

func addsvc(){
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = level.NewFilter(logger, level.AllowInfo())
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	
	
	var (
		svc addS.AddsvcService
		err error
	)

	consul.Init(logger);

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	conn, err := grpc.DialContext(ctx, target, grpc.WithBlock(), grpc.WithInsecure(), grpc.WithBalancerName("round_robin"))
	if err != nil {
		golog.Fatalf("did not connect: %v", err)
	}




	svc = addtransports.NewGRPCClient(conn, logger)
	var a int64 =14;
		var b int64 =10;
		ctx2, _ := context.WithTimeout(context.Background(), time.Second)
		v, err := svc.Sum(ctx2, a, b)
		fmt.Println(err);
		fmt.Fprintf(os.Stdout, "%d + %d = %d\n", a, b, v)
}

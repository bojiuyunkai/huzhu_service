package server
import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	
	
	"google.golang.org/grpc/health/grpc_health_v1"
	"huzhu_service/pb"
	
	
	service "huzhu_service/pkg/svc"
	"huzhu_service/pkg/transports"
	"huzhu_service/register/consul"
)

const (
	defConsulHost     string = "localhost"
	defConsulPort     string = "8500"
	defNameSpace      string = "microsoft"
	defServiceName    string = "micro_service"
	defLogLevel       string = "error"
	defServiceHost    string = "localhost"
	defHTTPPort       string = "8282"
	defGRPCPort       string = "8281"
	envConsulHost     string = "ENV_CONSULE_HOST"
	envConsultPort    string = "ENV_CONSULE_PORT"
	envNameSpace      string = "ENV_NAMESPACE"
	envServiceName    string = "ENV_SERVICE_NAME"
	envLogLevel       string = "ENV_LOG_LEVEL"
	envServiceHost    string = "ENV_SERVICE_HOST"
	envHTTPPort       string = "ENV_HTTP_PORT"
	envGRPCPort       string = "ENV_GRPC_PORT"
)

type config struct {
	nameSpace      string `json:"name_space"`
	serviceName    string `json:"service_name"`
	logLevel       string `json:"log_level"`
	serviceHost    string `json:"service_host"`
	httpPort       string `json:"http_port"`
	grpcPort       string `json:"grpc_port"`
	consulHost     string `json:"consul_host"`
	consultPort    string `json:"consult_port"`

}

// Env reads specified environment variable. If no value has been found,
// fallback is returned.
func env(key string, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func RunServer() error {
	//日志
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = level.NewFilter(logger, level.AllowInfo())
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	cfg := loadConfig(logger)

	// consul 服务注册
	{
		
		consulAddres := fmt.Sprintf("%s:%s", cfg.consulHost, cfg.consultPort)
		grpcPort, _ := strconv.Atoi(cfg.grpcPort)
		metricsPort, _ := strconv.Atoi(cfg.httpPort)
		consulReg := consul.NewConsulRegister(consulAddres,cfg.serviceName, grpcPort, metricsPort, []string{cfg.nameSpace, cfg.serviceName}, logger)
		svcRegistar, err := consulReg.NewConsulGRPCRegister()
		defer svcRegistar.Deregister()

		if err != nil {
				level.Error(logger).Log(
					"consulAddres", consulAddres,
					"serviceName", cfg.serviceName,
					"grpcPort", grpcPort,
					"metricsPort", metricsPort,
					"tags", []string{cfg.nameSpace, cfg.serviceName},
					"err", err,
				)
		}
		svcRegistar.Register()
	}

	
	errs := make(chan error, 2)
	//grpcServer, httpHandler := NewServer(cfg, logger)
	grpcServer := transports.MakeGRPCServer(  logger)
	//go startHTTPServer(cfg, httpHandler, logger, errs)
	go startGRPCServer(cfg, grpcServer, logger, errs)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err := <-errs
	level.Info(logger).Log("serviceName", cfg.serviceName, "terminated", err)
	return err;
}

func loadConfig(logger log.Logger) (cfg config) {

	cfg.nameSpace = env(envNameSpace, defNameSpace)
	cfg.serviceName = env(envServiceName, defServiceName)
	cfg.logLevel = env(envLogLevel, defLogLevel)
	cfg.serviceHost = env(envServiceHost, defServiceHost)
	cfg.httpPort = env(envHTTPPort, defHTTPPort)
	cfg.grpcPort = env(envGRPCPort, defGRPCPort)
	cfg.consulHost = env(envConsulHost, defConsulHost)
	cfg.consultPort = env(envConsultPort, defConsulPort)
	return cfg
}

/*func NewServer(cfg config, logger log.Logger) (*transports.GrpcServer, http.Handler) {
	service := service.New(logger)
	m:=msgsvc.New(logger);
	endpoints := endpoints.New(service, logger,m)
	httpHandler := transports.NewHTTPHandler(endpoints,  logger)
	grpcServer := transports.MakeGRPCServer(endpoints,  logger)

	return grpcServer, httpHandler
}*/

func startHTTPServer(cfg config, httpHandler http.Handler, logger log.Logger, errs chan error) {
	p := fmt.Sprintf(":%s", cfg.httpPort)
	
	level.Info(logger).Log("serviceName", cfg.serviceName, "protocol", "HTTP", "exposed", cfg.httpPort)
	errs <- http.ListenAndServe(p, httpHandler)
	
}

func startGRPCServer(cfg config, grpcServer *transports.GrpcServer, logger log.Logger, errs chan error) {
	p := fmt.Sprintf(":%s", cfg.grpcPort)
	listener, err := net.Listen("tcp", p)
	if err != nil {
		level.Error(logger).Log("serviceName", cfg.serviceName, "protocol", "GRPC", "listen", cfg.grpcPort, "err", err)
		os.Exit(1)
	}

	var server *grpc.Server
	
	level.Info(logger).Log("serviceName", cfg.serviceName, "protocol", "GRPC", "exposed", cfg.grpcPort)
	server = grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	pb.RegisterAddsvcServer(server, grpcServer)
	pb.RegisterMsgsvcServer(server, grpcServer);
	grpc_health_v1.RegisterHealthServer(server, &service.HealthImpl{})
	errs <- server.Serve(listener)
}

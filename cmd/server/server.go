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
	defConsulHost     string = ""
	defConsulPort     string = ""
	defZipkinV1URL    string = ""
	defZipkinV2URL    string = ""
	defLightstepToken string = ""
	defAppdashAddr    string = ""
	defNameSpace      string = "gokitconsul"
	defServiceName    string = "addsvc"
	defLogLevel       string = "error"
	defServiceHost    string = "localhost"
	defHTTPPort       string = "8180"
	defGRPCPort       string = "8181"
	defServerCert     string = ""
	defServerKey      string = ""
	defClientTLS      string = "false"
	defCACerts        string = ""
	envConsulHost     string = "QS_CONSULT_HOST"
	envConsultPort    string = "QS_CONSULT_PORT"
	envZipkinV1URL    string = "QS_ZIPKIN_V1_URL"
	envZipkinV2URL    string = "QS_ZIPKIN_V2_URL"
	envLightstepToken string = "QS_LIGHT_STEP_TOKEN"
	envAppdashAddr    string = "QS_APPDASH_ADDR"
	envNameSpace      string = "QS_addsvc_NAMESPACE"
	envServiceName    string = "QS_addsvc_SERVICE_NAME"
	envLogLevel       string = "QS_ADDSVC_LOG_LEVEL"
	envServiceHost    string = "QS_ADDSVC_SERVICE_HOST"
	envHTTPPort       string = "QS_ADDSVC_HTTP_PORT"
	envGRPCPort       string = "QS_ADDSVC_GRPC_PORT"
	envServerCert     string = "QS_ADDSVC_SERVER_CERT"
	envServerKey      string = "QS_ADDSVC_SERVER_KEY"
	envClientTLS      string = "QS_ADDSVC_CLIENT_TLS"
	envCACerts        string = "QS_ADDSVC_CA_CERTS"
)

type config struct {
	nameSpace      string `json:"name_space"`
	serviceName    string `json:"service_name"`
	logLevel       string `json:"log_level"`
	clientTLS      bool   `json:"client_tls"`
	caCerts        string `json:"ca_certs"`
	serviceHost    string `json:"service_host"`
	httpPort       string `json:"http_port"`
	grpcPort       string `json:"grpc_port"`
	serverCert     string `json:"server_cert"`
	serverKey      string `json:"server_key"`
	consulHost     string `json:"consul_host"`
	consultPort    string `json:"consult_port"`
	zipkinV1URL    string `json:"zipkin_v1url"`
	zipkinV2URL    string `json:"zipkin_v2url"`
	lightstepToken string `json:"lightstep_token"`
	appdashAddr    string `json:"appdash_addr"`
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
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = level.NewFilter(logger, level.AllowInfo())
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	cfg := loadConfig(logger)
	// consul
	{
		
			consulAddres := fmt.Sprintf("%s:%s", "127.0.0.1", "8500")
			grpcPort, _ := strconv.Atoi(cfg.grpcPort)
			metricsPort, _ := strconv.Atoi(cfg.httpPort)
			consulReg := consul.NewConsulRegister(consulAddres, "Addsvc", grpcPort, metricsPort, []string{"test","grpc.health.v1.Addsvc"}, logger)
			svcRegistar, err := consulReg.NewConsulGRPCRegister()
			fmt.Println(err);
			defer svcRegistar.Deregister()
			
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
	tls, err := strconv.ParseBool(env(envClientTLS, defClientTLS))
	if err != nil {
		level.Error(logger).Log("envClientTLS", envClientTLS, "error", err)
	}

	cfg.nameSpace = env(envNameSpace, defNameSpace)
	cfg.serviceName = env(envServiceName, defServiceName)
	cfg.logLevel = env(envLogLevel, defLogLevel)
	cfg.clientTLS = tls
	cfg.caCerts = env(envCACerts, defCACerts)
	cfg.serviceHost = env(envServiceHost, defServiceHost)
	cfg.httpPort = "8864";
	cfg.grpcPort = "8866"
	cfg.serverCert = env(envServerCert, defServerCert)
	cfg.serverKey = env(envServerKey, defServerKey)
	cfg.consulHost = env(envConsulHost, defConsulHost)
	cfg.consultPort = env(envConsultPort, defConsulPort)
	cfg.zipkinV1URL = env(envZipkinV1URL, defZipkinV1URL)
	cfg.zipkinV2URL = env(envZipkinV2URL, defZipkinV2URL)
	cfg.lightstepToken = env(envLightstepToken, defLightstepToken)
	cfg.appdashAddr = env(envAppdashAddr, defAppdashAddr)
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

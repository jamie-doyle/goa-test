package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"

	tableinfoservice "github.com/jamie-doyle/goa-test"
	tableinfo "github.com/jamie-doyle/goa-test/gen/table_info"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "0.0.0.0", "Server host (valid values: 0.0.0.0)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		grpcPortF = flag.String("grpc-port", "", "gRPC port (overrides host gRPC port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[tableinfoservice] ", log.Ltime)
	}

	// Initialize the services.
	var (
		tableInfoSvc tableinfo.Service
	)
	{
		tableInfoSvc = tableinfoservice.NewTableInfo(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		tableInfoEndpoints *tableinfo.Endpoints
	)
	{
		tableInfoEndpoints = tableinfo.NewEndpoints(tableInfoSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "0.0.0.0":
		{
			addr := "grpc://0.0.0.0:8000"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "grpcs"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *grpcPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *grpcPortF
			} else if u.Port() == "" {
				u.Host += ":8080"
			}
			handleGRPCServer(ctx, u, tableInfoEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: 0.0.0.0)", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}

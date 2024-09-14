package main

import (
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/kitex-contrib/registry-etcd/retry"

	"github.com/cylScripter/apiopen/admin/admin"
	"github.com/cylScripter/apiopenserver/adminserver/conf"
	"github.com/cylScripter/apiopenserver/adminserver/impl"
)

func main() {

	// Get configuration from conf package
	config := conf.GetConfig()

	addr, err := net.ResolveTCPAddr("tcp", config.Service.Address)
	if err != nil {
		panic(err)
	}

	// Configure retry policy for Etcd registry
	retryConfig := retry.NewRetryConfig(
		retry.WithMaxAttemptTimes(10),          // Maximum number of retry attempts
		retry.WithObserveDelay(20*time.Second), // Delay before observing retry
		retry.WithRetryDelay(5*time.Second),    // Delay between retries
	)

	// Initialize Etcd registry with retry mechanism
	r, err := etcd.NewEtcdRegistryWithRetry(config.Registry.RegistryAddress, retryConfig)
	if err != nil {
		klog.Errorf("failed to init etcd registry: %v", err)
		return
	}

	// Create a new Kitex server instance
	svr := admin.NewServer(
		new(impl.AdminImpl),
		server.WithRegistry(r), // Register the service with Etcd
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: config.Service.ServiceName, // Set the service name
		}),
		server.WithServiceAddr(addr),
	)

	impl.InitState()

	impl.InitMq()

	// Run the server
	err = svr.Run()
	if err != nil {
		klog.Errorf("server stopped with error: %v", err)
	}
}

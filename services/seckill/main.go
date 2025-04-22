package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"

	"gomall/kitex_gen/seckill/seckillservice"
	"gomall/services/seckill/config"
	"gomall/services/seckill/handler"
	"gomall/services/seckill/initialize"
)

func main() {
	// 加载配置
	if err := config.Init("config.json"); err != nil {
		klog.Fatalf("Load config failed: %v", err)
	}
	cfg := config.GetConf()

	// 初始化所有组件
	initialize.Init()
	defer func() {
		// 关闭数据库连接
		if db := initialize.GetMysql(); db != nil {
			sqlDB, _ := db.DB()
			if sqlDB != nil {
				sqlDB.Close()
			}
		}
		// 关闭Redis连接
		if redis := initialize.GetRedis(); redis != nil {
			redis.Close()
		}
		// 关闭RocketMQ
		initialize.CloseRocketMQ()
	}()

	// 服务注册
	r, err := etcd.NewEtcdRegistry([]string{cfg.Etcd.Address})
	if err != nil {
		klog.Fatalf("Create registry failed: %v", err)
	}

	// 服务地址
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", cfg.Service.Port))
	if err != nil {
		klog.Fatalf("Resolve TCP address failed: %v", err)
	}

	// 创建服务实例
	impl := handler.NewSeckillServiceImpl()

	// 设置连接限制
	connLimitOpts := server.WithLimit(&limit.Option{
		MaxConnections: int(cfg.Service.MaxConnections),
		MaxQPS:         int(cfg.Service.MaxQPS),
	})

	// 创建服务器
	svr := seckillservice.NewServer(
		impl,
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{
			ServiceName: cfg.Service.Name,
			Weight:      10,
		}),
		connLimitOpts,
	)

	// 启动定时任务管理器
	cronTaskManager := handler.NewCronTaskManager(impl)
	cronTaskManager.Start()
	defer cronTaskManager.Stop()

	// 优雅退出处理
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
	go func() {
		s := <-ch
		klog.Infof("Received signal: %s", s)
		cronTaskManager.Stop()
		svr.Stop()
	}()

	// 启动服务
	klog.Infof("Seckill service starting on %s", addr)
	err = svr.Run()
	if err != nil {
		klog.Fatalf("Service run failed: %v", err)
	}
}

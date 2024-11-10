package database

//func InitTracer(serviceName string) (func(context.Context) error, error) {
//	// 配置 Jaeger 的主机和端口
//	jaegerEndpoint := fmt.Sprintf("%s:%d", config.Common.Jaeger.Host, config.Common.Jaeger.Port)
//
//	// 创建一个使用 HTTP 协议连接到 Jaeger 的 Exporter
//	ctx := context.Background()
//	exporter, err := otlptracehttp.New(ctx,
//		otlptracehttp.WithEndpoint(jaegerEndpoint),
//		otlptracehttp.WithInsecure(),
//	)
//	if err != nil {
//		return nil, err
//	}
//
//	// 创建资源对象，包含服务的元数据
//	res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceName(serviceName)))
//	if err != nil {
//		return nil, err
//	}
//
//	// 创建 TracerProvider
//	tp := traceSDK.NewTracerProvider(
//		traceSDK.WithBatcher(exporter, traceSDK.WithBatchTimeout(time.Second)), // 处理批量处理时间间隔
//		traceSDK.WithSampler(traceSDK.AlwaysSample()),                          // 选择总是采样
//		traceSDK.WithResource(res),
//	)
//
//	// 设置全局 TracerProvider
//	otel.SetTracerProvider(tp) // 统一导出的trace
//
//	//设置传播提取器
//	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
//
//	// 返回用于关闭 TracerProvider 的函数
//	return tp.Shutdown, nil
//}

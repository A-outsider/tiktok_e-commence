package logs

import (
	"go.uber.org/zap"
)

// 获取Span以及相关信息
//func getSpanInfo(ctx context.Context) (context.Context, trace.Span, string, string) {
//	span := trace.SpanFromContext(ctx)
//	if !span.IsRecording() {
//		tracer := otel.Tracer("default")
//		ctx, span = tracer.Start(ctx, "default")
//		defer span.End()
//	}
//
//	traceID := span.SpanContext().TraceID().String()
//	spanID := span.SpanContext().SpanID().String()
//	return ctx, span, traceID, spanID
//}
//
//// Span 相关日志函数
//func ErrorWithSpan(ctx context.Context, err error) {
//	ctx, span, traceID, spanID := getSpanInfo(ctx)
//	span.RecordError(err)
//	span.SetStatus(codes.Error, "服务器异常")
//
//	zap.L().Error("Error occurred",
//		zap.String("traceID", traceID),
//		zap.String("spanID", spanID),
//		zap.Error(err),
//	)
//}

//func InfoWithSpan(ctx context.Context, msg string, fields ...zap.Field) {
//	_, _, traceID, spanID := getSpanInfo(ctx)
//
//	zap.L().Info(msg,
//		append(fields,
//			zap.String("traceID", traceID),
//			zap.String("spanID", spanID),
//		)...,
//	)
//}
//
//func DebugWithSpan(ctx context.Context, msg string, fields ...zap.Field) {
//	_, _, traceID, spanID := getSpanInfo(ctx)
//
//	zap.L().Debug(msg,
//		append(fields,
//			zap.String("traceID", traceID),
//			zap.String("spanID", spanID),
//		)...,
//	)
//}

// 非 Span 相关日志函数
func Error(msg string, fields ...zap.Field) {
	zap.L().Error(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	zap.L().Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	zap.L().Debug(msg, fields...)
}

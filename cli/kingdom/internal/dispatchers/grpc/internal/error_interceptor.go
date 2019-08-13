package internal

import (
	"context"

	"go.zenithar.org/pkg/errors"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var grpcErrorMap = map[errors.ErrorCode]codes.Code{
	errors.AlreadyExists:      codes.AlreadyExists,
	errors.NotFound:           codes.NotFound,
	errors.Internal:           codes.Internal,
	errors.InvalidArgument:    codes.InvalidArgument,
	errors.PermissionDenied:   codes.PermissionDenied,
	errors.FailedPrecondition: codes.FailedPrecondition,
	errors.Unauthenticated:    codes.Unauthenticated,
}

// ServiceErrorTranslationUnaryServerInterceptor ...
func ServiceErrorTranslationUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// Calls the handler
		h, err := handler(ctx, req)
		if err != nil {
			// Translate to grpc errors
			var srvErr *errors.Error
			if xerrors.As(err, &srvErr) {
				if code, ok := grpcErrorMap[srvErr.Code]; ok {
					return h, status.Errorf(code, srvErr.Error())
				}
				return h, status.Errorf(codes.Unknown, srvErr.Error())
			}
		}

		// Return result
		return h, err
	}
}

package mocks

import "cine/pkg/logger"

var _ logger.Logger = (*NopLogger)(nil)

type NopLogger struct{}

func (l NopLogger) Info(message string) {}

func (l NopLogger) Warn(context string, message string) {}

func (l NopLogger) Error(context string, err error) {}

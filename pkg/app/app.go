package app

import (
	nf_context "github.com/andy89923/lab4-af/internal/context"
	"github.com/andy89923/lab4-af/pkg/factory"
)

type App interface {
	SetLogEnable(enable bool)
	SetLogLevel(level string)
	SetReportCaller(reportCaller bool)

	Start()
	Terminate()

	Context() *nf_context.NFContext
	Config() *factory.Config
}

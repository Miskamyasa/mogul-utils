package flags

import (
	"context"
	"os"

	"github.com/Miskamyasa/utils/alerts"
	"github.com/go-logr/zerologr"
	flagd "github.com/open-feature/go-sdk-contrib/providers/flagd/pkg"
	"github.com/open-feature/go-sdk/openfeature"
)

var (
	client *openfeature.Client
	ctx    context.Context
)

func InitFlags() func() {
	zl := alerts.CreateLogger().With().Str("component", "flags").Logger()
	logger := zerologr.New(&zl)

	// Use flagd as the OpenFeature provider
	err := openfeature.SetProviderAndWait(flagd.NewProvider(
		flagd.WithLogger(logger),
	))
	if err != nil {
		alerts.Fatal("Failed to set OpenFeature (flagd) provider", err)
		return nil
	}

	ctx = context.Background()

	client = openfeature.NewClient(os.Getenv("SERVICE_NAME"))

	client.SetEvaluationContext(openfeature.NewTargetlessEvaluationContext(
		map[string]interface{}{
			"version": os.Getenv("SERVICE_VERSION"),
			"env":     os.Getenv("ENV"),
		},
	))

	// evalCtx = openfeature.NewEvaluationContext(
	//     os.Getenv("SERVICE_NAME"),
	//     os.Getenv("SERVICE_VERSION"),
	// )

	return openfeature.Shutdown
}

func GetClient() *openfeature.Client {
	return client
}

func GetBoolFlag(flag string, defaultValue bool) bool {
	val, err := client.BooleanValue(ctx, flag, defaultValue, client.EvaluationContext())
	if err != nil {
		alerts.Send("Failed to get flag", err)
		return defaultValue
	}
	return val
}

func GetStringFlag(flag string, defaultValue string) string {
	val, err := client.StringValue(ctx, flag, defaultValue, client.EvaluationContext())
	if err != nil {
		alerts.Send("Failed to get flag", err)
		return defaultValue
	}
	return val
}

func GetIntFlag(flag string, defaultValue int64) int64 {
	val, err := client.IntValue(ctx, flag, defaultValue, client.EvaluationContext())
	if err != nil {
		alerts.Send("Failed to get flag", err)
		return defaultValue
	}
	return val
}

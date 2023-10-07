package mongodb_test

import (
	"context"
	"testing"

	"github.com/audryus/miniature-octo-tribble/config"
	"github.com/audryus/miniature-octo-tribble/pkg/mongodb"
	"github.com/audryus/miniature-octo-tribble/types"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, types.Config, &config.Config{
		Mongo: config.Mongo{
			URL:      "mongodb://root:example@localhost:27017/",
			DATABASE: "tribble",
		},
	})
	ctx = mongodb.New(ctx)
	mg := ctx.Value(types.Mongodb)

	assert.NotNil(t, mg, "Sem conexão")
}

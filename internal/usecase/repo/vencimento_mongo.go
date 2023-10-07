package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/audryus/miniature-octo-tribble/internal/usecase/repo/entity"
	"github.com/audryus/miniature-octo-tribble/pkg/mongodb"
	"github.com/audryus/miniature-octo-tribble/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VencimentoRepo struct {
	col *mongo.Collection
}

func NewVencimentoRepo(ctx context.Context) context.Context {
	col := mongodb.Get(ctx).Db.Collection("vencimento")
	repo := &VencimentoRepo{
		col: col,
	}
	return context.WithValue(ctx, types.VencimentoRepo, repo)
}

func GetVencimentoRepo(ctx context.Context) *VencimentoRepo {
	return ctx.Value(types.VencimentoRepo).(*VencimentoRepo)
}

func (r *VencimentoRepo) SaveAll(series []entity.Serie) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	vencimentos := make(map[string]entity.Vencimento)

	for _, serie := range series {
		t, _ := time.Parse("20060102", serie.Vencimento)
		vencimentos[serie.Vencimento] = entity.Vencimento{
			Vencimento: t,
		}
	}

	for _, v := range vencimentos {
		r.col.InsertOne(ctx, v)
	}

}

func (r *VencimentoRepo) GetNextTwo() []entity.Vencimento {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	filter := bson.D{{}}
	opts := options.Find().SetSort(bson.D{{Key: "_id", Value: 1}}).SetLimit(2)

	cursor, err := r.col.Find(ctx, filter, opts)

	var results []entity.Vencimento
	if err != nil {
		fmt.Println("error", err.Error())
		return results
	}

	if err = cursor.All(ctx, &results); err != nil {
		fmt.Println("error", err.Error())
		return results
	}
	return results
}

func (r *VencimentoRepo) DeleteVencimentoPassado() *entity.Vencimento {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	corte := time.Now().AddDate(0, 0, 5).UTC()

	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$lt", Value: corte}}}}

	var result entity.Vencimento

	if err := r.col.FindOneAndDelete(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Printf("Nenhum Vencimento menor que %s encontrando.\n", corte)
			return nil
		}
		fmt.Println("Erro ao tentar remover Vencimento antigo.", err)
		return nil
	}

	fmt.Printf("%+v Removido com sucesso.\n", result)
	return &result
}

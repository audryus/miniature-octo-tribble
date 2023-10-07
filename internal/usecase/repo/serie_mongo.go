package repo

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/audryus/miniature-octo-tribble/internal/usecase/repo/entity"
	"github.com/audryus/miniature-octo-tribble/pkg/mongodb"
	"github.com/audryus/miniature-octo-tribble/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/exp/slices"
)

type SerieRepo interface {
	SaveAll(series []entity.Serie)
	UpdatePrice(serie string, price float64) error
	DeleteByVencimento(vencimento *entity.Vencimento) error
	FindAllByEmpresaAndVencimento(empresa entity.Empresa, vencimento entity.Vencimento) []entity.Serie
}

type serieRepo struct {
	col    *mongo.Collection
	letras []string
}

func NewSerieRepo(ctx context.Context) context.Context {
	col := mongodb.Get(ctx).Db.Collection("serie")
	repo := &serieRepo{
		col:    col,
		letras: strings.Split("A,B,C,D,E,F,G,H,I,J,K,L", ","),
	}
	return context.WithValue(ctx, types.SerieRepo, repo)
}

func GetSerieRepo(ctx context.Context) SerieRepo {
	return ctx.Value(types.SerieRepo).(SerieRepo)
}
func (r *serieRepo) DeleteByVencimento(vencimento *entity.Vencimento) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	filter := bson.D{{Key: "tsVencimento", Value: vencimento.Vencimento}}

	if _, err := r.col.DeleteMany(ctx, filter); err != nil {
		fmt.Println("Erro ao remover Series", err.Error())
		return err
	}
	fmt.Printf("Series com vencimento %+v removidas.\n", vencimento)
	return nil
}

func (r *serieRepo) SaveAll(series []entity.Serie) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	for i, _ := range series {
		t, _ := time.Parse("20060102", series[i].Vencimento)
		series[i].Serie = strings.TrimSpace(series[i].Serie)
		series[i].Tipo = strings.TrimSpace(series[i].Tipo)
		series[i].DtVencimento = t
		if !r.isCall(series[i].Serie) {
			continue
		}
		r.col.InsertOne(ctx, series[i])
	}
}
func (r *serieRepo) UpdatePrice(symbol string, price float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "price", Value: price},
			},
		},
	}

	if _, err := r.col.UpdateByID(ctx, symbol, update); err != nil {
		return err
	}

	return nil
}

func (r *serieRepo) FindAllByEmpresaAndVencimento(empresa entity.Empresa, vencimento entity.Vencimento) []entity.Serie {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	filter := bson.D{
		{Key: "tsVencimento", Value: vencimento.Vencimento},
		{Key: "razao_social", Value: empresa.RazSoc},
		{Key: "tipo", Value: empresa.Tipo},
	}
	cursor, err := r.col.Find(ctx, filter)
	var results []entity.Serie
	if err != nil {
		fmt.Println("Erro ao tentar encontrar Series", err.Error())
		return results
	}
	if err = cursor.All(ctx, &results); err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Printf("Nao existem Series para Empresa %s no vencimento %s\n", empresa.Symbol, vencimento.Vencimento)
			return results
		}
		fmt.Println("Erro ao tentar encontrar Series", err.Error())
		return results
	}

	return results
}

func (r *serieRepo) isCall(serie string) bool {
	var ser string
	for i := len(serie); i > 0; i-- {
		_, err := strconv.Atoi(string(serie[i-1]))
		if err != nil {
			ser = string(serie[i-1])
			break
		}
	}
	return slices.Contains(r.letras, ser)
}

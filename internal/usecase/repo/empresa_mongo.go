package repo

import (
	"context"
	"time"

	"github.com/audryus/miniature-octo-tribble/internal/usecase/repo/entity"
	"github.com/audryus/miniature-octo-tribble/pkg/mongodb"
	"github.com/audryus/miniature-octo-tribble/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmpresaRepo interface {
	Save(entity.Empresa) error
	FindAllWithoutTipo() ([]entity.Empresa, error)
	FindAll() []entity.Empresa
	UpdateSemTipo(entity.Empresa) error
	FindByID(symbol string) (*entity.Empresa, error)
	UpdatePrice(string, float64) error
}

type empresaRepo struct {
	col *mongo.Collection
}

func NewEmpresaRepo(ctx context.Context) context.Context {
	col := mongodb.Get(ctx).Db.Collection("empresa")
	repo := &empresaRepo{
		col: col,
	}
	return context.WithValue(ctx, types.EmpresaRepo, repo)
}

func GetEmpresaRepo(ctx context.Context) EmpresaRepo {
	return ctx.Value(types.EmpresaRepo).(EmpresaRepo)
}

func (r *empresaRepo) SaveAll(empresas []entity.Empresa) error {
	for _, empresa := range empresas {
		if err := r.Save(empresa); err != nil {
			return err
		}
	}
	return nil
}

func (r *empresaRepo) UpdatePrice(symbol string, price float64) error {
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

func (r *empresaRepo) Save(empresa entity.Empresa) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: empresa.Symbol}}
	var result entity.Empresa
	if err := r.col.FindOne(ctx, filter).Decode(&result); err != nil {
		if _, err := r.col.InsertOne(ctx, empresa); err != nil {
			return err
		}
	}

	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "razao_social", Value: empresa.RazSoc},
			},
		},
	}

	if _, err := r.col.UpdateByID(ctx, empresa.Symbol, update); err != nil {
		return err
	}

	return nil
}

func (r *empresaRepo) FindAllWithoutTipo() ([]entity.Empresa, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	filter := bson.D{{Key: "tipo", Value: ""}}
	cursor, err := r.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var empresas []entity.Empresa = make([]entity.Empresa, 0)
	// iterate the cursor and decode each item into an Employee
	if err := cursor.All(ctx, &empresas); err != nil {
		return nil, err

	}
	return empresas, nil

}

func (r *empresaRepo) FindByID(symbol string) (*entity.Empresa, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	filter := bson.D{{Key: "_id", Value: symbol}}
	var result entity.Empresa

	if err := r.col.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *empresaRepo) UpdateSemTipo(emp entity.Empresa) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: emp.Symbol}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "tipo", Value: emp.Tipo},
			},
		},
	}
	r.col.FindOneAndUpdate(ctx, filter, update)

	return nil
}

func (r *empresaRepo) FindAll() []entity.Empresa {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	var empresas []entity.Empresa = make([]entity.Empresa, 0)

	filter := bson.D{{}}
	cursor, err := r.col.Find(ctx, filter)
	if err != nil {
		return empresas
	}

	if err := cursor.All(ctx, &empresas); err != nil {
		return empresas

	}
	return empresas
}

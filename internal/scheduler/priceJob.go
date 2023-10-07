package scheduler

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/audryus/miniature-octo-tribble/internal/usecase"
	"golang.org/x/exp/slices"
)

type PriceJob interface {
	CopyFile() error
	CleanFile() error
	ReadFile() ([][]string, error)
	IsAtivo(symbols []string, symbol string) bool
	Run()
}

type priceJob struct {
	src   string
	dst   string
	emp   usecase.EmpresaUC
	serie usecase.SerieUC
}

func NewPriceJob(ctx context.Context) PriceJob {
	return &priceJob{
		src:   "C:/Users/audry/Documents/YourFile.csv",
		dst:   "D:/projs/miniature-octo-tribble/YourFile.csv",
		emp:   usecase.GetEmpresaUC(ctx),
		serie: usecase.GetSerieUC(ctx),
	}
}

func (j *priceJob) Run() {
	fmt.Println("JOB -- Price Updater  -- starting")
	errCode := j.run()
	if errCode == 1 {
		fmt.Println("JOB -- Price Updater  -- Tentaremos novamente em alguns segundos.")
		time.Sleep(10 * time.Second)
		fmt.Println("JOB -- Price Updater  -- Tentando novamente.")
		j.run()
	} else if errCode == 0 {
		fmt.Println("JOB -- Price Updater  -- sucesso")
	}
	fmt.Println("JOB -- Price Updater  -- ended")
}

func (j *priceJob) run() uint8 {
	if err := j.CopyFile(); err != nil {
		fmt.Printf("JOB -- Price Updater -- Error: %s\n", err.Error())
		return 1
	}
	defer j.CleanFile()

	records, err := j.ReadFile()
	if err != nil {
		fmt.Printf("JOB -- Price Updater -- Error: %s\n", err.Error())
		return 2
	}

	empresas := j.emp.FindAll()
	symbols := make([]string, 0)
	for _, emp := range empresas {
		symbols = append(symbols, emp.Symbol)
	}
	var errCode uint8 = 0
	for _, row := range records {
		symbol := row[0]
		price := strings.Replace(row[1], ",", ".", 1)
		s, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Printf("JOB -- Price Updater -- Error: %s\n", err.Error())
			errCode = 3
			continue
		}
		if j.IsAtivo(symbols, symbol) {
			err = j.emp.UpdatePrice(symbol, s)
		} else {
			err = j.serie.UpdatePrice(symbol, s)
		}
		if err != nil {
			fmt.Printf("JOB -- Price Updater -- Error: %s\n", err.Error())
			errCode = 4
			continue
		}
	}

	return errCode
}

func (j *priceJob) IsAtivo(symbols []string, symbol string) bool {
	return slices.Contains(symbols, symbol)
}

func (j *priceJob) ReadFile() ([][]string, error) {
	f, err := os.Open(j.dst)
	if err != nil {
		return nil, fmt.Errorf("%s Unable to read input file", j.dst)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("%s Unable to parse file as CSV for", j.dst)
	}

	return records, nil
}

func (j *priceJob) CleanFile() error {
	return os.Remove(j.dst)
}

func (j *priceJob) CopyFile() error {
	src := j.src
	dst := j.dst

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

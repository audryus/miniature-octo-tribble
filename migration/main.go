package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/audryus/miniature-octo-tribble/config"
	"github.com/audryus/miniature-octo-tribble/internal/usecase"
	"github.com/audryus/miniature-octo-tribble/internal/usecase/repo"
	"github.com/audryus/miniature-octo-tribble/internal/usecase/repo/entity"
	"github.com/audryus/miniature-octo-tribble/pkg/mongodb"
	"golang.org/x/exp/slices"
)

type Series struct {
	Empresa struct {
		Num3 []entity.Serie `json:"3"`
		A    []entity.Serie `json:"A"`
		B    []entity.Serie `json:"B"`
		C    []entity.Serie `json:"C"`
		D    []entity.Serie `json:"D"`
		E    []entity.Serie `json:"E"`
		F    []entity.Serie `json:"F"`
		G    []entity.Serie `json:"G"`
		H    []entity.Serie `json:"H"`
		I    []entity.Serie `json:"I"`
		J    []entity.Serie `json:"J"`
		K    []entity.Serie `json:"K"`
		L    []entity.Serie `json:"L"`
		M    []entity.Serie `json:"M"`
		N    []entity.Serie `json:"N"`
		O    []entity.Serie `json:"O"`
		P    []entity.Serie `json:"P"`
		Q    []entity.Serie `json:"Q"`
		R    []entity.Serie `json:"R"`
		S    []entity.Serie `json:"S"`
		T    []entity.Serie `json:"T"`
		U    []entity.Serie `json:"U"`
		V    []entity.Serie `json:"V"`
		W    []entity.Serie `json:"W"`
		Y    []entity.Serie `json:"Y"`
		X    []entity.Serie `json:"X"`
		Z    []entity.Serie `json:"Z"`
	} `json:"Empresa"`
}

func main() {
	ctx := context.Background()
	ctx = config.New(ctx)
	ctx = mongodb.New(ctx)
	ctx = repo.NewEmpresaRepo(ctx)
	ctx = repo.NewVencimentoRepo(ctx)
	ctx = repo.NewSerieRepo(ctx)
	ctx = usecase.NewSerieUC(ctx)

	todasEmpresas := strings.Split("AAPL34,ABCB4,ABEV3,AERI3,AESB3,AGRO3,ALPA4,ALSO3,ALUP11,AMAR3,AMBP3,AMER3,AMZO34,ANIM3,ARML3,ARZZ3,ASAI3,AURE3,AZUL4,B3SA3,BABA34,BBAS3,BBDC3,BBDC4,BBSE3,BEEF3,BMGB4,BMOB3,BOVA11,BOVB11,BOVV11,BPAC11,BPAN4,BRAP4,BRAX11,BRFS3,BRKM5,BRSR6,CAML3,CASH3,CBAV3,CCRO3,CEAB3,CIEL3,CLSA3,CMIG3,CMIG4,CMIN3,COGN3,CPFE3,CPLE6,CRFB3,CSAN3,CSMG3,CSNA3,CURY3,CVCB3,CXSE3,CYRE3,DIRR3,DXCO3,ECOR3,EGIE3,ELET3,ELET6,EMBR3,ENAT3,ENEV3,ENGI11,ENJU3,EQTL3,ESPA3,EVEN3,EXCO32,EZTC3,FESA4,FLRY3,GFSA3,GGBR4,GGPS3,GMAT3,GOAU4,GOGL34,GOLL4,GRND3,GUAR3,HAPV3,HBOR3,HBSA3,HYPE3,IBOV11,IFCM3,IGTI11,INBR31,INBR32,INTB3,IRBR3,ITSA4,ITUB3,ITUB4,ITUB99,IVVB11,JALL3,JBSS3,JHSF3,KEPL3,KLBN11,LAVV3,LEVE3,LIGT3,LJQQ3,LOGG3,LOGN3,LREN3,LWSA3,M1TA34,MATD3,MDIA3,MEAL3,MEGA3,MELI34,MGLU3,MILS3,MLAS3,MOVI3,MRFG3,MRVE3,MSFT11B,MSFT34,MULT3,MYPK3,NEOE3,NFLX34,NTCO3,NVDC34,ODPV3,ONCO3,PCAR3,PCAR99,PETR3,PETR4,PETZ3,PIBB11,PNVL3,POMO4,POSI3,PRIO3,PSSA3,PTBL3,PYPL34,QUAL3,RADL3,RAIL3,RAIZ4,RANI3,RAPT4,RDOR3,RECV3,RENT3,ROMI3,RRRP3,SANB11,SAPR11,SBFG3,SBSP3,SEER3,SEQL3,SIMH3,SLCE3,SMAL11,SMFT3,SMTO3,SOMA3,SPXI11,SQIA3,STBP3,SUZB3,TAEE11,TASA4,TCSA3,TEND3,TIMS3,TOTS3,TRAD3,TRIS3,TRPL4,TSLA34,TTEN3,TUPY3,UGPA3,UNIP6,USIM5,VALE3,VAMO3,VBBR3,VIIA3,VIVA3,VIVT3,VLID3,VULC3,WEGE3,WIZC3,WIZS3,XBOV11,XPBR31,YDUQ3,ZAMP3", ",")

	body, err := os.ReadFile("./migration/empresas.json")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	data := []entity.Empresa{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err.Error())
	}
	bodyTipo, err := os.ReadFile("./migration/empresa_tipo.json")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	tipos := []entity.Empresa{}
	err = json.Unmarshal(bodyTipo, &tipos)
	if err != nil {
		fmt.Println(err)
	}
	m := make(map[string]int)

	for _, emp := range tipos {
		i, ok := m[emp.RazSoc]
		if !ok {
			m[emp.RazSoc] = 1
		} else {
			m[emp.RazSoc] = i + 1
		}
	}

	empresaTipo := make(map[string]string)

	for _, emp := range tipos {
		i, ok := m[emp.RazSoc]
		if ok && i == 1 {
			empresaTipo[emp.RazSoc] = emp.Tipo
		}
	}

	empresaRepo := repo.GetEmpresaRepo(ctx)

	for idx, empresa := range data {
		data[idx].RazSoc = strings.Replace(empresa.RazSoc, " - ", "", -1)

		i, ok := empresaTipo[data[idx].RazSoc]
		if ok {
			data[idx].Tipo = i
		}
		if slices.Contains(todasEmpresas, data[idx].Symbol) {
			_, ok = m[data[idx].RazSoc]

			if ok {
				if err := empresaRepo.Save(data[idx]); err != nil {
					fmt.Println(err.Error())
				}
			}
		}
	}

	body, err = os.ReadFile("./migration/SI_C_OPCSEREMP.json")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var series Series
	err = json.Unmarshal(body, &series)
	if err != nil {
		fmt.Println(err.Error())
	}

	sRepo := usecase.GetSerieUC(ctx)
	sRepo.SaveAll(series.Empresa.Num3)
	sRepo.SaveAll(series.Empresa.A)
	sRepo.SaveAll(series.Empresa.C)
	sRepo.SaveAll(series.Empresa.D)
	sRepo.SaveAll(series.Empresa.E)
	sRepo.SaveAll(series.Empresa.F)
	sRepo.SaveAll(series.Empresa.G)
	sRepo.SaveAll(series.Empresa.H)
	sRepo.SaveAll(series.Empresa.I)
	sRepo.SaveAll(series.Empresa.J)
	sRepo.SaveAll(series.Empresa.K)
	sRepo.SaveAll(series.Empresa.L)
	sRepo.SaveAll(series.Empresa.M)
	sRepo.SaveAll(series.Empresa.N)
	sRepo.SaveAll(series.Empresa.O)
	sRepo.SaveAll(series.Empresa.P)
	sRepo.SaveAll(series.Empresa.Q)
	sRepo.SaveAll(series.Empresa.R)
	sRepo.SaveAll(series.Empresa.S)
	sRepo.SaveAll(series.Empresa.T)
	sRepo.SaveAll(series.Empresa.U)
	sRepo.SaveAll(series.Empresa.V)
	sRepo.SaveAll(series.Empresa.W)
	sRepo.SaveAll(series.Empresa.Y)
	sRepo.SaveAll(series.Empresa.X)
	sRepo.SaveAll(series.Empresa.Z)

}

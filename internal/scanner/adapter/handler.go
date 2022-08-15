package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/monopeelz/linear-avocado/internal/scanner/entity"
	"github.com/monopeelz/linear-avocado/internal/scanner/usecase"
	"log"
)

type Handler struct {
	u usecase.ScannerUseCase
}

func (h Handler) ScanProject(ctx context.Context, b []byte) {
	log.Printf("Received a message: %s", b)
	proj := new(entity.Project)
	if err := json.Unmarshal(b, proj); err != nil {
		fmt.Println(err)
	}
	finding, err := h.u.Exec(ctx, *proj)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Success %v", finding)
}

func NewHandler(u usecase.ScannerUseCase) *Handler {
	return &Handler{
		u,
	}
}

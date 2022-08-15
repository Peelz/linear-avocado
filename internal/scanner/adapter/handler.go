package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/monopeelz/linear-avocado/internal/scanner/models"
	"github.com/monopeelz/linear-avocado/internal/scanner/usecase"
	"log"
)

type Handler struct {
	u usecase.ScannerUseCase
}

func (h Handler) ScanProject(ctx context.Context, b []byte) {
	log.Printf("Received a message: %s", b)
	job := new(models.Job)
	if err := json.Unmarshal(b, job); err != nil {
		fmt.Println(err)
	}
	_, err := h.u.Exec(ctx, *job)
	if err != nil {
		fmt.Println(err)
	}
}

func NewHandler(u usecase.ScannerUseCase) *Handler {
	return &Handler{
		u,
	}
}

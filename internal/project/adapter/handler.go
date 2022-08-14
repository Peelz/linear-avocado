package adapter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/monopeelz/linear-avocado/internal/pkg/utils"
	"github.com/monopeelz/linear-avocado/internal/project/models"
	"github.com/monopeelz/linear-avocado/internal/project/usecase"
)

type Handler struct {
	uc usecase.ProjectUseCase
}

func (h Handler) List(ctx *gin.Context) {
	res, err := h.uc.All(ctx.Request.Context())
	if err != nil {
		utils.GinErrorWrapper(ctx, err, 0)
		return
	}
	utils.GinResponseWrapper(ctx, res, 200)
}

func (h Handler) Retrieve(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := h.uc.Get(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	utils.GinResponseWrapper(ctx, res, 200)
}

func (h Handler) Create(ctx *gin.Context) {
	req := new(CreateProjectRequest)
	if err := ctx.Bind(req); err != nil {
		utils.GinErrorWrapper(ctx, err, 400)
		return
	}
	proj := models.Project{
		Name: req.Name,
		URL:  req.URL,
	}
	res, err := h.uc.Create(ctx.Request.Context(), proj)
	if err != nil {
		utils.GinErrorWrapper(ctx, err, 400)
		return
	}
	utils.GinResponseWrapper(ctx, res, 200)
}

func (h Handler) Update(ctx *gin.Context) {
	req := new(UpdateProjectRequest)
	if err := ctx.Bind(req); err != nil {
		utils.GinErrorWrapper(ctx, err, 400)
		return
	}
	fmt.Println()
	proj := models.UpdateProject{
		// UUID: uuid.Must(uuid.FromBytes([]byte(ctx.Param("id")))),
		UUID: ctx.Param("id"),
		Name: req.Name,
		URL:  req.URL,
	}
	res, err := h.uc.Update(ctx.Request.Context(), proj)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	ctx.JSON(200, res)
}

func (h Handler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.uc.Delete(ctx.Request.Context(), id)
	if err != nil {
		utils.GinErrorWrapper(ctx, err, 400)
		return
	}
	utils.GinResponseWrapper(ctx, nil, 200)
}

func (h Handler) Scan(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.uc.Scan(ctx.Request.Context(), id)
	if err != nil {
		utils.GinErrorWrapper(ctx, err, 400)
		return
	}
	utils.GinResponseWrapper(ctx, nil, 200)
}

func NewHandler(uc usecase.ProjectUseCase) *Handler {
	return &Handler{
		uc,
	}
}

func RegisterRoute(e *gin.Engine, h *Handler) {
	group := e.Group("/api/v1/projects")
	group.GET("", h.List)
	group.GET("/:id", h.Retrieve)
	group.POST("", h.Create)
	group.PUT("/:id", h.Update)
	group.DELETE("/:id", h.Delete)
	group.POST("/:id/scan", h.Scan)
}

package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-study/go-backend/api/controller"
	"github.com/go-study/go-backend/bootstrap"
	"github.com/go-study/go-backend/domain"
	"github.com/go-study/go-backend/mongo"
	"github.com/go-study/go-backend/repository"
	"github.com/go-study/go-backend/usecase"
	"time"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}

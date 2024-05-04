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

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}

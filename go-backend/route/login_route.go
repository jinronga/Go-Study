package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-study/go-backend/bootstrap"
	"github.com/go-study/go-backend/controller"
	"github.com/go-study/go-backend/domain"
	"github.com/go-study/go-backend/repository"
	"github.com/go-study/go-backend/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {

	ur := repository.NewUserRepository(db, domain.CollectionUser)

	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}

	group.POST("/login", lc.Login)

}

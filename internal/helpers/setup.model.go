package helpers

import (
	"github.com/Deven-14/gin-mongo-trial/internal/interfaces"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpModel[R interfaces.Router, C interfaces.Controllers, S interfaces.Services, Re interfaces.Repositories](
	database *mongo.Database,
	routerGroup *gin.RouterGroup,
	newRepository func(*mongo.Database) Re,
	newService func(Re) S,
	newController func(S) C,
	newRouter func(*gin.RouterGroup, C) R) {

	repository := newRepository(database)
	service := newService(repository)
	controller := newController(service)
	router := newRouter(routerGroup, controller)
	router.RegisterRoutes()

}

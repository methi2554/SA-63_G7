package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/methi2554/app/controllers"
	_ "github.com/methi2554/app/docs"
	"github.com/methi2554/app/ent"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Employees struct {
	Employee []Employee
}

type Employee struct {
	Name   string
	Userid int
}

type Diseases struct {
	Disease []Disease
}

type Disease struct {
	Name string
}

type DrugTypes struct {
	DrugType []DrugType
}

type DrugType struct {
	Name string
}

// @title SUT SA Example API Drug
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:contagion.db?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewEmployeeController(v1, client)
	controllers.NewDrugTypeController(v1, client)
	controllers.NewDiseaseController(v1, client)
	controllers.NewDrugController(v1, client)

	// Set Employees Data
	employees := Employees{
		Employee: []Employee{
			Employee{"D12345", 1234},
			Employee{"D54321", 7890},
		},
	}

	for _, e := range employees.Employee {
		client.Employee.
			Create().
			SetName(e.Name).
			SetUserid(e.Userid).
			Save(context.Background())
	}

	// Set drugtype Data
	drugtypes := DrugTypes{
		DrugType: []DrugType{
			DrugType{"ยาเม็ด"},
			DrugType{"ยาน้ำ"},
			DrugType{"วัคซีน"},
		},
	}
	for _, dt := range drugtypes.DrugType {
		client.DrugType.
			Create().
			SetName(dt.Name).
			Save(context.Background())
	}

	// Set disease Data
	disease := Diseases{
		Disease: []Disease{
			Disease{"covid"},
			Disease{"ไข้หวัดใหญ่"},
			Disease{"เอดส์"},
			Disease{"ซาร์ส"},
			Disease{"เมอร์ส"},
		},
	}
	for _, di := range disease.Disease {
		client.Disease.
			Create().
			SetName(di.Name).
			Save(context.Background())
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}

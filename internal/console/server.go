package console

import (
	"net/http"

	"github.com/kodinggo/gb-2-api-story-service/internal/config"
	mysqldb "github.com/kodinggo/gb-2-api-story-service/internal/db/mysql"
	httphandler "github.com/kodinggo/gb-2-api-story-service/internal/delivery/http"
	"github.com/kodinggo/gb-2-api-story-service/internal/repository"
	"github.com/kodinggo/gb-2-api-story-service/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve is a command to run the service server",
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()
		dbConn := mysqldb.InitDBConn()
		commentRepository := repository.InitCommentRepository(dbConn)
		commentUseCase := usecase.InitCommentUsecase(commentRepository)
		commentHandler := httphandler.InitCommentHandler(commentUseCase)
		commentHandler.RegisterRoute(e)
		e.GET("/ping", func(c echo.Context) error {
			return c.String(http.StatusOK, "pong!")
		})
		e.Start(":" + config.Port())
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

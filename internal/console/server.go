package console

import (
	"fmt"
	"log"
	"net"

	"net/http"

	"github.com/kodinggo/gb-2-api-comment-service/internal/config"
	mysqldb "github.com/kodinggo/gb-2-api-comment-service/internal/db/mysql"
	pb "github.com/kodinggo/gb-2-api-comment-service/pb/comment_service"
	grpcHandler "github.com/kodinggo/gb-2-api-comment-service/internal/delivery/grpc"
	httphandler "github.com/kodinggo/gb-2-api-comment-service/internal/delivery/http"
	"github.com/kodinggo/gb-2-api-comment-service/internal/repository"
	"github.com/kodinggo/gb-2-api-comment-service/internal/usecase"
	"google.golang.org/grpc"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)
var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve is a command to run the service server",
	Run: func(cmd *cobra.Command, args []string) {
		dbConn := mysqldb.InitDBConn()
		commentRepository := repository.InitCommentRepository(dbConn)
		commentUseCase := usecase.InitCommentUsecase(commentRepository)
		quitChannel := make(chan bool, 1)
	go	func(){

			e := echo.New()
			commentHandler := httphandler.InitCommentHandler(commentUseCase)
			commentHandler.RegisterRoute(e)
			e.GET("/ping", func(c echo.Context) error {
				return c.String(http.StatusOK, "pong!")
			})
			fmt.Println("PORT :",config.Port())
			e.Start(":" + config.Port())
			}()
			go func() {
				grpcServer := grpc.NewServer()
				commentgRPCHandler := grpcHandler.InitgRPCHanlder(commentUseCase)
				pb.RegisterCommentServiceServer(grpcServer, commentgRPCHandler)
				httpListener, err := net.Listen("tcp", ":7778")
				if err != nil {
					log.Panic("create http listener %w", err)
				}
				log.Println("grpc server running....")
				grpcServer.Serve(httpListener)
				}()
				<-quitChannel
			},
	}
func init() {
	rootCmd.AddCommand(serverCmd)
}
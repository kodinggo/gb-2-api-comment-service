package console

import (
	"log"
	"net"
	"net/http"


	pbAccount "github.com/kodinggo/gb-2-api-account-service/pb/account"

	"github.com/kodinggo/gb-2-api-comment-service/internal/config"
	mysqldb "github.com/kodinggo/gb-2-api-comment-service/internal/db/mysql"
	grpcHandler "github.com/kodinggo/gb-2-api-comment-service/internal/delivery/grpc"
	httphandler "github.com/kodinggo/gb-2-api-comment-service/internal/delivery/http"
	"github.com/kodinggo/gb-2-api-comment-service/internal/repository"
	"github.com/kodinggo/gb-2-api-comment-service/internal/usecase"
	pb "github.com/kodinggo/gb-2-api-comment-service/pb/comment_service"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve is a command to run the service server",
	Run: func(cmd *cobra.Command, args []string) {
		dbConn := mysqldb.InitDBConn()
		commentRepository := repository.InitCommentRepository(dbConn)

		quitChannel := make(chan bool, 1)

		go func() {
			e := echo.New()

			placeholderUseCase := usecase.InitCommentUsecase(commentRepository, nil)

			commentHandler := httphandler.InitCommentHandler(placeholderUseCase)
			commentHandler.RegisterRoute(e)

			e.GET("/ping", func(c echo.Context) error {
				return c.String(http.StatusOK, "pong!")
			})

			log.Println("HTTP server running...")
			if err := e.Start(":" + config.Port()); err != nil {
				log.Fatalf("failed to start HTTP server: %v", err)
			}
		}()

		go func() {
			accountClient := func() pbAccount.AccountServiceClient {
				opts := []grpc.DialOption{
					grpc.WithTransportCredentials(insecure.NewCredentials()),
				}

				conn, err := grpc.NewClient("localhost:6000", opts...)
				if err != nil {
					log.Fatalf("failed to connect to account-service: %v", err)
				}

				return pbAccount.NewAccountServiceClient(conn)
			}()

			commentUseCase := usecase.InitCommentUsecase(commentRepository, accountClient)

			grpcServer := grpc.NewServer()
			commentgRPCHandler := grpcHandler.InitgRPCHanlder(commentUseCase)

			pb.RegisterCommentServiceServer(grpcServer, commentgRPCHandler)

			httpListener, err := net.Listen("tcp", ":7778")
			if err != nil {
				log.Panicf("failed to create HTTP listener: %v", err)
			}
			log.Println("gRPC server running...")
			if err := grpcServer.Serve(httpListener); err != nil {
				log.Fatalf("failed to start gRPC server: %v", err)
			}
		}()

		<-quitChannel

	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

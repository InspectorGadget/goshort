package main

import (
	"github.com/InspectorGadget/goshort/constants"
	"github.com/InspectorGadget/goshort/controllers"
	"github.com/InspectorGadget/goshort/initializers"
	"github.com/InspectorGadget/goshort/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "goshort",
	Short:   "A URL shortening service",
	Long:    `goshort is a URL shortening service that allows users to create short links for their URLs.`,
	Version: constants.Version,
}

func init() {
	godotenv.Load()

	if err := initializers.ConnectToDB(); err != nil {
		panic(err)
	}
}

func main() {
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "start",
			Short: "Start the GoShort server",
			Long:  `Starts the GoShort server, allowing users to create and manage short URLs.`,
			Run: func(cmd *cobra.Command, args []string) {
				startServer()
			},
		},
	)

	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "migrate",
			Short: "Run database migrations",
			Long:  `Runs the database migrations to set up the necessary tables for the GoShort service.`,
			Run: func(cmd *cobra.Command, args []string) {
				if err := initializers.Migrate(); err != nil {
					panic(err)
				}
			},
		},
	)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func startServer() {
	r := gin.Default()
	protected := r.Group("/protected")
	protected.Use(middlewares.Authenticate)

	r.GET("/", controllers.Index)

	// Login routes
	r.POST("/authenticate", controllers.Authenticate)

	// Redirect route
	r.GET("/redirect/:short", controllers.RedirectClient)

	// User routes
	protected.GET("/users", controllers.ListUsers)
	r.POST("/users", controllers.AddUser)
	protected.DELETE("/users/:id", controllers.DeleteUser)

	protected.POST("/users/:id/url", controllers.AddUrlToUser)
	protected.GET("/users/:id/url", controllers.ListUrlByUser)
	protected.DELETE("/users/:id/url/:urlid", controllers.DeleteUrlByUser)

	// Role routes
	protected.POST("/roles", controllers.AddRole)

	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}

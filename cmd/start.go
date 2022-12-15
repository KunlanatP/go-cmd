package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().Uint16VarP(&port, "port", "p", 8080, "Spacific port to start")
}

var port uint16

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start server",
	Long:  `Start server at localhost:8080`,
	Run:   startCommandLine,
}

func startCommandLine(cmd *cobra.Command, args []string) {

	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	router.POST("/user/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action" // true
		c.String(http.StatusOK, "%t", b)
	})

	// This handler will add a new router for /user/groups.
	// Exact routes are resolved before param routes, regardless of the order they were defined.
	// Routes starting with /user/groups are never interpreted as /user/:name/... routes
	router.GET("/user/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "The available groups are [...]")
	})

	router.Run(":8080")

	// app := fiber.New(
	// 	fiber.Config{
	// 		ErrorHandler: errs.ErrorHandler(),
	// 	},
	// )

	// v1Route := app.Group("/api/v1")
	// employee.RegisterRoute(v1Route)

	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)

	// go func() {
	// 	<-c
	// 	fmt.Println("Started server")
	// 	_ = app.Shutdown()
	// }()

	// log.Println("Starting server")

	// if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
	// 	log.Panic(err)
	// }
}

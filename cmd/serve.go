package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpDelivery "github.com/hobord/invst-portfolio-backend-golang/delivery/http"
	persistence "github.com/hobord/invst-portfolio-backend-golang/infrastructure/mysql"
	interactor "github.com/hobord/invst-portfolio-backend-golang/interactors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serverCMD represents the server command
var serverCMD = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Long:  `It is start the program as server mode`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Start server")
		httpPort, _ := cmd.Flags().GetInt("port")
		requestsLog, _ := cmd.Flags().GetBool("verbose")
		log.Printf("Requests log is: %v", requestsLog)

		dbUser, _ := cmd.Flags().GetString("db_user")
		dbPass, _ := cmd.Flags().GetString("db_password")
		dbHost, _ := cmd.Flags().GetString("db_host")
		dbName, _ := cmd.Flags().GetString("db_name")
		dbConnectionSTR := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

		dbConn, err := sql.Open("mysql", dbConnectionSTR)
		if err != nil {
			log.Fatal(err)
		}

		repository := persistence.NewInstrumentMysqlRepository(dbConn)
		interactor := interactor.CreateInstrumentInteractor(repository)

		r := mux.NewRouter()
		httpDelivery.MakeRouting(r, interactor)

		log.Printf("Listen on port: %v", httpPort)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), r))
	},
}

func init() {
	viper.AutomaticEnv() // read in environment variables that match

	rootCmd.AddCommand(serverCMD)
	serverCMD.Flags().IntP("port", "l", viper.GetInt("PORT"), "8080")
	serverCMD.Flags().BoolP("verbose", "v", viper.GetBool("LOG"), "Log requests into stdout")
	serverCMD.Flags().StringP("db_user", "u", viper.GetString("DB_USER"), "Database user")
	serverCMD.Flags().StringP("db_password", "P", viper.GetString("DB_PASSWORD"), "Database password")
	serverCMD.Flags().StringP("db_host", "H", viper.GetString("DB_HOST"), "Database host:port")
	serverCMD.Flags().StringP("db_name", "d", viper.GetString("DB_NAME"), "Database name")
	serverCMD.Flags().StringP("frontend", "f", viper.GetString("FRONTEND"), "Public frontend files direcotry path")
	// serverCMD.Flags().StringP("db_connection", "d", viper.GetString("DB_CONNECTION"), )
}

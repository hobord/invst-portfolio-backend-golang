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
		dbConnectionSTR, _ := cmd.Flags().GetString("db_connection")

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
	serverCMD.Flags().IntP("port", "p", viper.GetInt("PORT"), "8080")
	serverCMD.Flags().IntP("db_connection", "d", viper.GetInt("DB_CONNECTION"), "dbuser:secret@tcp(mysql:3306)/testdb?multiStatements=true")
}

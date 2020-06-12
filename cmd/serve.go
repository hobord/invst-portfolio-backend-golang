package cmd

import (
	"database/sql"
	"fmt"
	"log"

	http "github.com/hobord/invst-portfolio-backend-golang/delivery/http"
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
		allowedOrigins, _ := cmd.Flags().GetStringArray("cors")
		if len(allowedOrigins) == 0 {
			allowedOrigins = append(allowedOrigins, "*")
		}
		log.Printf("AllowedOrigins: %v", allowedOrigins)

		spaDir, _ := cmd.Flags().GetString("frontend")
		log.Printf("SPA dir: %s", spaDir)

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
		http.MakeWebServer(httpPort, allowedOrigins, spaDir, interactor)
	},
}

func init() {
	viper.AutomaticEnv() // read in environment variables that match

	rootCmd.AddCommand(serverCMD)
	serverCMD.Flags().IntP("port", "l", viper.GetInt("PORT"), "Listen on this port, default: 8080")
	serverCMD.Flags().StringArrayP("cors", "c", viper.GetStringSlice("CORS"), "CORS allowed origins You can use multiply this flag. If it is not set then *")
	serverCMD.Flags().StringP("db_user", "u", viper.GetString("DB_USER"), "Database user")
	serverCMD.Flags().StringP("db_password", "P", viper.GetString("DB_PASSWORD"), "Database password")
	serverCMD.Flags().StringP("db_host", "H", viper.GetString("DB_HOST"), "Database host:port")
	serverCMD.Flags().StringP("db_name", "d", viper.GetString("DB_NAME"), "Database name")
	serverCMD.Flags().StringP("frontend", "f", viper.GetString("FRONTEND"), "Public frontend files directory path")
}

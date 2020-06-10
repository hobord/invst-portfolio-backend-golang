package cmd

import (
	"database/sql"
	"fmt"
	"log"

	mysql_repository "github.com/hobord/invst-portfolio-backend-golang/infrastructure/mysql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// migrateCMD represents the migrate command
var migrateCMD = &cobra.Command{
	Use:   "migrate",
	Short: "Init / migrate database",
	Long:  `It is help to manage the database`,
	Run: func(cmd *cobra.Command, args []string) {
		migrations, _ := cmd.Flags().GetString("migrations")
		dbUser, _ := cmd.Flags().GetString("db_user")
		dbPass, _ := cmd.Flags().GetString("db_password")
		dbHost, _ := cmd.Flags().GetString("db_host")
		dbName, _ := cmd.Flags().GetString("db_name")
		dbConnectionSTR := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)
		down, _ := cmd.Flags().GetBool("down")

		dbConn, err := sql.Open("mysql", dbConnectionSTR)
		if err != nil {
			log.Fatal(err)
		}
		if down {
			mysql_repository.MigrationDown(dbConn, migrations)
		} else {
			mysql_repository.MigrationUp(dbConn, migrations)
		}
	},
}

func init() {
	viper.AutomaticEnv() // read in environment variables that match

	rootCmd.AddCommand(migrateCMD)
	migrateCMD.Flags().StringP("db_user", "u", viper.GetString("DB_USER"), "Database user")
	migrateCMD.Flags().StringP("db_password", "P", viper.GetString("DB_PASSWORD"), "Database password")
	migrateCMD.Flags().StringP("db_host", "H", viper.GetString("DB_HOST"), "Database host:port")
	migrateCMD.Flags().StringP("db_name", "d", viper.GetString("DB_NAME"), "Database name")
	migrateCMD.Flags().StringP("migrations", "m", viper.GetString("MIGRATIONS"), "Migrations files path")
	migrateCMD.Flags().BoolP("down", "", viper.GetBool("MIGRATE_DOWN"), "Migrate down")

}

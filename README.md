# Cobra/Viper Utils

## Problem
Often I find myself repeating essentially the same code in my cobra-cli projects.
I want my applications to pull in both flags and environment variables, which is
problematic when essentially the same code is repeated but includes small variations
for various types and I forget to bind both `pflags` and `viper` and have to start
debugging what should be a generally simple pattern.

## Usage

Add to your module or project
```bash
$ go get github.com/paul-nelson-baker/go-cobraviper-utils
```

Import in your root command or subcommand, wherever is most appropriate.
For example, if you had a base command that needed to add persistent flags to `root.go`
and `serve.go` and `migrate.go` subcommands, you might use these utils like this
```go
// root.go

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	cobraviperutils.BindFlag(rootCmd.PersistentFlags(), "postgres-host", "127.0.0.1", "Host or IP of database")
	cobraviperutils.BindFlag(rootCmd.PersistentFlags(), "postgres-port", 5432, "Port of database")
	cobraviperutils.BindFlag(rootCmd.PersistentFlags(), "postgres-password", "", "Password of database")
}
```

```go
// serve.go

func init() {
	rootCmd.AddCommand(serveCmd)
	// Here you will define your flags and configuration settings.
	cobraviperutils.BindFlag(serveCmd.Flags(), "host", "0.0.0.0", "Host to bind server")
	cobraviperutils.BindFlag(serveCmd.Flags(), "port", 8080, "Port to bind server")
}
```

You would then retrieve values like this
```go
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the rest API",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		host := cobraviperutils.GetFlag[string]("host")
		port := cobraviperutils.GetFlag[int]("port")
		log.Printf("Will run on: %s:%d", host, port)
		sqlHost := cobraviperutils.GetFlag[string]("postgres-host")
		sqlPort := cobraviperutils.GetFlag[int]("postgres-port")
		log.Printf("Will use postgres located: %s:%d", sqlHost, sqlPort)
	},
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Updates the database schema",
	Run: func(cmd *cobra.Command, args []string) {
		sqlHost := cobraviperutils.GetFlag[string]("postgres-host")
		sqlPort := cobraviperutils.GetFlag[int]("postgres-port")
		log.Printf("Will use postgres located: %s:%d", sqlHost, sqlPort)
	},
}
```

## Contributing

Most of this project is generated due to the toil intense nature of type permutations.
The templates directory is the place to start if you want to change the output code, but
new types can be added to `generate-permutations.json`.

Generation, formatting, and tidying can be all done with one command

```bash
$ make generate
```

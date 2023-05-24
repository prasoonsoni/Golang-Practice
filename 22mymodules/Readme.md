- We have to always do `go mod init github.com/<username>/<modulename>`

- To install any package from outside do `go get <url>`

- go.sum is used to check for integrity of the package

To run the module
1. go build .
2. go run main.go

Some more commands

1. `go mod tidy` - This command ensures that the go. mod file contains the correct dependencies and versions used in your codebase.

2. `go mod verify` - To verify integrity of the installed packages.

3. `go list all` - To list all the installed packages
    
    - `-m` flag to list packages used in current project 

4. `go list -m -versions github.com/gorilla/mux` - To get all the versions available for a package

5. `go mod why github.com/gorilla/mux` - To check which module is dependent on this package.

6. `go mod graph` - To tell which package is dependent on which package

7. `go mod edit` - To edit go.mod file using flags such as `-go 1.16`, `module 1.16`

8. `go mod vendor` - To bring all the code of packages used in vendor folder.

9. `go run -mod=vendor main.go` - This will first check the code in vendor 

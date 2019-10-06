# Lizard King

### Running

Here are the steps needed in order to run The Lizard King!

  - Install GoLang and ensure it works on your pc - https://golang.org/doc/install
  - Put The Lizard King directory in a new src folder in your gopath
  - Build the program with :
```sh
$ go build
```
- Run the built program
- Connect to the program using (or some other method of connecting to your pc's port):
```sh
$ telnet localhost 8080
```
- Interact from the Window you telnetted to

### Developing

  - Ensure you can run Lizard King 
  - The project is currently broken up into four .go files, however the way go works a function or data member in any of them can be seen by the whole program so these are just for organization
  - To add new commands add a function for it in commands.go, then insert the string to use the command and the name of the function in fillCommands()
  - Any changes to .go files will be added to the program be re-running 
```sh
$ go build
```
  - To print to the User use .Fprint commands, log prints or normal prints will print to the server
  - All the world information is contained in world.sql, to add changes to this file you must open the world.db file in sqlite3 (or a similar program) and either re-execute all the statements in the .sql file or simply execute the new commands you added
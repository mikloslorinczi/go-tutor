# Golang Tutor

A small collection of Go examples

A Tour of Go
https://tour.golang.org/welcome/1

The Go Playground
https://play.golang.org/

Writing Go Code (the traditional way in the GOPATH)
https://golang.org/doc/code.html

Go Modules
https://medium.com/rungo/anatomy-of-modules-in-go-c8274d215c16



# GOPATH
The GOPATH environment variable (one of the go envs, which we will see later) is the heel of Achilles of Go. In the past it was the only path where you could develop Go code. With Go 1.11 this has changed, now you shouldn't develop Go code in your GOPATH (make sense, eh?). The point is you should create a directory named **go** in your home folder and inside create these three subdirectories **bin** **pkg** **src**. The you should add the GOPATH environment variable the value of the path of your go directory (e.g.: export GOPATH=/Users/miki/go) and also modify your PATH to include the go/bin directory (e.g.: export PATH=$PATH:$GOPATH/bin). This needs to be done because Go installs the binaries there by default, and this way you can run them from any working directory.

# Then where should I develop Go code?
Basically outside of your GOPATH. The best practice is to have a folder with your GitHub (or another version control system) username, and every folder is named after your repos. You can create a new empty repo in GitHub and clone it, or you can create a new folder init the Git repo there and push it to GitHub. Anyways inside the repo you can run the **go mod init github.com/username/reponame** command to tell Go to create a new module, dependencies will be writen into the go.mod and go.sum files.

# Using the go command
When we tell the Go tool-chain (the **go** command) to run one (or more .go files) it creates a temporary directory (most likely in the /tmp folder) and builds our code there (we almost cannot notice this, because Go is dessigned to be compile FAST) finally it runs the binary. This is a very easy and convient way to run and test our code. Despite Go is a compiled language this practical and fast behaviour gives us the option to use Go as a scripting language (with a bit of BASH tinkering)

# Go envs and the build command
Issueing **go env** you can examine your current Go environment. These are just simple environment variables, which we can cange anytime. For example: the **go build** command builds our binary (It must be noticed that every Go program, no matter how many .go files it consists of, will be builded into exactly one binary file) The **build** command has a few flags like **-o** to change the binary name (by default it will be named after the folder, in our case go-tutor. This is because the module name is equals to the folders name where it lives, we will see this later). But if we change the GOOS and GOARCH Go envs we can build binaries for different operation systems and processor architectures with the same **go build** command. Check this gist for the complet list https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63

# Go install command
The **go install** command will install your application to GOPATH/bin. And because this folder is added to your PATH you can run your program from any working directory.

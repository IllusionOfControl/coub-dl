# coub-dl
Coub-dl is a Go application that allows you to download videos from the [coub.com](coub.com)

## Build
To build Coub-dl from source, you'll need to have Go and go mod installed on your system. Once you have Go and go mod installed, you can use the following commands to build and run Coub-dl:

```bash
# Clone the repository
git clone https://github.com/IllusionOfControl/coub-dl.git

# Change into the project directory
cd coub-dl

# Build the binary (add output extension .exe for Windows)
go build -o coub-dl ./cmd/main.go 

# Run the binary with the -url argument
./coub-dl -url https://coub.com/view/abc123
```

## Usage
To use Coub-dl, you'll need to provide a Coub URL as an argument. You can also use the loop argument to loop the video a specified number of times.

```shell
coub-dl --url <coub-url> [--loop <n>]
```
For example, to download a Coub video with the URL https://coub.com/view/abc123 and loop it 3 times, you would use the following command:

```shell
coub-dl -url https://coub.com/view/abc123 -loop 3
```
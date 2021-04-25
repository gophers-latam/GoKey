# - Tareas por definir de acuerdo al tipo de software a construir
#	y al entry point.
#---------------------------------------------------

all: win32 win64 linux64

win32:
	GOOS=windows GOARCH=386 go build -o ./build/gokey_x32.exe ./cmd

win64: 
	GOOS=windows GOARCH=amd64 go build -o ./build/gokey_x64.exe ./cmd

linux64:
	GOOS=linux GOARCH=amd64 go build -o ./build/gokey ./cmd

clean:
	rm ./build/gokey_x32.exe && rm ./build/gokey_x64.exe && rm ./build/gokey
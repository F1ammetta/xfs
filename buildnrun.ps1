Set-Location ~/goprojects/dir

go build -o ./build/xfs.exe

Set-Location ./build

gsudo ./xfs.exe


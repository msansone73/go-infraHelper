echo "get gin package"
go get -u github.com/gin-gonic/gin
echo "build "
go build main.go
echo "move for bin folder"
mv main /bin
echo "finish."
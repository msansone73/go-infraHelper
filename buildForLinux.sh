#echo "get gin package"
#go get -u github.com/gin-gonic/gin
echo "build "
go build go-infraHelper
echo "move for bin folder"
mv go-infraHelper /bin
echo "finish."
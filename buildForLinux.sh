#echo "get gin package"
go get github.com/gin-gonic/gin
go get gopkg.in/ini.v1
go get github.com/gin-gonic/autotls
go get golang.org/x/crypto/acme/autocert

echo "build "
go build 
echo "move for bin folder"
mv go-infraHelper /bin
echo "finish."
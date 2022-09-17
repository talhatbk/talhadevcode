<code style="font-size: 10.2375px; font-family: Consolas, Menlo, Monaco, "Courier New", monospace; color: rgba(51, 51, 51, 0.85); background-color: rgb(246, 247, 248); border-radius: 2px; display: block; line-height: 1.5; position: relative; top: 0px; outline: 0px !important;"> 
package main
import (
   "log"
   "net/http"
)
type Server struct{}
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
   w.WriteHeader(http.StatusOK)
   w.Header().Set("Content-Type", "application/json")
   w.Write([]byte(`{"message": "hello world"}`))
}
func main() {
   s := &Server{}
   http.Handle("/", s)
   log.Fatal(http.ListenAndServe(":8080", nil))
}
</code>
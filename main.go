package main




import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    "regexp"
    "strconv"
    "time"
    
)



func findemail(k int,n int, ch chan struct{}){

   for i:=k;i<=n;i++{

   response,err := http.Get("https://jsonplaceholder.typicode.com/posts/"+strconv.Itoa(i)+"/comments")
   
   
   if err!=nil{
      fmt.Println(err)
      os.Exit(1)
   }

   body,err := ioutil.ReadAll(response.Body)

  

   defer response.Body.Close() 

   
   
   
 
   

   if err == nil {
      re  := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)

      match := re.FindAllStringSubmatch(string(body),-1)

      for _, m := range match{
         fmt.Println(m[0])
      }
      
      
   }
  
   }
   ch <- struct{}{}

}
 


func main() {
   ch := make(chan struct{})
   count := 10

   currentTime1:=time.Now()
   
   for i:=1;i<=10;i++{
      k:=i*10
      t:=k-9
      go findemail(t,k,ch)
      
      
   }
   

   

   
   
       for range ch {
        count--
        if count == 0 {
            close(ch)
        }
    }  
     

   
   currentTime2:=time.Now()
   fmt.Println(currentTime2.Sub(currentTime1))


  
}
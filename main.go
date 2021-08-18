package main

import (
   "encoding/json"
   "fmt"
   "kubernetes_webhook/apps"
)

type Cla struct {
   Name string `json:"name"`
   test int
}

type Stu struct {
   Name string `json:"name"`
   Age int `json:"age"`
   sex int
   Cla `json:"Cla"`
}

func run()  {
   //apps.Server()
   stu := Stu{
      Name: "huazhong",
      Age: 28,
      sex: 1,
      Cla: Cla{
         Name: "huaha",
         test: 1,
      },
   }
   data, err := json.Marshal(stu)
   if err != nil{
      fmt.Println(err)
   } else {
      fmt.Println(string(data))
   }

   str_data := `{"name":"huazhong","age":28, "sex": 1}`
   ss := Stu{}
   json.Unmarshal([]byte(str_data), &ss)
   fmt.Println(ss)
}

func main() {
   apps.Server()
}

package main

import (
  "fmt" ;
  "net/http" ;
  "encoding/json" ;
  // "bytes" ;
  // "net/url"

)

var baseUrl = "http://localhost:8080"

type rental struct {
	Id          string `json:"id"`
	Brand       string `json:"brand"`
	Year        int    `json:"year"`
	OwnerId     string `json:"owner_id"`
	RentPrice   int    `json:"rent_per_hour"`
	IsAvailable int    `json:"availability"`
}

type response struct {
  Success bool `json:"success"`
  Message string `json:"message"`
  Data []Rental `json:"data"`
}

func ambil_api()([]rental, error){
   var err error
   var client = &http.Client{}
   // var data []makanan
   var data []rental

   // var param =  url.Values{}
   // param.Set("OwnerId", owner)
   // var payload = bytes.NewBufferString(param.Encode())
   request, err := http.NewRequest("GET", baseUrl+"/mobil", nil)
   // request, err := http.NewRequest("POST", baseUrl+"/rental", payload)
   if err != nil {
     return data, err
   }

   request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   response, err := client.Do(request)
   if err != nil {
     return data, err
   }
   defer response.Body.Close()

   err = json.NewDecoder(response.Body).Decode(&data)
   if err != nil {
     return data, err
   }

   return data, nil
}


func main(){
  // var mobil, err = ambil_api("O11")
  var response, err = ambil_api()
  if err != nil {
    fmt.Println("Error! ", err.Error())
    return
  }

  fmt.Println(response.Message)
  mobil = response.Data

  for _, each := range mobil {
    fmt.Print("ID : ",each.Id, ", Brand : ",each.Brand,", Tahun : ", each.Year )
    switch each.IsAvailable {
    case 1 :
      fmt.Println(", Status: Occupied")
    default :
      fmt.Println(", Status: Available")
    }
  }
  // fmt.Println(menu)
  //   fmt.Println("ID :",menu.Id, ", Menu :", menu.Nama,", Harga :", menu.Harga )

}

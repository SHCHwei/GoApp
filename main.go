package main

import(
    "GoApp/router"
    _ "GoApp/database"
)


func main(){
    router := router.SetupRouter()
    router.Run(":8080")
}
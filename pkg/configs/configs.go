package configs

import(
    "github.com/joho/godotenv"
    "os"
    "log"
)


type MariaDB struct{
    Host string `json:"host"`
    Port string `json:"port"`
    Account string `json:"account"`
    Pwd string `json:"pwd"`
    DB string `json:"db"`
}

var CfgMariaDB MariaDB


func init(){

    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    CfgMariaDB.Host = os.Getenv("DB_HOST")
    CfgMariaDB.Port = os.Getenv("DB_PORT")
    CfgMariaDB.DB = os.Getenv("DB_DATABASE")
    CfgMariaDB.Account = os.Getenv("DB_USERNAME")
    CfgMariaDB.Pwd = os.Getenv("DB_PASSWORD")

}


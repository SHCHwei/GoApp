package jwt

import(
    "github.com/golang-jwt/jwt/v5"
    "strconv"
    "time"
    "log"
)


type Claims struct {
    id string
    jwt.RegisteredClaims
}


const (
    secret string = "jaomoqhoi45frjih8r2q8"
)

func Build(id int, account string) (string, error){

    mid := strconv.Itoa(id)

    claims := Claims{
            mid,
            jwt.RegisteredClaims{
                ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
                IssuedAt:  jwt.NewNumericDate(time.Now()),
                NotBefore: jwt.NewNumericDate(time.Now()),
                Issuer: "GoApp",
                Subject: account,
            },
        }


    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString([]byte(secret))

    return tokenString, err
}

func Verify(tokenString string)(bool){

    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })

    if err != nil {
        log.Printf("%v", err)
        return false
    } else if _, ok := token.Claims.(*Claims); ok {
        return true
    } else {
        log.Fatal("unknown claims type, cannot proceed")
        return false
    }
}
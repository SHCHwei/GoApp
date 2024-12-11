package login

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "GoApp/pkg/jwt"
    "GoApp/pkg/hashing"
    "GoApp/member"
)


type Login struct{
    Account string `json:"account"`
    Pwd string `json:"password"`
}

var LoginInfo Login

func(st Login) UnSignIn(c *gin.Context){
    s := hashing.HashPassword("4654a9gr6ag")
    c.JSON(200, gin.H{"token": s})
}

func(st Login) SignIn(c *gin.Context){

    var err error

    if err = c.ShouldBindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"request error": err.Error()})
		return
	}

    st.Pwd = hashing.HashPassword(st.Pwd)

    if memberInfo, ok := member.FindOne(member.Member{Account: st.Account, PW : st.Pwd }); ok {
        if token, _ := jwt.Build(memberInfo); token != ""{
            c.Set("mid", memberInfo.Id)
            c.Set("account", memberInfo.Account)

            c.JSON(200, gin.H{"token": token})
        }else{
            c.JSON(403, gin.H{"error": "Login Failed(token error)"})
            return
        }

    } else{
        c.JSON(403, gin.H{"error": "Login Failed"})
        return
    }
}
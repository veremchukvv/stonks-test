package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	vkapi "github.com/go-vk-api/vk"
	"github.com/labstack/echo/v4"
	"github.com/veremchukvv/stonks-test/internal/config"
	"github.com/veremchukvv/stonks-test/pkg/jwt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//var cfg = struct {
//	Port int
//}{
//	Port: 8080,
//}

var TT struct {
	MovieList *template.Template
	Login     *template.Template
	Register  *template.Template
}

var OauthConf = &oauth2.Config{
	ClientID:     os.Getenv("CLIENT_ID"),
	ClientSecret: os.Getenv("CLIENT_SECRET"),
	RedirectURL:  os.Getenv("REDIRECT_URL"),
	Scopes:       []string{""},
	Endpoint:     vk.Endpoint,
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type VKuser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Movie struct {
	Category    int    `json:"category"`
	ID          int    `json:"id"`
	ReleaseDate string `json:"releasedate"`
	Title       string `json:"title"`
	MovieUrl    string `json:"movie_url,omitempty"`
}

type LoginPage struct {
	User  User
	Error string
	URL      string
	OauthLogin bool
}

type RegisterPage struct {
	PageId string
	User   User
	Error  string
	URL      string
	OauthLogin bool
}

type ListMovieResponse struct {
	PageNum  int   `json:"pagenum"`
	PageSize int32 `json:"pagesize"`
	Movies   *[]Movie
}

type MainPage struct {
	User     User
	PageNum  int   `json:"pagenum"`
	PageSize int32 `json:"pagesize"`
	Movies   *[]Movie
	URL      string
	OauthLogin bool
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"displayname"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginResponse struct {
	JWT string `json:"JWT"`
}

func main() {

	cfg := config.GetConfig()

	router := echo.New()

	router.GET("/", MainHandler)
	router.GET ("/login", LoginFormHandler)
	router.POST("/login", LoginHandler)
	router.POST("/oauth", OauthHandler)
	router.POST("/logout", LogoutHandler)
	router.GET("/portfolio", PortfolioHandler)

	//router.HandleFunc("/register", RegisterFormHandler).Methods("GET")
	//router.HandleFunc("/register", RegisterHandler).Methods("POST")

	// Настройка шаблонизатора

	var err error

	TT.MovieList, err = template.ParseFiles("template/layout/base.html", "template/main.html")
	if err != nil {
		log.Fatal(err)
	}

	TT.Register, err = template.ParseFiles("template/layout/base.html", "template/register.html")
	if err != nil {
		log.Fatal(err)
	}

	TT.Login, err = template.ParseFiles("template/layout/base.html", "template/login.html")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("server start at port: %v", cfg.ClientServer.Port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+cfg.ClientServer.Port, router))

}

func MainHandler(ctx echo.Context) error {
	http.Redirect(ctx.Response(), ctx.Request(), "/portfolio", http.StatusFound)
	return nil
}

func PortfolioHandler(ctx echo.Context) error {

	page := MainPage{}

	//isGeneralLogin, err := r.Cookie("jwt")
	isOauthLogin, err := ctx.Request().Cookie("oauthToken")
	oauthUserName, err := ctx.Request().Cookie("oauthUserName")

	if isOauthLogin == nil {
		page.URL = fmt.Sprintf("https://oauth.vk.com/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s&state=%s", OauthConf.ClientID, OauthConf.RedirectURL, OauthConf.Scopes, "state")
		page.OauthLogin = false
	} else {
		page.URL = ""
		page.OauthLogin = true
	}

	if oauthUserName != nil {
		page.User.Name = oauthUserName.Value
	}

	//page.Movies, err = listMovies()
	//if err != nil {
	//	log.Printf("Get movie error: %v", err)
	//}

	page.User, err = getUserByToken(ctx.Request())
	if err != nil {
		log.Printf("Get user error: %v", err)
	}

	err = TT.MovieList.ExecuteTemplate(ctx.Response(), "base", page)
	if err != nil {
		log.Printf("Render error: %v", err)
		ctx.Response().WriteHeader(http.StatusInternalServerError)
	}
	return nil
}

func LoginFormHandler(ctx echo.Context) error {
	page := &LoginPage{}

	isOauthLogin, err := ctx.Request().Cookie("oauthToken")
	oauthUserName, err := ctx.Request().Cookie("oauthUserName")

	if isOauthLogin == nil {
		page.URL = fmt.Sprintf("https://oauth.vk.com/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s&state=%s", OauthConf.ClientID, OauthConf.RedirectURL, OauthConf.Scopes, "state")
		page.OauthLogin = false
	} else {
		page.URL = ""
		page.OauthLogin = true
	}

	if oauthUserName != nil {
		page.User.Name = oauthUserName.Value
	}

	page.User, err = getUserByToken(ctx.Request())
	if err != nil {
		log.Printf("No user: %v", err)
		//В случае не валидного токена показываем страницу логина
		TT.Login.ExecuteTemplate(ctx.Response().Writer, "base", page)
		return nil
	}
	TT.Login.ExecuteTemplate(ctx.Response().Writer, "base", page)
	return nil
}

func LoginHandler(ctx echo.Context) error {
	page := &LoginPage{}

	isOauthLogin, err := ctx.Request().Cookie("oauthToken")
	oauthUserName, err := ctx.Request().Cookie("oauthUserName")

	if isOauthLogin == nil {
		page.URL = fmt.Sprintf("https://oauth.vk.com/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s&state=%s", OauthConf.ClientID, OauthConf.RedirectURL, OauthConf.Scopes, "state")
		page.OauthLogin = false
	} else {
		page.URL = ""
		page.OauthLogin = true
	}

	if oauthUserName != nil {
		page.User.Name = oauthUserName.Value
	}

	ctx.Request().ParseForm()
	email := ctx.Request().PostFormValue("email")
	pwd := ctx.Request().PostFormValue("pwd")

	req := &LoginRequest{Email: email, Password: pwd}
	jsonReq, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	jsonRes, err := http.Post("http://localhost:8001/login", "application/json", bytes.NewBuffer(jsonReq))

	body, err := ioutil.ReadAll(jsonRes.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(body))

	response := LoginResponse{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println(err)
	}

	// Что-то не так с сервисом user
	if err != nil {
		log.Printf("Get user error: %v", err)
		page.Error = "Сервис авторизации не доступен"
		TT.Login.ExecuteTemplate(ctx.Response(), "base", page)
		return nil
	}

	// Ошибка логина, ее можно показать пользователю
	if jsonRes.StatusCode != 200 {
		page.Error = jsonRes.Status
		TT.Login.ExecuteTemplate(ctx.Response(), "base", page)
		return nil
	}

	tok := response.JWT

	// Если пользователь успешно залогинен записываем токен в cookie

	http.SetCookie(ctx.Response(), &http.Cookie{Name: "jwt", Value: tok})

	jwtData, err := jwt.Parse(tok)
	if err != nil {
		// В случае не валидного токена показываем страницу логина
		log.Println("token is invalid")
		TT.Login.ExecuteTemplate(ctx.Response(), "base", page)
		return nil
	}
	log.Println(jwtData)
	log.Println(jwtData.Name)

	page.User = User{Name: jwtData.Name}
	log.Printf("%v+", page)

	TT.Login.ExecuteTemplate(ctx.Response(), "base", page)
	return nil
}

func LogoutHandler(ctx echo.Context) error {
	http.SetCookie(ctx.Response(), &http.Cookie{Name: "jwt", MaxAge: -1})
	http.Redirect(ctx.Response(), ctx.Request(), "/login", http.StatusFound)
	return nil
}

//func RegisterFormHandler(ctx echo.Context) error {
//	page := &RegisterPage{PageId: "register"}
//	page.User = User{
//		Name: "",
//	}
//
//	isOauthLogin, err := ctx.Request().Cookie("oauthToken")
//	if err != nil {
//		log.Println(err)
//	}
//
//	oauthUserName, err := ctx.Request().Cookie("oauthUserName")
//	if err != nil {
//		log.Println(err)
//	}
//
//	if isOauthLogin == nil {
//		page.URL = fmt.Sprintf("https://oauth.vk.com/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s&state=%s", OauthConf.ClientID, OauthConf.RedirectURL, OauthConf.Scopes, "state")
//		page.OauthLogin = false
//	} else {
//		page.URL = ""
//		page.OauthLogin = true
//	}
//
//	if oauthUserName != nil {
//		page.User.Name = oauthUserName.Value
//	}
//
//	TT.Register.ExecuteTemplate(ctx.Response(), "base", page)
//	return nil
//}
//
//func RegisterHandler(ctx echo.Context) error {
//	page := &RegisterPage{}
//
//	page.URL = ""
//
//	ctx.Request().ParseForm()
//	email := ctx.Request().PostFormValue("email")
//	name := ctx.Request().PostFormValue("name")
//	age := ctx.Request().PostFormValue("age")
//	tel := ctx.Request().PostFormValue("telephone")
//	pwd := ctx.Request().PostFormValue("pwd")
//
//	ageInt, err := strconv.Atoi(age)
//	if err != nil {
//		log.Println(err)
//	}
//
//	req := &RegisterRequest{Email: email, Name: name, Age: ageInt, Phone: tel, Password: pwd}
//	jsonReq, err := json.Marshal(req)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	log.Println(string(jsonReq))
//	jsonRes, err := http.Post("http://localhost:8081/register", "application/json", bytes.NewBuffer(jsonReq))
//
//	body, err := ioutil.ReadAll(jsonRes.Body)
//	if err != nil {
//		log.Println(err)
//	}
//	log.Println(string(body))
//
//	response := LoginResponse{}
//
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		log.Println(err)
//	}
//
//	// Что-то не так с сервисом user
//	if err != nil {
//		log.Printf("Get user error: %v", err)
//		page.Error = "Сервис авторизации не доступен"
//		TT.Login.ExecuteTemplate(ctx.Response(), "base", page)
//		return nil
//	}
//
//	// Ошибка логина, ее можно показать пользователю
//	if jsonRes.StatusCode != 200 {
//		page.Error = jsonRes.Status
//		TT.Login.ExecuteTemplate(ctx.Response(), "base", page)
//		return nil
//	}
//
//	tok := response.JWT
//
//	// Если пользователь успешно залогинен записываем токен в cookie
//
//	http.SetCookie(ctx.Response(), &http.Cookie{Name: "jwt", Value: tok})
//
//	jwtData, err := jwt.Parse(tok)
//	if err != nil {
//		// В случае не валидного токена показываем страницу логина
//		log.Println("token is invalid")
//		TT.Login.ExecuteTemplate(ctx.Response(), "base", page)
//		return nil
//	}
//
//	page.User = User{Name: jwtData.Name}
//
//	TT.Login.ExecuteTemplate(ctx.Response(), "base", page)
//	return nil
//}

func OauthHandler(ctx echo.Context) error {
	ctxt := context.Background()

	code := ctx.Request().URL.Query().Get("code")

	var tokOauth, err = OauthConf.Exchange(ctxt, code)
	if err != nil {
		log.Fatal(err)
	}

	vkClient, err := vkapi.NewClientWithOptions(vkapi.WithToken(tokOauth.AccessToken))
	if err != nil {
		log.Fatal(err)
	}

	vkUser := getCurrentOauthUser(vkClient)

	expire := time.Now().Add(10 * time.Minute)

	http.SetCookie(ctx.Response(), &http.Cookie{Name: "oauthUserName", Value: vkUser.FirstName, Expires: expire})
	http.SetCookie(ctx.Response(), &http.Cookie{Name: "oauthToken", Value: code, Expires: expire})

	http.Redirect(ctx.Response(), ctx.Request(), "/movies", http.StatusFound)
return nil
}

var ERR_NO_JWT = errors.New("No 'jwt' cookie")

func getUserByToken(r *http.Request) (u User, err error) {
	tok, err := r.Cookie("jwt")
	if tok == nil {
		return u, ERR_NO_JWT
	}

	jwtData, err := jwt.Parse(tok.Value)
	if err != nil {
		return u, fmt.Errorf("Can't parse token: %w", err)
	}

	u.Name = jwtData.Name
	return u, err
}

//func listMovies() (*[]Movie, error) {
//
//	client := &http.Client{}
//
//	url, err := http.NewRequest("GET", "http://localhost:8001/movies?limit=1", nil)
//	if err != nil {
//		log.Println(err)
//	}
//
//	url.Header.Add("Accept", "application/json")
//	jsonRes, err := client.Do(url)
//	if err != nil {
//		log.Printf("error here", err)
//	}
//	log.Println(jsonRes)
//
//	body, err := ioutil.ReadAll(jsonRes.Body)
//	if err != nil {
//		log.Printf("error here1", err)
//	}
//	log.Println("no payload!" + string(body))
//
//	response := ListMovieResponse{}
//
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		log.Println(err)
//	}
//	log.Println(response.Movies)
//
//	return response.Movies, nil
//}

func getCurrentOauthUser(api *vkapi.Client) VKuser {
	var users []VKuser
	api.CallMethod("users.get", vkapi.RequestParams{}, &users)
	return users[0]
}

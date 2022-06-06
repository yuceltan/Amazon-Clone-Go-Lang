package handlers

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"bitbucket.org/janpavtel/site/internal/drivers"
	"bitbucket.org/janpavtel/site/internal/forms"
	"bitbucket.org/janpavtel/site/internal/mails"
	"bitbucket.org/janpavtel/site/internal/models"
	"bitbucket.org/janpavtel/site/internal/repository"
	"bitbucket.org/janpavtel/site/internal/repository/dbrepo"
	"bitbucket.org/janpavtel/site/internal/tokens"
	"github.com/go-chi/chi/v5"
	"github.com/steambap/captcha"
	"golang.org/x/crypto/bcrypt"
)

type TemplateData struct {
	Form *forms.Form
	Data map[string]interface{}
}

type View struct {
	DB repository.DatabaseRepo
	MailChan chan mails.MailData
	Client *http.Client
}

func NewView(db *drivers.DB, mailChain chan mails.MailData) *View{
	return &View{
		DB: dbrepo.NewPostgresRepo(db.SQL),
		MailChan: mailChain,
		Client: httpClient(),
	}
}



func (view *View) Home(w http.ResponseWriter, r *http.Request){
	err := renderTemplate(w, "home.tpl", nil)
	if err != nil {
		log.Println("Can't render home page ", err)
	}
}

func (view *View) About(w http.ResponseWriter, r *http.Request){
	err := renderTemplate(w, "about.tpl", nil)
	if err != nil {
		log.Println("Can't render about page ", err)
	}
}

func (view *View) SignUp(w http.ResponseWriter, r *http.Request){
	err := renderTemplate(w, "signup.tpl", nil)
	if err != nil {
		log.Println("Can't render sing-up page ", err)
	}
}

func (view *View) ShowUsers(w http.ResponseWriter, r *http.Request){
	users, err := view.DB.LoadAllUsers()
	if err != nil {
		log.Println("Can't load users", err)
		return 
	}


	data := make(map[string]interface{})
	data["users"] = users

	err = renderTemplate(w, "users.tpl", &TemplateData{Data: data})
	if err != nil {
		log.Println("Can't render users page ", err)
		return
	}
}

func (view *View) ShowUser(w http.ResponseWriter, r *http.Request){
	userID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	user, err := view.DB.LoadUserById(userID)
	if err != nil {
		log.Println("Can't load user", err)
		return 
	}


	data := make(map[string]interface{})
	data["user"] = user

	err = renderTemplate(w, "user.tpl", &TemplateData{Data: data})
	if err != nil {
		log.Println("Can't render users page ", err)
		return
	}
}

func (view *View) EditUser(w http.ResponseWriter, r *http.Request){
	userID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	user, err := view.DB.LoadUserById(userID)
	if err != nil {
		log.Println("Can't load user", err)
		return 
	}


	data := make(map[string]interface{})
	data["user"] = user

	err = renderTemplate(w, "user-edit.tpl", &TemplateData{Data: data})
	if err != nil {
		log.Println("Can't render user edit page ", err)
		return
	}
}

func (view *View) UpdateUser(w http.ResponseWriter, r *http.Request){
	log.Println("UpdateUser func ")

	userID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := r.ParseForm()
	if err != nil {
		log.Println("Can't parse form ", err)
		return
	}

	form := forms.New(r.PostForm)
	form.CheckMinLengh("first_name", 3)
	form.CheckRequiredFields("email", "last_name") 


	if form.IsNotValid() {
		renderTemplate(w, "user-edit.tpl", &TemplateData{
			Form: form,
		})

		return
	}

	newUser := models.User{
		ID: userID,
		FirstName: r.Form.Get("first_name"),
		LastName: "Fix Me!",
		Email: r.Form.Get("email"),
	}

	err = view.DB.UpdateUser(newUser)
	if err != nil {
		log.Println("Can't add new user", err)
		return
	}

	http.Redirect(w, r, "/users/"+strconv.Itoa(userID), http.StatusSeeOther)
}



func (view *View) NewUserCreation(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		log.Println("Can't parse form ", err)
		return
	}

	form := forms.New(r.PostForm)
	form.CheckMinLengh("first_name", 3)
	form.CheckRequiredFields("email", "last_name", "password") 


	if form.IsNotValid() {
		renderTemplate(w, "signup.tpl", &TemplateData{
			Form: form,
		})

		return
	}

	encryptedPassword, err := encryptPassword(r.Form.Get("password"))
	if err != nil {
		log.Println("Can't encrypt password", err)
		renderTemplate(w, "signup.tpl", &TemplateData{
			Form: form,
		})

		return
	}

	newUser := models.User{
		FirstName: r.Form.Get("first_name"),
		LastName: r.Form.Get("last_name"),
		Email: r.Form.Get("email"),
		Password: encryptedPassword,

	}

	id, err := view.DB.AddUser(newUser)
	if err != nil {
		log.Println("Can't add new user", err)
		return
	}

	http.Redirect(w, r, "/users/"+strconv.Itoa(id), http.StatusSeeOther)
}

func (view *View) UserLoginShow(w http.ResponseWriter, r *http.Request){
	err := renderTemplate(w, "profile/login.tpl", nil)
	if err != nil {
		log.Println("Can't render login page ", err)
	}
}

const captchaCookieName = "mvccaptcha"
func (view *View) GenerateCaptcha(w http.ResponseWriter, r *http.Request){
	data, err := captcha.New(150, 50)
	if err != nil {
		log.Println("can't generate captcha ", err)
		return 
	}

	hashedCaptcha := hashSha256(data.Text)

	log.Println("generated captcha: ", data.Text)

	http.SetCookie(w, &http.Cookie{
		Name: captchaCookieName,
		Value: hashedCaptcha,
		Expires: time.Now().Add(5 * time.Minute),
		Path: "/",
	})

	data.WriteImage(w)
}
func (view *View) UserLoginPost(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		log.Println("Can't parse form ", err)
		return
	}

	form := forms.New(r.PostForm)
	form.CheckEmail("email")
	form.CheckRequiredFields("email", "password", "captcha") 

	captcha := r.Form.Get("captcha")

	hashedCaptcha := hashSha256(captcha)
	cookie, err := r.Cookie(captchaCookieName)
	if err != nil {
		log.Println("Error extracting cokie", err)
		return
	}

	expectedHashCaptcha := cookie.Value
	if hashedCaptcha != expectedHashCaptcha {
		form.Errors.Add("captcha", "Invalid Captcha")
	}


	if form.IsNotValid() {
		renderTemplate(w, "profile/login.tpl", &TemplateData{
			Form: form,
		})

		return
	}

	user, err := view.DB.LoadUserByEmail(r.Form.Get("email"))
	if err != nil {
		log.Println("Error loading user ", r.Form.Get("email"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Form.Get("password")))
	if err != nil {
		log.Println("Incorrect password for user ", r.Form.Get("email"))

		form.Errors.Add("password", "incorrect email or password")
		renderTemplate(w, "profile/login.tpl", &TemplateData{
			Form: form,
		})

		return
	}

	err = tokens.GenerateAndAddTokenToCookie(user.ID, user.Roles, w)
	if err != nil {
		log.Println("Can't generate token ", err)
		return
	}

	htmlMessage := fmt.Sprintf(`
	<strong>Detected New Login</strong><br>
	at %s.
	`, time.Now().Format("2006-01-02"))

	msg := mails.MailData{
		To: user.Email,
		From: "no-replay@example.com",
		Subject: "New Login - Site",
		Content: htmlMessage,
	}

    select {
		case view.MailChan <- msg: // Put msg in the channel unless it is full
		default:
			log.Println("Channel full. Discarding value")
	}



	http.Redirect(w, r, "/posts/allposts", http.StatusSeeOther)
}

func (view *View) PostsCreateShow(w http.ResponseWriter, r *http.Request){
	err := renderTemplate(w, "posts/create.tpl", nil)
	if err != nil {
		log.Println("Can't render posts create page ", err)
	}
}

func (view * View) PostsCreatePosts(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		log.Println("Can't parse form ", err)
		return
	}

	form := forms.New(r.PostForm)
	form.CheckRequiredFields("title", "body") 

	if form.IsNotValid() {
		renderTemplate(w, "posts/create.tpl", &TemplateData{
			Form: form,
		})

		return
	}

	userId, err  := strconv.Atoi(r.Context().Value("userId").(string))
	if err != nil {
		log.Println("can't retrieve userId from request's context", err)
		return
	}

	post := models.Post{
		Title: r.Form.Get("title"),
		Body: r.Form.Get("body"),
		OwnerId: userId,
	}

	_, err = view.DB.AddPost(post)
	if err != nil {
		log.Println("failed to add post ", err)
		return
	}


	user, err  := view.DB.LoadUserById(userId)
	if err != nil {
		log.Println("can't load user", err)
		return
	}
	htmlMessage := fmt.Sprintf(`
	<strong>New Post Created</strong><br>
	at %s.
	`, time.Now().Format("2006-01-02"))

	msg := mails.MailData{
		To: user.Email,
		From: "no-replay@example.com",
		Subject: "New Post - Site",
		Content: htmlMessage,
	}

    select {
		case view.MailChan <- msg: // Put msg in the channel unless it is full
		default:
			log.Println("Channel full. Discarding value")
	}


	http.Redirect(w, r, "/posts/create", http.StatusSeeOther)
	
}


func (view * View) PostsEditShow(w http.ResponseWriter, r *http.Request){
	postId, _ := strconv.Atoi(chi.URLParam(r, "id"))

	post, err := view.DB.LoadPostById(postId)
	if err != nil {
		log.Println("Can't load post ", err)
		return
	}

	data := make(map[string]interface{})
	data["post"] = post 

	err = renderTemplate(w, "posts/post-edit.tpl", &TemplateData{Data: data})
	if err != nil {
		log.Println("Can't render post edit page ", err)
		return
	}
	
}

func (view *View) PostsUpdate(w http.ResponseWriter, r *http.Request){
	postId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("Can't extract postID", err)
		return
	}

	err = r.ParseForm()
	if err != nil {
		log.Println("Can't parse form ", err)
		return
	}

	form := forms.New(r.PostForm)
	form.CheckRequiredFields("title", "body") 

	if form.IsNotValid() {
		renderTemplate(w, "posts/post-edit.tpl", &TemplateData{
			Form: form,
		})

		return
	}

	post := models.Post{
		ID: postId,
		Title: r.Form.Get("title"),
		Body: r.Form.Get("body"),
	}

	err = view.DB.UpdatePost(post)
	if err != nil {
		log.Println("can't update post", err)
		return
	}

	http.Redirect(w, r, "/posts/"+strconv.Itoa(postId)+"/edit", http.StatusSeeOther)
}



func (view *View) PostsMyPosts(w http.ResponseWriter, r *http.Request){
	userID, err  := strconv.Atoi(r.Context().Value("userId").(string))
	if err != nil {
		log.Println("can't retrieve userId from request's context", err)
		return
	}


	myposts, err := view.DB.LoadAllPostsByOwnerID(userID)
	if err != nil {
		log.Println("can't load myposts", err)
		return
	}

	data := make(map[string]interface{})
	data["posts"] = myposts

	user, err := view.DB.LoadUserById(userID)
	if err != nil {
		log.Println("can't load user", user)
		return
	}
	data["user"] = user

	err = renderTemplate(w, "posts/allmyposts.tpl", &TemplateData{
		Data: data,
	})
	if err != nil {
		log.Println("Can't render posts all my posts page ", err)
	}
}



func (view *View) PostsAllPosts(w http.ResponseWriter, r *http.Request){
	limit, err := extractIntParam(r, "limit", 2)
	if err != nil {
		log.Println("can't parse limit", err)
		return
	}
	offset, err := extractIntParam(r, "offset", 0)
	if err != nil {
		log.Println("can't parse offset", err)
		return
	}

	posts, err := view.DB.LoadAllPosts(limit, offset)
	if err != nil {
		log.Println("can't load myposts", err)
		return
	}

	data := make(map[string]interface{})
	data["posts"] = posts
	

	nextPagePagination := &models.Pagination{
		Limit: limit,
		Offset: offset+limit,
	}
	data["nextPagePagination"] = nextPagePagination


	roles := r.Context().Value("roles").([]string)
	user := &models.User{
		Roles: roles,
	}
	data["user"] = user


	err = renderTemplate(w, "posts/allposts.tpl", &TemplateData{
		Data: data,
	})
	if err != nil {
		log.Println("Can't render posts all posts page ", err)
	}
}


func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}


const techcrunchNewsEndpoint = "https://techcrunch.com/wp-json/wp/v2/posts?per_page=5&context=embed"
func (view *View) NewsAll(w http.ResponseWriter, r *http.Request){
	
	req, err := http.NewRequest("GET", techcrunchNewsEndpoint, nil)
	if err != nil {
		log.Println("can't load data ", err)
		return
	}

	response, err := view.Client.Do(req)
	if err != nil {
		log.Println("Can't execute request ", err)
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Can't parse response body ", err)
		return
	}

	var news []models.News
	if err := json.Unmarshal(body, &news); err != nil {
		log.Println("Can't unmarsh JSON with news", err)
		return 
	}
	

	data := make(map[string]interface{})
	data["news"] = news
	


	err = renderTemplate(w, "news/all.tpl", &TemplateData{
		Data: data,
	})
	if err != nil {
		log.Println("Can't render news all page ", err)
	}
}

func (view *View) ManagementGetUsers(w http.ResponseWriter, r *http.Request){
	users, err := view.DB.LoadAllUsers()
	if err != nil {
		log.Println("Can't load users", err)
		return 
	}


	data := make(map[string]interface{})
	data["users"] = users

	err = renderTemplate(w, "admin/users.tpl", &TemplateData{Data: data})
	if err != nil {
		log.Println("Can't render admin users page ", err)
		return
	}

}

func (view *View) ManagementEditUser(w http.ResponseWriter, r *http.Request){
	userID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	user, err := view.DB.LoadUserById(userID)
	if err != nil {
		log.Println("Can't load user", err)
		return 
	}


	data := make(map[string]interface{})
	data["user"] = user

	err = renderTemplate(w, "admin/user-edit.tpl", &TemplateData{Data: data})
	if err != nil {
		log.Println("Can't render user edit page ", err)
		return
	}
}

func (view *View) ManagementUpdateUser(w http.ResponseWriter, r *http.Request){
	log.Println("UpdateUser func ")

	userID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := r.ParseForm()
	if err != nil {
		log.Println("Can't parse form ", err)
		return
	}

	form := forms.New(r.PostForm)
	form.CheckMinLengh("first_name", 3)
	form.CheckRequiredFields("email", "last_name") 


	if form.IsNotValid() {
		renderTemplate(w, "admin/user-edit.tpl", &TemplateData{
			Form: form,
		})

		return
	}

	newUser := models.User{
		ID: userID,
		FirstName: r.Form.Get("first_name"),
		LastName: r.Form.Get("last_name"),
		Email: r.Form.Get("email"),
	}

	err = view.DB.UpdateUser(newUser)
	if err != nil {
		log.Println("Can't update  user", err)
		return
	}

	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}

func (view *View) UserProfileEditShow(w http.ResponseWriter, r *http.Request){
	userID, err  := strconv.Atoi(r.Context().Value("userId").(string))
	if err != nil {
		log.Println("can't retrieve userId from request's context", err)
		return
	}

	user, err := view.DB.LoadUserById(userID)
	if err != nil {
		log.Println("Can't load user", err)
		return 
	}


	data := make(map[string]interface{})
	data["user"] = user

	err = renderTemplate(w, "/profile/profile-edit.tpl", &TemplateData{Data: data})
	if err != nil {
		log.Println("Can't render profile edit page ", err)
		return
	}
}

func (view *View) UserProfileEditUpdate(w http.ResponseWriter, r *http.Request){
	userID, err  := strconv.Atoi(r.Context().Value("userId").(string))
	if err != nil {
		log.Println("can't retrieve userId from request's context", err)
		return
	}

	err = r.ParseForm()
	if err != nil {
		log.Println("Can't parse form ", err)
		return
	}

	form := forms.New(r.PostForm)
	form.CheckMinLengh("first_name", 3)
	form.CheckRequiredFields("email", "last_name") 


	if form.IsNotValid() {
		renderTemplate(w, "/profile/profile-edit.tpl", &TemplateData{
			Form: form,
		})

		return
	}

	newUser := models.User{
		ID: userID,
		FirstName: r.Form.Get("first_name"),
		LastName: r.Form.Get("last_name"),
		Email: r.Form.Get("email"),
	}

	err = view.DB.UpdateUser(newUser)
	if err != nil {
		log.Println("Can't add new user", err)
		return
	}

	http.Redirect(w, r, "/users/profile/edit", http.StatusSeeOther)
}

func (view *View) UserProfileUploadImage(w http.ResponseWriter, r *http.Request){
	file, fileHeader, err := r.FormFile("profile_image")
	if err != nil {
		log.Println("Wrong upload request", err)
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err = io.Copy(buf, file); err != nil {
		log.Println("Can't read file content", err)
		return 
	}

	userID, err  := strconv.Atoi(r.Context().Value("userId").(string))
	if err != nil {
		log.Println("can't retrieve userId from request's context", err)
		return
	}

	err = view.DB.UpdateUserProfilePicture(userID, fileHeader.Filename, buf)
	if err != nil {
		log.Println("can't insert picture ", err)
		return
	}

	http.Redirect(w, r, "/users/profile/edit", http.StatusSeeOther)
}

func (view *View) UserProfileImage(w http.ResponseWriter, r *http.Request){
	userID, err  := strconv.Atoi(r.Context().Value("userId").(string))
	if err != nil {
		log.Println("can't retrieve userId from request's context", err)
		return
	}

	buf, err := view.DB.LoadProfilePicture(userID)
	if err != nil {
		log.Println("Can't load profile picture", err)
		return 
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(buf.Bytes())
}


func extractIntParam(r *http.Request, paramName string, defaultValue int) (int, error) {
	if valueStr := r.URL.Query().Get(paramName); valueStr != ""{
		value, err := strconv.Atoi(valueStr)
		return value, err
	} 

	return defaultValue, nil
}



func renderTemplate(w http.ResponseWriter, templateName string, templateData *TemplateData) error{
	return renderTemplateWithLayout(w, "base.layout.tpl", templateName, templateData)
}

func renderTemplateWithLayout(w http.ResponseWriter, layoutName string, templateName string, templateData *TemplateData) error{
	if templateData == nil {
		templateData = &TemplateData{}
	}

	parsedTemplate, errTmp := template.ParseFiles("./www/pages/"+templateName)
	if errTmp != nil {
		return errTmp
	}
	resolvedTemplate, errLay := parsedTemplate.ParseFiles("./www/layouts/"+layoutName)
	if errLay != nil {
		return errLay
	}


	buf := new(bytes.Buffer)


	errBufData := resolvedTemplate.Execute(buf, templateData)
	if errBufData != nil {
		return errBufData
	}

	_, errExe := buf.WriteTo(w)
	if errExe != nil {
		return errExe
	}

	return nil
}

const bycryptCost = 12
func encryptPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bycryptCost)
    return string(bytes), err
}

func hashSha256(text string) string {
	data := []byte(text)
	hash := sha256.Sum256(data)
	return base64.StdEncoding.EncodeToString(hash[:])
}


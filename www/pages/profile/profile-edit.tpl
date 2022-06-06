{{template "base" .}} 





{{define "content"}}

{{$user := index .Data "user"}}

<div id='navBar'>
  <div id='topHalf'>
    <div id='logoWrapper'>
    <img id='logo' src="http://www.userlogos.org/files/logos/ArkAngel06/Amazon.png" />
    </div>
    <form class="example" action="action_page.php">
    <input type="text" placeholder="Search.." name="search">
    <button type="submit"><i class="fa fa-search"></i></button>
  </form>
    
  </div>

</div>
            <ul>  
              <li><a class="active" href="#home">Menu</a></li>  
              <li><a href="#">Prime</a></li>  
              <li><a href="http://localhost:8080/">Strona główna</a></li>  
             <li><a href="https://www.amazon.pl/gp/bestsellers?ref_=nav_cs_bestsellers">Bestsellery</a></li>  
              <li><a href="http://localhost:8080/posts/create">Twój produkt</a></li>  
              <li><a href="https://sell.amazon.pl/?ld=AZPLSOANavbar">Sprzedawaj na Amazon</a></li>  
              <li><a href="http://localhost:8080/users/profile/edit">użytkownik</a></li>  

            </ul>  

        


            
 <center><h1> Logowanie i bezpieczeństwo </h1></center>
 <br>
 
<center>
<div class ="loginpagediv">
  <div class="container">
    <img src="/users/profile/image" alt="profile image" width="150" height="150"> <br/>

      <form class = "loginpageform" action="/users/profile/image/upload" method="post" enctype="multipart/form-data">
        <input type="file" name="profile_image" multiple/>
        <br>
        <br>
        <input type="submit" value="Zaktualizuj obraz">
      </form>
    </div>

    <br>
    <br>


    <div class="container">
      <form  class = "loginpageform" action="/users/profile/edit" method="post">
        <label for="first_name">Nazwa użytkownika::</label>
        <input type="text" id="first_name" name="first_name" value="{{$user.FirstName}}" autocomplete="off">

        <label for="last_name">Nazwisko::</label>
        <input type="text" id="last_name" name="last_name" value="{{$user.LastName}}" autocomplete="off">

        <label for="email">Adres e-mail::</label>
        <input type="text" id="email" name="email" value="{{$user.Email}}" autocomplete="off">
        <br>

       <input type="submit" value="Gotowe">
      </form>
  </div>
  <br>
  

</div>
</center>


{{if .Form}}
<ul class="error">
{{ range $key, $value := .Form.Errors }}
   <li><strong>{{ $key }}</strong>: {{ $value }}</li>
{{ end }}
<ul>
{{end}}

{{end}}




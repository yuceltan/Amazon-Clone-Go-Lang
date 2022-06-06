{{template "base" .}} 

{{define "title"}} {{end}}



{{define "content"}}



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
              <li><a href="#">Wiosenne okazje</a></li>  
             <li><a href="https://www.amazon.pl/gp/bestsellers?ref_=nav_cs_bestsellers">Bestsellery</a></li>  
              <li><a href="http://localhost:8080/posts/create">Twój produkt</a></li>  
              <li><a href="https://sell.amazon.pl/?ld=AZPLSOANavbar">Sprzedawaj na Amazon</a></li>  
              <li><a href="http://localhost:8080/users/profile/edit">użytkownik</a></li>  

            </ul>  
<center>


  <div class="container">
    <img src="/users/profile/image" alt="profile image" width="150" height="150"> <br/>

      <form class = "loginpageform" action="/users/profile/image/upload" method="post" enctype="multipart/form-data">
        <input type="file" name="profile_image" multiple/>
        <br>
        <br>
        <input type="submit" value="Zaktualizuj obraz">
      </form>
    


  <div>
    <center><h1> Twój produkt </h1></center>

    <form action="/posts/create" method="post">
      <label for="title">Title:</label>
      <br>
      <input type="text" id="title" name="title" autocomplete="off">
      <br>

      <label for="body">Cechy produktu:</label>
      <br>
      <textarea type="text" id="body" name="body" autocomplete="off" rows="5" cols="33"> </textarea>
      <br>
      <input type="submit" value="Udział">
    </form>
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




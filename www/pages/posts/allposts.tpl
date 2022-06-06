{{template "base" .}} 

{{define "title"}} 



{{end}}



{{define "content"}}

<div id='navBar'>
  <div id='topHalf'>
    <div id='logoWrapper'>
    <img id='logo' src="http://www.userlogos.org/files/logos/ArkAngel06/Amazon.png" />
    </div>
    <form class="example" action="http://localhost:8080/posts/allposts">
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


{{$user := index .Data "user"}}

{{$posts := index .Data "posts"}}


    {{range $post := $posts}}
    <strong>{{$post.Title}}</strong>
    <p>{{$post.Body}}</p>
        {{if  $user.HaveOneOfRoles "ADMIN" "EDITOR" }}
            <a href="/posts/{{$post.ID}}/edit">Edit</a>
            <br/>
        {{end}}
    {{end}}


    <hr/>
    {{$nextPagePagination := index .Data "nextPagePagination"}}
    <a href="/posts/allposts">First Page </a> <br/>
    <a href="/posts/allposts?limit={{$nextPagePagination.Limit}}&offset={{$nextPagePagination.Offset}}">Next Page</a>


{{end}}
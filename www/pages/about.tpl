{{template "base" .}} 

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

{{end}}



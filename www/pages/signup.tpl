{{template "base" .}} 

{{define "title"}}  {{end}}


<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">

{{define "content"}}
<form class = "loginpageform" action="/users" method="post">
<center>
  <a href="http://localhost:8080/">
          <img alt="Qries" src="https://www.synointcdn.com/wp-content/uploads/2019/04/Amazon-Logo-PNG.png"
          width=160" height="70">
      </a>

      </center>

<center>

  <div class = "loginpagediv">
    <center><h1> Utwórz konto </h1></center>
    <center>
      <label for="first_name">Nazwa użytkownika:</label>
      <br>
      <input type="text" id="first_name" name="first_name" autocomplete="off">
      <br>
      <br>
      <label for="last_name">Nazwisko:</label>
      <input type="text" id="last_name" name="last_name" autocomplete="off">
      <br>
      <br>
      <label for="email">Adres e-mail:</label>
      <br>
      <input type="text" id="email" name="email" autocomplete="off">
      <br>
      <br>
      <label for="password">Hasło:</label>
      <br>
      <input type="password" id="password" name="password" autocomplete="off">
      <br>
      <br>
      <center>
      
    <label for="deal">Zarejestruj się teraz, aby otrzymywać e-maile z powiadomieniami o nowych produktach, pomysłach na prezenty, specjalnych okazjach, promocjach i nie tylko.
Wyrażam zgodę na otrzymywanie informacji marketingowych, w tym ofert i promocji, od Amazon za pomocą środków komunikacji elektronicznej takich jak e-mail. Możesz dostosować zakres lub zrezygnować z otrzymywanych od nas wiadomości, poprzez zakładkę Moje Konto lub w ustawieniach powiadomień w swojej aplikacji lub na urządzeniu mobilnym.</label></center>
      <br>
      <br>
      <img src="/users/login/captcha.jpg" alt="captcha" width="150" height="50"> <br/>
 	<label for="captcha">Wprowadź znaki, które widzisz:</label>
  	<input type="text" id="captcha" name="captcha" autocomplete="off">
	 <br>
         <br>
		
      <input type="submit" value="Utwórz konto Amazon">
      <br>
      <br>
      <p>By creating an account, you agree to Amazon's Conditions of Use and Privacy Notice.</p>
      <p>Posiadasz już konto?  <a href="/users/login">Zaloguj się</a>
      </p>

      </center>
  </div>
</center>
</form>






{{if .Form}}
<ul class="error">
{{ range $key, $value := .Form.Errors }}
   <li><strong>{{ $key }}</strong>: {{ $value }}</li>
{{ end }}
<ul>
{{end}}

{{end}}




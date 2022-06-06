{{template "base" .}} 


{{define "title"}}  {{end}}



{{define "content"}}

<center>
    <a href="http://localhost:8080/">
            <img alt="Qries" src="https://www.synointcdn.com/wp-content/uploads/2019/04/Amazon-Logo-PNG.png"
            width=160" height="70">
        </a>

        </center>
<center>
  <div class = "loginpagediv">
    <center><h1> Zaloguj się </h1></center>
      <center>     
      <form class = "loginpageform "action="/users/login" method="post">

        <label for="email">E-mail lub numer telefonu komórkowego:</label>
        <input type="text" id="email" name="email" autocomplete="off">

        <label for="password">Hasło:</label>
        <input type="password" id="password" name="password" autocomplete="off">
        <br>
	<img src="/users/login/captcha.jpg" alt="captcha" width="150" height="50"> <br/>
	<label for="captcha">Wprowadź znaki, które widzisz</label>
  	<input type="text" id="captcha" name="captcha" autocomplete="off">

        <br>

        <input type="submit" value="Dalej">
      </form>
      <p>
          Logując się wyrażasz zgodę na <a href="/gp/help/customer/display.html/ref=ap_signin_notification_condition_of_use?ie=UTF8&amp;nodeId=201909000">Warunki użytkowania i sprzedaży</a>  Amazon. Zobacz <a href="/gp/help/customer/display.html/ref=ap_signin_notification_privacy_notice?ie=UTF8&amp;nodeId=201909010">Informację o Prywatności</a> , <a href="/gp/help/customer/display.html/?nodeId=201890250">Informację o Plikach Cookies</a> oraz <a href="/gp/help/customer/display.html?nodeId=201909150">Informację o Reklamach dopasowanych do zainteresowań</a>
      </p>
      <br>

        <a href="/password">Potrzebujesz pomocy?</a> 
      <br>
     
    </center>
  </div>
  <h5>Pierwszy raz w serwisie Amazon?</h5>

<form class ="special" action="/signup" method="get">
    <input type="submit" value="Utwórz konto Amazon" 
         name="Submit" id="frm1_submit" />
</form>

</center>


{{if .Form}}
<ul class="error">
{{ range $key, $value := .Form.Errors }}
   <li><strong>{{ $key }}</strong>: {{ $value }}</li>
{{ end }}
<ul>
{{end}}

{{end}}




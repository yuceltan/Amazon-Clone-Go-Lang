{{template "base" .}} 

{{define "title"}} User Edit {{end}}



{{define "content"}}


{{$user := index .Data "user"}}

<form action="/admin/users/{{$user.ID}}" method="post">
  <label for="first_name">First Name:</label>
  <input type="text" id="first_name" name="first_name" value="{{$user.FirstName}}" autocomplete="off">

  <label for="last_name">Last Name:</label>
  <input type="text" id="last_name" name="last_name" value="{{$user.LastName}}" autocomplete="off">

  <label for="email">Email:</label>
  <input type="text" id="email" name="email" value="{{$user.Email}}" autocomplete="off">

  <input type="submit" value="Update">
</form>


{{if .Form}}
<ul class="error">
{{ range $key, $value := .Form.Errors }}
   <li><strong>{{ $key }}</strong>: {{ $value }}</li>
{{ end }}
<ul>
{{end}}

{{end}}




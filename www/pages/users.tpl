{{template "base" .}} 

{{define "title"}} Users {{end}}



{{define "content"}}

{{$users := index .Data "users"}}

<table border="1">
    <tr>
      <th>First Name</th>
      <th>Last Name</th>
      <th>Email</th>
    </tr>
    {{range $users}}
        <tr>
            <td>{{.FirstName}}</td>
            <td>{{.LastName}}</td>
            <td>{{.Email}}</td>
        </tr>
    {{end}}

</table>


{{end}}
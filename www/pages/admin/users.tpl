{{template "base" .}} 

{{define "title"}} Admin - Users {{end}}



{{define "content"}}
{{$users := index .Data "users"}}

<table class="table">
  <thead>
    <tr>
      <th scope="col">First</th>
      <th scope="col">Last</th>
      <th scope="col">Email</th>
      <th scope="col">Actions</th>
    </tr>
  </thead>
  <tbody>
  {{range $users}}
  <tr>
      <td>{{.FirstName}}</td>
      <td>{{.LastName}}</td>
      <td>{{.Email}}</td>
      <td><a href="/admin/users/{{.ID}}">Edit</a></td>
  </tr>
 {{end}}
  </tbody>
</table>







{{end}}



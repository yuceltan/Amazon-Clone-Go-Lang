{{template "base" .}} 

{{define "title"}} User {{end}}



{{define "content"}}

{{$user := index .Data "user"}}

<table>
    <tr>
        <td>Name</td>
        <td>{{$user.FirstName}} {{$user.LastName}}</td>
    </tr>
    <tr>
        <td>Email</td>
        <td>{{$user.Email}}</td>
    </tr>
</table>


{{end}}
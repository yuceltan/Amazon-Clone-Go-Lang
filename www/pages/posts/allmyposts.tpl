{{template "base" .}} 

{{define "title"}} 

{{$user := index .Data "user"}}

{{$user.FirstName}} {{$user.LastName}} - Posts 

{{end}}



{{define "content"}}

{{$posts := index .Data "posts"}}


    {{range $posts}}
    <strong>{{.Title}}</strong>
    <p>{{.Body}}</p>
    {{end}}


{{end}}
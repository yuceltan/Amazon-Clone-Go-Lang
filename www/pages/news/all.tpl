{{template "base" .}} 

{{define "title"}} 

All News 

{{end}}



{{define "content"}}


{{$news := index .Data "news"}}


    {{range $post := $news}}
    <strong>{{$post.Title.Rendered}}</strong> <br>
    <strong><a href="{{$post.Link}}">{{$post.ID}}</a></strong>
    

    <p>{{$post.Excerpt.Rendered}}</p>
    
    {{end}}

{{end}}
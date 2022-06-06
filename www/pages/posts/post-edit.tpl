{{template "base" .}} 

{{define "title"}} Edit Post {{end}}



{{define "content"}}

{{$post := index .Data "post"}}

<form action="/posts/{{$post.ID}}" method="post">
  <label for="title">Title:</label>
  <br>
  <input type="text" id="title" name="title" autocomplete="off" value="{{$post.Title}}">
  <br>

  <label for="body">Tell As story:</label>
  <br>
  <textarea type="text" id="body" name="body" autocomplete="off" rows="5" cols="33">{{$post.Body}} </textarea>
  <br>
  <input type="submit" value="Update Post">
</form>


{{if .Form}}
<ul class="error">
{{ range $key, $value := .Form.Errors }}
   <li><strong>{{ $key }}</strong>: {{ $value }}</li>
{{ end }}
<ul>
{{end}}

{{end}}




{{define "navbar"}}
<a class="navbar-brand" href="/">匿名社区</a>
<div>
    <ul class="nav navbar-nav">
        <li {{if .Ismessage}}class="active"{{end}}><a href="/">动态</a></li>
        <li {{if .Isconnect}}class="active"{{end}}><a href="/">联系作者</a></li>
    </ul>
</div>

<div class="pull-right">
    <ul class="nav navbar-nav">
    {{if .Islogin}}
    <li><a href="/login?exit=true">tuichu</a></li>
    {{else}}
    <li><a href="/login">denglu</a></li>
    {{end}}
    </ul>
</div>
{{end}}
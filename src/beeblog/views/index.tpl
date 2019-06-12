<!DOCTYPE html>
<html lang="zh-CN">
  {{template "header.tpl" .}}

  <body>

    <div class="blog-masthead">
      <div class="container">
        <nav class="blog-nav">
          <a class="blog-nav-item active" href="/">Home</a>          
        </nav>
      </div>
    </div>

    <div class="container">

      <div class="blog-header">
        {{if .IsTag}}
         <h1 class="blog-title">标签：{{.Tag.Name}} 文章列表</h1>
        {{end}}
      </div>

      <div class="row">

        <div class="col-sm-8 blog-main">
          {{template "list.tpl" .}}
          {{template "pager.tpl" .}}

        </div><!-- /.blog-main -->

        {{template "right.tpl" .}}

      </div><!-- /.row -->

    </div><!-- /.container -->

    <footer class="blog-footer">
      <p>Blog built by <a href="https://github.com/wyr6512">@wyr6512</a>.</p>
      <p>
        <a href="#">Back to top</a>
      </p>
    </footer>


  </body>
</html>

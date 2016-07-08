<!DOCTYPE html>

<html>
<head>
	<style type="text/css">@charset "UTF-8";[ng\:cloak],[ng-cloak],[data-ng-cloak],[x-ng-cloak],.ng-cloak,.x-ng-cloak,.ng-hide{display:none !important;}ng\:form{display:block;}.ng-animate-block-transitions{transition:0s all!important;-webkit-transition:0s all!important;}.ng-hide-add-active,.ng-hide-remove{display:block!important;}</style>
	<title>独孤影 - {{{.title}}}</title>
	<meta name="viewport" content="width=device-width, initial-scale=1"/>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<meta content="独孤影,博客,{{{.keywords}}}" name="keywords" />
	<meta content="{{{.description}}}" name="description" />
	<link rel="EditURI" type="application/rsd+xml" title="RSD" href="{{{.host}}}/xmlrpc" />
	<link rel="shortcut icon" href="/favicon.ico" />


	
	<link rel="stylesheet" href="../../static/css/style.css">
	<link rel="stylesheet" href="../../static/css/css.css">
	<link rel="stylesheet" href="../../static/css/handheld.css">
	<script src="../../static/js/jquery_003.js"></script>
	<script>window.jQuery || document.write("<script src='static/js/libs/jquery-1.5.1.min.js'>\x3C/script>")</script>
	<script src="../../static/js/modernizr-1.js"></script>
	<script src="../../static/js/easing.js" type="text/javascript"></script>
	<script src="../../static/js/jquery.js" type="text/javascript"></script>
	<script type="text/javascript">
		$(document).ready(function() {     
			$().UItoTop({ easingType: 'easeOutQuart' });      
		});
	</script> 
	<meta name="google-site-verification" content="ohMjRPHv0sKAahvl1H0GC7Dx0-z-zXbMNnWBfxp2PYY" />
	<meta name="baidu-site-verification" content="h3Y69jNgBz" />
</head>
<body >
	<div id="container">

		<header id="top">
			<ul id="social">
				<li class="twitter"><a href="#">Twitter</a></li><li class="rss"><a href="#">RSS</a></li><li class="facebook"><a href="#">Facebook</a></li>  
			</ul>
			<p class="subscribers">Subscribers <span class="subscribersno">2175</span></p>
			<p> 
				<a href="#" rel="home">
					<img src="../../static/img/logo.png" class="logo" alt="logo" height="69" width="299">  
				</a>
			</p>
		</header>

		<nav>
			<div id="access" role="navigation">
				<div class="menu"><ul><li class="current_page_item"><a href="/">Home</a></li><li><a href="#">About The Tests</a></li><li><a href="#">Clearing Floats</a></li><li><a href="#">Level 1</a><ul class="children"><li><a href="#">Level 2</a><ul class="children"><li><a href="#">Level 3</a></li></ul></li></ul></li><li><a href="#">Lorem Ipsum</a></li><li><a href="#">Parent page</a><ul class="children"><li><a href="#">Child page 1</a></li><li><a href="#">Child page 2</a></li></ul></li></ul></div>
			</div><!-- #access -->

			<form id="navsearchform" role="search" method="get" action="#">
				<input value="Search Here" onFocus="if (this.value == 'Search Here') {this.value = '';}" onBlur="if (this.value == '') {this.value = 'Search Here';}" name="s" id="navs" type="text">
				<input id="navsearchsubmit" value="" type="submit">
			</form>
		</nav>




		<div class="article-list">
			<div class="article" itemscope itemtype="http://schema.org/Article">
				<div class="post">

					<div class="post_title" ><h1>{{{.title}}}</h1></div>
					<div class="post_date">Posted on <time class="entry-date" datetime="2008-09-04T23:02:20+00:00" pubdate="">{{{.time}}}</time> by <a href="#">{{{.author}}}</a>  view <span title="{{{.count}}}次阅读" class="view-count">{{{.count}}}</span>
					</div>

					<div class="post_body">
						{{{str2html .content}}}

					</div>

					<div class="post_meta">
						Tag {{{.keywords|tags|str2html}}}

					</div>

				</div>



				{{{template "commentlist.tpl" .}}}

			</div>




			{{{template "inc/footer.tpl" .}}}
		</div>
	</body>
	</html>










<div id="main" role="main" class="clearfix">
	<div id="left">
		<div id="sliderwrap"><script type="text/javascript" src="../../static/js/jquery_002.js"></script>
			<script type="text/javascript">
				jQuery('#featured_slider ul').cycle({ 
					fx: 'fade',
					prev: '.feat_prev',
					next: '.feat_next',
					speed:  3000, 
					timeout: 2000, 
					pager:  null
				});
			</script>

			<style>

				#featured_slider {
					float: left;
					margin: 0px 0px;
					position: relative;
					border: 0px solid;
					width: 575px;
					height:266px;
					overflow:hidden;
				}

				#featured_slider ul, #featured_slider ul li {
					list-style: none !important;
					border: none !important;
					float: left;
					margin: 0px;
					width: 575px;
					height: 266px;
				}



				#featured_slider h2{
					position:absolute; bottom:0px; left:0px;
					width:575px;
					background-color:#000000;
					color:white;
					height:50px;
					line-height:47px;
					padding-left:10px;
				}

				#featured_slider a{
					color:#FFFFFF;
					font-weight:normal;
					font-family: Georgia, "Times New Roman", Times, serif;
					font-size: 24px;
					font-weight: normal;
					letter-spacing: -0.5px;
					width:575px;
				}

				#featured_slider .feat_prev {
					background: transparent url(../../static/img/sprite.png) no-repeat;
					background-position: 0px 0px;
					width: 17px;
					z-index: 10;
					height: 16px;
					position: absolute;
					left: 20px;
					cursor: pointer;
					top: 30px;
					float: left;
				}

				#featured_slider .feat_prev:hover {
					background-position: 0px -16px;
				}

				#featured_slider .feat_next {
					background: transparent url(../../static/img/sprite.png) no-repeat;
					background-position: -17px 0px;
					width: 17px;
					z-index: 10;
					height: 16px;
					position: absolute;
					left: 40px;
					top: 30px;
					cursor: pointer;
				}

				#featured_slider .feat_next:hover {
					background-position: -18px -16px;
				}

				.feat_link {
					float: right;
					position: relative;
					top: -5px;
				}

				.feat_link a {
					float: left;
					font-size: 20px;
					color: #CCC;
				}

			</style>

			<div id="featured_slider">
				<ul style="position: relative;" id="slider">							
					<li><h2><a href="#">Welcome to Busby – This is the Layout Test</a></h2><a href="#"><img src="../../static/img/skater.jpg"></a></li>

					<li><h2><a href="#">Look here for a Readability Test</a></h2><a href="#"><img src="../../static/img/skater-in-air.jpg"></a></li>

					<li><h2><a href="#">How about an Images Test ?</a></h2><a href="#"><img src="../../static/img/skaters.jpg"></a></li>

					<li><h2><a href="#">Comment Test</a></h2><a href="#"><img src="../../static/img/south-bank-graffiti.jpg"></a></li>

					<li><h2><a href="#">Many Tags Many Tags Many Tags</a></h2><a href="#"><img src="../../static/img/spray-paint.jpg"></a></li>
				</ul>
				<div class="feat_next"></div>
				<div class="feat_prev"></div>
			</div>
		</div>

		{{{range $k,$v := .articles_in_page}}}
		<article>

			<div class="post">
				<header>
					<h2 class="posttitle" title="{{{$v.title}}}"><a href="/article/{{{$v.uri}}}" rel="bookmark">{{{$v.title}}}</a></h2>

				</header>

				<div class="postdate"><p><span class="postdateno">{{{$v.year}}}</span><span class="postdatemonth">{{{$v.month}}}</span></p></div>
				<div class="postcontent">
					<p>{{{str2html $v.content1}}}</p>

				</div>

				<div class="postdetails">
					<p class="postedby"><span class="sep">Posted on </span><a href="#" rel="bookmark"><time class="entry-date" datetime="2008-09-04T23:02:20+00:00" pubdate="">{{{$v.time}}}</time></a> <span class="sep"> by </span> <span class="author vcard"><a class="url fn n" href="#author/chip-bennett">{{{$v.author}}}</a></span></p><p class="postcomments"><span title="{{{$v.count}}}次阅读" class="view-count">view {{{$v.count}}}</span></p>
				</div>
			</div>

		</article>

		<hr>
		{{{end}}}


		{{{if .prev_page_flag}}}
		<a href="{{{.prev_page}}}" class="page-nav">上一页</a>
		{{{end}}}
		{{{if .next_page_flag}}}
		<a href="{{{.next_page}}}" class="page-nav">下一页</a>
		{{{end}}}

	</div>

{{{template "inc/rightbar.tpl" .}}}
</div>





	<h3 class="left">Comments</h3>
	<div class="clearer">&nbsp;</div>
	<div class="comment_list">
	
		<ul>
			{{{range $k,$v := .commentlist}}}
			<div class="comment alt">
				<ul>
					<div class="comment_gravatar left">
						<img alt="" src="../../static/img/sample-gravatar.jpg" height="32" width="32" />
					</div>
					<div class="comment_author left">
						{{{ $v.author}}}
						<div class="comment_date"><a href="#">{{{date_cn $v.time}}}</a></div>
					</div>
					<div class="clearer">&nbsp;</div>
					<div class="comment_body">									
						<p>{{{ $v.content}}}</p>
					</div>


				</ul>
			</div>
			{{{end}}}
		</ul>


	</div>
	<form action="comment/add/{{{ .id}}}" method="post" id="reply">

		<fieldset>	

			<div class="legend">Leave a Reply</div>

			<div class="form_row">

				<div class="form_property form_required"><label for="name">Your name</label></div>
				<div class="form_value"><input type="text" size="32" name="name" value="" class="text" id="name" /></div>

				<div class="clearer">&nbsp;</div>

			</div>
			<div class="form_row">

				<div class="form_property form_required"><label for="comment">Comment</label></div>
				<div class="form_value"><textarea rows="10" cols="46" name="comment" id="comment"></textarea></div>

				<div class="clearer">&nbsp;</div>

			</div>

			<div class="form_row form_row_submit">

				<div class="form_value"><input type="submit" class="button" value="Submit Comment &#187;" /></div>

				<div class="clearer">&nbsp;</div>

			</div>

		</fieldset>

	</form>
	
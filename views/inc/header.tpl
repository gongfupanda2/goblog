



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
    <div class="menu"><ul><li class="current_page_item"><a href="#">Home</a></li><li><a href="#">About The Tests</a></li><li><a href="#">Clearing Floats</a></li><li><a href="#">Level 1</a><ul class="children"><li><a href="#">Level 2</a><ul class="children"><li><a href="#">Level 3</a></li></ul></li></ul></li><li><a href="#">Lorem Ipsum</a></li><li><a href="#">Parent page</a><ul class="children"><li><a href="#">Child page 1</a></li><li><a href="#">Child page 2</a></li></ul></li></ul></div>
  </div><!-- #access -->

  <form id="navsearchform" role="search" method="get" action="#">
    <input value="Search Here" onFocus="if (this.value == 'Search Here') {this.value = '';}" onBlur="if (this.value == '') {this.value = 'Search Here';}" name="s" id="navs" type="text">
    <input id="navsearchsubmit" value="" type="button"  onclick="check();" >
  </form>
</nav>

<script>

function check() { 

 window.location.href="/search/" + document.getElementById("navs").value + "/1";
 
}
</script>


<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<link rel="stylesheet" href="/static/css/main.css" type="text/css" media="screen, projection" /> <!-- main stylesheet -->
<link rel="stylesheet" type="text/css" media="all" href="/static/css/tipsy.css" /> <!-- Tipsy implementation -->

<!--[if lt IE 9]>
	<link rel="stylesheet" type="text/css" href="/static/css/ie8.css" />
<![endif]-->

<script type="text/javascript" src="/static/js/jquery-1.7.2.min.js"></script> <!-- uiToTop implementation -->
<script type="text/javascript" src="/static/js/custom-scripts.js"></script>
<script type="text/javascript" src="/static/js/jquery.tipsy.js"></script> <!-- Tipsy -->

<script type="text/javascript">

$(document).ready(function(){
			
	universalPreloader();
						   
});

$(window).load(function(){

	//remove Universal Preloader
	universalPreloaderRemove();
	
	rotate();
    dogRun();
	dogTalk();
	
	//Tipsy implementation
	$('.with-tooltip').tipsy({gravity: $.fn.tipsy.autoNS});
						   
});

</script>


<title>404 - Not found</title>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" /></head>

<body>

<!-- Universal preloader -->
<div id="universal-preloader">
    <div class="preloader">
        <img src="/static/img/universal-preloader.gif" alt="universal-preloader" class="universal-preloader-preloader" />
    </div>
</div>
<!-- Universal preloader -->

<div id="wrapper">
<div class="dog-wrapper">
<!-- dog running -->
	<div class="dog"></div>
<!-- dog running -->
	
<!-- dog bubble talking -->
	<div class="dog-bubble">
    	
    </div>
    
    <!-- The dog bubble rotates these -->
    <div class="bubble-options">
    	<p class="dog-bubble">
        	迷路了404，页面..页面找不到了！
        </p>
        <p class="dog-bubble">
        	点击<a onclick="history.go(-1)">返回</a>上一页或者<a href="/login">主页</a>
        </p>
        <p class="dog-bubble">
        	唉?下次要注意！ @_@
        </p>
    </div>
    <!-- The dog bubble rotates these -->
<!-- dog bubble talking -->
</div>

<!-- planet at the bottom -->
	<div class="planet"></div>
<!-- planet at the bottom -->
</div>


</body>
</html>
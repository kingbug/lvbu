<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<script type="text/javascript" src="/static/js/jquery-1.7.2.min.js"></script>
<title>503</title>
</head>
<body>
<h1>没有访问该资源的权限，<span id="sec">5</span>秒后跳转到<a id="redirect" onclick="history.go(-1)" style="color:blue;cursor:pointer">上个页面</a></h1>
</body>
<script>

$(document).ready(function(){
	var sec = $("#sec").text();
	setInterval(load,1000);
	function load(){
		
		if (sec < 0 ) {
			$("#redirect").click();
			
		}else {
			$("#sec").text(sec);
			sec = sec - 1;
		}
		
	}
})

</script>
</html>
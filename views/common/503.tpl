<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<script type="text/javascript" src="/static/js/jquery-1.7.2.min.js"></script>
<title>503</title>
</head>
<body>
<h1>没有访问该资源的权限，<span id="sec">5</span>秒后跳转到<a id="redirect" onclick="destination();" style="color:blue;cursor:pointer">上个页面</a></h1>
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

function destination() {
		
		var dest=GetQueryString("destination");
		if(dest !=null && dest.toString().length>1)
		{
		   location.href=dest;
		} else {
			history.go(-1);
		}
	}
function GetQueryString(name)
{
     var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)");
     var r = window.location.search.substr(1).match(reg);
     if(r!=null)return  unescape(r[2]); return null;
}
 

</script>
</html>
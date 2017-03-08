var socket;
var ms_modal_contr = false //实时信息打印框控制值，false 为隐藏
var deadcontainers = [];  //存放状态为停止的容器的列表.如为false 淡黄色闪烁提醒
var _interval ;//定时器
$(document).ready(function () {
    // Create a socket
//    socket = new WebSocket('ws://' + window.location.host + '/ws/join?uname=' + $('#uname').text());
    socket = new WebSocket('ws://' + window.location.host + '/ws/nodedeploy?proid=' + pro_id);
    // Message received on the socket
	$(".ms_modal .box-title").text("实时信息")
    socket.onmessage = function (event) {
		
		var data = JSON.parse(event.data);
        console.log(data);
        switch (data.Type) {
        case 0: // EVENT_STAT
			if (deadcontainers.length == 0) {
				deadcontainers.push(data);
			}
			var isexits = true
			$.each(deadcontainers,function(k, v){
				if (v.Containerid == data.Containerid) {
					if(v.Containerstats != data.Containerstats ){
						$("#nodelist td." + data.Containerid).find('i.fa-pause').removeClass('fa-pause').addClass('fa-play');
						$("#nodelist td." + data.Containerid).find('a.pause').removeClass('pause').addClass('play');
						$("#nodelist td." + data.Containerid).find('i.fa-play').removeClass('fa-play').addClass('fa-pause');
						$("#nodelist td." + data.Containerid).find('a.play').removeClass('play').addClass('pause');
						$("#nodelist td." + data.Containerid).parent().removeClass("bg-yellow");
					}
					v.Containerstats = data.Containerstats
					isexits = false
					return false// break
				}
			});
			if (isexits == true) {
				deadcontainers.push(data);
			}
            if (data.Containerstats == false) {//containerid
				console.log("false");
				$("#nodelist td." + data.Containerid).find('i.fa-pause').removeClass('fa-pause').addClass('fa-play');
				$("#nodelist td." + data.Containerid).find('a.pause').removeClass('pause').addClass('play');
            } else {
            }
            break;
        case 1: // EVENT_ERROR
			if (ms_modal_contr == false){
				$(".ms_modal").show(1000);
				ms_modal_contr = true;
			}
            $(".ms_modal .box-body").html($(".ms_modal .box-body").html() + '<span style=\'color:red;\'>' + data.Error + '</span><br/>')
            break;
        case 2: // EVENT_MESSAGE
			if (ms_modal_contr == false){
				$(".ms_modal").show(1000);
				ms_modal_contr = true;
			}
            $(".ms_modal .box-body").html($(".ms_modal .box-body").html() + data.Message + '<br/>')
            break;
        }

    };
	socket.onclose = function (event) {
		console.log("socket closed");
		clearInterval(_interval);
	}

    // Send messages.
    var postConecnt = function () {
        socket.send(JSON.stringify({
			
		}));
    }
	
	function flash() {
		$.each(deadcontainers,function(k, v){
			if (v.Containerstats) {
				return true;
			}
			if ($("td." + v.Containerid).parent().hasClass("bg-yellow")) {
				$("td." + v.Containerid).parent().removeClass("bg-yellow")
			} else {
				$("td." + v.Containerid).parent().addClass("bg-yellow")
			}
		});
	}
	
	
    _interval = setInterval(flash, 500);
});
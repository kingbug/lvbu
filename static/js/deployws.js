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
				console.log("k:" + k)
				if (v.Containerid == data.Containerid) {
					if(v.Containerstats != data.Containerstats ){
						if (data.Containerstats == 0) {
							$("#nodelist td." + data.Containerid).find('i.fa-pause').removeClass('fa-pause').addClass('fa-play');
							$("#nodelist td." + data.Containerid).find('a.pause').removeClass('pause').addClass('play');
							$("#nodelist td." + data.Containerid).find('i.fa-play').removeClass('fa-play').addClass('fa-pause');
							$("#nodelist td." + data.Containerid).find('a.play').removeClass('play').addClass('pause');
							$("#nodelist td." + data.Containerid).parent().removeClass("bg-yellow");
						}
						if (v.Containerstats == 2){ //否则如果之前是不存在状态时
							$("#nodelist td." + data.Containerid).parent().removeClass("bg-red");
							var td_a_list = $("." + v.Containerid).children();
							td_a_list.each(function(k,a){
								if ($(a).hasClass("rocket") || $(a).hasClass("trash")) {
									return true; //continue
								} else {
									$(a).removeClass("disabled")
								}
							});
						}
						deadcontainers.splice(k, 1, data);
					}
					
					isexits = false
					return false// break
				}
			});
			if (isexits == true) {
				deadcontainers.push(data); //添加新值
			}
			
			if (data.Containerstats == 0) {
				
			} else if (data.Containerstats == 2) {//containerid
				console.log(data.Containerid + "停止运行");
				$("#nodelist td." + data.Containerid).find('i.fa-pause').removeClass('fa-pause').addClass('fa-play');
				$("#nodelist td." + data.Containerid).find('a.pause').removeClass('pause').addClass('play');
            } else if (data.Containerstats == 3) {
				console.log(data.Containerid + "容器不存在");
            } else if (data.Containerstats == 1) {
				console.log(data.Containerid + "容器正在重启");
			}
            break;
        case 2: // EVENT_ERROR
			if (ms_modal_contr == false){
				$(".ms_modal").show(1000);
				ms_modal_contr = true;
			}
            $(".ms_modal .box-body").html($(".ms_modal .box-body").html() + '<span style=\'color:red;\'>' + data.Error + '</span><br/>')
            break;
        case 3: // EVENT_MESSAGE
			if (ms_modal_contr == false){
				$(".ms_modal").show(1000);
				ms_modal_contr = true;
			}
            $(".ms_modal .box-body").html($(".ms_modal .box-body").html() + data.Message + '<br/>')
            break;
		case 4: //node_update_container_id
			nodelist = $("#nodelist tbody tr").find("i.node_id");
			nodelist.each(function(){
				if($(this).text() == data.Nodeid){
					$(this).parent().removeAttr("class");
					$(this).parent().addClass(data.Containerid);
				}
			});
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
	
	//每0.5秒执行一次
	function flash() {
		$.each(deadcontainers,function(k, v){
			if (v.Containerstats == 0) {
				$("#nodelist td." + v.Containerid).find('i.fa-pause').removeClass('fa-pause').addClass('fa-play');
				$("#nodelist td." + v.Containerid).find('a.pause').removeClass('pause').addClass('play');
				$("#nodelist td." + v.Containerid).find('i.fa-play').removeClass('fa-play').addClass('fa-pause');
				$("#nodelist td." + v.Containerid).find('a.play').removeClass('play').addClass('pause');
				$("#nodelist td." + v.Containerid).parent().removeClass("bg-yellow");
				if ($("td." + v.Containerid).parent().hasClass("bg-red")) {
					$("td." + v.Containerid).parent().removeClass("bg-red")
				}
				var td_a_list = $("." + v.Containerid).children();
				td_a_list.each(function(k,a){
					if ($(a).hasClass("rocket") || $(a).hasClass("trash")) {
						return true; //continue
					} else {
						$(a).removeClass("disabled")
					}
				});
				return true;
			} else if (v.Containerstats == 2) {
				if ($("td." + v.Containerid).parent().hasClass("bg-yellow")) {
					$("td." + v.Containerid).parent().removeClass("bg-yellow")
				} else {
					$("td." + v.Containerid).parent().addClass("bg-yellow")
				}
			} else if (v.Containerstats == 3) {
				//var prev_td = $("td." + v.Containerid).prev().prev();
				//prev_td.text("容器已被删除")
				var td_a_list = $("." + v.Containerid).children();
				td_a_list.each(function(k,a){
					if ($(a).hasClass("rocket") || $(a).hasClass("trash")) {
						return true;
					} else {
						if ($(a).hasClass("disabled")){
							
						} else {
							$(a).addClass("disabled")
						}
					}
				});
				//改变闪烁颜色
				if ($("td." + v.Containerid).parent().hasClass("bg-red")) {
					$("td." + v.Containerid).parent().removeClass("bg-red")
				} else {
					$("td." + v.Containerid).parent().addClass("bg-red")
				}
				//不存在
			} else if (v.Containerstats == 1) {
				//正在重启
			}
			
		});
	}
	
	
    _interval = setInterval(flash, 500);
});
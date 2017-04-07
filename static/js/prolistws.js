$(document).ready(function () {
    // Create a socket
//    socket = new WebSocket('ws://' + window.location.host + '/ws/join?uname=' + $('#uname').text());
    socket = new WebSocket('ws://' + window.location.host + '/ws/project');
    socket.onmessage = function (event) {
		var data = $.parseJSON(event.data);
		console.log(data);
		$.each(data,function(index, value){
			console.log(index + "test:")
			console.log(value.Type)
			switch (value.Type) {
				case 1:
					prostats(value);
					break
				case 0:
					break
			}
		});
    };
	socket.onclose = function (event) {
		console.log("socket closed");
		
	}
	socket.onopen = function (event) {
		socket.send(envsign);
		$("#envul").show();
		nonsocket = function(env) {
			if (env.hasClass("DE")) {
				envsign = "DE"
			} else if (env.hasClass("QE")){
				envsign = "QE"
			} else if (env.hasClass("OE")) {
				envsign = "OE"
			}
			if (envsign== ""){
				console.log("环境标识<空>")
				return;
			}
			socket.send(envsign);
		}
	}
})

function prostats(pro) {
	if (pro.Prorunnodes == 0 && pro.Proallnodes > 0) {
		switchstatus(pro.Envsign, pro.Proid, "danger");
		
	} else if (pro.Prorunnodes == pro.Proallnodes) {
		switchstatus(pro.Envsign, pro.Proid, "success");
	} else {
		switchstatus(pro.Envsign, pro.Proid, "warning");
	}
	available(pro.Envsign, pro.Proid, pro.Proallnodes,pro.Prorunnodes);
}

function available(envsign, trproid, all, run) {
	trpro = $("table#" + envsign).find("tr.proid" +trproid);
	trpro.find("td.procount a").text(run + "/" + all);
}

function switchstatus(envsign, trproid, stat) {
	if (stat == "" || envsign == "" || trproid == "") {
		return;
	}
	trpro = $("table#" + envsign).find("tr.proid" +trproid);
	if (stat == "warning") {
		trpro.find("td.prostat").html("<span class=\"label label-warning\">运行警告</span>");
	} else if(stat == "success") {
		trpro.find("td.prostat").html("<span class=\"label label-success\">正常运行</span>");
	} else if(stat == "danger") {
		trpro.find("td.prostat").html("<span class=\"label label-danger\">项目异常</span>");
	}
	
}
<!DOCTYPE html>
<html>
<head>
    {{template "common/meta.tpl" .}}
    <!-- Bootstrap 3.3.6 -->
    <link rel="stylesheet" href="/static/bootstrap/css/bootstrap.min.css">
    <!-- Font Awesome -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.5.0/css/font-awesome.min.css">
    <!-- Ionicons -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/ionicons/2.0.1/css/ionicons.min.css">
    <!-- DataTables -->
    <link rel="stylesheet" href="/static/plugins/datatables/dataTables.bootstrap.css">
    <!-- Theme style -->
    <link rel="stylesheet" href="/static/css/AdminLTE.min.css">
    <!-- AdminLTE Skins. Choose a skin from the css/skins
         folder instead of downloading all of them to reduce the load. -->
    <link rel="stylesheet" href="/static/css/skins/_all-skins.min.css">

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
    <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->
	<style>
	.ms_modal {
	    position: fixed;
	    top: 0;
	    right: 0;
	    bottom: 0;
	    left: 0;
	    z-index: 1050;
	    display: none;
	    overflow: hidden;
	    -webkit-overflow-scrolling: touch;
	    outline: 0;
		background: rgba(0,0,0,0.3);
	}
	    .imageselection {
            cursor:pointer;
         }
		#fade_close {
			float:right;
		}
		.full-sreen {
			width: 100%;
		}
		.center {
			position: absolute;
		    top: 25%;
		    left: 25%;
		    width: 50%;
		    height: 50%;
		}
    </style>
</head>
<body class="hold-transition skin-blue sidebar-mini">
<span style="display:none" class="pro_id">{{.pro.Id}}</span>
<span style="display:none" class="env_sign">{{.env.Sign}}</span>
<div class="wrapper">
    {{template "common/headertitle.tpl" .}}
    <!-- Left side column. contains the logo and sidebar -->
    {{template "common/sidebar.tpl" .}}
    <!-- Content Wrapper. Contains page content -->
    <div class="content-wrapper">
        <!-- Content Header (Page header) -->
        <section class="content-header">
            <h1>
                节点列表
                <small><a href="/{{.pro.Id}}/{{.env.Sign}}/nodeadd">
                    <button type="button" class="btn btn-block btn-primary btn-xs">添加节点</button>
                </a></small>
            </h1>
            <ol class="breadcrumb">
                <li><a href="/"><i class="fa fa-dashboard"></i>主页</a></li>
                <li><a href="/prolist">项目列表({{.env.Name}})</a></li>
                <li class="active">{{.proname}}项目</li>
            </ol>
        </section>

        <!-- Main content -->
        <section class="content">
            <div class="row">
                <div class="col-xs-12">
                    <div class="box">
                        <div class="box-header">
                            <h3 class="box-title">项目<<b>{{.pro.Name}}</b>></h3>
							<h4>环境<<b>{{.env.Name}}</b>></h4>
                        </div>
                        <!-- /.box-header -->
                        <div class="box-body">
                            <table id="nodelist" class="table table-bordered table-striped">
                                <thead>
                                <tr>
                                    <th>节点名称</th>
                                    <th>隶属主机</th>
                                    <th>版本号</th>
                                    <th>可用版本</th>
                                    <th>操作</th>
                                </tr>
                                </thead>
                                <tbody>
								{{ range Getnode .pro.Id .env.Sign}}
                                <tr>
                                    <td><input type="checkbox"><a href="/{{$.pro.Id}}/{{$.env.Sign}}/nodedit/{{.Id}}">{{.Sign}}</a></td>
                                    <td><a href="/macedit/{{.Mac.Id}}">{{.Mac.Ipaddr2}}</a></td>
                                    <td>{{if .CurVer}}
											{{.CurVer}}
										{{else}}
											需要初始化
										{{end}}
                                    </td>
                                    <td class="ver">正在获取...</td>
                                    <td class="{{.DocId}}">
										<a class="btn rocket">
                                        	<i class="fa fa-rocket"></i>
                                    	</a>
                                        <a class="btn pause">
                                            <i class="fa fa-pause"></i>
                                        </a>
                                        <a class="btn repeat">
                                            <i class="fa fa-repeat"></i>
                                        </a>
                                        <a class="btn trash">
                                            <i class="fa fa-trash"></i>
                                        </a>
										<i class="node_id" style="display:none">{{.Id}}</i>
										
									</td>
										
                                </tr>
								{{end}} <!--end range Getnode   -->
                               
                                </tbody>
                                <tfoot>
                                <tr>
                                    <th>节点名称</th>
                                    <th>隶属主机</th>
                                    <th>版本号</th>
                                    <th>可用版本</th>
                                    <th>操作</th>
                                </tr>
                                </tfoot>
                            </table>
                            <div class="box-body no-padding">
                                <div class="mailbox-controls">
                                    <!-- Check all button -->
                                    <button type="button" class="btn btn-default btn-sm checkbox-toggle"><i
                                            class="fa fa-square-o"></i>
                                    </button>
                                    <div class="btn-group">
                                        <button type="button" class="btn btn-default btn-sm"><i
                                                class="fa fa-repeat"></i></button>
                                        <button type="button" class="btn btn-default btn-sm"><i class="fa fa-pause"></i>
                                        </button>
                                        <button type="button" class="btn btn-default btn-sm"><i
                                                class="fa fa-trash"></i></button>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <!-- /.box-body -->
                    </div>
                    <!-- /.box -->
                </div>
                <!-- /.col -->
            </div>
            <!-- /.row -->
        </section>
        <!-- /.content -->
    </div>
    <!-- /.content-wrapper -->
				<!-- 确认框 -->
	<div class="verify-modal">
        <div class="modal">
          <div class="modal-dialog">
            <div class="modal-content">
              <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                  <span id="ver_break" aria-hidden="true">×</span></button>
                <h4 class="modal-title">Default Modal</h4>
              </div>
              <div class="modal-body">
                <p>One fine body…</p>
              </div>
              <div class="modal-footer">
                <button id="ver_close" type="button" class=" btn btn-default pull-left" data-dismiss="modal">Close</button>
                <button id="ver_save" type="button" class="btn btn-primary">Save changes</button>
              </div>
            </div>
            <!-- /.modal-content -->
          </div>
          <!-- /.modal-dialog -->
        </div>
        <!-- /.modal -->
      </div>
	<!-- 确认框 end -->
    {{template "common/footitle.tpl" .}}    
</div>

<!-- 实时打印信息--->
	<div id="fade" class="ms_modal"> 
		<div class="col-md-3 center">
          <div class="box box-primary ">
            <div class="box-header with-border">
              <h3 class="box-title">Collapsable</h3>

              <div id="ms_title" class="box-tools pull-right">
                <span class="btn btn-box-tool"><i class="fa fa-times"></i></span>
              </div>
              <!-- /.box-tools -->
            </div>
            <!-- /.box-header -->
            <div class="box-body" style="display: block;">
              
            </div>
            <!-- /.box-body -->
          </div>
          <!-- /.box -->
        </div>
	</div>
<!--实时打印信息   end-->
<!-- ./wrapper -->

<!-- jQuery 2.2.3 -->
<script src="/static/plugins/jQuery/jquery-2.2.3.min.js"></script>
<!-- Bootstrap 3.3.6 -->
<script src="/static/bootstrap/js/bootstrap.min.js"></script>
<!-- DataTables -->
<script src="/static/plugins/datatables/jquery.dataTables.min.js"></script>
<script src="/static/plugins/datatables/dataTables.bootstrap.min.js"></script>
<!-- SlimScroll -->
<script src="/static/plugins/slimScroll/jquery.slimscroll.min.js"></script>
<!-- FastClick -->
<script src="/static/plugins/fastclick/fastclick.js"></script>
<!-- AdminLTE App -->
<script src="/static/js/app.min.js"></script>
<!-- AdminLTE for demo purposes -->
<script src="/static/js/demo.js"></script>
<!-- page script -->


<script>
	var pro_id = $("span.pro_id").text();
	var env_sign = $("span.env_sign").text();
    
    var table = $('#nodelist').DataTable({
            "paging": true,
            "lengthChange": false,
            "searching": false,
            "ordering": false,
            "info": true,
            "autoWidth": false
        });
	//加载所有可用版本
	asdf();
	function asdf() {
		$.ajax({
		url:"/jproverlist",
		type: "post",
		data:{
			id: pro_id,
		},
		dataType: "json",
		success: function(msg) {
			if (msg.message == "success"){
				console.log(msg.data);
				loadver(msg.data); 
				return;
			} else {
				alert(msg.error);
			}
			
		},
		error:function(XMLHttpRequest, textStatus, errorThrown) {
			if (XMLHttpRequest.status != 0) {
						alert(XMLHttpRequest.status);
					}
			if (XMLHttpRequest.status == "503" || XMLHttpRequest.status == "500" ) {
				location.href="/503"
			}
		}
	});
	}
	
	function loadver(tags) {
		$("#nodelist").children().find(".ver").html("");
		var verlist = "";
		$.each(tags,function(n,value) { 
			verlist = verlist + '<select name=\'' + n + '\' class=\'ver\'>' 
			$.each(value, function(k, tag){
				if (k == 0){
					verlist = verlist + '<option value=' + tag + '>' + n + '--' + tag + '</option>';
					return true;
				}
				verlist = verlist + '<option value=' + tag + '>' + tag + '</option>';
			});
			verlist = verlist + '</select>';
		});
		$("#nodelist").children().find(".ver").append(verlist);
		
	}
	var abcc = function(){}
	$("a.trash").click(function(){
		var node_id = $(this).siblings("i.node_id").text();
		var node_name = $(this).parent().siblings("td:nth-child(1)").text();
		var node_mac = $(this).parent().siblings("td:nth-child(2)").text();
		$(".modal").show(1000);
		$(".modal .modal-title").text("删除节点")
		$(".modal .modal-body").html("确定删除节点:<b>" + node_name + "</b>, 在主机:<b>" + node_mac+ "</b>");
		var tr = $(this).parent().parent();
		var tr_index = table.row(tr).index();
		console.log("node_id:" + node_id)
		abcc = function() {
			$.ajax({
				url:"/nodedel",
				type: "post",
				data:{
					node_id: node_id,
				},
				dataType: "json",
				success: function(msg) {
					if (msg.message == "success"){
						table.row(tr_index).remove().draw( false );
						$(".model").hide(1000);
						return;
					} else {
						alert(msg.error);
					}
					
				},
				error:function(XMLHttpRequest, textStatus, errorThrown) {
					if (XMLHttpRequest.status != 0) {
						alert(XMLHttpRequest.status);
					}
					if (XMLHttpRequest.status == "503" || XMLHttpRequest.status == "500" ) {
						location.href="/503"
					}
				}
			});//endajax
		}
	});
		//确认框关闭
	$("#ver_close").click(function(){
		$(".modal").hide(1000);
	})
	$("#ver_break").click(function(){
		$(".modal").hide(1000);
	})
			
	$("#ver_save").click(function(){
		abcc();
		$("#ver_close").click();
	})
	$("table#nodelist").on('click', 'td a.rocket', function(){
		console.log("Test")
		var ver_item = $(this).parent().siblings(".ver").find(".ver");
		var node_id = $(this).siblings(".node_id").text();
       	if (typeof(ver_item) == "undefined" || typeof(node_id)== "undefined") {
			return;
		}
		var verlist = "";
		var ver_contr = true;
		verlist = node_id;
		ver_item.each(function(){
			if ($(this).val() == "" || node_id == "") {
				ver_contr = false;
				return false;
			}
			//push() 方法可向数组的末尾添加一个或多个元素，并返回新的长度。
			verlist = verlist + "-" + $(this).attr("name") + "_" + $(this).val();
			console.log($(this).attr("name") + "_" + $(this).val());
		});
		if ( !ver_contr || ver_item.length == 0){
			return;
		}
		//清空实时信息打印框
		$(".ms_modal .box-body").html(" ");
		socket.send(verlist);
		console.log("Test2")
	});
	
	//实时信息打印完毕关闭按钮的单击事件
	$("div.center").on("click", "span.btn-box-tool", function(){
		$("div.ms_modal").hide(1000);
		ms_modal_contr = false;
		return;
	});
	

	
	$("table#nodelist").on('click', 'td a', function(){
		var node_id = $(this).siblings(".node_id").text();
		var signal;
		if ($(this).hasClass("repeat")) {
			signal = "RESTART"
		} else if ($(this).hasClass("pause")) {
			signal = "STOP"
		} else if ($(this).hasClass("play")) {
			signal = "START"
		} else {
			return;
		}
		$.ajax({
			url:"/jnodeopera",
			type: "post",
			data:{
				id: node_id,
				env_sign: env_sign,
				signal: signal,
			},
			dataType: "json",
			success: function(msg) {
				if (msg.message == "success"){
					console.log(msg.data);
					alert(msg.data);
					return;
				} else {
					alert(msg.data);
				}
				
			},
			error:function(XMLHttpRequest, textStatus, errorThrown) {
				if (XMLHttpRequest.status != 0) {
						alert(XMLHttpRequest.status);
					}
				if (XMLHttpRequest.status == "503" || XMLHttpRequest.status == "500" ) {
					location.href="/503"
				}
			}
		});
	});

</script>
<!-- deploy websocket 务必放到最后面--> 
<script src="/static/js/deployws.js"></script>
</body>
</html>
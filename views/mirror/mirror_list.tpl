<!DOCTYPE html>
<html>
<head>
    {{template "common/meta.tpl" .}}
    <!-- Bootstrap 3.3.6 -->
    <link rel="stylesheet" href="static/bootstrap/css/bootstrap.min.css">
    <!-- Font Awesome -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.5.0/css/font-awesome.min.css">
    <!-- Ionicons -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/ionicons/2.0.1/css/ionicons.min.css">
    <!-- DataTables -->
    <link rel="stylesheet" href="static/plugins/datatables/dataTables.bootstrap.css">
    <!-- Theme style -->
    <link rel="stylesheet" href="static/css/AdminLTE.min.css">
    <!-- AdminLTE Skins. Choose a skin from the css/skins
         folder instead of downloading all of them to reduce the load. -->
    <link rel="stylesheet" href="static/css/skins/_all-skins.min.css">

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
    <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->
</head>
<body class="hold-transition skin-blue sidebar-mini">
<span id="isadd" style="display:none">{{ .isadd }}</span>
<span id="isaddgid" style="display:none">{{.mirgid}}</span>
<div class="wrapper">
    {{template "common/headertitle.tpl" .}}
    <!-- Left side column. contains the logo and sidebar -->
    {{template "common/sidebar.tpl" .}}
    <!-- Content Wrapper. Contains page content -->
    <div class="content-wrapper">
        <!-- Content Header (Page header) -->
        <section class="content-header">
            <h1>
                镜像列表
                <small>
                    <a href="/mirradd">
                        <button type="button" class="btn btn-block btn-primary btn-xs">添加镜像</button>
                    </a>
                </small>
            </h1>
            <ol class="breadcrumb">
                <li><a href="#"><i class="fa fa-dashboard"></i>主页</a></li>
                <li class="active">镜像列表</li>
            </ol>
        </section>

        <!-- Main content -->
        <section class="content">
            <div class="row">
                <div class="col-md-12">
                    <div class="nav-tabs-custom">
                        <ul class="nav nav-tabs">
							{{ range .mirgs}}
								{{ if eq $.mirgid .Id }}
										<!-- 如果接收到上个页面传过来的类别ID就把该标签设为活动状态 class="active" -->
										<li class="mirg active" ><a href="#appmirr" data-toggle="tab">{{ .Name }}</a><span class="mirg_id" style="display:none">{{ .Id }}</span></li>
								{{else}}
										<li  class="mirg" ><a href="#appmirr" data-toggle="tab">{{ .Name }}</a><span class="mirg_id" style="display:none">{{ .Id }}</span></li>
				
								{{end}}
							{{end}}
                            
                            <li><a href="/mirrgroupadd"><i class="fa fa-plus text-aqua">新增分类</i></a></li>
                        </ul>
                        <div class="tab-content">
                            <div class="active tab-pane" id="basemirr">
                                <div class="post">
                                    <div class="box">
                                        <div class="box-header">
                                            <h3 class="box-title">基础镜像</h3>
                                        </div>
                                        <!-- /.box-header -->
                                        <div class="box-body">
                                            <div class="col-lg-3 col-xs-6">
                                                <!-- small box -->
                                                <div class="small-box bg-green">
                                                    <div class="inner">
                                                        <h3>53<sup style="font-size: 20px">%</sup></h3>

                                                        <p>Bounce Rate</p>
                                                    </div>
                                                    <div class="icon">
                                                        <i class="ion ion-stats-bars"></i>
                                                    </div>
                                                    <a href="#" class="small-box-footer">
                                                        More info <i class="fa fa-arrow-circle-right"></i>
                                                    </a>
                                                </div>
                                            </div>
                                            <div class="col-lg-3 col-xs-6">
                                                <!-- small box -->
                                                <div class="small-box bg-green">
                                                    <div class="inner">
                                                        <h3>53<sup style="font-size: 20px">%</sup></h3>

                                                        <p>Bounce Rate</p>
                                                    </div>
                                                    <div class="icon">
                                                        <i class="ion ion-stats-bars"></i>
                                                    </div>
                                                    <a href="#" class="small-box-footer">
                                                        More info <i class="fa fa-arrow-circle-right"></i>
                                                    </a>
                                                </div>
                                            </div>
                                            <div class="col-lg-3 col-xs-6">
                                                <!-- small box -->
                                                <div class="small-box bg-green">
                                                    <div class="inner">
                                                        <h3>53<sup style="font-size: 20px">%</sup></h3>

                                                        <p>Bounce Rate</p>
                                                    </div>
                                                    <div class="icon">
                                                        <i class="ion ion-stats-bars"></i>
                                                    </div>
                                                    <a href="#" class="small-box-footer">
                                                        More info <i class="fa fa-arrow-circle-right"></i>
                                                    </a>
                                                </div>
                                            </div>
                                            <div class="col-lg-3 col-xs-6">
                                                <!-- small box -->
                                                <div class="small-box bg-green">
                                                    <div class="inner">
                                                        <h3>53<sup style="font-size: 20px">%</sup></h3>

                                                        <p>Bounce Rate</p>
                                                    </div>
                                                    <div class="icon">
                                                        <i class="ion ion-stats-bars"></i>
                                                    </div>
                                                    <a href="#" class="small-box-footer">
                                                        More info <i class="fa fa-arrow-circle-right"></i>
                                                    </a>
                                                </div>
                                            </div>
                                            <div class="col-lg-3 col-xs-6">
                                                <!-- small box -->
                                                <div class="small-box bg-green">
                                                    <div class="inner">
                                                        <h3>53<sup style="font-size: 20px">%</sup></h3>

                                                        <p>Bounce Rate</p>
                                                    </div>
                                                    <div class="icon">
                                                        <i class="ion ion-stats-bars"></i>
                                                    </div>
                                                    <a href="#" class="small-box-footer">
                                                        More info <i class="fa fa-arrow-circle-right"></i>
                                                    </a>
                                                </div>
                                            </div>
                                        </div>
                                        <!-- /.box-body -->
                                    </div>
                                </div>
                            </div>
							
							
							{{ range .mirgs}}
								{{ if eq $.mirgid .Id }}
										<!-- 如果接收到上个页面传过来的类别ID就把该标签设为活动状态 class="active" -->
										<!-- /.tab-pane 每一个镜像类别一个div start -->
		                            <div class="tab-pane" id="appmirr{{ .Id }}">
		                                <!-- The timeline -->
		                                <div class="post">
		                                    <div class="box">                                        
		                                        <div class="box-body">
													<table id="example{{.Id}}" class="table table-bordered table-hover">
		                                                <thead>
		                                                <tr>
		                                                    <th>镜像名称</th>
		                                                    <!-- <th>基础镜像</th> -->
		                                                    <th>仓库地址</th>
															<th>操作</th>
		                                                </tr>
		                                                </thead>
		                                                <tbody>
														{{ if $.mirs }}
															{{ range $.mirs }}
																<tr>
				                                                    <td>{{ .Name }}</td>
				                                                    <td>{{ .Hubaddress }} </td>
				                                                    <td>
																		<span class="mir_edit">
				                                                        <a class="btn">
				                                                            <i class="fa fa-edit">编辑</i>
				                                                        </a>
																		</span>
																		<span id="mir_id" style="display:none">{{.Id}}</span>
																		<span class="mir_remove">
				                                                        <a class="btn">
				                                                            <i class="fa fa-trash">删除</i>
				                                                        </a></span>
				                                                    </td>
		                                                		</tr>
															{{ end }}
														
														{{else}}
															<tr>
			                                                    <td>该类别下没有镜像</td>
			                                                    <td></td>
			                                                    <td>
			                                                        <a class="btn">
			                                                            <i class="fa fa-edit"></i>
			                                                        </a>
			                                                        <a class="btn">
			                                                            <i class="fa fa-trash"></i>
			                                                        </a>
			                                                    </td>
		                                                	</tr>
														{{end}}
		                                                
		                                                </tbody>
		                                                <tfoot>
		                                                <tr>
		                                                    <th>镜像名称</th>
		                                                    <th>基础镜像</th>
		                                                    <th>仓库地址</th>
		                                                </tr>
		                                                </tfoot>
		                                            </table>
		                                        </div>
		                                        <!-- /.box-body -->
		                                    </div>
		                                </div>
		                                <!-- /.post -->
		                            </div>    <!-- 每个镜像类别一个div -->
								{{else}}
										<!-- /.tab-pane 每一个镜像类别一个div start -->
		                            <div class="tab-pane" id="appmirr{{.Id}}">
		                                <!-- The timeline -->
		                                <div class="post">
		                                    <div class="box">                                        
		                                        <div class="box-body">
													<table id="example{{.Id}}" class="table table-bordered table-hover">
		                                                <thead>
		                                                <tr>
		                                                    <th>镜像名称</th>
		                                                    <!-- <th>基础镜像</th> -->
		                                                    <th>仓库地址</th>
															<th>操作</th>
		                                                </tr>
		                                                </thead>
		                                                <tbody>
														
		                                                
		                                                </tbody>
		                                                <tfoot>
		                                                <tr>
		                                                    <th>镜像名称</th>
		                                                    <th>基础镜像</th>
		                                                    <th>仓库地址</th>
		                                                </tr>
		                                                </tfoot>
		                                            </table>
		                                        </div>
		                                        <!-- /.box-body -->
		                                    </div>
		                                </div>
		                                <!-- /.post -->
		                            </div>    <!-- 每个镜像类别一个div -->
				
								{{end}}
							{{end}}
                            
							
							                       
                        </div>
                        <!-- /.tab-content -->
                    </div>
                    <!-- /.nav-tabs-custom -->
                </div>
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
<!-- ./wrapper -->
<!-- jQuery 2.2.3 -->
<script src="static/plugins/jQuery/jquery-2.2.3.min.js"></script>
<!-- Bootstrap 3.3.6 -->
<script src="static/bootstrap/js/bootstrap.min.js"></script>
<!-- DataTables -->
<script src="static/plugins/datatables/jquery.dataTables.min.js"></script>
<script src="static/plugins/datatables/dataTables.bootstrap.min.js"></script>
<!-- SlimScroll -->
<script src="static/plugins/slimScroll/jquery.slimscroll.min.js"></script>
<!-- FastClick -->
<script src="static/plugins/fastclick/fastclick.js"></script>
<!-- AdminLTE App -->
<script src="static/js/app.js"></script>
<!-- AdminLTE for demo purposes -->
<script src="static/js/demo.js"></script>
<!-- page script -->
<script type="text/javascript">
 //   $(function () {
   //     $('#example1').DataTable();
     //   $('#example2').DataTable();
       // $('#example3').DataTable();
    //});
	
	//点击确认对话框执行的函数，可以动态绑定其它事件， abcc=function(){alert("cc")}
	var abcc = function(){
		alert("abc");
	}
	var mirs  = new Array();
	//强制刷新类别分组
	function refreshgroup(){
		gid = $("li.active.mirg").children("span").text();
		//$("#example" + gid).dataTable().fnDraw()
		//删除类别数组当前活动mirgid，然后模拟当前镜像类别点击刷新
		mirs.splice(gid,1)
		$("li.active").click();
	}
	
	// 点击镜像编辑 动态绑定(普通绑定事件不能用，因为要绑定的元素是动态生成 的。)
	var edit_status = "true";  //切换完成和编辑状态
	$(".box-body").on("click","span.mir_edit", function(){
		//获取镜像id
		mir_id = $(this).siblings("span#mir_id").text();
		mir_td1 = $(this).parent().siblings("td:nth-child(1)");
		mir_td2 = $(this).parent().siblings("td:nth-child(2)");
		
		if ( edit_status == "false" ) { //编辑完成，上传到服务器
			mir_td1text = mir_td1.find("input").attr("placeholder");
			mir_td2text = mir_td2.find("input").attr("placeholder");
			mir_td1input = mir_td1.find("input").val();
			mir_td2input = mir_td2.find("input").val();
			if (mir_td1input=="" && mir_td2input==""){$(this).children("a").children("i").text("编辑");mir_td1.html(mir_td1text);mir_td2.html(mir_td2text);edit_status="true";return;}
			if (mir_td1input==""){mir_td1input=mir_td1text}
			if (mir_td2input==""){mir_td2input=mir_td2text}
			$.get("jqmirr?mirid=" + mir_id + "&mirname="+ mir_td1input + "&mirhubaddress=" + mir_td2input, function(msg,status){
				if (status == "success"){
					if (msg.message == "success"){
						mir_td1.html(mir_td1input)
						mir_td2.html(mir_td2input)
						alertmessage(msg.type,msg.content+ '"' + mir_td1input + '"')
					}else {alertmessage(msg.type, msg.content); return;}
				}
			})
			edit_status="true"
			$(this).children("a").children("i").text("编辑");
			return;
		}
		
		
		input1 = "<input type=\"text\" class=\"form-control\" style=\"width:100%\;\" value=\"" + mir_td1.html() + "\" placeholder=\"" + mir_td1.html() + "\">"
		$(this).parent().siblings("td:nth-child(1)").html(input1);
		input2 = "<input type=\"text\" class=\"form-control\" style=\"width:100%\;\" value=\"" + mir_td2.html() + "\" placeholder=\"" + mir_td2.html() + "\">"
		$(this).parent().siblings("td:nth-child(2)").html(input2);
		//td1 = $(this).
		//设置“编辑” ==》 “完成”
		$(this).children("a").children("i").text("完成");
		edit_status = "false"
	});
	
	
	// 点击镜像删除 动态绑定(普通绑定事件不能用，因为要绑定的元素是动态生成 的。)
	$(".box-body").on("click","span.mir_remove", function(){
		//获取镜像id
		mir_id = $(this).siblings("span#mir_id").text();
		
		mir_td1 = $(this).parent().siblings("td:nth-child(1)");
		mir_td1text = mir_td1.find("input").attr("placeholder");
		if (typeof(mir_td1text) == "undefined"){
			$(".verify-modal").find(".modal-body").html("确认删除镜像:" + mir_td1.html());
		} else {
			$(".verify-modal").find(".modal-body").html("确认删除镜像:" + mir_td1text);
		}
		$(".modal").show();
		$(".verify-modal").find(".modal-title").html("你好")
		abcc = function (){
			$.get("/jqrmmir?mirid=" + mir_id, function(msg, status){
				if (status == "success"){
					if (msg.message == "success"){
						$("#ver_close").click();
						refreshgroup();
						alertmessage(msg.type,msg.content+ '"' + mir_td1.html() + '"');
						return;
					}else {
						alertmessage(msg.type,msg.content);
					}
				}else {
					alertmessage(3 ,'错误:'+status);
				}
			})
		}
	});

	
	$(document).ready(function(){
		
		
		//页面加载完成执行的事件：把显示基础镜像的百分比DIV隐藏并且把镜像列表DIV显示
		if ($("#isadd").text()=="true" ){
			gid = $("#isaddgid").text();
			$('#example' + gid).DataTable();
			mirs[gid] = ""
			$(".tab-content .tab-pane").removeClass("active");
			$("#appmirr"+ gid).addClass("active");
			//$("#basemirr").hide();
			//$("#appmirr").show();
		}

		//点击镜像类别事件
		var mirgid;
		$(".mirg").click(function(){
			//alert("asdfasdf")
			edit_status="true" //每点击一次就要把编辑状态还原
			
			//获取子元素SPAN隐藏的ID,去动态获取类别镜像 
			mirgid = Number($(this).children("span").text());
			$(".tab-content .tab-pane").removeClass("active");
			$("#appmirr"+ mirgid).addClass("active");
			if ( typeof(mirs[mirgid]) == "undefined"){
				$.get("jqmirrlist?mirgid=" + mirgid,function(data,status){
					$('#example' + mirgid + ' tbody').html("");
					$.each(data,function(i,event){
						$('#example' + mirgid + ' tbody').append("<tr><td> "+ event.Name +" </td><td> " + 
						event.Hubaddress + " </td><td><span  class=\"mir_edit\"><a class=\"btn\"><i class=\"fa fa-edit\">编辑</i></a></span><span id=\"mir_id\" style=\"display:none\">" +
						 event.Id + "</span><span  class=\"mir_remove\"><a class=\"btn\"><i class=\"fa fa-trash\">删除</i></a></span></td></tr>");
						
					})
					mirs[mirgid] = data;
					$('#example'+ mirgid).DataTable();
				});
				
			} else {
				/*$('#example1 tbody').html("");
				$.each(mirs[mirgid],function(i,event){
					$('#example'+ mirgid + ' tbody').append("<tr><td> "+ event.Name +" </td><td> " + 
					event.Hubaddress + " </td><td><span class=\"mir_edit\" ><a class=\"btn\"><i class=\"fa fa-edit\">编辑</i></a></span><span id=\"mir_id\" style=\"display:none\">" +
					 event.Id + "</span ><span  class=\"mir_remove\"><a class=\"btn\"><i class=\"fa fa-trash\">删除</i></a></span></td></tr>");
					
				});
				alert("aaa")*/
				
			}	
		});// end //点击镜像类别事件
		
		//确认框close事件
		$("#ver_close").click(function(){
			$(".modal").hide();
		})
		$("#ver_break").click(function(){
			$(".modal").hide();
		})
				
		$("#ver_save").click(function(){
			abcc();
		})
		
		
	})
	$("body").prepend('<nav class="navbar navbar-fixed-top" style="z-index: 1031; display:none;">' +
		'<div class="container"><div class="alert alert-warning"> '+
		'<a href="#" class="close" data-dismiss="alert">×</a>'+
		'<strong>警告！</strong>&nbsp;&nbsp;&nbsp;&nbsp;<span>您的网络连接有问题。<span></div></div></nav>');
	var t1;
	function alertmessage(type,content) {
		$(".navbar-fixed-top").show(1000);
		alertd = $("nav.navbar").children(".container").children(".alert");
		switch(type)
		{
			case 1://完成
				alertd.removeClass("alert-warning").removeClass("alert-error").addClass("alert-success");
				alertd.find("strong").text("成功!");
				break;
			case 2://警告
				alertd.removeClass("alert-success").removeClass("alert-error").addClass("alert-warning");
				alertd.find("strong").text("警告!");
				break;
			case 3://错误
				alertd.removeClass("alert-success").removeClass("alert-warning").addClass("alert-error");
				alertd.removeClass("alert-success").removeClass("alert-error").addClass("alert-warning");
				alertd.find("strong").text("错误!");
				break;
			default:
				break;
		}
		
		alertd.find("span").text(content);
		window.clearTimeout(t1);
		t1 = window.setTimeout("$('.alert').click();",5000);
	}
	
	$(".alert").click(function(){
		$(".navbar-fixed-top").hide(1000);
	});
</script>
</body>
</html>

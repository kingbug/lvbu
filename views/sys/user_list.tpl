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
    <link rel="stylesheet" href="/static/plugins/datatables/dataTables.bootstrap.css">
    <!-- Theme style -->
    <link rel="stylesheet" href="/static/css/AdminLTE.min.css">
    <!-- AdminLTE Skins. Choose a skin from the css/skins
         folder instead of downloading all of them to reduce the load. -->
    <link rel="stylesheet" href="/static/css/skins/_all-skins.min.css">
	<!-- Pace style -->
	<link rel="stylesheet" href="/static/plugins/pace/pace.min.css">
    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
    <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->
</head>
<body class="hold-transition skin-blue sidebar-mini">
<div class="wrapper">
    {{template "common/headertitle.tpl" .}}
    <!-- Left side column. contains the logo and sidebar -->
    {{template "common/sidebar.tpl" .}}
    <!-- Content Wrapper. Contains page content -->
    <div class="content-wrapper">
        <!-- Content Header (Page header) -->
        <section class="content-header">
            <h1>
                系统管理
            </h1>
            <ol class="breadcrumb">
                <li><a href="#"><i class="fa fa-dashboard"></i>主页</a></li>
                <li class="active">系统管理</li>
            </ol>
        </section>

        <!-- Main content -->
        <section class="content">
            <div class="row">
                <div class="col-md-12">
                    <div class="nav-tabs-custom">
                        <ul class="nav nav-tabs">
                            <li class="active"><a href="#userslist" data-toggle="tab">用户列表</a></li>
                            <li><a href="#poslist" data-toggle="tab">职位管理</a></li>
                            <li><a href="#settings" data-toggle="tab">系统参数</a></li>
                        </ul>
                        <div class="tab-content">
                            <div class="active tab-pane" id="userslist">
                                <div class="post">
                                    <div class="box">
                                        <div class="box-header">
                                            <h3 class="box-title">用户列表</h3>
                                            <label>
                                                <a href="/useradd">
                                                    <button type="button" class="btn btn-block btn-primary btn-xs">
                                                        添加用户
                                                    </button>
                                                </a></label>
                                        </div>
                                        <!-- /.box-header -->
                                        <div class="box-body">
                                            <table id="example1" class="table table-bordered table-hover">
                                                <thead>
                                                <tr>
                                                    <th>姓名</th>
                                                    <th>帐号</th>
                                                    <th>职位</th>
                                                    <th>性别</th>
                                                    <th>电话</th>
                                                    <th>创建时间</th>
                                                    <th>最后登录</th>
                                                    <th>操作</th>
                                                </tr>
                                                </thead>
                                                <tbody>
                                                <tr>
												 {{range $key, $val :=.users}}
                                                    <td class="usernick">{{$val.Nick}}
													<span style="display:none">{{.Id}}</span>
													<span class="editstatus">
													{{if eq $val.Status 0}}
													
														<a href="/lockuser/{{$val.Id}}"><i class="fa fa-lock text-blue"></i></a>
													{{else}}
														<a href="/unlockuser/{{$val.Id}}"><i class="fa fa-key text-red"></i></a>
													{{end}}
													</span>
                                                    </td>
                                                    <td>{{$val.UserName}}</td>
                                                    <td>{{Getposname $val.Position.Id}}</td>
                                                    <td>{{Getsex $val.Sex}}</td>
                                                    <td>{{$val.Phone}}</td>
                                                    <td>{{$val.Created}}</td>
                                                    <td>{{$val.Updated}}</td>
                                                    </td>
                                                    <td>
                                                        <a class="btn" href="/useredit/{{.Id}}">
                                                            <i class="fa fa-edit"></i>
                                                        </a>
														<span class="user_remove">
                                                        <a class="btn">
															
                                                            <i class="fa fa-trash"></i>
                                                        </a>
														</span>
														<span class="userid" style="display:none">{{.Id}}</span>
                                                    </td>
                                                </tr>
												{{end}}
                                                </tbody>
                                                <tfoot>
                                                <tr>
                                                    <th>姓名</th>
                                                    <th>帐号</th>
                                                    <th>职位</th>
                                                    <th>性别</th>
                                                    <th>电话</th>
                                                    <th>创建时间</th>
                                                    <th>最后登录</th>
                                                    <th>操作</th>
                                                </tr>
                                                </tfoot>
                                            </table>
                                        </div>
                                        <!-- /.box-body -->
                                    </div>
                                </div>
                            </div>
                            <!-- /.tab-pane -->
                            <div class="tab-pane" id="poslist">
                                <!-- The timeline -->
                                <div class="post">
                                    <div class="box">
                                        <div class="box-header">
                                            <h3 class="box-title">职位管理</h3>
                                            <span class="addpos">
                                                
												<label>
													<button type="button" class="btn btn-block btn-primary btn-xs">
                                                        添加职位
                                                    </button>
                                                </label>
												</span>
                                                    
                                        </div>
                                        <!-- /.box-header -->
                                        <div class="box-body">
                                            <table id="example2" class="table table-bordered table-hover">
                                                <thead>
                                                <tr>
                                                    <th>职位名称</th>
                                                    <th>职位标识</th>
                                                    <th>操作</th>
                                                </tr>
                                                </thead>
                                                <tbody>
												 {{range $key, $val :=.poss}}
                                                <tr>
                                                    <td>{{$val.Name}}</td>
                                                    <td>{{$val.Sign}}</td>
                                                    <td>
                                                        <a class="btn">
                                                            <i class="fa fa-users">成员</i>
                                                        </a>
                                                        <a class="btn" href="/permanage/{{$val.Id}}">
                                                            <i class="fa fa-shield">权限</i>
                                                        </a>
														<span class="pos_edit">
                                                        <a class="btn">
                                                            <i class="fa fa-edit">编辑</i>
                                                        </a>
														</span>
														<span class="pos_remove">
                                                        <a class="btn">
                                                            <i class="fa fa-trash">删除</i>
                                                        </a>
														</span>
														<span class="pos_id" style="display:none;">{{.Id}}</span>
                                                    </td>
                                                </tr>
                                               {{end}}
                                                </tbody>
                                                <tfoot>
                                                <tr>
                                                    <th>职位名称</th>
                                                    <th>职位标识</th>
                                                    <th>操作</th>
                                                </tr>
                                                </tfoot>
                                            </table>
                                        </div>
                                        <!-- /.box-body -->
                                    </div>
                                </div>
                            </div>
                            <!-- /.tab-pane -->

                            <div class="tab-pane" id="settings">
                                <div class="post">
                                    <div class="box">
                                        <div class="box-header">
                                            <h3 class="box-title">系统设置</h3>
                                        </div>
                                        <!-- /.box-header -->
                                        <div class="box-body">
                                            <form class="form-horizontal">
                                                <div class="form-group">
                                                    <label for="inputName" class="col-sm-2 control-label">邮箱主机</label>
                                                    <div class="col-sm-10">
                                                        <input type="text" class="form-control" id="inputName"
                                                               placeholder="邮箱主机地址">
                                                    </div>
                                                </div>
                                                <div class="form-group">
                                                    <label for="inputName" class="col-sm-2 control-label">邮箱帐号</label>
                                                    <div class="col-sm-10">
                                                        <input type="text" class="form-control" id="inputName"
                                                               placeholder="邮箱帐号">
                                                    </div>
                                                </div>
                                                <div class="form-group">
                                                    <label for="inputName" class="col-sm-2 control-label">邮箱密码</label>
                                                    <div class="col-sm-10">
                                                        <input type="text" class="form-control" id="inputName"
                                                               placeholder="邮箱密码">
                                                    </div>
                                                </div>
                                                
                                                <div class="form-group">
                                                    <div class="col-sm-offset-2 col-sm-10">
                                                        <button type="submit" class="btn btn-danger">Submit</button>
                                                    </div>
                                                </div>
                                            </form>
                                        </div>
                                        <!-- /.box-body -->
                                    </div>
                                </div>
                            </div>
                            <!-- /.tab-pane -->
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
		<!--start 添加职位弹出框  -->
			<div class="box box-info addpos" style="display:none">
		        <div class="box-header with-border">
		          <h3 class="box-title">添加职位</h3>
		
		          <div class="box-tools pull-right">
		            
		            <button type="button" class="btn btn-box-tool remove" data-widget="remove"><i class="fa fa-remove"></i></button>
		          </div>
		        </div>
		        <!-- /.box-header -->
		        <div class="box-body" style="display: block;">
					<form class="form-horizontal">
		              <div class="box-body">
		                <div class="form-group">
		                  <label for="inputEmail3" class="col-sm-2 control-label">职位名称</label>
		
		                  <div class="col-sm-10">
		                    <input type="text" class="form-control" id="inputEmail3" name="name" placeholder="职位名称">
		                  </div>
		                </div>
		                <div class="form-group">
		                  <label for="inputPassword3" class="col-sm-2 control-label">职位标识</label>
		
		                  <div class="col-sm-10">
		                    <input type="text" class="form-control" id="inputPassword3" name="sign" placeholder="example:DD">
		                  </div>
		                </div>
		                
		              </div>
		              <!-- /.box-body -->
		              <div class="box-footer">
						<span style="color:red"></span>
		                <button type="button" class="btn btn-info pull-right addpos">提交</button>
		              </div>
		              <!-- /.box-footer -->
            </form>
					
		          <div class="row">
		            
		            <!-- /.col -->
		            
		            <!-- /.col -->
		          </div>
		          <!-- /.row -->
		        </div>
		        <!-- /.box-body -->
		        
		      </div>
			<!--添加职位弹出框 end -->	
			<!---start 完成-->
				<div class="box box-info success" style="display:none">
		        <div class="box-header with-border">
		          <h3 class="box-title">返回结果</h3>
		
		          <div class="box-tools pull-right">
		            
		            <button type="button" class="btn btn-box-tool remove" data-widget="remove"><i class="fa fa-remove"></i></button>
		          </div>
		        </div>
		        <!-- /.box-header -->
		        <div class="box-body" style="display: block;">
					<form class="form-horizontal">
		              <div class="box-body">
		               	<span style="color:green;">OK</span>
		              </div>
		              <!-- /.box-body -->
		              <div class="box-footer">
		                <button type="button" class="btn btn-info pull-right success">完成</button>
		              </div>
		              <!-- /.box-footer -->
            </form>
					
		          <div class="row">
		            
		            <!-- /.col -->
		            
		            <!-- /.col -->
		          </div>
		          <!-- /.row -->
		        </div>
		        <!-- /.box-body -->
		        
		      </div>
			<!--完成--end-->
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
<script src="/static/plugins/jQuery/jquery-2.2.3.min.js"></script>
<!-- Bootstrap 3.3.6 -->
<script src="/static/bootstrap/js/bootstrap.min.js"></script>
<!-- DataTables -->
<script src="/static/plugins/datatables/jquery.dataTables.min.js"></script>
<script src="/static/plugins/datatables/dataTables.bootstrap.min.js"></script>
<!-- PACE -->
<script src="/static/plugins/pace/pace.min.js"></script>
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
	var table1 = $('#example1').DataTable();
	var table2 = $('#example2').DataTable();
	$(document).ajaxStart(function() { Pace.restart(); });

	

	$(".user_remove").click(function(){
		userid = $(this).siblings(".userid").text();
		usernick = $(this).parent().siblings(".usernick").text();
		tr = $(this).parent().parent().addClass("selected");
		$(".verify-modal").find(".modal-title").html("警告");
		$(".verify-modal").find(".modal-body").html("确认删除用户 : [" + usernick+ "]");
		$(".modal").show();
		abcc = function(){
			$.post("/jqrmuser",
			{
				userid:userid
			}, function(data, status){
				if (status == "success"){
					if (data == "success"){
						table1.row('.selected').remove().draw( false );
						$("#ver_close").click();
						return;
					}else {
						$(".verify-modal").find(".modal-title").html("警告")
						$(".verify-modal").find(".modal-body").html("服务器出错,返回数据data:" + data);
						abcc = function(){
							$("#ver_close").click();
						}
					}
				}else {
					alert("服务出错:" + data);
					$(".verify-modal").find(".modal-title").html("警告")
					$(".verify-modal").find(".modal-body").html("网络问题,statuscode:" + status);
					abcc = function(){
						$("#ver_close").click();
					}
				}
			})
		}
	})
	//点击确认对话框执行的函数，可以动态绑定其它事件， abcc=function(){alert("cc")}
	var abcc = function(){
		alert("abc");
	}
	//确认框close事件
	$("#ver_close").click(function(){
		$("tr.selected").removeClass("selected");
		$(".modal").hide();
	})
	$("#ver_break").click(function(){
		$("tr.selected").removeClass("selected");
		$(".modal").hide();
	})
	$(document).ready(function() {
		$("#ver_save").click(function(){
			abcc();
		})
	})	
	
	$("span.editstatus").click(function(){
		userid = $(this).siblings("span").text();
		var lockstatus;
		var icon = $(this).find("i.fa");
		if ($(this).find("i.fa").hasClass("fa-key")){
			lockstatus = "unlockuser"
		}else {lockstatus = "lockuser"}
		
		$.get("/"+ lockstatus + "/"+ userid,function(data, status){
			//$.get 里面不能用$(this)
			if (status == "success"){
				if (data == "lock"){
					icon.removeClass("fa-lock text-blue")
					icon.addClass("fa-key text-red")
				} else {
					if (data == "unlock"){
						icon.removeClass("fa-key text-red")
						icon.addClass("fa-lock text-blue")
					} else {
						alert("服务出错:" + data);
					}//endif
				}//endif
			}//endif
		})//endget
		return false;
	})//function
	
	var edit_status = "true";    //编辑按钮
	$(".box-body").on("click","span.pos_edit", function(){
		
		pos_id = $(this).siblings("span.pos_id").text();
		pos_td1 = $(this).parent().siblings("td:nth-child(1)");
		pos_td2 = $(this).parent().siblings("td:nth-child(2)");
		
		if ( edit_status == "false" ) { 
			pos_td1text = pos_td1.find("input").attr("placeholder");
			pos_td2text = pos_td2.find("input").attr("placeholder");
			pos_td1input = pos_td1.find("input").val();
			pos_td2input = pos_td2.find("input").val();
			if (pos_td1input=="" && pos_td2input==""){$(this).children("a").children("i").text("编辑");pos_td1.html(pos_td1text);pos_td2.html(pos_td2text);edit_status="true";return;}
			if (pos_td1input==""){pos_td1input=pos_td1text}
			if (pos_td2input==""){pos_td2input=pos_td2text}
			$.get("jqupdatepos?posid=" + pos_id + "&posname="+ pos_td1input + "&possign=" + pos_td2input, function(data,status){
				if (status == "success"){
					if (data == "success"){
						
						pos_td1.html(pos_td1input)
						pos_td2.html(pos_td2input)
						
					}else {alert("出错了，返回数据data:" + data); return;}
				}
			})
			edit_status="true"
			$(this).children("a").children("i").text("编辑");
			return;
		}
		
		
		input1 = "<input type=\"text\" class=\"form-control\" style=\"width:100%\;\" placeholder=\"" + pos_td1.html() + "\">"
		$(this).parent().siblings("td:nth-child(1)").html(input1);
		input2 = "<input type=\"text\" class=\"form-control\" style=\"width:100%\;\" placeholder=\"" + pos_td2.html() + "\">"
		$(this).parent().siblings("td:nth-child(2)").html(input2);
		
		
		$(this).children("a").children("i").text("完成");
		edit_status = "false"
	});
	
	// 点击镜像删除 动态绑定(普通绑定事件不能用，因为要绑定的元素是动态生成 的。)
	$(".box-body").on("click","span.pos_remove", function(){
		//获取镜像id
		pos_id = $(this).siblings("span.pos_id").text();
		
		pos_td1 = $(this).parent().siblings("td:nth-child(1)");
		pos_td1text = pos_td1.find("input").attr("placeholder");
		if (typeof(pos_td1text) == "undefined"){
			$(".verify-modal").find(".modal-body").html("确认删除职位:" + pos_td1.html());
		} else {
			$(".verify-modal").find(".modal-body").html("确认删除职位:" + pos_td1text);
		}
		$(".modal").show();
		$(".modal .modal-content").show();
		var tr = $(this).parent().parent();
		$(".verify-modal").find(".modal-title").html("职位删除")
		abcc = function (){
			$.get("/jqrmpos/" +pos_id, function(data, status){
				if (status == "success"){
					if (data == "success"){
						$("#ver_close").click();
						tr.addClass("selected");
						table2.row('.selected').remove().draw( false );
						return;
					}else {
						$(".verify-modal").find(".modal-title").html("警告")
						$(".verify-modal").find(".modal-body").html("服务器出错,返回数据data:" + data);
						abcc = function(){
						$("#ver_close").click();
						}
					}
				}else {
					
					$(".verify-modal").find(".modal-title").html("警告")
					$(".verify-modal").find(".modal-body").html("网络问题,statuscode:" + status);
					abcc = function(){
						$("#ver_close").click();
					}
				}
			})
		}
	});
	
	$("span.addpos").click(	function(){
		$("div.modal").show(1000);
		$("div.modal-content").hide();
		$("div.modal .addpos").show();
	})
	
	$(".box-tools .remove").click(function(){	
		$("div.modal").hide(1000);	
	})
	
	var name;
	var sign;
	var pos_id;
	$(".box-footer button.addpos").click(function(){
		var self = $(this)
		var box_body = $(this).parent().siblings(".box-body");
		name = box_body.find("input[name='name']").val();
		sign = box_body.find("input[name='sign']").val();
		var message = $(this).siblings("span");
		
		
		//判断不能为空
		if (name=="") { message.text("职位名不能为空");return;}
		if (sign=="") {message.text("标识不能为空");return;}
		message.text("");
		$.post("/jqaddpos",
		{
			name:name,
			sign:sign
		}, function(data, status){
			if (status=="success"){
				if (data.status=="success"){
					pos_id = data.id
					$(".modal .addpos").hide();
					$(".modal .success").show();
				}else {
					message.text("Error:" + data)
				}
				
			}else {
				message.text("请求出错!!!")
			}
		});
	})

	$(".box-footer").on("click","button.success", function(){
		$(".box-tools .remove").click();
		alert(name);
		table2.row.add([
			name+ '<span class="pull-right-container"><small class="label pull-right bg-green">new</small> </span>',
			sign,
			'<a class=\"btn\"><i class=\"fa fa-users\">成员</i></a><a class=\"btn\" href=\"\">'+
			' <i class=\"fa fa-shield\">权限</i>  '+
			'</a>'+
			'<span class=\"pos_edit\">'+
			'<a class=\"btn\">'+
			      '<i class=\"fa fa-edit\">编辑</i>'+
			   '</a>'+
			'</span>'+
			'<span class=\"pos_remove\">'+
			'<a class=\"btn\">'+
			'<i class=\"fa fa-trash\">删除</i>'+
			'  </a>'+
			'</span>'+
			'<span class=\"pos_id\" style=\"display:none;\">'+ pos_id +'</span>',
		]).draw();
 
	});
	
</script>
</body>
</html>

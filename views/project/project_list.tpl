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
</head>
<body class="hold-transition skin-blue sidebar-mini">
<nav class="navbar navbar-fixed-top" style="z-index: 1031; display:none;">

  <div class="container">

   <div class="alert alert-warning">
	<a href="#" class="close" data-dismiss="alert">
		×
	</a>
	<strong>警告！</strong>&nbsp;&nbsp;&nbsp;&nbsp;<span>您的网络连接有问题。<span>
</div>

  </div>

</nav>
<div class="wrapper">
    {{template "common/headertitle.tpl" .}}
    <!-- Left side column. contains the logo and sidebar -->
    {{template "common/sidebar.tpl" .}}
    <!-- Content Wrapper. Contains page content -->
    <div class="content-wrapper">
        <!-- Content Header (Page header) -->
        <section class="content-header">
            <h1>
                项目列表
				{{if(Isperitem "proa" .uid)}}
                <small>				
				<a href="/proadd">
                    <button type="button" class="btn btn-block btn-primary btn-xs">添加项目</button>
				</a>				
                </small>
				{{end}}
            </h1>
            <ol class="breadcrumb">
                <li><a href="#"><i class="fa fa-dashboard"></i>主页</a></li>
                <li class="active">项目列表</li>
            </ol>
        </section>

        <!-- Main content -->
        <section class="content">
            <div class="row">
                <div class="col-md-12">
                    <div class="nav-tabs-custom">
                        <ul class="nav nav-tabs" id="envul" style="display:none">
							{{if (Isuserper "DE" .uid)}}
                            	<li class="active"><a href="#activity" class="DE" data-toggle="tab">开发环境</a></li>
							{{end}}
                            
							{{if (Isuserper "QE" .uid)}}
                            	<li><a href="#timeline" class="QE" data-toggle="tab">测试环境</a></li>
							{{end}}
							{{if (Isuserper "OE" .uid)}}
                            	<li><a href="#settings" class="OE" data-toggle="tab">生产环境</a></li>	
							{{end}}					
                        </ul>
                        <div class="tab-content">
							{{if (Isuserper "DE" .uid)}}
                            <div class="active tab-pane" id="activity">
                                <div class="post">
                                    <div class="box">
                                        <div class="box-header">
                                            <h3 class="box-title">开发环境</h3>
                                        </div>
                                        <!-- /.box-header -->
                                        <div class="box-body">
                                            <table id="DE" class="table table-bordered table-hover">
                                                <thead>
                                                <tr>
                                                    <th>项目名称</th>
                                                    <th>项目状态</th>
                                                    <th>节点数量(s)</th>
                                                    <th>操作</th>
                                                </tr>
                                                </thead>
                                                <tbody>
												 {{range Getproject}}
                                                <tr class="proid{{.Id}}">
                                                    <td>{{.Name}}
														{{if $.newpro}}
															{{if eq .Id $.newpro}}
																	<span class="pull-right-container">
														              <small class="label pull-right bg-green">new</small>
														            </span>
															{{end}}
														{{end}}
														
													</td>
                                                    <td class="prostat">无可用节点
                                                    </td>
                                                    <td class="procount">
														<a href="/{{.Id}}/de/nodelist/" target="_blank" title="点击管理节点">
														
																无可用节点
														</a>
													</td>
                                                    <td>
                                                        {{if(Isperitem "proe" $.uid)}}
                                                        <a class="btn" href="/proedit/{{.Id}}"  target="_blank">
                                                            <i class="fa fa-edit">编辑</i>
                                                        </a>
														{{end}}
                                                        
														{{if(Isperitem "prod" $.uid)}}
                                                        <span class="rm_pro">
														<a class="btn">
                                                            <i class="fa fa-trash">删除</i>
                                                        </a>
														<i class="id" style="display:none">{{.Id}}</i>
														</span>
														{{end}}
														<!-- 权限配置查看验证 -->
														{{if (Isperitem "cons" $.uid)}}
                                                        
                                                            
															<button id="popover{{.Id}}" type="button" class="btn btn-link" style="padding-left:1px;" title="列表"  
																	data-container="body" data-toggle="popover" data-placement="right" 
																	data-content="
																	<a class='btn conffileadd' onclick='conffileadd({{.Id}})' style='color:#0033FF;'>添加文件</a>
																	{{$pid := .Id}}
																	
																		{{ range $index, $filename := .Conflist}}
																			<a class=' btn list-group-item' href='/{{$pid}}/conlist/?filename={{$filename}}'  >
																			{{ if $filename}}
																				{{$filename}}
																			{{else}}
																				默认文件
																			{{end}}
																			</a>
																		{{else}}
																			<a class='btn list-group-item' href='/{{$pid}}/conlist/'  >默认文件</a>
																		{{end}}
																			
																	">
																	<i class="fa fa-wrench btn-link"  title="点击管理项目配置文件">配置文件</i>
															</button>
															
                                                        
														{{end}}
														<i class="id" style="display:none">{{.Id}}</i>
                                                    </td>
                                                </tr>
                                               {{end}}
                                                </tfoot>
                                            </table>
                                        </div>
                                        <!-- /.box-body -->
                                    </div>
                                </div>
                            </div>
                            <!-- /.tab-pane -->
							{{end}}
							{{if (Isuserper "QE" .uid)}}
                            <div class="tab-pane" id="timeline">
                                <!-- The timeline -->
                                <div class="post">
                                    <div class="box">
                                        <div class="box-header">
                                            <h3 class="box-title">测试环境</h3>
                                        </div>
                                        <!-- /.box-header -->
                                        <div class="box-body">
                                            <table id="QE" class="table table-bordered table-hover">
                                                <thead>
                                                <tr>
                                                    <th>项目名称</th>
                                                    <th>项目状态</th>
                                                    <th>节点数量(s)</th>
                                                    <th>操作</th>
                                                </tr>
                                                </thead>
                                                <tbody>
												 {{range Getproject}}
                                                <tr class="proid{{.Id}}">
                                                    <td>{{.Name}}</td>
                                                    <td class="prostat">无可用节点
                                                    </td>
                                                    <td class="procount">
														<a href="/{{.Id}}/qe/nodelist/" target="_blank" title="点击管理节点">
														
																无可用节点
														</a>
													</td>
                                                    <td>
                                                        {{if(Isperitem "proe" $.uid)}}
                                                        <a class="btn" href="/proedit/{{.Id}}"  target="_blank">
                                                            <i class="fa fa-edit">编辑</i>
                                                        </a>
														{{end}}
                                                        
														{{if(Isperitem "prod" $.uid)}}
                                                        <span class="rm_pro">
														<a class="btn">
                                                            <i class="fa fa-trash">删除</i>
                                                        </a>
														<i class="id" style="display:none">{{.Id}}</i>
														</span>
														{{end}}
														{{if(Isperitem "cons" $.uid)}}
	                                                        <button id="popover{{.Id}}" type="button" class="btn btn-link" style="padding-left:1px;" title="列表"  
																	data-container="body" data-toggle="popover" data-placement="right" 
																	data-content="
																	<a class='btn conffileadd' onclick='conffileadd({{.Id}})' style='color:#0033FF;'>添加文件</a>
																	{{$pid := .Id}}
																	
																		{{ range $index, $filename := .Conflist}}
																			<a class='btn list-group-item' href='/{{$pid}}/conlist/?filename={{$filename}}' >
																			{{ if $filename}}
																				{{$filename}}
																			{{else}}
																				默认文件
																			{{end}}
																			</a>
																		{{else}}
																			<a class='btn list-group-item' href='/{{$pid}}/conlist/' >默认文件</a>
																		{{end}}
																			
																	">
																	<i class="fa fa-wrench btn-link"  title="点击管理项目配置文件">配置文件</i>
															</button>
														{{end}}
														
														<i class="id" style="display:none">{{.Id}}</i>
														
														
                                                    </td>
                                                </tr>
                                               {{end}}
                                                </tbody>
                                                <tfoot>
                                                <tr>
                                                    <th>项目名称</th>
                                                    <th>项目状态</th>
                                                    <th>节点数量(s)</th>
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
							{{end}}
							{{if (Isuserper "OE" .uid)}}
                            <div class="tab-pane" id="settings">
                                <div class="post">
                                    <div class="box">
                                        <div class="box-header">
                                            <h3 class="box-title">生产环境</h3>
                                        </div>
                                        <!-- /.box-header -->
                                        <div class="box-body">
                                            <table id="OE" class="table table-bordered table-hover"  target="_blank">
                                                <thead>
												
                                                <tr>
                                                    <th>项目名称</th>
                                                    <th>项目状态</th>
                                                    <th>节点数量(s)</th>
                                                    <th>操作</th>
                                                </tr>
                                                </thead>
                                                <tbody>
												{{range Getproject}}
                                                <tr class="proid{{.Id}}">
                                                    <td>{{.Name}}</td>
                                                    <td class="prostat">
													
																无可用节点
                                                    </td>
                                                    <td class="procount">
														<a href="/{{.Id}}/oe/nodelist/"  target="_blank" title="点击管理节点">
														
																无可用节点
														</a>
													</td>
                                                    <td>
                                                        {{if(Isperitem "proe" $.uid)}}
                                                        <a class="btn" href="/proedit/{{.Id}}"  target="_blank">
                                                            <i class="fa fa-edit">编辑</i>
                                                        </a>
														{{end}}
                                                        
														{{if(Isperitem "prod" $.uid)}}
                                                        <span class="rm_pro">
														<a class="btn">
                                                            <i class="fa fa-trash">删除</i>
                                                        </a>
														<i class="id" style="display:none">{{.Id}}</i>
														</span>
														{{end}}
														{{if(Isperitem "cons" $.uid)}}
	                                                        <button id="popover{{.Id}}" type="button" class="btn btn-link" style="padding-left:1px;" title="列表"  
																	data-container="body" data-toggle="popover" data-placement="right" 
																	data-content="
																	<a class='btn conffileadd' onclick='conffileadd({{.Id}})' style='color:#0033FF;'>添加文件</a>
																	{{$pid := .Id}}
																	
																		{{ range $index, $filename := .Conflist}}
																			<a class='btn list-group-item' href='/{{$pid}}/conlist/?filename={{$filename}}' >
																			{{ if $filename}}
																				{{$filename}}
																			{{else}}
																				默认文件
																			{{end}}
																			</a>
																		{{else}}
																			<a class='btn list-group-item' href='/{{$pid}}/conlist/' >默认文件</a>
																		{{end}}
																			
																	">
																	<i class="fa fa-wrench btn-link"  title="点击管理项目配置文件">配置文件</i>
															</button>
														{{end}}
														
														<i class="id" style="display:none">{{.Id}}</i>
														
                                                    </td>
                                                </tr>
                                                {{end}}
                                                </tbody>
                                                <tfoot>
                                                <tr>
                                                    <th>项目名称</th>
                                                    <th>项目状态</th>
                                                    <th>节点数量(s)</th>
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
							{{end}} <!-- //if OE -->
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

    var table1 = $('#DE').DataTable({
		"aLengthMenu" : [20, 40, 60],
		"iDisplayLength" : 40,
	});
    var table2 = $('#QE').DataTable({
		"aLengthMenu" : [20, 40, 60],
		"iDisplayLength" : 40,
	});
    var table3 = $('#OE').DataTable({
		"aLengthMenu" : [20, 40, 60],
		"iDisplayLength" : 40,
	});
	
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
	}
	var envsign;
	$("#envul li a").click(function(){
		nonsocket($(this));
		
	});
	$("#envul li:first a").click();	
	$(".rm_pro").click(function(){
		var pro_id = $(this).find("i.id").text();
		var mac_name = $(this).parent().siblings("td:nth-child(1)").text();
		var node_num = $(this).parent().siblings("td:nth-child(3)").text();
		$(".verify-modal").find(".modal-title").html("警告");
		$(".verify-modal").find(".modal-body").html("确认删除项目 : <b>" + mac_name+ "</b> ,节点数量 : <b>" + node_num + "</b>");
		$(".modal").show(1000);	
		var tr = $(this).parent().parent();
		var table;
		if (tr.parent().parent().attr("id") == "DE") {
			table = table1
		} else if (tr.parent().parent().attr("id") == "QE"){
			table = table2
		} else { table = table3}
		var tr_index = table.row(tr).index();
		//给确认按钮动态绑定事件
		abcc = function (){
			$.ajax({
				url:"/prodel",
				type: "post",
				data:{id: pro_id},
				dataType: "json",
				success: function(msg) {
					$("#ver_break").click();
					if (msg.message == "success"){
						table1.row(tr_index).remove().draw( false );
						table2.row(tr_index).remove().draw( false );
						table3.row(tr_index).remove().draw( false );
						alertmessage(msg.type,msg.content);
						return;
					} else {
						alertmessage(msg.type,msg.content);
					}
					
				},
				error:function(XMLHttpRequest, textStatus, errorThrown) {
					alertmessage(3,XMLHttpRequest.status);
					if (XMLHttpRequest.status == "503" || XMLHttpRequest.status == "500" ) {
						location.href="/503"
					}
				}
			});
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
	})
	
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
	
	
	var popoverhtmlstring;
	function conffileadd(pid) {
		$(".verify-modal").find(".modal-title").html("添加文件名");
		$(".verify-modal").find(".modal-body").html("<input type='text' name='filename' class='form-control'  placeholder='[0-9A-Za-z./]'/>");
		popoverhtmlstring = $("#popover"+pid).attr("data-content");
		//给确认按钮动态绑定事件
		abcc = function (){
			var filename;
			filename = $(".verify-modal").find(".modal-body").find("input").val();
			if (filename == ""){
				alertmessage(2, "文件名不能为空");
				return;
			}
			$.ajax({
				url:"/conffileadd",
				type: "post",
				data:{pid: pid, filename: filename},
				dataType: "json",
				success: function(msg) {
					$("#ver_break").click();
					if (msg.message == "success"){
						alertmessage(msg.type,msg.content);
						popoverhtmlstring = popoverhtmlstring + "<a class='btn list-group-item' href='/"+pid+"/conlist/?filename=" + filename + "' >" + filename + "</a>";
		
						$("#popover"+pid).attr("data-content", popoverhtmlstring);
						return;
					} else {
						alertmessage(msg.type,msg.content);
					}
					
				},
				error:function(XMLHttpRequest, textStatus, errorThrown) {
					alertmessage(3,XMLHttpRequest.status);
					if (XMLHttpRequest.status == "503" || XMLHttpRequest.status == "500" ) {
						location.href="/503";
					}
				}
			});// end ajax
		}
		$(".modal").show(1000);
	}
	$("[data-toggle='popover']").popover({html:true, trigger:"focus" });
</script>



<!--务必放在这里 -->
<script src="/static/js/prolistws.js"></script>
</body>
</html>

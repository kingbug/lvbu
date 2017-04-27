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
	<style type="text/css">
	.table>tbody>tr>td.key {
		max-width:150px;
	}
	.table>tbody>tr>td {
		padding:3px;
		vertical-align: initial; 
		font-size: 18px;
		max-width: 300px;
		overflow:hidden; /* 内容超出宽度时隐藏超出部分的内容 */ 
	}
	.input-h-w {
		height:38px;
		width:100%;
	}
	.open {
		color: #3c8dbc;
	}
	</style>
</head>
<body class="hold-transition skin-blue sidebar-mini">
<span style="display:none" id="pro_id">{{.pro.Id}}</span>

<div class="wrapper">

    {{template "common/headertitle.tpl" .}}
    <!-- Left side column. contains the logo and sidebar -->
    {{template "common/sidebar.tpl" .}}
    <!-- Content Wrapper. Contains page content -->
    <div class="content-wrapper">
        <!-- Content Header (Page header) -->
        <section class="content-header">
            <h1>
                {{.pro.Name}}配置列表
				
            </h1>
            <ol class="breadcrumb">
                <li><a href="#"><i class="fa fa-dashboard"></i>主页</a></li>
                <li class="active"><a href="/prolist">项目列表</a></li>
            </ol>
        </section>

        <!-- Main content -->
        <section class="content">
            <div class="row">
                <div class="col-md-12">
                    <div class="nav-tabs-custom" id="koko">
                        <ul class="nav nav-tabs">
							<li class="active"><a href="#All" data-toggle="tab">全部</a></li>
							{{ if Isuserper "DE" $.uid }}
                            <li><a href="#activity" data-toggle="tab">开发环境</a></li>
							{{end}}
							{{ if Isuserper "QE" $.uid }}
                            <li><a href="#timeline" data-toggle="tab">测试环境</a></li>
							{{end}}
							{{ if Isuserper "OE" $.uid }}
                            <li><a href="#settings" data-toggle="tab">生产环境</a></li>
							{{end}}							
                        </ul>
                        <div class="tab-content">
						
							{{ if Isuserper "DE" $.uid }}
                            <div class="tab-pane" id="activity"><!--  开发环境 start -->
                                <div class="post">
                                    <div class="box">
                                        <div class="box-header">
                                            <h3 class="box-title">开发环境</h3>
											<span class="addkey">
													<a class="btn" href="#" title="添加一行">
                                                            <i class="fa fa-plus" style="font-size:30px;"></i>
                                                    </a>
											</span>
                                        </div>
                                        <!-- /.box-header -->
                                        <div class="box-body">
                                            <table id="de" class="table table-bordered table-hover">
                                                <thead>
                                                <tr>
                                                    <th>KEY</th>
                                                    <th>VALUE</th>
													<th>Description</th>
                                                    <th>操作</th>
                                                </tr>
                                                </thead>
                                                <tbody>
												{{range .conf}}
													{{if lt .Dtstatus 2}} <!--小于2 未删除状态-->
														<tr>
		                                                    <td title="KEY不能修改，如必须修改，可以删除后，再添加" class="key">
																{{.Name}}
															</td>
		                                                    <td title="双击编辑,再次双击还原修改" class="value">
																{{.Dvalue}}
		                                                    </td>
															<td>
																{{.Description}}
															</td>
		                                                    <td>
		                                                        {{if(Isperitem "cone" $.uid)}} <!-- 判断是否有编辑权限 -->
																	<span class="save btn disabled" title="保存">
			                                                        
			                                                            <i class="fa  fa-save" style="font-size:20px;"></i>
			                                                        
																	</span>
																{{end}}
		                                                        
																{{if(Isperitem "cond" $.uid)}}
			                                                        <span class="rm_conf del" title="删除">
																	<a class="btn">
			                                                            <i class="fa fa-trash" style="font-size:20px;"></i>
			                                                        </a>
																	</span>
																{{end}}
																<i class="id" style="display:none">{{.Id}}</i>
		                                                    </td>
		                                                </tr>
													
													{{end}}
													
												{{end}} <!-- end range .conf-->
                                                
											
                                                </tbody>
												<tfoot>
                                                <tr>
                                                    <th>KEY</th>
                                                    <th>VALUE</th>
													<th>Description</th>
                                                    <th>操作</th>
                                                </tr>
                                                </tfoot>
                                            </table>
                                        </div>
                                        <!-- /.box-body -->
                                    </div>
                                </div>
                            </div> <!--  开发环境 end -->
							{{end}}
							
							{{if Isuserper "QE" $.uid }}
                            <!-- /.tab-pane -->
                            <div class="tab-pane" id="timeline"> <!-- 测试环境 start-->
                                <!-- The timeline -->
                                <div class="post">
                                    <div class="box">
                                        <div class="box-header">
                                            <h3 class="box-title">测试环境</h3>
                                        </div>
                                        <!-- /.box-header -->
                                        <div class="box-body">
                                            <table id="qe" class="table table-bordered table-hover">
                                                <thead>
                                                <tr>
                                                    <th>KEY</th>
                                                    <th>VALUE</th>
                                                    <th>操作</th>
                                                </tr>
                                                </thead>
                                                <tbody>
												
                                                {{range .conf}}
													{{ if lt .Dtstatus 3}}
													<tr>
	                                                    <td title="KEY不能修改，如必须修改，可以删除后，再添加" class="key">
															{{.Name}}
														</td>
	                                                    <td title="双击编辑,再次双击还原修改" class="value">
															{{if .Tvalue}}
																{{.Tvalue}}
															{{else}}
																等待同步
															{{end}}
	                                                    </td>
	                                                    <td>
	                                                        {{if(Isperitem "cone" $.uid)}} <!-- 判断是否有编辑权限 -->
																{{if .Dtstatus}}
																		<span class="sync btn open" title="同步:把测试环境同步到生产环境">
					                                                            <i class="fa fa-share"></i>
																		</span>
																		<span class="ignore btn open" title="忽略:忽视此KEY的更改" >
					                                                            <i class="fa fa-eye-slash"></i>
																		</span>
																	{{else}}
																		<span class="save btn disabled"  title="保存修改">
					                                                            <i class="fa fa-save"></i>
																		</span>
																	{{end}}
															{{end}}
	                                                        
															
															<i class="id" style="display:none">{{.Id}}</i>
	                                                    </td>
	                                                </tr>
													{{end}} <!--if lt 3 end -->
												{{end}} <!-- end range .conf-->
												
                                                </tbody>
												<tfoot>
                                                <tr>
                                                    <th>KEY</th>
                                                    <th>VALUE</th>
                                                    <th>操作</th>
                                                </tr>
                                                </tfoot>
                                            </table>
                                        </div>
                                        <!-- /.box-body -->
                                    </div>
                                </div>
                            </div><!-- 测试环境 end-->
							{{end}}
                            <!-- /.tab-pane -->
							{{if Isuserper "OE" $.uid }}
                            <div class="tab-pane" id="settings"><!-- 生产环境 start-->
                                <div class="post">
                                    <div class="box">
                                        <div class="box-header">
                                            <h3 class="box-title">生产环境</h3>
                                        </div>
                                        <!-- /.box-header -->
                                        <div class="box-body">
                                            <table id="oe" class="table table-bordered table-hover">
                                                <thead>
                                                <tr>
                                                    <th>KEY</th>
                                                    <th>VALUE</th>
                                                    <th>操作</th>
                                                </tr>
                                                </thead>
                                                <tbody>
												
                                                {{range .conf}}
													{{if lt .Tostatus 3}}
													<tr>
	                                                    <td title="KEY不能修改，如必须修改，可以删除后，再添加" class="key">
															{{.Name}}
														</td>
	                                                    <td title="双击编辑,再次双击还原修改" class="value">
															{{if .Ovalue}}
																{{.Ovalue}}
															{{else}}
																等待同步
															{{end}}
	                                                    </td>
														
	                                                    <td>
	                                                        {{if(Isperitem "cone" $.uid)}} <!-- 判断是否有编辑权限 -->
																{{if .Tostatus}}
																		<span class="sync btn open" title="同步:把测试环境同步到生产环境">
					                                                            <i class="fa fa-share"></i>
																		</span>
																		<span class="ignore btn open" title="忽略:忽视此KEY的更改">
					                                                            <i class="fa fa-eye-slash"></i>
																		</span>
																	{{else}}
																		<span class="save btn disabled" title="保存修改">
					                                                            <i class="fa fa-save"></i>
																		</span>
																	{{end}}
															{{end}}
	                                                        
															
															<i class="id" style="display:none">{{.Id}}</i>
	                                                    </td>
	                                                </tr>
													{{end}} <!--end if lt 3-->
												{{end}} <!-- end range .conf-->
                                                </tbody>
												<tfoot>
                                                <tr>
                                                    <th>KEY</th>
                                                    <th>VALUE</th>
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
							{{end}}  <!-- 生产环境 end-->
							<!-- 全部配置文件 -->
							<div class="active tab-pane" id="All">
                                <div class="post">
                                    <div class="box">
                                        <div class="box-header">
                                            <h3 class="box-title">汇总</h3>
                                        </div>
                                        <!-- /.box-header -->
                                        <div class="box-body">
                                            <table id="oo" class="table table-bordered table-hover">
                                                <thead>
                                                <tr>
                                                    <th>KEY</th>
                                                    <th>开发</th>
													{{ if Isuserper "QE" $.uid }}
                                                    <th>测试操作</th>
                                                    <th>测试</th>
													{{end}}
													{{ if Isuserper "OE" $.uid }}
													<th>生产操作</th>
													<th>生产</th>
													{{end}}
                                                </tr>
                                                </thead>
                                                <tbody>
												{{range $.conf}}
													{{if ne .Tostatus 3}} <!--不等于 -->
													<tr class="oo">
														<td title="{{.Name}}&#13;KEY不能修改，如必须修改，可以删除后，再添加" class="key">
															{{.Name}}
														</td>
														<td class="" title="{{.Dvalue}}&#13;{{.Description}}"><!--开发环境只能在开发页面修改-->
														{{if lt .Dtstatus 2}}
															{{.Dvalue}}
															{{ if .Dtstatus }}
																<small class="label pull-right bg-green">new</small>
															{{end}}
																
														{{else}}
															已删除
														{{end}}
                                                    	</td>
														{{ if Isuserper "QE" $.uid }}
															<td class="qe">
																{{if(Isperitem "cone" $.uid)}}
																	{{if .Dtstatus}}
																		<span class="sync btn open" title="同步:把开发环境同步到测试环境">
					                                                            <i class="fa fa-share"></i>
																		</span>
																		<span class="ignore btn open" title="忽略:忽视此KEY的更改">
					                                                            <i class="fa fa-eye-slash"></i>
																		</span>
																	{{else}}
																		<span class="save btn disabled" title="保存修改">
					                                                            <i class="fa fa-save"></i>
																		</span>
																	{{end}} <!-- end if 配置项状态-->
			                                                        
																{{end}} <!--end 配置编辑 权限 验证-->
																<i class="id" style="display:none">{{.Id}}</i>
															</td>
															<td class="value" title="{{.Tvalue}}&#13;双击修改VALUE">
															{{if lt .Dtstatus 3}}
	                                                        	{{.Tvalue}}
																{{if .Tostatus }}
																	<small class="label pull-right bg-green">new</small>
																{{end}}
															{{else}}
																已删除
															{{end}}
		                                                    </td>
														{{end}} <!--测试环境权限验证 end-->
														{{ if Isuserper "OE" $.uid }}
															<td class="oe">
																{{if(Isperitem "cone" $.uid)}}
																	{{if .Tostatus}}
																		<span class="sync btn open" title="同步:把测试环境同步到生产环境">
					                                                            <i class="fa fa-share"></i>
																		</span>
																		<span class="ignore btn open" title="忽略:忽视此KEY的更改">
					                                                            <i class="fa fa-eye-slash"></i>
																		</span>
																	{{else}}
																		<span class="save btn disabled" title="保存修改">
					                                                            <i class="fa fa-save"></i>
																		</span>
																	{{end}}
																	
			                                                        
																{{end}}
																<i class="id" style="display:none">{{.Id}}</i>
															</td>
															<td class="value" title="{{.Ovalue}}&#13;双击修改VALUE">
																{{.Ovalue}}
															</td>
														{{end}}<!--生产环境权限验证 end-->
													</tr><!--运维人员权限 --->
													{{end}} <!--end if ne 3-->
												{{end}} <!--end range-->
                                                
                                              
                                                </tbody>
												<tfoot>
	                                                <tr>
	                                                    <th>KEY</th>
	                                                    <th>开发</th>
														{{ if Isuserper "QE" $.uid }}
	                                                    <th>测试操作</th>
	                                                    <th>测试</th>
														{{end}}
														{{ if Isuserper "OE" $.uid }}
														<th>生产操作</th>
														<th>生产</th>
														{{end}}
	                                                </tr>
                                                </tfoot>
                                            </table>
                                        </div>
                                        <!-- /.box-body -->
                                    </div>
                                </div>
                            </div>
							
							<!-- 全部配置文件 end -->
							
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

    var table1 = $('#de').DataTable({
		"aLengthMenu" : [20, 40, 60],
		"iDisplayLength" : 40,
	});
    var table2 = $('#qe').DataTable({
		"aLengthMenu" : [20, 40, 60],
		"iDisplayLength" : 40,
	});
    var table3 = $('#oe').DataTable({
		"aLengthMenu" : [20, 40, 60],
		"iDisplayLength" : 40,
	});
	var table4 = $('#oo').DataTable({
		"aLengthMenu" : [20, 40, 60],
		"iDisplayLength" : 40,
	});
	
	
	$(".rm_pro").click(function(){
		var pro_id = $(this).find("i.id").text();
		var mac_name = $(this).parent().siblings("td:nth-child(1)").text();
		var node_num = $(this).parent().siblings("td:nth-child(3)").text();
		$(".verify-modal").find(".modal-title").html("警告");
		$(".verify-modal").find(".modal-body").html("确认删除项目 : <b>" + mac_name+ "</b> ,节点数量 : <b>" + node_num + "</b>");
		$(".modal").show(1000);	
		var tr = $(this).parent().parent();
		var table;
		if (tr.parent().parent().attr("id") == "example1") {
			table = table1
		} else if (tr.parent().parent().attr("id") == "example2"){
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
					if (msg == "success"){
						table1.row(tr_index).remove().draw( false );
						table2.row(tr_index).remove().draw( false );
						table3.row(tr_index).remove().draw( false );
						$("#ver_break").click();
						return;
					} else {
						alert(msg)
					}
					
				},
				error:function(XMLHttpRequest, textStatus, errorThrown) {
					alert(XMLHttpRequest.status);
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
	
	//点击添加按钮
	$(".addkey").click(function(){
		var pro_id = $("body span#pro_id").text();
		var env_sign = $(this).parent().parent().parent().parent().attr("id");
		var edit = '<span class=\"save btn disabled\" title=\"保存\">' +
	                   ' <i class=\"fa  fa-save\" style=\"font-size:20px;\"></i>' +
					'</span>' +
					
                    '<span class=\"rm_conf\" title=\"删除\">' +
					'<a class=\"btn\">' +
                        '<i class=\"fa fa-trash\" style=\"font-size:20px;\"></i>' +
                                               '</a>' +
					'</span>' +
					'<i class=\"pro_id\" style=\"display:none\">' + pro_id + '</i>'
					'<i class=\"id\" style=\"display:none\"></i>';
		t = table1
		t.row.add([
		'<input class=\"input-h-w form-control\" style="width:100%;" type=\"text\" name=\"key\"/>',
		'<input class=\" input-h-w form-control\" style="width:100%;" type=\"text\" name=\"value\"/>',
		'<input class=\" input-h-w form-control\" style="width:100%;" type=\"text\" name=\"description\"/>',
		edit
		]).draw( false );
		
		$("table tr:eq(1)").find("td:eq(1)").addClass("value");
	})
	//双击td
	var inputpla = "";//和保存按钮通信值，存放input 的placeholder
	var tdinputcontr = false;  //标识是否有正在修改的配置项
	var isnew = false;		   //当前编辑td 是否 new
	$("table tbody").on('dblclick','td', function(){
		if (inputpla == $(this).find("input").attr("placeholder")){ //第二次双击可还原
			$(this).text(inputpla);
			if (isnew == true){
				$(this).html($(this).html() + "<small class=\"label pull-right bg-green\">new</small>");
			}
			tdinputcontr = false;
			tr_obj = $(this).parent();
			if (tr_obj.find("span.save").hasClass('open')) {
				tr_obj.find("span.save").removeClass('open')
				tr_obj.find("span.save").addClass('disabled')
			}
			$(this).removeClass('select');
			inputpla = "";
			return;
		}
		if ($(this).hasClass('select')) {
			return;
		} 
		var tr = $(this).parent(); //获取当前行对象
		if(!$(this).hasClass("value")){return;}
		if (tdinputcontr == true){
			alert("当前有未保存的VALUE");
			return;
		}
		var env_sign = $(this).parent().parent().parent().attr("id");
		//var pro_id = $(this).find("i #id").text();
		if (env_sign == "de") {
			t = table1
		} else if (env_sign == "qe"){
			t = table2
		} else if (env_sign == "oe"){
			t = table3
		} else {
			t = table4
		}
		$(this).addClass("select");
		tdinputcontr = true;
		if ($(this).find("small").length>0){
			isnew = true;
		} else {
			isnew = false;
		}
		$(this).find("small").remove();
		var value = $.trim($(this).text()); //获取当前TD text
		//value = value.substring(0,value.length-4);
		inputpla = value;
		$(this).html('<input class=\" input-h-w form-control\" style="width:100%;" type=\"text\" name=\"value\" value=\"' + value + '\" placeholder=\"' +inputpla +'\"/>');

	});
	
	//把td 里面的input 去掉，直接显示<td>value</td>
	function inputtotd( tr_obj ) {
		if (tr_obj.hasClass("oo")) {
			var de_value = tr_obj.find("td:eq(1)")
			de_value.text(de_value.find("input[name='value']").val());
			var qe_value = tr_obj.find("td:eq(3)")
			qe_value.text(qe_value.find("input[name='value']").val());
			var oe_value = tr_obj.find("td:eq(5)")
			oe_value.text(oe_value.find("input[name='value']").val());
		} else {
			var td_key = tr_obj.find("td:eq(0)")
			if (td_key.text() == "") {
				td_key.text(td_key.find("input[name='key']").val());
			}
			var td_value = tr_obj.find("td:eq(1)")
			td_value.text(td_value.find("input[name='value']").val());
			var td_des = tr_obj.find("td:eq(2)")
			td_des.text(td_des.find("input[name='description']").val());
		}
	}
	
	//点击保存按钮可用和不可用切换
	function switch_save( tr_obj ) {
		if (tr_obj.find("span.save").hasClass('open')) {
			tr_obj.find("span.save").removeClass('open')
			tr_obj.find("span.save").addClass('disabled')
		} else {
			tr_obj.find("span.save").addClass('open')
			tr_obj.find("span.save").removeClass('disabled')
		}
	}
	//保存点击事件
	$("table tbody").on('click','span.save', function(){
		if ($(this).hasClass('disabled')){
			return;  //
		}

		var pro_id  = $("body span#pro_id").text();
		var conf_id = $(this).siblings("i.id").text();
		var td = $(this).parent();
		var tr = $(this).parent().parent();
		var value = tr.find("input[name='value']").val();
		if (typeof(value) == "undefined") {return;}  //value 所在td 不是input 状态，直接返回
		var key = tr.find("input[name='key']").val();//首先获取input的key 值，如没有，则直接获取td的值
		var description = tr.find("input[name='description']").val();
		if (typeof(key) == "undefined") {
			key = tr.find("td:eq(0)").text()
		}
		var td_index = $(this).parent().index();
		var table_id = $(this).parent().parent().parent().parent().attr("id");
		var previous_td_index = td_index - 1;
		var previous_td = $(this).parent().parent().find("td:nth-child(" + previous_td_index + ")");
		var env_sign;
		var tmp_env_sign = $(this).parent().parent().parent().parent().attr("id");
		if (tmp_env_sign=="oo") { //如果在汇总页面，那标识 就取当前td 的class属性值
			env_sign = $(this).parent().attr("class");
		}else {
			env_sign = tmp_env_sign;
		}
		var me = $(this)
		if (conf_id == ""){ //没有conf_id 说明是添加的
			$.ajax({
				url:"/confadd",
				type: "post",
				data:{pro_id: pro_id, key: key, value: value, sign: env_sign, description: description},
				dataType: "json",
				success: function(msg) {
					if (msg.message == "success"){
						inputtotd(tr);
						tr.find("td:eq(1)").removeClass("select");
						switch_save(tr);		//切换保存状态
						tdinputcontr = false;  //解锁编辑状态
						console.log("table_id=" + table_id);
						console.log("previous_td_index=" + previous_td_index);
						if (table_id == "de") {  //是否是开发环境div
							previous_td.addClass("value");//开发环境 td添加value  class
							td.find("i.pro_id").addClass("id").removeClass("pro_id").text(msg.confid);
						}
						return;
					} else {
						alert(msg.error);
					}
					
				},
				error:function(XMLHttpRequest, textStatus, errorThrown) {
					alert(XMLHttpRequest.status);
					if (XMLHttpRequest.status == "503" || XMLHttpRequest.status == "500" ) {
						location.href="/503"
					}
				}
			});
		} else {
			
			//成功后去掉select Class
			$.ajax({
				url:"/confedit",
				type: "post",
				data:{conf_id: conf_id, value: value, sign: env_sign,},
				dataType: "json",
				success: function(msg) {
					if (msg == "success"){
						if (tmp_env_sign == "oo"){
							index = td_index +1;
							tr.find("td:eq(" + index +")").removeClass("select");
						} else {
							tr.find("td:eq(1)").removeClass("select");
						}
						inputtotd(tr);
						me.siblings("span.ignore").remove();
						switch_save(tr);
						tdinputcontr = false;
						return;
					} else {
						alert(msg);
					}
					
				},
				error:function(XMLHttpRequest, textStatus, errorThrown) {
					alert(XMLHttpRequest.status);
					if (XMLHttpRequest.status == "503" || XMLHttpRequest.status == "500" ) {
						location.href="/503"
					}
				}
			});
			
		}
		
		
	});
	
	//input 数据改变
	$("table tbody").on('change','input', function(){
		var td = $(this).parent();
		var tr = $(this).parent().parent();
		var td_index = $(this).parent().index();
		var this_value = $.trim($(this).val());
		var table_id = $(this).parent().parent().parent().parent().attr("id");
		th_index = td_index + 1;
		if (this_value == ""){
			return;
		}
		var bro_input;
		if (td_index == 1 || td_index == 2 || td_index == 0){
			bro_input = tr.find("input[name='key']").val();
			
		}
		if (bro_input==""){
			tr.find("td span.save").removeClass("open");
			tr.find("td span.save").addClass("disabled");
			return;
		}
		//var td_span = td_index + 2;
		if (table_id == "oo") {
			tr.find("td:nth-child("+td_index+") span.save").removeClass("disabled");
			tr.find("td:nth-child("+td_index+") span.save").addClass("open");
		} else {
			tr.find("td span.save").removeClass("disabled");
			tr.find("td span.save").addClass("open");
		}
		
		sync_i = tr.find("span.sync").find("i.fa-share"); //同步按钮变保存
		sync_i.removeClass("fa-share");
		sync_i.addClass("fa-save");
		sync_span = tr.find("span.sync");
		sync_span.removeClass("sync");
		sync_span.addClass("save open");
		if (table_id == "oo"){
			
			var table_thead_tr_th_text = tr.parent().siblings("thead").find("tr th:nth-child("+th_index+")").text();
			tr.find("td:nth-child("+td_index+") span.save").attr("title","保存 : 保存" + table_thead_tr_th_text + "修改");//更改span 元素提示
		}
		
	})
	
	//删除按钮
	$("table tbody").on('click','span.rm_conf', function(){
		var conf_id = $(this).siblings("i.id").text();
		var tr = $(this).parent().parent();
		var key = tr.find("input[name='key']").val();//首先获取input的key 值，如没有，则直接获取td的值
		if (typeof(key) == "undefined") {
			key = tr.find("td:eq(0)").text()
		}
		if (conf_id == ""){
			var tr_index = table1.row(tr).index();
			table1.row(tr_index).remove().draw( false );
			return;
		}
		var inputval = tr.find("input[name='value']").val()
		if ( inputval == "") {
			alert("请先退出编辑模式");
			return;
		}
		var env_sign = $(this).parent().parent().parent().parent().attr("id");
		$(".verify-modal").find(".modal-title").html("警告");
		$(".verify-modal").find(".modal-body").html("确认删除KEY : <b>" + key+ "?</b>");
		$(".modal").show(1000);	
		abcc = function(){
			$.ajax({
				url:"/confdel",
				type: "post",
				data:{conf_id: conf_id, sign: env_sign,},
				dataType: "json",
				success: function(msg) {
					if (msg.message == "success"){
						var tr_index = table1.row(tr).index();
						table1.row(tr_index).remove().draw( false );
						$("#ver_break").click();
						return;
					} else {
						alert(msg.error);
					}
					
				},
				error:function(XMLHttpRequest, textStatus, errorThrown) {
					alert(XMLHttpRequest.status);
					if (XMLHttpRequest.status == "503" || XMLHttpRequest.status == "500" ) {
						location.href="/503"
					}
				}
			});
		}
		
	});
	
	//同步按钮
	$("table tbody").on('click','span.sync', function(){
		var conf_id = $(this).siblings("i.id").text();
		var tr = $(this).parent().parent();
		var td_index = $(this).parent().index();
		if (conf_id == ""){
			return;
		}
		var inputval = tr.find("input[name='value']").val()
		if ( inputval == "") {
			alert("请先退出编辑模式");
			return;
		}
		var env_sign;
		var tmp_env_sign = $(this).parent().parent().parent().parent().attr("id");
		if (tmp_env_sign=="oo") { //如果在汇总页面，那标识 就取当前td 的class属性值
			env_sign = $(this).parent().attr("class");
		}else {
			env_sign = tmp_env_sign;
		}
		value_td_index = td_index+1;
		var me = $(this);
		$.ajax({
				url:"/confsync",
				type: "post",
				data:{conf_id: conf_id, sign: env_sign,},
				dataType: "json",
				success: function(msg) {
					if (msg.message == "success"){
						if (tmp_env_sign =="oo"){
							tr.find("td:eq(" + value_td_index +")").text(msg.data);
							me.removeClass("sync open").addClass("save disabled");
							me.find("i.fa-share").removeClass("fa-share").addClass("fa-save");
							me.siblings("span.ignore").remove();
						}else {
							tr.find("td:eq(1)").text(msg.data);
						}
						
					} else {
						alert(msg.message);
					}
					
				},
				error:function(XMLHttpRequest, textStatus, errorThrown) {
					alert(XMLHttpRequest.status);
					if (XMLHttpRequest.status == "503" || XMLHttpRequest.status == "500" ) {
						location.href="/503"
					}
				}
			});
	});
	
	//忽略按钮 ignore
	$("table tbody").on('click','span.ignore', function(){
		var conf_id = $(this).siblings("i.id").text();
		var tr = $(this).parent().parent();
		var td_index = $(this).parent().index();
		if (conf_id == ""){
			return;
		}
		var inputval = tr.find("input[name='value']").val()
		if ( inputval == "") {
			alert("请先退出编辑模式");
			return;
		}
		var env_sign;
		var tmp_env_sign = $(this).parent().parent().parent().parent().attr("id");
		if (tmp_env_sign=="oo") { //如果在汇总页面，那标识 就取当前td 的class属性值
			env_sign = $(this).parent().attr("class");
		}else {
			env_sign = tmp_env_sign;
		}
		value_td_index = td_index+1;
		var me = $(this);
		
		$.ajax({
				url:"/confignore",
				type: "post",
				data:{conf_id: conf_id, sign: env_sign,},
				dataType: "json",
				success: function(msg) {
					if (msg.message == "success"){
						if (tmp_env_sign =="oo"){
							tr.find("td:eq(" + value_td_index +")").text(msg.data);
							me.siblings("span.sync").removeClass("sync open").addClass("save disabled");
							me.siblings("span.save").find("i.fa-share").removeClass("fa-share").addClass("fa-save");
							me.remove();
						}else {
							tr.find("td:eq(1)").text(msg.data);
						}
						
					} else {
						alert(msg.message);
					}
					
				},
				error:function(XMLHttpRequest, textStatus, errorThrown) {
					alert(XMLHttpRequest.status);
					if (XMLHttpRequest.status == "503" || XMLHttpRequest.status == "500" ) {
						location.href="/503"
					}
				}
		});
	});
</script>
</body>
</html>

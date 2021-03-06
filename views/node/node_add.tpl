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
	<!-- Select2 -->
	<link rel="stylesheet" href="/static/plugins/select2/select2.min.css">
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
        .white_content{
            display:none;
            position: absolute;
            top: 25%;  left: 25%;
            width: 50%;  height: 50%;
            padding: 16px;  border: 4px solid rgba(66, 62, 74, 0.43);
            background-color: white;  z-index:1002;  overflow: auto;
        }
         .black_overlay{
             display: none;
             position: absolute;
             top: 0%;  left: 0%;
             width: 100%;  height: 100%;
             background-color: black;  z-index:1001;  -moz-opacity: 0.8;  opacity:.80;  filter: alpha(opacity=80);
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
<div class="wrapper">
	<span style="display:none" class="pro_id">{{.pro.Id}}</span>
	<span style="display:none" class="env_sign">{{.env.Sign}}</span>
    {{template "common/headertitle.tpl" .}}
    <!-- Left side column. contains the logo and sidebar -->
    {{template "common/sidebar.tpl" .}}
    <!-- Content Wrapper. Contains page content -->
    <div class="content-wrapper">
        <!-- Content Header (Page header) -->
        <section class="content-header">
            <h1>添加节点
                <small>
                    <button type="button" class="btn btn-block btn-primary btn-xs" onclick="history.go(-1)">返回列表
                    </button>
                </small>
            </h1>
            <ol class="breadcrumb">
                <li><a href="/"><i class="fa fa-dashboard"></i>主页</a></li>
                <li><a href="/{{.pro.Id}}/{{.env.Sign}}/nodelist">{{.pro.Name}}{{.env.Name}}</a></li>
                <li class="active">节点添加</li>
            </ol>
        </section>

        <!-- Main content -->
        <section class="content">
            <div class="row">
                <div class="col-md-12">
                    <div class="box">
                        <div class="box-header">
                            <h3 class="box-title">项目<<b>{{.pro.Name}}</b>></h3>
							<h4>环境<<b>{{.env.Name}}</b>></h4>
                        </div>
                        <!-- /.box-header -->
                        <div class="box-body">
                            <form class="form-horizontal" action="#">
								<input type="text" name="mirror" style="display:none" />
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">节点别名</label>

                                    <div class="col-sm-10">
                                        <input type="text" class="form-control" id="inputName" name="sign" placeholder="用于标识节点唯一性">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">所属项目</label>

                                    <div class="col-sm-10">
                                        {{.pro.Name}}
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">所属环境</label>

                                    <div class="col-sm-10">
                                        {{.env.Name}}
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">隶属主机</label>

                                    <div class="col-sm-10">
                                        <select class="form-control select2" name="machine" data-placeholder="选择主机"style="width: 100%;">
                                            {{range .macs}}
												<option value="{{.Id}}">{{.Ipaddr2}}</option>
											{{end}}
                                        </select>

                                    </div>
                                </div>
                                <div class="form-group">
                                    <label for="inputExperience" class="col-sm-2 control-label">端口映射</label>

                                    <div class="col-sm-10">
                                        <textarea class="form-control" id="inputExperience" name="port"
                                                  placeholder="逗号分隔(英文!!)，如80:80,90:90"></textarea>
                                    </div>
                                </div>
                                <div class="form-group">
                                        <label for="inputName" class="col-sm-2 control-label">引用镜像:</label>
                                        <div class="col-sm-10">
                                            <span id="imageselection_id" class="imageselection btn" style="color:#3c8dbc;">
                                                点击选择
                                            </span>

                                        </div>
                                    </div>
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">cpu:</label>

                                    <div class="col-sm-10">
                                        Xen2520 3.5MHZ 2核心
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">内存信息:</label>

                                    <div class="col-sm-10">
                                        4098M
                                    </div>
                                </div>
								</form>
                                <div class="form-group">
                                    <div class="col-sm-offset-2 col-sm-10">
                                        <button type="submit" id="submit" class="btn btn-danger">Submit</button>
                                    </div>
                                </div>
                            
                        </div>
                    </div>
                </div>
            </div>
            <!-- /.row -->
        </section>
        <!-- /.content -->
    </div>
    {{template "common/footitle.tpl" .}}
</div>

<!--        添加节点弹窗---start---->
<div id="fade" class="black_overlay"></div> <!-- 半透明不可操作 -->
    <div class="white_content">
        <div class="pull-right box-tools full-sreen">
            <!-- button with a dropdown -->
            <div class="btn-group full-sreen">
			
            <table id="example1" class="table table-bordered table-striped dataTable" role="grid" aria-describedby="example1_info">
				{{range $i, $v := Getmirgroup}}
				<thead>
               		<tr role="row">
						<th class="" tabindex="0" aria-controls="example1" rowspan="1" colspan="3" aria-sort="ascending" aria-label="Rendering engine: activate to sort column descending" style="width: 290px;text-align: center;">{{$v.Name}}:</th>
					</tr>
					
				</thead>
				<tbody>
					
	               <tr role="row" class="odd">
						{{range $index, $val := Getmir $v.Id}}
							{{if Seek $index 3}}
								 <td width="33%" title="{{$val.Hubaddress}}">
										<input name="imagename"  type="radio" value="{{$val.Id}}" /><span>{{$val.Name}}</span>
								</td>
							{{else}}	
								</tr>
								<tr role="row" class="odd">
								<td width="33%" title="{{$val.Hubaddress}}">
										<input name="imagename"  type="radio" value="{{$val.Id}}" /><span>{{$val.Name}}</span>
								</td>
							{{end}}
						{{end}}
	               </tr>
					
				</tbody>
               	{{end}}
				<thead><!-- 添加按钮 -->
               			<tr role="row"><th class="" tabindex="0" aria-controls="example1" rowspan="1" colspan="3" aria-sort="ascending" aria-label="Rendering engine: activate to sort column descending" style="width: 290px; text-align: center;">
							<button id="addimage" type="button" class="btn btn-default pull-right"><i class="fa fa-plus"></i> Add </button>
						</th></tr>
               		</thead>

               
            </table>

        </div>

    </div>

</div>

<!--        添加节点弹窗---end------>

<!-- 实时打印信息--->
	<div id="fade" class="modal"> 
		<div class="col-md-3 center">
          <div class="box box-primary ">
            <div class="box-header with-border">
              <h3 class="box-title">Collapsable</h3>

              <div id="ms_title" class="box-tools pull-right">
                <button type="button" class="btn btn-box-tool" data-widget="collapse"><i class="fa fa-minus"></i>
                </button>
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

<script src="/static/plugins/select2/select2.full.min.js"></script>
<!-- SlimScroll -->
<script src="/static/plugins/slimScroll/jquery.slimscroll.min.js"></script>

<!-- AdminLTE App -->
<script src="/static/js/app.min.js"></script>
<!-- AdminLTE for demo purposes -->
<script src="/static/js/demo.js"></script>
<!--websocket-->
<script src="/static/js/nodews.js"></script>
<!-- page script -->
<script>
    $(function () {
        //Initialize Select2 Elements
        $(".select2").select2();
    });

    $(document).ready(function(){
		
		var imagename1 ;
        //点击显示选择镜像
        var cont_view = true;
       $("#imageselection_id").click(function(){
			
           if (cont_view){
               $(".white_content").show();
               $(".black_overlay").show();

               cont_view = false;
           }else {
               $(".white_content").hide();
               $(".black_overlay").hide();
               cont_view = true;
           }
       });

 
		$("#addimage").click(function(){
			imagename1 = $("input:radio[name='imagename']:checked + span").text() ;
			if (imagename1 == ""){
				alert("镜像不能为空");
				return;
			}
			image = $("input:radio[name='imagename']:checked").val();
			$("#imageselection_id").text(imagename1);
			$("form.form-horizontal").find("input[name='mirror']").val(image);
			if (cont_view){
              $(".white_content").show();
              $(".black_overlay").show();

              cont_view = false;
			}else {
			    $(".white_content").hide();
			    $(".black_overlay").hide();
			    cont_view = true;
			}
			});
    });
	
	var pro_id = $("span.pro_id").text();
	var env_sign = $("span.env_sign").text();
	$("#submit").click(function(){
		
		var n_form = $("form.form-horizontal");
		var n_mirror = n_form.find("input[name='mirror']").val();
		var n_sign = n_form.find("input[name='sign']").val();
		var n_machine = n_form.find("select[name='machine']").val();
		var n_port = n_form.find("textarea[name='port']").val();
		alert("m:" + n_mirror + "s:" + n_sign + "ma:" + n_machine + "p:" + n_port);
		$.ajax({
				url:"/" + pro_id + "/" + env_sign + "/nodeadd",
				type: "post",
				data:{
					mirr: n_mirror,
					sign: n_sign,
					machine: n_machine,
					port: n_port,
				},
				dataType: "json",
				success: function(msg) {
					if (msg.message == "success"){
						console.log(msg.data);
						socket.send(msg.data);
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
	});
	
	var aabbcc=function() {}
	//<!--$("#ms_title").append(addclose)-->
	var addclose = '<button type=\"button\" class=\"btn btn-box-tool\" data-widget=\"remove\"><i class=\"fa fa-times\"></i></button>';
	var ms_boxy = $("div.center").find(".box-body");
	var ms_title = $("div.center").find(".box-title");
	$("div.center").on("click", "button[data-widget='remove']", function(){
		$("div.modal").hide(1000);
		aabbcc();
	});
	
</script>
</body>
</html>
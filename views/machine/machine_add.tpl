<!DOCTYPE html>
<html>
<head>
    {{template "common/meta.tpl" .}}
    <!-- Bootstrap 3.3.6 -->
    <link rel="stylesheet" href="/static/bootstrap/css/bootstrap.min.css">
    <script src="/static/plugins/input-mask/jquery.inputmask.js"></script>
    <script src="/static/plugins/input-mask/jquery.inputmask.date.extensions.js"></script>
    <script src="/static/plugins/input-mask/jquery.inputmask.extensions.js"></script>
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
<div class="wrapper">
    {{template "common/headertitle.tpl" .}}
    <!-- Left side column. contains the logo and sidebar -->
    {{template "common/sidebar.tpl" .}}
    <!-- Content Wrapper. Contains page content -->
    <div class="content-wrapper">
        <!-- Content Header (Page header) -->
        <section class="content-header">
            <h1>添加主机
                <small>
				<a href="/maclist">
                    <button type="button" class="btn btn-block btn-primary btn-xs">返回列表
                    </button>
				</a>
                </small>
            </h1>
            <ol class="breadcrumb">
                <li><a href="#"><i class="fa fa-dashboard"></i>主页</a></li>
                <li class="active">主机添加</li>
            </ol>
        </section>

        <!-- Main content -->
        <section class="content">
            <div class="row">
                <div class="col-md-12">
                    <div class="box">
                        <div class="box-body">
                            <form class="form-horizontal" method="post">
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">主机名称</label>
                                    <div class="col-sm-10">
                                        <input type="text" class="form-control" id="inputName" name="name" value="{{.mac.Name}}"
											placeholder="中文，用于登记主机名称">
                                    </div>
									{{if .nameerr}}
										<lable style="margin-left:18%;">
												<small class="label label-danger"> {{.nameerr}}</small>
										</label>
									{{end}}
                                </div>                                
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">外网ip</label>
                                    <div class="col-sm-10">
                                        <input type="text" class="form-control" id="inputName" name="ipaddr1" placeholder="有外网请填写外网ip">
                                    </div>
									
                                </div>
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">内网ip</label>
                                    <div class="col-sm-10">
                                        <input type="text" class="form-control" id="inputName" name="ipaddr2" placeholder="有内网请填写内网ip">                                    	
									</div>
									{{if .ipaddr2err}}
										<lable style="margin-left:18%;">
												<small class="label label-danger"> {{.ipaddr2err}}</small>
										</label>
									{{end}}
                                </div>
								<div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">管理接口</label>
                                    <div class="col-sm-10">                                       
                                            <label>
                                                <input type="radio" name="interface" value="1" class="minimal" checked>内网
                                            </label>
                                            <label>
                                                <input type="radio" name="interface" value="2" class="minimal">外网
                                            </label> 
											<label>
												<input type="text" class="form-control" id="inputName" name="port" placeholder="端口"> 
											</label>                                      
                                    </div>
									{{if .porterr}}
										<lable style="margin-left:18%;">
												<small class="label label-danger"> {{.porterr}}</small>
										</label>
									{{end}}
                                </div>
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">硬件信息</label>
                                    <div class="col-sm-10">
                                        <input type="text" class="form-control" id="inputName" name="context" placeholder="硬件信息">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">所属环境</label>
                                    <div class="col-sm-10">
                                        <select class="form-control select2" name="env" style="width: 100%;">
											{{range Getenv}}
												<option value="{{.Id}}">{{.Name}}</option>
											{{end}}
                                            
                                        </select>
                                    </div>
									{{if .env2err}}
										<lable style="margin-left:18%;">
												<small class="label label-danger"> {{.enverr}}</small>
										</label>
									{{end}}
                                </div>
                        </div>
                        <div class="form-group">
                            <div class="col-sm-offset-2 col-sm-10">
                                <button type="submit" class="btn btn-danger">Submit</button>
                            </div>
                        </div>
                        </form>
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
<!-- ./wrapper -->

<!-- jQuery 2.2.3 -->
<script src="/static/plugins/jQuery/jquery-2.2.3.min.js"></script>
<!-- Bootstrap 3.3.6 -->
<script src="/static/bootstrap/js/bootstrap.min.js"></script>
<!-- DataTables -->
<script src="/static/plugins/datatables/jquery.dataTables.min.js"></script>
<script src="/static/plugins/datatables/dataTables.bootstrap.min.js"></script>
<!-- Select2 -->
<link rel="stylesheet" href="/static/plugins/select2/select2.min.css">
<script src="/static/plugins/select2/select2.full.min.js"></script>
<!-- SlimScroll -->
<script src="/static/plugins/slimScroll/jquery.slimscroll.min.js"></script>
<script src="/static/plugins/iCheck/icheck.min.js"></script>
<!-- FastClick -->
<script src="/static/plugins/fastclick/fastclick.js"></script>
<!-- AdminLTE App -->
<script src="/static/js/app.min.js"></script>
<!-- AdminLTE for demo purposes -->
<script src="/static/js/demo.js"></script>
<!-- page script -->
<script>
    $("[data-mask]").inputmask();
</script>
</body>
</html>
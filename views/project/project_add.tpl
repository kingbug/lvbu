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
<div class="wrapper">
    {{template "common/headertitle.tpl" .}}
    <!-- Left side column. contains the logo and sidebar -->
    {{template "common/sidebar.tpl" .}}
    <!-- Content Wrapper. Contains page content -->
    <div class="content-wrapper">
        <!-- Content Header (Page header) -->
        <section class="content-header">
            <h1>添加项目               
            </h1>
            <ol class="breadcrumb">
                <li><a href="/login"><i class="fa fa-dashboard"></i>主页</a></li>                
                <li class="active">项目添加</li>
            </ol>
        </section>

        <!-- Main content -->
        <section class="content">
            <div class="row">
                <div class="col-md-12">
                    <div class="box">                       
                        <div class="box-body">
                            <form class="form-horizontal">
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">项目名称</label>
                                    <div class="col-sm-10">
                                        <input type="text" class="form-control" id="inputName" placeholder="中文，用于登记项目名称">
                                    </div>
                                </div>
								<div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">项目标识</label>
                                    <div class="col-sm-10">
                                        <input type="text" class="form-control" id="inputName" placeholder="字母+数字，用于项目唯一性标识">
                                    </div>
                                </div>
								<div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">仓库地址</label>
                                    <div class="col-sm-10">
                                        <input type="text" class="form-control" id="inputName" placeholder="仅支持git">
                                    </div>
                                </div>
								<div class="form-group">								
                                    <label for="inputName" class="col-sm-2 control-label">帐号帐号</label>									
                                    <div class="col-sm-10">
                                        <input type="text" class="form-control" id="inputName" placeholder="git帐号密码，冒号分割。如admin:admin">
                                    </div>
                                </div>
								<div class="form-group">
                                    <label for="inputExperience" class="col-sm-2 control-label">忽略文件</label>

                                    <div class="col-sm-10">
                                        <textarea class="form-control" id="inputExperience"
                                                  placeholder="git需要忽略的文件"></textarea>
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
<!-- FastClick -->
<script src="/static/plugins/fastclick/fastclick.js"></script>
<!-- AdminLTE App -->
<script src="/static/js/app.min.js"></script>
<!-- AdminLTE for demo purposes -->
<script src="/static/js/demo.js"></script>
<!-- page script -->
<script>
    $(function () {
        //Initialize Select2 Elements
        $(".select2").select2();
    });
</script>
</body>
</html>
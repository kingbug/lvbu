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
            <h1>编辑节点
                <small>
                    <button type="button" class="btn btn-block btn-primary btn-xs" onclick="history.go(-1)">返回列表
                    </button>
                </small>
            </h1>
            <ol class="breadcrumb">
                <li><a href="#"><i class="fa fa-dashboard"></i>主页</a></li>
                <li><a href="#">XX项目(XX环境)</a></li>
                <li class="active">节点编辑</li>
            </ol>
        </section>

        <!-- Main content -->
        <section class="content">
            <div class="row">
                <div class="col-md-12">
                    <div class="box">
                        <div class="box-header">
                            <h3 class="box-title">XX项目XX环境XX节点</h3>
                        </div>
                        <!-- /.box-header -->
                        <div class="box-body">
                            <form class="form-horizontal">
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">节点ID</label>

                                    <div class="col-sm-10">
                                        asfdasfadsf111111
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">所属环境</label>

                                    <div class="col-sm-10">
                                        运维环境
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">隶属主机</label>

                                    <div class="col-sm-10">
                                        <a href="/machine">192.168.2.1</a>
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">节点别名</label>

                                    <div class="col-sm-10">
                                        <input type="text" class="form-control" id="inputName" placeholder="用于标识节点唯一性">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label for="inputExperience" class="col-sm-2 control-label">端口映射</label>

                                    <div class="col-sm-10">
                                        <textarea class="form-control" id="inputExperience"
                                                  placeholder="逗号分隔(英文!!)，如80:80,90:90"></textarea>
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">引用镜像:</label>
                                    <div class="col-sm-10">
                                        <a href="/mirror">J2EE TOMCAT1.7</a>
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
<!-- SlimScroll -->
<script src="/static/plugins/slimScroll/jquery.slimscroll.min.js"></script>
<!-- FastClick -->
<script src="/static/plugins/fastclick/fastclick.js"></script>
<!-- AdminLTE App -->
<script src="/static/js/app.min.js"></script>
<!-- AdminLTE for demo purposes -->
<script src="/static/js/demo.js"></script>
<!-- page script -->
</body>
</html>
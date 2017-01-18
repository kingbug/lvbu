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
            <h1>
                节点列表
                <small><a href="/nodeadd">
                    <button type="button" class="btn btn-block btn-primary btn-xs">添加节点</button>
                </a></small>
            </h1>
            <ol class="breadcrumb">
                <li><a href="#"><i class="fa fa-dashboard"></i>主页</a></li>
                <li><a href="#">项目列表(XX环境)</a></li>
                <li class="active">XX项目</li>
            </ol>
        </section>

        <!-- Main content -->
        <section class="content">
            <div class="row">
                <div class="col-xs-12">
                    <div class="box">
                        <div class="box-header">
                            <h3 class="box-title">XX项目XX环境节点列表</h3>
                        </div>
                        <!-- /.box-header -->
                        <div class="box-body">
                            <table id="example1" class="table table-bordered table-striped">
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
                                <tr>
                                    <td><input type="checkbox"><a href="/nodedit">接入服务1</a></td>
                                    <td><a href="/machine">192.168.2.1</a></td>
                                    <td>V1.0.0.1
                                    </td>
                                    <td>Win 95+</td>
                                    <td><a class="btn">
                                        <i class="fa fa-rocket"></i>
                                    </a>
                                        <a class="btn">
                                            <i class="fa fa-pause"></i>
                                        </a>
                                        <a class="btn">
                                            <i class="fa fa-repeat"></i>
                                        </a>
                                        <a class="btn">
                                            <i class="fa fa-trash"></i>
                                        </a></td>
                                </tr>
                                <td><input type="checkbox"><a href="/nodedit">接入服务2</a></td>
                                <td><a href="/machine">192.168.2.1</a></td>
                                <td>V1.0.0.1
                                </td>
                                <td>Win 95+</td>
                                <td><a class="btn">
                                    <i class="fa fa-rocket"></i>
                                </a>
                                    <a class="btn">
                                        <i class="fa fa-pause"></i>
                                    </a>
                                    <a class="btn">
                                        <i class="fa fa-repeat"></i>
                                    </a>
                                    <a class="btn">
                                        <i class="fa fa-trash"></i>
                                    </a></td>
                                </tr>
                                <td><input type="checkbox"><a href="/nodedit">接入服务3</a></td>
                                <td><a href="/machine">192.168.2.1</a></td>
                                <td>V1.0.0.1
                                </td>
                                <td>Win 95+</td>
                                <td><a class="btn">
                                    <i class="fa fa-rocket"></i>
                                </a>
                                    <a class="btn">
                                        <i class="fa fa-pause"></i>
                                    </a>
                                    <a class="btn">
                                        <i class="fa fa-repeat"></i>
                                    </a>
                                    <a class="btn">
                                        <i class="fa fa-trash"></i>
                                    </a></td>
                                </tr>
                                <td><input type="checkbox"><a href="/nodedit">接入服务4</a></td>
                                <td><a href="/machine">192.168.2.1</a></td>
                                <td>V1.0.0.1
                                </td>
                                <td>Win 95+</td>
                                <td><a class="btn">
                                    <i class="fa fa-rocket"></i>
                                </a>
                                    <a class="btn">
                                        <i class="fa fa-pause"></i>
                                    </a>
                                    <a class="btn">
                                        <i class="fa fa-repeat"></i>
                                    </a>
                                    <a class="btn">
                                        <i class="fa fa-trash"></i>
                                    </a></td>
                                </tr>
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
    $(function () {
        $('#example1').DataTable({
            "paging": true,
            "lengthChange": false,
            "searching": false,
            "ordering": false,
            "info": true,
            "autoWidth": false
        });
    });
</script>
</body>
</html>
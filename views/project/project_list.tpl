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
                        <ul class="nav nav-tabs">
                            <li class="active"><a href="#activity" data-toggle="tab">开发环境</a></li>
                            <li><a href="#timeline" data-toggle="tab">测试环境</a></li>
                            <li><a href="#settings" data-toggle="tab">生产环境</a></li>							
                        </ul>
                        <div class="tab-content">
                            <div class="active tab-pane" id="activity">
                                <div class="post">
                                    <div class="box">
                                        <div class="box-header">
                                            <h3 class="box-title">开发环境</h3>
                                        </div>
                                        <!-- /.box-header -->
                                        <div class="box-body">
                                            <table id="example1" class="table table-bordered table-hover">
                                                <thead>
                                                <tr>
                                                    <th>项目名称</th>
                                                    <th>项目状态</th>
                                                    <th>节点数量(s)</th>
                                                    <th>操作</th>
                                                </tr>
                                                </thead>
                                                <tbody>
												 {{range $key, $val :=.macd}}
                                                <tr>
                                                    <td>账户中心</td>
                                                    <td><span class="label label-success">正常运行</span>
                                                    </td>
                                                    <td><a href="/nodelist">5</a></td>
                                                    <td>
                                                        <a class="btn">
                                                            <i class="fa fa-edit">编辑</i>
                                                        </a>
                                                        <a class="btn">
                                                            <i class="fa fa-trash">删除</i>
                                                        </a>
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
                            <div class="tab-pane" id="timeline">
                                <!-- The timeline -->
                                <div class="post">
                                    <div class="box">
                                        <div class="box-header">
                                            <h3 class="box-title">测试环境</h3>
                                        </div>
                                        <!-- /.box-header -->
                                        <div class="box-body">
                                            <table id="example2" class="table table-bordered table-hover">
                                                <thead>
                                                <tr>
                                                    <th>项目名称</th>
                                                    <th>项目状态</th>
                                                    <th>节点数量(s)</th>
                                                    <th>操作</th>
                                                </tr>
                                                </thead>
                                                <tbody>
												{{range $key, $val :=.macq}}
                                                <tr>
                                                    <td>账户中心</td>
                                                    <td><span class="label label-success">正常运行</span>
                                                    </td>
                                                    <td><a href="/nodelist">5</a></td>
                                                    <td>
                                                        <a class="btn">
                                                            <i class="fa fa-edit">编辑</i>
                                                        </a>
                                                        <a class="btn">
                                                            <i class="fa fa-trash">删除</i>
                                                        </a>
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

                            <div class="tab-pane" id="settings">
                                <div class="post">
                                    <div class="box">
                                        <div class="box-header">
                                            <h3 class="box-title">生产环境</h3>
                                        </div>
                                        <!-- /.box-header -->
                                        <div class="box-body">
                                            <table id="example3" class="table table-bordered table-hover">
                                                <thead>
                                                <tr>
                                                    <th>项目名称</th>
                                                    <th>项目状态</th>
                                                    <th>节点数量(s)</th>
                                                    <th>操作</th>
                                                </tr>
                                                </thead>
                                                <tbody>
                                                <tr>
                                                    <td>账户中心</td>
                                                    <td><span class="label label-success">正常运行</span>
                                                    </td>
                                                    <td><a href="/nodelist">5</a></td>
                                                    <td>
                                                        <a class="btn">
                                                            <i class="fa fa-edit">编辑</i>
                                                        </a>
                                                        <a class="btn">
                                                            <i class="fa fa-trash">删除</i>
                                                        </a>
                                                    </td>
                                                </tr>
                                                <tr>
                                                    <td>运营后台</td>
                                                    <td><span class="label label-warning">运行警告</span>
                                                    </td>
                                                    <td><a href="/nodelist">5</a></td>
                                                    <td>
                                                        <a class="btn">
                                                            <i class="fa fa-edit">编辑</i>
                                                        </a>
                                                        <a class="btn">
                                                            <i class="fa fa-trash">删除</i>
                                                        </a>
                                                    </td>
                                                </tr>
                                                <tr>
                                                    <td>接入项目</td>
                                                    <td><span class="label label-danger">项目异常</span>
                                                    </td>
                                                    <td><a href="/nodelist">5</a></td>
                                                    <td>
                                                        <a class="btn">
                                                            <i class="fa fa-edit">编辑</i>
                                                        </a>
                                                        <a class="btn">
                                                            <i class="fa fa-trash">删除</i>
                                                        </a>
                                                    </td>
                                                </tr>
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
        $('#example1').DataTable();
        $('#example2').DataTable();
        $('#example3').DataTable();
    });
</script>
</body>
</html>

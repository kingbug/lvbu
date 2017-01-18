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
                            <li class="active"><a href="#basemirr" data-toggle="tab">基础镜像</a></li>
                            <li><a href="#appmirr" data-toggle="tab">应用镜像</a></li>
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
                            <!-- /.tab-pane -->
                            <div class="tab-pane" id="appmirr">
                                <!-- The timeline -->
                                <div class="post">
                                    <div class="box">                                        
                                        <div class="box-body">
											<table id="example1" class="table table-bordered table-hover">
                                                <thead>
                                                <tr>
                                                    <th>镜像名称</th>
                                                    <th>基础镜像</th>
                                                    <th>仓库地址</th>
													<th>操作</th>
                                                </tr>
                                                </thead>
                                                <tbody>
                                                <tr>
                                                    <td>JDK1.7+TOMCAT7.3.44</td>
                                                    <td>ubuntu</td>
                                                    <td>docker.cihi.cn:567/centosphpnginx:latest</td>
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
                                                    <td>JDK1.7+TOMCAT7.3.44</td>
                                                    <td>ubuntu</td>
                                                    <td>docker.cihi.cn:567/centosphpnginx:latest</td>
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
                                                    <td>nginx</td>
                                                    <td>ubuntu</td>
                                                    <td>docker.cihi.cn:567/centosphpnginx:latest1</td>
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
                            </div>                           
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

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
            <h1>关于系统
            </h1>
            <ol class="breadcrumb">
                <li><a href="#"><i class="fa fa-dashboard"></i>主页</a></li>
                <li class="active">关于</li>
            </ol>
        </section>
        <section class="content">

            <!-- row -->
            <div class="row">
                <div class="col-md-12">
                    <!-- The time line -->
                    <ul class="timeline">
                        <!-- timeline time label -->
                        <li class="time-label">
                  <span class="bg-red">
                    2017年1月1日
                  </span>
                        </li>
                        <!-- /.timeline-label -->
                        <!-- timeline item -->
                        <li>
                            <i class="fa fa-envelope bg-blue"></i>
                            <div class="timeline-item">
                                <span class="time"><i class="fa fa-clock-o"></i> 00:00</span>
                                <h3 class="timeline-header">系统立项V1.0</h3>
                                <div class="timeline-body">
                                    主要为解决开发测试及线上环境统一的快速部署化。
									1，采用docker作为容器，git作为仓库。
									2，采用golang作为开发语言，mysql存储。									
                                </div>                                
                            </div>
                        </li>
                        <!-- END timeline item -->
                        <!-- timeline item -->
                        <li>
                            <i class="fa fa-user bg-aqua"></i>
                            <div class="timeline-item">
                                <span class="time"><i class="fa fa-clock-o"></i> 00:00</span>
                                <h3 class="timeline-header no-border">使用memory、redis等可配置缓存</h3>
                            </div>
                        </li>
                        <!-- END timeline item -->
                        <!-- timeline item -->
                        <li>
                            <i class="fa fa-comments bg-yellow"></i>
                            <div class="timeline-item">
                                <span class="time"><i class="fa fa-clock-o"></i>00:00</span>
                                <h3 class="timeline-header">页面采用响应式布局</h3>
                                <div class="timeline-body">
                                    为解决运维的anywhere，采用bootstarp方案。
                                </div>                                
                            </div>
                        </li>
                        <!-- END timeline item -->
                        <!-- timeline time label -->
                        <li class="time-label">
                  <span class="bg-green">
                    2017年1月5日
                  </span>
                        </li>
                        <!-- /.timeline-label -->                       
                        <!-- timeline item -->
                        <li>
                            <i class="fa fa-comments bg-yellow"></i>
                            <div class="timeline-item">
                                <span class="time"><i class="fa fa-clock-o"></i>00:00</span>
                                <h3 class="timeline-header">增加安全设置</h3>
                                <div class="timeline-body">
                                    根据职位分配权限方案。
                                </div>                                
                            </div>
                        </li>
                        <!-- END timeline item -->
                        <li>
                            <i class="fa fa-clock-o bg-gray"></i>
                        </li>
                    </ul>
                </div>
                <!-- /.col -->
            </div>
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
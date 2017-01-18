<!DOCTYPE html>
<html>
<head>
    {{template "common/meta.tpl" .}}
    <!-- Bootstrap 3.3.6 -->
    <link rel="stylesheet" href="static/bootstrap/css/bootstrap.min.css">
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
                系统管理
            </h1>
            <ol class="breadcrumb">
                <li><a href="#"><i class="fa fa-dashboard"></i>主页</a></li>
                <li class="active">系统管理</li>
            </ol>
        </section>

        <!-- Main content -->
        <section class="content">
            <div class="row">
                <div class="col-md-12">
                    <div class="nav-tabs-custom">
                        <ul class="nav nav-tabs">
                            <li class="active"><a href="#userslist" data-toggle="tab">用户列表</a></li>
                            <li><a href="#poslist" data-toggle="tab">职位管理</a></li>
                            <li><a href="#settings" data-toggle="tab">系统参数</a></li>
                        </ul>
                        <div class="tab-content">
                            <div class="active tab-pane" id="userslist">
                                <div class="post">
                                    <div class="box">
                                        <div class="box-header">
                                            <h3 class="box-title">用户列表</h3>
                                            <label>
                                                <a href="/useradd">
                                                    <button type="button" class="btn btn-block btn-primary btn-xs">
                                                        添加用户
                                                    </button>
                                                </a></label>
                                        </div>
                                        <!-- /.box-header -->
                                        <div class="box-body">
                                            <table id="example1" class="table table-bordered table-hover">
                                                <thead>
                                                <tr>
                                                    <th>姓名</th>
                                                    <th>帐号</th>
                                                    <th>职位</th>
                                                    <th>性别</th>
                                                    <th>电话</th>
                                                    <th>创建时间</th>
                                                    <th>最后登录</th>
                                                    <th>操作</th>
                                                </tr>
                                                </thead>
                                                <tbody>
                                                <tr>
												 {{range $key, $val :=.users}}
                                                    <td>{{$val.Nick}}
													{{if eq $val.Status 0}}
													<a href="/lockuser/{{$val.Id}}"><i class="fa fa-lock text-blue"></i></a>
													{{else}}
													<a href="/unlockuser/{{$val.Id}}"><i class="fa fa-key text-red"></i></a>
													{{end}}
                                                    </td>
                                                    <td>{{$val.UserName}}</td>
                                                    <td>{{Getposname $val.Position.Id}}</td>
                                                    <td>{{Getsex $val.Sex}}</td>
                                                    <td>{{$val.Phone}}</td>
                                                    <td>{{$val.Created}}</td>
                                                    <td>{{$val.Updated}}</td>
                                                    </td>
                                                    <td>
                                                        <a class="btn" href="/useredit">
                                                            <i class="fa fa-edit"></i>
                                                        </a>
                                                        <a class="btn">
                                                            <i class="fa fa-trash"></i>
                                                        </a>
                                                    </td>
                                                </tr>
												{{end}}
                                                </tbody>
                                                <tfoot>
                                                <tr>
                                                    <th>姓名</th>
                                                    <th>帐号</th>
                                                    <th>职位</th>
                                                    <th>性别</th>
                                                    <th>电话</th>
                                                    <th>创建时间</th>
                                                    <th>最后登录</th>
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
                            <div class="tab-pane" id="poslist">
                                <!-- The timeline -->
                                <div class="post">
                                    <div class="box">
                                        <div class="box-header">
                                            <h3 class="box-title">职位管理</h3>
                                            <label>
                                                <a href="/posadd">
                                                    <button type="button" class="btn btn-block btn-primary btn-xs">
                                                        添加职位
                                                    </button>
                                                </a></label>
                                        </div>
                                        <!-- /.box-header -->
                                        <div class="box-body">
                                            <table id="example2" class="table table-bordered table-hover">
                                                <thead>
                                                <tr>
                                                    <th>职位名称</th>
                                                    <th>职位标识</th>
                                                    <th>操作</th>
                                                </tr>
                                                </thead>
                                                <tbody>
												 {{range $key, $val :=.poss}}
                                                <tr>
                                                    <td>{{$val.Name}}</td>
                                                    <td>{{$val.Sign}}</td>
                                                    <td>
                                                        <a class="btn">
                                                            <i class="fa fa-users">成员</i>
                                                        </a>
                                                        <a class="btn" href="/permanage/{{$val.Id}}">
                                                            <i class="fa fa-shield">权限</i>
                                                        </a>
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
                                                    <th>职位名称</th>
                                                    <th>职位标识</th>
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
                                            <h3 class="box-title">系统设置</h3>
                                        </div>
                                        <!-- /.box-header -->
                                        <div class="box-body">
                                            <form class="form-horizontal">
                                                <div class="form-group">
                                                    <label for="inputName" class="col-sm-2 control-label">邮箱主机</label>
                                                    <div class="col-sm-10">
                                                        <input type="text" class="form-control" id="inputName"
                                                               placeholder="邮箱主机地址">
                                                    </div>
                                                </div>
                                                <div class="form-group">
                                                    <label for="inputName" class="col-sm-2 control-label">邮箱帐号</label>
                                                    <div class="col-sm-10">
                                                        <input type="text" class="form-control" id="inputName"
                                                               placeholder="邮箱帐号">
                                                    </div>
                                                </div>
                                                <div class="form-group">
                                                    <label for="inputName" class="col-sm-2 control-label">邮箱密码</label>
                                                    <div class="col-sm-10">
                                                        <input type="text" class="form-control" id="inputName"
                                                               placeholder="邮箱密码">
                                                    </div>
                                                </div>
                                                
                                                <div class="form-group">
                                                    <div class="col-sm-offset-2 col-sm-10">
                                                        <button type="submit" class="btn btn-danger">Submit</button>
                                                    </div>
                                                </div>
                                            </form>
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
    });
</script>
</body>
</html>

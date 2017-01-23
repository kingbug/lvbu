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
<span id="checkedenvs" style="display:none;">
{{if .user.Permission}}
	{{range $v := Getenvidlist .user.Permission}}
		<label>
	        <input type="checkbox"  class="minimal" name="copypermission" value="{{$v}}" checked="checked">{{$v}}
	    </label>
	{{end}}
{{end}}

	
</span>
<div class="wrapper">
    {{template "common/headertitle.tpl" .}}
    <!-- Left side column. contains the logo and sidebar -->
    {{template "common/sidebar.tpl" .}}
    <!-- Content Wrapper. Contains page content -->
    <div class="content-wrapper">
        <!-- Content Header (Page header) -->
        <section class="content-header">
            <h1>添加用户               
            </h1>
            <ol class="breadcrumb">
                <li><a href="/login"><i class="fa fa-dashboard"></i>主页</a></li>                
                <li class="active">用户添加</li>
            </ol>
        </section>

        <!-- Main content -->
        <section class="content">
            <div class="row">
                <div class="col-md-12">
                    <div class="box">                       
                        <div class="box-body">
                            <form class="form-horizontal" action="" method="post">
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">用户名</label>
									{{if .user.UserName}}
										<div class="col-sm-10">
	                                        <input type="text" class="form-control" id="inputName" name="username" placeholder="" value="{{.user.UserName}}">
	                                    </div>
									{{else}}
										<div class="col-sm-10">
	                                        <input type="text" class="form-control" id="inputName" name="username" placeholder="用于登录系统">
	                                    </div>
									{{end}}
									{{if .usernameerr}}
										<lable style="margin-left:18%;">
												<small class="label label-danger"> {{.usernameerr}}</small>
										</label>
									{{end}}
                                    
                                </div>
								<div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">密码</label>
									{{if .user.Passwd}}
										<div class="col-sm-10">
	                                        <input type="password" class="form-control" id="inputName" name="passwd" placeholder="" value="{{.user.Passwd}}">
	                                    </div>
									{{else}}
										<div class="col-sm-10">
	                                        <input type="password" class="form-control" id="inputName" name="passwd" placeholder="密码">
	                                    </div>
									{{end}}
									{{if .passwderr}}
										<lable style="margin-left:18%;">
												<small class="label label-danger"> {{.passwderr}}</small>
										</label>
									{{end}}
 
                                </div>
								<div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">姓名</label>
									
									{{if .user.Nick}}
										<div class="col-sm-10">
	                                        <input type="text" class="form-control" id="inputName" name="nick" placeholder="" value="{{.user.Nick}}">
	                                    </div>
									{{else}}
	                                    <div class="col-sm-10">
	                                        <input type="text" class="form-control" id="inputName" name="nick" placeholder="姓名">
	                                    </div>
									{{end}}
									{{if .nickerr}}
										<lable style="margin-left:18%;">
												<small class="label label-danger"> {{.nickerr}}</small>
										</label>
									{{end}}
									
                                </div>
								<div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">性别</label>
                                    <div class="col-sm-10">
										{{ if .user.Sex}}
											{{if eq .usersex 1}}
												<label>
													<input type="radio" name="sex" class="minimal" value="1" checked="checked">男
												</label>
												<label style="margin-left:100px;">
													<input type="radio" name="sex" class="minimal" value="0">女
												</label>
											{{else}}
												<label>
													<input type="radio" name="sex" class="minimal" value="1" >男
												</label>
												<label style="margin-left:100px;">
													<input type="radio" name="sex" class="minimal" value="0" checked="checked">女
												</label>
											{{end}}
											
										{{else}}
									
	                                      	<label>
												<input type="radio" name="sex" class="minimal" value="1" checked="checked">男
											</label>
										
											<label style="margin-left:100px;">
												<input type="radio" name="sex" class="minimal" value="0">女
											</label>
										{{end}}
										{{if .sexerr}}
											<lable style="margin-left:18%;">
												<small class="label label-danger"> {{.sexerr}}</small>
											</label>
										{{end}}

                                    </div>
                                </div>
								<div class="form-group">								
                                    <label for="inputName" class="col-sm-2 control-label">电话</label>
									{{if .user.Phone}}
										<div class="col-sm-10">
	                                        <input type="text" class="form-control" id="inputName" name="phone" placeholder="" value="{{.user.Phone}}">
	                                    </div>
									{{else}}
										<div class="col-sm-10">
	                                        <input type="text" class="form-control" id="inputName" name="phone" placeholder="电话">
	                                    </div>
									{{end}}
									{{if .phoneerr}}
										<lable style="margin-left:18%;">
												<small class="label label-danger"> {{.phoneerr}}</small>
										</label>
									{{end}}									
                                   
                                </div>
								<div class="form-group">								
                                    <label for="inputName" class="col-sm-2 control-label">邮箱</label>	
									{{if .user.Email}}
										<div class="col-sm-10">
	                                        <input type="email" class="form-control" id="inputName" name="email" placeholder="" value="{{.user.Email}}">
	                                    </div>
									{{else}}
										<div class="col-sm-10">
	                                        <input type="email" class="form-control" id="inputName" name="email" placeholder="邮箱">
	                                    </div>
									{{end}}
									{{if .emailerr}}
										<lable style="margin-left:18%;">
												<small class="label label-danger"> {{.emailerr}}</small>
										</label>
									{{end}}											
                                   
                                </div>
								<div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">职位</label>
                                    <div class="col-sm-10">
                                        <select class="form-control select2" name="position" style="width: 100%;">
											{{if .user.Position}}
												<option value="0">选择职能↓↓↓</option>
											{{else}}
												<option selected="selected" value="0">选择职能↓↓↓</option>
											{{end}}
											{{ range Getposition }}
												{{if $.user.Position}}
													{{if eq .Id $.user.Position.Id}}
														<option selected="selected" value="{{.Id}}">{{.Name}}</option>
													{{else}}
														<option value="{{.Id}}">{{.Name}}</option>
													{{end}}
												{{else}}
													<option value="{{.Id}}">{{.Name}}</option>
												{{end}}
											{{end}}
                                        </select>
                                    </div>
									{{if .positionerr}}
										<lable style="margin-left:18%;">
												<small class="label label-danger"> {{.positionerr}}</small>
										</label>
									{{end}}
                                </div> 
								<div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">环境权限</label>
                                    <div class="col-sm-10">  
										{{ range Getenv}}
											<label>
                                                <input type="checkbox"  class="minimal" name="permission" value="{{.Id}}">{{.Name}}
                                            </label>
										{{end}}  
										{{if .enverr}}
											<lable style="margin-left:10%;">
													<small class="label label-danger"> {{.enverr}}</small>
											</label>
										{{end}}		                                   
                                           <!-- <label>
                                                <input type="checkbox"  class="minimal"  checked="checked">开发环境
                                            </label>
                                            <label>
                                                <input type="checkbox"  class="minimal">测试环境
                                            </label>  
											<label>
                                                <input type="checkbox"  class="minimal">运维环境
                                            </label>     -->                                  
                                    </div>
                                </div>           
                                <div class="form-group">
                                    <div class="col-sm-offset-2 col-sm-10">
                                        <button type="submit" class="btn btn-danger">Submit</button>
										{{if .message}}
											<div class="form-group has-error">
							                  <label class="control-label" for="inputError"><i class="fa fa-times-circle-o"></i> {{.message}}</label>
							                 
							                </div>
										{{end}}
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

<script>
$("#checkedenvs input[name='copypermission']:checked").each(function(){
		$("[name='permission'][value=" + $(this).val() + "]").attr("checked",'true');
	})
</script>
</body>
</html>
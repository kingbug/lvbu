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
					
						{{range $v := Getenvidlist .user.Permission}}
							<label>
                                <input type="checkbox"  class="minimal" name="copypermission" value="{{$v}}" checked="checked">{{$v}}
                            </label>
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
            <h1>编辑用户               
            </h1>
            <ol class="breadcrumb">
                <li><a href="/login"><i class="fa fa-dashboard"></i>主页</a></li>                
                <li class="active">用户编辑</li>
            </ol>
        </section>

        <!-- Main content -->
        <section class="content">
            <div class="row">
                <div class="col-md-12">
                    <div class="box">                       
                        <div class="box-body">
                            <form class="form-horizontal" action="" method="post" onsubmit="return toVaild()">
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">用户名</label>
                                    <div class="col-sm-10">
										
                                        <input type="text" class="form-control" id="inputName" placeholder="用于登录系统" name="username" value="{{.user.UserName}}">
                                    </div>
									<lable id="usernameerr" style="margin-left:18%; display:none;">
												<small id="usernameerrtext" class="label label-danger"></small>
									</label>
                                </div>
								<div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">密码</label>
                                    <div class="col-sm-10">
                                        <input type="password" class="form-control" id="inputName" placeholder="密码" name="passwd" value="{{.user.Passwd}}">
                                    </div>
									<lable id="passwderr" style="margin-left:18%;display:none;">
												<small id="passwderrtext" class="label label-danger"></small>
									</label>
                                </div>
								<div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">姓名</label>
                                    <div class="col-sm-10">
                                        <input type="text" class="form-control" id="inputName" placeholder="姓名" name="nick" value="{{.user.Nick}}">
                                    </div>
									<lable id="nickerr" style="margin-left:18%; display:none;">
												<small id="nickerrtext" class="label label-danger"></small>
									</label>
                                </div>
								<div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">性别</label>
                                    <div class="col-sm-10">
										{{ if .user.Sex}}
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
									<lable id="sexerr" style="margin-left:18%; display:none;">
												<small id="sexerrtext" class="label label-danger"></small>
									</label>
										

                                    </div>
                                </div>
								<div class="form-group">								
                                    <label for="inputName" class="col-sm-2 control-label">电话</label>									
                                    <div class="col-sm-10">
                                        <input type="text" class="form-control" id="inputName" placeholder="电话" name="phone" value="{{.user.Phone}}">
                                    </div>
									<lable id="phoneerr" style="margin-left:18%; display:none;">
												<small id="phoneerrtext" class="label label-danger"></small>
									</label>
                                </div>
								<div class="form-group">								
                                    <label for="inputName" class="col-sm-2 control-label">邮箱</label>									
                                    <div class="col-sm-10">
                                        <input type="text" class="form-control" id="inputName" placeholder="邮箱" name="email" value="{{.user.Email}}">
                                    </div>
									<lable id="emailerr" style="margin-left:18%; display:none;">
												<small id="emailerrtext" class="label label-danger"></small>
									</label>
                                </div>
								<div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">职位</label>
                                    <div class="col-sm-10">
                                        <select class="form-control select2" name="position" style="width: 100%;">
											{{range Getposition}}
												{{if $.user.Position.Id}}
													{{if eq $.user.Position.Id .Id}}
														<option value="{{.Id}}" selected="selected">{{.Name}}</option>
													{{else}}
														<option value="{{.Id}}">{{.Name}}</option>
													{{end}}
												
												{{end}}
												
											{{end}}
                                            
                                        </select>
                                    </div>
									<lable id="positionerr" style="margin-left:18%; display:none;">
												<small id="positionerrtext" class="label label-danger"></small>
									</label>
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
												<lable style="margin-left:18%;">
														<small class="label label-danger"> {{.enverr}}</small>
												</label>
											{{end}}		                                       
                                                                          
                                    </div>
									<lable id="permissionerr" style="margin-left:18%; display:none;">
												<small id="permissionerrtext" class="label label-danger"></small>
									</label>
                                </div> 
								<div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">状态</label>
                                    <div class="col-sm-10">  
									        {{if .user.Status}}     
												<label>
												    <i class="fa fa-lock text-blue">锁定</i>
	                                                <input type="radio" name="status" class="minimal"  value="1" checked="checked">
	                                            </label>
	                                            <label>
												   <i class="fa fa-key text-red">解锁</i>
	                                                <input type="radio" name="status" class="minimal"  value="0">
	                                            </label> 
											{{else}}      	   
	                                            
												<label>
												    <i class="fa fa-lock text-blue">锁定</i>
	                                                <input type="radio" name="status" class="minimal" value="1">
	                                            </label>
												<label>
												   <i class="fa fa-key text-red">解锁</i>
	                                               <input type="radio" name="status" class="minimal" value="0"  checked="checked">
	                                            </label>  
											{{end}} 										                                  
                                    </div>
									<lable id="statuserr" style="margin-left:18%; display:none;">
												<small id="statuserrtext" class="label label-danger"></small>
									</label>
                                </div>            
                                <div class="form-group">
                                    <div class="col-sm-offset-2 col-sm-10">
                                        <button type="submit" class="btn btn-danger">Submit</button>
                                    </div>
									{{if .message}}
										<div class="form-group has-error">
						                  <label class="control-label" for="inputError"><i class="fa fa-times-circle-o"></i> {{.message}}</label>
						                 
						                </div>
									{{end}}
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
	
	$("#checkedenvs input[name='copypermission']:checked").each(function(){
		$("[name='permission'][value=" + $(this).val() + "]").attr("checked",'true');
	})

	function toVaild(){
		var con = 0;
		// from 留空验证  start
		if ($("input[name='username']").val() == ""){
			$("#usernameerr").show();
			$("#usernameerrtext").html("用户名(登录名)不能为空");
			con = con +1;
		}else{
			$("#usernameerr").hide();
		}
		if ($("input[name='passwd']").val() == ""){
			$("#passwderr").show();
			$("#passwderrtext").html("密码不能为空");
			con = con +1;
		}else{
			$("#passwderr").hide();
		}
		if ($("input[name='nick']").val() == ""){
			$("#nickerr").show();
			$("#nickerrtext").html("姓名不能为空");
			con = con +1;
		}else{
			$("#nickerr").hide();
		}
		if ($("input[name='phone']").val() == ""){
			$("#phoneerr").show();
			$("#phoneerrtext").html("电话不能为空");
			con = con +1;
		}else{
			$("#phoneerr").hide();
		}
		if ($("input[name='email']").val() == ""){
			$("#emailerr").show();
			$("#emailerrtext").html("邮箱不能为空");
			con = con +1;
		}else{
			$("#emailerr").hide();
		}
		
		// from 留空验证 end
		if (con != 0){
			return false;
			con = 0;
		}else {
			return true;
		}
		
	}
	
</script>
</body>
</html>
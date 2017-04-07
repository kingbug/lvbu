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
                            <form class="form-horizontal" method="post" onSubmit="return toVaild();">
                                <div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">项目名称</label>
                                    <div class="col-sm-10">
                                        <input type="text" class="form-control" id="inputName" name="name" placeholder="中文，用于登记项目名称" value="{{.pro.Name}}">
                                    </div>
									{{if .nameerr}}
										<lable style="margin-left:18%;" class="nameerr">
												<small class="label label-danger"> {{.nameerr}}</small>
										</label>
									{{else}}
										<lable style="margin-left:18%; display:none;" class="nameerr">
												<small class="label label-danger"></small>
										</label>
									{{end}}
                                </div>
								<div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">项目标识</label>
                                    <div class="col-sm-10">
                                        <input type="text" class="form-control" id="inputName" name="sign" placeholder="字母+数字，用于项目唯一性标识" value="{{.pro.Sign}}">
                                    </div>
									{{if .signerr}}
										<lable style="margin-left:18%;" class="signerr">
												<small class="label label-danger"> {{.signerr}}</small>
										</label>
									{{else}}
										<lable style="margin-left:18%; display:none;" class="signerr">
												<small class="label label-danger"></small>
										</label>
									{{end}}
                                </div>
								<div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">代码标识</label>
                                    <div class="col-sm-10">
                                        <select class="form-control select2" name="compile" style="width: 20%;">
											{{range $k, $v := Compilever}}
												{{if $.pro.Compile}}
													{{if eq $.pro.Compile $k}}
														<option value="{{$k}}"  selected = "selected">{{$k}}</option>
													{{else}}
														<option value="{{$k}}">{{$k}}</option>
													{{end}}
												{{else}}
													<option value="{{$k}}">{{$k}}</option>
												{{end}}
											{{end}}
                                            
                                        </select>
										<label style="font-weight: 700;">语言版本:</label>
										<select class="form-control select2" name="compilever" style="width: 30%;">
											<option>先选择代码标识</option>
                                        </select>
                                    </div>
									{{if .compileerr}}
										<lable style="margin-left:18%;" class="compileerr">
												<small class="label label-danger"> {{.compileerr}}</small>
										</label>
									{{else}}
										<lable style="margin-left:18%; display:none;" class="compileerr">
												<small class="label label-danger"></small>
										</label>
									{{end}}
                                </div>
								<div class="form-group">
                                    <label for="inputName" class="col-sm-2 control-label">仓库地址</label>
                                    <div class="col-sm-10">
                                        <input type="text" class="form-control" id="inputName" name="git" placeholder="仅支持git">
                                    </div>
									{{if .giterr}}
										<lable style="margin-left:18%;" class="giterr">
												<small class="label label-danger"> {{.giterr}}</small>
										</label>
									{{else}}
										<lable style="margin-left:18%; display:none;" class="giterr">
												<small class="label label-danger"></small>
										</label>
									{{end}}
                                </div>
								
								<div class="form-group addsite">
									<label for="inputName" class="col-sm-2 control-label"></label>
									<div class="col-sm-10"  style="margin-top:-20px;">
                                        <a class="addsite btn" style="height:14px;">添加多个仓库</a>
                                    </div>
								</div>
								<div class="form-group">								
                                    <label for="inputName" class="col-sm-2 control-label">帐号密码</label>									
                                    <div class="col-sm-10">
                                        <input type="text" class="form-control" id="inputName" name="gituser" placeholder="git帐号密码，冒号分割。如admin:admin" value="{{.pro.Gituser}}">
                                    </div>
									{{if .gitusererr}}
										<lable style="margin-left:18%;" class="gitusererr">
												<small class="label label-danger"> {{.gitusererr}}</small>
										</label>
									{{else}}
										<lable style="margin-left:18%; display:none" class="gitusererr">
												<small class="label label-danger"></small>
										</label>
									{{end}}
                                </div>
								<div class="form-group">								
                                    <label for="inputName" class="col-sm-2 control-label">项目DNS</label>									
                                    <div class="col-sm-10">
										<input type="text" class="form-control" id="inputName" name="dns" placeholder="只支持一个" value="{{.pro.Dns}}">
                                    </div>
									{{if .dnserr}}
										<lable style="margin-left:18%;" class="dnserr">
												<small class="label label-danger"> {{.dnserr}}</small>
										</label>
									{{else}}
										<lable style="margin-left:18%; display:none;" class="dnserr">
												<small class="label label-danger"></small>
										</label>
									{{end}}
                                </div>
								<div class="form-group">								
                                    <label for="inputName" class="col-sm-2 control-label">共享文件</label>
									<div class="col-sm-10">
                                        <textarea class="form-control" id="inputExperience" name="sharedpath"
                                                  placeholder="开始&quot;/&quot;表示容器根目录开始,否则项目根目录。默认会在所属主机&quot;/lvbu/项目标识/sharedpath&quot;创建映射目录">{{.pro.Sharedpath}}</textarea>
                                    </div>
									{{if .sharedpatherr}}
										<lable style="margin-left:18%;">
												<small class="label label-danger"> {{.sharedpatherr}}</small>
										</label>
									{{else}}
										<lable style="margin-left:18%; display:none" class="sharedpatherr">
												<small class="label label-danger"></small>
										</label>
									{{end}}
                                </div>
								
								<div class="form-group">
                                    <label for="inputExperience" class="col-sm-2 control-label">忽略文件</label>

                                    <div class="col-sm-10">
                                        <textarea class="form-control" id="inputExperience" name="insfile"
                                                  placeholder="git需要忽略的文件">{{.pro.Insfile}}</textarea>
                                    </div>
									{{if .insfileerr}}
										<lable style="margin-left:18%;">
												<small class="label label-danger"> {{.insfileerr}}</small>
										</label>
									{{else}}
										<lable style="margin-left:18%; display:none" class="insfileerr">
												<small class="label label-danger"></small>
										</label>
									{{end}}
                                </div>
								      
                                <div class="form-group">
                                    <div class="col-sm-offset-2 col-sm-10">
                                        <input type="submit" class="btn btn-danger" value="Submit"></input>
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
        //$(".select2").select2();
    });
	
	var compilever = {{Compilever}} //编译环境map
	var compilever_post ={{$.pro.Compilever}}
	function toVaild() {
		
		var compile = $("form select[name='compile']").val();
		var name = $("form input[name='name']").val();
		var sign = $("form input[name='sign']").val();
		var gititem	 = $("form input[name='git']");
		var git = [];
		var tmp_git = "";
		gititem.each(function(){
			if ($(this).val() != ""){
				tmp_git = $(this).val();
			}
		});
		
		var gituser = $("form input[name='gituser']").val();
		var contr = true
		if ( name == "") {
			contr = false;
			$(".nameerr small").text("项目名称不能为空");
			$(".nameerr").show();
		}else {
			$(".nameerr").hide();
		}
		if (sign == "") {
			contr = false;
			$(".signerr small").text("项目标识不能为空");
			$(".signerr").show(); 
		}else {$(".signerr").hide();}
		if (compile == "") {
			contr = false;
			$(".compileerr small").text("代码标识不能为空");
			$(".compileerr").show(); 
		}else {$(".compileerr").hide();}
		if (tmp_git == "") {
			contr = false;
			$(".giterr small").text("仓库地址不能为空");
			$(".giterr").show(); 
		} else {$(".giterr").hide();}
		$(".dnserr").hide();
		return contr;
	}
	
	$("select[name='compile']").change(function(){
		selectlanguage = $(this);
		selectver = selectlanguage.siblings("select");
		selectver.empty();//清空 
		compile_key = selectlanguage.find("option:selected").text();
		tagjson = compilever[compile_key];
		$.each(tagjson, function(key,value){
			selectver.append("<option value='" + value + "'>" + key +"</option>");
		});
		return;
	});
	
	
	//提交错误时，初始化编译版本选择 
	selectlanguage = $("form select[name='compile']");
	selectver = selectlanguage.siblings("select");
	selectver.empty();//清空 
	compile_key = selectlanguage.find("option:selected").text();
	tagjson = compilever[compile_key];
	$.each(tagjson, function(key,value){
		selectver.append("<option value='" + value + "'>" + key +"</option>");
	});
	//提交错误 end
	console.log("compilever_post:" + compilever_post);
	if (compilever_post != ""){
		console.log("不为空");
		selectver.val(compilever_post);
	}
	//提交错误时，页面返回后初始化之前填写的git
	gitstr = {{$.pro.Git}}
	gitlist = gitstr.split(",");
	console.log("length:" + gitlist.length);
	
	$.each(gitlist, function(k,v){
		if ( k== 0) {
			$("form input[name='git']").val(v);
			return true;
		} 
		if ( v== "" ){
			return true;
		}
		$("div.addsite").before("<div class=\"form-group\">" +
									"<label for=\"inputName\" class=\"col-sm-2 control-label\">&nbsp;&nbsp;&nbsp;&nbsp;</label>"+
									"<div class=\"col-sm-10\">" +
									"<input type=\"text\" class=\"form-control\" id=\"inputName\" name=\"git\" placeholder=\"仅支持git\" value=\"" +v+ "\">"+
									"</div></div>");
	});
	
	
	
	$("a.addsite").click(function(){
		$("div.addsite").before("<div class=\"form-group\">" +
									"<label for=\"inputName\" class=\"col-sm-2 control-label\">&nbsp;&nbsp;&nbsp;&nbsp;</label>"+
									"<div class=\"col-sm-10\">" +
									"<input type=\"text\" class=\"form-control\" id=\"inputName\" name=\"git\" placeholder=\"仅支持git\">"+
									"</div></div>");
		
	});
</script>
</body>
</html>
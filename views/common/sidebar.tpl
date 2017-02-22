<aside class="main-sidebar">
    <!-- sidebar: style can be found in sidebar.less -->
    <section class="sidebar">
        <!-- Sidebar user panel -->
        <div class="user-panel">
            <div class="pull-left image">
                <img src="/static/img/{{(Getuser .uid).Avatar}}" class="img-circle" alt="User Image">
            </div>
            <div class="pull-left info">
                <p>{{(Getuser .uid).Nick}}</p>
				<i class="fa fa-circle text-success"></i>{{Getposname  (((Getuser .uid).Position).Id)}}
            </div>
        </div>
        <!-- sidebar menu: : style can be found in sidebar.less -->
        <ul class="sidebar-menu">            
            <li class="treeview">
                <a href="/login">
                    <i class="fa fa-dashboard"></i> <span>主页</span>
                </a>
            </li>
			<li class="header">部署相关功能</li>
			{{if(Isperitem "pros" .uid)}}
			<li class="treeview">
                <a href="#">
                    <i class="fa fa-th text-yellow"></i>
                    <span>项目管理</span>					
                    <span class="pull-right-container">	
					<span class="label label-primary pull-left">4</span>				
                  <i class="fa fa-angle-left pull-right"></i>
                </span>
                </a>
                <ul class="treeview-menu">
                    <li><a href="/prolist"><i class="fa fa-circle-o"></i>项目列表</a></li>
					{{if(Isperitem "proa" .uid)}}
                    <li><a href="/proadd"><i class="fa fa-circle-o"></i>添加项目</a></li> 
					{{end}}                   
                </ul>
            </li>
			{{end}}
			{{if(Isperitem "macs" .uid)}}           
            <li class="treeview">
                <a href="#">
                    <i class="fa fa-laptop text-yellow"></i>
                    <span>主机管理</span>
                    <span class="pull-right-container">
                  <i class="fa fa-angle-left pull-right"></i>
                </span>
                </a>
                <ul class="treeview-menu">
                    <li><a href="/maclist"><i class="fa fa-circle-o"></i>主机列表</a></li>
					{{if(Isperitem "maca" .uid)}}
                    <li><a href="/macadd"><i class="fa fa-circle-o"></i>添加主机</a></li>
					{{end}}
                </ul>
            </li>
			{{end}}
			<!--
			{{if(Isperitem "cons" .uid)}}
            <li class="treeview">
                <a href="#">
                    <i class="fa fa-edit text-yellow"></i> <span>配置管理</span>
                    <span class="pull-right-container">
                  <i class="fa fa-angle-left pull-right"></i>
                </span>
                </a>
                <ul class="treeview-menu">
					{{if(Isperitem "cons" .uid)}}
                    <li><a href="/conlist"><i class="fa fa-circle-o"></i>配置列表</a></li>
					{{end}}
                    <li><a href="#"><i class="fa fa-circle-o"></i>配置查询</a></li>
                </ul>
            </li>
			{{end}}
			-->
			
			{{if(Isperitem "mirs" .uid)}}
            <li class="treeview">
                <a href="#">
                    <i class="fa fa-files-o text-yellow"></i>
                    <span>镜像管理</span>
                    <span class="pull-right-container">
              <span class="label label-primary pull-right"></span>
            </span>
                </a>
                <ul class="treeview-menu">
                    <li><a href="/mirrlist"><i class="fa fa-circle-o"></i>镜像列表</a></li>
					{{if(Isperitem "mira" .uid)}}
                    <li><a href="/mirradd"><i class="fa fa-circle-o"></i>新增镜像</a></li>                    
                    {{end}}
					</li>
                </ul>
				
            </li>
			{{end}}            
            <li class="header">产品管理相关功能(暂不可用)</li>
			<li class="treeview">
                <a href="#">
                    <i class="fa fa-table text-aqua"></i> <span>产品管理</span>
                    <span class="pull-right-container">
                  <i class="fa fa-angle-left pull-right"></i>
                </span>
                </a>
                <ul class="treeview-menu">
                    <li><a href="#"><i class="fa fa-circle-o"></i> Simple tables</a></li>
                    <li><a href="#"><i class="fa fa-circle-o"></i> Data tables</a></li>
                </ul>
            </li>
            <li><a href="#"><i class="fa fa-book text-aqua"></i> <span>版本控制</span></a>
            </li>
            <li class="header">系统相关功能{{Isperitem "inds" .uid}}</li>
			{{if(Isperitem "sets" .uid)}}
            <li class="treeview">
                <a href="#">
                    <i class="fa fa-cog text-green"></i> <span>系统设置</span>
                    <span class="pull-right-container">
                  <i class="fa fa-angle-left pull-right"></i>
                </span>
                </a>
                <ul class="treeview-menu">
				 	<li><a href="/usermanager"><i class="fa fa-circle-o"></i>用户管理</a></li>
                    <li><a href="/env"><i class="fa fa-circle-o"></i>环境设置</a></li>
                    <li><a href="/reclist"><i class="fa fa-circle-o"></i>记录查询</a></li>
                </ul>
            </li> 			
			{{end}}
			<li><a href="/profile"><i class="fa fa-gg text-blue"></i> <span>个人设置</span></a></li>          
            <li><a href="/about"><i class="fa fa-gg text-purple"></i> <span>关于系统</span></a></li>
        </ul>
    </section>
    <!-- /.sidebar -->
</aside>
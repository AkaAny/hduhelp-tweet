//展开收起
var bottom = "65px";
var right = "122px";
var bottomBf = "35%";
var rightBf = "184px";
var setting = {
		imageWidth:1680,
		imageHeight:1050
	};
var windowHeight = $(window).height();
var windowWidth = $(window).width();
init = function(){
	var windowHeight = $(window).height();
	var windowWidth = $(window).width();
	$(".login_conatiner").height(windowHeight).width(windowWidth);
	$("#container_bg").height(windowHeight).width(windowWidth);
	$("#login_right_box").height(windowHeight);
	var imgW = setting.imageWidth;
	var imgH = setting.imageHeight;
	var ratio = imgH / imgW; //图片的高宽比

	imgW = windowWidth; //图片的宽度等于窗口宽度
	imgH = Math.round(windowWidth * ratio); //图片高度等于图片宽度 乘以 高宽比

	if(imgH < windowHeight){ //但如果图片高度小于窗口高度的话
		imgH = windowHeight; //让图片高度等于窗口高度
		imgW = Math.round(imgH / ratio); //图片宽度等于图片高度 除以 高宽比
	}
	$(".login_img_01").width(imgW).height(imgH);  //设置图片高度和宽度
	
};
openBtn = function(){
	
	if(1366 <= windowWidth && windowWidth <= 1440 ){
		rightBf = "10%";
	}
	if(1024 <= windowWidth && windowWidth <= 1250 ){
		right = "20px";
	}
	if(windowHeight <= 900 ){
		bottomBf = (windowHeight - 395) * 0.5;
	}
	if(windowHeight <= 680 ){
		bottomBf = (windowHeight - 395) * 0.3;
	}
};
var lqrcode;
$(function(){
	$(window).resize(function(){
		init();
		openBtn();
	});
	//加载账号密码登录
	password_login(true);
	
	$(".toggle_btn").hover(function(){
		if($(this).attr("data-shrink")=="0"){
			$(this).html("收起");

		}else{
			$(this).html("展开");

		}
	},function(){
		if($(this).attr("data-shrink")=="0"){
			$(this).html("<span class='minus'></span>");

		}else{
			$(this).html("<span class='plus'></span>");

		}
	});
	$(".toggle_btn").click(function(){
		if($(this).attr("data-shrink")=="0") {
			$(this).parent().animate({bottom: bottom, right: right, height: "41px"}, 400);
			$(this).parent().addClass("box-open");
			$(this).html("<span class='plus'></span>");
			$(this).attr("data-shrink", "1");
		}else {
			if(windowHeight <= 900 ){
				bottomBf = (windowHeight - 395) * 0.5;
			}
			if(windowHeight <= 680 ){
				bottomBf = (windowHeight - 395) * 0.3;
			}
			$(this).parent().animate({bottom: bottomBf, right: rightBf, height: "395px"}, 400);
			$(this).parent().removeClass("box-open");
			$(this).html("<span class='minus'></span>");
			$(this).attr("data-shrink", "0");
		}
	});
});
//加载账号登录功能
function password_login(firstclick){
	var passwordhtml = document.getElementById("password_template").innerHTML;	
	$("#template_container").html(passwordhtml);
	init();
	openBtn();
	if(lqrcode!=null){
		lqrcode.clear();
	}
	//每次切换都先让登录框最大化
	if(!firstclick){
		$(".toggle_btn").parent().animate({bottom: bottomBf, right: rightBf, height: "395px"}, 400);
		$(".toggle_btn").parent().removeClass("box-open");
		$(".toggle_btn").html("<span class='minus'></span>");
		$(".toggle_btn").attr("data-shrink", "0");		
	}
	
	$("#index_login_btn").click(function(){
		login();
	}); 
	
	//点击账号登陆
	$("#password_login").click(function(){
		$("#password_login").addClass("active");
		$("#qrcode_login").removeClass("active");
		$("#login_content").html(passwordhtml);
	});
	//点击扫码登陆
	$("#qrcode_login").click(function(){
		$("#password_login").removeClass("active");
		$("#qrcode_login").addClass("active");
		$("#login_content").html(qrcodehtml);
	});
	
	//点击记住用户名
	$("#rememberName").change(function(){
		if($(this).is(":checked")){
			var $u = $("#un").val() ;
			if($.trim($u)==''){
				$("#errormsg").text("账号不能为空。").show();
				$(this).removeAttr("checked");
			}else{
				//不等于空，写cookie
				setCookie('hdu_cas_un' , $u , 365);
			}
		}else{
			//反选之后清空cookie
			clearCookie('hdu_cas_un');
		}
	});
	
	//用户名文本域keyup事件
	$("#un").keyup(function(e){
		if(e.which == 13) {
			login();
	    }
	}).keydown(function(e){
		$("#errormsg").hide();
	}).focus();
	
	//密码文本域keyup事件
	$("#pd").keyup(function(e){
		if(e.which == 13) {
			login();
	    }
	}).keydown(function(e){
		$("#errormsg").hide();
	});
	//如果有错误信息，则显示
	if($("#errormsghide").text()){
		$("#errormsg").text($("#errormsghide").text()).show();
	}
	//获取cookie值
	var cookie = getCookie('hdu_cas_un');
	if(cookie){
		$("#un").val(cookie);
		$("#rememberName").attr("checked","checked");
	}
	//重新获取验证码
	$("#a_changeCode").click(function(){
    	$("#codeImage").attr("src", "code?"+Math.random()) ;
    });
	
	//360
	$("#open_360").mouseover(function(){
		$("#open_360_img").show();
	}).mouseout(function(){
		$("#open_360_img").hide();
	});
	
	//如果有错误信息，则显示
	if($("#errormsghide").text()){
		$("#errormsg").text($("#errormsghide").text()).show();
	}
};
//加载扫码登录功能
function qrcode_login(){
	var qrcodehtml = document.getElementById("qrcode_template").innerHTML;
	$("#template_container").html(qrcodehtml);
	init();
	openBtn();
	//每次切换都先让登录框最大化
	$(".toggle_btn").parent().animate({bottom: bottomBf, right: rightBf, height: "395px"}, 400);
	$(".toggle_btn").parent().removeClass("box-open");
	$(".toggle_btn").html("<span class='minus'></span>");
	$(".toggle_btn").attr("data-shrink", "0");
	
	lqrcode = new loginQRCode("qrcode",153,153);
	lqrcode.generateLoginQRCode(function(result){
		window.location.href = result.redirect_url;
	});
}

function login(){
	var $u = $("#un") , $p=$("#pd");
	
	var u = $u.val().trim();
	if(u==""){
		$u.focus();
		$u.parent().addClass("login_error_border");
		return ;
	}
	
	var p = $p.val().trim();
	if(p==""){
		$p.focus();
		$p.parent().addClass("login_error_border");
		return ;
	}
	
	$u.attr("disabled","disabled");
	$p.attr("disabled","disabled");
	
	var lt = $("#lt").val();
	
	$("#ul").val(u.length);
	$("#pl").val(p.length);
	$("#rsa").val(strEnc(u+p+lt , '1' , '2' , '3'));
	
	$("#loginForm")[0].submit();
	
}

//图片的轮播功能
jQuery(function() {

	jQuery('#camera_wrap_4').camera({

		height : 'auto',//高度

		hover : false,//鼠标经过幻灯片时暂停(true, false)

		//imagePath: 图片的目录

		loader : 'none',//加载图标(pie, bar, none)

		//loaderColor: 加载图标颜色( '颜色值,例如:#eee' )

		//loaderBgColor: 加载图标背景颜色

		loaderOpacity : '8',//加载图标的透明度( '.8'默认值, 其他的可以写 0, .1, .2, .3, .4, .5, .6, .7, .8, .9, 1 )

		loaderPadding : '2',//加载图标的大小( 填数字,默认为2 )

		navigation : false,//左右箭头显示/隐藏(true, false)

		navigationHover : false,//鼠标经过时左右箭头显示/隐藏(true, false)

		pagination : false,//是否显示分页(true, false)

		playPause : false,//暂停按钮显示/隐藏(true, false)

		pauseOnClick : false,//鼠标点击后是否暂停(true, false)

		portrait : false,//显示幻灯片里所有图片的实际大小(true, false)

		thumbnails : false,//是否显示缩略图(true, false)

		time : 500,// 幻灯片播放时间( 填数字 )

		//transPeriod: 4000,//动画速度( 填数字 )

		imagePath : '../images/',
		thumbnails : false

	});

	//触发如何使用360极速模式图片
	$("#open_360").mouseover(function() {
		$("#open_360_img").show();
	}).mouseout(function() {
		$("#open_360_img").hide();
	});

});

//设置cookie
function setCookie(cname, cvalue, exdays) {
  var d = new Date();
  d.setTime(d.getTime() + (exdays*24*60*60*1000));
  var expires = "expires="+d.toUTCString();
  document.cookie = cname + "=" + cvalue + "; " + expires;
}

//获取cookie
function getCookie(cname) {
  var name = cname + "=";
  var ca = document.cookie.split(';');
  for(var i=0; i<ca.length; i++) {
      var c = ca[i];
      while (c.charAt(0)==' ') c = c.substring(1);
      if (c.indexOf(name) != -1) return c.substring(name.length, c.length);
  }
  return "";
}

//清除cookie  
function clearCookie(name) {  
  setCookie(name, "", -1);  
}  


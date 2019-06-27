$(function () {
	// 初始化 iCheck 选框
	if ($(".icheck input").length > 0) {
	    var iCheckTime
        if (typeof $(".icheck input").iCheck == 'undefined') {
            iCheckTime = setInterval(initICheck, 50)
        } else {
            initICheck()
        }
	    function initICheck() {
            $(".icheck input").iCheck({
                checkboxClass: 'icheckbox_square-blue',
                radioClass: 'iradio_square-blue',
                increaseArea: '20%' /* optional */
            });
            clearInterval(iCheckTime)
        }
    }
    // 初始化 Switch 组件
    if ($(".switch input").length > 0) {
        var switchTime
        if (typeof $(".switch input").bootstrapSwitch == 'undefined') {
            switchTime = setInterval(initSwitch, 50)
        } else {
            initSwitch()
        }
        function initSwitch() {
            $(".switch input").bootstrapSwitch()
            if (switchTime != undefined) {
                clearInterval(switchTime)
            }
        }
        // $(".switch input").bootstrapSwitch()
	}
});
/**
 * 切换 switch 的状态
 * item 对象或者对象id属性
 * state 状态值
 */
function switchState(item, state) {
    if (state != 0 && state != '' && state != null && state != undefined) {
        state = true
    } else {
        state = false
    }
    if (typeof item == 'string') {
        item = $(item)
    }
    if (typeof item.bootstrapSwitch == 'undefined') {
        setTimeout(function () {
            switchState(item, state)
        }, 50)
        return false
    }
    item.bootstrapSwitch('state', state)
	item.attr('checked', state)
}

/* 密码显示 */
$(".show-password").on('click', function(e){
	var show = $(this).data('show');
	console.info(show)
	if ($(show).attr('type')=='text') {
		$(show).attr('type', 'password')
	} else {
		$(show).attr('type', 'text')
	}
})
/**
 * 去除空格
 */
function Trim(str,is_global)   {    
	var result;    
	result = str.replace(/(^\s+)|(\s+$)/g,"");    
	if(is_global === 1 || is_global === true || is_global.toLowerCase()=="g")    {     
		result = result.replace(/\s/g,"");     
	}    
	return result; 
} 
/**
 * 获取url参数
 */
function getQueryString(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return unescape(r[2]); return null;
}


$(function(){
	$('body').on('click','.jump-url',function () {
		location.href = $(this).data('url');
	})
	// 编辑框的最大化样式调整
	$('body').on('click', '.layui-layer-max', function () {
		if($(this).attr('class').indexOf('layui-layer-maxmin') == -1) {
            $('.layui-layer.layui-layer-iframe').css({'max-width':'1000px', 'max-height': '600px'})
		} else {
            $('.layui-layer.layui-layer-iframe').css({'max-width': '100%', 'max-height': '100%'})
        }
    })
})


$(function(){
	$("body").on('change', '.change-show', changeShow)
})
/**
 * 值改变现实不同的信息
 */
function changeShow(self, setself=1) {
	if (setself==1)
		self = this;
	//  显示类		隐藏类		是否设置disabled
	var showclass = hideclass = disa = '';
	console.log($(self).is(':checkbox'))
	console.log($(self))
	if($(self).is('select')) {
		showclass = $(self).find('option:selected').data('show');
		hideclass = $(self).find('option:selected').data('hide');
		disa = $(self).find('option:selected').data('disa');
	} else if($(self).is(':checkbox')) {
		if ($(self).is(':checked')) {
			showclass = $(self).data('show');
			hideclass = $(self).data('hide');
		} else {
			hideclass = $(self).data('show');
		}
	} else {
		showclass = $(self).data('show');
		hideclass = $(self).data('hide');
		disa = $(self).data('disa');
	}
	console.log('show', showclass, 'hide', hideclass, 'disa', disa)
	if(disa==1) {
		$(hideclass).removeClass('show').addClass('hide').hide();
        // $(hideclass).contents().find('input, textarea, select').attr({'disabled':true});
        $(hideclass).contents().find('input, select').attr({'disabled':true});
		$(showclass).removeClass('hide').addClass('show').show();
		$(showclass).contents().find('input, textarea, select').attr({'disabled':false});
	} else {
		$(hideclass).removeClass('show').addClass('hide').hide();
		$(showclass).removeClass('hide').addClass('show').show();
	}
}

// 判断是否 pc 设备
function isPC() {
  var userAgentInfo = navigator.userAgent;
  var Agents = ["Android", "iPhone",
        "SymbianOS", "Windows Phone",
        "iPad", "iPod"];
  var flag = true;
  for (var v = 0; v < Agents.length; v++) {
    if (userAgentInfo.indexOf(Agents[v]) > 0) {
      flag = false;
      break;
    }
  }
  return flag;
}
// 判断是否 wap 设备
function isMobile() {
    var regex_match = /(nokia|iphone|android|motorola|^mot-|softbank|foma|docomo|kddi|up.browser|up.link|htc|dopod|blazer|netfront|helio|hosin|huawei|novarra|CoolPad|webos|techfaith|palmsource|blackberry|alcatel|amoi|ktouch|nexian|samsung|^sam-|s[cg]h|^lge|ericsson|philips|sagem|wellcom|bunjalloo|maui|symbian|smartphone|midp|wap|phone|windows ce|iemobile|^spice|^bird|^zte-|longcos|pantech|gionee|^sie-|portalmmm|jigs browser|hiptop|^benq|haier|^lct|operas*mobi|opera*mini|320x320|240x320|176x220)/i;
    var u = navigator.userAgent;
    if (null == u) {
        return true;
    }
    var result = regex_match.exec(u);
    if (null == result) {
        return false
    } else {
        return true
    }
}

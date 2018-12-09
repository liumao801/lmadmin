/*!*********************************************
 * Copyright (C) Corporation. All rights reserved.
 *
 * Author      :  liumao801
 * Email        : liumao801@gmail.com
 * Create Date  : 2018-12-08
 * Description  : 自定义的简化ajax请求扩展,依赖layer插件(一款非常优秀的弹出层插件),默认datatype='json',async=true
 * Version      : V1.0.2
 *
 *
***********************************************
调用示例
$.lmpost('/home/index', {'name':'lht'});
$.lmpost('/home/index', {'name':'lht'}, function (re) {
    alert(re.code);
},false,'json');
//传入参数
参数1：请求的地址
参数2：提交的值
参数3：成功时的回调函数
参数4：async的值，默认true
参数5：dataType同ajax里的dataType,默认'josn'
*/
(function ($) {
    $.extend({
        lmpost: function (url, data, success, async, dataType) {
            if (typeof (data) === 'undefined' || data === null) data = {};
            if (typeof (async) === 'undefined' || async === null) async = true;
            if (typeof (dataType) === 'undefined' || dataType === null) dataType = 'json';
            var win = window;
            if(parent){
                win = parent;
            }
            $.ajax({
                url: url,
                data: data,
                type: 'post',
                async: async,
                dataType: dataType,
                beforeSend: function (XHR) {
                    win.layer.load();
                },
                complete: function (XHR, TS) {

                },
                success: function (data) {
                    win.layer.closeAll('loading');
                    if (success) {
                        success(data);
                    }
                },
                error: function (XHR, msg, e) {
                    win.layer.closeAll('loading');
                    if (typeof (XHR.responseText) !== 'undefined') {
                        if (XHR.responseText.indexOf('HttpRequestValidationException') > -1) {
                            win.layer.alert("请求失败：" + '您输入的内容里有潜在危险的字符，例如：“&#” 等', { icon: 2, title: '错误' });
                        } else {
                            win.layer.alert("请求失败：" + XHR.responseText, { icon: 2, title: '错误' });
                        }
                    }
                    else {
                        win.layer.alert("请求失败", { icon: 2, title: '错误' });
                    }
                }
            });
        },
        lmget: function (url, data, success, async, dataType) {
            if (typeof (data) === 'undefined') data = {};
            if (typeof (async) === 'undefined' || async === null) async = true;
            if (typeof (dataType) === 'undefined' || dataType === null) dataType = 'json';
            var win = window;
            if(parent){
                win = parent;
            }
            $.ajax({
                url: url,
                data: data,
                type: 'get',
                async: async,
                dataType: dataType,
                beforeSend: function (XHR) {
                    win.layer.load();
                },
                complete: function (XHR, TS) {
                },
                success: function (data) {
                    win.layer.closeAll('loading');
                    if (success) {
                        success(data);
                    }
                },
                error: function (XHR, msg, e) {
                    win.layer.closeAll('loading');
                    if (typeof (XHR.responseText) !== 'undefined') {
                        if (XHR.responseText.indexOf('HttpRequestValidationException') > -1) {
                            win.layer.alert("请求失败：" + '您输入的内容里有潜在危险的字符，例如：“&#” 等', { icon: 2, title: '错误' });
                        } else {
                            win.layer.alert("请求失败：" + XHR.responseText, { icon: 2, title: '错误' });
                        }
                    }
                    else {
                        win.layer.alert("请求失败", { icon: 2, title: '错误' });
                    }
                }
            });
        }
    });
})(jQuery);
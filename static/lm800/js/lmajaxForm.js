/**
 * ajax form
 */
function ajax2Form(obj) {
    if (typeof obj == 'string') {
        obj = $(obj)
    }
    var options = {
        error: ajaxFormError,
        success: ajaxFormNoJump,
        forceSync: true,
        dataType: 'JSON',
        // clearForm: true,
        // resetForm: true,
        timeout: 3000
    }
    obj.ajaxSubmit(options)
    return false
    obj.ajaxForm(options)
}
/**
 * ajax 提交失败不跳转
 */
function ajaxFormError(data) {
    //配置一个透明的询问框
    layer.msg(data,  {
            time: 2000, //2s后自动关闭
            skin: 'layer-black-translucent'
        }
    );
}
/**
 * ajax 提交失败不跳转
 */
function ajaxFormNoJump(data){
    console.log('ajaxForm data', data)
    var btn = [];
    var options = {
        time: 2000,
        skin: 'layer-black-translucent'
    }
    if (data.obj.btn) {
        options.btn = data.obj.btn
    }
    if (data.code == 1) {
        // 提交失败；返回错误信息并清理相应字段
        layer.msg(data.msg, options, function(){
            if (data.obj.click!=undefined) {
                $(data.obj.click).click()
            }
            if (data.obj.reset_val!=undefined) {
                $(data.obj.reset_val).val("")
            }
            if (data.obj.focus!=undefined) {
                $(data.obj.focus).focus()
            }
        });
    } else {
        // 提交成功；显示提示信息
        //配置一个透明的询问框
        layer.msg(data.msg, options, function() {
                if (data.code == 302) {
                    if (data.obj.url != undefined) {
                        window.location.href = data.obj.url;
                    } else {
                        window.location.href = data.obj;
                    }
                }
            }
        );
    }
    en_submit = false;
}










/**
 * 搜索form导出数据
 */
$(function () {
    $('body').on('click','.ajax-export-csv',function () {
        var query = $(this).parents("form").serialize();
        var url = $(this).attr('href');

        window.location.href = url + '?' + query;
        // $.ajax({
        // 	type:"POST",
        // 	data:{},
        // 	success: function(data,status){
        // 		$("#ajax-data").html(data);
        // 	}
        // });
        return false;
    });
});


/**
 * ajaxform 表单提交
 * 提交前不检测
 * 失败不跳转
 */
$(function(){
    $(".ajaxFormNoCheckNoJump").ajaxForm({
        success: ajaxFormNoJump,
        dataType: 'json',
        forceSync: true,
    })
})
/**
 * ajaxform 表单提交
 * 提交前检测
 * 失败不跳转
 */
$(function () {
    $('.ajaxFormCheckNoJump').ajaxForm({
        // $('.ajaxFormCheckNoJump').ajaxSubmit({
        beforeSubmit: ajaxFormCheck, // 此方法主要是提交前执行的方法，根据需要设置
        success: ajaxFormNoJump, // 这是提交后的方法
        dataType: 'json',
        forceSync: true,
    });

    // 同步 ckeditor 内容到文本框
    $("body").on('click', ckeditorSync)
});

function  ckeditorSync() {
    if (CKEDITOR.instances == undefined || typeof CKEDITOR.instances == 'undefined') {
        return '';
    }
    // $('#Content').val(CKEDITOR.instances.Content.getData());
    for ( instance in CKEDITOR.instances )
    {
        // console.log('CKEDITOR.instances', CKEDITOR.instances[instance].getData())
        $("#"+ CKEDITOR.instances[instance].name).val(CKEDITOR.instances[instance].getData())
        CKEDITOR.instances[instance].updateElement();
    }
}
//提交数据之前检测信息合法
function ajaxFormCheck() {
    console.log('ajaxFormCheckNoJump')
    console.info($(this))
    validateInfo()
    ckeditorSync()
    var check_ok = true;
    $(".no-space").each(function(){
        var nowVal = $.trim($(this).val());
        if (nowVal.indexOf(" ")>=0 || nowVal.indexOf("	")>=0 || nowVal.indexOf("　")>=0) {
            layer.msg("输入信息包含有异常空格，请检查重新输入！",  {
                time: 2000, //2s后自动关闭
                btn: [],
                skin: 'layer-black-translucent'
            });
            $(this).focus();
            check_ok = false;
            return false;
        }
    })
    if (check_ok==false) {
        return false;
    }
}


/* get执行并返回结果，执行后带跳转 */
$(function () {
    $('body').on('click','.confirm-rst-url-btn6',function () {
        var $url = this.href,
            $info = $(this).data('info');
        layer.confirm($info, {skin: 'layer-skin6'}, function (index) {
            layer.close(index);
            $.get($url, function (data) {
                var btn = [];
                if (data.obj.btn) {
                    btn = data.obj.btn
                }
                if (data.code==1) {
                    layer.msg(data.msg,  {
                        time: 2000, //2s后自动关闭
                        btn: btn,
                        skin: 'layer-black-translucent'
                    });
                } else {
                    layer.msg(data.msg,  {
                            time: 2000, //2s后自动关闭
                            btn: btn,
                            skin: 'layer-black-translucent'
                        }, function() {
                            window.location.href = data.obj.url;
                        }
                    );
                }
            }, "json");
        });
        return false;
    });
});

/* 多选删除操作 */
en_submit = false;
$(function () {
    $('#alldel').ajaxForm({
        beforeSubmit: confirmSelectForm2, // 此方法主要是提交前执行的方法，根据需要设置，一般是判断为空获取其他规则
        success: ajaxFormNoJump, // 这是提交后的方法
        dataType: 'json'
    });
});

/**
 * 所有删除确认操作
 */
function confirmSelectForm2() {
    console.log(en_submit)
    if (en_submit) { return true; }
    var chk_value = [];
    $('input[id="navid"]:checked').each(function () {
        chk_value.push($(this).val());
    });

    if (!chk_value.length) {
        layer.msg('至少选择一个删除项',  {time: 2000, btn: [], skin: 'layer-black-translucent' });
        return false;
    }
    var msg = "确认执行操作？";

    layer.confirm(msg, {skin: 'layer-skin6'}, function (index) {
        layer.close(index);
        en_submit = true;
        $('#alldel').submit();
        return true;
    });
    return false;
}

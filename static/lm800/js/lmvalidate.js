/*!
 * 对jquery.validate简单封装，针对不同的前端框架，默认参数值可能不一样
 *
 */
(function ($) {
    $.fn.extend({
        lmvalidate: function (options) {
            // if nothing is selected, return nothing; can't chain anyway
            if (!this.length) {
                if (options && options.debug && window.console) {
                    console.warn("Nothing selected, can't validate, returning nothing.");
                }
                return;
            }
            var defaults={
                errorElement: 'span', //default input error message container
                errorClass: 'label label-danger margin-l5px form-error-msg', // default input error message class
                focusInvalid: false, // do not focus the last invalid input
                ignore: "", // validate all fields including form hidden input
                rules: {},
                messages: {},
                errorPlacement: function (error, element) { // render error placement for each input type
                    if (element.parent(".input-group").size() > 0) {
                        if (element.parent(".input-group").prev("div.tip-title").length > 0) {
                            element.parent(".input-group").prev("div.tip-title").append(error)
                        } else {
                            error.insertBefore(element.parent(".input-group"));
                        }
                    } else if (element.attr("data-error-container")) {
                        if (element.prev("div.tip-title").length > 0) {
                            element.prev("div.tip-title").append(error)
                        } else {
                            error.insertBefore(element.attr("data-error-container"));
                        }
                    } else if (element.parents('.radio-list').size() > 0) {
                        if (element.parents('.radio-list').prev("div.tip-title").length > 0) {
                            element.parents('.radio-list').prev("div.tip-title").append(error)
                        } else {
                            error.insertBefore(element.parents('.radio-list').attr("data-error-container"));
                        }
                    } else if (element.parents('.radio-inline').size() > 0) {
                        if (element.parents('.radio-inline').prev("div.tip-title").length > 0) {
                            element.parents('.radio-inline').prev("div.tip-title").append(error)
                        } else {
                            error.insertBefore(element.parents('.radio-inline').attr("data-error-container"));
                        }
                    } else if (element.parents('.checkbox-list').size() > 0) {
                        if (element.parents('.checkbox-list').prev("div.tip-title").length > 0) {
                            element.parents('.checkbox-list').prev("div.tip-title").append(error)
                        } else {
                            error.insertBefore(element.parents('.checkbox-list').attr("data-error-container"));
                        }
                    } else if (element.parents('.checkbox-inline').size() > 0) {
                        if (element.parents('.checkbox-inline').prev("div.tip-title").length > 0) {
                            element.parents('.checkbox-inline').prev("div.tip-title").append(error)
                        } else {
                            error.insertBefore(element.parents('.checkbox-inline').attr("data-error-container"));
                        }
                    } else if (element.parents('.bs-select').size() > 0) {
                        if (element.parents('.bs-select').prev("div.tip-title").length > 0) {
                            element.parents('.bs-select').prev("div.tip-title").append(error)
                        } else {
                            error.insertBefore(element.parents('.bootstrap-select:eq(0)'));
                        }
                    } else {
                        if (element.prev("div.tip-title").length > 0) {
                            element.prev("div.tip-title").append(error)
                        } else {
                            error.insertBefore(element); // for other inputs, just perform default behavior
                        }
                    }
                },
                invalidHandler: function (event, validator) { //display error alert on form submit
                    //验证不通过时
                },
                highlight: function (element) { // hightlight error inputs
                    $(element).closest('.form-group').removeClass('has-success').addClass('has-error');
                },
                unhighlight: function (element) {
                    $(element).closest('.form-group').removeClass('has-error');
                },
                success: function (label) {
                    label.closest('.form-group').removeClass('has-error').addClass("has-success");
                },
                submitHandler: null
            };
            var destOptions = $.extend({},defaults, options);
            //调用jquuery.validate插件方法
            this.validate(destOptions);
        },
    });
})(jQuery);
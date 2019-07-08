/*
* @Author: Actor Liu
* @Date:   2018-12-26 21:42:30
* @Last Modified by:   Actor Liu
* @Last Modified time: 2019-07-08 17:08:02
*/
CKEDITOR.editorConfig = function( config ) {
	config.toolbarGroups = [
		{ name: 'document', groups: [ 'mode', 'document', 'doctools' ] },
		{ name: 'clipboard', groups: [ 'clipboard', 'undo' ] },
		{ name: 'editing', groups: [ 'find', 'selection', 'spellchecker', 'editing' ] },
		{ name: 'forms', groups: [ 'forms' ] },
		{ name: 'styles', groups: [ 'styles' ] },
		{ name: 'colors', groups: [ 'colors' ] },
		{ name: 'tools', groups: [ 'tools' ] },
		{ name: 'others', groups: [ 'others' ] },
		{ name: 'about', groups: [ 'about' ] },
		'/',
		{ name: 'basicstyles', groups: [ 'basicstyles', 'cleanup' ] },
		{ name: 'paragraph', groups: [ 'list', 'indent', 'blocks', 'align', 'bidi', 'paragraph' ] },
		{ name: 'links', groups: [ 'links' ] },
		{ name: 'insert', groups: [ 'insert' ] },
	];

	config.removeButtons = 'Save,NewPage,Preview,Print,Templates,Cut,Copy,Paste,PasteText,PasteFromWord,Find,Replace,SelectAll,Scayt,Form,Checkbox,Radio,TextField,Textarea,Select,Button,ImageButton,HiddenField,CopyFormatting,RemoveFormat,Blockquote,CreateDiv,BidiRtl,BidiLtr,Language,SpecialChar,Iframe,Flash,ShowBlocks,About';

	// config.defaultLanguage = 'zh-cn';
	config.language = 'zh-cn';

	config.allowedContent = true; 	// 抓取什么就是什么；不过滤html标签
	config.autoGrow_minHeight = 200;	// 编辑器可以采用的最小高度
	
	// Configure your file manager integration. This example uses CKFinder 3 for PHP.
	// config.filebrowserBrowseUrl = '/apps/ckfinder/3.4.5/ckfinder.html';
	// config.filebrowserImageBrowseUrl = '/apps/ckfinder/3.4.5/ckfinder.html?type=Images';
	config.filebrowserUploadUrl = '/common/upload.html?command=QuickUpload&type=Files&responseType=json&refer=CKEDITOR';	// 超链接上传地址 
	config.filebrowserImageUploadUrl = '/common/upload.html?command=QuickUpload&type=Images&responseType=json&refer=CKEDITOR';	// 超链接上传地址 
	// config.removeDialogTabs = 'link:advanced;link:target;image:Upload;image:target;flash:Upload;flash:target;'; // 移除超链接的“目标”、“上传”工具栏
	config.removeDialogTabs = 'link:advanced;image:advanced;flash:Upload;'; // 移除超链接的“高级”、图片的“高级”、flash的“上传”工具栏



	// Adding drag and drop image upload.
	// config.extraPlugins = 'print,format,font,colorbutton,justify,uploadimage';
	config.uploadUrl = '/common/upload.html?command=QuickUpload&type=Files&responseType=json&refer=CKEDITOR';

	config.height = 460;

	// config.removeDialogTabs = 'image:advanced;link:advanced';
};

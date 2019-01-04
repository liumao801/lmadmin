/*
* @Author: Actor Liu
* @Date:   2018-12-26 21:42:30
* @Last Modified by:   Actor Liu
* @Last Modified time: 2019-01-01 08:19:02
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
	
	// config.filebrowserUploadUrl = '/common/upload.html?command=QuickUpload&type=Files&responseType=json&refer=CKEDITOR';	// 超链接上传地址 
	// config.filebrowserUploadUrl = '/upload/commonupload.html?command=QuickUpload&type=Files&responseType=json&refer=CKEDITOR';	// 超链接上传地址 
	config.filebrowserUploadUrl = '/test/upload?command=QuickUpload&type=Files&responseType=json&refer=CKEDITOR';	// 超链接上传地址 
	// config.filebrowserUploadUrl = '/home/test/upload.html?command=QuickUpload&type=Files&responseType=json&refer=CKEDITOR';	// 超链接上传地址 
	// config.filebrowserUploadUrl = '/admin/test/upload.html?command=QuickUpload&type=Files&responseType=json&refer=CKEDITOR';	// 超链接上传地址 
    config.removeDialogTabs = 'link:advanced;link:target;image:Upload;image:target;flash:Upload;flash:target;'; // 移除超链接的“目标”、“上传”工具栏

};
// 获取状态文字
function GetStatusText(value) {
    var texts = [{ Text: '停用', Value: 0 }, { Text: '正常', Value: 1 }, { Text: '已删除', Value: 2 } ]
    return lmtheme.showstatus(value, texts);
}
// 获取显示文字
function GetShowText(value) {
    var texts = [{ Text: '隐藏', Value: 0 }, { Text: '显示', Value: 1 }, { Text: '已删除', Value: 2 } ]
    return lmtheme.showstatus(value, texts);
}


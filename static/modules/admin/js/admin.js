
function GetStatusText(value) {
    var texts = [{ Text: '停用', Value: 0 }, { Text: '正常', Value: 1 }, { Text: '已删除', Value: 2 } ]
    return lmtheme.showstatus(value, texts);
}
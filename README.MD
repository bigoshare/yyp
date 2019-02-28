## yyp
本项目用于yy协议的组包和解包。

####用法
	// 定义包结构和uri
    const URI_PServerPing2Res = 32<<8 | 18
    type PServerPing2Res struct {
        ServerId uint32
        IsRecord bool
    }

    // marshal
    var pkg PServerPing2Res
    yyp.Marshal(URI_PServerPing2Res, pkg)
    buff, _ := yyp.Marshal(URI_PServerPing2Res, pkg)

    // unmarshal
    request := yyp.LoadRequest(buff)

    if request.Uri == URI_PServerPing2Res {
        res := PServerPing2Res{}
        yyp.Unmarshal(request.Data, &res)

        // 处理PServerPing2Res包
        return
    }
	
#### 结构体定义注意事项
#####1、所有字段都会被marshall
#####2、字段名必须大写（golang权限控制，小写为私有结构体无法被外部访问）

#### 关于超大string类型字段的组包说明
string类型默认为marshall的时候是16个字节长度，再加上字符串本身。
有的协议需要用到超大的string类型，因此定义为32个字节的长度。定义结构体的时候需要定义为string32
	
	
	

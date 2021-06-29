




/**
 * 2 手机分辨率 1080*1920
 * 国内抖音版本 16.1.0
 */

/**
 * 修改系统 toast 方法!
 */
var _toast_ = toast;
toast = function (message) {
    _toast_(message);
    sleep(2000);
}

// 请求超时! 10秒
http.__okhttp__.setTimeout(10000);

//获取 百度的SK AK

手机屏幕宽 = 1080

var nickname = ""
var 域名 = ""
var 全局抖音ID = ""
path = "/sdcard/Pictures/nickname.txt";
var 全局的最后一个加粉名字 = ""

if (files.exists(path)) {
    nickname = files.read(path);
    toast("昵称:" + nickname)
} else {
    toast("昵称不存在,请写好配置文件!!!在运行")
    exit();
}



path = "/sdcard/Pictures/ip.txt";
if (files.exists(path)) {
    var 域名 = files.read(path);
    toast("使用域名:" + 域名)
} else {
    toast("域名不存在,请写好配置文件!!!在运行")
    exit();

}

域名 = "172.81.248.168:2345"







function 上传用户抖音号(用户抖音) {

    try {
        域名 = "47.242.44.105:9501"
        url = "http://" + 域名 + "/two_add_id?nickname=" + nickname + "&douyi_id=" + 用户抖音

        log(url)

        var res = http.get(url)

    } catch (error) {
        console.log("上传用户抖音号 异常:" + error);
    }
}



function 关闭应用(packageName) {
    var 关闭应用现在的时间戳 = timest()
    var name = getPackageName(packageName);
    if (!name) {
        if (getAppName(packageName)) {
            name = packageName;
        } else {
            return false;
        }
    }
    app.openAppSetting(name);
    text(app.getAppName(name)).waitFor();
    while (true) {
        if (timest() - 关闭应用现在的时间戳 > 30) {
            //等待超时大于30 秒
            var name = getPackageName(packageName);
            if (!name) {
                if (getAppName(packageName)) {
                    name = packageName;
                } else {
                    return false;
                }
            }
            app.openAppSetting(name);
            text(app.getAppName(name)).waitFor();
            var 关闭应用现在的时间戳 = timest()
        } else {
            if (id("right_button").exists()) {
                text("FORCE STOP").click();
                sleep(2000)
                text("OK").click()
                break  //退出循环
            } else {
                toast("right_button 按钮不存在")
                back() //返回
            }
        }


        sleep(2000)
    }


}









function 判断粉丝群是否存在() {
    //说明 粉丝群存在
    if (id('jka').exists()) {
        toast("这个粉丝存在粉丝群....")
        //获取 抖音号
        if (id('kce').exists()) {
            var douyin_id = id('kce').findOne().text()
            toast(douyin_id)
            上传用户抖音号(douyin_id)

        } else[
            toast("获取粉丝抖音号失败....,原因,没有找到 id:kce")

        ]
    } else {
        toast("不存在粉丝群.....")

    }
}








function 切换短视频() {
    toast("准备切换一下个视频.....")
    swipe(500, 1300, 500, 200, 300)


}




function 抖音判断是否首页() {
    if (text("我").exists()) {
        var 我的坐标 = text("我").findOne().bounds().centerX()
        if (我的坐标 > 0 && 我的坐标 < 1080) {
            return true
        } else {
            return false
        }
    } else {
        return false
    }
}



function 抖音判断是否在个人资料界面() {
    if (desc("更多").exists()) {
        var 我的坐标 = desc("更多").findOne().bounds().centerX()
        if (我的坐标 > 0 && 我的坐标 < 1080) {
            return true
        } else {
            return false
        }
    } else {
        return false
    }
}







/**
 *  收集粉丝的主程序
 */
function 采集主程序() {
    var hint = ""
    var 循环延迟时间 = 2   //单位秒
    var 是否点击 = false
    var 是否滑动 = false
    var 等待进入资料界面次数 = 0
    var 未知界面次数 = 0
    var 等待个人资料次数 = 0
    var 是否点击返回 = false
    while (true) {
        if (抖音判断是否首页()) {
            if (是否滑动) {
                if (是否点击) {
                    toast("等待进入资料界面.......")
                    if (等待进入资料界面次数 > 5) {
                        是否滑动 = false
                    } else {
                        等待进入资料界面次数++
                        sleep(100)
                    }

                } else {
                    if (等待个人资料次数 > 10) {
                        // 可能在其他的界面
                        // back()
                        // sleep(1000)
                        // back()

                        系统返回方法重写()

                        是否滑动 = false
                    } else {
                        toast("在抖音主页......准备点击进入资料....")
                        // click(970, 780)
                        weight = id("user_avatar").find()
                        if (weight) {
                            if (weight.length <= 1) {
                                toast("不是正常的首页界面......")
                                是否滑动 = false

                            } else {
                                if (weight[0].bounds().centerY() > 0 && weight[0].bounds().centerY() < 1920) {
                                    click(970, weight[0].bounds().centerY())
                                } else {

                                    if (weight[1].bounds().centerY() > 0 && weight[1].bounds().centerY() < 1920) {
                                        click(970, weight[1].bounds().centerY())
                                    } else {
                                        click(970, 780)
                                    }
                                }
                            }

                        } else {
                            console.log("789")
                            click(970, 780)
                        }

                        是否点击 = true
                        等待个人资料次数++
                    }


                }

            } else {
                等待进入资料界面次数 = 0
                切换短视频()
                是否滑动 = true
                是否点击 = false
                未知界面次数 = 0
                等待个人资料次数 = 0
                是否点击返回 = false
            }
        } else if (抖音判断是否在个人资料界面()) {
            if (是否点击返回) {
                系统返回方法重写()
            } else {
                toast("个人资料页....")
                判断粉丝群是否存在()
                toast("准备点击返回.....")
                系统返回方法重写()
                是否滑动 = false
                是否点击 = false
                是否点击返回 = true
            }

            // 

        } else if (id('a2w').exists()) {
            toast("进入直播了...")
            id("a2w").findOne().click()
        } else {
            if (未知界面次数 > 10) {
                // back()
                // sleep(1000)
                // back()
                系统返回方法重写()
                是否滑动 = false
            } else {
                toast("采集程序的未知界面.....")
                未知界面次数++
            }

        }
        sleep(循环延迟时间 * 1000)
    }

}



function 关闭通讯录发现好友() {
    if (id("tv_cancel").exists()) {
        toast("关闭通讯录")
        click(300, 1300)
    }
}


function 获取抖音ID() {
    for (var i = 0; i < 10000000; i++) {
        try {
            域名 = "47.242.44.105:9501"
            url = "http://" + 域名 + "/get_two_dy_id?username=" + nickname

            log(url)
            var res = http.get(url);

            if (res.statusCode != 200) {
                toast("请求失败: " + res.statusCode + " " + res.statusMessage);
            } else {
                if (res != "") {
                    var weather = res.body.json();
                    //获取 百度的access_token
                    if (weather.code == 200) {  //等于1  的时候就是运行 脚本了..
                        //toast("开始运行!!")
                        return weather.msg;
                    }
                }
            }
        } catch (error) {
            console.log("获取联系方式:" + error)
        }
        sleep(2000)
    }
}






/**
 * 添加群
 */


function 添加群主程序() {
    launch("com.ss.android.ugc.aweme");
    toast("开启抖音")
    sleep(3000)
    hint = "加粉"
    var 循环延迟时间 = 2
    var 是否输入 = false
    var 等待搜索信息 = 0
    var 未知界面次数 = 0
    while (true) {

        if (hint == "") {
            关闭应用("com.ss.android.ugc.aweme")
            sleep(3000)
            未知界面次数 = 0
            launch("com.ss.android.ugc.aweme");
            toast("开启抖音")
            sleep(3000)
            hint = "加粉"

        } else if (hint == "加粉") {

            if (抖音判断是否首页()) {
                toast("首页....点击搜索按钮")
                id("clc").findOne().click()

            } else if (id("et_search_kw").exists()) {
                if (是否输入) {
                    if (id("drh").exists()) {
                        toast("准备点击进入")
                        click(100, 300)
                        是否输入 = false


                    } else if (text("用户").exists()) {  //领一种方式哦
                        toast("点击第一个用户..")
                        click(100, 620)

                    } else {
                        if (等待搜索信息 > 5) {
                            toast("点击搜索....")
                            click(950, 100)  //点击按钮
                            等待搜索信息 = 0

                        } else {

                            toast("没有搜索任何信息.....")
                            等待搜索信息++

                        }

                    }
                } else {
                    toast("准备输入抖音号")
                    var ID = 获取抖音ID()
                    id("et_search_kw").findOne().setText(ID)
                    是否输入 = true
                }



            } else if (抖音判断是否在个人资料界面()) {
                toast("个人资料界面....")

                if (text("粉丝群").exists()) {

                    if (text("粉丝群").findOne().bounds().centerX() > 0 && text("粉丝群").findOne().bounds().centerX() < 1080 && text("粉丝群").findOne().bounds().centerY() > 0 && text("粉丝群").findOne().bounds().centerY() < 1920) {
                        click(text("粉丝群").findOne().bounds().centerX(), text("粉丝群").findOne().bounds().centerY())

                    } else {
                        toast("框架异常.......返回切换下一个")
                        if (id("back_btn").exists()) {
                            id("back_btn").findOne().click()
                        } else {
                            hint = ""
                        }
                    }
                } else {
                    toast("粉丝群不存在.....准备点击返回")
                    if (id("back_btn").exists()) {
                        id("back_btn").findOne().click()
                    } else {
                        hint = ""

                    }
                }

            } else if (id("f45").exists()) {
                toast("加群界面......")
                if (text("申请加入").exists()) {
                    var uc = text("申请加入").find();
                    for (var i = 0; i < uc.length; i++) {
                        var Y = uc[i].bounds().centerY()
                        click(540, Y)  // 点击申请加群



                        toast("申请加群..")


                        if (text("进群条件不符").exists()) {
                            //这里 可能有各种原因  如果有

                            if (text("我知道了").exists()) {
                                text("我知道了").findOne().click()
                            }

                            hint = ""


                        }

                    }

                } else {
                    toast("没有可以加的群")
                    hint = ""
                }


            } else {

                if (未知界面次数 > 30) {
                    hint = ""
                } else {
                    未知界面次数++
                    toast("添加群主程序未知界面.......")
                }

            }




        }
        sleep(循环延迟时间 * 1000)
    }

}




// 添加群主程序()
function timest() {
    var tmp = Date.parse(new Date()).toString();
    tmp = tmp.substr(0, 10);
    return tmp;
}



function 系统返回方法重写() {
    var 当前时间戳 = timest()
    while (true) {
        if (timest() - 当前时间戳 > 30) {
            关闭应用("com.ss.android.ugc.aweme")
            sleep(3000)
            未知界面次数 = 0
            launch("com.ss.android.ugc.aweme");
            toast("开启抖音")
            sleep(3000)
            当前时间戳 = timest()
        } else {
            if (抖音判断是否首页()) {
                toast("返回成功....在首页.....")
                break
            } else {
                toast("点击返回")
                back()

            }
        }
        sleep(1000)
    }

}









// 采集主程序()




function 关注返回手机通讯录界面() {
    if (id("back_btn").exists()) {
        id("back_btn").findOne().click()
    } else {
        click(80, 130)
    }
}


function 发现通讯录好友(value) {
    if (text("发现通讯录好友").exists()) {
        if (value) {
            //如果为真 点击暂时不要
            text("暂时不要").findOne().click()
        } else {
            //点击发现好友
            text("发现好友").findOne().click()
        }
    }
}


function 是否需要关注界面() {
    if (text("关注").exists()) {
        var ll = text("关注").find()
        if (ll.length == 3) {
            return true
        }
    }
    return false

}



function 私聊() {
    for (var i = 0; i < 5; i++) {
        if (id("d80").exists()) {
            //判断是否私聊过
            if (text("已送达").exists()) {
                toast("不要重复的聊天!!")
                sleep(1000)
                toast("返回")
                id("enl").click()
            } else {
                toast("准备聊天....")
                id("foe").findOne().setText("你好!")
                sleep(2000)
                desc("发送").click()
                toast("发送完毕,准备返回")
                id("enl").click()
            }
            break
        } else {
            toast("等待进入私聊界面......")
        }

    }
}


function 关注并且私信() {
    var 现在的时间戳 = timest()
    while (true) {
        if (timest() - 现在的时间戳 > 60) {
            toast("私信超时")
            return false
        } else {
            if (text("回关").exists()) {
                text("回关").findOne().click()
                sleep(1000)
            } else if (desc("私信").exists()) {
                toast("准备点击私信")
                desc("私信").findOne().click()
                sleep(2000)
                私聊()
                sleep(2000)
                关注返回手机通讯录界面()
                return true
            } else if (是否需要关注界面()) {
                toast("点击关注")
                // text("关注").findOne().click()
                click(600, 420)
                sleep(1000)
            } else if (text("已请求").exists()) {
                //点击的返回退出循环
                关注返回手机通讯录界面()
                return true   //退出循环
            } else {
                toast("正在关注未知界面........")
            }
            toast("正在关注.......")
            sleep(1000)
        }
    }

}






function 遍历通讯录可关注数() {
    if (id("kc-").exists()) {
        var list = id("kc-").find();
        if (list.length == 0) {
            toast("需要遍历的关注的控件不存在")
            return true
        }
        for (var i = 0; i < list.length; i++) {
            var tv = list[i];
            var X = 100
            var Y = tv.bounds().centerY()
            if (Y > 1843 || Y < 0) {
                continue   //不符合 继续
            }

            toast("准备点点击坐标:(900," + Y + ")")
            click(X, Y)
            if (!关注并且私信()) {
                return false
            }


            // if (i == list.launch - 1) {
            //     if (id("kd5").exists()) {
            //         var a = id("kd5").find()
            //         全局的最后一个加粉名字 = a[i].text()
            //     }
            // }

        }



        return true
    } else {
        toast("需要遍历的关注的控件不存在")   //这种情况下就是没有通讯录  可以个
        console.log(1121212);
        return true

    }

}







function 关注通讯判断是否到底() {
    if (id("kd5").exists()) {
        var a = id("kd5").find()
        return a[a.length - 1].text()
    } else {
        return ""
    }
}


function 判断直播间观众是否到底() {
    try {
        if (id("user_name").exists()) {
            var a = id("user_name").find.find()
            return a[a.length - 2].text()
        } else {
            return ""
        }
    } catch (error) {
        return ""
    }


}

function 添加通讯录好友() {
    var 是否点击我 = false
    var 现在时间戳 = 0
    var 是否可以滑动 = false
    var hint = ""
    var 判断是否到底 = ""
    while (true) {
        if (hint == "") {
            if (抖音判断是否首页()) {
                if (是否点击我) {
                    if (timest() - 现在时间戳 > 10) {
                        toast("等待个人信息界面超时")
                        现在时间戳 = timest()
                        是否点击我 = false
                    } else {
                        if (抖音判断是否在个人资料界面()) {
                            toast("个人资料界面...准备点击+朋友")
                            id("hc").findOne().click()  //点击加好友
                            是否点击我 = false
                        } else {
                            toast("等待进入个人资料界面......")
                        }
                    }
                } else {
                    现在时间戳 = timest()
                    click(text("我").findOne().bounds().centerX(), text("我").findOne().bounds().centerY())
                    toast("在首页.....点击我")
                    是否点击我 = true
                }

            } else if (text("发现朋友").exists()) {
                if (text("查看通讯录朋友").exists()) {
                    toast("点击查看通讯录朋友")
                    click(100, 400)
                } else {
                    toast("没有发现..........")   //这里我后面需要补充
                }
            } else if (text("手机通讯录").exists()) {
                if (是否可以滑动) {
                    //判断是否到底
                    判断是否到底 = 关注通讯判断是否到底()
                    swipe(100, 1877, 100, 351, 1200)
                    sleep(2000)
                    是否可以滑动 = false
                    if (关注通讯判断是否到底() == 判断是否到底) {
                        log("---------------------------到底了,程序结束")
                        exit()
                    }
                } else {
                    if (!遍历通讯录可关注数()) {
                        hint = "重启"
                    }
                    是否可以滑动 = true

                }
            } else {
                toast("未知界面.....")
            }
        } else if (hint == "重启") {
            关闭应用("com.ss.android.ugc.aweme")
            sleep(3000)
            未知界面次数 = 0
            launch("com.ss.android.ugc.aweme");
            toast("开启抖音")
            sleep(3000)
            是否点击我 = false
            现在时间戳 = 0
            是否可以滑动 = false
            hint = ""

        }
        sleep(2000)
    }
}



/**
 * @version 抖音版本 16.1.0
 * 
 */
function 直播间关注() {
    var 是否滑动 = false
    var 最后一个观众名字 = ""
    hint = ""
    while (true) {
        if (hint == "") {
            if (id("f7c").exists() || text("说点什么...").exists()) {
                toast("正在直播间主界面.......")
                click(900, 120)
                // } else if (id("cp6").exists()) {
            } else if (id("user_avatar").exists()) {
                if (是否滑动) {
                    var 最后一个观众名字 = 判断直播间观众是否到底()
                    toast("准备滑动.....")
                    swipe(160, 1633, 160, 351, 856)
                    if (判断直播间观众是否到底() == 最后一个观众名字) {
                        log("到底了................")
                        hint = "切换"

                    }
                    是否滑动 = false
                } else {
                    var list = id("cox").find();
                    if (list.length == 0) {
                        toast("需要遍历的关注的控件不存在")
                    }
                    for (var i = 0; i < list.length; i++) {
                        var tv = list[i];
                        var X = 180
                        var Y = tv.bounds().centerY()
                        // log(a[i].text())
                        // click(X, Y)
                        if (id("ivj").exists()) {
                            log("坐标:" + X + "," + Y)
                            var a = id("ivj").find()
                            if (a[i].text() > 0 && a[i].text() < 1000) {
                                click(X, Y)
                                //这里给一个延迟时间
                                for (var i = 0; i < 10; i++) {
                                    toast("关注成功.......休息时间.....")
                                }
                            }

                        }
                        sleep(2000)
                    }

                    是否滑动 = true
                }

            } else if (id("g6v").exists()) {  //举报
                toast("点错了.......")
                back()


            } else {
                toast("观看直播未知界面..............")


            }
        } else if (hint == "切换") {
            toast("准备切换下一个直播")
            click(500, 300)
            sleep(2000)
            swipe(160, 1633, 160, 351, 856)
            sleep(2000)
            hint = ""
        }
        sleep(2000)
    }

}




//直播间关注()

//live: "snssdk1128://live?room_id={{room_id}}&user_id={{user_id}}&from=webview&refer=web",

// app.startActivity({
//     data: "snssdk1128://aweme/detail/6975844403248253734" 
// });

// app.startActivity({
//     data: "snssdk1128://live?room_id=6975684821527841577" 
// });



function 直播间粉丝采集数据上传(年龄, 地址, 获赞, 关注, 粉丝, 个性签名, 性别, 抖音号) {

    var host = "http://" + 域名 + "/collect/UploadInformation?age=" + 年龄 + "&address=" + 地址 + "&like=" + 获赞 + "&following=" + 关注 + "&follows=" + 粉丝 + "&sign=" + 个性签名 + "&sex=" + 性别 + "&dyNumber=" + 抖音号 + "&username=" + nickname + "&type=1"

    log(host)
    try {
        var res = http.get(host);
    } catch (error) {
        console.log(" 直播间粉丝采集数据上传  异常抛出:" + error)
    }

}





function 进入采集用户列表() {
    var time1 = timest()  //初始时间
    var time2 = timest()
    var 年龄 = ""
    var 地址 = ""
    var 获赞 = ""
    var 关注 = ""
    var 粉丝 = ""
    var 个性签名 = ""
    var 性别 = ""
    var 抖音号 = ""
    var 是否点击了返回 = false
    var 是否采集 = false
    while (true) {
        try {
            if (timest() - time1 > 60) {  //两分钟的超时
                toast("超时......")
                break
            } else {
                if (text("举报").exists()) {
                    if (是否采集) {
                        if (是否点击了返回) {
                            if (timest() - time2 > 10) {
                                是否点击了返回 = false
                            } else {
                                toast("等待返回.....")
                            }
                        } else {
                            back()
                            是否采集 = false
                            是否点击了返回 = true
                        }
                    } else {
                        //获取性别
                        if (id("cvd").exists()) {
                            性别 = id("cvd").findOne().desc()
                            //点击主页  无性别不点
                            toast("点击主页.....")
                            text("主页").findOne().click()
                        } else {
                            //没有性别  点击返回
                            if (是否点击了返回) {
                                if (timest() - time2 > 10) {
                                    是否点击了返回 = false
                                } else {
                                    toast("等待返回.....")
                                }
                            } else {
                                back()
                                是否点击了返回 = true
                            }
                        }

                    }
                } else if (id("ix5").exists()) {
                    log("采集资料界面.................")
                    var a = className("android.widget.TextView").find()
                    for (var i = 0; i < a.length; i++) {
                        try {
                            // log(a[i].text())
                            if (a[i].text() == "获赞") {
                                if (i == 0) {
                                    //磨板2
                                    获赞 = a[1].text()
                                    关注 = a[5].text()
                                    粉丝 = a[3].text()
                                    抖音号 = a[7].text()
                                    if (判断控件性别属性(a[8].text())) {
                                        年龄 = a[8].text()
                                        地址 = a[9].text()
                                    } else {
                                        个性签名 = a[8].text()
                                        if (判断控件性别属性(a[9].text())) {
                                            年龄 = a[9].text()
                                            地址 = a[10].text()
                                        }
                                    }

                                } else {
                                    //磨板2
                                    if (判断控件性别属性(a[3].text())) {
                                        //说明了没有  个人签名
                                        年龄 = a[3].text()
                                        if (i > 5) {
                                            地址 = a[4].text()
                                        }
                                    } else {
                                        个性签名 = a[3].text()
                                        年龄 = a[4].text()
                                        //判断是否有 家乡地址
                                        if (i < 7) { //没有家乡
                                        } else {
                                            //有家乡地址
                                            地址 = a[5].text()
                                        }
                                    }
                                    抖音号 = a[2].text()
                                    获赞 = a[i - 1].text()
                                    关注 = a[i + 1].text()
                                    粉丝 = a[i + 3].text()
                                   

                                }
                            }

                        } catch (error) {
                            toast("文本读取异常:" + error)
                            log("文本读取异常:" + error)

                        }

                    }

                    抖音号 = 抖音号.slice(4)
                    log("抖音号:" + 抖音号)
                    直播间粉丝采集数据上传(年龄, 地址, 获赞, 关注, 粉丝, 个性签名, 性别, 抖音号)
                    //采集的数据上传 
                    if (id("back_btn").exists()) {
                        toast("点击返回.......")
                        id("back_btn").findOne().click()
                    } else {
                        toast("点击返回.......")
                        click(80, 140)
                    }

                    是否采集 = true
                    sleep(2000)
                } else if (text("帮助").exists()) {
                    toast("点错了,准备返回")
                    click(80, 170)
                } else if (id("user_avatar").exists()) {
                    toast("返回成功")
                    break;
                } else {
                    toast("进入采集用户列表---位置界面")

                }
            }
            sleep(2000)

        } catch (error) {
            toast("方法 进入采集用户列表() 异常抛出:" + error)
            sleep(2000)
        }
    }
}













function 遍历观众数量() {
    if (id("user_avatar").exists()) {
        var list = id("user_avatar").find()
        if (list.length == 0) {
            toast("需要遍历的关注的控件不存在")
        }
        for (var i = 0; i < list.length; i++) {
            var tv = list[i];
            var X = 180
            var Y = tv.bounds().centerY()
            if (Y > 0 && Y < 1696) {
                log(X, Y)
                click(X, Y)
                进入采集用户列表()
            }
            sleep(2000)
        }

    } else {
        toast("没有找到头像控件,准备滑动......")

    }
}








function 判断控件性别属性(sex) {
    if (sex.length > 3) {
        return false;
    } else {
        if (sex.indexOf("男") != -1 && sex.length == 1) {
            return true
        } else if (sex.indexOf("女") != -1 && sex.length == 1) {

            return true
        } else if (sex.indexOf("岁") != -1 && sex.length == 3) {
            return true
        }
    }
    return false
}






function 直播已经结束() {

    if (text("直播已结束").exists()) {
        return true
    }

    return false
}

function 获取直播间roomId() {
    host = 'http://' + 域名 + '/link/GetOneLink?username=' + nickname + "&type=1"  // type=1 直播间id
    for (var i = 0; i < 10000000; i++) {
        try {
            var res = http.get(host);
            if (res.statusCode != 200) {
                toast("请求失败: " + res.statusCode + " " + res.statusMessage);
            } else {
                var weather = res.body.json()
                if (weather) {
                    log(weather)
                    if (weather.code == 200) {
                        toast("获取到id:" + weather.msg);
                        全局抖音ID = weather.msg
                        app.startActivity({
                            data: "snssdk1128://live?room_id=" + weather.msg,
                        });
                        break;
                    } else {
                        toast(weather.msg)
                    }
                }

            }
            sleep(2000)

        } catch (error) {
            toast("异常退出!!!!");

        }
    }
}




// 遍历观众数量()




/***
 * 
 * @version 16.1.0
 * 
 * 服务器获取连接
 */

function 直播界面采集数据() {
    var hint = ""
    var 是否已经获取连接 = false
    var 现在时间戳 = timest()
    while (true) {
        if (hint == "") {
            if (是否已经获取连接) {
                if (直播已经结束()) {
                    // 获取直播间roomId()
                    toast("播主已经下播了,切换直播间")
                    是否已经获取连接 = false
                } else if (id("f2=").exists()) {
                    toast("进入直播间成功...准备采集....")
                    hint = "采集"
                } else {
                    //给一个定时
                    if (timest() - 现在时间戳 > 30) {
                        是否已经获取连接 = false
                    } else {
                        toast("未知界面..........")
                        sleep(1000)
                    }
                }

            } else {
                现在时间戳 = timest()
                获取直播间roomId()
                是否已经获取连接 = true
            }


        } else if (hint == "采集") {
            if (text("说点什么...").exists()) {  //直播间首页
                toast("直播间首页,准备点击在线关注数,准备采集")
                是否已经获取连接 = false
                click(900, 120)
            } else if (id("user_avatar").exists()) {
                if (遍历观众数量()) {

                } else {
                    hint = "滑动"
                }

            } else {
                //这里进行 进行超时 判断
                toast("未知界面,hint=采集")

            }


        } else if (hint == "滑动") {
            var 最后一个观众名字 = 判断直播间观众是否到底()
            toast("准备滑动.....")
            swipe(160, 1633, 160, 351, 856)
            if (判断直播间观众是否到底() == 最后一个观众名字) {
                log("到底了................")
                hint = ""
            } else {
                hint = "采集"
            }

        } else {
            toast("hint 位置属性")

        }
        sleep(2000)
    }
}


/***
 * 手动滑获取  直播间采集
 * 
 */


function 直播间采集数据手动滑动() {
    var 是否滑动 = false
    var 最后一个观众名字 = ""
    var 现在时间戳 = timest()
    var 是否点击观众 = false
    hint = "采集"
    while (true) {
        if (hint == "采集") {
            if (text("说点什么...").exists()) {  //直播间首页
                if (是否点击观众) {
                    if (timest() - 现在时间戳 > 30) {
                        toast("点击观众列表超时......")
                        是否点击观众 = false
                        hint = "切换"
                    } else {
                        toast("等待进入观众列表")
                    }
                } else {
                    现在时间戳 = timest()
                    toast("直播间首页,准备点击在线关注数,准备采集")
                    是否已经获取连接 = false
                    click(900, 120)
                    是否点击观众 = true
                }
                
            } else if (text("检测到更新").exists()) {
                toast("关闭更新,以后再说......")
                text("以后再说").findOne().click()
                现在时间戳 = timest()
            } else if (text("推荐").exists()) {
                toast("准备点击直播...")
                id("i17").findOne().click()
                现在时间戳 = timest()
            } else if (id("user_avatar").exists()) {
                if (遍历观众数量()) {
                    现在时间戳 = timest()
                } else {
                    hint = "滑动"
                }
            } else {
                //这里进行 进行超时 判断
                if (timest() - 现在时间戳 > 60) {
                    toast("hint=采集超时准备重启....")
                    hint = "重启"
                } else {
                    toast("未知界面,hint=采集")

                }
            }
        } else if (hint == "切换") {
            toast("准备切换下一个直播")
            click(500, 300)
            sleep(2000)
            swipe(160, 1633, 160, 351, 856)
            sleep(2000)
            hint = "采集"
        } else if (hint == "滑动") {
            var 最后一个观众名字 = 判断直播间观众是否到底()
            toast("准备滑动.....")
            swipe(160, 1633, 160, 351, 856)
            if (判断直播间观众是否到底() == 最后一个观众名字) {
                log("到底了................")
                hint = "切换"
            } else {
                hint = "采集"
            }

        } else if (hint == "重启") {
            关闭应用("com.ss.android.ugc.aweme")
            sleep(3000)
            launch("com.ss.android.ugc.aweme");
            toast("开启抖音")
            sleep(3000)
            hint = "采集"
        } else {



            toast("hint 位置属性")

        }
        sleep(2000)
    }


}






直播间采集数据手动滑动()








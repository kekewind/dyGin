
    //2021年5月22日22:26:59 优化修改

    js_global_requestAddress="";

    //添加链接
    js_global_requestAddress = "http://172.81.248.168:2345/link/AnalysisLink"


    //获取链接列表
    js_global_requestAddress = "http://172.81.248.168:2345/link/GetLinksList"


    js_global_requestloginAddress="/login";










    //获取目录路径方法
    function getRootPath_web() {

        //获取当前网址，如： http://localhost:8888/eeeeeee/aaaa/vvvv.html
        var curWwwPath = window.document.location.href;
        //获取主机地址之后的目录，如： uimcardprj/share/meun.jsp
        var pathName = window.document.location.pathname;
        var pos = curWwwPath.indexOf(pathName);
        //获取主机地址，如： http://localhost:8888
        var localhostPaht = curWwwPath.substring(0, pos);
        //获取带"/"的项目名，如：/abcd
        var projectName = pathName.substring(0, pathName.substr(1).indexOf('/') + 1);

        // return (localhostPaht + projectName);


        // console.log("当前网址:"+curWwwPath);
        // console.log("主机地址后的目录:"+pos+"----"+pathName);
        // console.log("主机地址:"+localhostPaht);
        // console.log("项目名:"+projectName);


        return projectName;
    }

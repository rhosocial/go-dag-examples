// JavaScript函数，用于与应用交互
function login(username, password) {
    // 调用应用提供的登录接口
    var result = app.login(username, password);
    if (result) {
        console.log("登录成功");
    } else {
        console.log("登录失败");
    }
}

function logout() {
    // 调用应用提供的退出接口
    app.logout();
    console.log("退出成功");
}

function getUserInfo() {
    // 调用应用提供的获取用户信息接口
    var userInfo = app.getUserInfo();
    console.log("用户信息：", userInfo);
}

login("admin", "123456");
getUserInfo();
logout();
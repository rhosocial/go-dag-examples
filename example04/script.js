// JavaScript 需要实现一个接口，其中包含 sayHello 方法
const jsObject = {
    sayHello: function(name) {
        return `Hello, ${name} from JavaScript!`;
    },
    add: function(a, b) {
        return a + b;
    }
};

// 导出函数供 Go 调用
globalThis.getObject = function() {
    return jsObject;
}
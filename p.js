// 先定义三个常量表示状态
var PENDING = 'pending';
var FULFILLED = 'fulfilled';
var REJECTED = 'rejected';

function MyPromise(fn) {
    this.status = PENDING;    // 初始状态为pending
    this.value = null;        // 初始化value
    this.reason = null;       // 初始化reason

    // 构造函数里面添加两个数组存储成功和失败的回调
    this.onFulfilledCallbacks = [];
    this.onRejectedCallbacks = [];

    // 存一下this,以便resolve和reject里面访问
    var that = this;
    // resolve方法参数是value
    function resolve(value) {
        if (that.status === PENDING) {
            that.status = FULFILLED;
            that.value = value;
            // resolve里面将所有成功的回调拿出来执行
            that.onFulfilledCallbacks.forEach(callback => {
                callback(that.value);
            });
        }
    }

    // reject方法参数是reason
    function reject(reason) {
        if (that.status === PENDING) {
            that.status = REJECTED;
            that.reason = reason;

            // resolve里面将所有失败的回调拿出来执行
            that.onRejectedCallbacks.forEach(callback => {
                callback(that.reason);
            });
        }
    }

    try {
        fn(resolve, reject);
    } catch (error) {
        reject(error);
    }
}

function resolvePromise(promise, x, resolve, reject) {
    if (promise === x) {
        return reject(new TypeError('The promise and the return value are the same'));
    }

    if (typeof x === 'object' || typeof x === 'function') {
        if (x === null) {
            return resolve(x);
        }

        try {
            var then = x.then;
        } catch (error) {
            return reject(error);
        }

        if (typeof then === 'function') {
            var called = false;
            try {
                then.call(
                    x,
                    function (y) {
                        if (called) return;
                        called = true;
                        resolvePromise(promise, y, resolve, reject);
                    },
                    function (r) {
                        if (called) return;
                        called = true;
                        reject(r);
                    });
            } catch (error) {
                if (called) return;
                reject(error);
            }
        } else {
            resolve(x);
        }
    } else {
        resolve(x);
    }
}

MyPromise.prototype.then = function (onFulfilled, onRejected) {
    var realOnFulfilled = onFulfilled;
    if (typeof realOnFulfilled !== 'function') {
        realOnFulfilled = function (value) {
            return value;
        }
    }

    var realOnRejected = onRejected;
    if (typeof realOnRejected !== 'function') {
        realOnRejected = function (reason) {
            throw reason;
        }
    }

    var that = this;   // 保存一下this

    if (this.status === FULFILLED) {
        var promise2 = new MyPromise(function (resolve, reject) {
            setTimeout(function () {
                try {
                    var x = realOnFulfilled(that.value);
                    resolvePromise(promise2, x, resolve, reject);
                } catch (error) {
                    reject(error);
                }
            }, 0);
        });

        return promise2;
    }

    if (this.status === REJECTED) {
        var promise2 = new MyPromise(function (resolve, reject) {
            setTimeout(function () {
                try {
                    var x = realOnRejected(that.reason);
                    resolvePromise(promise2, x, resolve, reject);
                } catch (error) {
                    reject(error);
                }
            }, 0);
        });

        return promise2;
    }

    // 如果还是PENDING状态，将回调保存下来
    if (this.status === PENDING) {
        var promise2 = new MyPromise(function (resolve, reject) {
            that.onFulfilledCallbacks.push(function () {
                setTimeout(function () {
                    try {
                        var x = realOnFulfilled(that.value);
                        resolvePromise(promise2, x, resolve, reject);
                    } catch (error) {
                        reject(error);
                    }
                })

            });
            that.onRejectedCallbacks.push(function () {
                setTimeout(function () {
                    try {
                        var x = realOnRejected(that.reason);
                        resolvePromise(promise2, x, resolve, reject);
                    } catch (error) {
                        reject(error);
                    }
                })

            });
        });

        return promise2;
    }
}

MyPromise.deferred = function () {
    var result = {};
    result.promise = new MyPromise(function (resolve, reject) {
        result.resolve = resolve;
        result.reject = reject;
    });

    return result;
}

module.exports = MyPromise;

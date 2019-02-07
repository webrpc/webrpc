"use strict";
/* tslint:disable */
var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    }
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
var __assign = (this && this.__assign) || function () {
    __assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return __assign.apply(this, arguments);
};
Object.defineProperty(exports, "__esModule", { value: true });
var Empty = /** @class */ (function () {
    function Empty(_data) {
        this._data = {};
        if (_data) {
        }
    }
    Empty.prototype.toJSON = function () {
        return this._data;
    };
    return Empty;
}());
exports.Empty = Empty;
var GetUserRequest = /** @class */ (function () {
    function GetUserRequest(_data) {
        this._data = {};
        if (_data) {
            this._data['UserID'] = _data['UserID'];
        }
    }
    Object.defineProperty(GetUserRequest.prototype, "UserID", {
        get: function () {
            return this._data['UserID'];
        },
        set: function (value) {
            this._data['UserID'] = value;
        },
        enumerable: true,
        configurable: true
    });
    GetUserRequest.prototype.toJSON = function () {
        return this._data;
    };
    return GetUserRequest;
}());
exports.GetUserRequest = GetUserRequest;
var User = /** @class */ (function () {
    function User(_data) {
        this._data = {};
        if (_data) {
            this._data['ID'] = _data['ID'];
            this._data['Username'] = _data['Username'];
            this._data['CreatedAt'] = _data['CreatedAt'];
        }
    }
    Object.defineProperty(User.prototype, "ID", {
        get: function () {
            return this._data['ID'];
        },
        set: function (value) {
            this._data['ID'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(User.prototype, "Username", {
        get: function () {
            return this._data['Username'];
        },
        set: function (value) {
            this._data['Username'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(User.prototype, "CreatedAt", {
        get: function () {
            return this._data['CreatedAt'];
        },
        set: function (value) {
            this._data['CreatedAt'] = value;
        },
        enumerable: true,
        configurable: true
    });
    User.prototype.toJSON = function () {
        return this._data;
    };
    return User;
}());
exports.User = User;
var RandomStuff = /** @class */ (function () {
    function RandomStuff(_data) {
        this._data = {};
        if (_data) {
            this._data['Meta'] = _data['Meta'];
            this._data['MetaNestedExample'] = _data['MetaNestedExample'];
            this._data['NamesList'] = _data['NamesList'];
            this._data['NumsList'] = _data['NumsList'];
            this._data['DoubleArray'] = _data['DoubleArray'];
            this._data['ListOfMaps'] = _data['ListOfMaps'];
            this._data['ListOfUsers'] = _data['ListOfUsers'];
            this._data['MapOfUsers'] = _data['MapOfUsers'];
            this._data['User'] = _data['User'];
        }
    }
    Object.defineProperty(RandomStuff.prototype, "Meta", {
        get: function () {
            return this._data['Meta'];
        },
        set: function (value) {
            this._data['Meta'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(RandomStuff.prototype, "MetaNestedExample", {
        get: function () {
            return this._data['MetaNestedExample'];
        },
        set: function (value) {
            this._data['MetaNestedExample'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(RandomStuff.prototype, "NamesList", {
        get: function () {
            return this._data['NamesList'];
        },
        set: function (value) {
            this._data['NamesList'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(RandomStuff.prototype, "NumsList", {
        get: function () {
            return this._data['NumsList'];
        },
        set: function (value) {
            this._data['NumsList'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(RandomStuff.prototype, "DoubleArray", {
        get: function () {
            return this._data['DoubleArray'];
        },
        set: function (value) {
            this._data['DoubleArray'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(RandomStuff.prototype, "ListOfMaps", {
        get: function () {
            return this._data['ListOfMaps'];
        },
        set: function (value) {
            this._data['ListOfMaps'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(RandomStuff.prototype, "ListOfUsers", {
        get: function () {
            return this._data['ListOfUsers'];
        },
        set: function (value) {
            this._data['ListOfUsers'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(RandomStuff.prototype, "MapOfUsers", {
        get: function () {
            return this._data['MapOfUsers'];
        },
        set: function (value) {
            this._data['MapOfUsers'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(RandomStuff.prototype, "User", {
        get: function () {
            return this._data['User'];
        },
        set: function (value) {
            this._data['User'] = value;
        },
        enumerable: true,
        configurable: true
    });
    RandomStuff.prototype.toJSON = function () {
        return this._data;
    };
    return RandomStuff;
}());
exports.RandomStuff = RandomStuff;
// Client
var ExampleServicePathPrefix = "/rpc/ExampleService/";
var ExampleService = /** @class */ (function () {
    function ExampleService(hostname, fetch) {
        this.path = '/rpc/ExampleService/';
        this.hostname = hostname;
        this.fetch = fetch;
    }
    ExampleService.prototype.url = function (name) {
        return this.hostname + this.path + name;
    };
    ExampleService.prototype.Ping = function (headers) {
        return this.fetch(this.url('Ping'), exports.createHTTPRequest({}, headers)).then(function (res) {
            if (!res.ok) {
                return exports.throwHTTPError(res);
            }
            return res.json().then(function (_data) { return (_data); });
        });
    };
    ExampleService.prototype.GetUser = function (params, headers) {
        return this.fetch(this.url('GetUser'), exports.createHTTPRequest(params, headers)).then(function (res) {
            if (!res.ok) {
                return exports.throwHTTPError(res);
            }
            return res.json().then(function (_data) { return new User(_data); });
        });
    };
    return ExampleService;
}());
exports.ExampleService = ExampleService;
var WebRPCError = /** @class */ (function (_super) {
    __extends(WebRPCError, _super);
    function WebRPCError(te) {
        var _this = _super.call(this, te.msg) || this;
        _this.code = te.code;
        _this.meta = te.meta;
        return _this;
    }
    return WebRPCError;
}(Error));
exports.WebRPCError = WebRPCError;
exports.throwHTTPError = function (resp) {
    return resp.json().then(function (err) { throw new WebRPCError(err); });
};
exports.createHTTPRequest = function (body, headers) {
    if (body === void 0) { body = {}; }
    if (headers === void 0) { headers = {}; }
    return {
        method: 'POST',
        headers: __assign({}, headers, { 'Content-Type': 'application/json' }),
        body: JSON.stringify(body || {})
    };
};

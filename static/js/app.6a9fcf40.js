(function (t) {
    function s(s) {
        for (var a, n, r = s[0], l = s[1], c = s[2], d = 0, m = []; d < r.length; d++) n = r[d], Object.prototype.hasOwnProperty.call(o, n) && o[n] && m.push(o[n][0]), o[n] = 0;
        for (a in l) Object.prototype.hasOwnProperty.call(l, a) && (t[a] = l[a]);
        u && u(s);
        while (m.length) m.shift()();
        return i.push.apply(i, c || []), e()
    }

    function e() {
        for (var t, s = 0; s < i.length; s++) {
            for (var e = i[s], a = !0, r = 1; r < e.length; r++) {
                var l = e[r];
                0 !== o[l] && (a = !1)
            }
            a && (i.splice(s--, 1), t = n(n.s = e[0]))
        }
        return t
    }

    var a = {}, o = {app: 0}, i = [];

    function n(s) {
        if (a[s]) return a[s].exports;
        var e = a[s] = {i: s, l: !1, exports: {}};
        return t[s].call(e.exports, e, e.exports, n), e.l = !0, e.exports
    }

    n.m = t, n.c = a, n.d = function (t, s, e) {
        n.o(t, s) || Object.defineProperty(t, s, {enumerable: !0, get: e})
    }, n.r = function (t) {
        "undefined" !== typeof Symbol && Symbol.toStringTag && Object.defineProperty(t, Symbol.toStringTag, {value: "Module"}), Object.defineProperty(t, "__esModule", {value: !0})
    }, n.t = function (t, s) {
        if (1 & s && (t = n(t)), 8 & s) return t;
        if (4 & s && "object" === typeof t && t && t.__esModule) return t;
        var e = Object.create(null);
        if (n.r(e), Object.defineProperty(e, "default", {
            enumerable: !0,
            value: t
        }), 2 & s && "string" != typeof t) for (var a in t) n.d(e, a, function (s) {
            return t[s]
        }.bind(null, a));
        return e
    }, n.n = function (t) {
        var s = t && t.__esModule ? function () {
            return t["default"]
        } : function () {
            return t
        };
        return n.d(s, "a", s), s
    }, n.o = function (t, s) {
        return Object.prototype.hasOwnProperty.call(t, s)
    }, n.p = "/";
    var r = window["webpackJsonp"] = window["webpackJsonp"] || [], l = r.push.bind(r);
    r.push = s, r = r.slice();
    for (var c = 0; c < r.length; c++) s(r[c]);
    var u = l;
    i.push([0, "chunk-vendors"]), e()
})({
    0: function (t, s, e) {
        t.exports = e("56d7")
    }, "02ea": function (t, s, e) {
    }, "0944": function (t, s, e) {
        "use strict";
        e("75cf")
    }, "0a82": function (t, s, e) {
    }, "0c12": function (t, s, e) {
        "use strict";
        e("f167")
    }, "108d": function (t, s, e) {
        "use strict";
        e("4b24")
    }, "18de": function (t, s, e) {
        "use strict";
        e("5c31")
    }, "47f3": function (t, s, e) {
    }, "4b24": function (t, s, e) {
    }, "56d7": function (t, s, e) {
        "use strict";
        e.r(s);
        var a = e("2b0e"), o = function () {
                var t = this, s = t._self._c;
                return s("div", {attrs: {id: "app"}}, [s("div", {staticClass: "page"}, [s("HeadBar"), s("router-view")], 1)])
            }, i = [], n = function () {
                var t = this, s = t._self._c;
                return s("header", {staticClass: "header"}, [s("span", {
                    staticClass: "logo",
                    on: {click: t.goIndex}
                }, [t._v("bluebell")]), t._m(0), s("div", {staticClass: "btns"}, [s("div", {
                    directives: [{
                        name: "show",
                        rawName: "v-show",
                        value: !t.isLogin,
                        expression: "!isLogin"
                    }]
                }, [s("a", {
                    staticClass: "login-btn",
                    on: {click: t.goLogin}
                }, [t._v("登录")]), s("a", {
                    staticClass: "login-btn",
                    on: {click: t.goSignUp}
                }, [t._v("注册")])]), s("div", {
                    directives: [{
                        name: "show",
                        rawName: "v-show",
                        value: t.isLogin,
                        expression: "isLogin"
                    }], staticClass: "user-box"
                }, [s("span", {staticClass: "user"}, [t._v(t._s(t.currUsername))]), s("div", {staticClass: "dropdown-content"}, [s("a", {on: {click: t.goLogout}}, [t._v("登出")])])])])])
            }, r = [function () {
                var t = this, s = t._self._c;
                return s("div", {staticClass: "search"}, [s("label", {staticClass: "s-logo"}), s("input", {
                    staticClass: "s-input",
                    attrs: {type: "text", placeholder: "搜索"}
                })])
            }], l = (e("14d9"), {
                name: "HeadBar", created() {
                    this.$store.commit("init")
                }, computed: {
                    isLogin() {
                        return this.$store.getters.isLogin
                    }, currUsername() {
                        return console.log(this.$store.getters.username), this.$store.getters.username
                    }
                }, methods: {
                    goIndex() {
                        this.$router.push({name: "Home"})
                    }, goLogin() {
                        this.$router.push({name: "Login"})
                    }, goSignUp() {
                        this.$router.push({name: "SignUp"})
                    }, goLogout() {
                        this.$store.commit("logout")
                    }
                }
            }), c = l, u = (e("0944"), e("2877")), d = Object(u["a"])(c, n, r, !1, null, "56de6ff9", null), m = d.exports,
            p = {components: {HeadBar: m}}, h = p, v = (e("18de"), Object(u["a"])(h, o, i, !1, null, null, null)),
            g = v.exports, f = e("8c4f"), C = function () {
                var t = this, s = t._self._c;
                return s("div", {staticClass: "satic-area"}, [s("div", {staticClass: "dynamic-area1"}), s("div", {staticClass: "dynamic-area2"}), s("div", {staticClass: "content"}, [t._m(0), s("div", {staticClass: "left"}, [s("div", {staticClass: "c-l-header"}, [s("div", {
                    staticClass: "new btn-iconfont",
                    class: {active: t.timeOrder},
                    on: {
                        click: function (s) {
                            return t.selectOrder("time")
                        }
                    }
                }, [s("i", {staticClass: "iconfont icon-polygonred"}), t._v("最新发布 ")]), s("div", {
                    staticClass: "top btn-iconfont",
                    class: {active: t.scoreOrder},
                    on: {
                        click: function (s) {
                            return t.selectOrder("score")
                        }
                    }
                }, [s("i", {staticClass: "iconfont icon-top"}), t._v("热度最高 ")]), s("button", {
                    staticClass: "btn-publish",
                    on: {click: t.goPublish}
                }, [t._v("发布新帖")])]), s("ul", {staticClass: "c-l-list"}, t._l(t.postList, (function (e) {
                    return s("li", {
                        key: e.id,
                        staticClass: "c-l-item"
                    }, [s("div", {staticClass: "post"}, [s("a", {staticClass: "vote"}, [s("span", {
                        staticClass: "iconfont icon-up",
                        on: {
                            click: function (s) {
                                return t.vote(e.id, "1")
                            }
                        }
                    })]), s("span", {staticClass: "text"}, [t._v(t._s(e.vote_num))]), s("a", {staticClass: "vote"}, [s("span", {
                        staticClass: "iconfont icon-down",
                        on: {
                            click: function (s) {
                                return t.vote(e.id, "-1")
                            }
                        }
                    })])]), s("div", {
                        staticClass: "l-container", on: {
                            click: function (s) {
                                return t.goDetail(e.id)
                            }
                        }
                    }, [s("h4", {staticClass: "con-title"}, [t._v(t._s(e.title))]), s("div", {staticClass: "con-memo"}, [s("p", [t._v(t._s(e.content.slice(0, 20)) + "...")])])])])
                })), 0)]), s("div", {staticClass: "right"}, [s("div", {staticClass: "communities"}, [s("h2", {staticClass: "r-c-title"}, [t._v("今日热门社区")]), s("ul", {staticClass: "r-c-content"}, t._l(t.hotChannels, (function (e) {
                    return s("li", {
                        key: e.id,
                        staticClass: "r-c-item"
                    }, [s("span", {staticClass: "index"}, [t._v(t._s(e.rank))]), s("div", {staticClass: "icon"}), s("span", {staticClass: "channel-name"}, [t._v(t._s(e.name))])])
                })), 0), s("button", {staticClass: "view-all"}, [t._v("查看所有")])]), s("div", {staticClass: "r-trending"}, [s("h2", {staticClass: "r-t-title"}, [t._v("所有社区")]), s("ul", {staticClass: "rank"}, t._l(t.trendingChannels, (function (e) {
                    return s("li", {
                        key: e.id,
                        staticClass: "r-t-cell"
                    }, [s("div", {staticClass: "r-t-cell-info"}, [s("div", {staticClass: "avatar"}), s("div", {staticClass: "info"}, [s("span", {staticClass: "info-title"}, [t._v(t._s(e.name))]), s("p", {staticClass: "info-num"}, [t._v(t._s(e.members) + " members")])])]), s("button", {staticClass: "join-btn"}, [t._v("JOIN")])])
                })), 0)])])])])
            }, _ = [function () {
                var t = this, s = t._self._c;
                return s("div", {staticClass: "sidebar-container"}, [s("div", {staticClass: "left-sidebar"}, [s("h1", [t._v("热门网站")]), s("ul", [s("li", [s("a", {
                    attrs: {
                        href: "https://google.com",
                        target: "_blank"
                    }
                }, [t._v("谷歌主页")])]), s("li", [s("a", {
                    attrs: {
                        href: "https://github.com/",
                        target: "_blank"
                    }
                }, [t._v("github")])]), s("li", [s("a", {
                    attrs: {
                        href: "https://poe.com/",
                        target: "_blank"
                    }
                }, [t._v("poe:AI模型")])]), s("li", [s("a", {
                    attrs: {
                        href: "https://leetcode.cn/problemset/",
                        target: "_blank"
                    }
                }, [t._v("力扣")])])])]), s("div", {staticClass: "right-sidebar"}, [s("h1", [t._v("学习资料")]), s("ul", [s("li", [s("a", {attrs: {href: "https://hit-alibaba.github.io/interview/basic/network/TCP.html"}}, [t._v("计网面经")])]), s("li", [s("a", {attrs: {href: "https://wx.zsxq.com/dweb2/index/group/51122858222824"}}, [t._v("知识星球")])]), s("li", [s("a", {attrs: {href: "https://hellogithub.com/"}}, [t._v("HelloGitHub:开源项目分享")])]), s("li", [s("a", {attrs: {href: "https://pdai.tech/md/algorithm/alg-domain-id-snowflake.html"}}, [t._v("Java知识体系")])])])])])
            }], b = {
                name: "Home", data() {
                    return {order: "time", page: 1, postList: [], hotChannels: [], trendingChannels: []}
                }, methods: {
                    selectOrder(t) {
                        this.order = t, this.loadData()
                    }, goPublish() {
                        this.$router.push({name: "Publish"})
                    }, goDetail(t) {
                        this.$router.push({name: "Content", params: {id: t}})
                    }, loadData() {
                        this.getPostList(), this.getHotChannels(), this.getTrendingChannels()
                    }, getPostList() {
                        this.$axios({
                            method: "get",
                            url: "/posts2",
                            params: {page: this.page, order: this.order}
                        }).then(t => {
                            1e3 === t.data.code ? this.postList = t.data.data : console.log(t.data.msg)
                        }).catch(t => {
                            console.log(t)
                        })
                    }, getHotChannels() {
                        this.$axios({method: "get", url: "/community/:id", params: {order: this.order}}).then(t => {
                            1e3 === t.data.code ? this.hotChannels = t.data.data : console.log(t.data.msg)
                        }).catch(t => {
                            console.log(t)
                        })
                    }, getTrendingChannels() {
                        this.$axios({method: "get", url: "/community", params: {order: this.order}}).then(t => {
                            1e3 === t.data.code ? this.trendingChannels = t.data.data : console.log(t.data.msg)
                        }).catch(t => {
                            console.log(t)
                        })
                    }, vote(t, s) {
                        this.$axios({
                            method: "post",
                            url: "/vote",
                            data: JSON.stringify({post_id: t, direction: s})
                        }).then(t => {
                            1e3 === t.data.code ? console.log("vote success") : console.log(t.data.msg)
                        }).catch(t => {
                            console.log(t)
                        })
                    }
                }, mounted() {
                    this.loadData()
                }, computed: {
                    timeOrder() {
                        return "time" === this.order
                    }, scoreOrder() {
                        return "score" === this.order
                    }
                }
            }, y = b, w = (e("108d"), Object(u["a"])(y, C, _, !1, null, "744650e4", null)), x = w.exports, k = function () {
                var t = this, s = t._self._c;
                return s("div", {staticClass: "content"}, [s("div", {staticClass: "left"}, [s("div", {staticClass: "container"}, [t._m(0), s("div", {staticClass: "l-container"}, [s("h4", {staticClass: "con-title"}, [t._v(t._s(t.post.title))]), s("div", {staticClass: "con-info"}, [t._v(t._s(t.post.content))]), t._m(1)])])]), s("div", {staticClass: "right"}, [s("div", {staticClass: "topic-info"}, [s("h5", {staticClass: "t-header"}), s("div", {staticClass: "t-info"}, [s("a", {staticClass: "avatar"}), s("span", {staticClass: "topic-name"}, [t._v("b/" + t._s(t.post.community_name))])]), s("p", {staticClass: "t-desc"}, [t._v("树洞 树洞 无限树洞的树洞")]), t._m(2), s("div", {staticClass: "date"}, [t._v("Created Apr 10, 2008")]), s("button", {staticClass: "topic-btn"}, [t._v("JOIN")])])])])
            }, L = [function () {
                var t = this, s = t._self._c;
                return s("div", {staticClass: "post"}, [s("a", {staticClass: "vote"}, [s("span", {staticClass: "iconfont icon-up"})]), s("span", {staticClass: "text"}, [t._v("50.2k")]), s("a", {staticClass: "vote"}, [s("span", {staticClass: "iconfont icon-down"})])])
            }, function () {
                var t = this, s = t._self._c;
                return s("div", {staticClass: "user-btn"}, [s("span", {staticClass: "btn-item"}, [s("i", {staticClass: "iconfont icon-comment"}), t._v("comment ")])])
            }, function () {
                var t = this, s = t._self._c;
                return s("ul", {staticClass: "t-num"}, [s("li", {staticClass: "t-num-item"}, [s("p", {staticClass: "number"}, [t._v("5.2m")]), s("span", {staticClass: "unit"}, [t._v("Members")])]), s("li", {staticClass: "t-num-item"}, [s("p", {staticClass: "number"}, [t._v("5.2m")]), s("span", {staticClass: "unit"}, [t._v("Members")])])])
            }], O = {
                name: "Content", data() {
                    return {post: {}}
                }, methods: {
                    getPostDetail() {
                        this.$axios({method: "get", url: "/post/" + this.$route.params.id}).then(t => {
                            console.log(1, t.data), 1e3 == t.code ? this.post = t.data : console.log(t.msg)
                        }).catch(t => {
                            console.log(t)
                        })
                    }
                }, mounted: function () {
                    this.getPostDetail()
                }
            }, P = O, S = (e("0c12"), Object(u["a"])(P, k, L, !1, null, "090b76ef", null)), $ = S.exports, j = function () {
                var t = this, s = t._self._c;
                return s("div", {staticClass: "content"}, [s("div", {staticClass: "left"}, [s("div", {staticClass: "post-name"}, [t._v("我好想写点什么")]), s("div", {staticClass: "post-type"}, [s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.selectCommunity.name,
                        expression: "selectCommunity.name"
                    }],
                    staticClass: "post-type-value",
                    attrs: {type: "text", placeholder: "选择一个频道"},
                    domProps: {value: t.selectCommunity.name},
                    on: {
                        click: function (s) {
                            return t.showCommunity()
                        }, input: function (s) {
                            s.target.composing || t.$set(t.selectCommunity, "name", s.target.value)
                        }
                    }
                }), s("ul", {
                    directives: [{
                        name: "show",
                        rawName: "v-show",
                        value: t.showCommunityList,
                        expression: "showCommunityList"
                    }], staticClass: "post-type-options"
                }, t._l(t.communityList, (function (e, a) {
                    return s("li", {
                        key: e.id, staticClass: "post-type-cell", on: {
                            click: function (s) {
                                return t.selected(a)
                            }
                        }
                    }, [t._v(" " + t._s(e.name) + " ")])
                })), 0), s("i", {staticClass: "p-icon"})]), s("div", {staticClass: "post-content"}, [t._m(0), s("div", {staticClass: "post-sub-container"}, [s("div", {staticClass: "post-sub-header"}, [s("textarea", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.title,
                        expression: "title"
                    }],
                    staticClass: "post-title",
                    attrs: {id: "", cols: "30", rows: "10", placeholder: "标题"},
                    domProps: {value: t.title},
                    on: {
                        input: function (s) {
                            s.target.composing || (t.title = s.target.value)
                        }
                    }
                }), s("span", {staticClass: "textarea-num"}, [t._v("0/300")])]), s("div", {staticClass: "post-text-con"}, [s("textarea", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.content,
                        expression: "content"
                    }],
                    staticClass: "post-content-t",
                    attrs: {id: "", cols: "30", rows: "10", placeholder: "内容"},
                    domProps: {value: t.content},
                    on: {
                        input: function (s) {
                            s.target.composing || (t.content = s.target.value)
                        }
                    }
                })])]), s("div", {staticClass: "post-footer"}, [s("div", {staticClass: "btns"}, [s("button", {staticClass: "btn"}, [t._v("取消")]), s("button", {
                    staticClass: "btn",
                    on: {
                        click: function (s) {
                            return t.submit()
                        }
                    }
                }, [t._v("发表")])])])])]), t._m(1)])
            }, N = [function () {
                var t = this, s = t._self._c;
                return s("ul", {staticClass: "cat"}, [s("li", {staticClass: "cat-item active"}, [s("i", {staticClass: "iconfont icon-edit"}), t._v("post ")]), s("li", {staticClass: "cat-item"}, [s("i", {staticClass: "iconfont icon-image"}), t._v("image/video ")])])
            }, function () {
                var t = this, s = t._self._c;
                return s("div", {staticClass: "right"}, [s("div", {staticClass: "post-rank"}, [s("h5", {staticClass: "p-r-title"}, [s("i", {staticClass: "p-r-icon"}), t._v("发帖规范 ")]), s("ul", {staticClass: "p-r-content"}, [s("li", {staticClass: "p-r-item"}, [t._v("1.网络不是法外之地")]), s("li", {staticClass: "p-r-item"}, [t._v("2.网络不是法外之地")]), s("li", {staticClass: "p-r-item"}, [t._v("3.网络不是法外之地")])])])])
            }], M = {
                name: "Publish", data() {
                    return {title: "", content: "", showCommunityList: !1, selectCommunity: {}, communityList: []}
                }, methods: {
                    submit() {
                        this.$axios({
                            method: "post",
                            url: "/post",
                            data: JSON.stringify({
                                title: this.title,
                                content: this.content,
                                community_id: this.selectCommunity.id
                            })
                        }).then(t => {
                            console.log(t.data), 1e3 == t.code ? this.$router.push({path: this.redirect || "/"}) : console.log(t.msg)
                        }).catch(t => {
                            console.log(t)
                        })
                    }, getCommunityList() {
                        this.$axios({method: "get", url: "/community"}).then(t => {
                            console.log(t.data), 1e3 == t.code ? this.communityList = t.data : console.log(t.msg)
                        }).catch(t => {
                            console.log(t)
                        })
                    }, showCommunity() {
                        this.showCommunityList = !this.showCommunityList
                    }, selected(t) {
                        this.selectCommunity = this.communityList[t], this.showCommunityList = !1, console.log(this.selectCommunity)
                    }
                }, mounted: function () {
                    this.getCommunityList()
                }
            }, R = M, J = (e("af7a"), Object(u["a"])(R, j, N, !1, null, "6236a644", null)), I = J.exports, H = function () {
                var t = this, s = t._self._c;
                return s("div", {staticClass: "main"}, [t._m(0), s("div", {staticClass: "container"}, [s("h2", {staticClass: "form-title"}, [t._v("登录")]), s("div", {staticClass: "form-group login-method"}, [s("button", {
                    class: {active: "username" === t.loginMethod},
                    on: {
                        click: function (s) {
                            return t.setLoginMethod("username")
                        }
                    }
                }, [t._v("用户名登录")]), s("button", {
                    class: {active: "verification" === t.loginMethod},
                    on: {
                        click: function (s) {
                            return t.setLoginMethod("verification")
                        }
                    }
                }, [t._v("验证码登录 ")])]), "username" === t.loginMethod ? s("div", [s("div", {staticClass: "form-group margin-adjust"}, [s("label", {attrs: {for: "name"}}, [t._v("用户名")]), s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.username,
                        expression: "username"
                    }],
                    staticClass: "form-control",
                    attrs: {type: "text", name: "name", id: "name", placeholder: "用户名"},
                    domProps: {value: t.username},
                    on: {
                        input: function (s) {
                            s.target.composing || (t.username = s.target.value)
                        }
                    }
                })]), s("div", {staticClass: "form-group margin-adjust"}, [s("label", {attrs: {for: "pass"}}, [t._v("密码")]), s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.password,
                        expression: "password"
                    }],
                    staticClass: "form-control",
                    attrs: {type: "password", name: "pass", id: "pass", placeholder: "密码"},
                    domProps: {value: t.password},
                    on: {
                        input: function (s) {
                            s.target.composing || (t.password = s.target.value)
                        }
                    }
                })])]) : t._e(), "verification" === t.loginMethod ? s("div", [s("div", {staticClass: "form-group margin-adjust"}, [s("label", {attrs: {for: "phone"}}, [t._v("手机号")]), s("div", {staticClass: "verification-group"}, [s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.phone,
                        expression: "phone"
                    }],
                    staticClass: "form-control",
                    attrs: {type: "text", name: "phone", id: "phone", placeholder: "手机号"},
                    domProps: {value: t.phone},
                    on: {
                        input: function (s) {
                            s.target.composing || (t.phone = s.target.value)
                        }
                    }
                }), s("button", {
                    staticClass: "btn-verification",
                    on: {click: t.sendVerificationCode}
                }, [t._v("发送验证码")])])]), s("div", {staticClass: "form-group margin-adjust"}, [s("label", {attrs: {for: "code"}}, [t._v("验证码")]), s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.verificationCode,
                        expression: "verificationCode"
                    }],
                    staticClass: "form-control",
                    attrs: {type: "text", name: "code", id: "code", placeholder: "验证码"},
                    domProps: {value: t.verificationCode},
                    on: {
                        input: function (s) {
                            s.target.composing || (t.verificationCode = s.target.value)
                        }
                    }
                })])]) : t._e(), s("div", {staticClass: "form-btn"}, [s("button", {
                    staticClass: "btn btn-info",
                    attrs: {type: "button"},
                    on: {click: t.submit}
                }, [t._v("登录")])])]), t._m(1)])
            }, U = [function () {
                var t = this, s = t._self._c;
                return s("div", {staticClass: "header"}, [s("h1", [t._v("BlueBell")]), s("p", [t._v("风铃草论坛，发现你喜欢")])])
            }, function () {
                var t = this, s = t._self._c;
                return s("footer", [s("div", {staticClass: "footer-content"}, [s("span", {staticClass: "footer-info"}, [t._v("开发组：软件三班第四组")]), s("span", {staticClass: "footer-info"}, [t._v("联系方式：6666666")]), s("span", {staticClass: "footer-info"}, [t._v("版权所有 © 2024 XUPT")])])])
            }], T = {
                name: "Login", data() {
                    return {
                        loginMethod: "username",
                        username: "",
                        password: "",
                        phone: "",
                        verificationCode: "",
                        submitted: !1
                    }
                }, methods: {
                    setLoginMethod(t) {
                        this.loginMethod = t
                    }, submit() {
                        "username" === this.loginMethod ? this.$axios({
                            method: "post",
                            url: "/login",
                            data: JSON.stringify({username: this.username, password: this.password})
                        }).then(t => {
                            console.log(t.data), 1e3 === t.code ? (localStorage.setItem("loginResult", JSON.stringify(t.data)), this.$store.commit("login", t.data), this.$router.push({path: this.redirect || "/"})) : console.log(t.message)
                        }).catch(t => {
                            console.log(t)
                        }) : "verification" === this.loginMethod && this.$axios({
                            method: "post",
                            url: "/login",
                            data: JSON.stringify({phone: this.phone, verificationCode: this.verificationCode})
                        }).then(t => {
                            1e3 === t.code ? (localStorage.setItem("loginResult", JSON.stringify(t.data)), this.$store.commit("login", t.data), this.$router.push({path: this.redirect || "/"})) : console.log(t.message)
                        }).catch(t => {
                            console.log(t)
                        })
                    }, sendVerificationCode() {
                        this.$axios({
                            method: "post",
                            url: "/loginSMS",
                            data: JSON.stringify({phone: this.phone})
                        }).then(t => {
                            console.log(t.data)
                        }).catch(t => {
                            console.log(t)
                        })
                    }
                }
            }, B = T, D = (e("b605"), Object(u["a"])(B, H, U, !1, null, "6145479d", null)), A = D.exports, q = function () {
                var t = this, s = t._self._c;
                return s("div", {staticClass: "main"}, [t._m(0), s("div", {staticClass: "container"}, [s("h2", {staticClass: "form-title"}, [t._v("注册")]), s("div", {staticClass: "form-group margin-adjust"}, [s("label", {attrs: {for: "name"}}, [t._v("用户名")]), s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.username,
                        expression: "username"
                    }],
                    staticClass: "form-control",
                    attrs: {type: "text", name: "name", id: "name", placeholder: "用户名"},
                    domProps: {value: t.username},
                    on: {
                        input: function (s) {
                            s.target.composing || (t.username = s.target.value)
                        }
                    }
                })]), s("div", {staticClass: "form-group margin-adjust"}, [s("label", {attrs: {for: "pass"}}, [t._v("密码")]), s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.password,
                        expression: "password"
                    }],
                    staticClass: "form-control",
                    attrs: {type: "password", name: "pass", id: "pass", placeholder: "密码"},
                    domProps: {value: t.password},
                    on: {
                        input: function (s) {
                            s.target.composing || (t.password = s.target.value)
                        }
                    }
                })]), s("div", {staticClass: "form-group margin-adjust"}, [s("label", {attrs: {for: "re_pass"}}, [t._v("确认密码")]), s("input", {
                    directives: [{
                        name: "model",
                        rawName: "v-model",
                        value: t.confirm_password,
                        expression: "confirm_password"
                    }],
                    staticClass: "form-control",
                    attrs: {type: "password", name: "re_pass", id: "re_pass", placeholder: "确认密码"},
                    domProps: {value: t.confirm_password},
                    on: {
                        input: function (s) {
                            s.target.composing || (t.confirm_password = s.target.value)
                        }
                    }
                })]), s("div", {staticClass: "form-btn"}, [s("button", {
                    staticClass: "btn btn-info",
                    attrs: {type: "button"},
                    on: {click: t.submit}
                }, [t._v("注册")])])]), t._m(1)])
            }, z = [function () {
                var t = this, s = t._self._c;
                return s("div", {staticClass: "header"}, [s("h1", [t._v("BlueBell")]), s("p", [t._v("风铃草论坛，发现你喜欢")])])
            }, function () {
                var t = this, s = t._self._c;
                return s("footer", [s("div", {staticClass: "footer-content"}, [s("span", {staticClass: "footer-info"}, [t._v("开发组：软件三班第四组")]), s("span", {staticClass: "footer-info"}, [t._v("联系方式：6666666")]), s("span", {staticClass: "footer-info"}, [t._v("版权所有 © 2024 XUPT")])])])
            }], V = {
                name: "SignUp", data() {
                    return {username: "", password: "", confirm_password: "", submitted: !1}
                }, computed: {}, created() {
                }, methods: {
                    submit() {
                        this.$axios({
                            method: "post",
                            url: "/signup",
                            data: JSON.stringify({
                                username: this.username,
                                password: this.password,
                                confirm_password: this.confirm_password
                            })
                        }).then(t => {
                            console.log(t.data), 1e3 === t.code ? (console.log("signup success"), this.$router.push({name: "Login"})) : console.log(t.msg)
                        }).catch(t => {
                            console.log(t)
                        })
                    }
                }
            }, X = V, E = (e("df34"), Object(u["a"])(X, q, z, !1, null, "62117e1e", null)), G = E.exports;
        const F = f["a"].prototype.push;
        f["a"].prototype.push = function (t) {
            return F.call(this, t).catch(t => t)
        }, a["a"].use(f["a"]);
        const K = [{path: "/", name: "Home", component: x}, {
            path: "/post/:id",
            name: "Content",
            component: $
        }, {path: "/publish", name: "Publish", component: I, meta: {requireAuth: !0}}, {
            path: "/login",
            name: "Login",
            component: A
        }, {path: "/signup", name: "SignUp", component: G}], Q = new f["a"]({mode: "history", base: "/", routes: K});
        var W = Q, Y = e("2f62");
        a["a"].use(Y["a"]);
        const Z = {token: null, user_id: null, user_name: null};
        var tt = new Y["a"].Store({
            state: {isLogin: !1, loginResult: Z},
            mutations: {
                init(t) {
                    let s = JSON.parse(localStorage.getItem("loginResult"));
                    console.log(localStorage.getItem("loginResult")), null != s && (t.loginResult = s)
                }, login(t, s) {
                    t.loginResult = s
                }, logout(t) {
                    localStorage.removeItem("loginResult"), t.loginResult = Z
                }
            },
            actions: {},
            getters: {
                isLogin: t => null !== t.loginResult.user_id,
                userID: t => t.loginResult.user_id,
                username: t => t.loginResult.user_name,
                accessToken: t => t.loginResult.token
            }
        }), st = e("bc3a"), et = e.n(st);
        et.a.defaults.baseURL = "/api/v1/", et.a.interceptors.request.use(t => {
            let s = JSON.parse(localStorage.getItem("loginResult"));
            if (s) {
                const e = s.token;
                t.headers.Authorization = "Bearer " + e
            }
            return t
        }, t => Promise.reject(t)), et.a.interceptors.response.use(t => 200 === t.status ? Promise.resolve(t.data) : Promise.reject(t.data), t => {
            console.log("error", t)
        });
        var at = et.a;
        a["a"].prototype.$axios = at, a["a"].config.productionTip = !1, W.beforeEach((t, s, e) => {
            console.log(t), console.log(s), t.meta.requireAuth ? localStorage.getItem("loginResult") || "/login" === t.path ? e() : e({path: "/login"}) : e(), "/login" == t.fullPath && (localStorage.getItem("loginResult") ? e({path: s.fullPath}) : e())
        }), new a["a"]({router: W, store: tt, render: t => t(g)}).$mount("#app")
    }, "5c31": function (t, s, e) {
    }, "75cf": function (t, s, e) {
    }, af7a: function (t, s, e) {
        "use strict";
        e("02ea")
    }, b605: function (t, s, e) {
        "use strict";
        e("0a82")
    }, df34: function (t, s, e) {
        "use strict";
        e("47f3")
    }, f167: function (t, s, e) {
    }
});
//# sourceMappingURL=app.6a9fcf40.js.map
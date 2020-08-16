(window["webpackJsonp"] = window["webpackJsonp"] || []).push([["./assets/js/dashboard"], {

    /***/
    "./node_modules/babel-loader/lib/index.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/components/Dashboard.vue?vue&type=script&lang=js&":
    /*!***************************************************************************************************************************************************************************!*\
      !*** ./node_modules/babel-loader/lib??ref--4-0!./node_modules/vue-loader/lib??vue-loader-options!./resources/assets/js/components/Dashboard.vue?vue&type=script&lang=js& ***!
      \***************************************************************************************************************************************************************************/
    /*! exports provided: default */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var vuex__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! vuex */ "./node_modules/vuex/dist/vuex.esm.js");
        /* harmony import */
        var _Sidebar__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Sidebar */ "./resources/assets/js/components/Sidebar.vue");
        /* harmony import */
        var _Navbar__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./Navbar */ "./resources/assets/js/components/Navbar.vue");
        /* harmony import */
        var _Footer__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./Footer */ "./resources/assets/js/components/Footer.vue");

        function ownKeys(object, enumerableOnly) {
            var keys = Object.keys(object);
            if (Object.getOwnPropertySymbols) {
                var symbols = Object.getOwnPropertySymbols(object);
                if (enumerableOnly) symbols = symbols.filter(function (sym) {
                    return Object.getOwnPropertyDescriptor(object, sym).enumerable;
                });
                keys.push.apply(keys, symbols);
            }
            return keys;
        }

        function _objectSpread(target) {
            for (var i = 1; i < arguments.length; i++) {
                var source = arguments[i] != null ? arguments[i] : {};
                if (i % 2) {
                    ownKeys(Object(source), true).forEach(function (key) {
                        _defineProperty(target, key, source[key]);
                    });
                } else if (Object.getOwnPropertyDescriptors) {
                    Object.defineProperties(target, Object.getOwnPropertyDescriptors(source));
                } else {
                    ownKeys(Object(source)).forEach(function (key) {
                        Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key));
                    });
                }
            }
            return target;
        }

        function _defineProperty(obj, key, value) {
            if (key in obj) {
                Object.defineProperty(obj, key, {value: value, enumerable: true, configurable: true, writable: true});
            } else {
                obj[key] = value;
            }
            return obj;
        }

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//


        /* harmony default export */
        __webpack_exports__["default"] = ({
            name: 'Dashboard',
            computed: _objectSpread({}, Object(vuex__WEBPACK_IMPORTED_MODULE_0__["mapState"])(['sideBarOpen'])),
            components: {
                Sidebar: _Sidebar__WEBPACK_IMPORTED_MODULE_1__["default"],
                Navbar: _Navbar__WEBPACK_IMPORTED_MODULE_2__["default"],
                Footer: _Footer__WEBPACK_IMPORTED_MODULE_3__["default"]
            }
        });

        /***/
    }),

    /***/
    "./node_modules/babel-loader/lib/index.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/components/Footer.vue?vue&type=script&lang=js&":
    /*!************************************************************************************************************************************************************************!*\
      !*** ./node_modules/babel-loader/lib??ref--4-0!./node_modules/vue-loader/lib??vue-loader-options!./resources/assets/js/components/Footer.vue?vue&type=script&lang=js& ***!
      \************************************************************************************************************************************************************************/
    /*! exports provided: default */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
//
//
//
//
//
//
//
//
//
//
//
//
        /* harmony default export */
        __webpack_exports__["default"] = ({
            name: 'Footer'
        });

        /***/
    }),

    /***/
    "./node_modules/babel-loader/lib/index.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/components/Navbar.vue?vue&type=script&lang=js&":
    /*!************************************************************************************************************************************************************************!*\
      !*** ./node_modules/babel-loader/lib??ref--4-0!./node_modules/vue-loader/lib??vue-loader-options!./resources/assets/js/components/Navbar.vue?vue&type=script&lang=js& ***!
      \************************************************************************************************************************************************************************/
    /*! exports provided: default */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var vuex__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! vuex */ "./node_modules/vuex/dist/vuex.esm.js");

        function ownKeys(object, enumerableOnly) {
            var keys = Object.keys(object);
            if (Object.getOwnPropertySymbols) {
                var symbols = Object.getOwnPropertySymbols(object);
                if (enumerableOnly) symbols = symbols.filter(function (sym) {
                    return Object.getOwnPropertyDescriptor(object, sym).enumerable;
                });
                keys.push.apply(keys, symbols);
            }
            return keys;
        }

        function _objectSpread(target) {
            for (var i = 1; i < arguments.length; i++) {
                var source = arguments[i] != null ? arguments[i] : {};
                if (i % 2) {
                    ownKeys(Object(source), true).forEach(function (key) {
                        _defineProperty(target, key, source[key]);
                    });
                } else if (Object.getOwnPropertyDescriptors) {
                    Object.defineProperties(target, Object.getOwnPropertyDescriptors(source));
                } else {
                    ownKeys(Object(source)).forEach(function (key) {
                        Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key));
                    });
                }
            }
            return target;
        }

        function _defineProperty(obj, key, value) {
            if (key in obj) {
                Object.defineProperty(obj, key, {value: value, enumerable: true, configurable: true, writable: true});
            } else {
                obj[key] = value;
            }
            return obj;
        }

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

        /* harmony default export */
        __webpack_exports__["default"] = ({
            name: 'Navbar',
            computed: _objectSpread({}, Object(vuex__WEBPACK_IMPORTED_MODULE_0__["mapState"])(['sideBarOpen'])),
            data: function data() {
                return {
                    dropDownOpen: false
                };
            },
            methods: {
                toggleSidebar: function toggleSidebar() {
                    this.$store.dispatch('toggleSidebar');
                }
            }
        });

        /***/
    }),

    /***/
    "./node_modules/babel-loader/lib/index.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/components/Sidebar.vue?vue&type=script&lang=js&":
    /*!*************************************************************************************************************************************************************************!*\
      !*** ./node_modules/babel-loader/lib??ref--4-0!./node_modules/vue-loader/lib??vue-loader-options!./resources/assets/js/components/Sidebar.vue?vue&type=script&lang=js& ***!
      \*************************************************************************************************************************************************************************/
    /*! exports provided: default */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var vuex__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! vuex */ "./node_modules/vuex/dist/vuex.esm.js");

        function ownKeys(object, enumerableOnly) {
            var keys = Object.keys(object);
            if (Object.getOwnPropertySymbols) {
                var symbols = Object.getOwnPropertySymbols(object);
                if (enumerableOnly) symbols = symbols.filter(function (sym) {
                    return Object.getOwnPropertyDescriptor(object, sym).enumerable;
                });
                keys.push.apply(keys, symbols);
            }
            return keys;
        }

        function _objectSpread(target) {
            for (var i = 1; i < arguments.length; i++) {
                var source = arguments[i] != null ? arguments[i] : {};
                if (i % 2) {
                    ownKeys(Object(source), true).forEach(function (key) {
                        _defineProperty(target, key, source[key]);
                    });
                } else if (Object.getOwnPropertyDescriptors) {
                    Object.defineProperties(target, Object.getOwnPropertyDescriptors(source));
                } else {
                    ownKeys(Object(source)).forEach(function (key) {
                        Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key));
                    });
                }
            }
            return target;
        }

        function _defineProperty(obj, key, value) {
            if (key in obj) {
                Object.defineProperty(obj, key, {value: value, enumerable: true, configurable: true, writable: true});
            } else {
                obj[key] = value;
            }
            return obj;
        }

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

        /* harmony default export */
        __webpack_exports__["default"] = ({
            name: "Sidebar",
            computed: _objectSpread({}, Object(vuex__WEBPACK_IMPORTED_MODULE_0__["mapState"])(["sideBarOpen"]))
        });

        /***/
    }),

    /***/
    "./node_modules/babel-loader/lib/index.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/pages/Home.vue?vue&type=script&lang=js&":
    /*!*****************************************************************************************************************************************************************!*\
      !*** ./node_modules/babel-loader/lib??ref--4-0!./node_modules/vue-loader/lib??vue-loader-options!./resources/assets/js/pages/Home.vue?vue&type=script&lang=js& ***!
      \*****************************************************************************************************************************************************************/
    /*! exports provided: default */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var chart_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! chart.js */ "./node_modules/chart.js/dist/Chart.js");
        /* harmony import */
        var chart_js__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(chart_js__WEBPACK_IMPORTED_MODULE_0__);
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

        /* harmony default export */
        __webpack_exports__["default"] = ({
            name: 'DashboardHome',
            data: function data() {
                return {
                    buyersData: {
                        type: 'line',
                        data: {
                            labels: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
                            datasets: [{
                                backgroundColor: "rgba(99,179,237,0.4)",
                                strokeColor: "#63b3ed",
                                pointColor: "#fff",
                                pointStrokeColor: "#63b3ed",
                                data: [203, 156, 99, 251, 305, 247, 256]
                            }, {
                                backgroundColor: "rgba(198,198,198,0.4)",
                                strokeColor: "#f7fafc",
                                pointColor: "#fff",
                                pointStrokeColor: "#f7fafc",
                                data: [86, 97, 144, 114, 94, 108, 156]
                            }]
                        },
                        options: {
                            legend: {
                                display: false
                            },
                            scales: {
                                yAxes: [{
                                    gridLines: {
                                        display: false
                                    },
                                    ticks: {
                                        display: false
                                    }
                                }],
                                xAxes: [{
                                    gridLines: {
                                        display: false
                                    }
                                }]
                            }
                        }
                    },
                    reviewsData: {
                        type: 'bar',
                        data: {
                            labels: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
                            datasets: [{
                                backgroundColor: "rgba(99,179,237,0.4)",
                                strokeColor: "#63b3ed",
                                pointColor: "#fff",
                                pointStrokeColor: "#63b3ed",
                                data: [203, 156, 99, 251, 305, 247, 256]
                            }]
                        },
                        options: {
                            legend: {
                                display: false
                            },
                            scales: {
                                yAxes: [{
                                    gridLines: {
                                        display: false
                                    },
                                    ticks: {
                                        display: false
                                    }
                                }],
                                xAxes: [{
                                    gridLines: {
                                        display: false
                                    }
                                }]
                            }
                        }
                    }
                };
            },
            mounted: function mounted() {
                new chart_js__WEBPACK_IMPORTED_MODULE_0___default.a(document.getElementById('buyers-chart'), this.buyersData);
                new chart_js__WEBPACK_IMPORTED_MODULE_0___default.a(document.getElementById('reviews-chart'), this.reviewsData);
            }
        });

        /***/
    }),

    /***/
    "./node_modules/moment/locale sync recursive ^\\.\\/.*$":
    /*!**************************************************!*\
      !*** ./node_modules/moment/locale sync ^\.\/.*$ ***!
      \**************************************************/
    /*! no static exports found */
    /***/ (function (module, exports, __webpack_require__) {

        var map = {
            "./af": "./node_modules/moment/locale/af.js",
            "./af.js": "./node_modules/moment/locale/af.js",
            "./ar": "./node_modules/moment/locale/ar.js",
            "./ar-dz": "./node_modules/moment/locale/ar-dz.js",
            "./ar-dz.js": "./node_modules/moment/locale/ar-dz.js",
            "./ar-kw": "./node_modules/moment/locale/ar-kw.js",
            "./ar-kw.js": "./node_modules/moment/locale/ar-kw.js",
            "./ar-ly": "./node_modules/moment/locale/ar-ly.js",
            "./ar-ly.js": "./node_modules/moment/locale/ar-ly.js",
            "./ar-ma": "./node_modules/moment/locale/ar-ma.js",
            "./ar-ma.js": "./node_modules/moment/locale/ar-ma.js",
            "./ar-sa": "./node_modules/moment/locale/ar-sa.js",
            "./ar-sa.js": "./node_modules/moment/locale/ar-sa.js",
            "./ar-tn": "./node_modules/moment/locale/ar-tn.js",
            "./ar-tn.js": "./node_modules/moment/locale/ar-tn.js",
            "./ar.js": "./node_modules/moment/locale/ar.js",
            "./az": "./node_modules/moment/locale/az.js",
            "./az.js": "./node_modules/moment/locale/az.js",
            "./be": "./node_modules/moment/locale/be.js",
            "./be.js": "./node_modules/moment/locale/be.js",
            "./bg": "./node_modules/moment/locale/bg.js",
            "./bg.js": "./node_modules/moment/locale/bg.js",
            "./bm": "./node_modules/moment/locale/bm.js",
            "./bm.js": "./node_modules/moment/locale/bm.js",
            "./bn": "./node_modules/moment/locale/bn.js",
            "./bn.js": "./node_modules/moment/locale/bn.js",
            "./bo": "./node_modules/moment/locale/bo.js",
            "./bo.js": "./node_modules/moment/locale/bo.js",
            "./br": "./node_modules/moment/locale/br.js",
            "./br.js": "./node_modules/moment/locale/br.js",
            "./bs": "./node_modules/moment/locale/bs.js",
            "./bs.js": "./node_modules/moment/locale/bs.js",
            "./ca": "./node_modules/moment/locale/ca.js",
            "./ca.js": "./node_modules/moment/locale/ca.js",
            "./cs": "./node_modules/moment/locale/cs.js",
            "./cs.js": "./node_modules/moment/locale/cs.js",
            "./cv": "./node_modules/moment/locale/cv.js",
            "./cv.js": "./node_modules/moment/locale/cv.js",
            "./cy": "./node_modules/moment/locale/cy.js",
            "./cy.js": "./node_modules/moment/locale/cy.js",
            "./da": "./node_modules/moment/locale/da.js",
            "./da.js": "./node_modules/moment/locale/da.js",
            "./de": "./node_modules/moment/locale/de.js",
            "./de-at": "./node_modules/moment/locale/de-at.js",
            "./de-at.js": "./node_modules/moment/locale/de-at.js",
            "./de-ch": "./node_modules/moment/locale/de-ch.js",
            "./de-ch.js": "./node_modules/moment/locale/de-ch.js",
            "./de.js": "./node_modules/moment/locale/de.js",
            "./dv": "./node_modules/moment/locale/dv.js",
            "./dv.js": "./node_modules/moment/locale/dv.js",
            "./el": "./node_modules/moment/locale/el.js",
            "./el.js": "./node_modules/moment/locale/el.js",
            "./en-au": "./node_modules/moment/locale/en-au.js",
            "./en-au.js": "./node_modules/moment/locale/en-au.js",
            "./en-ca": "./node_modules/moment/locale/en-ca.js",
            "./en-ca.js": "./node_modules/moment/locale/en-ca.js",
            "./en-gb": "./node_modules/moment/locale/en-gb.js",
            "./en-gb.js": "./node_modules/moment/locale/en-gb.js",
            "./en-ie": "./node_modules/moment/locale/en-ie.js",
            "./en-ie.js": "./node_modules/moment/locale/en-ie.js",
            "./en-il": "./node_modules/moment/locale/en-il.js",
            "./en-il.js": "./node_modules/moment/locale/en-il.js",
            "./en-in": "./node_modules/moment/locale/en-in.js",
            "./en-in.js": "./node_modules/moment/locale/en-in.js",
            "./en-nz": "./node_modules/moment/locale/en-nz.js",
            "./en-nz.js": "./node_modules/moment/locale/en-nz.js",
            "./en-sg": "./node_modules/moment/locale/en-sg.js",
            "./en-sg.js": "./node_modules/moment/locale/en-sg.js",
            "./eo": "./node_modules/moment/locale/eo.js",
            "./eo.js": "./node_modules/moment/locale/eo.js",
            "./es": "./node_modules/moment/locale/es.js",
            "./es-do": "./node_modules/moment/locale/es-do.js",
            "./es-do.js": "./node_modules/moment/locale/es-do.js",
            "./es-us": "./node_modules/moment/locale/es-us.js",
            "./es-us.js": "./node_modules/moment/locale/es-us.js",
            "./es.js": "./node_modules/moment/locale/es.js",
            "./et": "./node_modules/moment/locale/et.js",
            "./et.js": "./node_modules/moment/locale/et.js",
            "./eu": "./node_modules/moment/locale/eu.js",
            "./eu.js": "./node_modules/moment/locale/eu.js",
            "./fa": "./node_modules/moment/locale/fa.js",
            "./fa.js": "./node_modules/moment/locale/fa.js",
            "./fi": "./node_modules/moment/locale/fi.js",
            "./fi.js": "./node_modules/moment/locale/fi.js",
            "./fil": "./node_modules/moment/locale/fil.js",
            "./fil.js": "./node_modules/moment/locale/fil.js",
            "./fo": "./node_modules/moment/locale/fo.js",
            "./fo.js": "./node_modules/moment/locale/fo.js",
            "./fr": "./node_modules/moment/locale/fr.js",
            "./fr-ca": "./node_modules/moment/locale/fr-ca.js",
            "./fr-ca.js": "./node_modules/moment/locale/fr-ca.js",
            "./fr-ch": "./node_modules/moment/locale/fr-ch.js",
            "./fr-ch.js": "./node_modules/moment/locale/fr-ch.js",
            "./fr.js": "./node_modules/moment/locale/fr.js",
            "./fy": "./node_modules/moment/locale/fy.js",
            "./fy.js": "./node_modules/moment/locale/fy.js",
            "./ga": "./node_modules/moment/locale/ga.js",
            "./ga.js": "./node_modules/moment/locale/ga.js",
            "./gd": "./node_modules/moment/locale/gd.js",
            "./gd.js": "./node_modules/moment/locale/gd.js",
            "./gl": "./node_modules/moment/locale/gl.js",
            "./gl.js": "./node_modules/moment/locale/gl.js",
            "./gom-deva": "./node_modules/moment/locale/gom-deva.js",
            "./gom-deva.js": "./node_modules/moment/locale/gom-deva.js",
            "./gom-latn": "./node_modules/moment/locale/gom-latn.js",
            "./gom-latn.js": "./node_modules/moment/locale/gom-latn.js",
            "./gu": "./node_modules/moment/locale/gu.js",
            "./gu.js": "./node_modules/moment/locale/gu.js",
            "./he": "./node_modules/moment/locale/he.js",
            "./he.js": "./node_modules/moment/locale/he.js",
            "./hi": "./node_modules/moment/locale/hi.js",
            "./hi.js": "./node_modules/moment/locale/hi.js",
            "./hr": "./node_modules/moment/locale/hr.js",
            "./hr.js": "./node_modules/moment/locale/hr.js",
            "./hu": "./node_modules/moment/locale/hu.js",
            "./hu.js": "./node_modules/moment/locale/hu.js",
            "./hy-am": "./node_modules/moment/locale/hy-am.js",
            "./hy-am.js": "./node_modules/moment/locale/hy-am.js",
            "./id": "./node_modules/moment/locale/id.js",
            "./id.js": "./node_modules/moment/locale/id.js",
            "./is": "./node_modules/moment/locale/is.js",
            "./is.js": "./node_modules/moment/locale/is.js",
            "./it": "./node_modules/moment/locale/it.js",
            "./it-ch": "./node_modules/moment/locale/it-ch.js",
            "./it-ch.js": "./node_modules/moment/locale/it-ch.js",
            "./it.js": "./node_modules/moment/locale/it.js",
            "./ja": "./node_modules/moment/locale/ja.js",
            "./ja.js": "./node_modules/moment/locale/ja.js",
            "./jv": "./node_modules/moment/locale/jv.js",
            "./jv.js": "./node_modules/moment/locale/jv.js",
            "./ka": "./node_modules/moment/locale/ka.js",
            "./ka.js": "./node_modules/moment/locale/ka.js",
            "./kk": "./node_modules/moment/locale/kk.js",
            "./kk.js": "./node_modules/moment/locale/kk.js",
            "./km": "./node_modules/moment/locale/km.js",
            "./km.js": "./node_modules/moment/locale/km.js",
            "./kn": "./node_modules/moment/locale/kn.js",
            "./kn.js": "./node_modules/moment/locale/kn.js",
            "./ko": "./node_modules/moment/locale/ko.js",
            "./ko.js": "./node_modules/moment/locale/ko.js",
            "./ku": "./node_modules/moment/locale/ku.js",
            "./ku.js": "./node_modules/moment/locale/ku.js",
            "./ky": "./node_modules/moment/locale/ky.js",
            "./ky.js": "./node_modules/moment/locale/ky.js",
            "./lb": "./node_modules/moment/locale/lb.js",
            "./lb.js": "./node_modules/moment/locale/lb.js",
            "./lo": "./node_modules/moment/locale/lo.js",
            "./lo.js": "./node_modules/moment/locale/lo.js",
            "./lt": "./node_modules/moment/locale/lt.js",
            "./lt.js": "./node_modules/moment/locale/lt.js",
            "./lv": "./node_modules/moment/locale/lv.js",
            "./lv.js": "./node_modules/moment/locale/lv.js",
            "./me": "./node_modules/moment/locale/me.js",
            "./me.js": "./node_modules/moment/locale/me.js",
            "./mi": "./node_modules/moment/locale/mi.js",
            "./mi.js": "./node_modules/moment/locale/mi.js",
            "./mk": "./node_modules/moment/locale/mk.js",
            "./mk.js": "./node_modules/moment/locale/mk.js",
            "./ml": "./node_modules/moment/locale/ml.js",
            "./ml.js": "./node_modules/moment/locale/ml.js",
            "./mn": "./node_modules/moment/locale/mn.js",
            "./mn.js": "./node_modules/moment/locale/mn.js",
            "./mr": "./node_modules/moment/locale/mr.js",
            "./mr.js": "./node_modules/moment/locale/mr.js",
            "./ms": "./node_modules/moment/locale/ms.js",
            "./ms-my": "./node_modules/moment/locale/ms-my.js",
            "./ms-my.js": "./node_modules/moment/locale/ms-my.js",
            "./ms.js": "./node_modules/moment/locale/ms.js",
            "./mt": "./node_modules/moment/locale/mt.js",
            "./mt.js": "./node_modules/moment/locale/mt.js",
            "./my": "./node_modules/moment/locale/my.js",
            "./my.js": "./node_modules/moment/locale/my.js",
            "./nb": "./node_modules/moment/locale/nb.js",
            "./nb.js": "./node_modules/moment/locale/nb.js",
            "./ne": "./node_modules/moment/locale/ne.js",
            "./ne.js": "./node_modules/moment/locale/ne.js",
            "./nl": "./node_modules/moment/locale/nl.js",
            "./nl-be": "./node_modules/moment/locale/nl-be.js",
            "./nl-be.js": "./node_modules/moment/locale/nl-be.js",
            "./nl.js": "./node_modules/moment/locale/nl.js",
            "./nn": "./node_modules/moment/locale/nn.js",
            "./nn.js": "./node_modules/moment/locale/nn.js",
            "./oc-lnc": "./node_modules/moment/locale/oc-lnc.js",
            "./oc-lnc.js": "./node_modules/moment/locale/oc-lnc.js",
            "./pa-in": "./node_modules/moment/locale/pa-in.js",
            "./pa-in.js": "./node_modules/moment/locale/pa-in.js",
            "./pl": "./node_modules/moment/locale/pl.js",
            "./pl.js": "./node_modules/moment/locale/pl.js",
            "./pt": "./node_modules/moment/locale/pt.js",
            "./pt-br": "./node_modules/moment/locale/pt-br.js",
            "./pt-br.js": "./node_modules/moment/locale/pt-br.js",
            "./pt.js": "./node_modules/moment/locale/pt.js",
            "./ro": "./node_modules/moment/locale/ro.js",
            "./ro.js": "./node_modules/moment/locale/ro.js",
            "./ru": "./node_modules/moment/locale/ru.js",
            "./ru.js": "./node_modules/moment/locale/ru.js",
            "./sd": "./node_modules/moment/locale/sd.js",
            "./sd.js": "./node_modules/moment/locale/sd.js",
            "./se": "./node_modules/moment/locale/se.js",
            "./se.js": "./node_modules/moment/locale/se.js",
            "./si": "./node_modules/moment/locale/si.js",
            "./si.js": "./node_modules/moment/locale/si.js",
            "./sk": "./node_modules/moment/locale/sk.js",
            "./sk.js": "./node_modules/moment/locale/sk.js",
            "./sl": "./node_modules/moment/locale/sl.js",
            "./sl.js": "./node_modules/moment/locale/sl.js",
            "./sq": "./node_modules/moment/locale/sq.js",
            "./sq.js": "./node_modules/moment/locale/sq.js",
            "./sr": "./node_modules/moment/locale/sr.js",
            "./sr-cyrl": "./node_modules/moment/locale/sr-cyrl.js",
            "./sr-cyrl.js": "./node_modules/moment/locale/sr-cyrl.js",
            "./sr.js": "./node_modules/moment/locale/sr.js",
            "./ss": "./node_modules/moment/locale/ss.js",
            "./ss.js": "./node_modules/moment/locale/ss.js",
            "./sv": "./node_modules/moment/locale/sv.js",
            "./sv.js": "./node_modules/moment/locale/sv.js",
            "./sw": "./node_modules/moment/locale/sw.js",
            "./sw.js": "./node_modules/moment/locale/sw.js",
            "./ta": "./node_modules/moment/locale/ta.js",
            "./ta.js": "./node_modules/moment/locale/ta.js",
            "./te": "./node_modules/moment/locale/te.js",
            "./te.js": "./node_modules/moment/locale/te.js",
            "./tet": "./node_modules/moment/locale/tet.js",
            "./tet.js": "./node_modules/moment/locale/tet.js",
            "./tg": "./node_modules/moment/locale/tg.js",
            "./tg.js": "./node_modules/moment/locale/tg.js",
            "./th": "./node_modules/moment/locale/th.js",
            "./th.js": "./node_modules/moment/locale/th.js",
            "./tl-ph": "./node_modules/moment/locale/tl-ph.js",
            "./tl-ph.js": "./node_modules/moment/locale/tl-ph.js",
            "./tlh": "./node_modules/moment/locale/tlh.js",
            "./tlh.js": "./node_modules/moment/locale/tlh.js",
            "./tr": "./node_modules/moment/locale/tr.js",
            "./tr.js": "./node_modules/moment/locale/tr.js",
            "./tzl": "./node_modules/moment/locale/tzl.js",
            "./tzl.js": "./node_modules/moment/locale/tzl.js",
            "./tzm": "./node_modules/moment/locale/tzm.js",
            "./tzm-latn": "./node_modules/moment/locale/tzm-latn.js",
            "./tzm-latn.js": "./node_modules/moment/locale/tzm-latn.js",
            "./tzm.js": "./node_modules/moment/locale/tzm.js",
            "./ug-cn": "./node_modules/moment/locale/ug-cn.js",
            "./ug-cn.js": "./node_modules/moment/locale/ug-cn.js",
            "./uk": "./node_modules/moment/locale/uk.js",
            "./uk.js": "./node_modules/moment/locale/uk.js",
            "./ur": "./node_modules/moment/locale/ur.js",
            "./ur.js": "./node_modules/moment/locale/ur.js",
            "./uz": "./node_modules/moment/locale/uz.js",
            "./uz-latn": "./node_modules/moment/locale/uz-latn.js",
            "./uz-latn.js": "./node_modules/moment/locale/uz-latn.js",
            "./uz.js": "./node_modules/moment/locale/uz.js",
            "./vi": "./node_modules/moment/locale/vi.js",
            "./vi.js": "./node_modules/moment/locale/vi.js",
            "./x-pseudo": "./node_modules/moment/locale/x-pseudo.js",
            "./x-pseudo.js": "./node_modules/moment/locale/x-pseudo.js",
            "./yo": "./node_modules/moment/locale/yo.js",
            "./yo.js": "./node_modules/moment/locale/yo.js",
            "./zh-cn": "./node_modules/moment/locale/zh-cn.js",
            "./zh-cn.js": "./node_modules/moment/locale/zh-cn.js",
            "./zh-hk": "./node_modules/moment/locale/zh-hk.js",
            "./zh-hk.js": "./node_modules/moment/locale/zh-hk.js",
            "./zh-mo": "./node_modules/moment/locale/zh-mo.js",
            "./zh-mo.js": "./node_modules/moment/locale/zh-mo.js",
            "./zh-tw": "./node_modules/moment/locale/zh-tw.js",
            "./zh-tw.js": "./node_modules/moment/locale/zh-tw.js"
        };


        function webpackContext(req) {
            var id = webpackContextResolve(req);
            return __webpack_require__(id);
        }

        function webpackContextResolve(req) {
            if (!__webpack_require__.o(map, req)) {
                var e = new Error("Cannot find module '" + req + "'");
                e.code = 'MODULE_NOT_FOUND';
                throw e;
            }
            return map[req];
        }

        webpackContext.keys = function webpackContextKeys() {
            return Object.keys(map);
        };
        webpackContext.resolve = webpackContextResolve;
        module.exports = webpackContext;
        webpackContext.id = "./node_modules/moment/locale sync recursive ^\\.\\/.*$";

        /***/
    }),

    /***/
    "./node_modules/vue-loader/lib/loaders/templateLoader.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/components/Dashboard.vue?vue&type=template&id=1f65406d&":
    /*!*******************************************************************************************************************************************************************************************************************!*\
      !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib??vue-loader-options!./resources/assets/js/components/Dashboard.vue?vue&type=template&id=1f65406d& ***!
      \*******************************************************************************************************************************************************************************************************************/
    /*! exports provided: render, staticRenderFns */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony export (binding) */
        __webpack_require__.d(__webpack_exports__, "render", function () {
            return render;
        });
        /* harmony export (binding) */
        __webpack_require__.d(__webpack_exports__, "staticRenderFns", function () {
            return staticRenderFns;
        });
        var render = function () {
            var _vm = this
            var _h = _vm.$createElement
            var _c = _vm._self._c || _h
            return _c(
                "div",
                {
                    staticClass: "leading-normal tracking-normal",
                    attrs: {id: "main-body"}
                },
                [
                    _c(
                        "div",
                        {staticClass: "flex flex-wrap"},
                        [
                            _c("Sidebar"),
                            _vm._v(" "),
                            _c(
                                "div",
                                {
                                    staticClass: "w-full bg-gray-100 pl-0 lg:pl-64 min-h-screen",
                                    class: _vm.sideBarOpen ? "overlay" : "",
                                    attrs: {id: "main-content"}
                                },
                                [
                                    _c("Navbar"),
                                    _vm._v(" "),
                                    _c(
                                        "div",
                                        {staticClass: "p-6 bg-gray-100 mb-20"},
                                        [_c("router-view")],
                                        1
                                    ),
                                    _vm._v(" "),
                                    _c("Footer")
                                ],
                                1
                            )
                        ],
                        1
                    )
                ]
            )
        }
        var staticRenderFns = []
        render._withStripped = true


        /***/
    }),

    /***/
    "./node_modules/vue-loader/lib/loaders/templateLoader.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/components/Footer.vue?vue&type=template&id=083ff5dc&":
    /*!****************************************************************************************************************************************************************************************************************!*\
      !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib??vue-loader-options!./resources/assets/js/components/Footer.vue?vue&type=template&id=083ff5dc& ***!
      \****************************************************************************************************************************************************************************************************************/
    /*! exports provided: render, staticRenderFns */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony export (binding) */
        __webpack_require__.d(__webpack_exports__, "render", function () {
            return render;
        });
        /* harmony export (binding) */
        __webpack_require__.d(__webpack_exports__, "staticRenderFns", function () {
            return staticRenderFns;
        });
        var render = function () {
            var _vm = this
            var _h = _vm.$createElement
            var _c = _vm._self._c || _h
            return _vm._m(0)
        }
        var staticRenderFns = [
            function () {
                var _vm = this
                var _h = _vm.$createElement
                var _c = _vm._self._c || _h
                return _c(
                    "div",
                    {
                        staticClass:
                            "w-full border-t-2 px-8 py-6 lg:flex justify-between items-center"
                    },
                    [
                        _c("p", {staticClass: "mb-2 lg:mb-0"}, [_vm._v("© Copyright 2020")]),
                        _vm._v(" "),
                        _c("div", {staticClass: "flex"}, [
                            _c(
                                "a",
                                {staticClass: "mr-6 hover:text-gray-900", attrs: {href: "#"}},
                                [_vm._v("Terms of Service")]
                            ),
                            _vm._v(" "),
                            _c(
                                "a",
                                {staticClass: "mr-6 hover:text-gray-900", attrs: {href: "#"}},
                                [_vm._v("Privacy Policy")]
                            ),
                            _vm._v(" "),
                            _c(
                                "a",
                                {staticClass: "hover:text-gray-900", attrs: {href: "#"}},
                                [_vm._v("About Us")]
                            )
                        ])
                    ]
                )
            }
        ]
        render._withStripped = true


        /***/
    }),

    /***/
    "./node_modules/vue-loader/lib/loaders/templateLoader.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/components/Navbar.vue?vue&type=template&id=cadbadf2&":
    /*!****************************************************************************************************************************************************************************************************************!*\
      !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib??vue-loader-options!./resources/assets/js/components/Navbar.vue?vue&type=template&id=cadbadf2& ***!
      \****************************************************************************************************************************************************************************************************************/
    /*! exports provided: render, staticRenderFns */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony export (binding) */
        __webpack_require__.d(__webpack_exports__, "render", function () {
            return render;
        });
        /* harmony export (binding) */
        __webpack_require__.d(__webpack_exports__, "staticRenderFns", function () {
            return staticRenderFns;
        });
        var render = function () {
            var _vm = this
            var _h = _vm.$createElement
            var _c = _vm._self._c || _h
            return _c("div", {staticClass: "sticky top-0 z-40"}, [
                _c(
                    "div",
                    {
                        staticClass:
                            "w-full h-20 px-6 bg-gray-100 border-b flex items-center justify-between"
                    },
                    [
                        _c("div", {staticClass: "flex"}, [
                            _c(
                                "div",
                                {staticClass: "inline-block lg:hidden flex items-center mr-4"},
                                [
                                    _c(
                                        "button",
                                        {
                                            staticClass:
                                                "hover:text-blue-500 hover:border-white focus:outline-none navbar-burger",
                                            on: {
                                                click: function ($event) {
                                                    return _vm.toggleSidebar()
                                                }
                                            }
                                        },
                                        [
                                            _c(
                                                "svg",
                                                {
                                                    staticClass: "h-5 w-5",
                                                    style: {fill: "black"},
                                                    attrs: {
                                                        viewBox: "0 0 20 20",
                                                        xmlns: "http://www.w3.org/2000/svg"
                                                    }
                                                },
                                                [
                                                    _c("title", [_vm._v("Menu")]),
                                                    _c("path", {
                                                        attrs: {
                                                            d: "M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z"
                                                        }
                                                    })
                                                ]
                                            )
                                        ]
                                    )
                                ]
                            ),
                            _vm._v(" "),
                            _c("div", {staticClass: "relative text-gray-600"}, [
                                _c("input", {
                                    staticClass:
                                        "bg-white h-10 w-full xl:w-64 px-5 rounded-lg border text-sm focus:outline-none",
                                    attrs: {
                                        type: "search",
                                        name: "serch",
                                        placeholder: "Search products..."
                                    }
                                }),
                                _vm._v(" "),
                                _c(
                                    "button",
                                    {
                                        staticClass: "absolute right-0 top-0 mt-3 mr-4",
                                        attrs: {type: "submit"}
                                    },
                                    [
                                        _c(
                                            "svg",
                                            {
                                                staticClass: "h-4 w-4 fill-current",
                                                staticStyle: {
                                                    "enable-background": "new 0 0 56.966 56.966"
                                                },
                                                attrs: {
                                                    xmlns: "http://www.w3.org/2000/svg",
                                                    "xmlns:xlink": "http://www.w3.org/1999/xlink",
                                                    version: "1.1",
                                                    id: "Capa_1",
                                                    x: "0px",
                                                    y: "0px",
                                                    viewBox: "0 0 56.966 56.966",
                                                    "xml:space": "preserve",
                                                    width: "512px",
                                                    height: "512px"
                                                }
                                            },
                                            [
                                                _c("path", {
                                                    attrs: {
                                                        d:
                                                            "M55.146,51.887L41.588,37.786c3.486-4.144,5.396-9.358,5.396-14.786c0-12.682-10.318-23-23-23s-23,10.318-23,23  s10.318,23,23,23c4.761,0,9.298-1.436,13.177-4.162l13.661,14.208c0.571,0.593,1.339,0.92,2.162,0.92  c0.779,0,1.518-0.297,2.079-0.837C56.255,54.982,56.293,53.08,55.146,51.887z M23.984,6c9.374,0,17,7.626,17,17s-7.626,17-17,17  s-17-7.626-17-17S14.61,6,23.984,6z"
                                                    }
                                                })
                                            ]
                                        )
                                    ]
                                )
                            ])
                        ]),
                        _vm._v(" "),
                        _c("div", {staticClass: "flex items-center relative"}, [
                            _c(
                                "svg",
                                {
                                    staticClass: "fill-current mr-3 hover:text-blue-500",
                                    attrs: {
                                        xmlns: "http://www.w3.org/2000/svg",
                                        height: "24",
                                        viewBox: "0 0 24 24",
                                        width: "24"
                                    }
                                },
                                [
                                    _c("path", {attrs: {d: "M0 0h24v24H0z", fill: "none"}}),
                                    _c("path", {
                                        attrs: {
                                            d:
                                                "M12 22c1.1 0 2-.9 2-2h-4c0 1.1.9 2 2 2zm6-6v-5c0-3.07-1.63-5.64-4.5-6.32V4c0-.83-.67-1.5-1.5-1.5s-1.5.67-1.5 1.5v.68C7.64 5.36 6 7.92 6 11v5l-2 2v1h16v-1l-2-2zm-2 1H8v-6c0-2.48 1.51-4.5 4-4.5s4 2.02 4 4.5v6z"
                                        }
                                    })
                                ]
                            ),
                            _vm._v(" "),
                            _c("img", {
                                staticClass: "w-12 h-12 rounded-full shadow-lg",
                                attrs: {
                                    src: "https://a7sas.net/wp-content/uploads/2019/07/4060.jpeg"
                                },
                                on: {
                                    click: function ($event) {
                                        _vm.dropDownOpen = !_vm.dropDownOpen
                                    }
                                }
                            })
                        ])
                    ]
                ),
                _vm._v(" "),
                _c(
                    "div",
                    {
                        staticClass:
                            "absolute bg-gray-100 border border-t-0 shadow-xl text-gray-700 rounded-b-lg w-48 bottom-10 right-0 mr-6",
                        class: _vm.dropDownOpen ? "" : "hidden"
                    },
                    [
                        _c(
                            "a",
                            {
                                staticClass: "block px-4 py-2 hover:bg-gray-200",
                                attrs: {href: "#"}
                            },
                            [_vm._v("Account")]
                        ),
                        _vm._v(" "),
                        _c(
                            "a",
                            {
                                staticClass: "block px-4 py-2 hover:bg-gray-200",
                                attrs: {href: "#"}
                            },
                            [_vm._v("Settings")]
                        ),
                        _vm._v(" "),
                        _c(
                            "a",
                            {
                                staticClass: "block px-4 py-2 hover:bg-gray-200",
                                attrs: {href: "#"}
                            },
                            [_vm._v("Logout")]
                        )
                    ]
                )
            ])
        }
        var staticRenderFns = []
        render._withStripped = true


        /***/
    }),

    /***/
    "./node_modules/vue-loader/lib/loaders/templateLoader.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/components/Sidebar.vue?vue&type=template&id=28cb1975&":
    /*!*****************************************************************************************************************************************************************************************************************!*\
      !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib??vue-loader-options!./resources/assets/js/components/Sidebar.vue?vue&type=template&id=28cb1975& ***!
      \*****************************************************************************************************************************************************************************************************************/
    /*! exports provided: render, staticRenderFns */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony export (binding) */
        __webpack_require__.d(__webpack_exports__, "render", function () {
            return render;
        });
        /* harmony export (binding) */
        __webpack_require__.d(__webpack_exports__, "staticRenderFns", function () {
            return staticRenderFns;
        });
        var render = function () {
            var _vm = this
            var _h = _vm.$createElement
            var _c = _vm._self._c || _h
            return _c(
                "div",
                {
                    staticClass:
                        "w-1/2 md:w-1/3 lg:w-64 fixed md:top-0 md:left-0 h-screen lg:block bg-gray-100 border-r z-30",
                    class: _vm.sideBarOpen ? "" : "hidden",
                    attrs: {id: "main-nav"}
                },
                [
                    _c(
                        "div",
                        {staticClass: "w-full h-20 border-b flex px-4 items-center mb-8"},
                        [
                            _c("router-link", {attrs: {to: "/dashboard/home"}}, [
                                _c(
                                    "div",
                                    {staticClass: "font-semibold text-3xl text-blue-400 pl-4"},
                                    [_vm._v("LOGO")]
                                )
                            ])
                        ],
                        1
                    ),
                    _vm._v(" "),
                    _c(
                        "div",
                        {staticClass: "mb-4 px-4"},
                        [
                            _c("p", {staticClass: "pl-4 text-sm font-semibold mb-1"}, [
                                _vm._v("MAIN")
                            ]),
                            _vm._v(" "),
                            _c("router-link", {attrs: {to: "/dashboard/home"}}, [
                                _c(
                                    "div",
                                    {
                                        staticClass:
                                            "w-full flex items-center text-blue-400 h-10 pl-4 bg-gray-200 hover:bg-gray-200 rounded-lg cursor-pointer"
                                    },
                                    [
                                        _c(
                                            "svg",
                                            {
                                                staticClass: "h-6 w-6 fill-current mr-2",
                                                attrs: {viewBox: "0 0 20 20"}
                                            },
                                            [
                                                _c("path", {
                                                    attrs: {
                                                        d:
                                                            "M18.121,9.88l-7.832-7.836c-0.155-0.158-0.428-0.155-0.584,0L1.842,9.913c-0.262,0.263-0.073,0.705,0.292,0.705h2.069v7.042c0,0.227,0.187,0.414,0.414,0.414h3.725c0.228,0,0.414-0.188,0.414-0.414v-3.313h2.483v3.313c0,0.227,0.187,0.414,0.413,0.414h3.726c0.229,0,0.414-0.188,0.414-0.414v-7.042h2.068h0.004C18.331,10.617,18.389,10.146,18.121,9.88 M14.963,17.245h-2.896v-3.313c0-0.229-0.186-0.415-0.414-0.415H8.342c-0.228,0-0.414,0.187-0.414,0.415v3.313H5.032v-6.628h9.931V17.245z M3.133,9.79l6.864-6.868l6.867,6.868H3.133z"
                                                    }
                                                })
                                            ]
                                        ),
                                        _vm._v(" "),
                                        _c("span", {staticClass: "text-gray-700"}, [
                                            _vm._v("Dashboard")
                                        ])
                                    ]
                                )
                            ]),
                            _vm._v(" "),
                            _c(
                                "div",
                                {
                                    staticClass:
                                        "w-full flex items-center text-blue-400 h-10 pl-4 hover:bg-gray-200 rounded-lg cursor-pointer"
                                },
                                [
                                    _c(
                                        "svg",
                                        {
                                            staticClass: "h-6 w-6 fill-current mr-2",
                                            attrs: {viewBox: "0 0 20 20"}
                                        },
                                        [
                                            _c("path", {
                                                attrs: {
                                                    d:
                                                        "M17.431,2.156h-3.715c-0.228,0-0.413,0.186-0.413,0.413v6.973h-2.89V6.687c0-0.229-0.186-0.413-0.413-0.413H6.285c-0.228,0-0.413,0.184-0.413,0.413v6.388H2.569c-0.227,0-0.413,0.187-0.413,0.413v3.942c0,0.228,0.186,0.413,0.413,0.413h14.862c0.228,0,0.413-0.186,0.413-0.413V2.569C17.844,2.342,17.658,2.156,17.431,2.156 M5.872,17.019h-2.89v-3.117h2.89V17.019zM9.587,17.019h-2.89V7.1h2.89V17.019z M13.303,17.019h-2.89v-6.651h2.89V17.019z M17.019,17.019h-2.891V2.982h2.891V17.019z"
                                                }
                                            })
                                        ]
                                    ),
                                    _vm._v(" "),
                                    _c("span", {staticClass: "text-gray-700"}, [
                                        _vm._v("Analytics")
                                    ])
                                ]
                            ),
                            _vm._v(" "),
                            _c(
                                "div",
                                {
                                    staticClass:
                                        "w-full flex items-center text-blue-400 h-10 pl-4 hover:bg-gray-200 rounded-lg cursor-pointer"
                                },
                                [
                                    _c(
                                        "svg",
                                        {
                                            staticClass: "h-6 w-6 fill-current mr-2",
                                            attrs: {viewBox: "0 0 20 20"}
                                        },
                                        [
                                            _c("path", {
                                                attrs: {
                                                    d:
                                                        "M17.283,5.549h-5.26V4.335c0-0.222-0.183-0.404-0.404-0.404H8.381c-0.222,0-0.404,0.182-0.404,0.404v1.214h-5.26c-0.223,0-0.405,0.182-0.405,0.405v9.71c0,0.223,0.182,0.405,0.405,0.405h14.566c0.223,0,0.404-0.183,0.404-0.405v-9.71C17.688,5.731,17.506,5.549,17.283,5.549 M8.786,4.74h2.428v0.809H8.786V4.74z M16.879,15.26H3.122v-4.046h5.665v1.201c0,0.223,0.182,0.404,0.405,0.404h1.618c0.222,0,0.405-0.182,0.405-0.404v-1.201h5.665V15.26z M9.595,9.583h0.81v2.428h-0.81V9.583zM16.879,10.405h-5.665V9.19c0-0.222-0.183-0.405-0.405-0.405H9.191c-0.223,0-0.405,0.183-0.405,0.405v1.215H3.122V6.358h13.757V10.405z"
                                                }
                                            })
                                        ]
                                    ),
                                    _vm._v(" "),
                                    _c("span", {staticClass: "text-gray-700"}, [
                                        _vm._v("Inventory")
                                    ])
                                ]
                            ),
                            _vm._v(" "),
                            _c(
                                "div",
                                {
                                    staticClass:
                                        "w-full flex items-center text-blue-400 h-10 pl-4 hover:bg-gray-200 rounded-lg cursor-pointer"
                                },
                                [
                                    _c(
                                        "svg",
                                        {
                                            staticClass: "h-6 w-6 fill-current mr-2",
                                            attrs: {viewBox: "0 0 20 20"}
                                        },
                                        [
                                            _c("path", {
                                                attrs: {
                                                    d:
                                                        "M18.303,4.742l-1.454-1.455c-0.171-0.171-0.475-0.171-0.646,0l-3.061,3.064H2.019c-0.251,0-0.457,0.205-0.457,0.456v9.578c0,0.251,0.206,0.456,0.457,0.456h13.683c0.252,0,0.457-0.205,0.457-0.456V7.533l2.144-2.146C18.481,5.208,18.483,4.917,18.303,4.742 M15.258,15.929H2.476V7.263h9.754L9.695,9.792c-0.057,0.057-0.101,0.13-0.119,0.212L9.18,11.36h-3.98c-0.251,0-0.457,0.205-0.457,0.456c0,0.253,0.205,0.456,0.457,0.456h4.336c0.023,0,0.899,0.02,1.498-0.127c0.312-0.077,0.55-0.137,0.55-0.137c0.08-0.018,0.155-0.059,0.212-0.118l3.463-3.443V15.929z M11.241,11.156l-1.078,0.267l0.267-1.076l6.097-6.091l0.808,0.808L11.241,11.156z"
                                                }
                                            })
                                        ]
                                    ),
                                    _vm._v(" "),
                                    _c("span", {staticClass: "text-gray-700"}, [
                                        _vm._v("Enquiries")
                                    ])
                                ]
                            ),
                            _vm._v(" "),
                            _c(
                                "div",
                                {
                                    staticClass:
                                        "w-full flex items-center text-blue-400 h-10 pl-4 hover:bg-gray-200 rounded-lg cursor-pointer"
                                },
                                [
                                    _c(
                                        "svg",
                                        {
                                            staticClass: "h-6 w-6 fill-current mr-2",
                                            attrs: {viewBox: "0 0 20 20"}
                                        },
                                        [
                                            _c("path", {
                                                attrs: {
                                                    d:
                                                        "M16.557,4.467h-1.64v-0.82c0-0.225-0.183-0.41-0.409-0.41c-0.226,0-0.41,0.185-0.41,0.41v0.82H5.901v-0.82c0-0.225-0.185-0.41-0.41-0.41c-0.226,0-0.41,0.185-0.41,0.41v0.82H3.442c-0.904,0-1.64,0.735-1.64,1.639v9.017c0,0.904,0.736,1.64,1.64,1.64h13.114c0.904,0,1.64-0.735,1.64-1.64V6.106C18.196,5.203,17.461,4.467,16.557,4.467 M17.377,15.123c0,0.453-0.366,0.819-0.82,0.819H3.442c-0.453,0-0.82-0.366-0.82-0.819V8.976h14.754V15.123z M17.377,8.156H2.623V6.106c0-0.453,0.367-0.82,0.82-0.82h1.639v1.23c0,0.225,0.184,0.41,0.41,0.41c0.225,0,0.41-0.185,0.41-0.41v-1.23h8.196v1.23c0,0.225,0.185,0.41,0.41,0.41c0.227,0,0.409-0.185,0.409-0.41v-1.23h1.64c0.454,0,0.82,0.367,0.82,0.82V8.156z"
                                                }
                                            })
                                        ]
                                    ),
                                    _vm._v(" "),
                                    _c("span", {staticClass: "text-gray-700"}, [_vm._v("Calender")])
                                ]
                            )
                        ],
                        1
                    ),
                    _vm._v(" "),
                    _c("div", {staticClass: "mb-4 px-4"}, [
                        _c("p", {staticClass: "pl-4 text-sm font-semibold mb-1"}, [
                            _vm._v("PRODUCTS")
                        ]),
                        _vm._v(" "),
                        _c(
                            "div",
                            {
                                staticClass:
                                    "w-full flex items-center text-blue-400 h-10 pl-4 hover:bg-gray-200 rounded-lg cursor-pointer"
                            },
                            [
                                _c(
                                    "svg",
                                    {
                                        staticClass: "h-6 w-6 fill-current mr-2",
                                        attrs: {viewBox: "0 0 20 20"}
                                    },
                                    [
                                        _c("path", {
                                            attrs: {
                                                d:
                                                    "M14.613,10c0,0.23-0.188,0.419-0.419,0.419H10.42v3.774c0,0.23-0.189,0.42-0.42,0.42s-0.419-0.189-0.419-0.42v-3.774H5.806c-0.23,0-0.419-0.189-0.419-0.419s0.189-0.419,0.419-0.419h3.775V5.806c0-0.23,0.189-0.419,0.419-0.419s0.42,0.189,0.42,0.419v3.775h3.774C14.425,9.581,14.613,9.77,14.613,10 M17.969,10c0,4.401-3.567,7.969-7.969,7.969c-4.402,0-7.969-3.567-7.969-7.969c0-4.402,3.567-7.969,7.969-7.969C14.401,2.031,17.969,5.598,17.969,10 M17.13,10c0-3.932-3.198-7.13-7.13-7.13S2.87,6.068,2.87,10c0,3.933,3.198,7.13,7.13,7.13S17.13,13.933,17.13,10"
                                            }
                                        })
                                    ]
                                ),
                                _vm._v(" "),
                                _c("span", {staticClass: "text-gray-700"}, [
                                    _vm._v("Add Product")
                                ])
                            ]
                        ),
                        _vm._v(" "),
                        _c(
                            "div",
                            {
                                staticClass:
                                    "w-full flex items-center text-blue-400 h-10 pl-4 hover:bg-gray-200 rounded-lg cursor-pointer"
                            },
                            [
                                _c(
                                    "svg",
                                    {
                                        staticClass: "h-6 w-6 fill-current mr-2",
                                        attrs: {viewBox: "0 0 20 20"}
                                    },
                                    [
                                        _c("path", {
                                            attrs: {
                                                d:
                                                    "M15.396,2.292H4.604c-0.212,0-0.385,0.174-0.385,0.386v14.646c0,0.212,0.173,0.385,0.385,0.385h10.792c0.211,0,0.385-0.173,0.385-0.385V2.677C15.781,2.465,15.607,2.292,15.396,2.292 M15.01,16.938H4.99v-2.698h1.609c0.156,0.449,0.586,0.771,1.089,0.771c0.638,0,1.156-0.519,1.156-1.156s-0.519-1.156-1.156-1.156c-0.503,0-0.933,0.321-1.089,0.771H4.99v-3.083h1.609c0.156,0.449,0.586,0.771,1.089,0.771c0.638,0,1.156-0.518,1.156-1.156c0-0.638-0.519-1.156-1.156-1.156c-0.503,0-0.933,0.322-1.089,0.771H4.99V6.531h1.609C6.755,6.98,7.185,7.302,7.688,7.302c0.638,0,1.156-0.519,1.156-1.156c0-0.638-0.519-1.156-1.156-1.156c-0.503,0-0.933,0.322-1.089,0.771H4.99V3.062h10.02V16.938z M7.302,13.854c0-0.212,0.173-0.386,0.385-0.386s0.385,0.174,0.385,0.386s-0.173,0.385-0.385,0.385S7.302,14.066,7.302,13.854 M7.302,10c0-0.212,0.173-0.385,0.385-0.385S8.073,9.788,8.073,10s-0.173,0.385-0.385,0.385S7.302,10.212,7.302,10 M7.302,6.146c0-0.212,0.173-0.386,0.385-0.386s0.385,0.174,0.385,0.386S7.899,6.531,7.688,6.531S7.302,6.358,7.302,6.146"
                                            }
                                        })
                                    ]
                                ),
                                _vm._v(" "),
                                _c("span", {staticClass: "text-gray-700"}, [
                                    _vm._v("View Products")
                                ])
                            ]
                        )
                    ]),
                    _vm._v(" "),
                    _c(
                        "div",
                        {staticClass: "mb-4 px-4"},
                        [
                            _c("p", {staticClass: "pl-4 text-sm font-semibold mb-1"}, [
                                _vm._v("MISC")
                            ]),
                            _vm._v(" "),
                            _c(
                                "div",
                                {
                                    staticClass:
                                        "w-full flex items-center text-blue-400 h-10 pl-4 hover:bg-gray-200 rounded-lg cursor-pointer"
                                },
                                [
                                    _c(
                                        "svg",
                                        {
                                            staticClass: "h-6 w-6 fill-current mr-2",
                                            attrs: {viewBox: "0 0 20 20"}
                                        },
                                        [
                                            _c("path", {
                                                attrs: {
                                                    d:
                                                        "M17.592,8.936l-6.531-6.534c-0.593-0.631-0.751-0.245-0.751,0.056l0.002,2.999L5.427,9.075H2.491c-0.839,0-0.162,0.901-0.311,0.752l3.683,3.678l-3.081,3.108c-0.17,0.171-0.17,0.449,0,0.62c0.169,0.17,0.448,0.17,0.618,0l3.098-3.093l3.675,3.685c-0.099-0.099,0.773,0.474,0.773-0.296v-2.965l3.601-4.872l2.734-0.005C17.73,9.688,18.326,9.669,17.592,8.936 M3.534,9.904h1.906l4.659,4.66v1.906L3.534,9.904z M10.522,13.717L6.287,9.48l4.325-3.124l3.088,3.124L10.522,13.717z M14.335,8.845l-3.177-3.177V3.762l5.083,5.083H14.335z"
                                                }
                                            })
                                        ]
                                    ),
                                    _vm._v(" "),
                                    _c("span", {staticClass: "text-gray-700"}, [_vm._v("Notices")])
                                ]
                            ),
                            _vm._v(" "),
                            _c(
                                "div",
                                {
                                    staticClass:
                                        "w-full flex items-center text-blue-400 h-10 pl-4 hover:bg-gray-200 rounded-lg cursor-pointer"
                                },
                                [
                                    _c(
                                        "svg",
                                        {
                                            staticClass: "h-6 w-6 fill-current mr-2",
                                            attrs: {viewBox: "0 0 20 20"}
                                        },
                                        [
                                            _c("path", {
                                                attrs: {
                                                    d:
                                                        "M6.176,7.241V6.78c0-0.221-0.181-0.402-0.402-0.402c-0.221,0-0.403,0.181-0.403,0.402v0.461C4.79,7.416,4.365,7.955,4.365,8.591c0,0.636,0.424,1.175,1.006,1.35v3.278c0,0.222,0.182,0.402,0.403,0.402c0.222,0,0.402-0.181,0.402-0.402V9.941c0.582-0.175,1.006-0.714,1.006-1.35C7.183,7.955,6.758,7.416,6.176,7.241 M5.774,9.195c-0.332,0-0.604-0.272-0.604-0.604c0-0.332,0.272-0.604,0.604-0.604c0.332,0,0.604,0.272,0.604,0.604C6.377,8.923,6.105,9.195,5.774,9.195 M10.402,10.058V6.78c0-0.221-0.181-0.402-0.402-0.402c-0.222,0-0.402,0.181-0.402,0.402v3.278c-0.582,0.175-1.006,0.714-1.006,1.35c0,0.637,0.424,1.175,1.006,1.351v0.461c0,0.222,0.181,0.402,0.402,0.402c0.221,0,0.402-0.181,0.402-0.402v-0.461c0.582-0.176,1.006-0.714,1.006-1.351C11.408,10.772,10.984,10.233,10.402,10.058M10,12.013c-0.333,0-0.604-0.272-0.604-0.604S9.667,10.805,10,10.805c0.332,0,0.604,0.271,0.604,0.604S10.332,12.013,10,12.013M14.629,8.448V6.78c0-0.221-0.182-0.402-0.403-0.402c-0.221,0-0.402,0.181-0.402,0.402v1.668c-0.581,0.175-1.006,0.714-1.006,1.35c0,0.636,0.425,1.176,1.006,1.351v2.07c0,0.222,0.182,0.402,0.402,0.402c0.222,0,0.403-0.181,0.403-0.402v-2.07c0.581-0.175,1.006-0.715,1.006-1.351C15.635,9.163,15.21,8.624,14.629,8.448 M14.226,10.402c-0.331,0-0.604-0.272-0.604-0.604c0-0.332,0.272-0.604,0.604-0.604c0.332,0,0.604,0.272,0.604,0.604C14.83,10.13,14.558,10.402,14.226,10.402 M17.647,3.962H2.353c-0.221,0-0.402,0.181-0.402,0.402v11.27c0,0.222,0.181,0.402,0.402,0.402h15.295c0.222,0,0.402-0.181,0.402-0.402V4.365C18.05,4.144,17.869,3.962,17.647,3.962 M17.245,15.232H2.755V4.768h14.49V15.232z"
                                                }
                                            })
                                        ]
                                    ),
                                    _vm._v(" "),
                                    _c("span", {staticClass: "text-gray-700"}, [_vm._v("Controls")])
                                ]
                            ),
                            _vm._v(" "),
                            _c(
                                "div",
                                {
                                    staticClass:
                                        "w-full flex items-center text-blue-400 h-10 pl-4 hover:bg-gray-200 rounded-lg cursor-pointer"
                                },
                                [
                                    _c(
                                        "svg",
                                        {
                                            staticClass: "h-6 w-6 fill-current mr-2",
                                            attrs: {viewBox: "0 0 20 20"}
                                        },
                                        [
                                            _c("path", {
                                                attrs: {
                                                    d:
                                                        "M17.498,11.697c-0.453-0.453-0.704-1.055-0.704-1.697c0-0.642,0.251-1.244,0.704-1.697c0.069-0.071,0.15-0.141,0.257-0.22c0.127-0.097,0.181-0.262,0.137-0.417c-0.164-0.558-0.388-1.093-0.662-1.597c-0.075-0.141-0.231-0.22-0.391-0.199c-0.13,0.02-0.238,0.027-0.336,0.027c-1.325,0-2.401-1.076-2.401-2.4c0-0.099,0.008-0.207,0.027-0.336c0.021-0.158-0.059-0.316-0.199-0.391c-0.503-0.274-1.039-0.498-1.597-0.662c-0.154-0.044-0.32,0.01-0.416,0.137c-0.079,0.106-0.148,0.188-0.22,0.257C11.244,2.956,10.643,3.207,10,3.207c-0.642,0-1.244-0.25-1.697-0.704c-0.071-0.069-0.141-0.15-0.22-0.257C7.987,2.119,7.821,2.065,7.667,2.109C7.109,2.275,6.571,2.497,6.07,2.771C5.929,2.846,5.85,3.004,5.871,3.162c0.02,0.129,0.027,0.237,0.027,0.336c0,1.325-1.076,2.4-2.401,2.4c-0.098,0-0.206-0.007-0.335-0.027C3.001,5.851,2.845,5.929,2.77,6.07C2.496,6.572,2.274,7.109,2.108,7.667c-0.044,0.154,0.01,0.32,0.137,0.417c0.106,0.079,0.187,0.148,0.256,0.22c0.938,0.936,0.938,2.458,0,3.394c-0.069,0.072-0.15,0.141-0.256,0.221c-0.127,0.096-0.181,0.262-0.137,0.416c0.166,0.557,0.388,1.096,0.662,1.596c0.075,0.143,0.231,0.221,0.392,0.199c0.129-0.02,0.237-0.027,0.335-0.027c1.325,0,2.401,1.076,2.401,2.402c0,0.098-0.007,0.205-0.027,0.334C5.85,16.996,5.929,17.154,6.07,17.23c0.501,0.273,1.04,0.496,1.597,0.66c0.154,0.047,0.32-0.008,0.417-0.137c0.079-0.105,0.148-0.186,0.22-0.256c0.454-0.453,1.055-0.703,1.697-0.703c0.643,0,1.244,0.25,1.697,0.703c0.071,0.07,0.141,0.15,0.22,0.256c0.073,0.098,0.188,0.152,0.307,0.152c0.036,0,0.073-0.004,0.109-0.016c0.558-0.164,1.096-0.387,1.597-0.66c0.141-0.076,0.22-0.234,0.199-0.393c-0.02-0.129-0.027-0.236-0.027-0.334c0-1.326,1.076-2.402,2.401-2.402c0.098,0,0.206,0.008,0.336,0.027c0.159,0.021,0.315-0.057,0.391-0.199c0.274-0.5,0.496-1.039,0.662-1.596c0.044-0.154-0.01-0.32-0.137-0.416C17.648,11.838,17.567,11.77,17.498,11.697 M16.671,13.334c-0.059-0.002-0.114-0.002-0.168-0.002c-1.749,0-3.173,1.422-3.173,3.172c0,0.053,0.002,0.109,0.004,0.166c-0.312,0.158-0.64,0.295-0.976,0.406c-0.039-0.045-0.077-0.086-0.115-0.123c-0.601-0.6-1.396-0.93-2.243-0.93s-1.643,0.33-2.243,0.93c-0.039,0.037-0.077,0.078-0.116,0.123c-0.336-0.111-0.664-0.248-0.976-0.406c0.002-0.057,0.004-0.113,0.004-0.166c0-1.75-1.423-3.172-3.172-3.172c-0.054,0-0.11,0-0.168,0.002c-0.158-0.312-0.293-0.639-0.405-0.975c0.044-0.039,0.085-0.078,0.124-0.115c1.236-1.236,1.236-3.25,0-4.486C3.009,7.719,2.969,7.68,2.924,7.642c0.112-0.336,0.247-0.664,0.405-0.976C3.387,6.668,3.443,6.67,3.497,6.67c1.75,0,3.172-1.423,3.172-3.172c0-0.054-0.002-0.11-0.004-0.168c0.312-0.158,0.64-0.293,0.976-0.405C7.68,2.969,7.719,3.01,7.757,3.048c0.6,0.6,1.396,0.93,2.243,0.93s1.643-0.33,2.243-0.93c0.038-0.039,0.076-0.079,0.115-0.123c0.336,0.112,0.663,0.247,0.976,0.405c-0.002,0.058-0.004,0.114-0.004,0.168c0,1.749,1.424,3.172,3.173,3.172c0.054,0,0.109-0.002,0.168-0.004c0.158,0.312,0.293,0.64,0.405,0.976c-0.045,0.038-0.086,0.077-0.124,0.116c-0.6,0.6-0.93,1.396-0.93,2.242c0,0.847,0.33,1.645,0.93,2.244c0.038,0.037,0.079,0.076,0.124,0.115C16.964,12.695,16.829,13.021,16.671,13.334 M10,5.417c-2.528,0-4.584,2.056-4.584,4.583c0,2.529,2.056,4.584,4.584,4.584s4.584-2.055,4.584-4.584C14.584,7.472,12.528,5.417,10,5.417 M10,13.812c-2.102,0-3.812-1.709-3.812-3.812c0-2.102,1.71-3.812,3.812-3.812c2.102,0,3.812,1.71,3.812,3.812C13.812,12.104,12.102,13.812,10,13.812"
                                                }
                                            })
                                        ]
                                    ),
                                    _vm._v(" "),
                                    _c("span", {staticClass: "text-gray-700"}, [_vm._v("Settings")])
                                ]
                            ),
                            _vm._v(" "),
                            _c("router-link", {attrs: {to: "/dashboard/user"}}, [
                                _c(
                                    "div",
                                    {
                                        staticClass:
                                            "w-full flex items-center text-blue-400 h-10 pl-4 hover:bg-gray-200 rounded-lg cursor-pointer"
                                    },
                                    [
                                        _c("FontAwesome", {
                                            staticClass: "h-6 w-6 fill-current mr-3",
                                            attrs: {icon: "users"}
                                        }),
                                        _vm._v(" "),
                                        _c("span", {staticClass: "text-gray-700"}, [_vm._v("Users")])
                                    ],
                                    1
                                )
                            ])
                        ],
                        1
                    )
                ]
            )
        }
        var staticRenderFns = []
        render._withStripped = true


        /***/
    }),

    /***/
    "./node_modules/vue-loader/lib/loaders/templateLoader.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/pages/Home.vue?vue&type=template&id=440dff1c&":
    /*!*********************************************************************************************************************************************************************************************************!*\
      !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib??vue-loader-options!./resources/assets/js/pages/Home.vue?vue&type=template&id=440dff1c& ***!
      \*********************************************************************************************************************************************************************************************************/
    /*! exports provided: render, staticRenderFns */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony export (binding) */
        __webpack_require__.d(__webpack_exports__, "render", function () {
            return render;
        });
        /* harmony export (binding) */
        __webpack_require__.d(__webpack_exports__, "staticRenderFns", function () {
            return staticRenderFns;
        });
        var render = function () {
            var _vm = this
            var _h = _vm.$createElement
            var _c = _vm._self._c || _h
            return _c("div", {attrs: {id: "home"}}, [
                _c(
                    "nav",
                    {
                        staticClass: "text-sm font-semibold mb-6",
                        attrs: {"aria-label": "Breadcrumb"}
                    },
                    [
                        _c("ol", {staticClass: "list-none p-0 inline-flex"}, [
                            _c("li", {staticClass: "flex items-center text-blue-500"}, [
                                _c("a", {staticClass: "text-gray-700", attrs: {href: "#"}}, [
                                    _vm._v("Home")
                                ]),
                                _vm._v(" "),
                                _c(
                                    "svg",
                                    {
                                        staticClass: "fill-current w-3 h-3 mx-3",
                                        attrs: {
                                            xmlns: "http://www.w3.org/2000/svg",
                                            viewBox: "0 0 320 512"
                                        }
                                    },
                                    [
                                        _c("path", {
                                            attrs: {
                                                d:
                                                    "M285.476 272.971L91.132 467.314c-9.373 9.373-24.569 9.373-33.941 0l-22.667-22.667c-9.357-9.357-9.375-24.522-.04-33.901L188.505 256 34.484 101.255c-9.335-9.379-9.317-24.544.04-33.901l22.667-22.667c9.373-9.373 24.569-9.373 33.941 0L285.475 239.03c9.373 9.372 9.373 24.568.001 33.941z"
                                            }
                                        })
                                    ]
                                )
                            ]),
                            _vm._v(" "),
                            _vm._m(0)
                        ])
                    ]
                ),
                _vm._v(" "),
                _vm._m(1),
                _vm._v(" "),
                _c("div", {staticClass: "flex flex-wrap -mx-3 mb-20"}, [
                    _c("div", {staticClass: "w-1/2 xl:w-1/4 px-3"}, [
                        _c(
                            "div",
                            {
                                staticClass:
                                    "w-full bg-white border text-blue-400 rounded-lg flex items-center p-6 mb-6 xl:mb-0"
                            },
                            [
                                _c(
                                    "svg",
                                    {
                                        staticClass: "w-16 h-16 fill-current mr-4 hidden lg:block",
                                        attrs: {viewBox: "0 0 20 20"}
                                    },
                                    [
                                        _c("path", {
                                            attrs: {
                                                d:
                                                    "M17.35,2.219h-5.934c-0.115,0-0.225,0.045-0.307,0.128l-8.762,8.762c-0.171,0.168-0.171,0.443,0,0.611l5.933,5.934c0.167,0.171,0.443,0.169,0.612,0l8.762-8.763c0.083-0.083,0.128-0.192,0.128-0.307V2.651C17.781,2.414,17.587,2.219,17.35,2.219M16.916,8.405l-8.332,8.332l-5.321-5.321l8.333-8.332h5.32V8.405z M13.891,4.367c-0.957,0-1.729,0.772-1.729,1.729c0,0.957,0.771,1.729,1.729,1.729s1.729-0.772,1.729-1.729C15.619,5.14,14.848,4.367,13.891,4.367 M14.502,6.708c-0.326,0.326-0.896,0.326-1.223,0c-0.338-0.342-0.338-0.882,0-1.224c0.342-0.337,0.881-0.337,1.223,0C14.84,5.826,14.84,6.366,14.502,6.708"
                                            }
                                        })
                                    ]
                                ),
                                _vm._v(" "),
                                _vm._m(2)
                            ]
                        )
                    ]),
                    _vm._v(" "),
                    _c("div", {staticClass: "w-1/2 xl:w-1/4 px-3"}, [
                        _c(
                            "div",
                            {
                                staticClass:
                                    "w-full bg-white border text-blue-400 rounded-lg flex items-center p-6 mb-6 xl:mb-0"
                            },
                            [
                                _c(
                                    "svg",
                                    {
                                        staticClass: "w-16 h-16 fill-current mr-4 hidden lg:block",
                                        attrs: {viewBox: "0 0 20 20"}
                                    },
                                    [
                                        _c("path", {
                                            attrs: {
                                                d:
                                                    "M17.684,7.925l-5.131-0.67L10.329,2.57c-0.131-0.275-0.527-0.275-0.658,0L7.447,7.255l-5.131,0.67C2.014,7.964,1.892,8.333,2.113,8.54l3.76,3.568L4.924,17.21c-0.056,0.297,0.261,0.525,0.533,0.379L10,15.109l4.543,2.479c0.273,0.153,0.587-0.089,0.533-0.379l-0.949-5.103l3.76-3.568C18.108,8.333,17.986,7.964,17.684,7.925 M13.481,11.723c-0.089,0.083-0.129,0.205-0.105,0.324l0.848,4.547l-4.047-2.208c-0.055-0.03-0.116-0.045-0.176-0.045s-0.122,0.015-0.176,0.045l-4.047,2.208l0.847-4.547c0.023-0.119-0.016-0.241-0.105-0.324L3.162,8.54L7.74,7.941c0.124-0.016,0.229-0.093,0.282-0.203L10,3.568l1.978,4.17c0.053,0.11,0.158,0.187,0.282,0.203l4.578,0.598L13.481,11.723z"
                                            }
                                        })
                                    ]
                                ),
                                _vm._v(" "),
                                _vm._m(3)
                            ]
                        )
                    ]),
                    _vm._v(" "),
                    _c("div", {staticClass: "w-1/2 xl:w-1/4 px-3"}, [
                        _c(
                            "div",
                            {
                                staticClass:
                                    "w-full bg-white border text-blue-400 rounded-lg flex items-center p-6"
                            },
                            [
                                _c(
                                    "svg",
                                    {
                                        staticClass: "w-16 h-16 fill-current mr-4 hidden lg:block",
                                        attrs: {viewBox: "0 0 20 20"}
                                    },
                                    [
                                        _c("path", {
                                            attrs: {
                                                d:
                                                    "M14.999,8.543c0,0.229-0.188,0.417-0.416,0.417H5.417C5.187,8.959,5,8.772,5,8.543s0.188-0.417,0.417-0.417h9.167C14.812,8.126,14.999,8.314,14.999,8.543 M12.037,10.213H5.417C5.187,10.213,5,10.4,5,10.63c0,0.229,0.188,0.416,0.417,0.416h6.621c0.229,0,0.416-0.188,0.416-0.416C12.453,10.4,12.266,10.213,12.037,10.213 M14.583,6.046H5.417C5.187,6.046,5,6.233,5,6.463c0,0.229,0.188,0.417,0.417,0.417h9.167c0.229,0,0.416-0.188,0.416-0.417C14.999,6.233,14.812,6.046,14.583,6.046 M17.916,3.542v10c0,0.229-0.188,0.417-0.417,0.417H9.373l-2.829,2.796c-0.117,0.116-0.71,0.297-0.71-0.296v-2.5H2.5c-0.229,0-0.417-0.188-0.417-0.417v-10c0-0.229,0.188-0.417,0.417-0.417h15C17.729,3.126,17.916,3.313,17.916,3.542 M17.083,3.959H2.917v9.167H6.25c0.229,0,0.417,0.187,0.417,0.416v1.919l2.242-2.215c0.079-0.077,0.184-0.12,0.294-0.12h7.881V3.959z"
                                            }
                                        })
                                    ]
                                ),
                                _vm._v(" "),
                                _vm._m(4)
                            ]
                        )
                    ]),
                    _vm._v(" "),
                    _c("div", {staticClass: "w-1/2 xl:w-1/4 px-3"}, [
                        _c(
                            "div",
                            {
                                staticClass:
                                    "w-full bg-white border text-blue-400 rounded-lg flex items-center p-6"
                            },
                            [
                                _c(
                                    "svg",
                                    {
                                        staticClass: "w-16 h-16 fill-current mr-4 hidden lg:block",
                                        attrs: {viewBox: "0 0 20 20"}
                                    },
                                    [
                                        _c("path", {
                                            attrs: {
                                                d:
                                                    "M17.431,2.156h-3.715c-0.228,0-0.413,0.186-0.413,0.413v6.973h-2.89V6.687c0-0.229-0.186-0.413-0.413-0.413H6.285c-0.228,0-0.413,0.184-0.413,0.413v6.388H2.569c-0.227,0-0.413,0.187-0.413,0.413v3.942c0,0.228,0.186,0.413,0.413,0.413h14.862c0.228,0,0.413-0.186,0.413-0.413V2.569C17.844,2.342,17.658,2.156,17.431,2.156 M5.872,17.019h-2.89v-3.117h2.89V17.019zM9.587,17.019h-2.89V7.1h2.89V17.019z M13.303,17.019h-2.89v-6.651h2.89V17.019z M17.019,17.019h-2.891V2.982h2.891V17.019z"
                                            }
                                        })
                                    ]
                                ),
                                _vm._v(" "),
                                _vm._m(5)
                            ]
                        )
                    ]),
                    _vm._v(" "),
                    _c("section", {staticClass: "text-gray-700 body-font"}, [
                        _c("div", {staticClass: "container px-5 py-24 mx-auto"}, [
                            _vm._m(6),
                            _vm._v(" "),
                            _c("div", {staticClass: "flex flex-wrap -m-4 text-center"}, [
                                _c("div", {staticClass: "p-4 md:w-1/4 sm:w-1/2 w-full"}, [
                                    _c(
                                        "div",
                                        {
                                            staticClass: "border-2 border-gray-200 px-4 py-6 rounded-lg"
                                        },
                                        [
                                            _c(
                                                "svg",
                                                {
                                                    staticClass:
                                                        "text-indigo-500 w-12 h-12 mb-3 inline-block",
                                                    attrs: {
                                                        fill: "none",
                                                        stroke: "currentColor",
                                                        "stroke-linecap": "round",
                                                        "stroke-linejoin": "round",
                                                        "stroke-width": "2",
                                                        viewBox: "0 0 24 24"
                                                    }
                                                },
                                                [
                                                    _c("path", {attrs: {d: "M8 17l4 4 4-4m-4-5v9"}}),
                                                    _vm._v(" "),
                                                    _c("path", {
                                                        attrs: {
                                                            d: "M20.88 18.09A5 5 0 0018 9h-1.26A8 8 0 103 16.29"
                                                        }
                                                    })
                                                ]
                                            ),
                                            _vm._v(" "),
                                            _c(
                                                "h2",
                                                {
                                                    staticClass:
                                                        "title-font font-medium text-3xl text-gray-900"
                                                },
                                                [_vm._v("2.7K")]
                                            ),
                                            _vm._v(" "),
                                            _c("p", {staticClass: "leading-relaxed"}, [
                                                _vm._v("Downloads")
                                            ])
                                        ]
                                    )
                                ]),
                                _vm._v(" "),
                                _c("div", {staticClass: "p-4 md:w-1/4 sm:w-1/2 w-full"}, [
                                    _c(
                                        "div",
                                        {
                                            staticClass: "border-2 border-gray-200 px-4 py-6 rounded-lg"
                                        },
                                        [
                                            _c(
                                                "svg",
                                                {
                                                    staticClass:
                                                        "text-indigo-500 w-12 h-12 mb-3 inline-block",
                                                    attrs: {
                                                        fill: "none",
                                                        stroke: "currentColor",
                                                        "stroke-linecap": "round",
                                                        "stroke-linejoin": "round",
                                                        "stroke-width": "2",
                                                        viewBox: "0 0 24 24"
                                                    }
                                                },
                                                [
                                                    _c("path", {
                                                        attrs: {d: "M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2"}
                                                    }),
                                                    _vm._v(" "),
                                                    _c("circle", {attrs: {cx: "9", cy: "7", r: "4"}}),
                                                    _vm._v(" "),
                                                    _c("path", {
                                                        attrs: {
                                                            d: "M23 21v-2a4 4 0 00-3-3.87m-4-12a4 4 0 010 7.75"
                                                        }
                                                    })
                                                ]
                                            ),
                                            _vm._v(" "),
                                            _c(
                                                "h2",
                                                {
                                                    staticClass:
                                                        "title-font font-medium text-3xl text-gray-900"
                                                },
                                                [_vm._v("1.3K")]
                                            ),
                                            _vm._v(" "),
                                            _c("p", {staticClass: "leading-relaxed"}, [_vm._v("Users")])
                                        ]
                                    )
                                ]),
                                _vm._v(" "),
                                _c("div", {staticClass: "p-4 md:w-1/4 sm:w-1/2 w-full"}, [
                                    _c(
                                        "div",
                                        {
                                            staticClass: "border-2 border-gray-200 px-4 py-6 rounded-lg"
                                        },
                                        [
                                            _c(
                                                "svg",
                                                {
                                                    staticClass:
                                                        "text-indigo-500 w-12 h-12 mb-3 inline-block",
                                                    attrs: {
                                                        fill: "none",
                                                        stroke: "currentColor",
                                                        "stroke-linecap": "round",
                                                        "stroke-linejoin": "round",
                                                        "stroke-width": "2",
                                                        viewBox: "0 0 24 24"
                                                    }
                                                },
                                                [
                                                    _c("path", {attrs: {d: "M3 18v-6a9 9 0 0118 0v6"}}),
                                                    _vm._v(" "),
                                                    _c("path", {
                                                        attrs: {
                                                            d:
                                                                "M21 19a2 2 0 01-2 2h-1a2 2 0 01-2-2v-3a2 2 0 012-2h3zM3 19a2 2 0 002 2h1a2 2 0 002-2v-3a2 2 0 00-2-2H3z"
                                                        }
                                                    })
                                                ]
                                            ),
                                            _vm._v(" "),
                                            _c(
                                                "h2",
                                                {
                                                    staticClass:
                                                        "title-font font-medium text-3xl text-gray-900"
                                                },
                                                [_vm._v("74")]
                                            ),
                                            _vm._v(" "),
                                            _c("p", {staticClass: "leading-relaxed"}, [_vm._v("Files")])
                                        ]
                                    )
                                ]),
                                _vm._v(" "),
                                _c("div", {staticClass: "p-4 md:w-1/4 sm:w-1/2 w-full"}, [
                                    _c(
                                        "div",
                                        {
                                            staticClass: "border-2 border-gray-200 px-4 py-6 rounded-lg"
                                        },
                                        [
                                            _c(
                                                "svg",
                                                {
                                                    staticClass:
                                                        "text-indigo-500 w-12 h-12 mb-3 inline-block",
                                                    attrs: {
                                                        fill: "none",
                                                        stroke: "currentColor",
                                                        "stroke-linecap": "round",
                                                        "stroke-linejoin": "round",
                                                        "stroke-width": "2",
                                                        viewBox: "0 0 24 24"
                                                    }
                                                },
                                                [
                                                    _c("path", {
                                                        attrs: {
                                                            d: "M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"
                                                        }
                                                    })
                                                ]
                                            ),
                                            _vm._v(" "),
                                            _c(
                                                "h2",
                                                {
                                                    staticClass:
                                                        "title-font font-medium text-3xl text-gray-900"
                                                },
                                                [_vm._v("46")]
                                            ),
                                            _vm._v(" "),
                                            _c("p", {staticClass: "leading-relaxed"}, [
                                                _vm._v("Places")
                                            ])
                                        ]
                                    )
                                ])
                            ])
                        ])
                    ])
                ]),
                _vm._v(" "),
                _vm._m(7)
            ])
        }
        var staticRenderFns = [
            function () {
                var _vm = this
                var _h = _vm.$createElement
                var _c = _vm._self._c || _h
                return _c("li", {staticClass: "flex items-center"}, [
                    _c("a", {staticClass: "text-gray-600", attrs: {href: "#"}}, [
                        _vm._v("Dashboard")
                    ])
                ])
            },
            function () {
                var _vm = this
                var _h = _vm.$createElement
                var _c = _vm._self._c || _h
                return _c(
                    "div",
                    {staticClass: "lg:flex justify-between items-center mb-6"},
                    [
                        _c("p", {staticClass: "text-2xl font-semibold mb-2 lg:mb-0"}, [
                            _vm._v("Good afternoon, Joe!")
                        ]),
                        _vm._v(" "),
                        _c(
                            "button",
                            {
                                staticClass:
                                    "bg-blue-500 hover:bg-blue-600 focus:outline-none rounded-lg px-6 py-2 text-white font-semibold shadow"
                            },
                            [_vm._v("\n            View Logs\n        ")]
                        )
                    ]
                )
            },
            function () {
                var _vm = this
                var _h = _vm.$createElement
                var _c = _vm._self._c || _h
                return _c("div", {staticClass: "text-gray-700"}, [
                    _c("p", {staticClass: "font-semibold text-3xl"}, [_vm._v("237")]),
                    _vm._v(" "),
                    _c("p", [_vm._v("Products Sold")])
                ])
            },
            function () {
                var _vm = this
                var _h = _vm.$createElement
                var _c = _vm._self._c || _h
                return _c("div", {staticClass: "text-gray-700"}, [
                    _c("p", {staticClass: "font-semibold text-3xl"}, [_vm._v("177")]),
                    _vm._v(" "),
                    _c("p", [_vm._v("Product Reviews")])
                ])
            },
            function () {
                var _vm = this
                var _h = _vm.$createElement
                var _c = _vm._self._c || _h
                return _c("div", {staticClass: "text-gray-700"}, [
                    _c("p", {staticClass: "font-semibold text-3xl"}, [_vm._v("31")]),
                    _vm._v(" "),
                    _c("p", [_vm._v("New Enquiries")])
                ])
            },
            function () {
                var _vm = this
                var _h = _vm.$createElement
                var _c = _vm._self._c || _h
                return _c("div", {staticClass: "text-gray-700"}, [
                    _c("p", {staticClass: "font-semibold text-3xl"}, [_vm._v("1,653")]),
                    _vm._v(" "),
                    _c("p", [_vm._v("Product Views")])
                ])
            },
            function () {
                var _vm = this
                var _h = _vm.$createElement
                var _c = _vm._self._c || _h
                return _c(
                    "div",
                    {staticClass: "flex flex-col text-center w-full mb-20"},
                    [
                        _c(
                            "h1",
                            {
                                staticClass:
                                    "sm:text-3xl text-2xl font-medium title-font mb-4 text-gray-900"
                            },
                            [_vm._v("Master Cleanse Reliac Heirloom")]
                        ),
                        _vm._v(" "),
                        _c("p", {staticClass: "lg:w-2/3 mx-auto leading-relaxed text-base"}, [
                            _vm._v(
                                "Whatever cardigan tote bag tumblr hexagon brooklyn asymmetrical gentrify, subway tile poke farm-to-table. Franzen you probably haven't heard of them man bun deep jianbing selfies heirloom prism food truck ugh squid celiac humblebrag."
                            )
                        ])
                    ]
                )
            },
            function () {
                var _vm = this
                var _h = _vm.$createElement
                var _c = _vm._self._c || _h
                return _c("div", {staticClass: "flex flex-wrap -mx-3"}, [
                    _c("div", {staticClass: "w-full xl:w-1/3 px-3"}, [
                        _c("p", {staticClass: "text-xl font-semibold mb-4"}, [
                            _vm._v("Recent Sales")
                        ]),
                        _vm._v(" "),
                        _c(
                            "div",
                            {staticClass: "w-full bg-white border rounded-lg p-4 mb-8 xl:mb-0"},
                            [
                                _c("canvas", {
                                    attrs: {id: "buyers-chart", width: "600", height: "400"}
                                })
                            ]
                        )
                    ]),
                    _vm._v(" "),
                    _c("div", {staticClass: "w-full xl:w-1/3 px-3"}, [
                        _c("p", {staticClass: "text-xl font-semibold mb-4"}, [
                            _vm._v("Recent Reviews")
                        ]),
                        _vm._v(" "),
                        _c(
                            "div",
                            {staticClass: "w-full bg-white border rounded-lg p-4 mb-8 xl:mb-0"},
                            [
                                _c("canvas", {
                                    attrs: {id: "reviews-chart", width: "600", height: "400"}
                                })
                            ]
                        )
                    ]),
                    _vm._v(" "),
                    _c("div", {staticClass: "w-full xl:w-1/3 px-3"}, [
                        _c("p", {staticClass: "text-xl font-semibold mb-4"}, [
                            _vm._v("Recent Transactions")
                        ]),
                        _vm._v(" "),
                        _c("div", {staticClass: "w-full bg-white border rounded-lg p-4"}, [
                            _c(
                                "div",
                                {
                                    staticClass:
                                        "w-full bg-gray-100 border rounded-lg flex justify-between items-center px-4 py-2 mb-4"
                                },
                                [
                                    _c("div", [
                                        _c("p", {staticClass: "font-semibold text-xl"}, [
                                            _vm._v("Trent Murphy")
                                        ]),
                                        _vm._v(" "),
                                        _c("p", [_vm._v("Product 1")])
                                    ]),
                                    _vm._v(" "),
                                    _c(
                                        "span",
                                        {staticClass: "text-green-500 font-semibold text-lg"},
                                        [_vm._v("$25.00")]
                                    )
                                ]
                            ),
                            _vm._v(" "),
                            _c(
                                "div",
                                {
                                    staticClass:
                                        "w-full bg-gray-100 border rounded-lg flex justify-between items-center px-4 py-2 mb-4"
                                },
                                [
                                    _c("div", [
                                        _c("p", {staticClass: "font-semibold text-xl"}, [
                                            _vm._v("Joseph Brent")
                                        ]),
                                        _vm._v(" "),
                                        _c("p", [_vm._v("Product 34")])
                                    ]),
                                    _vm._v(" "),
                                    _c(
                                        "span",
                                        {staticClass: "text-red-500 font-semibold text-lg"},
                                        [_vm._v("$74.99")]
                                    )
                                ]
                            ),
                            _vm._v(" "),
                            _c(
                                "div",
                                {
                                    staticClass:
                                        "w-full bg-gray-100 border rounded-lg flex justify-between items-center px-4 py-2 mb-4"
                                },
                                [
                                    _c("div", [
                                        _c("p", {staticClass: "font-semibold text-xl"}, [
                                            _vm._v("Jacob Bator")
                                        ]),
                                        _vm._v(" "),
                                        _c("p", [_vm._v("Product 23")])
                                    ]),
                                    _vm._v(" "),
                                    _c(
                                        "span",
                                        {staticClass: "text-green-500 font-semibold text-lg"},
                                        [_vm._v("$14.95")]
                                    )
                                ]
                            ),
                            _vm._v(" "),
                            _c(
                                "div",
                                {
                                    staticClass:
                                        "w-full bg-gray-100 border rounded-lg flex justify-between items-center px-4 py-2"
                                },
                                [
                                    _c("div", [
                                        _c("p", {staticClass: "font-semibold text-xl"}, [
                                            _vm._v("Alex Mason")
                                        ]),
                                        _vm._v(" "),
                                        _c("p", [_vm._v("Product 66")])
                                    ]),
                                    _vm._v(" "),
                                    _c(
                                        "span",
                                        {staticClass: "text-green-500 font-semibold text-lg"},
                                        [_vm._v("$44.99")]
                                    )
                                ]
                            )
                        ])
                    ])
                ])
            }
        ]
        render._withStripped = true


        /***/
    }),

    /***/
    "./resources/assets/js/components/Dashboard.vue":
    /*!******************************************************!*\
      !*** ./resources/assets/js/components/Dashboard.vue ***!
      \******************************************************/
    /*! exports provided: default */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var _Dashboard_vue_vue_type_template_id_1f65406d___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Dashboard.vue?vue&type=template&id=1f65406d& */ "./resources/assets/js/components/Dashboard.vue?vue&type=template&id=1f65406d&");
        /* harmony import */
        var _Dashboard_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Dashboard.vue?vue&type=script&lang=js& */ "./resources/assets/js/components/Dashboard.vue?vue&type=script&lang=js&");
        /* empty/unused harmony star reexport *//* harmony import */
        var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");


        /* normalize component */

        var component = Object(_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__["default"])(
            _Dashboard_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__["default"],
            _Dashboard_vue_vue_type_template_id_1f65406d___WEBPACK_IMPORTED_MODULE_0__["render"],
            _Dashboard_vue_vue_type_template_id_1f65406d___WEBPACK_IMPORTED_MODULE_0__["staticRenderFns"],
            false,
            null,
            null,
            null
        )

        /* hot reload */
        if (false) {
            var api;
        }
        component.options.__file = "resources/assets/js/components/Dashboard.vue"
        /* harmony default export */
        __webpack_exports__["default"] = (component.exports);

        /***/
    }),

    /***/
    "./resources/assets/js/components/Dashboard.vue?vue&type=script&lang=js&":
    /*!*******************************************************************************!*\
      !*** ./resources/assets/js/components/Dashboard.vue?vue&type=script&lang=js& ***!
      \*******************************************************************************/
    /*! exports provided: default */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var _node_modules_babel_loader_lib_index_js_ref_4_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Dashboard_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/babel-loader/lib??ref--4-0!../../../../node_modules/vue-loader/lib??vue-loader-options!./Dashboard.vue?vue&type=script&lang=js& */ "./node_modules/babel-loader/lib/index.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/components/Dashboard.vue?vue&type=script&lang=js&");
        /* empty/unused harmony star reexport */ /* harmony default export */
        __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_ref_4_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Dashboard_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__["default"]);

        /***/
    }),

    /***/
    "./resources/assets/js/components/Dashboard.vue?vue&type=template&id=1f65406d&":
    /*!*************************************************************************************!*\
      !*** ./resources/assets/js/components/Dashboard.vue?vue&type=template&id=1f65406d& ***!
      \*************************************************************************************/
    /*! exports provided: render, staticRenderFns */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Dashboard_vue_vue_type_template_id_1f65406d___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../node_modules/vue-loader/lib??vue-loader-options!./Dashboard.vue?vue&type=template&id=1f65406d& */ "./node_modules/vue-loader/lib/loaders/templateLoader.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/components/Dashboard.vue?vue&type=template&id=1f65406d&");
        /* harmony reexport (safe) */
        __webpack_require__.d(__webpack_exports__, "render", function () {
            return _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Dashboard_vue_vue_type_template_id_1f65406d___WEBPACK_IMPORTED_MODULE_0__["render"];
        });

        /* harmony reexport (safe) */
        __webpack_require__.d(__webpack_exports__, "staticRenderFns", function () {
            return _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Dashboard_vue_vue_type_template_id_1f65406d___WEBPACK_IMPORTED_MODULE_0__["staticRenderFns"];
        });


        /***/
    }),

    /***/
    "./resources/assets/js/components/Footer.vue":
    /*!***************************************************!*\
      !*** ./resources/assets/js/components/Footer.vue ***!
      \***************************************************/
    /*! exports provided: default */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var _Footer_vue_vue_type_template_id_083ff5dc___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Footer.vue?vue&type=template&id=083ff5dc& */ "./resources/assets/js/components/Footer.vue?vue&type=template&id=083ff5dc&");
        /* harmony import */
        var _Footer_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Footer.vue?vue&type=script&lang=js& */ "./resources/assets/js/components/Footer.vue?vue&type=script&lang=js&");
        /* empty/unused harmony star reexport *//* harmony import */
        var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");


        /* normalize component */

        var component = Object(_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__["default"])(
            _Footer_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__["default"],
            _Footer_vue_vue_type_template_id_083ff5dc___WEBPACK_IMPORTED_MODULE_0__["render"],
            _Footer_vue_vue_type_template_id_083ff5dc___WEBPACK_IMPORTED_MODULE_0__["staticRenderFns"],
            false,
            null,
            null,
            null
        )

        /* hot reload */
        if (false) {
            var api;
        }
        component.options.__file = "resources/assets/js/components/Footer.vue"
        /* harmony default export */
        __webpack_exports__["default"] = (component.exports);

        /***/
    }),

    /***/
    "./resources/assets/js/components/Footer.vue?vue&type=script&lang=js&":
    /*!****************************************************************************!*\
      !*** ./resources/assets/js/components/Footer.vue?vue&type=script&lang=js& ***!
      \****************************************************************************/
    /*! exports provided: default */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var _node_modules_babel_loader_lib_index_js_ref_4_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Footer_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/babel-loader/lib??ref--4-0!../../../../node_modules/vue-loader/lib??vue-loader-options!./Footer.vue?vue&type=script&lang=js& */ "./node_modules/babel-loader/lib/index.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/components/Footer.vue?vue&type=script&lang=js&");
        /* empty/unused harmony star reexport */ /* harmony default export */
        __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_ref_4_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Footer_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__["default"]);

        /***/
    }),

    /***/
    "./resources/assets/js/components/Footer.vue?vue&type=template&id=083ff5dc&":
    /*!**********************************************************************************!*\
      !*** ./resources/assets/js/components/Footer.vue?vue&type=template&id=083ff5dc& ***!
      \**********************************************************************************/
    /*! exports provided: render, staticRenderFns */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Footer_vue_vue_type_template_id_083ff5dc___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../node_modules/vue-loader/lib??vue-loader-options!./Footer.vue?vue&type=template&id=083ff5dc& */ "./node_modules/vue-loader/lib/loaders/templateLoader.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/components/Footer.vue?vue&type=template&id=083ff5dc&");
        /* harmony reexport (safe) */
        __webpack_require__.d(__webpack_exports__, "render", function () {
            return _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Footer_vue_vue_type_template_id_083ff5dc___WEBPACK_IMPORTED_MODULE_0__["render"];
        });

        /* harmony reexport (safe) */
        __webpack_require__.d(__webpack_exports__, "staticRenderFns", function () {
            return _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Footer_vue_vue_type_template_id_083ff5dc___WEBPACK_IMPORTED_MODULE_0__["staticRenderFns"];
        });


        /***/
    }),

    /***/
    "./resources/assets/js/components/Navbar.vue":
    /*!***************************************************!*\
      !*** ./resources/assets/js/components/Navbar.vue ***!
      \***************************************************/
    /*! exports provided: default */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var _Navbar_vue_vue_type_template_id_cadbadf2___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Navbar.vue?vue&type=template&id=cadbadf2& */ "./resources/assets/js/components/Navbar.vue?vue&type=template&id=cadbadf2&");
        /* harmony import */
        var _Navbar_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Navbar.vue?vue&type=script&lang=js& */ "./resources/assets/js/components/Navbar.vue?vue&type=script&lang=js&");
        /* empty/unused harmony star reexport *//* harmony import */
        var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");


        /* normalize component */

        var component = Object(_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__["default"])(
            _Navbar_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__["default"],
            _Navbar_vue_vue_type_template_id_cadbadf2___WEBPACK_IMPORTED_MODULE_0__["render"],
            _Navbar_vue_vue_type_template_id_cadbadf2___WEBPACK_IMPORTED_MODULE_0__["staticRenderFns"],
            false,
            null,
            null,
            null
        )

        /* hot reload */
        if (false) {
            var api;
        }
        component.options.__file = "resources/assets/js/components/Navbar.vue"
        /* harmony default export */
        __webpack_exports__["default"] = (component.exports);

        /***/
    }),

    /***/
    "./resources/assets/js/components/Navbar.vue?vue&type=script&lang=js&":
    /*!****************************************************************************!*\
      !*** ./resources/assets/js/components/Navbar.vue?vue&type=script&lang=js& ***!
      \****************************************************************************/
    /*! exports provided: default */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var _node_modules_babel_loader_lib_index_js_ref_4_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Navbar_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/babel-loader/lib??ref--4-0!../../../../node_modules/vue-loader/lib??vue-loader-options!./Navbar.vue?vue&type=script&lang=js& */ "./node_modules/babel-loader/lib/index.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/components/Navbar.vue?vue&type=script&lang=js&");
        /* empty/unused harmony star reexport */ /* harmony default export */
        __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_ref_4_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Navbar_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__["default"]);

        /***/
    }),

    /***/
    "./resources/assets/js/components/Navbar.vue?vue&type=template&id=cadbadf2&":
    /*!**********************************************************************************!*\
      !*** ./resources/assets/js/components/Navbar.vue?vue&type=template&id=cadbadf2& ***!
      \**********************************************************************************/
    /*! exports provided: render, staticRenderFns */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Navbar_vue_vue_type_template_id_cadbadf2___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../node_modules/vue-loader/lib??vue-loader-options!./Navbar.vue?vue&type=template&id=cadbadf2& */ "./node_modules/vue-loader/lib/loaders/templateLoader.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/components/Navbar.vue?vue&type=template&id=cadbadf2&");
        /* harmony reexport (safe) */
        __webpack_require__.d(__webpack_exports__, "render", function () {
            return _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Navbar_vue_vue_type_template_id_cadbadf2___WEBPACK_IMPORTED_MODULE_0__["render"];
        });

        /* harmony reexport (safe) */
        __webpack_require__.d(__webpack_exports__, "staticRenderFns", function () {
            return _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Navbar_vue_vue_type_template_id_cadbadf2___WEBPACK_IMPORTED_MODULE_0__["staticRenderFns"];
        });


        /***/
    }),

    /***/
    "./resources/assets/js/components/Sidebar.vue":
    /*!****************************************************!*\
      !*** ./resources/assets/js/components/Sidebar.vue ***!
      \****************************************************/
    /*! exports provided: default */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var _Sidebar_vue_vue_type_template_id_28cb1975___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Sidebar.vue?vue&type=template&id=28cb1975& */ "./resources/assets/js/components/Sidebar.vue?vue&type=template&id=28cb1975&");
        /* harmony import */
        var _Sidebar_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Sidebar.vue?vue&type=script&lang=js& */ "./resources/assets/js/components/Sidebar.vue?vue&type=script&lang=js&");
        /* empty/unused harmony star reexport *//* harmony import */
        var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");


        /* normalize component */

        var component = Object(_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__["default"])(
            _Sidebar_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__["default"],
            _Sidebar_vue_vue_type_template_id_28cb1975___WEBPACK_IMPORTED_MODULE_0__["render"],
            _Sidebar_vue_vue_type_template_id_28cb1975___WEBPACK_IMPORTED_MODULE_0__["staticRenderFns"],
            false,
            null,
            null,
            null
        )

        /* hot reload */
        if (false) {
            var api;
        }
        component.options.__file = "resources/assets/js/components/Sidebar.vue"
        /* harmony default export */
        __webpack_exports__["default"] = (component.exports);

        /***/
    }),

    /***/
    "./resources/assets/js/components/Sidebar.vue?vue&type=script&lang=js&":
    /*!*****************************************************************************!*\
      !*** ./resources/assets/js/components/Sidebar.vue?vue&type=script&lang=js& ***!
      \*****************************************************************************/
    /*! exports provided: default */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var _node_modules_babel_loader_lib_index_js_ref_4_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Sidebar_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/babel-loader/lib??ref--4-0!../../../../node_modules/vue-loader/lib??vue-loader-options!./Sidebar.vue?vue&type=script&lang=js& */ "./node_modules/babel-loader/lib/index.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/components/Sidebar.vue?vue&type=script&lang=js&");
        /* empty/unused harmony star reexport */ /* harmony default export */
        __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_ref_4_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Sidebar_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__["default"]);

        /***/
    }),

    /***/
    "./resources/assets/js/components/Sidebar.vue?vue&type=template&id=28cb1975&":
    /*!***********************************************************************************!*\
      !*** ./resources/assets/js/components/Sidebar.vue?vue&type=template&id=28cb1975& ***!
      \***********************************************************************************/
    /*! exports provided: render, staticRenderFns */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Sidebar_vue_vue_type_template_id_28cb1975___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../node_modules/vue-loader/lib??vue-loader-options!./Sidebar.vue?vue&type=template&id=28cb1975& */ "./node_modules/vue-loader/lib/loaders/templateLoader.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/components/Sidebar.vue?vue&type=template&id=28cb1975&");
        /* harmony reexport (safe) */
        __webpack_require__.d(__webpack_exports__, "render", function () {
            return _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Sidebar_vue_vue_type_template_id_28cb1975___WEBPACK_IMPORTED_MODULE_0__["render"];
        });

        /* harmony reexport (safe) */
        __webpack_require__.d(__webpack_exports__, "staticRenderFns", function () {
            return _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Sidebar_vue_vue_type_template_id_28cb1975___WEBPACK_IMPORTED_MODULE_0__["staticRenderFns"];
        });


        /***/
    }),

    /***/
    "./resources/assets/js/pages/Home.vue":
    /*!********************************************!*\
      !*** ./resources/assets/js/pages/Home.vue ***!
      \********************************************/
    /*! exports provided: default */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var _Home_vue_vue_type_template_id_440dff1c___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Home.vue?vue&type=template&id=440dff1c& */ "./resources/assets/js/pages/Home.vue?vue&type=template&id=440dff1c&");
        /* harmony import */
        var _Home_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Home.vue?vue&type=script&lang=js& */ "./resources/assets/js/pages/Home.vue?vue&type=script&lang=js&");
        /* empty/unused harmony star reexport *//* harmony import */
        var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");


        /* normalize component */

        var component = Object(_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__["default"])(
            _Home_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__["default"],
            _Home_vue_vue_type_template_id_440dff1c___WEBPACK_IMPORTED_MODULE_0__["render"],
            _Home_vue_vue_type_template_id_440dff1c___WEBPACK_IMPORTED_MODULE_0__["staticRenderFns"],
            false,
            null,
            null,
            null
        )

        /* hot reload */
        if (false) {
            var api;
        }
        component.options.__file = "resources/assets/js/pages/Home.vue"
        /* harmony default export */
        __webpack_exports__["default"] = (component.exports);

        /***/
    }),

    /***/
    "./resources/assets/js/pages/Home.vue?vue&type=script&lang=js&":
    /*!*********************************************************************!*\
      !*** ./resources/assets/js/pages/Home.vue?vue&type=script&lang=js& ***!
      \*********************************************************************/
    /*! exports provided: default */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var _node_modules_babel_loader_lib_index_js_ref_4_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Home_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/babel-loader/lib??ref--4-0!../../../../node_modules/vue-loader/lib??vue-loader-options!./Home.vue?vue&type=script&lang=js& */ "./node_modules/babel-loader/lib/index.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/pages/Home.vue?vue&type=script&lang=js&");
        /* empty/unused harmony star reexport */ /* harmony default export */
        __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_ref_4_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Home_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__["default"]);

        /***/
    }),

    /***/
    "./resources/assets/js/pages/Home.vue?vue&type=template&id=440dff1c&":
    /*!***************************************************************************!*\
      !*** ./resources/assets/js/pages/Home.vue?vue&type=template&id=440dff1c& ***!
      \***************************************************************************/
    /*! exports provided: render, staticRenderFns */
    /***/ (function (module, __webpack_exports__, __webpack_require__) {

        "use strict";
        __webpack_require__.r(__webpack_exports__);
        /* harmony import */
        var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Home_vue_vue_type_template_id_440dff1c___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../node_modules/vue-loader/lib??vue-loader-options!./Home.vue?vue&type=template&id=440dff1c& */ "./node_modules/vue-loader/lib/loaders/templateLoader.js?!./node_modules/vue-loader/lib/index.js?!./resources/assets/js/pages/Home.vue?vue&type=template&id=440dff1c&");
        /* harmony reexport (safe) */
        __webpack_require__.d(__webpack_exports__, "render", function () {
            return _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Home_vue_vue_type_template_id_440dff1c___WEBPACK_IMPORTED_MODULE_0__["render"];
        });

        /* harmony reexport (safe) */
        __webpack_require__.d(__webpack_exports__, "staticRenderFns", function () {
            return _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Home_vue_vue_type_template_id_440dff1c___WEBPACK_IMPORTED_MODULE_0__["staticRenderFns"];
        });


        /***/
    })

}]);
!function(){function a(a){global.setImmediate?setImmediate(a):global.importScripts?setTimeout(a):(c++,d[c]=a,global.postMessage(c,"*"))}function b(c){"use strict";function d(a,b,c,d){if("object"!=typeof j&&"function"!=typeof j||"function"!=typeof a)d();else try{var e=0;a.call(j,function(a){e++||(j=a,b())},function(a){e++||(j=a,c())})}catch(f){j=f,c()}}function e(){var a;try{a=j&&j.then}catch(b){return j=b,i=2,e()}d(a,function(){i=1,e()},function(){i=2,e()},function(){try{1==i&&"function"==typeof f?j=f(j):2==i&&"function"==typeof g&&(j=g(j),i=1)}catch(b){return j=b,l()}j==h?(j=TypeError(),l()):d(a,function(){l(3)},l,function(){l(1==i&&3)})})}if("function"!=typeof c&&void 0!=c)throw TypeError();if("object"!=typeof this||this&&this.then)throw TypeError();var f,g,h=this,i=0,j=0,k=[];h.promise=h,h.resolve=function(b){return f=h.fn,g=h.er,i||(j=b,i=1,a(e)),h},h.reject=function(b){return f=h.fn,g=h.er,i||(j=b,i=2,a(e)),h},h._d=1,h.then=function(a,c){if(1!=this._d)throw TypeError();var d=new b;return d.fn=a,d.er=c,3==i?d.resolve(j):4==i?d.reject(j):k.push(d),d},h["catch"]=function(a){return h.then(null,a)};var l=function(a){i=a||4,k.map(function(a){3==i&&a.resolve(j)||a.reject(j)})};try{"function"==typeof c&&c(h.resolve,h.reject)}catch(m){h.reject(m)}return h}global=this;var c=1,d={},e=!1;global.setImmediate||global.addEventListener("message",function(b){if(b.source==global)if(e)a(d[b.data]);else{e=!0;try{d[b.data]()}catch(b){}delete d[b.data],e=!1}}),b.resolve=function(a){if(1!=this._d)throw TypeError();return a instanceof b?a:new b(function(b){b(a)})},b.reject=function(a){if(1!=this._d)throw TypeError();return new b(function(b,c){c(a)})},b.all=function(a){function c(b,e){if(e)return d.resolve(e);if(b)return d.reject(b);var f=a.reduce(function(a,b){return b&&b.then?a+1:a},0);0==f&&d.resolve(a),a.map(function(b,d){b&&b.then&&b.then(function(b){return a[d]=b,c(),b},c)})}if(1!=this._d)throw TypeError();if(!(a instanceof Array))return b.reject(TypeError());var d=new b;return c(),d},b.race=function(a){function c(b,e){if(e)return d.resolve(e);if(b)return d.reject(b);var f=a.reduce(function(a,b){return b&&b.then?a+1:a},0);0==f&&d.resolve(a),a.map(function(a,b){a&&a.then&&a.then(function(a){c(null,a)},c)})}if(1!=this._d)throw TypeError();if(!(a instanceof Array))return b.reject(TypeError());if(0==a.length)return new b;var d=new b;return c(),d},b._d=1,"undefined"!=typeof module?module.exports=b:global.Promise=global.Promise||b}();
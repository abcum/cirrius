/*!
 * Copyright (c) 2014 Chris O'Hara <cohara87@gmail.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining
 * a copy of this software and associated documentation files (the
 * "Software"), to deal in the Software without restriction, including
 * without limitation the rights to use, copy, modify, merge, publish,
 * distribute, sublicense, and/or sell copies of the Software, and to
 * permit persons to whom the Software is furnished to do so, subject to
 * the following conditions:
 *
 * The above copyright notice and this permission notice shall be
 * included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
 * LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
 * OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
 * WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */
!function(t,e){"undefined"!=typeof module?module.exports=e():"function"==typeof define&&"object"==typeof define.amd?define(e):this[t]=e()}("validator",function(t){"use strict";function e(t,e){t=t||{};for(var u in e)"undefined"==typeof t[u]&&(t[u]=e[u]);return t}t={version:"3.16.1"};var u=/^((([a-z]|\d|[!#\$%&'\*\+\-\/=\?\^_`{\|}~]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])+(\.([a-z]|\d|[!#\$%&'\*\+\-\/=\?\^_`{\|}~]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])+)*)|((\x22)((((\x20|\x09)*(\x0d\x0a))?(\x20|\x09)+)?(([\x01-\x08\x0b\x0c\x0e-\x1f\x7f]|\x21|[\x23-\x5b]|[\x5d-\x7e]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(\\([\x01-\x09\x0b\x0c\x0d-\x7f]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]))))*(((\x20|\x09)*(\x0d\x0a))?(\x20|\x09)+)?(\x22)))@((([a-z]|\d|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(([a-z]|\d|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])*([a-z]|\d|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])))\.)+(([a-z]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(([a-z]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])*([a-z]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])))\.?$/i,n=/^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$/,r=/^(?:[0-9]{9}X|[0-9]{10})$/,F=/^(?:[0-9]{13})$/,i=/^(\d?\d?\d)\.(\d?\d?\d)\.(\d?\d?\d)\.(\d?\d?\d)$/,o=/^::|^::1|^([a-fA-F0-9]{1,4}::?){1,7}([a-fA-F0-9]{1,4})$/,a={3:/^[0-9A-F]{8}-[0-9A-F]{4}-3[0-9A-F]{3}-[0-9A-F]{4}-[0-9A-F]{12}$/i,4:/^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i,5:/^[0-9A-F]{8}-[0-9A-F]{4}-5[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i,all:/^[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}$/i},f=/^[a-zA-Z]+$/,s=/^[a-zA-Z0-9]+$/,c=/^-?[0-9]+$/,l=/^(?:-?(?:0|[1-9][0-9]*))$/,d=/^(?:-?(?:[0-9]+))?(?:\.[0-9]*)?(?:[eE][\+\-]?(?:[0-9]+))?$/,p=/^[0-9a-fA-F]+$/,D=/^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$/,A=/^[\x00-\x7F]+$/,x=/[^\x00-\x7F]/,g=/[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]/,h=/[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]/,E=/[\uD800-\uDBFF][\uDC00-\uDFFF]/,$=/^(?:[A-Za-z0-9+\/]{4})*(?:[A-Za-z0-9+\/]{2}==|[A-Za-z0-9+\/]{3}=|[A-Za-z0-9+\/]{4})$/;t.extend=function(e,u){t[e]=function(){var e=Array.prototype.slice.call(arguments);return e[0]=t.toString(e[0]),u.apply(t,e)}},t.init=function(){for(var e in t)"function"==typeof t[e]&&"toString"!==e&&"toDate"!==e&&"extend"!==e&&"init"!==e&&t.extend(e,t[e])},t.toString=function(t){return"object"==typeof t&&null!==t&&t.toString?t=t.toString():null===t||"undefined"==typeof t||isNaN(t)&&!t.length?t="":"string"!=typeof t&&(t+=""),t},t.toDate=function(t){return"[object Date]"===Object.prototype.toString.call(t)?t:(t=Date.parse(t),isNaN(t)?null:new Date(t))},t.toFloat=function(t){return parseFloat(t)},t.toInt=function(t,e){return parseInt(t,e||10)},t.toBoolean=function(t,e){return e?"1"===t||"true"===t:"0"!==t&&"false"!==t&&""!==t},t.equals=function(e,u){return e===t.toString(u)},t.contains=function(e,u){return e.indexOf(t.toString(u))>=0},t.matches=function(t,e,u){return"[object RegExp]"!==Object.prototype.toString.call(e)&&(e=new RegExp(e,u)),e.test(t)},t.isEmail=function(t){return u.test(t)};var v={protocols:["http","https","ftp"],require_tld:!0,require_protocol:!1,allow_underscores:!1};return t.isURL=function(t,u){if(!t||t.length>=2083)return!1;u=e(u,v);var n="-?-?"+(u.allow_underscores?"_?":""),r=new RegExp("^(?!mailto:)(?:(?:"+u.protocols.join("|")+")://)"+(u.require_protocol?"":"?")+"(?:\\S+(?::\\S*)?@)?(?:(?:(?:[1-9]\\d?|1\\d\\d|2[01]\\d|22[0-3])(?:\\.(?:1?\\d{1,2}|2[0-4]\\d|25[0-5])){2}(?:\\.(?:[0-9]\\d?|1\\d\\d|2[0-4]\\d|25[0-4]))|(?:(?:www.)?)?(?:(?:[a-z\\u00a1-\\uffff0-9]+"+n+")*[a-z\\u00a1-\\uffff0-9]+)(?:\\.(?:[a-z\\u00a1-\\uffff0-9]+"+n+")*[a-z\\u00a1-\\uffff0-9]+)*(?:\\.(?:[a-z\\u00a1-\\uffff]{2,}))"+(u.require_tld?"":"?")+")|localhost)(?::(\\d{1,5}))?(?:(?:/|\\?|#)[^\\s]*)?$","i"),F=t.match(r),i=F?F[1]:0;return!(!F||i&&!(i>0&&65535>=i))},t.isIP=function(e,u){if(u=t.toString(u),!u)return t.isIP(e,4)||t.isIP(e,6);if("4"===u){if(!i.test(e))return!1;var n=e.split(".").sort();return n[3]<=255}return"6"===u&&o.test(e)},t.isAlpha=function(t){return f.test(t)},t.isAlphanumeric=function(t){return s.test(t)},t.isNumeric=function(t){return c.test(t)},t.isHexadecimal=function(t){return p.test(t)},t.isHexColor=function(t){return D.test(t)},t.isLowercase=function(t){return t===t.toLowerCase()},t.isUppercase=function(t){return t===t.toUpperCase()},t.isInt=function(t){return l.test(t)},t.isFloat=function(t){return""!==t&&d.test(t)},t.isDivisibleBy=function(e,u){return t.toFloat(e)%t.toInt(u)===0},t.isNull=function(t){return 0===t.length},t.isLength=function(t,e,u){var n=t.match(/[\uD800-\uDBFF][\uDC00-\uDFFF]/g)||[],r=t.length-n.length;return r>=e&&("undefined"==typeof u||u>=r)},t.isByteLength=function(t,e,u){return t.length>=e&&("undefined"==typeof u||t.length<=u)},t.isUUID=function(t,e){var u=a[e?e:"all"];return u&&u.test(t)},t.isDate=function(t){return!isNaN(Date.parse(t))},t.isAfter=function(e,u){var n=t.toDate(u||new Date),r=t.toDate(e);return!!(r&&n&&r>n)},t.isBefore=function(e,u){var n=t.toDate(u||new Date),r=t.toDate(e);return r&&n&&n>r},t.isIn=function(e,u){if(!u||"function"!=typeof u.indexOf)return!1;if("[object Array]"===Object.prototype.toString.call(u)){for(var n=[],r=0,F=u.length;F>r;r++)n[r]=t.toString(u[r]);u=n}return u.indexOf(e)>=0},t.isCreditCard=function(t){var e=t.replace(/[^0-9]+/g,"");if(!n.test(e))return!1;for(var u,r,F,i=0,o=e.length-1;o>=0;o--)u=e.substring(o,o+1),r=parseInt(u,10),F?(r*=2,i+=r>=10?r%10+1:r):i+=r,F=!F;return!!(i%10===0?e:!1)},t.isISBN=function(e,u){if(u=t.toString(u),!u)return t.isISBN(e,10)||t.isISBN(e,13);var n,i=e.replace(/[\s-]+/g,""),o=0;if("10"===u){if(!r.test(i))return!1;for(n=0;9>n;n++)o+=(n+1)*i.charAt(n);if(o+="X"===i.charAt(9)?100:10*i.charAt(9),o%11===0)return!!i}else if("13"===u){if(!F.test(i))return!1;var a=[1,3];for(n=0;12>n;n++)o+=a[n%2]*i.charAt(n);if(i.charAt(12)-(10-o%10)%10===0)return!!i}return!1},t.isJSON=function(t){try{JSON.parse(t)}catch(e){return!1}return!0},t.isMultibyte=function(t){return x.test(t)},t.isAscii=function(t){return A.test(t)},t.isFullWidth=function(t){return g.test(t)},t.isHalfWidth=function(t){return h.test(t)},t.isVariableWidth=function(t){return g.test(t)&&h.test(t)},t.isSurrogatePair=function(t){return E.test(t)},t.isBase64=function(t){return $.test(t)},t.ltrim=function(t,e){var u=e?new RegExp("^["+e+"]+","g"):/^\s+/g;return t.replace(u,"")},t.rtrim=function(t,e){var u=e?new RegExp("["+e+"]+$","g"):/\s+$/g;return t.replace(u,"")},t.trim=function(t,e){var u=e?new RegExp("^["+e+"]+|["+e+"]+$","g"):/^\s+|\s+$/g;return t.replace(u,"")},t.escape=function(t){return t.replace(/&/g,"&amp;").replace(/"/g,"&quot;").replace(/'/g,"&apos;").replace(/</g,"&lt;").replace(/>/g,"&gt;")},t.stripLow=function(e,u){var n=u?"\x00-	\f-":"\x00-";return t.blacklist(e,n)},t.whitelist=function(t,e){return t.replace(new RegExp("[^"+e+"]+","g"),"")},t.blacklist=function(t,e){return t.replace(new RegExp("["+e+"]+","g"),"")},t.init(),t});
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
!function(t,e){"undefined"!=typeof module?module.exports=e():"function"==typeof define&&"object"==typeof define.amd?define(e):this[t]=e()}("validator",function(t){"use strict";function e(t,e){t=t||{};for(var n in e)"undefined"==typeof t[n]&&(t[n]=e[n]);return t}t={version:"3.11.2"};var n=/^((([a-z]|\d|[!#\$%&'\*\+\-\/=\?\^_`{\|}~]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])+(\.([a-z]|\d|[!#\$%&'\*\+\-\/=\?\^_`{\|}~]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])+)*)|((\x22)((((\x20|\x09)*(\x0d\x0a))?(\x20|\x09)+)?(([\x01-\x08\x0b\x0c\x0e-\x1f\x7f]|\x21|[\x23-\x5b]|[\x5d-\x7e]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(\\([\x01-\x09\x0b\x0c\x0d-\x7f]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]))))*(((\x20|\x09)*(\x0d\x0a))?(\x20|\x09)+)?(\x22)))@((([a-z]|\d|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(([a-z]|\d|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])*([a-z]|\d|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])))\.)+(([a-z]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(([a-z]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])*([a-z]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])))\.?$/i,u=/^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$/,r=/^(?:[0-9]{9}X|[0-9]{10})$/,i=/^(?:[0-9]{13})$/,F=/^(\d?\d?\d)\.(\d?\d?\d)\.(\d?\d?\d)\.(\d?\d?\d)$/,o=/^::|^::1|^([a-fA-F0-9]{1,4}::?){1,7}([a-fA-F0-9]{1,4})$/,a={3:/^[0-9A-F]{8}-[0-9A-F]{4}-3[0-9A-F]{3}-[0-9A-F]{4}-[0-9A-F]{12}$/i,4:/^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i,5:/^[0-9A-F]{8}-[0-9A-F]{4}-5[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i,all:/^[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}$/i},f=/^[a-zA-Z]+$/,c=/^[a-zA-Z0-9]+$/,s=/^-?[0-9]+$/,l=/^(?:-?(?:0|[1-9][0-9]*))$/,d=/^(?:-?(?:[0-9]+))?(?:\.[0-9]*)?(?:[eE][\+\-]?(?:[0-9]+))?$/,p=/^[0-9a-fA-F]+$/,x=/^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$/,A=/^[\x00-\x7F]+$/,D=/[^\x00-\x7F]/,g=/[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]/,E=/[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]/;t.extend=function(e,n){t[e]=function(){var e=Array.prototype.slice.call(arguments);return e[0]=t.toString(e[0]),n.apply(t,e)}},t.init=function(){for(var e in t)"function"==typeof t[e]&&"toString"!==e&&"toDate"!==e&&"extend"!==e&&"init"!==e&&t.extend(e,t[e])},t.toString=function(t){return"object"==typeof t&&null!==t&&t.toString?t=t.toString():null===t||"undefined"==typeof t||isNaN(t)&&!t.length?t="":"string"!=typeof t&&(t+=""),t},t.toDate=function(t){return"[object Date]"===Object.prototype.toString.call(t)?t:(t=Date.parse(t),isNaN(t)?null:new Date(t))},t.toFloat=function(t){return parseFloat(t)},t.toInt=function(t,e){return parseInt(t,e||10)},t.toBoolean=function(t,e){return e?"1"===t||"true"===t:"0"!==t&&"false"!==t&&""!==t},t.equals=function(e,n){return e===t.toString(n)},t.contains=function(e,n){return e.indexOf(t.toString(n))>=0},t.matches=function(t,e,n){return"[object RegExp]"!==Object.prototype.toString.call(e)&&(e=new RegExp(e,n)),e.test(t)},t.isEmail=function(t){return n.test(t)};var h={protocols:["http","https","ftp"],require_tld:!0,require_protocol:!1};return t.isURL=function(t,n){if(!t||t.length>=2083)return!1;n=e(n,h);var u=new RegExp("^(?!mailto:)(?:(?:"+n.protocols.join("|")+")://)"+(n.require_protocol?"":"?")+"(?:\\S+(?::\\S*)?@)?(?:(?:(?:[1-9]\\d?|1\\d\\d|2[01]\\d|22[0-3])(?:\\.(?:1?\\d{1,2}|2[0-4]\\d|25[0-5])){2}(?:\\.(?:[0-9]\\d?|1\\d\\d|2[0-4]\\d|25[0-4]))|(?:(?:www.)?xn--)?(?:(?:[a-z\\u00a1-\\uffff0-9]+-?)*[a-z\\u00a1-\\uffff0-9]+)(?:\\.(?:[a-z\\u00a1-\\uffff0-9]+-?)*[a-z\\u00a1-\\uffff0-9]+)*(?:\\.(?:[a-z\\u00a1-\\uffff]{2,}))"+(n.require_tld?"":"?")+")|localhost)(?::(\\d{1,5}))?(?:(?:/|\\?|#)[^\\s]*)?$","i"),r=t.match(u),i=r?r[1]:0;return r&&(!i||i>0&&65535>=i)},t.isIP=function(e,n){if(n=t.toString(n),!n)return t.isIP(e,4)||t.isIP(e,6);if("4"===n){if(!F.test(e))return!1;var u=e.split(".").sort();return u[3]<=255}return"6"===n&&o.test(e)},t.isAlpha=function(t){return f.test(t)},t.isAlphanumeric=function(t){return c.test(t)},t.isNumeric=function(t){return s.test(t)},t.isHexadecimal=function(t){return p.test(t)},t.isHexColor=function(t){return x.test(t)},t.isLowercase=function(t){return t===t.toLowerCase()},t.isUppercase=function(t){return t===t.toUpperCase()},t.isInt=function(t){return l.test(t)},t.isFloat=function(t){return""!==t&&d.test(t)},t.isDivisibleBy=function(e,n){return t.toFloat(e)%t.toInt(n)===0},t.isNull=function(t){return 0===t.length},t.isLength=function(t,e,n){return t.length>=e&&("undefined"==typeof n||t.length<=n)},t.isUUID=function(t,e){var n=a[e?e:"all"];return n&&n.test(t)},t.isDate=function(t){return!isNaN(Date.parse(t))},t.isAfter=function(e,n){var u=t.toDate(n||new Date),r=t.toDate(e);return r&&u&&r>u},t.isBefore=function(e,n){var u=t.toDate(n||new Date),r=t.toDate(e);return r&&u&&u>r},t.isIn=function(e,n){if(!n||"function"!=typeof n.indexOf)return!1;if("[object Array]"===Object.prototype.toString.call(n)){for(var u=[],r=0,i=n.length;i>r;r++)u[r]=t.toString(n[r]);n=u}return n.indexOf(e)>=0},t.isCreditCard=function(t){var e=t.replace(/[^0-9]+/g,"");if(!u.test(e))return!1;for(var n,r,i,F=0,o=e.length-1;o>=0;o--)n=e.substring(o,o+1),r=parseInt(n,10),i?(r*=2,F+=r>=10?r%10+1:r):F+=r,i=!i;return F%10===0?e:!1},t.isISBN=function(e,n){if(n=t.toString(n),!n)return t.isISBN(e,10)||t.isISBN(e,13);var u,F=e.replace(/[\s-]+/g,""),o=0;if("10"===n){if(!r.test(F))return!1;for(u=0;9>u;u++)o+=(u+1)*F.charAt(u);if(o+="X"===F.charAt(9)?100:10*F.charAt(9),o%11===0)return F}else if("13"===n){if(!i.test(F))return!1;var a=[1,3];for(u=0;12>u;u++)o+=a[u%2]*F.charAt(u);if(F.charAt(12)-(10-o%10)%10===0)return F}return!1},t.isJSON=function(t){try{JSON.parse(t)}catch(e){if(e instanceof SyntaxError)return!1}return!0},t.isMultibyte=function(t){return D.test(t)},t.isAscii=function(t){return A.test(t)},t.isFullWidth=function(t){return g.test(t)},t.isHalfWidth=function(t){return E.test(t)},t.isVariableWidth=function(t){return g.test(t)&&E.test(t)},t.ltrim=function(t,e){var n=e?new RegExp("^["+e+"]+","g"):/^\s+/g;return t.replace(n,"")},t.rtrim=function(t,e){var n=e?new RegExp("["+e+"]+$","g"):/\s+$/g;return t.replace(n,"")},t.trim=function(t,e){var n=e?new RegExp("^["+e+"]+|["+e+"]+$","g"):/^\s+|\s+$/g;return t.replace(n,"")},t.escape=function(t){return t.replace(/&/g,"&amp;").replace(/"/g,"&quot;").replace(/</g,"&lt;").replace(/>/g,"&gt;")},t.stripLow=function(e,n){var u=n?"\x00-	\f-":"\x00-";return t.blacklist(e,u)},t.whitelist=function(t,e){return t.replace(new RegExp("[^"+e+"]+","g"),"")},t.blacklist=function(t,e){return t.replace(new RegExp("["+e+"]+","g"),"")},t.init(),t});
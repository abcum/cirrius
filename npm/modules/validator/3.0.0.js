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
!function(t,e){"undefined"!=typeof module?module.exports=e():"function"==typeof define&&"object"==typeof define.amd?define(e):this[t]=e()}("validator",function(t){t={version:"2.1.0"};var e=/^(?:[\w\!\#\$\%\&\'\*\+\-\/\=\?\^\`\{\|\}\~]+\.)*[\w\!\#\$\%\&\'\*\+\-\/\=\?\^\`\{\|\}\~]+@(?:(?:(?:[a-zA-Z0-9](?:[a-zA-Z0-9\-](?!\.)){0,61}[a-zA-Z0-9]?\.)+[a-zA-Z0-9](?:[a-zA-Z0-9\-](?!$)){0,61}[a-zA-Z0-9]?)|(?:\[(?:(?:[01]?\d{1,2}|2[0-4]\d|25[0-5])\.){3}(?:[01]?\d{1,2}|2[0-4]\d|25[0-5])\]))$/,n=/^(?!mailto:)(?:(?:https?|ftp):\/\/)?(?:\S+(?::\S*)?@)?(?:(?:(?:[1-9]\d?|1\d\d|2[01]\d|22[0-3])(?:\.(?:1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.(?:[0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(?:(?:[a-z\u00a1-\uffff0-9]+-?)*[a-z\u00a1-\uffff0-9]+)(?:\.(?:[a-z\u00a1-\uffff0-9]+-?)*[a-z\u00a1-\uffff0-9]+)*(?:\.(?:[a-z\u00a1-\uffff]{2,})))|localhost)(?::\d{2,5})?(?:\/[^\s]*)?$/i,r=/^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$/,i=/^(\d?\d?\d)\.(\d?\d?\d)\.(\d?\d?\d)\.(\d?\d?\d)$/,o=/^::|^::1|^([a-fA-F0-9]{1,4}::?){1,7}([a-fA-F0-9]{1,4})$/,u={3:/^[0-9A-F]{8}-[0-9A-F]{4}-3[0-9A-F]{3}-[0-9A-F]{4}-[0-9A-F]{12}$/i,4:/^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i,5:/^[0-9A-F]{8}-[0-9A-F]{4}-5[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i,all:/^[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}$/i},a=/^[a-zA-Z]+$/,f=/^[a-zA-Z0-9]+$/,c=/^-?[0-9]+$/,s=/^(?:-?(?:0|[1-9][0-9]*))$/,l=/^(?:-?(?:[0-9]+))?(?:\.[0-9]*)?(?:[eE][\+\-]?(?:[0-9]+))?$/,p=/^[0-9a-fA-F]+$/,d=/^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$/;t.toString=function(t){return null===t||"undefined"==typeof t||isNaN(t)&&!t.length?t="":"object"==typeof t&&null!==t&&t.toString&&(t=t.toString()),t+""},t.toDate=function(t){return"[object Date]"===Object.prototype.toString.call(t)?t:(t=Date.parse(t),isNaN(t)?null:new Date(t))},t.toFloat=function(t){return parseFloat(t)},t.toInt=function(t,e){return parseInt(t,e||10)},t.toBoolean=function(t,e){return e?"1"===t||"true"===t:"0"!==t&&"false"!==t&&""!==t},t.equals=function(e,n){return e===t.toString(n)},t.contains=function(e,n){return e.indexOf(t.toString(n))>=0},t.matches=function(t,e,n){return"[object RegExp]"!==Object.prototype.toString.call(e)&&(e=new RegExp(e,n)),e.test(t)},t.isEmail=function(t){return e.test(t)},t.isURL=function(t){return t.length<2083&&n.test(t)},t.isIP=function(e,n){if(n=t.toString(n),!n)return t.isIP(e,4)||t.isIP(e,6);if("4"===n){if(!i.test(e))return!1;var r=e.split(".").sort();return r[3]<=255}return"6"===n&&o.test(e)},t.isAlpha=function(t){return a.test(t)},t.isAlphanumeric=function(t){return f.test(t)},t.isNumeric=function(t){return c.test(t)},t.isHexadecimal=function(t){return p.test(t)},t.isHexColor=function(t){return d.test(t)},t.isLowercase=function(t){return t===t.toLowerCase()},t.isUppercase=function(t){return t===t.toUpperCase()},t.isInt=function(t){return s.test(t)},t.isFloat=function(t){return""!==t&&l.test(t)},t.isDivisibleBy=function(e,n){return t.toFloat(e)%t.toInt(n)===0},t.isNull=function(t){return 0===t.length},t.isLength=function(t,e,n){return t.length>=e&&("undefined"==typeof n||t.length<=n)},t.isUUID=function(t,e){var n=u[e?e:"all"];return n&&n.test(t)},t.isDate=function(t){return!isNaN(Date.parse(t))},t.isAfter=function(e,n){var r=t.toDate(n||new Date),i=t.toDate(e);return i&&r&&i>r},t.isBefore=function(e,n){var r=t.toDate(n||new Date),i=t.toDate(e);return i&&r&&r>i},t.isIn=function(e,n){if(!n||"function"!=typeof n.indexOf)return!1;if("[object Array]"===Object.prototype.toString.call(n)){for(var r=[],i=0,o=n.length;o>i;i++)r[i]=t.toString(n[i]);n=r}return n.indexOf(e)>=0},t.isCreditCard=function(t){var e=t.replace(/[^0-9]+/g,"");if(!r.test(e))return!1;for(var n,i,o,u=0,a=e.length-1;a>=0;a--)n=e.substring(a,a+1),i=parseInt(n,10),o?(i*=2,u+=i>=10?i%10+1:i):u+=i,o=!o;return u%10===0?e:!1},t.ltrim=function(t,e){var n=e?new RegExp("^["+e+"]+","g"):/^\s+/g;return t.replace(n,"")},t.rtrim=function(t,e){var n=e?new RegExp("["+e+"]+$","g"):/\s+$/g;return t.replace(n,"")},t.trim=function(t,e){var n=e?new RegExp("^["+e+"]+|["+e+"]+$","g"):/^\s+|\s+$/g;return t.replace(n,"")},t.escape=function(t){return t.replace(/&/g,"&amp;").replace(/"/g,"&quot;").replace(/</g,"&lt;").replace(/>/g,"&gt;")},t.whitelist=function(t,e){return t.replace(new RegExp("[^"+e+"]+","g"),"")},t.blacklist=function(t,e){return t.replace(new RegExp("["+e+"]+","g"),"")};for(var g in t)"toString"!==g&&"toDate"!==g&&"function"==typeof t[g]&&!function(e){var n=t[e];t[e]=function(){var e=Array.prototype.slice.call(arguments);return e[0]=t.toString(e[0]),n.apply(t,e)}}(g);return t});
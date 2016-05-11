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
!function(t,e){"undefined"!=typeof module?module.exports=e():"function"==typeof define&&"object"==typeof define.amd?define(e):this[t]=e()}("validator",function(t){"use strict";t={version:"3.2.0"};var e=/^(?:[\w\!\#\$\%\&\'\*\+\-\/\=\?\^\`\{\|\}\~]+\.)*[\w\!\#\$\%\&\'\*\+\-\/\=\?\^\`\{\|\}\~]+@(?:(?:(?:[a-zA-Z0-9](?:[a-zA-Z0-9\-](?!\.)){0,61}[a-zA-Z0-9]?\.)+[a-zA-Z0-9](?:[a-zA-Z0-9\-](?!$)){0,61}[a-zA-Z0-9]?)|(?:\[(?:(?:[01]?\d{1,2}|2[0-4]\d|25[0-5])\.){3}(?:[01]?\d{1,2}|2[0-4]\d|25[0-5])\]))$/,n=/^(?!mailto:)(?:(?:https?|ftp):\/\/)?(?:\S+(?::\S*)?@)?(?:(?:(?:[1-9]\d?|1\d\d|2[01]\d|22[0-3])(?:\.(?:1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.(?:[0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(?:(?:[a-z\u00a1-\uffff0-9]+-?)*[a-z\u00a1-\uffff0-9]+)(?:\.(?:[a-z\u00a1-\uffff0-9]+-?)*[a-z\u00a1-\uffff0-9]+)*(?:\.(?:[a-z\u00a1-\uffff]{2,})))|localhost)(?::\d{2,5})?(?:\/[^\s]*)?$/i,r=/^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$/,i=/^(?:[0-9]{9}X|[0-9]{10})$/,o=/^(?:[0-9]{13})$/,u=/^(\d?\d?\d)\.(\d?\d?\d)\.(\d?\d?\d)\.(\d?\d?\d)$/,a=/^::|^::1|^([a-fA-F0-9]{1,4}::?){1,7}([a-fA-F0-9]{1,4})$/,f={3:/^[0-9A-F]{8}-[0-9A-F]{4}-3[0-9A-F]{3}-[0-9A-F]{4}-[0-9A-F]{12}$/i,4:/^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i,5:/^[0-9A-F]{8}-[0-9A-F]{4}-5[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i,all:/^[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}$/i},c=/^[a-zA-Z]+$/,s=/^[a-zA-Z0-9]+$/,l=/^-?[0-9]+$/,p=/^(?:-?(?:0|[1-9][0-9]*))$/,d=/^(?:-?(?:[0-9]+))?(?:\.[0-9]*)?(?:[eE][\+\-]?(?:[0-9]+))?$/,g=/^[0-9a-fA-F]+$/,A=/^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$/;t.toString=function(t){return null===t||"undefined"==typeof t||isNaN(t)&&!t.length?t="":"object"==typeof t&&t.toString&&(t=t.toString()),t+""},t.toDate=function(t){return"[object Date]"===Object.prototype.toString.call(t)?t:(t=Date.parse(t),isNaN(t)?null:new Date(t))},t.toFloat=function(t){return parseFloat(t)},t.toInt=function(t,e){return parseInt(t,e||10)},t.toBoolean=function(t,e){return e?"1"===t||"true"===t:"0"!==t&&"false"!==t&&""!==t},t.equals=function(e,n){return e===t.toString(n)},t.contains=function(e,n){return e.indexOf(t.toString(n))>=0},t.matches=function(t,e,n){return"[object RegExp]"!==Object.prototype.toString.call(e)&&(e=new RegExp(e,n)),e.test(t)},t.isEmail=function(t){return e.test(t)},t.isURL=function(t){return t.length<2083&&n.test(t)},t.isIP=function(e,n){if(n=t.toString(n),!n)return t.isIP(e,4)||t.isIP(e,6);if("4"===n){if(!u.test(e))return!1;var r=e.split(".").sort();return r[3]<=255}return"6"===n&&a.test(e)},t.isAlpha=function(t){return c.test(t)},t.isAlphanumeric=function(t){return s.test(t)},t.isNumeric=function(t){return l.test(t)},t.isHexadecimal=function(t){return g.test(t)},t.isHexColor=function(t){return A.test(t)},t.isLowercase=function(t){return t===t.toLowerCase()},t.isUppercase=function(t){return t===t.toUpperCase()},t.isInt=function(t){return p.test(t)},t.isFloat=function(t){return""!==t&&d.test(t)},t.isDivisibleBy=function(e,n){return t.toFloat(e)%t.toInt(n)===0},t.isNull=function(t){return 0===t.length},t.isLength=function(t,e,n){return t.length>=e&&("undefined"==typeof n||t.length<=n)},t.isUUID=function(t,e){var n=f[e?e:"all"];return n&&n.test(t)},t.isDate=function(t){return!isNaN(Date.parse(t))},t.isAfter=function(e,n){var r=t.toDate(n||new Date),i=t.toDate(e);return i&&r&&i>r},t.isBefore=function(e,n){var r=t.toDate(n||new Date),i=t.toDate(e);return i&&r&&r>i},t.isIn=function(e,n){if(!n||"function"!=typeof n.indexOf)return!1;if("[object Array]"===Object.prototype.toString.call(n)){for(var r=[],i=0,o=n.length;o>i;i++)r[i]=t.toString(n[i]);n=r}return n.indexOf(e)>=0},t.isCreditCard=function(t){var e=t.replace(/[^0-9]+/g,"");if(!r.test(e))return!1;for(var n,i,o,u=0,a=e.length-1;a>=0;a--)n=e.substring(a,a+1),i=parseInt(n,10),o?(i*=2,u+=i>=10?i%10+1:i):u+=i,o=!o;return u%10===0?e:!1},t.isISBN=function(e,n){if(n=t.toString(n),!n)return t.isISBN(e,10)||t.isISBN(e,13);var r,u=e.replace(/[\s-]+/g,""),a=0;if("10"===n){if(!i.test(u))return!1;for(r=0;9>r;r++)a+=(r+1)*u.charAt(r);if(a+="X"===u.charAt(9)?100:10*u.charAt(9),a%11===0)return u}else if("13"===n){if(!o.test(u))return!1;var f=[1,3];for(r=0;12>r;r++)a+=f[r%2]*u.charAt(r);if(u.charAt(12)-(10-a%10)%10===0)return u}return!1},t.ltrim=function(t,e){var n=e?new RegExp("^["+e+"]+","g"):/^\s+/g;return t.replace(n,"")},t.rtrim=function(t,e){var n=e?new RegExp("["+e+"]+$","g"):/\s+$/g;return t.replace(n,"")},t.trim=function(t,e){var n=e?new RegExp("^["+e+"]+|["+e+"]+$","g"):/^\s+|\s+$/g;return t.replace(n,"")},t.escape=function(t){return t.replace(/&/g,"&amp;").replace(/"/g,"&quot;").replace(/</g,"&lt;").replace(/>/g,"&gt;")},t.whitelist=function(t,e){return t.replace(new RegExp("[^"+e+"]+","g"),"")},t.blacklist=function(t,e){return t.replace(new RegExp("["+e+"]+","g"),"")},t.extend=function(e,n){t[e]=function(){var e=Array.prototype.slice.call(arguments);return e[0]=t.toString(e[0]),n.apply(t,e)}},t.noCoerce=["toString","toDate","extend"];for(var F in t)"function"!=typeof t[F]||t.noCoerce.indexOf(F)>=0||t.extend(F,t[F]);return t});
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
!function(t,e){"undefined"!=typeof exports&&"undefined"!=typeof module?module.exports=e():"function"==typeof define&&"object"==typeof define.amd?define(e):this[t]=e()}("validator",function(t){"use strict";function e(t,e){t=t||{};for(var r in e)"undefined"==typeof t[r]&&(t[r]=e[r]);return t}function r(t){var e="(\\"+t.symbol.replace(/\./g,"\\.")+")"+(t.require_symbol?"":"?"),r="-?",n="[1-9]\\d*",i="[1-9]\\d{0,2}(\\"+t.thousands_separator+"\\d{3})*",u=["0",n,i],o="("+u.join("|")+")?",a="(\\"+t.decimal_separator+"\\d{2})?",s=o+a;return t.allow_negatives&&!t.parens_for_negatives&&(t.negative_sign_after_digits?s+=r:t.negative_sign_before_digits&&(s=r+s)),t.allow_negative_sign_placeholder?s="( (?!\\-))?"+s:t.allow_space_after_symbol?s=" ?"+s:t.allow_space_after_digits&&(s+="( (?!$))?"),t.symbol_after_digits?s+=e:s=e+s,t.allow_negatives&&(t.parens_for_negatives?s="(\\("+s+"\\)|"+s+")":t.negative_sign_before_digits||t.negative_sign_after_digits||(s=r+s)),new RegExp("^(?!-? )(?=.*\\d)"+s+"$")}t={version:"3.37.0"};var n=/^((([a-z]|\d|[!#\$%&'\*\+\-\/=\?\^_`{\|}~])+(\.([a-z]|\d|[!#\$%&'\*\+\-\/=\?\^_`{\|}~])+)*)|((\x22)((((\x20|\x09)*(\x0d\x0a))?(\x20|\x09)+)?(([\x01-\x08\x0b\x0c\x0e-\x1f\x7f]|\x21|[\x23-\x5b]|[\x5d-\x7e])|(\\[\x01-\x09\x0b\x0c\x0d-\x7f])))*(((\x20|\x09)*(\x0d\x0a))?(\x20|\x09)+)?(\x22)))$/i,i=/^((([a-z]|\d|[!#\$%&'\*\+\-\/=\?\^_`{\|}~]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])+(\.([a-z]|\d|[!#\$%&'\*\+\-\/=\?\^_`{\|}~]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])+)*)|((\x22)((((\x20|\x09)*(\x0d\x0a))?(\x20|\x09)+)?(([\x01-\x08\x0b\x0c\x0e-\x1f\x7f]|\x21|[\x23-\x5b]|[\x5d-\x7e]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(\\([\x01-\x09\x0b\x0c\x0d-\x7f]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]))))*(((\x20|\x09)*(\x0d\x0a))?(\x20|\x09)+)?(\x22)))$/i,u=/^(?:[a-z]|\d|[!#\$%&'\*\+\-\/=\?\^_`{\|}~\.]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])+(?:[a-z]|\d|[!#\$%&'\*\+\-\/=\?\^_`{\|}~\.]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]|\s)*<(.+)>$/i,o=/^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$/,a=/^[A-Z]{2}[0-9A-Z]{9}[0-9]$/,s=/^(?:[0-9]{9}X|[0-9]{10})$/,l=/^(?:[0-9]{13})$/,f=/^(\d+)\.(\d+)\.(\d+)\.(\d+)$/,c=/^[0-9A-F]{1,4}$/i,F={3:/^[0-9A-F]{8}-[0-9A-F]{4}-3[0-9A-F]{3}-[0-9A-F]{4}-[0-9A-F]{12}$/i,4:/^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i,5:/^[0-9A-F]{8}-[0-9A-F]{4}-5[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i,all:/^[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}$/i},p=/^[A-Z]+$/i,g=/^[0-9A-Z]+$/i,x=/^[-+]?[0-9]+$/,d=/^(?:[-+]?(?:0|[1-9][0-9]*))$/,_=/^(?:[-+]?(?:[0-9]+))?(?:\.[0-9]*)?(?:[eE][\+\-]?(?:[0-9]+))?$/,h=/^[0-9A-F]+$/i,A=/^#?([0-9A-F]{3}|[0-9A-F]{6})$/i,v=/^[\x00-\x7F]+$/,D=/[^\x00-\x7F]/,$=/[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]/,w=/[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]/,b=/[\uD800-\uDBFF][\uDC00-\uDFFF]/,m=/^(?:[A-Z0-9+\/]{4})*(?:[A-Z0-9+\/]{2}==|[A-Z0-9+\/]{3}=|[A-Z0-9+\/]{4})$/i,y={"zh-CN":/^(\+?0?86\-?)?1[345789]\d{9}$/,"en-ZA":/^(\+?27|0)\d{9}$/,"en-AU":/^(\+?61|0)4\d{8}$/,"en-HK":/^(\+?852\-?)?[569]\d{3}\-?\d{4}$/,"fr-FR":/^(\+?33|0)[67]\d{8}$/,"pt-PT":/^(\+351)?9[1236]\d{7}$/,"el-GR":/^(\+30)?((2\d{9})|(69\d{8}))$/,"en-GB":/^(\+?44|0)7\d{9}$/};t.extend=function(e,r){t[e]=function(){var e=Array.prototype.slice.call(arguments);return e[0]=t.toString(e[0]),r.apply(t,e)}},t.init=function(){for(var e in t)"function"==typeof t[e]&&"toString"!==e&&"toDate"!==e&&"extend"!==e&&"init"!==e&&t.extend(e,t[e])},t.toString=function(t){return"object"==typeof t&&null!==t&&t.toString?t=t.toString():null===t||"undefined"==typeof t||isNaN(t)&&!t.length?t="":"string"!=typeof t&&(t+=""),t},t.toDate=function(t){return"[object Date]"===Object.prototype.toString.call(t)?t:(t=Date.parse(t),isNaN(t)?null:new Date(t))},t.toFloat=function(t){return parseFloat(t)},t.toInt=function(t,e){return parseInt(t,e||10)},t.toBoolean=function(t,e){return e?"1"===t||"true"===t:"0"!==t&&"false"!==t&&""!==t},t.equals=function(e,r){return e===t.toString(r)},t.contains=function(e,r){return e.indexOf(t.toString(r))>=0},t.matches=function(t,e,r){return"[object RegExp]"!==Object.prototype.toString.call(e)&&(e=new RegExp(e,r)),e.test(t)};var E={allow_display_name:!1,allow_utf8_local_part:!0,require_tld:!0};t.isEmail=function(r,o){if(o=e(o,E),o.allow_display_name){var a=r.match(u);a&&(r=a[1])}else if(/\s/.test(r))return!1;var s=r.split("@"),l=s.pop(),f=s.join("@");return t.isFQDN(l,{require_tld:o.require_tld})?o.allow_utf8_local_part?i.test(f):n.test(f):!1};var C={protocols:["http","https","ftp"],require_tld:!0,require_protocol:!1,allow_underscores:!1,allow_trailing_dot:!1,allow_protocol_relative_urls:!1};t.isURL=function(r,n){if(!r||r.length>=2083||/\s/.test(r))return!1;if(0===r.indexOf("mailto:"))return!1;n=e(n,C);var i,u,o,a,s,l,f;if(f=r.split("://"),f.length>1){if(i=f.shift(),-1===n.protocols.indexOf(i))return!1}else{if(n.require_protocol)return!1;n.allow_protocol_relative_urls&&"//"===r.substr(0,2)&&(f[0]=r.substr(2))}return r=f.join("://"),f=r.split("#"),r=f.shift(),f=r.split("?"),r=f.shift(),f=r.split("/"),r=f.shift(),f=r.split("@"),f.length>1&&(u=f.shift(),u.indexOf(":")>=0&&u.split(":").length>2)?!1:(a=f.join("@"),f=a.split(":"),o=f.shift(),f.length&&(l=f.join(":"),s=parseInt(l,10),!/^[0-9]+$/.test(l)||0>=s||s>65535)?!1:t.isIP(o)||t.isFQDN(o,n)||"localhost"===o?n.host_whitelist&&-1===n.host_whitelist.indexOf(o)?!1:n.host_blacklist&&-1!==n.host_blacklist.indexOf(o)?!1:!0:!1)},t.isIP=function(e,r){if(r=t.toString(r),!r)return t.isIP(e,4)||t.isIP(e,6);if("4"===r){if(!f.test(e))return!1;var n=e.split(".").sort(function(t,e){return t-e});return n[3]<=255}if("6"===r){var i=e.split(":"),u=!1;if(i.length>8)return!1;if("::"===e)return!0;"::"===e.substr(0,2)?(i.shift(),i.shift(),u=!0):"::"===e.substr(e.length-2)&&(i.pop(),i.pop(),u=!0);for(var o=0;o<i.length;++o)if(""===i[o]&&o>0&&o<i.length-1){if(u)return!1;u=!0}else if(!c.test(i[o]))return!1;return u?i.length>=1:8===i.length}return!1};var I={require_tld:!0,allow_underscores:!1,allow_trailing_dot:!1};t.isFQDN=function(t,r){r=e(r,I),r.allow_trailing_dot&&"."===t[t.length-1]&&(t=t.substring(0,t.length-1));var n=t.split(".");if(r.require_tld){var i=n.pop();if(!n.length||!/^([a-z\u00a1-\uffff]{2,}|xn[a-z0-9-]{2,})$/i.test(i))return!1}for(var u,o=0;o<n.length;o++){if(u=n[o],r.allow_underscores){if(u.indexOf("__")>=0)return!1;u=u.replace(/_/g,"")}if(!/^[a-z\u00a1-\uffff0-9-]+$/i.test(u))return!1;if("-"===u[0]||"-"===u[u.length-1]||u.indexOf("---")>=0)return!1}return!0},t.isAlpha=function(t){return p.test(t)},t.isAlphanumeric=function(t){return g.test(t)},t.isNumeric=function(t){return x.test(t)},t.isHexadecimal=function(t){return h.test(t)},t.isHexColor=function(t){return A.test(t)},t.isLowercase=function(t){return t===t.toLowerCase()},t.isUppercase=function(t){return t===t.toUpperCase()},t.isInt=function(t){return d.test(t)},t.isFloat=function(t){return""!==t&&_.test(t)},t.isDivisibleBy=function(e,r){return t.toFloat(e)%t.toInt(r)===0},t.isNull=function(t){return 0===t.length},t.isLength=function(t,e,r){var n=t.match(/[\uD800-\uDBFF][\uDC00-\uDFFF]/g)||[],i=t.length-n.length;return i>=e&&("undefined"==typeof r||r>=i)},t.isByteLength=function(t,e,r){return t.length>=e&&("undefined"==typeof r||t.length<=r)},t.isUUID=function(t,e){var r=F[e?e:"all"];return r&&r.test(t)},t.isDate=function(t){return!isNaN(Date.parse(t))},t.isAfter=function(e,r){var n=t.toDate(r||new Date),i=t.toDate(e);return!!(i&&n&&i>n)},t.isBefore=function(e,r){var n=t.toDate(r||new Date),i=t.toDate(e);return i&&n&&n>i},t.isIn=function(e,r){var n;if("[object Array]"===Object.prototype.toString.call(r)){var i=[];for(n in r)i[n]=t.toString(r[n]);return i.indexOf(e)>=0}return"object"==typeof r?r.hasOwnProperty(e):r&&"function"==typeof r.indexOf?r.indexOf(e)>=0:!1},t.isCreditCard=function(t){var e=t.replace(/[^0-9]+/g,"");if(!o.test(e))return!1;for(var r,n,i,u=0,a=e.length-1;a>=0;a--)r=e.substring(a,a+1),n=parseInt(r,10),i?(n*=2,u+=n>=10?n%10+1:n):u+=n,i=!i;return!!(u%10===0?e:!1)},t.isISIN=function(t){if(!a.test(t))return!1;for(var e,r,n=t.replace(/[A-Z]/g,function(t){return parseInt(t,36)}),i=0,u=!0,o=n.length-2;o>=0;o--)e=n.substring(o,o+1),r=parseInt(e,10),u?(r*=2,i+=r>=10?r+1:r):i+=r,u=!u;return parseInt(t.substr(t.length-1),10)===(1e4-i)%10},t.isISBN=function(e,r){if(r=t.toString(r),!r)return t.isISBN(e,10)||t.isISBN(e,13);var n,i=e.replace(/[\s-]+/g,""),u=0;if("10"===r){if(!s.test(i))return!1;for(n=0;9>n;n++)u+=(n+1)*i.charAt(n);if(u+="X"===i.charAt(9)?100:10*i.charAt(9),u%11===0)return!!i}else if("13"===r){if(!l.test(i))return!1;var o=[1,3];for(n=0;12>n;n++)u+=o[n%2]*i.charAt(n);if(i.charAt(12)-(10-u%10)%10===0)return!!i}return!1},t.isMobilePhone=function(t,e){return e in y?y[e].test(t):!1};var S={symbol:"$",require_symbol:!1,allow_space_after_symbol:!1,symbol_after_digits:!1,allow_negatives:!0,parens_for_negatives:!1,negative_sign_before_digits:!1,negative_sign_after_digits:!1,allow_negative_sign_placeholder:!1,thousands_separator:",",decimal_separator:".",allow_space_after_digits:!1};t.isCurrency=function(t,n){return n=e(n,S),r(n).test(t)},t.isJSON=function(t){try{JSON.parse(t)}catch(e){return!1}return!0},t.isMultibyte=function(t){return D.test(t)},t.isAscii=function(t){return v.test(t)},t.isFullWidth=function(t){return $.test(t)},t.isHalfWidth=function(t){return w.test(t)},t.isVariableWidth=function(t){return $.test(t)&&w.test(t)},t.isSurrogatePair=function(t){return b.test(t)},t.isBase64=function(t){return m.test(t)},t.isMongoId=function(e){return t.isHexadecimal(e)&&24===e.length},t.ltrim=function(t,e){var r=e?new RegExp("^["+e+"]+","g"):/^\s+/g;return t.replace(r,"")},t.rtrim=function(t,e){var r=e?new RegExp("["+e+"]+$","g"):/\s+$/g;return t.replace(r,"")},t.trim=function(t,e){var r=e?new RegExp("^["+e+"]+|["+e+"]+$","g"):/^\s+|\s+$/g;return t.replace(r,"")},t.escape=function(t){return t.replace(/&/g,"&amp;").replace(/"/g,"&quot;").replace(/'/g,"&#x27;").replace(/</g,"&lt;").replace(/>/g,"&gt;").replace(/\//g,"&#x2F;").replace(/\`/g,"&#96;")},t.stripLow=function(e,r){var n=r?"\\x00-\\x09\\x0B\\x0C\\x0E-\\x1F\\x7F":"\\x00-\\x1F\\x7F";return t.blacklist(e,n)},t.whitelist=function(t,e){return t.replace(new RegExp("[^"+e+"]+","g"),"")},t.blacklist=function(t,e){return t.replace(new RegExp("["+e+"]+","g"),"")};var N={lowercase:!0};return t.normalizeEmail=function(r,n){if(n=e(n,N),!t.isEmail(r))return!1;var i=r.split("@",2);if(i[1]=i[1].toLowerCase(),"gmail.com"===i[1]||"googlemail.com"===i[1]){if(i[0]=i[0].toLowerCase().replace(/\./g,""),"+"===i[0][0])return!1;i[0]=i[0].split("+")[0],i[1]="gmail.com"}else n.lowercase&&(i[0]=i[0].toLowerCase());return i.join("@")},t.init(),t});
/*!
 * Copyright (c) 2015 Chris O'Hara <cohara87@gmail.com>
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
!function(e,t){"undefined"!=typeof exports&&"undefined"!=typeof module?module.exports=t():"function"==typeof define&&"object"==typeof define.amd?define(t):"function"==typeof define&&"object"==typeof define.petal?define(e,[],t):this[e]=t()}("validator",function(e){"use strict";function t(e){var t,r,n,i,o=e.match(j);if(o){if(t=o[21],!t)return o[12]?null:0;if("z"===t||"Z"===t)return 0;r=o[22],-1!==t.indexOf(":")?(n=parseInt(o[23]),i=parseInt(o[24])):(n=0,i=parseInt(o[23]))}else{if(e=e.toLowerCase(),t=e.match(/(?:\s|gmt\s*)(-|\+)(\d{1,4})(\s|$)/),!t)return-1!==e.indexOf("gmt")?0:null;r=t[1];var a=t[2];3===a.length&&(a="0"+a),a.length<=2?(n=0,i=parseInt(a)):(n=parseInt(a.slice(0,2)),i=parseInt(a.slice(2,4)))}return(60*n+i)*("-"===r?1:-1)}function r(e,t){e=e||{};for(var r in t)"undefined"==typeof e[r]&&(e[r]=t[r]);return e}function n(e){var t="(\\"+e.symbol.replace(/\./g,"\\.")+")"+(e.require_symbol?"":"?"),r="-?",n="[1-9]\\d*",i="[1-9]\\d{0,2}(\\"+e.thousands_separator+"\\d{3})*",o=["0",n,i],a="("+o.join("|")+")?",s="(\\"+e.decimal_separator+"\\d{2})?",u=a+s;return e.allow_negatives&&!e.parens_for_negatives&&(e.negative_sign_after_digits?u+=r:e.negative_sign_before_digits&&(u=r+u)),e.allow_negative_sign_placeholder?u="( (?!\\-))?"+u:e.allow_space_after_symbol?u=" ?"+u:e.allow_space_after_digits&&(u+="( (?!$))?"),e.symbol_after_digits?u+=t:u=t+u,e.allow_negatives&&(e.parens_for_negatives?u="(\\("+u+"\\)|"+u+")":e.negative_sign_before_digits||e.negative_sign_after_digits||(u=r+u)),new RegExp("^(?!-? )(?=.*\\d)"+u+"$")}e={version:"4.7.1",coerce:!0};var i=/^[a-z\d!#\$%&'\*\+\-\/=\?\^_`{\|}~]+$/i,o=/^([\s\x01-\x08\x0b\x0c\x0e-\x1f\x7f\x21\x23-\x5b\x5d-\x7e]|(\\[\x01-\x09\x0b\x0c\x0d-\x7f]))*$/i,a=/^[a-z\d!#\$%&'\*\+\-\/=\?\^_`{\|}~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]+$/i,s=/^([\s\x01-\x08\x0b\x0c\x0e-\x1f\x7f\x21\x23-\x5b\x5d-\x7e\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]|(\\[\x01-\x09\x0b\x0c\x0d-\x7f\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]))*$/i,u=/^[a-z\d!#\$%&'\*\+\-\/=\?\^_`{\|}~\.\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]+[a-z\d!#\$%&'\*\+\-\/=\?\^_`{\|}~\.\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF\s]*<(.+)>$/i,l=/^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$/,f=/^[A-Z]{2}[0-9A-Z]{9}[0-9]$/,c=/^(?:[0-9]{9}X|[0-9]{10})$/,d=/^(?:[0-9]{13})$/,p=/^([0-9a-fA-F][0-9a-fA-F]:){5}([0-9a-fA-F][0-9a-fA-F])$/,g=/^(\d+)\.(\d+)\.(\d+)\.(\d+)$/,F=/^[0-9A-F]{1,4}$/i,_={3:/^[0-9A-F]{8}-[0-9A-F]{4}-3[0-9A-F]{3}-[0-9A-F]{4}-[0-9A-F]{12}$/i,4:/^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i,5:/^[0-9A-F]{8}-[0-9A-F]{4}-5[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i,all:/^[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}$/i},x={"en-US":/^[A-Z]+$/i,"de-DE":/^[A-ZÄÖÜß]+$/i},h={"en-US":/^[0-9A-Z]+$/i,"de-DE":/^[0-9A-ZÄÖÜß]+$/i},v=/^[-+]?[0-9]+$/,m=/^(?:[-+]?(?:0|[1-9][0-9]*))$/,A=/^(?:[-+]?(?:[0-9]+))?(?:\.[0-9]*)?(?:[eE][\+\-]?(?:[0-9]+))?$/,$=/^[0-9A-F]+$/i,w=/^[-+]?([0-9]+|\.[0-9]+|[0-9]+\.[0-9]+)$/,y=/^#?([0-9A-F]{3}|[0-9A-F]{6})$/i,b=/^[\x00-\x7F]+$/,D=/[^\x00-\x7F]/,I=/[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]/,E=/[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]/,O=/[\uD800-\uDBFF][\uDC00-\uDFFF]/,S=/^(?:[A-Z0-9+\/]{4})*(?:[A-Z0-9+\/]{2}==|[A-Z0-9+\/]{3}=|[A-Z0-9+\/]{4})$/i,C={"zh-CN":/^(\+?0?86\-?)?((13\d|14[57]|15[^4,\D]|17[678]|18\d)\d{8}|170[059]\d{7})$/,"zh-TW":/^(\+?886\-?|0)?9\d{8}$/,"en-ZA":/^(\+?27|0)\d{9}$/,"en-AU":/^(\+?61|0)4\d{8}$/,"en-HK":/^(\+?852\-?)?[569]\d{3}\-?\d{4}$/,"fr-FR":/^(\+?33|0)[67]\d{8}$/,"pt-PT":/^(\+?351)?9[1236]\d{7}$/,"el-GR":/^(\+?30)?(69\d{8})$/,"en-GB":/^(\+?44|0)7\d{9}$/,"en-US":/^(\+?1)?[2-9]\d{2}[2-9](?!11)\d{6}$/,"en-ZM":/^(\+?26)?09[567]\d{7}$/,"ru-RU":/^(\+?7|8)?9\d{9}$/,"nb-NO":/^(\+?47)?[49]\d{7}$/,"nn-NO":/^(\+?47)?[49]\d{7}$/,"vi-VN":/^(\+?84|0)?((1(2([0-9])|6([2-9])|88|99))|(9((?!5)[0-9])))([0-9]{7})$/,"en-NZ":/^(\+?64|0)2\d{7,9}$/,"en-IN":/^(\+?91|0)?[789]\d{9}$/,"es-ES":/^(\+?34)?(6\d{1}|7[1234])\d{7}$/,"de-DE":/^(\+?49[ \.\-])?([\(]{1}[0-9]{1,6}[\)])?([0-9 \.\-\/]{3,20})((x|ext|extension)[ ]?[0-9]{1,4})?$/,"fi-FI":/^(\+?358|0)\s?(4(0|1|2|4|5)?|50)\s?(\d\s?){4,8}\d$/},j=/^([\+-]?\d{4}(?!\d{2}\b))((-?)((0[1-9]|1[0-2])(\3([12]\d|0[1-9]|3[01]))?|W([0-4]\d|5[0-2])(-?[1-7])?|(00[1-9]|0[1-9]\d|[12]\d{2}|3([0-5]\d|6[1-6])))([T\s]((([01]\d|2[0-3])((:?)[0-5]\d)?|24\:?00)([\.,]\d+(?!:))?)?(\17[0-5]\d([\.,]\d+)?)?([zZ]|([\+-])([01]\d|2[0-3]):?([0-5]\d)?)?)?)?$/;e.extend=function(t,r){e[t]=function(){var t=Array.prototype.slice.call(arguments);return t[0]=e.toString(t[0]),r.apply(e,t)}},e.init=function(){for(var t in e)"function"==typeof e[t]&&"toString"!==t&&"toDate"!==t&&"extend"!==t&&"init"!==t&&e.extend(t,e[t])};var N=null;e.deprecation=function(e){if(null===N){if("function"!=typeof require)return;N=require("depd")("validator")}N(e)},e.toString=function(t){if("string"!=typeof t){if(!e.coerce)throw new Error("this library validates strings only");e.deprecation("you tried to validate a "+typeof t+" but this library (validator.js) validates strings only. Please update your code as this will be an error soon.")}return"object"==typeof t&&null!==t?t="function"==typeof t.toString?t.toString():"[object Object]":(null===t||"undefined"==typeof t||isNaN(t)&&!t.length)&&(t=""),""+t},e.toDate=function(e){return"[object Date]"===Object.prototype.toString.call(e)?e:(e=Date.parse(e),isNaN(e)?null:new Date(e))},e.toFloat=function(e){return parseFloat(e)},e.toInt=function(e,t){return parseInt(e,t||10)},e.toBoolean=function(e,t){return t?"1"===e||"true"===e:"0"!==e&&"false"!==e&&""!==e},e.equals=function(t,r){return t===e.toString(r)},e.contains=function(t,r){return t.indexOf(e.toString(r))>=0},e.matches=function(e,t,r){return"[object RegExp]"!==Object.prototype.toString.call(t)&&(t=new RegExp(t,r)),t.test(e)};var Z={allow_display_name:!1,allow_utf8_local_part:!0,require_tld:!0};e.isEmail=function(t,n){if(n=r(n,Z),n.allow_display_name){var l=t.match(u);l&&(t=l[1])}var f=t.split("@"),c=f.pop(),d=f.join("@"),p=c.toLowerCase();if(("gmail.com"===p||"googlemail.com"===p)&&(d=d.replace(/\./g,"").toLowerCase()),!e.isByteLength(d,{max:64})||!e.isByteLength(c,{max:256}))return!1;if(!e.isFQDN(c,{require_tld:n.require_tld}))return!1;if('"'===d[0])return d=d.slice(1,d.length-1),n.allow_utf8_local_part?s.test(d):o.test(d);for(var g=n.allow_utf8_local_part?a:i,F=d.split("."),_=0;_<F.length;_++)if(!g.test(F[_]))return!1;return!0};var B={protocols:["http","https","ftp"],require_tld:!0,require_protocol:!1,require_valid_protocol:!0,allow_underscores:!1,allow_trailing_dot:!1,allow_protocol_relative_urls:!1};e.isURL=function(t,n){if(!t||t.length>=2083||/\s/.test(t))return!1;if(0===t.indexOf("mailto:"))return!1;n=r(n,B);var i,o,a,s,u,l,f;if(f=t.split("://"),f.length>1){if(i=f.shift(),n.require_valid_protocol&&-1===n.protocols.indexOf(i))return!1}else{if(n.require_protocol)return!1;n.allow_protocol_relative_urls&&"//"===t.substr(0,2)&&(f[0]=t.substr(2))}return t=f.join("://"),f=t.split("#"),t=f.shift(),f=t.split("?"),t=f.shift(),f=t.split("/"),t=f.shift(),f=t.split("@"),f.length>1&&(o=f.shift(),o.indexOf(":")>=0&&o.split(":").length>2)?!1:(s=f.join("@"),f=s.split(":"),a=f.shift(),f.length&&(l=f.join(":"),u=parseInt(l,10),!/^[0-9]+$/.test(l)||0>=u||u>65535)?!1:e.isIP(a)||e.isFQDN(a,n)||"localhost"===a?n.host_whitelist&&-1===n.host_whitelist.indexOf(a)?!1:n.host_blacklist&&-1!==n.host_blacklist.indexOf(a)?!1:!0:!1)},e.isMACAddress=function(e){return p.test(e)},e.isIP=function(t,r){if(r=r?r+"":"",!r)return e.isIP(t,4)||e.isIP(t,6);if("4"===r){if(!g.test(t))return!1;var n=t.split(".").sort(function(e,t){return e-t});return n[3]<=255}if("6"===r){var i=t.split(":"),o=!1,a=e.isIP(i[i.length-1],4),s=a?7:8;if(i.length>s)return!1;if("::"===t)return!0;"::"===t.substr(0,2)?(i.shift(),i.shift(),o=!0):"::"===t.substr(t.length-2)&&(i.pop(),i.pop(),o=!0);for(var u=0;u<i.length;++u)if(""===i[u]&&u>0&&u<i.length-1){if(o)return!1;o=!0}else if(a&&u==i.length-1);else if(!F.test(i[u]))return!1;return o?i.length>=1:i.length===s}return!1};var q={require_tld:!0,allow_underscores:!1,allow_trailing_dot:!1};e.isFQDN=function(e,t){t=r(t,q),t.allow_trailing_dot&&"."===e[e.length-1]&&(e=e.substring(0,e.length-1));var n=e.split(".");if(t.require_tld){var i=n.pop();if(!n.length||!/^([a-z\u00a1-\uffff]{2,}|xn[a-z0-9-]{2,})$/i.test(i))return!1}for(var o,a=0;a<n.length;a++){if(o=n[a],t.allow_underscores){if(o.indexOf("__")>=0)return!1;o=o.replace(/_/g,"")}if(!/^[a-z\u00a1-\uffff0-9-]+$/i.test(o))return!1;if(/[\uff01-\uff5e]/.test(o))return!1;if("-"===o[0]||"-"===o[o.length-1])return!1;if(o.indexOf("---")>=0&&"xn--"!==o.slice(0,4))return!1}return!0},e.isBoolean=function(e){return["true","false","1","0"].indexOf(e)>=0},e.isAlpha=function(e,t){return t=t||"en-US",x[t].test(e)},e.isAlphanumeric=function(e,t){return t=t||"en-US",h[t].test(e)},e.isNumeric=function(e){return v.test(e)},e.isDecimal=function(e){return""!==e&&w.test(e)},e.isHexadecimal=function(e){return $.test(e)},e.isHexColor=function(e){return y.test(e)},e.isLowercase=function(e){return e===e.toLowerCase()},e.isUppercase=function(e){return e===e.toUpperCase()},e.isInt=function(e,t){return t=t||{},m.test(e)&&(!t.hasOwnProperty("min")||e>=t.min)&&(!t.hasOwnProperty("max")||e<=t.max)},e.isFloat=function(e,t){return t=t||{},""===e||"."===e?!1:A.test(e)&&(!t.hasOwnProperty("min")||e>=t.min)&&(!t.hasOwnProperty("max")||e<=t.max)},e.isDivisibleBy=function(t,r){return e.toFloat(t)%parseInt(r,10)===0},e.isNull=function(e){return 0===e.length},e.isLength=function(e,t){var r,n;"object"==typeof t?(r=t.min||0,n=t.max):(r=arguments[1],n=arguments[2]);var i=e.match(/[\uD800-\uDBFF][\uDC00-\uDFFF]/g)||[],o=e.length-i.length;return o>=r&&("undefined"==typeof n||n>=o)},e.isByteLength=function(e,t){var r,n;"object"==typeof t?(r=t.min||0,n=t.max):(r=arguments[1],n=arguments[2]);var i=encodeURI(e).split(/%..|./).length-1;return i>=r&&("undefined"==typeof n||n>=i)},e.isUUID=function(e,t){var r=_[t?t:"all"];return r&&r.test(e)},e.isDate=function(e){var r=new Date(Date.parse(e));if(isNaN(r))return!1;var n=t(e);if(null!==n){var i=r.getTimezoneOffset()-n;r=new Date(r.getTime()+6e4*i)}var o,a,s,u=String(r.getDate());return(a=e.match(/(^|[^:\d])[23]\d([^:\d]|$)/g))?(o=a.map(function(e){return e.match(/\d+/g)[0]}).join("/"),s=String(r.getFullYear()).slice(-2),o===u||o===s?!0:o===u+"/"+s||o===s+"/"+u?!0:!1):!0},e.isAfter=function(t,r){var n=e.toDate(r||new Date),i=e.toDate(t);return!!(i&&n&&i>n)},e.isBefore=function(t,r){var n=e.toDate(r||new Date),i=e.toDate(t);return!!(i&&n&&n>i)},e.isIn=function(t,r){var n;if("[object Array]"===Object.prototype.toString.call(r)){var i=[];for(n in r)i[n]=e.toString(r[n]);return i.indexOf(t)>=0}return"object"==typeof r?r.hasOwnProperty(t):r&&"function"==typeof r.indexOf?r.indexOf(t)>=0:!1},e.isWhitelisted=function(e,t){for(var r=e.length-1;r>=0;r--)if(-1===t.indexOf(e[r]))return!1;return!0},e.isCreditCard=function(e){var t=e.replace(/[^0-9]+/g,"");if(!l.test(t))return!1;for(var r,n,i,o=0,a=t.length-1;a>=0;a--)r=t.substring(a,a+1),n=parseInt(r,10),i?(n*=2,o+=n>=10?n%10+1:n):o+=n,i=!i;return!!(o%10===0?t:!1)},e.isISIN=function(e){if(!f.test(e))return!1;for(var t,r,n=e.replace(/[A-Z]/g,function(e){return parseInt(e,36)}),i=0,o=!0,a=n.length-2;a>=0;a--)t=n.substring(a,a+1),r=parseInt(t,10),o?(r*=2,i+=r>=10?r+1:r):i+=r,o=!o;return parseInt(e.substr(e.length-1),10)===(1e4-i)%10},e.isISBN=function(t,r){if(r=r?r+"":"",!r)return e.isISBN(t,10)||e.isISBN(t,13);var n,i=t.replace(/[\s-]+/g,""),o=0;if("10"===r){if(!c.test(i))return!1;for(n=0;9>n;n++)o+=(n+1)*i.charAt(n);if(o+="X"===i.charAt(9)?100:10*i.charAt(9),o%11===0)return!!i}else if("13"===r){if(!d.test(i))return!1;var a=[1,3];for(n=0;12>n;n++)o+=a[n%2]*i.charAt(n);if(i.charAt(12)-(10-o%10)%10===0)return!!i}return!1},e.isMobilePhone=function(e,t){return t in C?C[t].test(e):!1};var z={symbol:"$",require_symbol:!1,allow_space_after_symbol:!1,symbol_after_digits:!1,allow_negatives:!0,parens_for_negatives:!1,negative_sign_before_digits:!1,negative_sign_after_digits:!1,allow_negative_sign_placeholder:!1,thousands_separator:",",decimal_separator:".",allow_space_after_digits:!1};e.isCurrency=function(e,t){return t=r(t,z),n(t).test(e)},e.isJSON=function(e){try{var t=JSON.parse(e);return!!t&&"object"==typeof t}catch(r){}return!1},e.isMultibyte=function(e){return D.test(e)},e.isAscii=function(e){return b.test(e)},e.isFullWidth=function(e){return I.test(e)},e.isHalfWidth=function(e){return E.test(e)},e.isVariableWidth=function(e){return I.test(e)&&E.test(e)},e.isSurrogatePair=function(e){return O.test(e)},e.isBase64=function(e){return S.test(e)},e.isMongoId=function(t){return e.isHexadecimal(t)&&24===t.length},e.isISO8601=function(e){return j.test(e)},e.ltrim=function(e,t){var r=t?new RegExp("^["+t+"]+","g"):/^\s+/g;return e.replace(r,"")},e.rtrim=function(e,t){var r=t?new RegExp("["+t+"]+$","g"):/\s+$/g;return e.replace(r,"")},e.trim=function(e,t){var r=t?new RegExp("^["+t+"]+|["+t+"]+$","g"):/^\s+|\s+$/g;return e.replace(r,"")},e.escape=function(e){return e.replace(/&/g,"&amp;").replace(/"/g,"&quot;").replace(/'/g,"&#x27;").replace(/</g,"&lt;").replace(/>/g,"&gt;").replace(/\//g,"&#x2F;").replace(/\`/g,"&#96;")},e.stripLow=function(t,r){var n=r?"\\x00-\\x09\\x0B\\x0C\\x0E-\\x1F\\x7F":"\\x00-\\x1F\\x7F";return e.blacklist(t,n)},e.whitelist=function(e,t){return e.replace(new RegExp("[^"+t+"]+","g"),"")},e.blacklist=function(e,t){return e.replace(new RegExp("["+t+"]+","g"),"")};var L={lowercase:!0,remove_dots:!0,remove_extension:!0};return e.normalizeEmail=function(t,n){if(n=r(n,L),!e.isEmail(t))return!1;var i=t.split("@",2);if(i[1]=i[1].toLowerCase(),"gmail.com"===i[1]||"googlemail.com"===i[1]){if(n.remove_extension&&(i[0]=i[0].split("+")[0]),n.remove_dots&&(i[0]=i[0].replace(/\./g,"")),!i[0].length)return!1;i[0]=i[0].toLowerCase(),i[1]="gmail.com"}else n.lowercase&&(i[0]=i[0].toLowerCase());return i.join("@")},e.init(),e});
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
!function(t,e){"undefined"!=typeof exports&&"undefined"!=typeof module?module.exports=e():"function"==typeof define&&"object"==typeof define.amd?define(e):"function"==typeof define&&"object"==typeof define.petal?define(t,[],e):this[t]=e()}("validator",function(t){"use strict";function e(t){var e,r,n,i,o=t.match(N);if(o){if(e=o[21],!e)return o[12]?null:0;if("z"===e||"Z"===e)return 0;r=o[22],-1!==e.indexOf(":")?(n=parseInt(o[23]),i=parseInt(o[24])):(n=0,i=parseInt(o[23]))}else{if(t=t.toLowerCase(),e=t.match(/(?:\s|gmt\s*)(-|\+)(\d{1,4})(\s|$)/),!e)return-1!==t.indexOf("gmt")?0:null;r=e[1];var a=e[2];3===a.length&&(a="0"+a),a.length<=2?(n=0,i=parseInt(a)):(n=parseInt(a.slice(0,2)),i=parseInt(a.slice(2,4)))}return(60*n+i)*("-"===r?1:-1)}function r(t,e){t=t||{};for(var r in e)"undefined"==typeof t[r]&&(t[r]=e[r]);return t}function n(t){var e="(\\"+t.symbol.replace(/\./g,"\\.")+")"+(t.require_symbol?"":"?"),r="-?",n="[1-9]\\d*",i="[1-9]\\d{0,2}(\\"+t.thousands_separator+"\\d{3})*",o=["0",n,i],a="("+o.join("|")+")?",u="(\\"+t.decimal_separator+"\\d{2})?",s=a+u;return t.allow_negatives&&!t.parens_for_negatives&&(t.negative_sign_after_digits?s+=r:t.negative_sign_before_digits&&(s=r+s)),t.allow_negative_sign_placeholder?s="( (?!\\-))?"+s:t.allow_space_after_symbol?s=" ?"+s:t.allow_space_after_digits&&(s+="( (?!$))?"),t.symbol_after_digits?s+=e:s=e+s,t.allow_negatives&&(t.parens_for_negatives?s="(\\("+s+"\\)|"+s+")":t.negative_sign_before_digits||t.negative_sign_after_digits||(s=r+s)),new RegExp("^(?!-? )(?=.*\\d)"+s+"$")}t={version:"4.6.0"};var i=/^[a-z\d!#\$%&'\*\+\-\/=\?\^_`{\|}~]+$/i,o=/^([\s\x01-\x08\x0b\x0c\x0e-\x1f\x7f\x21\x23-\x5b\x5d-\x7e]|(\\[\x01-\x09\x0b\x0c\x0d-\x7f]))*$/i,a=/^[a-z\d!#\$%&'\*\+\-\/=\?\^_`{\|}~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]+$/i,u=/^([\s\x01-\x08\x0b\x0c\x0e-\x1f\x7f\x21\x23-\x5b\x5d-\x7e\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]|(\\[\x01-\x09\x0b\x0c\x0d-\x7f\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]))*$/i,s=/^[a-z\d!#\$%&'\*\+\-\/=\?\^_`{\|}~\.\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]+[a-z\d!#\$%&'\*\+\-\/=\?\^_`{\|}~\.\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF\s]*<(.+)>$/i,l=/^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$/,f=/^[A-Z]{2}[0-9A-Z]{9}[0-9]$/,c=/^(?:[0-9]{9}X|[0-9]{10})$/,d=/^(?:[0-9]{13})$/,g=/^([0-9a-fA-F][0-9a-fA-F]:){5}([0-9a-fA-F][0-9a-fA-F])$/,p=/^(\d+)\.(\d+)\.(\d+)\.(\d+)$/,F=/^[0-9A-F]{1,4}$/i,_={3:/^[0-9A-F]{8}-[0-9A-F]{4}-3[0-9A-F]{3}-[0-9A-F]{4}-[0-9A-F]{12}$/i,4:/^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i,5:/^[0-9A-F]{8}-[0-9A-F]{4}-5[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i,all:/^[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}$/i},x=/^[A-Z]+$/i,h=/^[0-9A-Z]+$/i,m=/^[-+]?[0-9]+$/,v=/^(?:[-+]?(?:0|[1-9][0-9]*))$/,A=/^(?:[-+]?(?:[0-9]+))?(?:\.[0-9]*)?(?:[eE][\+\-]?(?:[0-9]+))?$/,$=/^[0-9A-F]+$/i,w=/^[-+]?([0-9]+|\.[0-9]+|[0-9]+\.[0-9]+)$/,b=/^#?([0-9A-F]{3}|[0-9A-F]{6})$/i,D=/^[\x00-\x7F]+$/,y=/[^\x00-\x7F]/,I=/[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]/,O=/[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]/,E=/[\uD800-\uDBFF][\uDC00-\uDFFF]/,S=/^(?:[A-Z0-9+\/]{4})*(?:[A-Z0-9+\/]{2}==|[A-Z0-9+\/]{3}=|[A-Z0-9+\/]{4})$/i,C={"zh-CN":/^(\+?0?86\-?)?((13\d|14[57]|15[^4,\D]|17[678]|18\d)\d{8}|170[059]\d{7})$/,"zh-TW":/^(\+?886\-?|0)?9\d{8}$/,"en-ZA":/^(\+?27|0)\d{9}$/,"en-AU":/^(\+?61|0)4\d{8}$/,"en-HK":/^(\+?852\-?)?[569]\d{3}\-?\d{4}$/,"fr-FR":/^(\+?33|0)[67]\d{8}$/,"pt-PT":/^(\+351)?9[1236]\d{7}$/,"el-GR":/^(\+?30)?(69\d{8})$/,"en-GB":/^(\+?44|0)7\d{9}$/,"en-US":/^(\+?1)?[2-9]\d{2}[2-9](?!11)\d{6}$/,"en-ZM":/^(\+26)?09[567]\d{7}$/,"ru-RU":/^(\+?7|8)?9\d{9}$/,"nb-NO":/^(\+?47)?[49]\d{7}$/,"nn-NO":/^(\+?47)?[49]\d{7}$/,"vi-VN":/^(0|\+?84)?((1(2([0-9])|6([2-9])|88|99))|(9((?!5)[0-9])))([0-9]{7})$/,"en-NZ":/^(\+?64|0)2\d{7,9}$/,"en-IN":/^(\+?91|0)?[789]\d{9}$/,"es-ES":/^(\+34)?(6\d{1}|7[1234])\d{7}$/},N=/^([\+-]?\d{4}(?!\d{2}\b))((-?)((0[1-9]|1[0-2])(\3([12]\d|0[1-9]|3[01]))?|W([0-4]\d|5[0-2])(-?[1-7])?|(00[1-9]|0[1-9]\d|[12]\d{2}|3([0-5]\d|6[1-6])))([T\s]((([01]\d|2[0-3])((:?)[0-5]\d)?|24\:?00)([\.,]\d+(?!:))?)?(\17[0-5]\d([\.,]\d+)?)?([zZ]|([\+-])([01]\d|2[0-3]):?([0-5]\d)?)?)?)?$/;t.extend=function(e,r){t[e]=function(){var e=Array.prototype.slice.call(arguments);return e[0]=t.toString(e[0]),r.apply(t,e)}},t.init=function(){for(var e in t)"function"==typeof t[e]&&"toString"!==e&&"toDate"!==e&&"extend"!==e&&"init"!==e&&t.extend(e,t[e])},t.toString=function(t){return"object"==typeof t?t=null!==t&&"function"==typeof t.toString?t.toString():"":("undefined"==typeof t||isNaN(t)&&!t.length)&&(t=""),""+t},t.toDate=function(t){return"[object Date]"===Object.prototype.toString.call(t)?t:(t=Date.parse(t),isNaN(t)?null:new Date(t))},t.toFloat=function(t){return parseFloat(t)},t.toInt=function(t,e){return parseInt(t,e||10)},t.toBoolean=function(t,e){return e?"1"===t||"true"===t:"0"!==t&&"false"!==t&&""!==t},t.equals=function(e,r){return e===t.toString(r)},t.contains=function(e,r){return e.indexOf(t.toString(r))>=0},t.matches=function(t,e,r){return"[object RegExp]"!==Object.prototype.toString.call(e)&&(e=new RegExp(e,r)),e.test(t)};var j={allow_display_name:!1,allow_utf8_local_part:!0,require_tld:!0};t.isEmail=function(e,n){if(n=r(n,j),n.allow_display_name){var l=e.match(s);l&&(e=l[1])}var f=e.split("@"),c=f.pop(),d=f.join("@"),g=c.toLowerCase();if(("gmail.com"===g||"googlemail.com"===g)&&(d=d.replace(/\./g,"").toLowerCase()),!t.isByteLength(d,{max:64})||!t.isByteLength(c,{max:256}))return!1;if(!t.isFQDN(c,{require_tld:n.require_tld}))return!1;if('"'===d[0])return d=d.slice(1,d.length-1),n.allow_utf8_local_part?u.test(d):o.test(d);for(var p=n.allow_utf8_local_part?a:i,F=d.split("."),_=0;_<F.length;_++)if(!p.test(F[_]))return!1;return!0};var B={protocols:["http","https","ftp"],require_tld:!0,require_protocol:!1,require_valid_protocol:!0,allow_underscores:!1,allow_trailing_dot:!1,allow_protocol_relative_urls:!1};t.isURL=function(e,n){if(!e||e.length>=2083||/\s/.test(e))return!1;if(0===e.indexOf("mailto:"))return!1;n=r(n,B);var i,o,a,u,s,l,f;if(f=e.split("://"),f.length>1){if(i=f.shift(),n.require_valid_protocol&&-1===n.protocols.indexOf(i))return!1}else{if(n.require_protocol)return!1;n.allow_protocol_relative_urls&&"//"===e.substr(0,2)&&(f[0]=e.substr(2))}return e=f.join("://"),f=e.split("#"),e=f.shift(),f=e.split("?"),e=f.shift(),f=e.split("/"),e=f.shift(),f=e.split("@"),f.length>1&&(o=f.shift(),o.indexOf(":")>=0&&o.split(":").length>2)?!1:(u=f.join("@"),f=u.split(":"),a=f.shift(),f.length&&(l=f.join(":"),s=parseInt(l,10),!/^[0-9]+$/.test(l)||0>=s||s>65535)?!1:t.isIP(a)||t.isFQDN(a,n)||"localhost"===a?n.host_whitelist&&-1===n.host_whitelist.indexOf(a)?!1:n.host_blacklist&&-1!==n.host_blacklist.indexOf(a)?!1:!0:!1)},t.isMACAddress=function(t){return g.test(t)},t.isIP=function(e,r){if(r=t.toString(r),!r)return t.isIP(e,4)||t.isIP(e,6);if("4"===r){if(!p.test(e))return!1;var n=e.split(".").sort(function(t,e){return t-e});return n[3]<=255}if("6"===r){var i=e.split(":"),o=!1,a=t.isIP(i[i.length-1],4),u=a?7:8;if(i.length>u)return!1;if("::"===e)return!0;"::"===e.substr(0,2)?(i.shift(),i.shift(),o=!0):"::"===e.substr(e.length-2)&&(i.pop(),i.pop(),o=!0);for(var s=0;s<i.length;++s)if(""===i[s]&&s>0&&s<i.length-1){if(o)return!1;o=!0}else if(a&&s==i.length-1);else if(!F.test(i[s]))return!1;return o?i.length>=1:i.length===u}return!1};var Z={require_tld:!0,allow_underscores:!1,allow_trailing_dot:!1};t.isFQDN=function(t,e){e=r(e,Z),e.allow_trailing_dot&&"."===t[t.length-1]&&(t=t.substring(0,t.length-1));var n=t.split(".");if(e.require_tld){var i=n.pop();if(!n.length||!/^([a-z\u00a1-\uffff]{2,}|xn[a-z0-9-]{2,})$/i.test(i))return!1}for(var o,a=0;a<n.length;a++){if(o=n[a],e.allow_underscores){if(o.indexOf("__")>=0)return!1;o=o.replace(/_/g,"")}if(!/^[a-z\u00a1-\uffff0-9-]+$/i.test(o))return!1;if(/[\uff01-\uff5e]/.test(o))return!1;if("-"===o[0]||"-"===o[o.length-1])return!1;if(o.indexOf("---")>=0&&"xn--"!==o.slice(0,4))return!1}return!0},t.isBoolean=function(t){return["true","false","1","0"].indexOf(t)>=0},t.isAlpha=function(t){return x.test(t)},t.isAlphanumeric=function(t){return h.test(t)},t.isNumeric=function(t){return m.test(t)},t.isDecimal=function(t){return""!==t&&w.test(t)},t.isHexadecimal=function(t){return $.test(t)},t.isHexColor=function(t){return b.test(t)},t.isLowercase=function(t){return t===t.toLowerCase()},t.isUppercase=function(t){return t===t.toUpperCase()},t.isInt=function(t,e){return e=e||{},v.test(t)&&(!e.hasOwnProperty("min")||t>=e.min)&&(!e.hasOwnProperty("max")||t<=e.max)},t.isFloat=function(t,e){return e=e||{},""===t||"."===t?!1:A.test(t)&&(!e.hasOwnProperty("min")||t>=e.min)&&(!e.hasOwnProperty("max")||t<=e.max)},t.isDivisibleBy=function(e,r){return t.toFloat(e)%t.toInt(r)===0},t.isNull=function(t){return 0===t.length},t.isLength=function(t,e){var r,n;"object"==typeof e?(r=e.min||0,n=e.max):(r=arguments[1],n=arguments[2]);var i=t.match(/[\uD800-\uDBFF][\uDC00-\uDFFF]/g)||[],o=t.length-i.length;return o>=r&&("undefined"==typeof n||n>=o)},t.isByteLength=function(t,e){var r,n;"object"==typeof e?(r=e.min||0,n=e.max):(r=arguments[1],n=arguments[2]);var i=encodeURI(t).split(/%..|./).length-1;return i>=r&&("undefined"==typeof n||n>=i)},t.isUUID=function(t,e){var r=_[e?e:"all"];return r&&r.test(t)},t.isDate=function(t){var r=new Date(Date.parse(t));if(isNaN(r))return!1;var n=e(t);if(null!==n){var i=r.getTimezoneOffset()-n;r=new Date(r.getTime()+6e4*i)}var o,a,u,s=String(r.getDate());return(a=t.match(/(^|[^:\d])[23]\d([^:\d]|$)/g))?(o=a.map(function(t){return t.match(/\d+/g)[0]}).join("/"),u=String(r.getFullYear()).slice(-2),o===s||o===u?!0:o===s+"/"+u||o===u+"/"+s?!0:!1):!0},t.isAfter=function(e,r){var n=t.toDate(r||new Date),i=t.toDate(e);return!!(i&&n&&i>n)},t.isBefore=function(e,r){var n=t.toDate(r||new Date),i=t.toDate(e);return!!(i&&n&&n>i)},t.isIn=function(e,r){var n;if("[object Array]"===Object.prototype.toString.call(r)){var i=[];for(n in r)i[n]=t.toString(r[n]);return i.indexOf(e)>=0}return"object"==typeof r?r.hasOwnProperty(e):r&&"function"==typeof r.indexOf?r.indexOf(e)>=0:!1},t.isWhitelisted=function(t,e){for(var r=t.length-1;r>=0;r--)if(-1===e.indexOf(t[r]))return!1;return!0},t.isCreditCard=function(t){var e=t.replace(/[^0-9]+/g,"");if(!l.test(e))return!1;for(var r,n,i,o=0,a=e.length-1;a>=0;a--)r=e.substring(a,a+1),n=parseInt(r,10),i?(n*=2,o+=n>=10?n%10+1:n):o+=n,i=!i;return!!(o%10===0?e:!1)},t.isISIN=function(t){if(!f.test(t))return!1;for(var e,r,n=t.replace(/[A-Z]/g,function(t){return parseInt(t,36)}),i=0,o=!0,a=n.length-2;a>=0;a--)e=n.substring(a,a+1),r=parseInt(e,10),o?(r*=2,i+=r>=10?r+1:r):i+=r,o=!o;return parseInt(t.substr(t.length-1),10)===(1e4-i)%10},t.isISBN=function(e,r){if(r=t.toString(r),!r)return t.isISBN(e,10)||t.isISBN(e,13);var n,i=e.replace(/[\s-]+/g,""),o=0;if("10"===r){if(!c.test(i))return!1;for(n=0;9>n;n++)o+=(n+1)*i.charAt(n);if(o+="X"===i.charAt(9)?100:10*i.charAt(9),o%11===0)return!!i}else if("13"===r){if(!d.test(i))return!1;var a=[1,3];for(n=0;12>n;n++)o+=a[n%2]*i.charAt(n);if(i.charAt(12)-(10-o%10)%10===0)return!!i}return!1},t.isMobilePhone=function(t,e){return e in C?C[e].test(t):!1};var z={symbol:"$",require_symbol:!1,allow_space_after_symbol:!1,symbol_after_digits:!1,allow_negatives:!0,parens_for_negatives:!1,negative_sign_before_digits:!1,negative_sign_after_digits:!1,allow_negative_sign_placeholder:!1,thousands_separator:",",decimal_separator:".",allow_space_after_digits:!1};t.isCurrency=function(t,e){return e=r(e,z),n(e).test(t)},t.isJSON=function(t){try{var e=JSON.parse(t);return!!e&&"object"==typeof e}catch(r){}return!1},t.isMultibyte=function(t){return y.test(t)},t.isAscii=function(t){return D.test(t)},t.isFullWidth=function(t){return I.test(t)},t.isHalfWidth=function(t){return O.test(t)},t.isVariableWidth=function(t){return I.test(t)&&O.test(t)},t.isSurrogatePair=function(t){return E.test(t)},t.isBase64=function(t){return S.test(t)},t.isMongoId=function(e){return t.isHexadecimal(e)&&24===e.length},t.isISO8601=function(t){return N.test(t)},t.ltrim=function(t,e){var r=e?new RegExp("^["+e+"]+","g"):/^\s+/g;return t.replace(r,"")},t.rtrim=function(t,e){var r=e?new RegExp("["+e+"]+$","g"):/\s+$/g;return t.replace(r,"")},t.trim=function(t,e){var r=e?new RegExp("^["+e+"]+|["+e+"]+$","g"):/^\s+|\s+$/g;return t.replace(r,"")},t.escape=function(t){return t.replace(/&/g,"&amp;").replace(/"/g,"&quot;").replace(/'/g,"&#x27;").replace(/</g,"&lt;").replace(/>/g,"&gt;").replace(/\//g,"&#x2F;").replace(/\`/g,"&#96;")},t.stripLow=function(e,r){var n=r?"\\x00-\\x09\\x0B\\x0C\\x0E-\\x1F\\x7F":"\\x00-\\x1F\\x7F";return t.blacklist(e,n)},t.whitelist=function(t,e){return t.replace(new RegExp("[^"+e+"]+","g"),"")},t.blacklist=function(t,e){return t.replace(new RegExp("["+e+"]+","g"),"")};var q={lowercase:!0,remove_dots:!0,remove_extension:!0};return t.normalizeEmail=function(e,n){if(n=r(n,q),!t.isEmail(e))return!1;var i=e.split("@",2);if(i[1]=i[1].toLowerCase(),"gmail.com"===i[1]||"googlemail.com"===i[1]){if(n.remove_extension&&(i[0]=i[0].split("+")[0]),n.remove_dots&&(i[0]=i[0].replace(/\./g,"")),!i[0].length)return!1;i[0]=i[0].toLowerCase(),i[1]="gmail.com"}else n.lowercase&&(i[0]=i[0].toLowerCase());return i.join("@")},t.init(),t});

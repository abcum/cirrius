!function(t){if("object"==typeof exports)module.exports=t();else if("function"==typeof define&&define.amd)define(t);else{var r;"undefined"!=typeof window?r=window:"undefined"!=typeof global?r=global:"undefined"!=typeof self&&(r=self),r.S=t()}}(function(){var t;return function r(t,e,n){function i(o,u){if(!e[o]){if(!t[o]){var a="function"==typeof require&&require;if(!u&&a)return a(o,!0);if(s)return s(o,!0);throw new Error("Cannot find module '"+o+"'")}var c=e[o]={exports:{}};t[o][0].call(c.exports,function(r){var e=t[o][1][r];return i(e?e:r)},c,c.exports,r,t,e,n)}return e[o].exports}for(var s="function"==typeof require&&require,o=0;o<n.length;o++)i(n[o]);return i}({1:[function(t,r){function e(t,r){for(var e=0,n=t.indexOf(r);n>=0;)e+=1,n=t.indexOf(r,n+1);return e}r.exports=e},{}],2:[function(r,e){!function(){"use strict";function n(t,r){t.s=null!==r&&void 0!==r?"string"==typeof r?r:r.toString():r,t.orig=r,null!==r&&void 0!==r?t.__defineGetter__?t.__defineGetter__("length",function(){return t.s.length}):t.length=r.length:t.length=-1}function i(t){n(this,t)}function s(){for(var t in v)!function(t){var r=v[t];m.hasOwnProperty(t)||(w.push(t),m[t]=function(){return String.prototype.s=this,r.apply(this,arguments)})}(t)}function o(){for(var t=0;t<w.length;++t)delete String.prototype[w[t]];w.length=0}function u(){for(var t=a(),r={},e=0;e<t.length;++e){var n=t[e],i=m[n];try{var s=typeof i.apply("test",["string"]);r[n]=s}catch(o){}}return r}function a(){var t=[];if(Object.getOwnPropertyNames)return t=Object.getOwnPropertyNames(m),t.splice(t.indexOf("valueOf"),1),t.splice(t.indexOf("toString"),1),t;var r={};for(var e in String.prototype)r[e]=e;for(var e in Object.prototype)delete r[e];for(var e in r)t.push(e);return t}function c(t){return new i(t)}function h(t,r){var e,n=[];for(e=0;e<t.length;e++)n.push(t[e]),r&&r.call(t,t[e],e);return n}function l(t){var r,e,n=[],i=/^[A-Za-z0-9]+$/;for(t=f(t),e=0;e<t.length;++e)r=t.charAt(e),n.push(i.test(r)?r:"\\000"===r?"\\000":"\\"+r);return n.join("")}function f(t){return null==t?"":""+t}var p="3.1.2",g={},d={"Á":"A","Ă":"A","Ắ":"A","Ặ":"A","Ằ":"A","Ẳ":"A","Ẵ":"A","Ǎ":"A","Â":"A","Ấ":"A","Ậ":"A","Ầ":"A","Ẩ":"A","Ẫ":"A","Ä":"A","Ǟ":"A","Ȧ":"A","Ǡ":"A","Ạ":"A","Ȁ":"A","À":"A","Ả":"A","Ȃ":"A","Ā":"A","Ą":"A","Å":"A","Ǻ":"A","Ḁ":"A","Ⱥ":"A","Ã":"A","Ꜳ":"AA","Æ":"AE","Ǽ":"AE","Ǣ":"AE","Ꜵ":"AO","Ꜷ":"AU","Ꜹ":"AV","Ꜻ":"AV","Ꜽ":"AY","Ḃ":"B","Ḅ":"B","Ɓ":"B","Ḇ":"B","Ƀ":"B","Ƃ":"B","Ć":"C","Č":"C","Ç":"C","Ḉ":"C","Ĉ":"C","Ċ":"C","Ƈ":"C","Ȼ":"C","Ď":"D","Ḑ":"D","Ḓ":"D","Ḋ":"D","Ḍ":"D","Ɗ":"D","Ḏ":"D","ǲ":"D","ǅ":"D","Đ":"D","Ƌ":"D","Ǳ":"DZ","Ǆ":"DZ","É":"E","Ĕ":"E","Ě":"E","Ȩ":"E","Ḝ":"E","Ê":"E","Ế":"E","Ệ":"E","Ề":"E","Ể":"E","Ễ":"E","Ḙ":"E","Ë":"E","Ė":"E","Ẹ":"E","Ȅ":"E","È":"E","Ẻ":"E","Ȇ":"E","Ē":"E","Ḗ":"E","Ḕ":"E","Ę":"E","Ɇ":"E","Ẽ":"E","Ḛ":"E","Ꝫ":"ET","Ḟ":"F","Ƒ":"F","Ǵ":"G","Ğ":"G","Ǧ":"G","Ģ":"G","Ĝ":"G","Ġ":"G","Ɠ":"G","Ḡ":"G","Ǥ":"G","Ḫ":"H","Ȟ":"H","Ḩ":"H","Ĥ":"H","Ⱨ":"H","Ḧ":"H","Ḣ":"H","Ḥ":"H","Ħ":"H","Í":"I","Ĭ":"I","Ǐ":"I","Î":"I","Ï":"I","Ḯ":"I","İ":"I","Ị":"I","Ȉ":"I","Ì":"I","Ỉ":"I","Ȋ":"I","Ī":"I","Į":"I","Ɨ":"I","Ĩ":"I","Ḭ":"I","Ꝺ":"D","Ꝼ":"F","Ᵹ":"G","Ꞃ":"R","Ꞅ":"S","Ꞇ":"T","Ꝭ":"IS","Ĵ":"J","Ɉ":"J","Ḱ":"K","Ǩ":"K","Ķ":"K","Ⱪ":"K","Ꝃ":"K","Ḳ":"K","Ƙ":"K","Ḵ":"K","Ꝁ":"K","Ꝅ":"K","Ĺ":"L","Ƚ":"L","Ľ":"L","Ļ":"L","Ḽ":"L","Ḷ":"L","Ḹ":"L","Ⱡ":"L","Ꝉ":"L","Ḻ":"L","Ŀ":"L","Ɫ":"L","ǈ":"L","Ł":"L","Ǉ":"LJ","Ḿ":"M","Ṁ":"M","Ṃ":"M","Ɱ":"M","Ń":"N","Ň":"N","Ņ":"N","Ṋ":"N","Ṅ":"N","Ṇ":"N","Ǹ":"N","Ɲ":"N","Ṉ":"N","Ƞ":"N","ǋ":"N","Ñ":"N","Ǌ":"NJ","Ó":"O","Ŏ":"O","Ǒ":"O","Ô":"O","Ố":"O","Ộ":"O","Ồ":"O","Ổ":"O","Ỗ":"O","Ö":"O","Ȫ":"O","Ȯ":"O","Ȱ":"O","Ọ":"O","Ő":"O","Ȍ":"O","Ò":"O","Ỏ":"O","Ơ":"O","Ớ":"O","Ợ":"O","Ờ":"O","Ở":"O","Ỡ":"O","Ȏ":"O","Ꝋ":"O","Ꝍ":"O","Ō":"O","Ṓ":"O","Ṑ":"O","Ɵ":"O","Ǫ":"O","Ǭ":"O","Ø":"O","Ǿ":"O","Õ":"O","Ṍ":"O","Ṏ":"O","Ȭ":"O","Ƣ":"OI","Ꝏ":"OO","Ɛ":"E","Ɔ":"O","Ȣ":"OU","Ṕ":"P","Ṗ":"P","Ꝓ":"P","Ƥ":"P","Ꝕ":"P","Ᵽ":"P","Ꝑ":"P","Ꝙ":"Q","Ꝗ":"Q","Ŕ":"R","Ř":"R","Ŗ":"R","Ṙ":"R","Ṛ":"R","Ṝ":"R","Ȑ":"R","Ȓ":"R","Ṟ":"R","Ɍ":"R","Ɽ":"R","Ꜿ":"C","Ǝ":"E","Ś":"S","Ṥ":"S","Š":"S","Ṧ":"S","Ş":"S","Ŝ":"S","Ș":"S","Ṡ":"S","Ṣ":"S","Ṩ":"S","ẞ":"SS","Ť":"T","Ţ":"T","Ṱ":"T","Ț":"T","Ⱦ":"T","Ṫ":"T","Ṭ":"T","Ƭ":"T","Ṯ":"T","Ʈ":"T","Ŧ":"T","Ɐ":"A","Ꞁ":"L","Ɯ":"M","Ʌ":"V","Ꜩ":"TZ","Ú":"U","Ŭ":"U","Ǔ":"U","Û":"U","Ṷ":"U","Ü":"U","Ǘ":"U","Ǚ":"U","Ǜ":"U","Ǖ":"U","Ṳ":"U","Ụ":"U","Ű":"U","Ȕ":"U","Ù":"U","Ủ":"U","Ư":"U","Ứ":"U","Ự":"U","Ừ":"U","Ử":"U","Ữ":"U","Ȗ":"U","Ū":"U","Ṻ":"U","Ų":"U","Ů":"U","Ũ":"U","Ṹ":"U","Ṵ":"U","Ꝟ":"V","Ṿ":"V","Ʋ":"V","Ṽ":"V","Ꝡ":"VY","Ẃ":"W","Ŵ":"W","Ẅ":"W","Ẇ":"W","Ẉ":"W","Ẁ":"W","Ⱳ":"W","Ẍ":"X","Ẋ":"X","Ý":"Y","Ŷ":"Y","Ÿ":"Y","Ẏ":"Y","Ỵ":"Y","Ỳ":"Y","Ƴ":"Y","Ỷ":"Y","Ỿ":"Y","Ȳ":"Y","Ɏ":"Y","Ỹ":"Y","Ź":"Z","Ž":"Z","Ẑ":"Z","Ⱬ":"Z","Ż":"Z","Ẓ":"Z","Ȥ":"Z","Ẕ":"Z","Ƶ":"Z","Ĳ":"IJ","Œ":"OE","ᴀ":"A","ᴁ":"AE","ʙ":"B","ᴃ":"B","ᴄ":"C","ᴅ":"D","ᴇ":"E","ꜰ":"F","ɢ":"G","ʛ":"G","ʜ":"H","ɪ":"I","ʁ":"R","ᴊ":"J","ᴋ":"K","ʟ":"L","ᴌ":"L","ᴍ":"M","ɴ":"N","ᴏ":"O","ɶ":"OE","ᴐ":"O","ᴕ":"OU","ᴘ":"P","ʀ":"R","ᴎ":"N","ᴙ":"R","ꜱ":"S","ᴛ":"T","ⱻ":"E","ᴚ":"R","ᴜ":"U","ᴠ":"V","ᴡ":"W","ʏ":"Y","ᴢ":"Z","á":"a","ă":"a","ắ":"a","ặ":"a","ằ":"a","ẳ":"a","ẵ":"a","ǎ":"a","â":"a","ấ":"a","ậ":"a","ầ":"a","ẩ":"a","ẫ":"a","ä":"a","ǟ":"a","ȧ":"a","ǡ":"a","ạ":"a","ȁ":"a","à":"a","ả":"a","ȃ":"a","ā":"a","ą":"a","ᶏ":"a","ẚ":"a","å":"a","ǻ":"a","ḁ":"a","ⱥ":"a","ã":"a","ꜳ":"aa","æ":"ae","ǽ":"ae","ǣ":"ae","ꜵ":"ao","ꜷ":"au","ꜹ":"av","ꜻ":"av","ꜽ":"ay","ḃ":"b","ḅ":"b","ɓ":"b","ḇ":"b","ᵬ":"b","ᶀ":"b","ƀ":"b","ƃ":"b","ɵ":"o","ć":"c","č":"c","ç":"c","ḉ":"c","ĉ":"c","ɕ":"c","ċ":"c","ƈ":"c","ȼ":"c","ď":"d","ḑ":"d","ḓ":"d","ȡ":"d","ḋ":"d","ḍ":"d","ɗ":"d","ᶑ":"d","ḏ":"d","ᵭ":"d","ᶁ":"d","đ":"d","ɖ":"d","ƌ":"d","ı":"i","ȷ":"j","ɟ":"j","ʄ":"j","ǳ":"dz","ǆ":"dz","é":"e","ĕ":"e","ě":"e","ȩ":"e","ḝ":"e","ê":"e","ế":"e","ệ":"e","ề":"e","ể":"e","ễ":"e","ḙ":"e","ë":"e","ė":"e","ẹ":"e","ȅ":"e","è":"e","ẻ":"e","ȇ":"e","ē":"e","ḗ":"e","ḕ":"e","ⱸ":"e","ę":"e","ᶒ":"e","ɇ":"e","ẽ":"e","ḛ":"e","ꝫ":"et","ḟ":"f","ƒ":"f","ᵮ":"f","ᶂ":"f","ǵ":"g","ğ":"g","ǧ":"g","ģ":"g","ĝ":"g","ġ":"g","ɠ":"g","ḡ":"g","ᶃ":"g","ǥ":"g","ḫ":"h","ȟ":"h","ḩ":"h","ĥ":"h","ⱨ":"h","ḧ":"h","ḣ":"h","ḥ":"h","ɦ":"h","ẖ":"h","ħ":"h","ƕ":"hv","í":"i","ĭ":"i","ǐ":"i","î":"i","ï":"i","ḯ":"i","ị":"i","ȉ":"i","ì":"i","ỉ":"i","ȋ":"i","ī":"i","į":"i","ᶖ":"i","ɨ":"i","ĩ":"i","ḭ":"i","ꝺ":"d","ꝼ":"f","ᵹ":"g","ꞃ":"r","ꞅ":"s","ꞇ":"t","ꝭ":"is","ǰ":"j","ĵ":"j","ʝ":"j","ɉ":"j","ḱ":"k","ǩ":"k","ķ":"k","ⱪ":"k","ꝃ":"k","ḳ":"k","ƙ":"k","ḵ":"k","ᶄ":"k","ꝁ":"k","ꝅ":"k","ĺ":"l","ƚ":"l","ɬ":"l","ľ":"l","ļ":"l","ḽ":"l","ȴ":"l","ḷ":"l","ḹ":"l","ⱡ":"l","ꝉ":"l","ḻ":"l","ŀ":"l","ɫ":"l","ᶅ":"l","ɭ":"l","ł":"l","ǉ":"lj","ſ":"s","ẜ":"s","ẛ":"s","ẝ":"s","ḿ":"m","ṁ":"m","ṃ":"m","ɱ":"m","ᵯ":"m","ᶆ":"m","ń":"n","ň":"n","ņ":"n","ṋ":"n","ȵ":"n","ṅ":"n","ṇ":"n","ǹ":"n","ɲ":"n","ṉ":"n","ƞ":"n","ᵰ":"n","ᶇ":"n","ɳ":"n","ñ":"n","ǌ":"nj","ó":"o","ŏ":"o","ǒ":"o","ô":"o","ố":"o","ộ":"o","ồ":"o","ổ":"o","ỗ":"o","ö":"o","ȫ":"o","ȯ":"o","ȱ":"o","ọ":"o","ő":"o","ȍ":"o","ò":"o","ỏ":"o","ơ":"o","ớ":"o","ợ":"o","ờ":"o","ở":"o","ỡ":"o","ȏ":"o","ꝋ":"o","ꝍ":"o","ⱺ":"o","ō":"o","ṓ":"o","ṑ":"o","ǫ":"o","ǭ":"o","ø":"o","ǿ":"o","õ":"o","ṍ":"o","ṏ":"o","ȭ":"o","ƣ":"oi","ꝏ":"oo","ɛ":"e","ᶓ":"e","ɔ":"o","ᶗ":"o","ȣ":"ou","ṕ":"p","ṗ":"p","ꝓ":"p","ƥ":"p","ᵱ":"p","ᶈ":"p","ꝕ":"p","ᵽ":"p","ꝑ":"p","ꝙ":"q","ʠ":"q","ɋ":"q","ꝗ":"q","ŕ":"r","ř":"r","ŗ":"r","ṙ":"r","ṛ":"r","ṝ":"r","ȑ":"r","ɾ":"r","ᵳ":"r","ȓ":"r","ṟ":"r","ɼ":"r","ᵲ":"r","ᶉ":"r","ɍ":"r","ɽ":"r","ↄ":"c","ꜿ":"c","ɘ":"e","ɿ":"r","ś":"s","ṥ":"s","š":"s","ṧ":"s","ş":"s","ŝ":"s","ș":"s","ṡ":"s","ṣ":"s","ṩ":"s","ʂ":"s","ᵴ":"s","ᶊ":"s","ȿ":"s","ɡ":"g","ß":"ss","ᴑ":"o","ᴓ":"o","ᴝ":"u","ť":"t","ţ":"t","ṱ":"t","ț":"t","ȶ":"t","ẗ":"t","ⱦ":"t","ṫ":"t","ṭ":"t","ƭ":"t","ṯ":"t","ᵵ":"t","ƫ":"t","ʈ":"t","ŧ":"t","ᵺ":"th","ɐ":"a","ᴂ":"ae","ǝ":"e","ᵷ":"g","ɥ":"h","ʮ":"h","ʯ":"h","ᴉ":"i","ʞ":"k","ꞁ":"l","ɯ":"m","ɰ":"m","ᴔ":"oe","ɹ":"r","ɻ":"r","ɺ":"r","ⱹ":"r","ʇ":"t","ʌ":"v","ʍ":"w","ʎ":"y","ꜩ":"tz","ú":"u","ŭ":"u","ǔ":"u","û":"u","ṷ":"u","ü":"u","ǘ":"u","ǚ":"u","ǜ":"u","ǖ":"u","ṳ":"u","ụ":"u","ű":"u","ȕ":"u","ù":"u","ủ":"u","ư":"u","ứ":"u","ự":"u","ừ":"u","ử":"u","ữ":"u","ȗ":"u","ū":"u","ṻ":"u","ų":"u","ᶙ":"u","ů":"u","ũ":"u","ṹ":"u","ṵ":"u","ᵫ":"ue","ꝸ":"um","ⱴ":"v","ꝟ":"v","ṿ":"v","ʋ":"v","ᶌ":"v","ⱱ":"v","ṽ":"v","ꝡ":"vy","ẃ":"w","ŵ":"w","ẅ":"w","ẇ":"w","ẉ":"w","ẁ":"w","ⱳ":"w","ẘ":"w","ẍ":"x","ẋ":"x","ᶍ":"x","ý":"y","ŷ":"y","ÿ":"y","ẏ":"y","ỵ":"y","ỳ":"y","ƴ":"y","ỷ":"y","ỿ":"y","ȳ":"y","ẙ":"y","ɏ":"y","ỹ":"y","ź":"z","ž":"z","ẑ":"z","ʑ":"z","ⱬ":"z","ż":"z","ẓ":"z","ȥ":"z","ẕ":"z","ᵶ":"z","ᶎ":"z","ʐ":"z","ƶ":"z","ɀ":"z","ﬀ":"ff","ﬃ":"ffi","ﬄ":"ffl","ﬁ":"fi","ﬂ":"fl","ĳ":"ij","œ":"oe","ﬆ":"st","ₐ":"a","ₑ":"e","ᵢ":"i","ⱼ":"j","ₒ":"o","ᵣ":"r","ᵤ":"u","ᵥ":"v","ₓ":"x"},m=String.prototype,v=i.prototype={between:function(t,r){var e=this.s,n=e.indexOf(t),i=e.indexOf(r,n+t.length);return new this.constructor(-1==i&&null!=r?"":-1==i&&null==r?e.substring(n+t.length):e.slice(n+t.length,i))},camelize:function(){var t=this.trim().s.replace(/(\-|_|\s)+(.)?/g,function(t,r,e){return e?e.toUpperCase():""});return new this.constructor(t)},capitalize:function(){return new this.constructor(this.s.substr(0,1).toUpperCase()+this.s.substring(1).toLowerCase())},charAt:function(t){return this.s.charAt(t)},chompLeft:function(t){var r=this.s;return 0===r.indexOf(t)?(r=r.slice(t.length),new this.constructor(r)):this},chompRight:function(t){if(this.endsWith(t)){var r=this.s;return r=r.slice(0,r.length-t.length),new this.constructor(r)}return this},collapseWhitespace:function(){var t=this.s.replace(/[\s\xa0]+/g," ").replace(/^\s+|\s+$/g,"");return new this.constructor(t)},contains:function(t){return this.s.indexOf(t)>=0},count:function(t){return r("./_count")(this.s,t)},dasherize:function(){var t=this.trim().s.replace(/[_\s]+/g,"-").replace(/([A-Z])/g,"-$1").replace(/-+/g,"-").toLowerCase();return new this.constructor(t)},latinise:function(){var t=this.replace(/[^A-Za-z0-9\[\] ]/g,function(t){return d[t]||t});return new this.constructor(t)},decodeHtmlEntities:function(){var t=this.s;return t=t.replace(/&#(\d+);?/g,function(t,r){return String.fromCharCode(r)}).replace(/&#[xX]([A-Fa-f0-9]+);?/g,function(t,r){return String.fromCharCode(parseInt(r,16))}).replace(/&([^;\W]+;?)/g,function(t,r){var e=r.replace(/;$/,""),n=g[r]||r.match(/;$/)&&g[e];return"number"==typeof n?String.fromCharCode(n):"string"==typeof n?n:t}),new this.constructor(t)},endsWith:function(){for(var t=Array.prototype.slice.call(arguments,0),r=0;r<t.length;++r){var e=this.s.length-t[r].length;if(e>=0&&this.s.indexOf(t[r],e)===e)return!0}return!1},escapeHTML:function(){return new this.constructor(this.s.replace(/[&<>"']/g,function(t){return"&"+E[t]+";"}))},ensureLeft:function(t){var r=this.s;return 0===r.indexOf(t)?this:new this.constructor(t+r)},ensureRight:function(t){var r=this.s;return this.endsWith(t)?this:new this.constructor(r+t)},humanize:function(){if(null===this.s||void 0===this.s)return new this.constructor("");var t=this.underscore().replace(/_id$/,"").replace(/_/g," ").trim().capitalize();return new this.constructor(t)},isAlpha:function(){return!/[^a-z\xDF-\xFF]|^$/.test(this.s.toLowerCase())},isAlphaNumeric:function(){return!/[^0-9a-z\xDF-\xFF]/.test(this.s.toLowerCase())},isEmpty:function(){return null===this.s||void 0===this.s?!0:/^[\s\xa0]*$/.test(this.s)},isLower:function(){return this.isAlpha()&&this.s.toLowerCase()===this.s},isNumeric:function(){return!/[^0-9]/.test(this.s)},isUpper:function(){return this.isAlpha()&&this.s.toUpperCase()===this.s},left:function(t){if(t>=0){var r=this.s.substr(0,t);return new this.constructor(r)}return this.right(-t)},lines:function(){return this.replaceAll("\r\n","\n").s.split("\n")},pad:function(t,r){if(null==r&&(r=" "),this.s.length>=t)return new this.constructor(this.s);t-=this.s.length;var e=Array(Math.ceil(t/2)+1).join(r),n=Array(Math.floor(t/2)+1).join(r);return new this.constructor(e+this.s+n)},padLeft:function(t,r){return null==r&&(r=" "),new this.constructor(this.s.length>=t?this.s:Array(t-this.s.length+1).join(r)+this.s)},padRight:function(t,r){return null==r&&(r=" "),new this.constructor(this.s.length>=t?this.s:this.s+Array(t-this.s.length+1).join(r))},parseCSV:function(t,r,e,n){t=t||",",e=e||"\\","undefined"==typeof r&&(r='"');var i=0,s=[],o=[],u=this.s.length,a=!1,c=!1,h=this,l=function(t){return h.s.charAt(t)};if("undefined"!=typeof n)var f=[];for(r||(a=!0);u>i;){var p=l(i);switch(p){case e:if(a&&(e!==r||l(i+1)===r)){i+=1,s.push(l(i));break}if(e!==r)break;case r:a=!a;break;case t:c&&(a=!1,c=!1),a&&r?s.push(p):(o.push(s.join("")),s.length=0);break;case n:c?(a=!1,c=!1,o.push(s.join("")),f.push(o),o=[],s.length=0):a?s.push(p):f&&(o.push(s.join("")),f.push(o),o=[],s.length=0);break;case" ":a&&s.push(p);break;default:a?s.push(p):p!==r&&(s.push(p),a=!0,c=!0)}i+=1}return o.push(s.join("")),f?(f.push(o),f):o},replaceAll:function(t,r){var e=this.s.split(t).join(r);return new this.constructor(e)},strip:function(){for(var t=this.s,r=0,e=arguments.length;e>r;r++)t=t.split(arguments[r]).join("");return new this.constructor(t)},stripLeft:function(t){var r,e,n=f(this.s);return void 0===t?e=/^\s+/g:(r=l(t),e=new RegExp("^["+r+"]+","g")),new this.constructor(n.replace(e,""))},stripRight:function(t){var r,e,n=f(this.s);return void 0===t?e=/\s+$/g:(r=l(t),e=new RegExp("["+r+"]+$","g")),new this.constructor(n.replace(e,""))},right:function(t){if(t>=0){var r=this.s.substr(this.s.length-t,t);return new this.constructor(r)}return this.left(-t)},setValue:function(t){return n(this,t),this},slugify:function(){var t=new i(new i(this.s).latinise().s.replace(/[^\w\s-]/g,"").toLowerCase()).dasherize().s;return"-"===t.charAt(0)&&(t=t.substr(1)),new this.constructor(t)},startsWith:function(){for(var t=Array.prototype.slice.call(arguments,0),r=0;r<t.length;++r)if(0===this.s.lastIndexOf(t[r],0))return!0;return!1},stripPunctuation:function(){return new this.constructor(this.s.replace(/[^\w\s]|_/g,"").replace(/\s+/g," "))},stripTags:function(){var t=this.s,r=arguments.length>0?arguments:[""];return h(r,function(r){t=t.replace(RegExp("</?"+r+"[^<>]*>","gi"),"")}),new this.constructor(t)},template:function(t,r,e){var n=this.s,r=r||c.TMPL_OPEN,e=e||c.TMPL_CLOSE,i=r.replace(/[-[\]()*\s]/g,"\\$&").replace(/\$/g,"\\$"),s=e.replace(/[-[\]()*\s]/g,"\\$&").replace(/\$/g,"\\$"),o=new RegExp(i+"(.+?)"+s,"g"),u=n.match(o)||[];return u.forEach(function(i){var s=i.substring(r.length,i.length-e.length).trim(),o="undefined"==typeof t[s]?"":t[s];n=n.replace(i,o)}),new this.constructor(n)},times:function(t){return new this.constructor(new Array(t+1).join(this.s))},toBoolean:function(){if("string"==typeof this.orig){var t=this.s.toLowerCase();return"true"===t||"yes"===t||"on"===t||"1"===t}return this.orig===!0||1===this.orig},toFloat:function(t){var r=parseFloat(this.s);return t?parseFloat(r.toFixed(t)):r},toInt:function(){return/^\s*-?0x/i.test(this.s)?parseInt(this.s,16):parseInt(this.s,10)},trim:function(){var t;return t="undefined"==typeof m.trim?this.s.replace(/(^\s*|\s*$)/g,""):this.s.trim(),new this.constructor(t)},trimLeft:function(){var t;return t=m.trimLeft?this.s.trimLeft():this.s.replace(/(^\s*)/g,""),new this.constructor(t)},trimRight:function(){var t;return t=m.trimRight?this.s.trimRight():this.s.replace(/\s+$/,""),new this.constructor(t)},truncate:function(t,r){var e=this.s;if(t=~~t,r=r||"...",e.length<=t)return new this.constructor(e);var n=function(t){return t.toUpperCase()!==t.toLowerCase()?"A":" "},s=e.slice(0,t+1).replace(/.(?=\W*\w*$)/g,n);return s=s.slice(s.length-2).match(/\w\w/)?s.replace(/\s*\S+$/,""):new i(s.slice(0,s.length-1)).trimRight().s,new i((s+r).length>e.length?e:e.slice(0,s.length)+r)},toCSV:function(){function t(t){return null!==t&&""!==t}var r=",",e='"',n="\\",s=!0,o=!1,u=[];if("object"==typeof arguments[0]?(r=arguments[0].delimiter||r,r=arguments[0].separator||r,e=arguments[0].qualifier||e,s=!!arguments[0].encloseNumbers,n=arguments[0].escape||n,o=!!arguments[0].keys):"string"==typeof arguments[0]&&(r=arguments[0]),"string"==typeof arguments[1]&&(e=arguments[1]),null===arguments[1]&&(e=null),this.orig instanceof Array)u=this.orig;else for(var a in this.orig)this.orig.hasOwnProperty(a)&&u.push(o?a:this.orig[a]);for(var c=n+e,h=[],l=0;l<u.length;++l){var f=t(e);if("number"==typeof u[l]&&(f&=s),f&&h.push(e),null!==u[l]&&void 0!==u[l]){var p=new i(u[l]).replaceAll(e,c).s;h.push(p)}else h.push("");f&&h.push(e),r&&h.push(r)}return h.length=h.length-1,new this.constructor(h.join(""))},toString:function(){return this.s},underscore:function(){var t=this.trim().s.replace(/([a-z\d])([A-Z]+)/g,"$1_$2").replace(/([A-Z\d]+)([A-Z][a-z])/,"$1_$2").replace(/[-\s]+/g,"_").toLowerCase();return new this.constructor(t)},unescapeHTML:function(){return new this.constructor(this.s.replace(/\&([^;]+);/g,function(t,r){var e;return r in y?y[r]:(e=r.match(/^#x([\da-fA-F]+)$/))?String.fromCharCode(parseInt(e[1],16)):(e=r.match(/^#(\d+)$/))?String.fromCharCode(~~e[1]):t}))},valueOf:function(){return this.s.valueOf()},wrapHTML:function(t,r){var e=this.s,n=null==t?"span":t,i="",s="";if("object"==typeof r)for(var o in r)i+=" "+o+'="'+new this.constructor(r[o]).escapeHTML()+'"';return e=s.concat("<",n,i,">",this,"</",n,">"),new this.constructor(e)}},w=[],O=u();for(var A in O)!function(t){var r=m[t];"function"==typeof r&&(v[t]||(v[t]="string"===O[t]?function(){return new this.constructor(r.apply(this,arguments))}:r))}(A);v.repeat=v.times,v.include=v.contains,v.toInteger=v.toInt,v.toBool=v.toBoolean,v.decodeHTMLEntities=v.decodeHtmlEntities,v.constructor=i,c.extendPrototype=s,c.restorePrototype=o,c.VERSION=p,c.TMPL_OPEN="{{",c.TMPL_CLOSE="}}",c.ENTITIES=g,"undefined"!=typeof e&&"undefined"!=typeof e.exports?e.exports=c:"function"==typeof t&&t.amd?t([],function(){return c}):window.S=c;var y={lt:"<",gt:">",quot:'"',apos:"'",amp:"&"},E={};for(var b in y)E[y[b]]=b;g={amp:"&",gt:">",lt:"<",quot:'"',apos:"'",AElig:198,Aacute:193,Acirc:194,Agrave:192,Aring:197,Atilde:195,Auml:196,Ccedil:199,ETH:208,Eacute:201,Ecirc:202,Egrave:200,Euml:203,Iacute:205,Icirc:206,Igrave:204,Iuml:207,Ntilde:209,Oacute:211,Ocirc:212,Ograve:210,Oslash:216,Otilde:213,Ouml:214,THORN:222,Uacute:218,Ucirc:219,Ugrave:217,Uuml:220,Yacute:221,aacute:225,acirc:226,aelig:230,agrave:224,aring:229,atilde:227,auml:228,ccedil:231,eacute:233,ecirc:234,egrave:232,eth:240,euml:235,iacute:237,icirc:238,igrave:236,iuml:239,ntilde:241,oacute:243,ocirc:244,ograve:242,oslash:248,otilde:245,ouml:246,szlig:223,thorn:254,uacute:250,ucirc:251,ugrave:249,uuml:252,yacute:253,yuml:255,copy:169,reg:174,nbsp:160,iexcl:161,cent:162,pound:163,curren:164,yen:165,brvbar:166,sect:167,uml:168,ordf:170,laquo:171,not:172,shy:173,macr:175,deg:176,plusmn:177,sup1:185,sup2:178,sup3:179,acute:180,micro:181,para:182,middot:183,cedil:184,ordm:186,raquo:187,frac14:188,frac12:189,frac34:190,iquest:191,times:215,divide:247,"OElig;":338,"oelig;":339,"Scaron;":352,"scaron;":353,"Yuml;":376,"fnof;":402,"circ;":710,"tilde;":732,"Alpha;":913,"Beta;":914,"Gamma;":915,"Delta;":916,"Epsilon;":917,"Zeta;":918,"Eta;":919,"Theta;":920,"Iota;":921,"Kappa;":922,"Lambda;":923,"Mu;":924,"Nu;":925,"Xi;":926,"Omicron;":927,"Pi;":928,"Rho;":929,"Sigma;":931,"Tau;":932,"Upsilon;":933,"Phi;":934,"Chi;":935,"Psi;":936,"Omega;":937,"alpha;":945,"beta;":946,"gamma;":947,"delta;":948,"epsilon;":949,"zeta;":950,"eta;":951,"theta;":952,"iota;":953,"kappa;":954,"lambda;":955,"mu;":956,"nu;":957,"xi;":958,"omicron;":959,"pi;":960,"rho;":961,"sigmaf;":962,"sigma;":963,"tau;":964,"upsilon;":965,"phi;":966,"chi;":967,"psi;":968,"omega;":969,"thetasym;":977,"upsih;":978,"piv;":982,"ensp;":8194,"emsp;":8195,"thinsp;":8201,"zwnj;":8204,"zwj;":8205,"lrm;":8206,"rlm;":8207,"ndash;":8211,"mdash;":8212,"lsquo;":8216,"rsquo;":8217,"sbquo;":8218,"ldquo;":8220,"rdquo;":8221,"bdquo;":8222,"dagger;":8224,"Dagger;":8225,"bull;":8226,"hellip;":8230,"permil;":8240,"prime;":8242,"Prime;":8243,"lsaquo;":8249,"rsaquo;":8250,"oline;":8254,"frasl;":8260,"euro;":8364,"image;":8465,"weierp;":8472,"real;":8476,"trade;":8482,"alefsym;":8501,"larr;":8592,"uarr;":8593,"rarr;":8594,"darr;":8595,"harr;":8596,"crarr;":8629,"lArr;":8656,"uArr;":8657,"rArr;":8658,"dArr;":8659,"hArr;":8660,"forall;":8704,"part;":8706,"exist;":8707,"empty;":8709,"nabla;":8711,"isin;":8712,"notin;":8713,"ni;":8715,"prod;":8719,"sum;":8721,"minus;":8722,"lowast;":8727,"radic;":8730,"prop;":8733,"infin;":8734,"ang;":8736,"and;":8743,"or;":8744,"cap;":8745,"cup;":8746,"int;":8747,"there4;":8756,"sim;":8764,"cong;":8773,"asymp;":8776,"ne;":8800,"equiv;":8801,"le;":8804,"ge;":8805,"sub;":8834,"sup;":8835,"nsub;":8836,"sube;":8838,"supe;":8839,"oplus;":8853,"otimes;":8855,"perp;":8869,"sdot;":8901,"lceil;":8968,"rceil;":8969,"lfloor;":8970,"rfloor;":8971,"lang;":9001,"rang;":9002,"loz;":9674,"spades;":9824,"clubs;":9827,"hearts;":9829,"diams;":9830}}.call(this)},{"./_count":1}]},{},[2])(2)});
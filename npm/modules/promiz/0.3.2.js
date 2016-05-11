!function(){function a(a,b){var c=this;c.promise=c,c.state="pending",c.val=null,c.fn=a||null,c.er=b||null,c.next=[]}var b="undefined"!=typeof setImmediate?setImmediate:function(a){setTimeout(a,0)};a.prototype.resolve=function(a){var c=this;"pending"===c.state&&(c.val=a,c.state="resolving",b(function(){c.fire()}))},a.prototype.reject=function(a){var c=this;"pending"===c.state&&(c.val=a,c.state="rejecting",b(function(){c.fire()}))},a.prototype.then=function(b,c){var d=this,e=new a(b,c);return d.next.push(e),"resolved"===d.state&&e.resolve(d.val),"rejected"===d.state&&e.reject(d.val),e},a.prototype.fail=function(a){return this.then(null,a)},a.prototype.finish=function(a){var b=this;b.state=a,"resolved"===b.state&&b.next.map(function(a){a.resolve(b.val)}),"rejected"===b.state&&b.next.map(function(a){a.reject(b.val)})},a.prototype.thennable=function(a,b,c,d,e){var f=this;if(e=e||f.val,"object"==typeof e&&"function"==typeof a)try{var g=0;a.call(e,function(a){0===g++&&b(a)},function(a){0===g++&&c(a)})}catch(h){c(h)}else d(e)},a.prototype.fire=function(){var a,b=this;try{a=b.val&&b.val.then}catch(c){return b.val=c,b.state="rejecting",b.fire()}b.thennable(a,function(a){b.val=a,b.state="resolving",b.fire()},function(a){b.val=a,b.state="rejecting",b.fire()},function(c){if(b.val=c,"resolving"===b.state&&"function"==typeof b.fn)try{b.val=b.fn.call(void 0,b.val)}catch(d){return b.val=d,b.finish("rejected")}if("rejecting"===b.state&&"function"==typeof b.er)try{b.val=b.er.call(void 0,b.val),b.state="resolving"}catch(d){return b.val=d,b.finish("rejected")}return b.val===b?(b.val=TypeError(),b.finish("rejected")):void b.thennable(a,function(a){b.val=a,b.finish("resolved")},function(a){b.val=a,b.finish("rejected")},function(a){b.val=a,b.finish("resolving"===b.state?"resolved":"rejected")})})},a.prototype.done=function(){if("rejected"===this.state)throw this.val;return null},a.prototype.nodeify=function(a){return"function"==typeof a?this.then(function(b){try{a(null,b)}catch(c){setImmediate(function(){throw c})}return b},function(b){try{a(b)}catch(c){setImmediate(function(){throw c})}return b}):this},a.prototype.spread=function(a,b){return this.all().then(function(b){return"function"==typeof a&&a.apply(null,b)},b)},a.prototype.all=function(){var b=this;return this.then(function(c){function d(){++f===g&&e.resolve(c)}var e=new a;if(!(c instanceof Array))return e.reject(TypeError),e;for(var f=0,g=c.length,h=0,i=c.length;i>h;h++){var j,k=c[h];try{j=k&&k.then}catch(l){e.reject(l);break}!function(a){b.thennable(j,function(b){c[a]=b,d()},function(b){c[a]=b,d()},function(){d()},k)}(h)}return e})};var c={defer:function(){return new a(null,null)},fcall:function(){var b=new a,c=Array.apply([],arguments),d=c.shift();try{var e=d.apply(null,c);b.resolve(e)}catch(f){b.reject(f)}return b},nfcall:function(){var b=new a,c=Array.apply([],arguments),d=c.shift();try{c.push(function(a,c){return a?b.reject(a):b.resolve(c)}),d.apply(null,c)}catch(e){b.reject(e)}return b}};"undefined"!=typeof module?module.exports=c:self.Promiz=c}();
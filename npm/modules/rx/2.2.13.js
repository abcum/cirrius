(function(a){function b(){}function c(a){return a}function d(a,b){return R(a,b)}function e(a,b){return a-b}function f(a){return a.toString()}function g(a){throw a}function h(){if(this.isDisposed)throw new Error(D)}function i(a){return"function"!=typeof a.toString&&"string"==typeof(a+"")}function j(a){return a&&"object"==typeof a?N.call(a)==E:!1}function k(a){return"function"==typeof a}function l(a,b,c,d){var e;if(a===b)return 0!==a||1/a==1/b;var f=typeof a,g=typeof b;if(!(a!==a||a&&s[f]||b&&s[g]))return!1;if(null==a||null==b)return a===b;var h=N.call(a),m=N.call(b);if(h==E&&(h=K),m==E&&(m=K),h!=m)return!1;switch(h){case G:case H:return+a==+b;case J:return a!=+a?b!=+b:0==a?1/a==1/b:a==+b;case L:case M:return a==String(b)}var n=h==F;if(!n){if(h!=K||!y&&(i(a)||i(b)))return!1;var o=!P&&j(a)?Object:a.constructor,p=!P&&j(b)?Object:b.constructor;if(o!=p&&!(k(o)&&o instanceof o&&k(p)&&p instanceof p))return!1}for(var q=c.length;q--;)if(c[q]==a)return d[q]==b;var r=0;if(e=!0,c.push(a),d.push(b),n){for(q=a.length,r=b.length,e=r==a.length;r--;){var t=b[r];if(!(e=l(a[r],t,c,d)))break}return e}for(var u in b)if(O.call(b,u))return r++,e=O.call(a,u)&&l(a[u],b[u],c,d);if(e)for(var u in a)if(O.call(a,u))return e=--r>-1;return c.pop(),d.pop(),e}function m(a,b){return 1===a.length&&Array.isArray(a[b])?a[b]:S.call(a)}function n(a,b){for(var c=new Array(a),d=0;a>d;d++)c[d]=b();return c}function o(a,b){this.scheduler=a,this.disposable=b,this.isDisposed=!1}function p(a,b){return new Qb(function(c){var d=new db,e=new eb;return e.setDisposable(d),d.setDisposable(a.subscribe(c.onNext.bind(c),function(a){var d,f;try{f=b(a)}catch(g){return c.onError(g),void 0}d=new db,e.setDisposable(d),d.setDisposable(f.subscribe(c))},c.onCompleted.bind(c))),e})}function q(a,b){var c=this;return new Qb(function(d){var e=0,f=a.length;return c.subscribe(function(c){if(f>e){var g,h=a[e++];try{g=b(c,h)}catch(i){return d.onError(i),void 0}d.onNext(g)}else d.onCompleted()},d.onError.bind(d),d.onCompleted.bind(d))})}function r(a){return this.select(a).mergeObservable()}var s={"boolean":!1,"function":!0,object:!0,number:!1,string:!1,undefined:!1},t=s[typeof window]&&window||this,u=s[typeof exports]&&exports&&!exports.nodeType&&exports,v=s[typeof module]&&module&&!module.nodeType&&module,w=v&&v.exports===u&&u,x=s[typeof global]&&global;!x||x.global!==x&&x.window!==x||(t=x);var y,z={Internals:{}},A=Date.now,B="Sequence contains no elements.",C="Argument out of range",D="Object has been disposed",E="[object Arguments]",F="[object Array]",G="[object Boolean]",H="[object Date]",I="[object Function]",J="[object Number]",K="[object Object]",L="[object RegExp]",M="[object String]",N=Object.prototype.toString,O=Object.prototype.hasOwnProperty,P=N.call(arguments)==E;try{y=!(N.call(document)==K&&!({toString:0}+""))}catch(Q){y=!0}P||(j=function(a){return a&&"object"==typeof a?O.call(a,"callee"):!1}),k(/x/)&&(k=function(a){return"function"==typeof a&&N.call(a)==I});var R=z.Internals.isEqual=function(a,b){return l(a,b,[],[])},S=Array.prototype.slice,T=({}.hasOwnProperty,this.inherits=z.Internals.inherits=function(a,b){function c(){this.constructor=a}c.prototype=b.prototype,a.prototype=new c}),U=z.Internals.addProperties=function(a){for(var b=S.call(arguments,1),c=0,d=b.length;d>c;c++){var e=b[c];for(var f in e)a[f]=e[f]}},V=z.Internals.addRef=function(a,b){return new Qb(function(c){return new Z(b.getDisposable(),a.subscribe(c))})},W=function(a,b){this.id=a,this.value=b};W.prototype.compareTo=function(a){var b=this.value.compareTo(a.value);return 0===b&&(b=this.id-a.id),b};var X=z.Internals.PriorityQueue=function(a){this.items=new Array(a),this.length=0},Y=X.prototype;Y.isHigherPriority=function(a,b){return this.items[a].compareTo(this.items[b])<0},Y.percolate=function(a){if(!(a>=this.length||0>a)){var b=a-1>>1;if(!(0>b||b===a)&&this.isHigherPriority(a,b)){var c=this.items[a];this.items[a]=this.items[b],this.items[b]=c,this.percolate(b)}}},Y.heapify=function(b){if(b===a&&(b=0),!(b>=this.length||0>b)){var c=2*b+1,d=2*b+2,e=b;if(c<this.length&&this.isHigherPriority(c,e)&&(e=c),d<this.length&&this.isHigherPriority(d,e)&&(e=d),e!==b){var f=this.items[b];this.items[b]=this.items[e],this.items[e]=f,this.heapify(e)}}},Y.peek=function(){return this.items[0].value},Y.removeAt=function(a){this.items[a]=this.items[--this.length],delete this.items[this.length],this.heapify()},Y.dequeue=function(){var a=this.peek();return this.removeAt(0),a},Y.enqueue=function(a){var b=this.length++;this.items[b]=new W(X.count++,a),this.percolate(b)},Y.remove=function(a){for(var b=0;b<this.length;b++)if(this.items[b].value===a)return this.removeAt(b),!0;return!1},X.count=0;var Z=z.CompositeDisposable=function(){this.disposables=m(arguments,0),this.isDisposed=!1,this.length=this.disposables.length},$=Z.prototype;$.add=function(a){this.isDisposed?a.dispose():(this.disposables.push(a),this.length++)},$.remove=function(a){var b=!1;if(!this.isDisposed){var c=this.disposables.indexOf(a);-1!==c&&(b=!0,this.disposables.splice(c,1),this.length--,a.dispose())}return b},$.dispose=function(){if(!this.isDisposed){this.isDisposed=!0;var a=this.disposables.slice(0);this.disposables=[],this.length=0;for(var b=0,c=a.length;c>b;b++)a[b].dispose()}},$.clear=function(){var a=this.disposables.slice(0);this.disposables=[],this.length=0;for(var b=0,c=a.length;c>b;b++)a[b].dispose()},$.contains=function(a){return-1!==this.disposables.indexOf(a)},$.toArray=function(){return this.disposables.slice(0)};var _=z.Disposable=function(a){this.isDisposed=!1,this.action=a||b};_.prototype.dispose=function(){this.isDisposed||(this.action(),this.isDisposed=!0)};var ab=_.create=function(a){return new _(a)},bb=_.empty={dispose:b},cb=function(){function a(a){this.isSingle=a,this.isDisposed=!1,this.current=null}var b=a.prototype;return b.getDisposable=function(){return this.current},b.setDisposable=function(a){if(this.current&&this.isSingle)throw new Error("Disposable has already been assigned");var b,c=this.isDisposed;c||(b=this.current,this.current=a),b&&b.dispose(),c&&a&&a.dispose()},b.dispose=function(){var a;this.isDisposed||(this.isDisposed=!0,a=this.current,this.current=null),a&&a.dispose()},a}(),db=z.SingleAssignmentDisposable=function(a){function b(){a.call(this,!0)}return T(b,a),b}(cb),eb=z.SerialDisposable=function(a){function b(){a.call(this,!1)}return T(b,a),b}(cb),fb=z.RefCountDisposable=function(){function a(a){this.disposable=a,this.disposable.count++,this.isInnerDisposed=!1}function b(a){this.underlyingDisposable=a,this.isDisposed=!1,this.isPrimaryDisposed=!1,this.count=0}return a.prototype.dispose=function(){this.disposable.isDisposed||this.isInnerDisposed||(this.isInnerDisposed=!0,this.disposable.count--,0===this.disposable.count&&this.disposable.isPrimaryDisposed&&(this.disposable.isDisposed=!0,this.disposable.underlyingDisposable.dispose()))},b.prototype.dispose=function(){this.isDisposed||this.isPrimaryDisposed||(this.isPrimaryDisposed=!0,0===this.count&&(this.isDisposed=!0,this.underlyingDisposable.dispose()))},b.prototype.getDisposable=function(){return this.isDisposed?bb:new a(this)},b}();o.prototype.dispose=function(){var a=this;this.scheduler.schedule(function(){a.isDisposed||(a.isDisposed=!0,a.disposable.dispose())})};var gb=z.Internals.ScheduledItem=function(a,b,c,d,f){this.scheduler=a,this.state=b,this.action=c,this.dueTime=d,this.comparer=f||e,this.disposable=new db};gb.prototype.invoke=function(){this.disposable.setDisposable(this.invokeCore())},gb.prototype.compareTo=function(a){return this.comparer(this.dueTime,a.dueTime)},gb.prototype.isCancelled=function(){return this.disposable.isDisposed},gb.prototype.invokeCore=function(){return this.action(this.scheduler,this.state)};var hb,ib=z.Scheduler=function(){function a(a,b,c,d){this.now=a,this._schedule=b,this._scheduleRelative=c,this._scheduleAbsolute=d}function b(a,b){var c=b.first,d=b.second,e=new Z,f=function(b){d(b,function(b){var c=!1,d=!1,g=a.scheduleWithState(b,function(a,b){return c?e.remove(g):d=!0,f(b),bb});d||(e.add(g),c=!0)})};return f(c),e}function c(a,b,c){var d=b.first,e=b.second,f=new Z,g=function(b){e(b,function(b,d){var e=!1,h=!1,i=a[c].call(a,b,d,function(a,b){return e?f.remove(i):h=!0,g(b),bb});h||(f.add(i),e=!0)})};return g(d),f}function d(a,b){return b(),bb}var e=a.prototype;return e.catchException=e["catch"]=function(a){return new nb(this,a)},e.schedulePeriodic=function(a,b){return this.schedulePeriodicWithState(null,a,function(){b()})},e.schedulePeriodicWithState=function(a,b,c){var d=a,e=setInterval(function(){d=c(d)},b);return ab(function(){clearInterval(e)})},e.schedule=function(a){return this._schedule(a,d)},e.scheduleWithState=function(a,b){return this._schedule(a,b)},e.scheduleWithRelative=function(a,b){return this._scheduleRelative(b,a,d)},e.scheduleWithRelativeAndState=function(a,b,c){return this._scheduleRelative(a,b,c)},e.scheduleWithAbsolute=function(a,b){return this._scheduleAbsolute(b,a,d)},e.scheduleWithAbsoluteAndState=function(a,b,c){return this._scheduleAbsolute(a,b,c)},e.scheduleRecursive=function(a){return this.scheduleRecursiveWithState(a,function(a,b){a(function(){b(a)})})},e.scheduleRecursiveWithState=function(a,c){return this.scheduleWithState({first:a,second:c},function(a,c){return b(a,c)})},e.scheduleRecursiveWithRelative=function(a,b){return this.scheduleRecursiveWithRelativeAndState(b,a,function(a,b){a(function(c){b(a,c)})})},e.scheduleRecursiveWithRelativeAndState=function(a,b,d){return this._scheduleRelative({first:a,second:d},b,function(a,b){return c(a,b,"scheduleWithRelativeAndState")})},e.scheduleRecursiveWithAbsolute=function(a,b){return this.scheduleRecursiveWithAbsoluteAndState(b,a,function(a,b){a(function(c){b(a,c)})})},e.scheduleRecursiveWithAbsoluteAndState=function(a,b,d){return this._scheduleAbsolute({first:a,second:d},b,function(a,b){return c(a,b,"scheduleWithAbsoluteAndState")})},a.now=A,a.normalize=function(a){return 0>a&&(a=0),a},a}(),jb=ib.normalize,kb=(z.Internals.SchedulePeriodicRecursive=function(){function a(a,b){b(0,this._period);try{this._state=this._action(this._state)}catch(c){throw this._cancel.dispose(),c}}function b(a,b,c,d){this._scheduler=a,this._state=b,this._period=c,this._action=d}return b.prototype.start=function(){var b=new db;return this._cancel=b,b.setDisposable(this._scheduler.scheduleRecursiveWithRelativeAndState(0,this._period,a.bind(this))),b},b}(),ib.immediate=function(){function a(a,b){return b(this,a)}function b(a,b,c){for(var d=jb(d);d-this.now()>0;);return c(this,a)}function c(a,b,c){return this.scheduleWithRelativeAndState(a,b-this.now(),c)}return new ib(A,a,b,c)}()),lb=ib.currentThread=function(){function a(a){for(var b;a.length>0;)if(b=a.dequeue(),!b.isCancelled()){for(;b.dueTime-ib.now()>0;);b.isCancelled()||b.invoke()}}function b(a,b){return this.scheduleWithRelativeAndState(a,0,b)}function c(b,c,d){var f=this.now()+ib.normalize(c),g=new gb(this,b,d,f);if(e)e.enqueue(g);else{e=new X(4),e.enqueue(g);try{a(e)}catch(h){throw h}finally{e=null}}return g.disposable}function d(a,b,c){return this.scheduleWithRelativeAndState(a,b-this.now(),c)}var e,f=new ib(A,b,c,d);return f.scheduleRequired=function(){return null===e},f.ensureTrampoline=function(a){return null===e?this.schedule(a):a()},f}(),mb=b;!function(){function a(){if(!t.postMessage||t.importScripts)return!1;var a=!1,b=t.onmessage;return t.onmessage=function(){a=!0},t.postMessage("","*"),t.onmessage=b,a}function b(a){for(var b=0,c=a.length;c>b;b++){var d=a[b].target.getAttribute("drainQueue");h[d](),delete h[d]}}function c(a){if("string"==typeof a.data&&a.data.substring(0,l.length)===l){var b=a.data.substring(l.length),c=m[b];c(),delete m[b]}}var d=RegExp("^"+String(N).replace(/[.*+?^${}()|[\]\\]/g,"\\$&").replace(/toString| for [^\]]+/g,".*?")+"$"),e="function"==typeof(e=x&&w&&x.setImmediate)&&!d.test(e)&&e,f="function"==typeof(f=x&&w&&x.clearImmediate)&&!d.test(f)&&f,g=t.MutationObserver||t.WebKitMutationObserver;if(g){var h={},i=0,j=new g(b),k=document.createElement("div");j.observe(k,{attributes:!0}),t.addEventListener("unload",function(){j.disconnect(),j=null}),hb=function(a){var b=i++;return h[b]=a,k.setAttribute("drainQueue",b),b},mb=function(a){delete h[a]}}else if("undefined"!=typeof process&&"[object process]"==={}.toString.call(process))hb=process.nextTick;else if("function"==typeof e)hb=e,mb=f;else if(a()){var l="ms.rx.schedule"+Math.random(),m={},n=0;t.addEventListener?t.addEventListener("message",c,!1):t.attachEvent("onmessage",c,!1),hb=function(a){var b=n++;m[b]=a,t.postMessage(l+b,"*")}}else if(t.MessageChannel){var o=new t.MessageChannel,p={},q=0;o.port1.onmessage=function(a){var b=a.data,c=p[b];c(),delete p[b]},hb=function(a){var b=q++;p[b]=a,o.port2.postMessage(b)}}else"document"in t&&"onreadystatechange"in t.document.createElement("script")?hb=function(a){var b=t.document.createElement("script");b.onreadystatechange=function(){a(),b.onreadystatechange=null,b.parentNode.removeChild(b),b=null},t.document.documentElement.appendChild(b)}:(hb=function(a){return setTimeout(a,0)},mb=clearTimeout)}();var nb=(ib.timeout=function(){function a(a,b){var c=this,d=new db,e=hb(function(){d.isDisposed||d.setDisposable(b(c,a))});return new Z(d,ab(function(){mb(e)}))}function b(a,b,c){var d=this,e=ib.normalize(b);if(0===e)return d.scheduleWithState(a,c);var f=new db,g=setTimeout(function(){f.isDisposed||f.setDisposable(c(d,a))},e);return new Z(f,ab(function(){clearTimeout(g)}))}function c(a,b,c){return this.scheduleWithRelativeAndState(a,b-this.now(),c)}return new ib(A,a,b,c)}(),function(a){function b(){return this._scheduler.now()}function c(a,b){return this._scheduler.scheduleWithState(a,this._wrap(b))}function d(a,b,c){return this._scheduler.scheduleWithRelativeAndState(a,b,this._wrap(c))}function e(a,b,c){return this._scheduler.scheduleWithAbsoluteAndState(a,b,this._wrap(c))}function f(f,g){this._scheduler=f,this._handler=g,this._recursiveOriginal=null,this._recursiveWrapper=null,a.call(this,b,c,d,e)}return T(f,a),f.prototype._clone=function(a){return new f(a,this._handler)},f.prototype._wrap=function(a){var b=this;return function(c,d){try{return a(b._getRecursiveWrapper(c),d)}catch(e){if(!b._handler(e))throw e;return bb}}},f.prototype._getRecursiveWrapper=function(a){if(this._recursiveOriginal!==a){this._recursiveOriginal=a;var b=this._clone(a);b._recursiveOriginal=a,b._recursiveWrapper=b,this._recursiveWrapper=b}return this._recursiveWrapper},f.prototype.schedulePeriodicWithState=function(a,b,c){var d=this,e=!1,f=new db;return f.setDisposable(this._scheduler.schedulePeriodicWithState(a,b,function(a){if(e)return null;try{return c(a)}catch(b){if(e=!0,!d._handler(b))throw b;return f.dispose(),null}})),f},f}(ib)),ob=z.Notification=function(){function a(a,b){this.hasValue=null==b?!1:b,this.kind=a}var b=a.prototype;return b.accept=function(a,b,c){return 1===arguments.length&&"object"==typeof a?this._acceptObservable(a):this._accept(a,b,c)},b.toObservable=function(a){var b=this;return a||(a=kb),new Qb(function(c){return a.schedule(function(){b._acceptObservable(c),"N"===b.kind&&c.onCompleted()})})},a}(),pb=ob.createOnNext=function(){function a(a){return a(this.value)}function b(a){return a.onNext(this.value)}function c(){return"OnNext("+this.value+")"}return function(d){var e=new ob("N",!0);return e.value=d,e._accept=a,e._acceptObservable=b,e.toString=c,e}}(),qb=ob.createOnError=function(){function a(a,b){return b(this.exception)}function b(a){return a.onError(this.exception)}function c(){return"OnError("+this.exception+")"}return function(d){var e=new ob("E");return e.exception=d,e._accept=a,e._acceptObservable=b,e.toString=c,e}}(),rb=ob.createOnCompleted=function(){function a(a,b,c){return c()}function b(a){return a.onCompleted()}function c(){return"OnCompleted()"}return function(){var d=new ob("C");return d._accept=a,d._acceptObservable=b,d.toString=c,d}}(),sb=z.Internals.Enumerator=function(a,b){this.moveNext=a,this.getCurrent=b},tb=sb.create=function(a,b){var c=!1;return new sb(function(){if(c)return!1;var b=a();return b||(c=!0),b},function(){return b()})},ub=z.Internals.Enumerable=function(a){this.getEnumerator=a};ub.prototype.concat=function(){var a=this;return new Qb(function(b){var c,d=a.getEnumerator(),e=new eb,f=kb.scheduleRecursive(function(a){var f,g;if(!c){try{g=d.moveNext(),g&&(f=d.getCurrent())}catch(h){return b.onError(h),void 0}if(!g)return b.onCompleted(),void 0;var i=new db;e.setDisposable(i),i.setDisposable(f.subscribe(b.onNext.bind(b),b.onError.bind(b),function(){a()}))}});return new Z(e,f,ab(function(){c=!0}))})},ub.prototype.catchException=function(){var a=this;return new Qb(function(b){var c,d,e=a.getEnumerator(),f=new eb,g=kb.scheduleRecursive(function(a){var g,h;if(!c){try{h=e.moveNext(),h&&(g=e.getCurrent())}catch(i){return b.onError(i),void 0}if(!h)return d?b.onError(d):b.onCompleted(),void 0;var j=new db;f.setDisposable(j),j.setDisposable(g.subscribe(b.onNext.bind(b),function(b){d=b,a()},b.onCompleted.bind(b)))}});return new Z(f,g,ab(function(){c=!0}))})};var vb=ub.repeat=function(a,b){return 1===arguments.length&&(b=-1),new ub(function(){var c,d=b;return tb(function(){return 0===d?!1:(d>0&&d--,c=a,!0)},function(){return c})})},wb=ub.forEach=function(a,b,d){return b||(b=c),new ub(function(){var c,e=-1;return tb(function(){return++e<a.length?(c=b.call(d,a[e],e,a),!0):!1},function(){return c})})},xb=z.Observer=function(){};xb.prototype.toNotifier=function(){var a=this;return function(b){return b.accept(a)}},xb.prototype.asObserver=function(){return new Bb(this.onNext.bind(this),this.onError.bind(this),this.onCompleted.bind(this))},xb.prototype.checked=function(){return new Cb(this)};var yb=xb.create=function(a,c,d){return a||(a=b),c||(c=g),d||(d=b),new Bb(a,c,d)};xb.fromNotifier=function(a){return new Bb(function(b){return a(pb(b))},function(b){return a(qb(b))},function(){return a(rb())})},xb.notifyOn=function(a){return new Eb(a,this)};var zb,Ab=z.Internals.AbstractObserver=function(a){function b(){this.isStopped=!1,a.call(this)}return T(b,a),b.prototype.onNext=function(a){this.isStopped||this.next(a)},b.prototype.onError=function(a){this.isStopped||(this.isStopped=!0,this.error(a))},b.prototype.onCompleted=function(){this.isStopped||(this.isStopped=!0,this.completed())},b.prototype.dispose=function(){this.isStopped=!0},b.prototype.fail=function(a){return this.isStopped?!1:(this.isStopped=!0,this.error(a),!0)},b}(xb),Bb=z.AnonymousObserver=function(a){function b(b,c,d){a.call(this),this._onNext=b,this._onError=c,this._onCompleted=d}return T(b,a),b.prototype.next=function(a){this._onNext(a)},b.prototype.error=function(a){this._onError(a)},b.prototype.completed=function(){this._onCompleted()},b}(Ab),Cb=function(a){function b(b){a.call(this),this._observer=b,this._state=0}T(b,a);var c=b.prototype;return c.onNext=function(a){this.checkAccess();try{this._observer.onNext(a)}catch(b){throw b}finally{this._state=0}},c.onError=function(a){this.checkAccess();try{this._observer.onError(a)}catch(b){throw b}finally{this._state=2}},c.onCompleted=function(){this.checkAccess();try{this._observer.onCompleted()}catch(a){throw a}finally{this._state=2}},c.checkAccess=function(){if(1===this._state)throw new Error("Re-entrancy detected");if(2===this._state)throw new Error("Observer completed");0===this._state&&(this._state=1)},b}(xb),Db=z.Internals.ScheduledObserver=function(a){function b(b,c){a.call(this),this.scheduler=b,this.observer=c,this.isAcquired=!1,this.hasFaulted=!1,this.queue=[],this.disposable=new eb}return T(b,a),b.prototype.next=function(a){var b=this;this.queue.push(function(){b.observer.onNext(a)})},b.prototype.error=function(a){var b=this;this.queue.push(function(){b.observer.onError(a)})},b.prototype.completed=function(){var a=this;this.queue.push(function(){a.observer.onCompleted()})},b.prototype.ensureActive=function(){var a=!1,b=this;!this.hasFaulted&&this.queue.length>0&&(a=!this.isAcquired,this.isAcquired=!0),a&&this.disposable.setDisposable(this.scheduler.scheduleRecursive(function(a){var c;if(!(b.queue.length>0))return b.isAcquired=!1,void 0;c=b.queue.shift();try{c()}catch(d){throw b.queue=[],b.hasFaulted=!0,d}a()}))},b.prototype.dispose=function(){a.prototype.dispose.call(this),this.disposable.dispose()},b}(Ab),Eb=function(a){function b(){a.apply(this,arguments)}return T(b,a),b.prototype.next=function(b){a.prototype.next.call(this,b),this.ensureActive()},b.prototype.error=function(b){a.prototype.error.call(this,b),this.ensureActive()},b.prototype.completed=function(){a.prototype.completed.call(this),this.ensureActive()},b}(Db),Fb=z.Observable=function(){function a(a){this._subscribe=a}return zb=a.prototype,zb.finalValue=function(){var a=this;return new Qb(function(b){var c,d=!1;return a.subscribe(function(a){d=!0,c=a},b.onError.bind(b),function(){d?(b.onNext(c),b.onCompleted()):b.onError(new Error(B))})})},zb.subscribe=zb.forEach=function(a,b,c){var d;return d="object"==typeof a?a:yb(a,b,c),this._subscribe(d)},zb.toArray=function(){function a(a,b){var c=a.slice(0);return c.push(b),c}return this.scan([],a).startWith([]).finalValue()},a}();zb.observeOn=function(a){var b=this;return new Qb(function(c){return b.subscribe(new Eb(a,c))})},zb.subscribeOn=function(a){var b=this;return new Qb(function(c){var d=new db,e=new eb;return e.setDisposable(d),d.setDisposable(a.schedule(function(){e.setDisposable(new o(a,b.subscribe(c)))})),e})},Fb.create=Fb.createWithDisposable=function(a){return new Qb(a)};var Gb=(Fb.defer=function(a){return new Qb(function(b){var c;try{c=a()}catch(d){return Kb(d).subscribe(b)}return c.subscribe(b)})},Fb.empty=function(a){return a||(a=kb),new Qb(function(b){return a.schedule(function(){b.onCompleted()})})}),Hb=Fb.fromArray=function(a,b){return b||(b=lb),new Qb(function(c){var d=0;return b.scheduleRecursive(function(b){d<a.length?(c.onNext(a[d++]),b()):c.onCompleted()})})};Fb.generate=function(a,b,c,d,e){return e||(e=lb),new Qb(function(f){var g=!0,h=a;return e.scheduleRecursive(function(a){var e,i;try{g?g=!1:h=c(h),e=b(h),e&&(i=d(h))}catch(j){return f.onError(j),void 0}e?(f.onNext(i),a()):f.onCompleted()})})};var Ib=Fb.never=function(){return new Qb(function(){return bb})};Fb.range=function(a,b,c){return c||(c=lb),new Qb(function(d){return c.scheduleRecursiveWithState(0,function(c,e){b>c?(d.onNext(a+c),e(c+1)):d.onCompleted()})})},Fb.repeat=function(a,b,c){return c||(c=lb),null==b&&(b=-1),Jb(a,c).repeat(b)};var Jb=Fb["return"]=Fb.returnValue=function(a,b){return b||(b=kb),new Qb(function(c){return b.schedule(function(){c.onNext(a),c.onCompleted()})})},Kb=Fb["throw"]=Fb.throwException=function(a,b){return b||(b=kb),new Qb(function(c){return b.schedule(function(){c.onError(a)})})};Fb.using=function(a,b){return new Qb(function(c){var d,e,f=bb;try{d=a(),d&&(f=d),e=b(d)}catch(g){return new Z(Kb(g).subscribe(c),f)}return new Z(e.subscribe(c),f)})},zb.amb=function(a){var b=this;return new Qb(function(c){function d(){f||(f=g,j.dispose())}function e(){f||(f=h,i.dispose())}var f,g="L",h="R",i=new db,j=new db;return i.setDisposable(b.subscribe(function(a){d(),f===g&&c.onNext(a)},function(a){d(),f===g&&c.onError(a)},function(){d(),f===g&&c.onCompleted()})),j.setDisposable(a.subscribe(function(a){e(),f===h&&c.onNext(a)},function(a){e(),f===h&&c.onError(a)},function(){e(),f===h&&c.onCompleted()})),new Z(i,j)})},Fb.amb=function(){function a(a,b){return a.amb(b)}for(var b=Ib(),c=m(arguments,0),d=0,e=c.length;e>d;d++)b=a(b,c[d]);return b},zb["catch"]=zb.catchException=function(a){return"function"==typeof a?p(this,a):Lb([this,a])};var Lb=Fb.catchException=Fb["catch"]=function(){var a=m(arguments,0);return wb(a).catchException()};zb.combineLatest=function(){var a=S.call(arguments);return Array.isArray(a[0])?a[0].unshift(this):a.unshift(this),Mb.apply(this,a)};var Mb=Fb.combineLatest=function(){var a=S.call(arguments),b=a.pop();return Array.isArray(a[0])&&(a=a[0]),new Qb(function(d){function e(a){var e;if(i[a]=!0,j||(j=i.every(c))){try{e=b.apply(null,l)}catch(f){return d.onError(f),void 0}d.onNext(e)}else k.filter(function(b,c){return c!==a}).every(c)&&d.onCompleted()}function f(a){k[a]=!0,k.every(c)&&d.onCompleted()}for(var g=function(){return!1},h=a.length,i=n(h,g),j=!1,k=n(h,g),l=new Array(h),m=new Array(h),o=0;h>o;o++)!function(b){m[b]=new db,m[b].setDisposable(a[b].subscribe(function(a){l[b]=a,e(b)},d.onError.bind(d),function(){f(b)}))}(o);return new Z(m)})};zb.concat=function(){var a=S.call(arguments,0);return a.unshift(this),Nb.apply(this,a)};var Nb=Fb.concat=function(){var a=m(arguments,0);return wb(a).concat()};zb.concatObservable=zb.concatAll=function(){return this.merge(1)},zb.merge=function(a){if("number"!=typeof a)return Ob(this,a);var b=this;return new Qb(function(c){var d=0,e=new Z,f=!1,g=[],h=function(a){var b=new db;e.add(b),b.setDisposable(a.subscribe(c.onNext.bind(c),c.onError.bind(c),function(){var a;e.remove(b),g.length>0?(a=g.shift(),h(a)):(d--,f&&0===d&&c.onCompleted())}))};return e.add(b.subscribe(function(b){a>d?(d++,h(b)):g.push(b)},c.onError.bind(c),function(){f=!0,0===d&&c.onCompleted()})),e})};var Ob=Fb.merge=function(){var a,b;return arguments[0]?arguments[0].now?(a=arguments[0],b=S.call(arguments,1)):(a=kb,b=S.call(arguments,0)):(a=kb,b=S.call(arguments,1)),Array.isArray(b[0])&&(b=b[0]),Hb(b,a).mergeObservable()};zb.mergeObservable=zb.mergeAll=function(){var a=this;return new Qb(function(b){var c=new Z,d=!1,e=new db;return c.add(e),e.setDisposable(a.subscribe(function(a){var e=new db;c.add(e),e.setDisposable(a.subscribe(function(a){b.onNext(a)},b.onError.bind(b),function(){c.remove(e),d&&1===c.length&&b.onCompleted()}))},b.onError.bind(b),function(){d=!0,1===c.length&&b.onCompleted()})),c})},zb.onErrorResumeNext=function(a){if(!a)throw new Error("Second observable is required");return Pb([this,a])};var Pb=Fb.onErrorResumeNext=function(){var a=m(arguments,0);return new Qb(function(b){var c=0,d=new eb,e=kb.scheduleRecursive(function(e){var f,g;c<a.length?(f=a[c++],g=new db,d.setDisposable(g),g.setDisposable(f.subscribe(b.onNext.bind(b),function(){e()},function(){e()}))):b.onCompleted()});return new Z(d,e)})};zb.skipUntil=function(a){var b=this;return new Qb(function(c){var d=!1,e=new Z(b.subscribe(function(a){d&&c.onNext(a)},c.onError.bind(c),function(){d&&c.onCompleted()})),f=new db;return e.add(f),f.setDisposable(a.subscribe(function(){d=!0,f.dispose()},c.onError.bind(c),function(){f.dispose()})),e})},zb["switch"]=zb.switchLatest=function(){var a=this;return new Qb(function(b){var c=!1,d=new eb,e=!1,f=0,g=a.subscribe(function(a){var g=new db,h=++f;c=!0,d.setDisposable(g),g.setDisposable(a.subscribe(function(a){f===h&&b.onNext(a)},function(a){f===h&&b.onError(a)},function(){f===h&&(c=!1,e&&b.onCompleted())}))},b.onError.bind(b),function(){e=!0,c||b.onCompleted()});return new Z(g,d)})},zb.takeUntil=function(a){var c=this;return new Qb(function(d){return new Z(c.subscribe(d),a.subscribe(d.onCompleted.bind(d),d.onError.bind(d),b))})},zb.zip=function(){if(Array.isArray(arguments[0]))return q.apply(this,arguments);var a=this,b=S.call(arguments),d=b.pop();return b.unshift(a),new Qb(function(e){function f(a){i[a]=!0,i.every(function(a){return a})&&e.onCompleted()}for(var g=b.length,h=n(g,function(){return[]}),i=n(g,function(){return!1}),j=function(b){var f,g;if(h.every(function(a){return a.length>0})){try{g=h.map(function(a){return a.shift()}),f=d.apply(a,g)}catch(j){return e.onError(j),void 0}e.onNext(f)}else i.filter(function(a,c){return c!==b}).every(c)&&e.onCompleted()},k=new Array(g),l=0;g>l;l++)!function(a){k[a]=new db,k[a].setDisposable(b[a].subscribe(function(b){h[a].push(b),j(a)},e.onError.bind(e),function(){f(a)}))}(l);return new Z(k)})},Fb.zip=function(){var a=S.call(arguments,0),b=a.shift();return b.zip.apply(b,a)},Fb.zipArray=function(){var a=m(arguments,0);return new Qb(function(b){function d(a){if(g.every(function(a){return a.length>0})){var d=g.map(function(a){return a.shift()});b.onNext(d)}else if(h.filter(function(b,c){return c!==a}).every(c))return b.onCompleted(),void 0}function e(a){return h[a]=!0,h.every(c)?(b.onCompleted(),void 0):void 0}for(var f=a.length,g=n(f,function(){return[]}),h=n(f,function(){return!1}),i=new Array(f),j=0;f>j;j++)!function(c){i[c]=new db,i[c].setDisposable(a[c].subscribe(function(a){g[c].push(a),d(c)},b.onError.bind(b),function(){e(c)}))}(j);var k=new Z(i);return k.add(ab(function(){for(var a=0,b=g.length;b>a;a++)g[a]=[]})),k})},zb.asObservable=function(){var a=this;return new Qb(function(b){return a.subscribe(b)})},zb.bufferWithCount=function(a,b){return 1===arguments.length&&(b=a),this.windowWithCount(a,b).selectMany(function(a){return a.toArray()}).where(function(a){return a.length>0})},zb.dematerialize=function(){var a=this;return new Qb(function(b){return a.subscribe(function(a){return a.accept(b)},b.onError.bind(b),b.onCompleted.bind(b))})},zb.distinctUntilChanged=function(a,b){var e=this;return a||(a=c),b||(b=d),new Qb(function(c){var d,f=!1;return e.subscribe(function(e){var g,h=!1;try{g=a(e)}catch(i){return c.onError(i),void 0}if(f)try{h=b(d,g)}catch(i){return c.onError(i),void 0}f&&h||(f=!0,d=g,c.onNext(e))},c.onError.bind(c),c.onCompleted.bind(c))})},zb["do"]=zb.doAction=function(a,b,c){var d,e=this;return"function"==typeof a?d=a:(d=a.onNext.bind(a),b=a.onError.bind(a),c=a.onCompleted.bind(a)),new Qb(function(a){return e.subscribe(function(b){try{d(b)}catch(c){a.onError(c)}a.onNext(b)},function(c){if(b){try{b(c)}catch(d){a.onError(d)}a.onError(c)}else a.onError(c)},function(){if(c){try{c()}catch(b){a.onError(b)}a.onCompleted()}else a.onCompleted()})})},zb["finally"]=zb.finallyAction=function(a){var b=this;return new Qb(function(c){var d=b.subscribe(c);return ab(function(){try{d.dispose()}catch(b){throw b}finally{a()}})})},zb.ignoreElements=function(){var a=this;return new Qb(function(c){return a.subscribe(b,c.onError.bind(c),c.onCompleted.bind(c))})},zb.materialize=function(){var a=this;return new Qb(function(b){return a.subscribe(function(a){b.onNext(pb(a))},function(a){b.onNext(qb(a)),b.onCompleted()},function(){b.onNext(rb()),b.onCompleted()})})},zb.repeat=function(a){return vb(this,a).concat()},zb.retry=function(a){return vb(this,a).catchException()},zb.scan=function(){var a,b,c=!1,d=this;return 2===arguments.length?(c=!0,a=arguments[0],b=arguments[1]):b=arguments[0],new Qb(function(e){var f,g,h;return d.subscribe(function(d){try{h||(h=!0),f?g=b(g,d):(g=c?b(a,d):d,f=!0)}catch(i){return e.onError(i),void 0}e.onNext(g)},e.onError.bind(e),function(){!h&&c&&e.onNext(a),e.onCompleted()})})},zb.skipLast=function(a){var b=this;return new Qb(function(c){var d=[];return b.subscribe(function(b){d.push(b),d.length>a&&c.onNext(d.shift())},c.onError.bind(c),c.onCompleted.bind(c))})},zb.startWith=function(){var a,b,c=0;return arguments.length&&"now"in Object(arguments[0])?(b=arguments[0],c=1):b=kb,a=S.call(arguments,c),wb([Hb(a,b),this]).concat()},zb.takeLast=function(a,b){return this.takeLastBuffer(a).selectMany(function(a){return Hb(a,b)})},zb.takeLastBuffer=function(a){var b=this;return new Qb(function(c){var d=[];return b.subscribe(function(b){d.push(b),d.length>a&&d.shift()},c.onError.bind(c),function(){c.onNext(d),c.onCompleted()})})},zb.windowWithCount=function(a,b){var c=this;if(0>=a)throw new Error(C);if(1===arguments.length&&(b=a),0>=b)throw new Error(C);return new Qb(function(d){var e=new db,f=new fb(e),g=0,h=[],i=function(){var a=new Ub;h.push(a),d.onNext(V(a,f))};return i(),e.setDisposable(c.subscribe(function(c){for(var d,e=0,f=h.length;f>e;e++)h[e].onNext(c);var j=g-a+1;j>=0&&j%b===0&&(d=h.shift(),d.onCompleted()),g++,g%b===0&&i()},function(a){for(;h.length>0;)h.shift().onError(a);d.onError(a)},function(){for(;h.length>0;)h.shift().onCompleted();d.onCompleted()})),f})},zb.defaultIfEmpty=function(b){var c=this;return b===a&&(b=null),new Qb(function(a){var d=!1;return c.subscribe(function(b){d=!0,a.onNext(b)},a.onError.bind(a),function(){d||a.onNext(b),a.onCompleted()})})},zb.distinct=function(a,b){var d=this;return a||(a=c),b||(b=f),new Qb(function(c){var e={};return d.subscribe(function(d){var f,g,h,i=!1;
try{f=a(d),g=b(f)}catch(j){return c.onError(j),void 0}for(h in e)if(g===h){i=!0;break}i||(e[g]=null,c.onNext(d))},c.onError.bind(c),c.onCompleted.bind(c))})},zb.groupBy=function(a,b,c){return this.groupByUntil(a,b,function(){return Ib()},c)},zb.groupByUntil=function(a,d,e,g){var h=this;return d||(d=c),g||(g=f),new Qb(function(c){var f={},i=new Z,j=new fb(i);return i.add(h.subscribe(function(h){var k,l,m,n,o,p,q,r,s,t;try{p=a(h),q=g(p)}catch(u){for(t in f)f[t].onError(u);return c.onError(u),void 0}n=!1;try{s=f[q],s||(s=new Ub,f[q]=s,n=!0)}catch(u){for(t in f)f[t].onError(u);return c.onError(u),void 0}if(n){o=new Sb(p,s,j),l=new Sb(p,s);try{k=e(l)}catch(u){for(t in f)f[t].onError(u);return c.onError(u),void 0}c.onNext(o),r=new db,i.add(r);var v=function(){q in f&&(delete f[q],s.onCompleted()),i.remove(r)};r.setDisposable(k.take(1).subscribe(b,function(a){for(t in f)f[t].onError(a);c.onError(a)},function(){v()}))}try{m=d(h)}catch(u){for(t in f)f[t].onError(u);return c.onError(u),void 0}s.onNext(m)},function(a){for(var b in f)f[b].onError(a);c.onError(a)},function(){for(var a in f)f[a].onCompleted();c.onCompleted()})),j})},zb.select=zb.map=function(a,b){var c=this;return new Qb(function(d){var e=0;return c.subscribe(function(f){var g;try{g=a.call(b,f,e++,c)}catch(h){return d.onError(h),void 0}d.onNext(g)},d.onError.bind(d),d.onCompleted.bind(d))})},zb.pluck=function(a){return this.select(function(b){return b[a]})},zb.selectMany=zb.flatMap=function(a,b){return b?this.selectMany(function(c){return a(c).select(function(a){return b(c,a)})}):"function"==typeof a?r.call(this,a):r.call(this,function(){return a})},zb.selectSwitch=zb.flatMapLatest=function(a,b){return this.select(a,b).switchLatest()},zb.skip=function(a){if(0>a)throw new Error(C);var b=this;return new Qb(function(c){var d=a;return b.subscribe(function(a){0>=d?c.onNext(a):d--},c.onError.bind(c),c.onCompleted.bind(c))})},zb.skipWhile=function(a,b){var c=this;return new Qb(function(d){var e=0,f=!1;return c.subscribe(function(g){if(!f)try{f=!a.call(b,g,e++,c)}catch(h){return d.onError(h),void 0}f&&d.onNext(g)},d.onError.bind(d),d.onCompleted.bind(d))})},zb.take=function(a,b){if(0>a)throw new Error(C);if(0===a)return Gb(b);var c=this;return new Qb(function(b){var d=a;return c.subscribe(function(a){d>0&&(d--,b.onNext(a),0===d&&b.onCompleted())},b.onError.bind(b),b.onCompleted.bind(b))})},zb.takeWhile=function(a,b){var c=this;return new Qb(function(d){var e=0,f=!0;return c.subscribe(function(g){if(f){try{f=a.call(b,g,e++,c)}catch(h){return d.onError(h),void 0}f?d.onNext(g):d.onCompleted()}},d.onError.bind(d),d.onCompleted.bind(d))})},zb.where=zb.filter=function(a,b){var c=this;return new Qb(function(d){var e=0;return c.subscribe(function(f){var g;try{g=a.call(b,f,e++,c)}catch(h){return d.onError(h),void 0}g&&d.onNext(f)},d.onError.bind(d),d.onCompleted.bind(d))})};var Qb=z.Internals.AnonymousObservable=function(a){function b(a){return"undefined"==typeof a?a=bb:"function"==typeof a&&(a=ab(a)),a}function c(d){function e(a){var c=new Rb(a);if(lb.scheduleRequired())lb.schedule(function(){try{c.setDisposable(b(d(c)))}catch(a){if(!c.fail(a))throw a}});else try{c.setDisposable(b(d(c)))}catch(e){if(!c.fail(e))throw e}return c}return this instanceof c?(a.call(this,e),void 0):new c(d)}return T(c,a),c}(Fb),Rb=function(a){function b(b){a.call(this),this.observer=b,this.m=new db}T(b,a);var c=b.prototype;return c.next=function(a){var b=!1;try{this.observer.onNext(a),b=!0}catch(c){throw c}finally{b||this.dispose()}},c.error=function(a){try{this.observer.onError(a)}catch(b){throw b}finally{this.dispose()}},c.completed=function(){try{this.observer.onCompleted()}catch(a){throw a}finally{this.dispose()}},c.setDisposable=function(a){this.m.setDisposable(a)},c.getDisposable=function(){return this.m.getDisposable()},c.disposable=function(a){return arguments.length?this.getDisposable():setDisposable(a)},c.dispose=function(){a.prototype.dispose.call(this),this.m.dispose()},b}(Ab),Sb=function(a){function b(a){return this.underlyingObservable.subscribe(a)}function c(c,d,e){a.call(this,b),this.key=c,this.underlyingObservable=e?new Qb(function(a){return new Z(e.getDisposable(),d.subscribe(a))}):d}return T(c,a),c}(Fb),Tb=function(a,b){this.subject=a,this.observer=b};Tb.prototype.dispose=function(){if(!this.subject.isDisposed&&null!==this.observer){var a=this.subject.observers.indexOf(this.observer);this.subject.observers.splice(a,1),this.observer=null}};var Ub=z.Subject=function(a){function b(a){return h.call(this),this.isStopped?this.exception?(a.onError(this.exception),bb):(a.onCompleted(),bb):(this.observers.push(a),new Tb(this,a))}function c(){a.call(this,b),this.isDisposed=!1,this.isStopped=!1,this.observers=[]}return T(c,a),U(c.prototype,xb,{hasObservers:function(){return this.observers.length>0},onCompleted:function(){if(h.call(this),!this.isStopped){var a=this.observers.slice(0);this.isStopped=!0;for(var b=0,c=a.length;c>b;b++)a[b].onCompleted();this.observers=[]}},onError:function(a){if(h.call(this),!this.isStopped){var b=this.observers.slice(0);this.isStopped=!0,this.exception=a;for(var c=0,d=b.length;d>c;c++)b[c].onError(a);this.observers=[]}},onNext:function(a){if(h.call(this),!this.isStopped)for(var b=this.observers.slice(0),c=0,d=b.length;d>c;c++)b[c].onNext(a)},dispose:function(){this.isDisposed=!0,this.observers=null}}),c.create=function(a,b){return new Vb(a,b)},c}(Fb),Vb=(z.AsyncSubject=function(a){function b(a){if(h.call(this),!this.isStopped)return this.observers.push(a),new Tb(this,a);var b=this.exception,c=this.hasValue,d=this.value;return b?a.onError(b):c?(a.onNext(d),a.onCompleted()):a.onCompleted(),bb}function c(){a.call(this,b),this.isDisposed=!1,this.isStopped=!1,this.value=null,this.hasValue=!1,this.observers=[],this.exception=null}return T(c,a),U(c.prototype,xb,{hasObservers:function(){return h.call(this),this.observers.length>0},onCompleted:function(){var a,b,c;if(h.call(this),!this.isStopped){this.isStopped=!0;var d=this.observers.slice(0),e=this.value,f=this.hasValue;if(f)for(b=0,c=d.length;c>b;b++)a=d[b],a.onNext(e),a.onCompleted();else for(b=0,c=d.length;c>b;b++)d[b].onCompleted();this.observers=[]}},onError:function(a){if(h.call(this),!this.isStopped){var b=this.observers.slice(0);this.isStopped=!0,this.exception=a;for(var c=0,d=b.length;d>c;c++)b[c].onError(a);this.observers=[]}},onNext:function(a){h.call(this),this.isStopped||(this.value=a,this.hasValue=!0)},dispose:function(){this.isDisposed=!0,this.observers=null,this.exception=null,this.value=null}}),c}(Fb),function(a){function b(a){return this.observable.subscribe(a)}function c(c,d){a.call(this,b),this.observer=c,this.observable=d}return T(c,a),U(c.prototype,xb,{onCompleted:function(){this.observer.onCompleted()},onError:function(a){this.observer.onError(a)},onNext:function(a){this.observer.onNext(a)}}),c}(Fb));"function"==typeof define&&"object"==typeof define.amd&&define.amd?(t.Rx=z,define(function(){return z})):u&&v?w?(v.exports=z).Rx=z:u.Rx=z:t.Rx=z}).call(this);
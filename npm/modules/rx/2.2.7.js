(function(t,e){function n(){}function r(t){return t}function i(t,e){return V(t,e)}function o(t,e){return t-e}function s(t){return""+t}function u(t){throw t}function c(){if(this.isDisposed)throw Error(_)}function a(t){return"function"!=typeof t.toString&&"string"==typeof(t+"")}function h(t){return t&&"object"==typeof t?T.call(t)==N:!1}function l(t){return"function"==typeof t}function f(t,e,n,r){var i;if(t===e)return 0!==t||1/t==1/e;var o=typeof t,s=typeof e;if(!(t!==t||t&&S[o]||e&&S[s]))return!1;if(null==t||null==e)return t===e;var u=T.call(t),c=T.call(e);if(u==N&&(u=I),c==N&&(c=I),u!=c)return!1;switch(u){case R:case W:return+t==+e;case k:return t!=+t?e!=+e:0==t?1/t==1/e:t==+e;case q:case P:return t==e+""}var p=u==O;if(!p){if(u!=I||!E&&(a(t)||a(e)))return!1;var d=!L&&h(t)?Object:t.constructor,b=!L&&h(e)?Object:e.constructor;if(d!=b&&!(l(d)&&d instanceof d&&l(b)&&b instanceof b))return!1}for(var v=n.length;v--;)if(n[v]==t)return r[v]==e;var m=0;if(i=!0,n.push(t),r.push(e),p){for(v=t.length,m=e.length,i=m==t.length;m--;){var y=e[m];if(!(i=f(t[m],y,n,r)))break}return i}for(var w in e)if(M.call(e,w))return m++,i=M.call(t,w)&&f(t[w],e[w],n,r);if(i)for(var w in t)if(M.call(t,w))return i=--m>-1;return n.pop(),r.pop(),i}function p(t,e){return 1===t.length&&Array.isArray(t[e])?t[e]:F.call(t)}function d(t,e){for(var n=Array(t),r=0;t>r;r++)n[r]=e();return n}function b(t,e){this.scheduler=t,this.disposable=e,this.isDisposed=!1}function v(t,n){return new Ve(function(r){var i=new te,o=new ne;return o.setDisposable(i),i.setDisposable(t.subscribe(r.onNext.bind(r),function(t){var i,s;try{s=n(t)}catch(u){return r.onError(u),e}i=new te,o.setDisposable(i),i.setDisposable(s.subscribe(r))},r.onCompleted.bind(r))),o})}function m(t,n){var r=this;return new Ve(function(i){var o=0,s=t.length;return r.subscribe(function(r){if(s>o){var u,c=t[o++];try{u=n(r,c)}catch(a){return i.onError(a),e}i.onNext(u)}else i.onCompleted()},i.onError.bind(i),i.onCompleted.bind(i))})}function y(t){return this.select(t).mergeObservable()}var w="object"==typeof exports&&exports,g=("object"==typeof module&&module&&module.exports==w&&module,"object"==typeof global&&global);g.global===g&&(t=g);var E,D={Internals:{}},x=Date.now,C="Sequence contains no elements.",A="Argument out of range",_="Object has been disposed",S={"boolean":!1,"function":!0,object:!0,number:!1,string:!1,undefined:!1},N="[object Arguments]",O="[object Array]",R="[object Boolean]",W="[object Date]",j="[object Function]",k="[object Number]",I="[object Object]",q="[object RegExp]",P="[object String]",T=Object.prototype.toString,M=Object.prototype.hasOwnProperty,L=T.call(arguments)==N;try{E=!(T.call(document)==I&&!({toString:0}+""))}catch(z){E=!0}L||(h=function(t){return t&&"object"==typeof t?M.call(t,"callee"):!1}),l(/x/)&&(l=function(t){return"function"==typeof t&&T.call(t)==j});var V=D.Internals.isEqual=function(t,e){return f(t,e,[],[])},F=Array.prototype.slice;({}).hasOwnProperty;var B=this.inherits=D.Internals.inherits=function(t,e){function n(){this.constructor=t}n.prototype=e.prototype,t.prototype=new n},U=D.Internals.addProperties=function(t){for(var e=F.call(arguments,1),n=0,r=e.length;r>n;n++){var i=e[n];for(var o in i)t[o]=i[o]}},H=D.Internals.addRef=function(t,e){return new Ve(function(n){return new K(e.getDisposable(),t.subscribe(n))})},Q=function(t,e){this.id=t,this.value=e};Q.prototype.compareTo=function(t){var e=this.value.compareTo(t.value);return 0===e&&(e=this.id-t.id),e};var G=D.Internals.PriorityQueue=function(t){this.items=Array(t),this.length=0},J=G.prototype;J.isHigherPriority=function(t,e){return 0>this.items[t].compareTo(this.items[e])},J.percolate=function(t){if(!(t>=this.length||0>t)){var e=t-1>>1;if(!(0>e||e===t)&&this.isHigherPriority(t,e)){var n=this.items[t];this.items[t]=this.items[e],this.items[e]=n,this.percolate(e)}}},J.heapify=function(t){if(t===e&&(t=0),!(t>=this.length||0>t)){var n=2*t+1,r=2*t+2,i=t;if(this.length>n&&this.isHigherPriority(n,i)&&(i=n),this.length>r&&this.isHigherPriority(r,i)&&(i=r),i!==t){var o=this.items[t];this.items[t]=this.items[i],this.items[i]=o,this.heapify(i)}}},J.peek=function(){return this.items[0].value},J.removeAt=function(t){this.items[t]=this.items[--this.length],delete this.items[this.length],this.heapify()},J.dequeue=function(){var t=this.peek();return this.removeAt(0),t},J.enqueue=function(t){var e=this.length++;this.items[e]=new Q(G.count++,t),this.percolate(e)},J.remove=function(t){for(var e=0;this.length>e;e++)if(this.items[e].value===t)return this.removeAt(e),!0;return!1},G.count=0;var K=D.CompositeDisposable=function(){this.disposables=p(arguments,0),this.isDisposed=!1,this.length=this.disposables.length},X=K.prototype;X.add=function(t){this.isDisposed?t.dispose():(this.disposables.push(t),this.length++)},X.remove=function(t){var e=!1;if(!this.isDisposed){var n=this.disposables.indexOf(t);-1!==n&&(e=!0,this.disposables.splice(n,1),this.length--,t.dispose())}return e},X.dispose=function(){if(!this.isDisposed){this.isDisposed=!0;var t=this.disposables.slice(0);this.disposables=[],this.length=0;for(var e=0,n=t.length;n>e;e++)t[e].dispose()}},X.clear=function(){var t=this.disposables.slice(0);this.disposables=[],this.length=0;for(var e=0,n=t.length;n>e;e++)t[e].dispose()},X.contains=function(t){return-1!==this.disposables.indexOf(t)},X.toArray=function(){return this.disposables.slice(0)};var Y=D.Disposable=function(t){this.isDisposed=!1,this.action=t||n};Y.prototype.dispose=function(){this.isDisposed||(this.action(),this.isDisposed=!0)};var Z=Y.create=function(t){return new Y(t)},$=Y.empty={dispose:n},te=D.SingleAssignmentDisposable=function(){this.isDisposed=!1,this.current=null},ee=te.prototype;ee.getDisposable=function(){return this.current},te.disposable=function(t){return arguments.length?this.getDisposable():this.setDisposable(t)},ee.setDisposable=function(t){if(this.current)throw Error("Disposable has already been assigned");var e=this.isDisposed;e||(this.current=t),e&&t&&t.dispose()},ee.dispose=function(){var t;this.isDisposed||(this.isDisposed=!0,t=this.current,this.current=null),t&&t.dispose()};var ne=D.SerialDisposable=function(){this.isDisposed=!1,this.current=null},re=ne.prototype;re.getDisposable=function(){return this.current},re.setDisposable=function(t){var e,n=this.isDisposed;n||(e=this.current,this.current=t),e&&e.dispose(),n&&t&&t.dispose()},re.disposable=function(t){return t?(this.setDisposable(t),e):this.getDisposable()},re.dispose=function(){var t;this.isDisposed||(this.isDisposed=!0,t=this.current,this.current=null),t&&t.dispose()};var ie=D.RefCountDisposable=function(){function t(t){this.disposable=t,this.disposable.count++,this.isInnerDisposed=!1}function e(t){this.underlyingDisposable=t,this.isDisposed=!1,this.isPrimaryDisposed=!1,this.count=0}return t.prototype.dispose=function(){this.disposable.isDisposed||this.isInnerDisposed||(this.isInnerDisposed=!0,this.disposable.count--,0===this.disposable.count&&this.disposable.isPrimaryDisposed&&(this.disposable.isDisposed=!0,this.disposable.underlyingDisposable.dispose()))},e.prototype.dispose=function(){this.isDisposed||this.isPrimaryDisposed||(this.isPrimaryDisposed=!0,0===this.count&&(this.isDisposed=!0,this.underlyingDisposable.dispose()))},e.prototype.getDisposable=function(){return this.isDisposed?$:new t(this)},e}();b.prototype.dispose=function(){var t=this;this.scheduler.schedule(function(){t.isDisposed||(t.isDisposed=!0,t.disposable.dispose())})};var oe=D.Internals.ScheduledItem=function(t,e,n,r,i){this.scheduler=t,this.state=e,this.action=n,this.dueTime=r,this.comparer=i||o,this.disposable=new te};oe.prototype.invoke=function(){this.disposable.setDisposable(this.invokeCore())},oe.prototype.compareTo=function(t){return this.comparer(this.dueTime,t.dueTime)},oe.prototype.isCancelled=function(){return this.disposable.isDisposed},oe.prototype.invokeCore=function(){return this.action(this.scheduler,this.state)};var se=D.Scheduler=function(){function e(t,e,n,r){this.now=t,this._schedule=e,this._scheduleRelative=n,this._scheduleAbsolute=r}function n(t,e){var n=e.first,r=e.second,i=new K,o=function(e){r(e,function(e){var n=!1,r=!1,s=t.scheduleWithState(e,function(t,e){return n?i.remove(s):r=!0,o(e),$});r||(i.add(s),n=!0)})};return o(n),i}function r(t,e,n){var r=e.first,i=e.second,o=new K,s=function(e){i(e,function(e,r){var i=!1,u=!1,c=t[n].call(t,e,r,function(t,e){return i?o.remove(c):u=!0,s(e),$});u||(o.add(c),i=!0)})};return s(r),o}function i(t,e){return e(),$}var o=e.prototype;return o.catchException=o["catch"]=function(t){return new fe(this,t)},o.schedulePeriodic=function(t,e){return this.schedulePeriodicWithState(null,t,function(){e()})},o.schedulePeriodicWithState=function(e,n,r){var i=e,o=t.setInterval(function(){i=r(i)},n);return Z(function(){t.clearInterval(o)})},o.schedule=function(t){return this._schedule(t,i)},o.scheduleWithState=function(t,e){return this._schedule(t,e)},o.scheduleWithRelative=function(t,e){return this._scheduleRelative(e,t,i)},o.scheduleWithRelativeAndState=function(t,e,n){return this._scheduleRelative(t,e,n)},o.scheduleWithAbsolute=function(t,e){return this._scheduleAbsolute(e,t,i)},o.scheduleWithAbsoluteAndState=function(t,e,n){return this._scheduleAbsolute(t,e,n)},o.scheduleRecursive=function(t){return this.scheduleRecursiveWithState(t,function(t,e){t(function(){e(t)})})},o.scheduleRecursiveWithState=function(t,e){return this.scheduleWithState({first:t,second:e},function(t,e){return n(t,e)})},o.scheduleRecursiveWithRelative=function(t,e){return this.scheduleRecursiveWithRelativeAndState(e,t,function(t,e){t(function(n){e(t,n)})})},o.scheduleRecursiveWithRelativeAndState=function(t,e,n){return this._scheduleRelative({first:t,second:n},e,function(t,e){return r(t,e,"scheduleWithRelativeAndState")})},o.scheduleRecursiveWithAbsolute=function(t,e){return this.scheduleRecursiveWithAbsoluteAndState(e,t,function(t,e){t(function(n){e(t,n)})})},o.scheduleRecursiveWithAbsoluteAndState=function(t,e,n){return this._scheduleAbsolute({first:t,second:n},e,function(t,e){return r(t,e,"scheduleWithAbsoluteAndState")})},e.now=x,e.normalize=function(t){return 0>t&&(t=0),t},e}();D.Internals.SchedulePeriodicRecursive=function(){function t(t,e){e(0,this._period);try{this._state=this._action(this._state)}catch(n){throw this._cancel.dispose(),n}}function e(t,e,n,r){this._scheduler=t,this._state=e,this._period=n,this._action=r}return e.prototype.start=function(){var e=new te;return this._cancel=e,e.setDisposable(this._scheduler.scheduleRecursiveWithRelativeAndState(0,this._period,t.bind(this))),e},e}();var ue,ce="Scheduler is not allowed to block the thread",ae=se.immediate=function(){function t(t,e){return e(this,t)}function e(t,e,n){if(e>0)throw Error(ce);return n(this,t)}function n(t,e,n){return this.scheduleWithRelativeAndState(t,e-this.now(),n)}return new se(x,t,e,n)}(),he=se.currentThread=function(){function t(){i=new G(4)}function e(t,e){return this.scheduleWithRelativeAndState(t,0,e)}function n(e,n,r){var o,s=this.now()+se.normalize(n),u=new oe(this,e,r,s);if(i)i.enqueue(u);else{o=new t;try{i.enqueue(u),o.run()}catch(c){throw c}finally{o.dispose()}}return u.disposable}function r(t,e,n){return this.scheduleWithRelativeAndState(t,e-this.now(),n)}var i;t.prototype.dispose=function(){i=null},t.prototype.run=function(){for(var t;i.length>0;)if(t=i.dequeue(),!t.isCancelled()){for(;t.dueTime-se.now()>0;);t.isCancelled()||t.invoke()}};var o=new se(x,e,n,r);return o.scheduleRequired=function(){return null===i},o.ensureTrampoline=function(t){return null===i?this.schedule(t):t()},o}(),le=n;(function(){function e(){if(!t.postMessage||t.importScripts)return!1;var e=!1,n=t.onmessage;return t.onmessage=function(){e=!0},t.postMessage("","*"),t.onmessage=n,e}function n(t){if("string"==typeof t.data&&t.data.substring(0,r.length)===r){var e=t.data.substring(r.length),n=i[e];n(),delete i[e]}}if("function"==typeof t.setImmediate)ue=t.setImmediate,le=clearImmediate;else if("undefined"!=typeof process&&"[object process]"==={}.toString.call(process))ue=process.nextTick;else if(e()){var r="ms.rx.schedule"+Math.random(),i={},o=0;t.addEventListener?t.addEventListener("message",n,!1):t.attachEvent("onmessage",n,!1),ue=function(e){var n=o++;i[n]=e,t.postMessage(r+n,"*")}}else if(t.MessageChannel){var s=new t.MessageChannel,u={},c=0;s.port1.onmessage=function(t){var e=t.data,n=u[e];n(),delete u[e]},ue=function(t){var e=c++;u[e]=t,s.port2.postMessage(e)}}else"document"in t&&"onreadystatechange"in t.document.createElement("script")?ue=function(e){var n=t.document.createElement("script");n.onreadystatechange=function(){e(),n.onreadystatechange=null,n.parentNode.removeChild(n),n=null},t.document.documentElement.appendChild(n)}:(ue=function(e){return t.setTimeout(e,0)},le=t.clearTimeout)})(),se.timeout=function(){function e(t,e){var n=this,r=new te,i=ue(function(){r.isDisposed||r.setDisposable(e(n,t))});return new K(r,Z(function(){le(i)}))}function n(e,n,r){var i=this,o=se.normalize(n);if(0===o)return i.scheduleWithState(e,r);var s=new te,u=t.setTimeout(function(){s.isDisposed||s.setDisposable(r(i,e))},o);return new K(s,Z(function(){t.clearTimeout(u)}))}function r(t,e,n){return this.scheduleWithRelativeAndState(t,e-this.now(),n)}return new se(x,e,n,r)}();var fe=function(t){function e(){return this._scheduler.now()}function n(t,e){return this._scheduler.scheduleWithState(t,this._wrap(e))}function r(t,e,n){return this._scheduler.scheduleWithRelativeAndState(t,e,this._wrap(n))}function i(t,e,n){return this._scheduler.scheduleWithAbsoluteAndState(t,e,this._wrap(n))}function o(o,s){this._scheduler=o,this._handler=s,this._recursiveOriginal=null,this._recursiveWrapper=null,t.call(this,e,n,r,i)}return B(o,t),o.prototype._clone=function(t){return new o(t,this._handler)},o.prototype._wrap=function(t){var e=this;return function(n,r){try{return t(e._getRecursiveWrapper(n),r)}catch(i){if(!e._handler(i))throw i;return $}}},o.prototype._getRecursiveWrapper=function(t){if(this._recursiveOriginal!==t){this._recursiveOriginal=t;var e=this._clone(t);e._recursiveOriginal=t,e._recursiveWrapper=e,this._recursiveWrapper=e}return this._recursiveWrapper},o.prototype.schedulePeriodicWithState=function(t,e,n){var r=this,i=!1,o=new te;return o.setDisposable(this._scheduler.schedulePeriodicWithState(t,e,function(t){if(i)return null;try{return n(t)}catch(e){if(i=!0,!r._handler(e))throw e;return o.dispose(),null}})),o},o}(se),pe=D.Notification=function(){function t(t,e){this.hasValue=null==e?!1:e,this.kind=t}var e=t.prototype;return e.accept=function(t,e,n){return 1===arguments.length&&"object"==typeof t?this._acceptObservable(t):this._accept(t,e,n)},e.toObservable=function(t){var e=this;return t||(t=ae),new Ve(function(n){return t.schedule(function(){e._acceptObservable(n),"N"===e.kind&&n.onCompleted()})})},t}(),de=pe.createOnNext=function(){function t(t){return t(this.value)}function e(t){return t.onNext(this.value)}function n(){return"OnNext("+this.value+")"}return function(r){var i=new pe("N",!0);return i.value=r,i._accept=t,i._acceptObservable=e,i.toString=n,i}}(),be=pe.createOnError=function(){function t(t,e){return e(this.exception)}function e(t){return t.onError(this.exception)}function n(){return"OnError("+this.exception+")"}return function(r){var i=new pe("E");return i.exception=r,i._accept=t,i._acceptObservable=e,i.toString=n,i}}(),ve=pe.createOnCompleted=function(){function t(t,e,n){return n()}function e(t){return t.onCompleted()}function n(){return"OnCompleted()"}return function(){var r=new pe("C");return r._accept=t,r._acceptObservable=e,r.toString=n,r}}(),me=D.Internals.Enumerator=function(t,e,n){this.moveNext=t,this.getCurrent=e,this.dispose=n},ye=me.create=function(t,e,r){var i=!1;return r||(r=n),new me(function(){if(i)return!1;var e=t();return e||(i=!0,r()),e},function(){return e()},function(){i||(r(),i=!0)})},we=D.Internals.Enumerable=function(){function t(t){this.getEnumerator=t}return t.prototype.concat=function(){var t=this;return new Ve(function(n){var r=t.getEnumerator(),i=!1,o=new ne,s=ae.scheduleRecursive(function(t){var s,u,c=!1;if(!i){try{c=r.moveNext(),c?s=r.getCurrent():r.dispose()}catch(a){u=a,r.dispose()}if(u)return n.onError(u),e;if(!c)return n.onCompleted(),e;var h=new te;o.setDisposable(h),h.setDisposable(s.subscribe(n.onNext.bind(n),n.onError.bind(n),function(){t()}))}});return new K(o,s,Z(function(){i=!0,r.dispose()}))})},t.prototype.catchException=function(){var t=this;return new Ve(function(n){var r,i=t.getEnumerator(),o=!1,s=new ne,u=ae.scheduleRecursive(function(t){var u,c,a;if(a=!1,!o){try{a=i.moveNext(),a&&(u=i.getCurrent())}catch(h){c=h}if(c)return n.onError(c),e;if(!a)return r?n.onError(r):n.onCompleted(),e;var l=new te;s.setDisposable(l),l.setDisposable(u.subscribe(n.onNext.bind(n),function(e){r=e,t()},n.onCompleted.bind(n)))}});return new K(s,u,Z(function(){o=!0}))})},t}(),ge=we.repeat=function(t,n){return n===e&&(n=-1),new we(function(){var e,r=n;return ye(function(){return 0===r?!1:(r>0&&r--,e=t,!0)},function(){return e})})},Ee=we.forEach=function(t,e){return e||(e=r),new we(function(){var n,r=-1;return ye(function(){return++r<t.length?(n=e(t[r],r),!0):!1},function(){return n})})},De=D.Observer=function(){};De.prototype.toNotifier=function(){var t=this;return function(e){return e.accept(t)}},De.prototype.asObserver=function(){return new _e(this.onNext.bind(this),this.onError.bind(this),this.onCompleted.bind(this))},De.prototype.checked=function(){return new Se(this)};var xe=De.create=function(t,e,r){return t||(t=n),e||(e=u),r||(r=n),new _e(t,e,r)};De.fromNotifier=function(t){return new _e(function(e){return t(de(e))},function(e){return t(be(e))},function(){return t(ve())})},De.notifyOn=function(t){return new Oe(t,this)};var Ce,Ae=D.Internals.AbstractObserver=function(t){function e(){this.isStopped=!1,t.call(this)}return B(e,t),e.prototype.onNext=function(t){this.isStopped||this.next(t)},e.prototype.onError=function(t){this.isStopped||(this.isStopped=!0,this.error(t))},e.prototype.onCompleted=function(){this.isStopped||(this.isStopped=!0,this.completed())},e.prototype.dispose=function(){this.isStopped=!0},e.prototype.fail=function(t){return this.isStopped?!1:(this.isStopped=!0,this.error(t),!0)},e}(De),_e=D.AnonymousObserver=function(t){function e(e,n,r){t.call(this),this._onNext=e,this._onError=n,this._onCompleted=r}return B(e,t),e.prototype.next=function(t){this._onNext(t)},e.prototype.error=function(t){this._onError(t)},e.prototype.completed=function(){this._onCompleted()},e}(Ae),Se=function(t){function e(e){t.call(this),this._observer=e,this._state=0}B(e,t);var n=e.prototype;return n.onNext=function(t){this.checkAccess();try{this._observer.onNext(t)}catch(e){throw e}finally{this._state=0}},n.onError=function(t){this.checkAccess();try{this._observer.onError(t)}catch(e){throw e}finally{this._state=2}},n.onCompleted=function(){this.checkAccess();try{this._observer.onCompleted()}catch(t){throw t}finally{this._state=2}},n.checkAccess=function(){if(1===this._state)throw Error("Re-entrancy detected");if(2===this._state)throw Error("Observer completed");0===this._state&&(this._state=1)},e}(De),Ne=D.Internals.ScheduledObserver=function(t){function n(e,n){t.call(this),this.scheduler=e,this.observer=n,this.isAcquired=!1,this.hasFaulted=!1,this.queue=[],this.disposable=new ne}return B(n,t),n.prototype.next=function(t){var e=this;this.queue.push(function(){e.observer.onNext(t)})},n.prototype.error=function(t){var e=this;this.queue.push(function(){e.observer.onError(t)})},n.prototype.completed=function(){var t=this;this.queue.push(function(){t.observer.onCompleted()})},n.prototype.ensureActive=function(){var t=!1,n=this;!this.hasFaulted&&this.queue.length>0&&(t=!this.isAcquired,this.isAcquired=!0),t&&this.disposable.setDisposable(this.scheduler.scheduleRecursive(function(t){var r;if(!(n.queue.length>0))return n.isAcquired=!1,e;r=n.queue.shift();try{r()}catch(i){throw n.queue=[],n.hasFaulted=!0,i}t()}))},n.prototype.dispose=function(){t.prototype.dispose.call(this),this.disposable.dispose()},n}(Ae),Oe=function(t){function e(){t.apply(this,arguments)}return B(e,t),e.prototype.next=function(e){t.prototype.next.call(this,e),this.ensureActive()},e.prototype.error=function(e){t.prototype.error.call(this,e),this.ensureActive()},e.prototype.completed=function(){t.prototype.completed.call(this),this.ensureActive()},e}(Ne),Re=D.Observable=function(){function t(t){this._subscribe=t}return Ce=t.prototype,Ce.finalValue=function(){var t=this;return new Ve(function(e){var n,r=!1;return t.subscribe(function(t){r=!0,n=t},e.onError.bind(e),function(){r?(e.onNext(n),e.onCompleted()):e.onError(Error(C))})})},Ce.subscribe=Ce.forEach=function(t,e,n){var r;return r="object"==typeof t?t:xe(t,e,n),this._subscribe(r)},Ce.toArray=function(){function t(t,e){var n=t.slice(0);return n.push(e),n}return this.scan([],t).startWith([]).finalValue()},t}();Ce.observeOn=function(t){var e=this;return new Ve(function(n){return e.subscribe(new Oe(t,n))})},Ce.subscribeOn=function(t){var e=this;return new Ve(function(n){var r=new te,i=new ne;return i.setDisposable(r),r.setDisposable(t.schedule(function(){i.setDisposable(new b(t,e.subscribe(n)))})),i})},Re.create=Re.createWithDisposable=function(t){return new Ve(t)},Re.defer=function(t){return new Ve(function(e){var n;try{n=t()}catch(r){return qe(r).subscribe(e)}return n.subscribe(e)})};var We=Re.empty=function(t){return t||(t=ae),new Ve(function(e){return t.schedule(function(){e.onCompleted()})})},je=Re.fromArray=function(t,e){return e||(e=he),new Ve(function(n){var r=0;return e.scheduleRecursive(function(e){t.length>r?(n.onNext(t[r++]),e()):n.onCompleted()})})};Re.generate=function(t,n,r,i,o){return o||(o=he),new Ve(function(s){var u=!0,c=t;return o.scheduleRecursive(function(t){var o,a;try{u?u=!1:c=r(c),o=n(c),o&&(a=i(c))}catch(h){return s.onError(h),e}o?(s.onNext(a),t()):s.onCompleted()})})};var ke=Re.never=function(){return new Ve(function(){return $})};Re.range=function(t,e,n){return n||(n=he),new Ve(function(r){return n.scheduleRecursiveWithState(0,function(n,i){e>n?(r.onNext(t+n),i(n+1)):r.onCompleted()})})},Re.repeat=function(t,e,n){return n||(n=he),null==e&&(e=-1),Ie(t,n).repeat(e)};var Ie=Re["return"]=Re.returnValue=function(t,e){return e||(e=ae),new Ve(function(n){return e.schedule(function(){n.onNext(t),n.onCompleted()})})},qe=Re["throw"]=Re.throwException=function(t,e){return e||(e=ae),new Ve(function(n){return e.schedule(function(){n.onError(t)})})};Re.using=function(t,e){return new Ve(function(n){var r,i,o=$;try{r=t(),r&&(o=r),i=e(r)}catch(s){return new K(qe(s).subscribe(n),o)}return new K(i.subscribe(n),o)})},Ce.amb=function(t){var e=this;return new Ve(function(n){function r(){o||(o=s,a.dispose())}function i(){o||(o=u,c.dispose())}var o,s="L",u="R",c=new te,a=new te;return c.setDisposable(e.subscribe(function(t){r(),o===s&&n.onNext(t)},function(t){r(),o===s&&n.onError(t)},function(){r(),o===s&&n.onCompleted()})),a.setDisposable(t.subscribe(function(t){i(),o===u&&n.onNext(t)},function(t){i(),o===u&&n.onError(t)},function(){i(),o===u&&n.onCompleted()})),new K(c,a)})},Re.amb=function(){function t(t,e){return t.amb(e)}for(var e=ke(),n=p(arguments,0),r=0,i=n.length;i>r;r++)e=t(e,n[r]);return e},Ce["catch"]=Ce.catchException=function(t){return"function"==typeof t?v(this,t):Pe([this,t])};var Pe=Re.catchException=Re["catch"]=function(){var t=p(arguments,0);return Ee(t).catchException()};Ce.combineLatest=function(){var t=F.call(arguments);return Array.isArray(t[0])?t[0].unshift(this):t.unshift(this),Te.apply(this,t)};var Te=Re.combineLatest=function(){var t=F.call(arguments),n=t.pop();return Array.isArray(t[0])&&(t=t[0]),new Ve(function(i){function o(t){var o;if(a[t]=!0,h||(h=a.every(r))){try{o=n.apply(null,f)}catch(s){return i.onError(s),e}i.onNext(o)}else l.filter(function(e,n){return n!==t}).every(r)&&i.onCompleted()}function s(t){l[t]=!0,l.every(r)&&i.onCompleted()}for(var u=function(){return!1},c=t.length,a=d(c,u),h=!1,l=d(c,u),f=Array(c),p=Array(c),b=0;c>b;b++)(function(e){p[e]=new te,p[e].setDisposable(t[e].subscribe(function(t){f[e]=t,o(e)},i.onError.bind(i),function(){s(e)}))})(b);return new K(p)})};Ce.concat=function(){var t=F.call(arguments,0);return t.unshift(this),Me.apply(this,t)};var Me=Re.concat=function(){var t=p(arguments,0);return Ee(t).concat()};Ce.concatObservable=Ce.concatAll=function(){return this.merge(1)},Ce.merge=function(t){if("number"!=typeof t)return Le(this,t);var e=this;return new Ve(function(n){var r=0,i=new K,o=!1,s=[],u=function(t){var e=new te;i.add(e),e.setDisposable(t.subscribe(n.onNext.bind(n),n.onError.bind(n),function(){var t;i.remove(e),s.length>0?(t=s.shift(),u(t)):(r--,o&&0===r&&n.onCompleted())}))};return i.add(e.subscribe(function(e){t>r?(r++,u(e)):s.push(e)},n.onError.bind(n),function(){o=!0,0===r&&n.onCompleted()})),i})};var Le=Re.merge=function(){var t,e;return arguments[0]?arguments[0].now?(t=arguments[0],e=F.call(arguments,1)):(t=ae,e=F.call(arguments,0)):(t=ae,e=F.call(arguments,1)),Array.isArray(e[0])&&(e=e[0]),je(e,t).mergeObservable()};Ce.mergeObservable=Ce.mergeAll=function(){var t=this;return new Ve(function(e){var n=new K,r=!1,i=new te;return n.add(i),i.setDisposable(t.subscribe(function(t){var i=new te;n.add(i),i.setDisposable(t.subscribe(function(t){e.onNext(t)},e.onError.bind(e),function(){n.remove(i),r&&1===n.length&&e.onCompleted()}))},e.onError.bind(e),function(){r=!0,1===n.length&&e.onCompleted()})),n})},Ce.onErrorResumeNext=function(t){if(!t)throw Error("Second observable is required");return ze([this,t])};var ze=Re.onErrorResumeNext=function(){var t=p(arguments,0);return new Ve(function(e){var n=0,r=new ne,i=ae.scheduleRecursive(function(i){var o,s;t.length>n?(o=t[n++],s=new te,r.setDisposable(s),s.setDisposable(o.subscribe(e.onNext.bind(e),function(){i()},function(){i()}))):e.onCompleted()});return new K(r,i)})};Ce.skipUntil=function(t){var e=this;return new Ve(function(n){var r=!1,i=new K(e.subscribe(function(t){r&&n.onNext(t)},n.onError.bind(n),function(){r&&n.onCompleted()})),o=new te;return i.add(o),o.setDisposable(t.subscribe(function(){r=!0,o.dispose()},n.onError.bind(n),function(){o.dispose()})),i})},Ce["switch"]=Ce.switchLatest=function(){var t=this;return new Ve(function(e){var n=!1,r=new ne,i=!1,o=0,s=t.subscribe(function(t){var s=new te,u=++o;n=!0,r.setDisposable(s),s.setDisposable(t.subscribe(function(t){o===u&&e.onNext(t)},function(t){o===u&&e.onError(t)},function(){o===u&&(n=!1,i&&e.onCompleted())}))},e.onError.bind(e),function(){i=!0,n||e.onCompleted()});return new K(s,r)})},Ce.takeUntil=function(t){var e=this;return new Ve(function(r){return new K(e.subscribe(r),t.subscribe(r.onCompleted.bind(r),r.onError.bind(r),n))})},Ce.zip=function(){if(Array.isArray(arguments[0]))return m.apply(this,arguments);var t=this,n=F.call(arguments),i=n.pop();return n.unshift(t),new Ve(function(o){function s(t){a[t]=!0,a.every(function(t){return t})&&o.onCompleted()}for(var u=n.length,c=d(u,function(){return[]}),a=d(u,function(){return!1}),h=function(n){var s,u;if(c.every(function(t){return t.length>0})){try{u=c.map(function(t){return t.shift()}),s=i.apply(t,u)}catch(h){return o.onError(h),e}o.onNext(s)}else a.filter(function(t,e){return e!==n}).every(r)&&o.onCompleted()},l=Array(u),f=0;u>f;f++)(function(t){l[t]=new te,l[t].setDisposable(n[t].subscribe(function(e){c[t].push(e),h(t)},o.onError.bind(o),function(){s(t)}))})(f);return new K(l)})},Re.zip=function(){var t=F.call(arguments,0),e=t.shift();return e.zip.apply(e,t)},Re.zipArray=function(){var t=F.call(arguments);return new Ve(function(n){function i(t){if(u.every(function(t){return t.length>0})){var i=u.map(function(t){return t.shift()});n.onNext(i)}else if(c.filter(function(e,n){return n!==t}).every(r))return n.onCompleted(),e}function o(t){return c[t]=!0,c.every(r)?(n.onCompleted(),e):e}for(var s=t.length,u=d(s,function(){return[]}),c=d(s,function(){return!1}),a=Array(s),h=0;s>h;h++)(function(e){a[e]=new te,a[e].setDisposable(t[e].subscribe(function(t){u[e].push(t),i(e)},n.onError.bind(n),function(){o(e)}))})(h);var l=new K(a);return l.add(Z(function(){for(var t=0,e=u.length;e>t;t++)u[t]=[]})),l})},Ce.asObservable=function(){var t=this;return new Ve(function(e){return t.subscribe(e)})},Ce.bufferWithCount=function(t,e){return 1===arguments.length&&(e=t),this.windowWithCount(t,e).selectMany(function(t){return t.toArray()}).where(function(t){return t.length>0})},Ce.dematerialize=function(){var t=this;return new Ve(function(e){return t.subscribe(function(t){return t.accept(e)},e.onError.bind(e),e.onCompleted.bind(e))})},Ce.distinctUntilChanged=function(t,n){var o=this;return t||(t=r),n||(n=i),new Ve(function(r){var i,s=!1;return o.subscribe(function(o){var u,c=!1;try{u=t(o)}catch(a){return r.onError(a),e}if(s)try{c=n(i,u)}catch(a){return r.onError(a),e}s&&c||(s=!0,i=u,r.onNext(o))},r.onError.bind(r),r.onCompleted.bind(r))})},Ce["do"]=Ce.doAction=function(t,e,n){var r,i=this;return"function"==typeof t?r=t:(r=t.onNext.bind(t),e=t.onError.bind(t),n=t.onCompleted.bind(t)),new Ve(function(t){return i.subscribe(function(e){try{r(e)}catch(n){t.onError(n)}t.onNext(e)},function(n){if(e){try{e(n)}catch(r){t.onError(r)}t.onError(n)}else t.onError(n)},function(){if(n){try{n()}catch(e){t.onError(e)}t.onCompleted()}else t.onCompleted()})})},Ce["finally"]=Ce.finallyAction=function(t){var e=this;return new Ve(function(n){var r=e.subscribe(n);return Z(function(){try{r.dispose()}catch(e){throw e}finally{t()}})})},Ce.ignoreElements=function(){var t=this;return new Ve(function(e){return t.subscribe(n,e.onError.bind(e),e.onCompleted.bind(e))})},Ce.materialize=function(){var t=this;return new Ve(function(e){return t.subscribe(function(t){e.onNext(de(t))},function(t){e.onNext(be(t)),e.onCompleted()},function(){e.onNext(ve()),e.onCompleted()})})},Ce.repeat=function(t){return ge(this,t).concat()},Ce.retry=function(t){return ge(this,t).catchException()},Ce.scan=function(){var t,n,r=!1,i=this;return 2===arguments.length?(r=!0,t=arguments[0],n=arguments[1]):n=arguments[0],new Ve(function(o){var s,u,c;return i.subscribe(function(i){try{c||(c=!0),s?u=n(u,i):(u=r?n(t,i):i,s=!0)}catch(a){return o.onError(a),e}o.onNext(u)},o.onError.bind(o),function(){!c&&r&&o.onNext(t),o.onCompleted()})})},Ce.skipLast=function(t){var e=this;return new Ve(function(n){var r=[];return e.subscribe(function(e){r.push(e),r.length>t&&n.onNext(r.shift())},n.onError.bind(n),n.onCompleted.bind(n))})},Ce.startWith=function(){var t,e,n=0;return arguments.length&&"now"in Object(arguments[0])?(e=arguments[0],n=1):e=ae,t=F.call(arguments,n),Ee([je(t,e),this]).concat()},Ce.takeLast=function(t,e){return this.takeLastBuffer(t).selectMany(function(t){return je(t,e)})},Ce.takeLastBuffer=function(t){var e=this;return new Ve(function(n){var r=[];return e.subscribe(function(e){r.push(e),r.length>t&&r.shift()},n.onError.bind(n),function(){n.onNext(r),n.onCompleted()})})},Ce.windowWithCount=function(t,e){var n=this;if(0>=t)throw Error(A);if(1===arguments.length&&(e=t),0>=e)throw Error(A);return new Ve(function(r){var i=new te,o=new ie(i),s=0,u=[],c=function(){var t=new He;u.push(t),r.onNext(H(t,o))};return c(),i.setDisposable(n.subscribe(function(n){for(var r,i=0,o=u.length;o>i;i++)u[i].onNext(n);var a=s-t+1;a>=0&&0===a%e&&(r=u.shift(),r.onCompleted()),s++,0===s%e&&c()},function(t){for(;u.length>0;)u.shift().onError(t);r.onError(t)},function(){for(;u.length>0;)u.shift().onCompleted();r.onCompleted()})),o})},Ce.defaultIfEmpty=function(t){var n=this;return t===e&&(t=null),new Ve(function(e){var r=!1;return n.subscribe(function(t){r=!0,e.onNext(t)},e.onError.bind(e),function(){r||e.onNext(t),e.onCompleted()})})},Ce.distinct=function(t,n){var i=this;return t||(t=r),n||(n=s),new Ve(function(r){var o={};return i.subscribe(function(i){var s,u,c,a=!1;try{s=t(i),u=n(s)}catch(h){return r.onError(h),e}for(c in o)if(u===c){a=!0;break}a||(o[u]=null,r.onNext(i))},r.onError.bind(r),r.onCompleted.bind(r))})},Ce.groupBy=function(t,e,n){return this.groupByUntil(t,e,function(){return ke()},n)},Ce.groupByUntil=function(t,i,o,u){var c=this;
return i||(i=r),u||(u=s),new Ve(function(r){var s={},a=new K,h=new ie(a);return a.add(c.subscribe(function(c){var l,f,p,d,b,v,m,y,w,g;try{v=t(c),m=u(v)}catch(E){for(g in s)s[g].onError(E);return r.onError(E),e}d=!1;try{w=s[m],w||(w=new He,s[m]=w,d=!0)}catch(E){for(g in s)s[g].onError(E);return r.onError(E),e}if(d){b=new Be(v,w,h),f=new Be(v,w);try{l=o(f)}catch(E){for(g in s)s[g].onError(E);return r.onError(E),e}r.onNext(b),y=new te,a.add(y);var D=function(){m in s&&(delete s[m],w.onCompleted()),a.remove(y)};y.setDisposable(l.take(1).subscribe(n,function(t){for(g in s)s[g].onError(t);r.onError(t)},function(){D()}))}try{p=i(c)}catch(E){for(g in s)s[g].onError(E);return r.onError(E),e}w.onNext(p)},function(t){for(var e in s)s[e].onError(t);r.onError(t)},function(){for(var t in s)s[t].onCompleted();r.onCompleted()})),h})},Ce.select=Ce.map=function(t,n){var r=this;return new Ve(function(i){var o=0;return r.subscribe(function(s){var u;try{u=t.call(n,s,o++,r)}catch(c){return i.onError(c),e}i.onNext(u)},i.onError.bind(i),i.onCompleted.bind(i))})},Ce.pluck=function(t){return this.select(function(e){return e[t]})},Ce.selectMany=Ce.flatMap=function(t,e){return e?this.selectMany(function(n){return t(n).select(function(t){return e(n,t)})}):"function"==typeof t?y.call(this,t):y.call(this,function(){return t})},Ce.selectSwitch=Ce.flatMapLatest=function(t,e){return this.select(t,e).switchLatest()},Ce.skip=function(t){if(0>t)throw Error(A);var e=this;return new Ve(function(n){var r=t;return e.subscribe(function(t){0>=r?n.onNext(t):r--},n.onError.bind(n),n.onCompleted.bind(n))})},Ce.skipWhile=function(t,n){var r=this;return new Ve(function(i){var o=0,s=!1;return r.subscribe(function(u){if(!s)try{s=!t.call(n,u,o++,r)}catch(c){return i.onError(c),e}s&&i.onNext(u)},i.onError.bind(i),i.onCompleted.bind(i))})},Ce.take=function(t,e){if(0>t)throw Error(A);if(0===t)return We(e);var n=this;return new Ve(function(e){var r=t;return n.subscribe(function(t){r>0&&(r--,e.onNext(t),0===r&&e.onCompleted())},e.onError.bind(e),e.onCompleted.bind(e))})},Ce.takeWhile=function(t,n){var r=this;return new Ve(function(i){var o=0,s=!0;return r.subscribe(function(u){if(s){try{s=t.call(n,u,o++,r)}catch(c){return i.onError(c),e}s?i.onNext(u):i.onCompleted()}},i.onError.bind(i),i.onCompleted.bind(i))})},Ce.where=Ce.filter=function(t,n){var r=this;return new Ve(function(i){var o=0;return r.subscribe(function(s){var u;try{u=t.call(n,s,o++,r)}catch(c){return i.onError(c),e}u&&i.onNext(s)},i.onError.bind(i),i.onCompleted.bind(i))})};var Ve=D.Internals.AnonymousObservable=function(t){function n(t){return t===e?t=$:"function"==typeof t&&(t=Z(t)),t}function r(i){function o(t){var e=new Fe(t);if(he.scheduleRequired())he.schedule(function(){try{e.setDisposable(n(i(e)))}catch(t){if(!e.fail(t))throw t}});else try{e.setDisposable(n(i(e)))}catch(r){if(!e.fail(r))throw r}return e}return this instanceof r?(t.call(this,o),e):new r(i)}return B(r,t),r}(Re),Fe=function(t){function e(e){t.call(this),this.observer=e,this.m=new te}B(e,t);var n=e.prototype;return n.next=function(t){var e=!1;try{this.observer.onNext(t),e=!0}catch(n){throw n}finally{e||this.dispose()}},n.error=function(t){try{this.observer.onError(t)}catch(e){throw e}finally{this.dispose()}},n.completed=function(){try{this.observer.onCompleted()}catch(t){throw t}finally{this.dispose()}},n.setDisposable=function(t){this.m.setDisposable(t)},n.getDisposable=function(){return this.m.getDisposable()},n.disposable=function(t){return arguments.length?this.getDisposable():setDisposable(t)},n.dispose=function(){t.prototype.dispose.call(this),this.m.dispose()},e}(Ae),Be=function(t){function e(t){return this.underlyingObservable.subscribe(t)}function n(n,r,i){t.call(this,e),this.key=n,this.underlyingObservable=i?new Ve(function(t){return new K(i.getDisposable(),r.subscribe(t))}):r}return B(n,t),n}(Re),Ue=function(t,e){this.subject=t,this.observer=e};Ue.prototype.dispose=function(){if(!this.subject.isDisposed&&null!==this.observer){var t=this.subject.observers.indexOf(this.observer);this.subject.observers.splice(t,1),this.observer=null}};var He=D.Subject=function(t){function e(t){return c.call(this),this.isStopped?this.exception?(t.onError(this.exception),$):(t.onCompleted(),$):(this.observers.push(t),new Ue(this,t))}function n(){t.call(this,e),this.isDisposed=!1,this.isStopped=!1,this.observers=[]}return B(n,t),U(n.prototype,De,{hasObservers:function(){return this.observers.length>0},onCompleted:function(){if(c.call(this),!this.isStopped){var t=this.observers.slice(0);this.isStopped=!0;for(var e=0,n=t.length;n>e;e++)t[e].onCompleted();this.observers=[]}},onError:function(t){if(c.call(this),!this.isStopped){var e=this.observers.slice(0);this.isStopped=!0,this.exception=t;for(var n=0,r=e.length;r>n;n++)e[n].onError(t);this.observers=[]}},onNext:function(t){if(c.call(this),!this.isStopped)for(var e=this.observers.slice(0),n=0,r=e.length;r>n;n++)e[n].onNext(t)},dispose:function(){this.isDisposed=!0,this.observers=null}}),n.create=function(t,e){return new Qe(t,e)},n}(Re);D.AsyncSubject=function(t){function e(t){if(c.call(this),!this.isStopped)return this.observers.push(t),new Ue(this,t);var e=this.exception,n=this.hasValue,r=this.value;return e?t.onError(e):n?(t.onNext(r),t.onCompleted()):t.onCompleted(),$}function n(){t.call(this,e),this.isDisposed=!1,this.isStopped=!1,this.value=null,this.hasValue=!1,this.observers=[],this.exception=null}return B(n,t),U(n.prototype,De,{hasObservers:function(){return this.observers.length>0},onCompleted:function(){var t,e,n;if(c.call(this),!this.isStopped){var r=this.observers.slice(0);this.isStopped=!0;var i=this.value,o=this.hasValue;if(o)for(e=0,n=r.length;n>e;e++)t=r[e],t.onNext(i),t.onCompleted();else for(e=0,n=r.length;n>e;e++)r[e].onCompleted();this.observers=[]}},onError:function(t){if(c.call(this),!this.isStopped){var e=this.observers.slice(0);this.isStopped=!0,this.exception=t;for(var n=0,r=e.length;r>n;n++)e[n].onError(t);this.observers=[]}},onNext:function(t){c.call(this),this.isStopped||(this.value=t,this.hasValue=!0)},dispose:function(){this.isDisposed=!0,this.observers=null,this.exception=null,this.value=null}}),n}(Re);var Qe=function(t){function e(t){return this.observable.subscribe(t)}function n(n,r){t.call(this,e),this.observer=n,this.observable=r}return B(n,t),U(n.prototype,De,{onCompleted:function(){this.observer.onCompleted()},onError:function(t){this.observer.onError(t)},onNext:function(t){this.observer.onNext(t)}}),n}(Re);return"function"==typeof define&&"object"==typeof define.amd&&define.amd?(t.Rx=D,define(function(){return D})):(w?"object"==typeof module&&module&&module.exports==w?module.exports=D:w=D:t.Rx=D,e)})(this);
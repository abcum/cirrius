/*jshint node:true*/
/* global require, module */
var EmberApp = require('ember-cli/lib/broccoli/ember-app');

module.exports = function(defaults) {

	if (process.env.EMBER_ENV === 'development') {

		var opts = {
			hinting: false,
			sourcemaps: {
				enabled: true,
			},
			minifyHTML: {
				enabled: false,
			},
			SRI: {
				enabled: false,
			},
			inlineContent: {
				browser: '',
				tracker: '',
			},
		};

	}

	if (process.env.EMBER_ENV === 'production') {

		var opts = {
			hinting: false,
			storeConfigInMeta: false,
			sourcemaps: {
				enabled: false,
			},
			minifyJS: {
				enabled: true,
			},
			minifyCSS: {
				enabled: true,
			},
			minifyHTML: {
				enabled: true,
			},
			fingerprint: {
				enabled: true,
			},
			SRI: {
				enabled: true,
			},
			inlineContent: {
				browser: 'public/browser.html',
				tracker: 'public/tracker.html',
				appcache: { content: ' manifest="/cache.appcache"' },
			},
		};

	}

	var app = new EmberApp(defaults, opts);

	// Use `app.import` to add additional libraries to the generated
	// output files.
	//
	// If you need to use different assets in different
	// environments, specify an object as the first parameter. That
	// object's keys should be the environment name and the values
	// should be the asset to use in that environment.
	//
	// If the library that you are including contains AMD or ES6
	// modules that you would like to import into your application
	// please specify an object with the list of modules as keys
	// along with the exports of each module as its value.

	return app.toTree();
};

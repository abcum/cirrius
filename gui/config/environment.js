/* jshint node: true */

module.exports = function(environment) {
	
	var ENV = {
	
		modulePrefix: 'gui',
		environment: environment,
		baseURL: '/',
		locationType: 'auto',

		editables: {
            includeTags: true,
            includeComments: true,
            includeBrackets: true,
            includeWhitespace: true,
            modes: [ "htmlmixed", "css",  "sass",  "javascript",  "markdown",  "handlebars", "sql" ],
        },
		
		// Set ember flags / options for the
        // ember runtime application
        // environment

        APP: { },

        // Set experimental ember features
        // to be used when using ember
        // canary builds

        EmberENV: { FEATURES: { } },

	};

	if (environment === 'development') {
		
		ENV.APP.BINDINGS = false;
        ENV.APP.LOG_RESOLVER = false;
        ENV.APP.LOG_TRANSITIONS = false;
        ENV.APP.LOG_VIEW_LOOKUPS = false;
        ENV.APP.LOG_ACTIVE_GENERATION = false;
        ENV.APP.LOG_TRANSITIONS_INTERNAL = false;

        ENV.APP.RAISE_ON_DEPRECATION = false;
        ENV.APP.LOG_STACKTRACE_ON_DEPRECATION = false;

	}

	if (environment === 'test') {
		
		// Testem prefers this...
		ENV.baseURL = '/';
		ENV.locationType = 'none';
		ENV.APP.rootElement = '#ember-testing';

		ENV.APP.BINDINGS = false;
        ENV.APP.LOG_RESOLVER = false;
        ENV.APP.LOG_TRANSITIONS = false;
        ENV.APP.LOG_VIEW_LOOKUPS = false;
        ENV.APP.LOG_ACTIVE_GENERATION = false;
        ENV.APP.LOG_TRANSITIONS_INTERNAL = false;

        ENV.APP.RAISE_ON_DEPRECATION = false;
        ENV.APP.LOG_STACKTRACE_ON_DEPRECATION = false;

	}

	if (environment === 'production') {

		ENV.APP.BINDINGS = false;
        ENV.APP.LOG_RESOLVER = false;
        ENV.APP.LOG_TRANSITIONS = false;
        ENV.APP.LOG_VIEW_LOOKUPS = false;
        ENV.APP.LOG_ACTIVE_GENERATION = false;
        ENV.APP.LOG_TRANSITIONS_INTERNAL = false;

        ENV.APP.RAISE_ON_DEPRECATION = false;
        ENV.APP.LOG_STACKTRACE_ON_DEPRECATION = false;

	}

	return ENV;

};

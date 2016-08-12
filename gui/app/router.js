import Ember from 'ember';
import config from './config/environment';

const Router = Ember.Router.extend({
		location: config.locationType
});

Router.map(function() {
		this.route('index', { path: '/' });

		this.route('project', { resetNamespace: true, path: '/:project_id' }, function() {
			this.route('file', { resetNamespace: true, path: '/:file_id' });
			this.route('analytics', { resetNamespace: true });
			this.route('settings', { resetNamespace: true });

			this.route('requests', { path: '/requests' }, function() {
				this.route('request', { path: '/:request_id' });
			});
		});

	this.route('login');
	this.route('reset');
	this.route('forgot');
	this.route('register');
	this.route('logout');

	this.route('request');
});

export default Router;
